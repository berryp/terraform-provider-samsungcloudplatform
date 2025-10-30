package certificate

type CreatePrdCertificateRequest struct {
	// Certificate Chain
	CertificateChain string
	// Certificate Name
	CertificateName string
	// Private Key
	PrivateKey string
	// Public Certificate
	PublicCertificate string
	// Recipient List
	Recipients []Recipient
	// 태그
	Tags map[string]interface{}
}

type CreateDevCertificateRequest struct {
	// Certificate Expiration Date
	CertificateExpirationDate string
	// Certificate Name
	CertificateName string
	// Certificate Start Date
	CertificateStartDate string
	// Common Name
	CommonName string
	// Organization Name
	OrganizationName string
	// Recipient List
	Recipients []Recipient
	// 태그
	Tags map[string]interface{}
}

type Recipient struct {
	// Email
	Email string
	// User Id
	UserId string
	// User Name
	UserName string
}
