package v2saml

// SAMLConnector represents SAML connectors which implement the HTTP POST binding.
//  RelayState is handled by the server.
//
// See: https://docs.oasis-open.org/security/saml/v2.0/saml-bindings-2.0-os.pdf
// "3.5 HTTP POST Binding"
type SAMLConnector interface {
	// POSTData returns an encoded SAML request and SSO URL for the server to
	// render a POST form with.
	//
	// POSTData should encode the provided request ID in the returned serialized
	// SAML request.
	POSTData(s Scopes, requestID string) (sooURL, samlRequest string, err error)

	// HandlePOST decodes, verifies, and maps attributes from the SAML response.
	// It passes the expected value of the "InResponseTo" response field, which
	// the connector must ensure matches the response value.
	//
	// See: https://www.oasis-open.org/committees/download.php/35711/sstc-saml-core-errata-2.0-wd-06-diff.pdf
	// "3.2.2 Complex Type StatusResponseType"
	HandlePOST(s Scopes, samlResponse, inResponseTo string) (identity Identity, err error)
}
