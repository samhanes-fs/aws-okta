package lib

// http://developer.okta.com/docs/api/resources/authn.html
type OktaUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type OktaStateToken struct {
	StateToken string `json:"stateToken"`
	PassCode   string `json:"passCode"`
}

type OktaU2fResponse struct {
	StateToken    string `json:"stateToken"`
	ClientData    string `json:"clientData"`
	SignatureData string `json:"signatureData"`
}

type OktaUserAuthn struct {
	StateToken   string                `json:"stateToken"`
	SessionToken string                `json:"sessionToken"`
	ExpiresAt    string                `json:"expiresAt"`
	Status       string                `json:"status"`
	Embedded     OktaUserAuthnEmbedded `json:"_embedded"`
	FactorResult string                `json:"factorResult"`
}

type OktaUserAuthnEmbedded struct {
	Factors []OktaUserAuthnFactor `json:"factors"`
	Factor  OktaUserAuthnFactor   `json:"factor"`
}

type OktaUserAuthnFactor struct {
	Id         string                      `json:"id"`
	FactorType string                      `json:"factorType"`
	Provider   string                      `json:"provider"`
	Profile    OktaUserAuthnFactorProfile  `json:"profile"`
	Embedded   OktaUserAuthnFactorEmbedded `json:"_embedded"`
}

type OktaUserAuthnFactorProfile struct {
	CredentialId string `json:"credentialId"`
	AppId        string `json:"appId"`
}

type OktaUserAuthnFactorEmbedded struct {
	Challenge OktaUserAuthnFactorEmbeddedChallenge `json:"challenge"`
	Verification OktaUserAuthnFactorEmbeddedVerification `json:"verification"`
}

type OktaUserAuthnFactorEmbeddedChallenge struct {
	Nonce          string `json:"nonce"`
	TimeoutSeconds int    `json:"timeoutSeconds"`
}

type OktaUserAuthnFactorEmbeddedVerification struct {
	Host         string                                       `json:"host"`
	Signature    string                                       `json:"signature"`
	FactorResult string                                       `json:"factorResult"`
	Links        OktaUserAuthnFactorEmbeddedVerificationLinks `json:"_links"`
}

type OktaUserAuthnFactorEmbeddedVerificationLinks struct {
	Complete OktaUserAuthnFactorEmbeddedVerificationLinksComplete `json:"complete"`
}

type OktaUserAuthnFactorEmbeddedVerificationLinksComplete struct {
	Href string `json:"href"`
}
