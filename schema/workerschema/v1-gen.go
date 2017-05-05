// Package workerschema provides access to the Dex API.
//
// See http://github.com/coreos/dex
//
// Usage example:
//
//   import "google.golang.org/api/workerschema/v1"
//   ...
//   workerschemaService, err := workerschema.New(oauthHttpClient)
package workerschema

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"google.golang.org/api/googleapi"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace

const apiId = "dex:v1"
const apiName = "workerschema"
const apiVersion = "v1"
const basePath = "$ENDPOINT/api/v1/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.RefreshClient = NewRefreshClientService(s)
	s.Users = NewUsersService(s)
	return s, nil
}

type Service struct {
	client   *http.Client
	BasePath string // API endpoint base URL

	RefreshClient *RefreshClientService

	Users *UsersService
}

func NewRefreshClientService(s *Service) *RefreshClientService {
	rs := &RefreshClientService{s: s}
	return rs
}

type RefreshClientService struct {
	s *Service
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	return rs
}

type UsersService struct {
	s *Service
}

type AddRemoteIdentityRequest struct {
	Identity *RemoteIdentity `json:"identity,omitempty"`
}

type AddRemoteIdentityResponse struct {
	Identity *RemoteIdentity `json:"identity,omitempty"`
}

type DeleteRemoteIdentityRequest struct {
	Identity *RemoteIdentity `json:"identity,omitempty"`
}

type DeleteRemoteIdentityResponse struct {
	Identity *RemoteIdentity `json:"identity,omitempty"`
}

type Error struct {
	Error string `json:"error,omitempty"`

	Error_description string `json:"error_description,omitempty"`
}

type GetRemoteIdentityResponse struct {
	Identity *RemoteIdentity `json:"identity,omitempty"`
}

type ListRemoteIdentityResponse struct {
	Identities []*RemoteIdentity `json:"identities,omitempty"`
}

type RefreshClient struct {
	ClientID string `json:"clientID,omitempty"`

	ClientName string `json:"clientName,omitempty"`

	ClientURI string `json:"clientURI,omitempty"`

	LogoURI string `json:"logoURI,omitempty"`
}

type RefreshClientList struct {
	Clients []*RefreshClient `json:"clients,omitempty"`
}

type RemoteIdentity struct {
	ConnectorID string `json:"connectorID,omitempty"`

	RemoteID string `json:"remoteID,omitempty"`
}

type RemoteIdentityDeleteResponse struct {
	Ok bool `json:"ok,omitempty"`
}

type ResendEmailInvitationRequest struct {
	RedirectURL string `json:"redirectURL,omitempty"`
}

type ResendEmailInvitationResponse struct {
	EmailSent bool `json:"emailSent,omitempty"`

	ResetPasswordLink string `json:"resetPasswordLink,omitempty"`
}

type User struct {
	Admin bool `json:"admin,omitempty"`

	CreatedAt string `json:"createdAt,omitempty"`

	Disabled bool `json:"disabled,omitempty"`

	DisplayName string `json:"displayName,omitempty"`

	Email string `json:"email,omitempty"`

	EmailVerified bool `json:"emailVerified,omitempty"`

	Id string `json:"id,omitempty"`

	Metadata string `json:"metadata,omitempty"`
}

type UserCreateRequest struct {
	RedirectURL string `json:"redirectURL,omitempty"`

	User *User `json:"user,omitempty"`
}

type UserCreateResponse struct {
	EmailSent bool `json:"emailSent,omitempty"`

	ResetPasswordLink string `json:"resetPasswordLink,omitempty"`

	User *User `json:"user,omitempty"`
}

type UserCreateResponseUser struct {
}

type UserDeleteResponse struct {
	Ok bool `json:"ok,omitempty"`
}

type UserDisableRequest struct {
	// Disable: If true, disable this user, if false, enable them. No error
	// is signaled if the user state doesn't change.
	Disable bool `json:"disable,omitempty"`
}

type UserDisableResponse struct {
	Ok bool `json:"ok,omitempty"`
}

type UserGetMetadataResponse struct {
	Metadata string `json:"metadata,omitempty"`
}

type UserResponse struct {
	User *User `json:"user,omitempty"`
}

type UserSetMetadataRequest struct {
	Metadata string `json:"metadata,omitempty"`
}

type UserSetMetadataResponse struct {
	Ok bool `json:"ok,omitempty"`
}

type UsersResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty"`

	Users []*User `json:"users,omitempty"`
}

// method id "dex.RefreshClient.List":

type RefreshClientListCall struct {
	s      *Service
	userid string
	opt_   map[string]interface{}
}

// List: List all clients that hold refresh tokens for the specified
// user.
func (r *RefreshClientService) List(userid string) *RefreshClientListCall {
	c := &RefreshClientListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userid = userid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RefreshClientListCall) Fields(s ...googleapi.Field) *RefreshClientListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *RefreshClientListCall) Do() (*RefreshClientList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "account/{userid}/refresh")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userid": c.userid,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *RefreshClientList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all clients that hold refresh tokens for the specified user.",
	//   "httpMethod": "GET",
	//   "id": "dex.RefreshClient.List",
	//   "parameterOrder": [
	//     "userid"
	//   ],
	//   "parameters": {
	//     "userid": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "account/{userid}/refresh",
	//   "response": {
	//     "$ref": "RefreshClientList"
	//   }
	// }

}

// method id "dex.RefreshClient.Revoke":

type RefreshClientRevokeCall struct {
	s        *Service
	userid   string
	clientid string
	opt_     map[string]interface{}
}

// Revoke: Revoke all refresh tokens issues to the client for the
// specified user.
func (r *RefreshClientService) Revoke(userid string, clientid string) *RefreshClientRevokeCall {
	c := &RefreshClientRevokeCall{s: r.s, opt_: make(map[string]interface{})}
	c.userid = userid
	c.clientid = clientid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RefreshClientRevokeCall) Fields(s ...googleapi.Field) *RefreshClientRevokeCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *RefreshClientRevokeCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "account/{userid}/refresh/{clientid}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userid":   c.userid,
		"clientid": c.clientid,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Revoke all refresh tokens issues to the client for the specified user.",
	//   "httpMethod": "DELETE",
	//   "id": "dex.RefreshClient.Revoke",
	//   "parameterOrder": [
	//     "userid",
	//     "clientid"
	//   ],
	//   "parameters": {
	//     "clientid": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userid": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "account/{userid}/refresh/{clientid}"
	// }

}

// method id "dex.User.AddRemoteIdentity":

type UsersAddRemoteIdentityCall struct {
	s                        *Service
	id                       string
	addremoteidentityrequest *AddRemoteIdentityRequest
	opt_                     map[string]interface{}
}

// AddRemoteIdentity: Add a remote identity for a user.
func (r *UsersService) AddRemoteIdentity(id string, addremoteidentityrequest *AddRemoteIdentityRequest) *UsersAddRemoteIdentityCall {
	c := &UsersAddRemoteIdentityCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.addremoteidentityrequest = addremoteidentityrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersAddRemoteIdentityCall) Fields(s ...googleapi.Field) *UsersAddRemoteIdentityCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersAddRemoteIdentityCall) Do() (*AddRemoteIdentityResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.addremoteidentityrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users/{id}/remote-identity")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *AddRemoteIdentityResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Add a remote identity for a user.",
	//   "httpMethod": "POST",
	//   "id": "dex.User.AddRemoteIdentity",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{id}/remote-identity",
	//   "request": {
	//     "$ref": "AddRemoteIdentityRequest"
	//   },
	//   "response": {
	//     "$ref": "AddRemoteIdentityResponse"
	//   }
	// }

}

// method id "dex.User.Create":

type UsersCreateCall struct {
	s                 *Service
	usercreaterequest *UserCreateRequest
	opt_              map[string]interface{}
}

// Create: Create a new User.
func (r *UsersService) Create(usercreaterequest *UserCreateRequest) *UsersCreateCall {
	c := &UsersCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.usercreaterequest = usercreaterequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersCreateCall) Fields(s ...googleapi.Field) *UsersCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersCreateCall) Do() (*UserCreateResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.usercreaterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *UserCreateResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a new User.",
	//   "httpMethod": "POST",
	//   "id": "dex.User.Create",
	//   "path": "users",
	//   "request": {
	//     "$ref": "UserCreateRequest"
	//   },
	//   "response": {
	//     "$ref": "UserCreateResponse"
	//   }
	// }

}

// method id "dex.User.DeleteRemoteIdentity":

type UsersDeleteRemoteIdentityCall struct {
	s                           *Service
	id                          string
	deleteremoteidentityrequest *DeleteRemoteIdentityRequest
	opt_                        map[string]interface{}
}

// DeleteRemoteIdentity: Delete a single RemoteIdentity object by user
// and supplied remote identity object.
func (r *UsersService) DeleteRemoteIdentity(id string, deleteremoteidentityrequest *DeleteRemoteIdentityRequest) *UsersDeleteRemoteIdentityCall {
	c := &UsersDeleteRemoteIdentityCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.deleteremoteidentityrequest = deleteremoteidentityrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDeleteRemoteIdentityCall) Fields(s ...googleapi.Field) *UsersDeleteRemoteIdentityCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersDeleteRemoteIdentityCall) Do() (*DeleteRemoteIdentityResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.deleteremoteidentityrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users/{id}/remote-identity")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *DeleteRemoteIdentityResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Delete a single RemoteIdentity object by user and supplied remote identity object.",
	//   "httpMethod": "DELETE",
	//   "id": "dex.User.DeleteRemoteIdentity",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{id}/remote-identity",
	//   "request": {
	//     "$ref": "DeleteRemoteIdentityRequest"
	//   },
	//   "response": {
	//     "$ref": "DeleteRemoteIdentityResponse"
	//   }
	// }

}

// method id "dex.User.Disable":

type UsersDisableCall struct {
	s                  *Service
	id                 string
	userdisablerequest *UserDisableRequest
	opt_               map[string]interface{}
}

// Disable: Enable or disable a user.
func (r *UsersService) Disable(id string, userdisablerequest *UserDisableRequest) *UsersDisableCall {
	c := &UsersDisableCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.userdisablerequest = userdisablerequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDisableCall) Fields(s ...googleapi.Field) *UsersDisableCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersDisableCall) Do() (*UserDisableResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.userdisablerequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users/{id}/disable")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *UserDisableResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Enable or disable a user.",
	//   "httpMethod": "POST",
	//   "id": "dex.User.Disable",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{id}/disable",
	//   "request": {
	//     "$ref": "UserDisableRequest"
	//   },
	//   "response": {
	//     "$ref": "UserDisableResponse"
	//   }
	// }

}

// method id "dex.User.Get":

type UsersGetCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Get: Get a single User object by id.
func (r *UsersService) Get(id string) *UsersGetCall {
	c := &UsersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersGetCall) Fields(s ...googleapi.Field) *UsersGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersGetCall) Do() (*UserResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *UserResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a single User object by id.",
	//   "httpMethod": "GET",
	//   "id": "dex.User.Get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{id}",
	//   "response": {
	//     "$ref": "UserResponse"
	//   }
	// }

}

// method id "dex.User.GetRemoteIdentity":

type UsersGetRemoteIdentityCall struct {
	s           *Service
	id          string
	connectorid string
	opt_        map[string]interface{}
}

// GetRemoteIdentity: Get a single RemoteIdentity object by user and
// remote ids.
func (r *UsersService) GetRemoteIdentity(id string, connectorid string) *UsersGetRemoteIdentityCall {
	c := &UsersGetRemoteIdentityCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.connectorid = connectorid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersGetRemoteIdentityCall) Fields(s ...googleapi.Field) *UsersGetRemoteIdentityCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersGetRemoteIdentityCall) Do() (*GetRemoteIdentityResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users/{id}/remote-identity/{connectorid}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id":          c.id,
		"connectorid": c.connectorid,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *GetRemoteIdentityResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a single RemoteIdentity object by user and remote ids.",
	//   "httpMethod": "GET",
	//   "id": "dex.User.GetRemoteIdentity",
	//   "parameterOrder": [
	//     "id",
	//     "connectorid"
	//   ],
	//   "parameters": {
	//     "connectorid": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{id}/remote-identity/{connectorid}",
	//   "response": {
	//     "$ref": "GetRemoteIdentityResponse"
	//   }
	// }

}

// method id "dex.User.List":

type UsersListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Retrieve a page of User objects.
func (r *UsersService) List() *UsersListCall {
	c := &UsersListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// MaxResults sets the optional parameter "maxResults":
func (c *UsersListCall) MaxResults(maxResults int64) *UsersListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// NextPageToken sets the optional parameter "nextPageToken":
func (c *UsersListCall) NextPageToken(nextPageToken string) *UsersListCall {
	c.opt_["nextPageToken"] = nextPageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersListCall) Fields(s ...googleapi.Field) *UsersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersListCall) Do() (*UsersResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["nextPageToken"]; ok {
		params.Set("nextPageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *UsersResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve a page of User objects.",
	//   "httpMethod": "GET",
	//   "id": "dex.User.List",
	//   "parameters": {
	//     "maxResults": {
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "nextPageToken": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "users",
	//   "response": {
	//     "$ref": "UsersResponse"
	//   }
	// }

}

// method id "dex.User.ListRemoteIdentity":

type UsersListRemoteIdentityCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// ListRemoteIdentity: Retrieve a page of RemoteIdentity objects.
func (r *UsersService) ListRemoteIdentity(id string) *UsersListRemoteIdentityCall {
	c := &UsersListRemoteIdentityCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersListRemoteIdentityCall) Fields(s ...googleapi.Field) *UsersListRemoteIdentityCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersListRemoteIdentityCall) Do() (*ListRemoteIdentityResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users/{id}/remote-identity")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListRemoteIdentityResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve a page of RemoteIdentity objects.",
	//   "httpMethod": "GET",
	//   "id": "dex.User.ListRemoteIdentity",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{id}/remote-identity",
	//   "response": {
	//     "$ref": "ListRemoteIdentityResponse"
	//   }
	// }

}

// method id "dex.User.ResendEmailInvitation":

type UsersResendEmailInvitationCall struct {
	s                            *Service
	id                           string
	resendemailinvitationrequest *ResendEmailInvitationRequest
	opt_                         map[string]interface{}
}

// ResendEmailInvitation: Resend invitation email to an existing user
// with unverified email.
func (r *UsersService) ResendEmailInvitation(id string, resendemailinvitationrequest *ResendEmailInvitationRequest) *UsersResendEmailInvitationCall {
	c := &UsersResendEmailInvitationCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.resendemailinvitationrequest = resendemailinvitationrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersResendEmailInvitationCall) Fields(s ...googleapi.Field) *UsersResendEmailInvitationCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersResendEmailInvitationCall) Do() (*ResendEmailInvitationResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.resendemailinvitationrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users/{id}/resend-invitation")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ResendEmailInvitationResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Resend invitation email to an existing user with unverified email.",
	//   "httpMethod": "POST",
	//   "id": "dex.User.ResendEmailInvitation",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{id}/resend-invitation",
	//   "request": {
	//     "$ref": "ResendEmailInvitationRequest"
	//   },
	//   "response": {
	//     "$ref": "ResendEmailInvitationResponse"
	//   }
	// }

}
