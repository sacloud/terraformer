// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acm

type CertificateStatus string

// Enum values for CertificateStatus
const (
	CertificateStatusPendingValidation  CertificateStatus = "PENDING_VALIDATION"
	CertificateStatusIssued             CertificateStatus = "ISSUED"
	CertificateStatusInactive           CertificateStatus = "INACTIVE"
	CertificateStatusExpired            CertificateStatus = "EXPIRED"
	CertificateStatusValidationTimedOut CertificateStatus = "VALIDATION_TIMED_OUT"
	CertificateStatusRevoked            CertificateStatus = "REVOKED"
	CertificateStatusFailed             CertificateStatus = "FAILED"
)

func (enum CertificateStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CertificateStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CertificateTransparencyLoggingPreference string

// Enum values for CertificateTransparencyLoggingPreference
const (
	CertificateTransparencyLoggingPreferenceEnabled  CertificateTransparencyLoggingPreference = "ENABLED"
	CertificateTransparencyLoggingPreferenceDisabled CertificateTransparencyLoggingPreference = "DISABLED"
)

func (enum CertificateTransparencyLoggingPreference) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CertificateTransparencyLoggingPreference) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CertificateType string

// Enum values for CertificateType
const (
	CertificateTypeImported     CertificateType = "IMPORTED"
	CertificateTypeAmazonIssued CertificateType = "AMAZON_ISSUED"
	CertificateTypePrivate      CertificateType = "PRIVATE"
)

func (enum CertificateType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CertificateType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DomainStatus string

// Enum values for DomainStatus
const (
	DomainStatusPendingValidation DomainStatus = "PENDING_VALIDATION"
	DomainStatusSuccess           DomainStatus = "SUCCESS"
	DomainStatusFailed            DomainStatus = "FAILED"
)

func (enum DomainStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DomainStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ExtendedKeyUsageName string

// Enum values for ExtendedKeyUsageName
const (
	ExtendedKeyUsageNameTlsWebServerAuthentication ExtendedKeyUsageName = "TLS_WEB_SERVER_AUTHENTICATION"
	ExtendedKeyUsageNameTlsWebClientAuthentication ExtendedKeyUsageName = "TLS_WEB_CLIENT_AUTHENTICATION"
	ExtendedKeyUsageNameCodeSigning                ExtendedKeyUsageName = "CODE_SIGNING"
	ExtendedKeyUsageNameEmailProtection            ExtendedKeyUsageName = "EMAIL_PROTECTION"
	ExtendedKeyUsageNameTimeStamping               ExtendedKeyUsageName = "TIME_STAMPING"
	ExtendedKeyUsageNameOcspSigning                ExtendedKeyUsageName = "OCSP_SIGNING"
	ExtendedKeyUsageNameIpsecEndSystem             ExtendedKeyUsageName = "IPSEC_END_SYSTEM"
	ExtendedKeyUsageNameIpsecTunnel                ExtendedKeyUsageName = "IPSEC_TUNNEL"
	ExtendedKeyUsageNameIpsecUser                  ExtendedKeyUsageName = "IPSEC_USER"
	ExtendedKeyUsageNameAny                        ExtendedKeyUsageName = "ANY"
	ExtendedKeyUsageNameNone                       ExtendedKeyUsageName = "NONE"
	ExtendedKeyUsageNameCustom                     ExtendedKeyUsageName = "CUSTOM"
)

func (enum ExtendedKeyUsageName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ExtendedKeyUsageName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FailureReason string

// Enum values for FailureReason
const (
	FailureReasonNoAvailableContacts            FailureReason = "NO_AVAILABLE_CONTACTS"
	FailureReasonAdditionalVerificationRequired FailureReason = "ADDITIONAL_VERIFICATION_REQUIRED"
	FailureReasonDomainNotAllowed               FailureReason = "DOMAIN_NOT_ALLOWED"
	FailureReasonInvalidPublicDomain            FailureReason = "INVALID_PUBLIC_DOMAIN"
	FailureReasonDomainValidationDenied         FailureReason = "DOMAIN_VALIDATION_DENIED"
	FailureReasonCaaError                       FailureReason = "CAA_ERROR"
	FailureReasonPcaLimitExceeded               FailureReason = "PCA_LIMIT_EXCEEDED"
	FailureReasonPcaInvalidArn                  FailureReason = "PCA_INVALID_ARN"
	FailureReasonPcaInvalidState                FailureReason = "PCA_INVALID_STATE"
	FailureReasonPcaRequestFailed               FailureReason = "PCA_REQUEST_FAILED"
	FailureReasonPcaResourceNotFound            FailureReason = "PCA_RESOURCE_NOT_FOUND"
	FailureReasonPcaInvalidArgs                 FailureReason = "PCA_INVALID_ARGS"
	FailureReasonPcaInvalidDuration             FailureReason = "PCA_INVALID_DURATION"
	FailureReasonPcaAccessDenied                FailureReason = "PCA_ACCESS_DENIED"
	FailureReasonOther                          FailureReason = "OTHER"
)

func (enum FailureReason) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FailureReason) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type KeyAlgorithm string

// Enum values for KeyAlgorithm
const (
	KeyAlgorithmRsa2048      KeyAlgorithm = "RSA_2048"
	KeyAlgorithmRsa1024      KeyAlgorithm = "RSA_1024"
	KeyAlgorithmRsa4096      KeyAlgorithm = "RSA_4096"
	KeyAlgorithmEcPrime256v1 KeyAlgorithm = "EC_prime256v1"
	KeyAlgorithmEcSecp384r1  KeyAlgorithm = "EC_secp384r1"
	KeyAlgorithmEcSecp521r1  KeyAlgorithm = "EC_secp521r1"
)

func (enum KeyAlgorithm) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum KeyAlgorithm) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type KeyUsageName string

// Enum values for KeyUsageName
const (
	KeyUsageNameDigitalSignature   KeyUsageName = "DIGITAL_SIGNATURE"
	KeyUsageNameNonRepudiation     KeyUsageName = "NON_REPUDIATION"
	KeyUsageNameKeyEncipherment    KeyUsageName = "KEY_ENCIPHERMENT"
	KeyUsageNameDataEncipherment   KeyUsageName = "DATA_ENCIPHERMENT"
	KeyUsageNameKeyAgreement       KeyUsageName = "KEY_AGREEMENT"
	KeyUsageNameCertificateSigning KeyUsageName = "CERTIFICATE_SIGNING"
	KeyUsageNameCrlSigning         KeyUsageName = "CRL_SIGNING"
	KeyUsageNameEncipherOnly       KeyUsageName = "ENCIPHER_ONLY"
	KeyUsageNameDecipherOnly       KeyUsageName = "DECIPHER_ONLY"
	KeyUsageNameAny                KeyUsageName = "ANY"
	KeyUsageNameCustom             KeyUsageName = "CUSTOM"
)

func (enum KeyUsageName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum KeyUsageName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RecordType string

// Enum values for RecordType
const (
	RecordTypeCname RecordType = "CNAME"
)

func (enum RecordType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RecordType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RenewalEligibility string

// Enum values for RenewalEligibility
const (
	RenewalEligibilityEligible   RenewalEligibility = "ELIGIBLE"
	RenewalEligibilityIneligible RenewalEligibility = "INELIGIBLE"
)

func (enum RenewalEligibility) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RenewalEligibility) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RenewalStatus string

// Enum values for RenewalStatus
const (
	RenewalStatusPendingAutoRenewal RenewalStatus = "PENDING_AUTO_RENEWAL"
	RenewalStatusPendingValidation  RenewalStatus = "PENDING_VALIDATION"
	RenewalStatusSuccess            RenewalStatus = "SUCCESS"
	RenewalStatusFailed             RenewalStatus = "FAILED"
)

func (enum RenewalStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RenewalStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RevocationReason string

// Enum values for RevocationReason
const (
	RevocationReasonUnspecified          RevocationReason = "UNSPECIFIED"
	RevocationReasonKeyCompromise        RevocationReason = "KEY_COMPROMISE"
	RevocationReasonCaCompromise         RevocationReason = "CA_COMPROMISE"
	RevocationReasonAffiliationChanged   RevocationReason = "AFFILIATION_CHANGED"
	RevocationReasonSuperceded           RevocationReason = "SUPERCEDED"
	RevocationReasonCessationOfOperation RevocationReason = "CESSATION_OF_OPERATION"
	RevocationReasonCertificateHold      RevocationReason = "CERTIFICATE_HOLD"
	RevocationReasonRemoveFromCrl        RevocationReason = "REMOVE_FROM_CRL"
	RevocationReasonPrivilegeWithdrawn   RevocationReason = "PRIVILEGE_WITHDRAWN"
	RevocationReasonAACompromise         RevocationReason = "A_A_COMPROMISE"
)

func (enum RevocationReason) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RevocationReason) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ValidationMethod string

// Enum values for ValidationMethod
const (
	ValidationMethodEmail ValidationMethod = "EMAIL"
	ValidationMethodDns   ValidationMethod = "DNS"
)

func (enum ValidationMethod) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ValidationMethod) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
