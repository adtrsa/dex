package connector

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/coreos/dex/connector/v2saml"
	phttp "github.com/coreos/dex/pkg/http"
	"github.com/coreos/go-oidc/oauth2"
	"github.com/coreos/go-oidc/oidc"
)

const (
	SAMLConnectorType = "saml"
)

type SAMLConnectorConfig struct {
	ID       string        `json:"id"`
	V2Config v2saml.Config `json:"v2config"`
}

type SAMLConnector struct {
	id          string
	namespace   url.URL
	loginFunc   oidc.LoginFunc
	v2connector v2saml.SAMLConnector
}

func init() {
	RegisterConnectorConfigType(SAMLConnectorType,
		func() ConnectorConfig { return &SAMLConnectorConfig{} })
}

func (cfg *SAMLConnectorConfig) ConnectorID() string {
	return cfg.ID
}

func (cfg *SAMLConnectorConfig) ConnectorType() string {
	return SAMLConnectorType
}

func (cfg *SAMLConnectorConfig) Connector(ns url.URL,
	lf oidc.LoginFunc,
	tpls *template.Template) (Connector, error) {

	idpc := &SAMLConnector{
		id:        cfg.ID,
		namespace: ns,
		loginFunc: lf,
	}

	var err error
	idpc.v2connector, err = cfg.V2Config.Open(nil)
	return idpc, err
}

func (c *SAMLConnector) ID() string {
	return c.id
}

func (c *SAMLConnector) Healthy() error {
	return nil
}

func (c *SAMLConnector) LoginURL(sessionKey, prompt string) (string, error) {
	q := url.Values{}
	q.Set("session_key", sessionKey)
	q.Set("prompt", prompt)
	enc := q.Encode()
	return path.Join(c.namespace.Path, "login") + "?" + enc, nil
}

func (c *SAMLConnector) Sync() chan struct{} {
	return make(chan struct{})
}

func (c *SAMLConnector) TrustedEmailProvider() bool {
	return false
}

func (c *SAMLConnector) Identity(email string) (*oidc.Identity, error) {
	//TODO: [adtrsa] Erm.... is this a fair assumption?
	//The remote identity format will depend on what's beyond the SAML
	//connection (LDAP etc.).
	id := &oidc.Identity{Email: email, ID: email}
	return id, nil
}

func (c *SAMLConnector) Handler(errorURL url.URL) http.Handler {
	route := path.Join(c.namespace.Path, "/login")
	return c.handleLogin(c.loginFunc, route, errorURL)
}

func (c *SAMLConnector) handleLogin(lf oidc.LoginFunc,
	localErrorPath string, errorURL url.URL) http.HandlerFunc {
	handleGET := func(w http.ResponseWriter, r *http.Request, errMsg string) {
		q := r.URL.Query()
		sessionKey := q.Get("session_key")

		var scopes v2saml.Scopes

		// build authentication request POST data.
		action, value, err := c.v2connector.POSTData(scopes, sessionKey)

		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("SAMLConnector.handleGET v2 POSTData failed")
			return
		}

		// TODO(adtrsa) May need to add other things to RelayState apart from session key to auth
		// TODO(ericchiang): Don't inline this.
		fmt.Fprintf(w, `<!DOCTYPE html>
			  <html lang="en">
			  <head>
			    <meta http-equiv="content-type" content="text/html; charset=utf-8">
			    <title>SAML login</title>
			  </head>
			  <body>
			    <form method="post" action="%s" >
				    <input type="hidden" name="SAMLRequest" value="%s" />
				    <input type="hidden" name="RelayState" value="%s" />
			    </form>
				<script>
				    document.forms[0].submit();
				</script>
			  </body>
			  </html>`, action, value, sessionKey)

	} //handleGET

	handlePOST := func(w http.ResponseWriter,
		r *http.Request, errMsg string) {
		q := r.URL.Query()
		// ensure POST was done with correct content type
		contenTypeHeader := r.Header.Get("Content-Type")
		expectedContentTypeHeader := "application/x-www-form-urlencoded"

		if contenTypeHeader != expectedContentTypeHeader {
			log.WithFields(
				log.Fields{
					"actual":   contenTypeHeader,
					"expected": expectedContentTypeHeader,
				}).Error(
				"SAMLConnector.handlePOST Wrong Content-Type.")
			return
		}

		err := r.ParseForm()
		if err != nil {
			log.Error("SAMLConnector.handlePOST Cannot parse form.")
			return
		}

		// recover session key from relay state
		sessionKey := r.PostFormValue("RelayState")

		if sessionKey == "" {
			log.Error("SAMLConnector.handlePOST Empty session key/RelayState.")
			return
		}

		samlResponse := r.PostFormValue("SAMLResponse")

		if samlResponse == "" {
			log.Error("SAMLConnector.handlePOST Empty SAMLResponse.")
			return
		}

		var s v2saml.Scopes
		ident2, err := c.v2connector.HandlePOST(s, samlResponse,
			sessionKey)

		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("SAMLConnector.handlePOST Unable to handle response.")
			return
		}

		ident, err := c.Identity(ident2.Email)

		if err != nil {
			log.WithFields(
				log.Fields{
					"email": ident2.Email,
					"err":   err,
				}).Error("SAMLConnector.handlePOST Identity retrieval failure.")
			return
		}

		redirectURL, err := lf(*ident, sessionKey)

		if err != nil {
			log.Errorf("Unable to log in %#v: %v", ident, err)
			q.Set("error", oauth2.ErrorAccessDenied)
			q.Set("error_description", "login failed")
			redirectError(w, errorURL, q)
		}

		w.Header().Set("Location", redirectURL)
		w.WriteHeader(http.StatusFound)
	} //handlePOST

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			handlePOST(w, r, "")
		case "GET":
			handleGET(w, r, "")
		default:
			w.Header().Set("Allow", "GET, POST")
			phttp.WriteError(w, http.StatusMethodNotAllowed,
				"GET and POST only acceptable methods")
		}
	}
}