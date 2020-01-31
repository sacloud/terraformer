// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateMultipartUploadInput struct {
	_ struct{} `type:"structure"`

	// The canned ACL to apply to the object.
	ACL ObjectCannedACL `location:"header" locationName:"x-amz-acl" type:"string" enum:"true"`

	// The name of the bucket to which to initiate the upload
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`

	// Specifies presentational information for the object.
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`

	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field.
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`

	// The language the content is in.
	ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`

	// A standard MIME type describing the format of the object data.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// The date and time at which the object is no longer cacheable.
	Expires *time.Time `location:"header" locationName:"Expires" type:"timestamp"`

	// Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
	GrantFullControl *string `location:"header" locationName:"x-amz-grant-full-control" type:"string"`

	// Allows grantee to read the object data and its metadata.
	GrantRead *string `location:"header" locationName:"x-amz-grant-read" type:"string"`

	// Allows grantee to read the object ACL.
	GrantReadACP *string `location:"header" locationName:"x-amz-grant-read-acp" type:"string"`

	// Allows grantee to write the ACL for the applicable object.
	GrantWriteACP *string `location:"header" locationName:"x-amz-grant-write-acp" type:"string"`

	// Object key for which the multipart upload is to be initiated.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// A map of metadata to store with the object in S3.
	Metadata map[string]string `location:"headers" locationName:"x-amz-meta-" type:"map"`

	// Specifies whether you want to apply a Legal Hold to the uploaded object.
	ObjectLockLegalHoldStatus ObjectLockLegalHoldStatus `location:"header" locationName:"x-amz-object-lock-legal-hold" type:"string" enum:"true"`

	// Specifies the Object Lock mode that you want to apply to the uploaded object.
	ObjectLockMode ObjectLockMode `location:"header" locationName:"x-amz-object-lock-mode" type:"string" enum:"true"`

	// Specifies the date and time when you want the Object Lock to expire.
	ObjectLockRetainUntilDate *time.Time `location:"header" locationName:"x-amz-object-lock-retain-until-date" type:"timestamp" timestampFormat:"iso8601"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// Specifies the algorithm to use to when encrypting the object (e.g., AES256).
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting
	// data. This value is used to store the object and then it is discarded; Amazon
	// does not store the encryption key. The key must be appropriate for use with
	// the algorithm specified in the x-amz-server-side​-encryption​-customer-algorithm
	// header.
	SSECustomerKey *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key" type:"string" sensitive:"true"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure the encryption
	// key was transmitted without error.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// Specifies the AWS KMS Encryption Context to use for object encryption. The
	// value of this header is a base64-encoded UTF-8 string holding JSON with the
	// encryption context key-value pairs.
	SSEKMSEncryptionContext *string `location:"header" locationName:"x-amz-server-side-encryption-context" type:"string" sensitive:"true"`

	// Specifies the AWS KMS key ID to use for object encryption. All GET and PUT
	// requests for an object protected by AWS KMS will fail if not made via SSL
	// or using SigV4. Documentation on configuring any of the officially supported
	// AWS SDKs and CLI can be found at http://docs.aws.amazon.com/AmazonS3/latest/dev/UsingAWSSDK.html#specify-signature-version
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`

	// The type of storage to use for the object. Defaults to 'STANDARD'.
	StorageClass StorageClass `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"true"`

	// The tag-set for the object. The tag-set must be encoded as URL Query parameters
	Tagging *string `location:"header" locationName:"x-amz-tagging" type:"string"`

	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata.
	WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
}

// String returns the string representation
func (s CreateMultipartUploadInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMultipartUploadInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMultipartUploadInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *CreateMultipartUploadInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

func (s *CreateMultipartUploadInput) getSSECustomerKey() (v string) {
	if s.SSECustomerKey == nil {
		return v
	}
	return *s.SSECustomerKey
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMultipartUploadInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.ACL) > 0 {
		v := s.ACL

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-acl", v, metadata)
	}
	if s.CacheControl != nil {
		v := *s.CacheControl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Cache-Control", protocol.StringValue(v), metadata)
	}
	if s.ContentDisposition != nil {
		v := *s.ContentDisposition

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Disposition", protocol.StringValue(v), metadata)
	}
	if s.ContentEncoding != nil {
		v := *s.ContentEncoding

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Encoding", protocol.StringValue(v), metadata)
	}
	if s.ContentLanguage != nil {
		v := *s.ContentLanguage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Language", protocol.StringValue(v), metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue(v), metadata)
	}
	if s.Expires != nil {
		v := *s.Expires

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Expires",
			protocol.TimeValue{V: v, Format: protocol.RFC822TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.GrantFullControl != nil {
		v := *s.GrantFullControl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-full-control", protocol.StringValue(v), metadata)
	}
	if s.GrantRead != nil {
		v := *s.GrantRead

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-read", protocol.StringValue(v), metadata)
	}
	if s.GrantReadACP != nil {
		v := *s.GrantReadACP

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-read-acp", protocol.StringValue(v), metadata)
	}
	if s.GrantWriteACP != nil {
		v := *s.GrantWriteACP

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-write-acp", protocol.StringValue(v), metadata)
	}
	if len(s.ObjectLockLegalHoldStatus) > 0 {
		v := s.ObjectLockLegalHoldStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-legal-hold", v, metadata)
	}
	if len(s.ObjectLockMode) > 0 {
		v := s.ObjectLockMode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-mode", v, metadata)
	}
	if s.ObjectLockRetainUntilDate != nil {
		v := *s.ObjectLockRetainUntilDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-retain-until-date",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: false}, metadata)
	}
	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.SSECustomerAlgorithm != nil {
		v := *s.SSECustomerAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-algorithm", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKey != nil {
		v := *s.SSECustomerKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKeyMD5 != nil {
		v := *s.SSECustomerKeyMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key-MD5", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSEncryptionContext != nil {
		v := *s.SSEKMSEncryptionContext

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-context", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSKeyId != nil {
		v := *s.SSEKMSKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-aws-kms-key-id", protocol.StringValue(v), metadata)
	}
	if len(s.ServerSideEncryption) > 0 {
		v := s.ServerSideEncryption

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption", v, metadata)
	}
	if len(s.StorageClass) > 0 {
		v := s.StorageClass

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-storage-class", v, metadata)
	}
	if s.Tagging != nil {
		v := *s.Tagging

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-tagging", protocol.StringValue(v), metadata)
	}
	if s.WebsiteRedirectLocation != nil {
		v := *s.WebsiteRedirectLocation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-website-redirect-location", protocol.StringValue(v), metadata)
	}
	if s.Metadata != nil {
		v := s.Metadata

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.HeadersTarget, "x-amz-meta-", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.StringValue(v1))
		}
		ms0.End()

	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	return nil
}

type CreateMultipartUploadOutput struct {
	_ struct{} `type:"structure"`

	// If the bucket has a lifecycle rule configured with an action to abort incomplete
	// multipart uploads and the prefix in the lifecycle rule matches the object
	// name in the request, the response includes this header. The header indicates
	// when the initiated multipart upload becomes eligible for an abort operation.
	// For more information, see Aborting Incomplete Multipart Uploads Using a Bucket
	// Lifecycle Policy (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config).
	//
	// The response also includes the x-amz-abort-rule-id header that provides the
	// ID of the lifecycle configuration rule that defines this action.
	AbortDate *time.Time `location:"header" locationName:"x-amz-abort-date" type:"timestamp"`

	// This header is returned along with the x-amz-abort-date header. It identifies
	// the applicable lifecycle configuration rule that defines the action to abort
	// incomplete multipart uploads.
	AbortRuleId *string `location:"header" locationName:"x-amz-abort-rule-id" type:"string"`

	// Name of the bucket to which the multipart upload was initiated.
	Bucket *string `locationName:"Bucket" type:"string"`

	// Object key for which the multipart upload was initiated.
	Key *string `min:"1" type:"string"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// If present, specifies the AWS KMS Encryption Context to use for object encryption.
	// The value of this header is a base64-encoded UTF-8 string holding JSON with
	// the encryption context key-value pairs.
	SSEKMSEncryptionContext *string `location:"header" locationName:"x-amz-server-side-encryption-context" type:"string" sensitive:"true"`

	// If present, specifies the ID of the AWS Key Management Service (KMS) customer
	// master key (CMK) that was used for the object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`

	// ID for the initiated multipart upload.
	UploadId *string `type:"string"`
}

// String returns the string representation
func (s CreateMultipartUploadOutput) String() string {
	return awsutil.Prettify(s)
}

func (s *CreateMultipartUploadOutput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMultipartUploadOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.UploadId != nil {
		v := *s.UploadId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UploadId", protocol.StringValue(v), metadata)
	}
	if s.AbortDate != nil {
		v := *s.AbortDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-abort-date",
			protocol.TimeValue{V: v, Format: protocol.RFC822TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.AbortRuleId != nil {
		v := *s.AbortRuleId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-abort-rule-id", protocol.StringValue(v), metadata)
	}
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	if s.SSECustomerAlgorithm != nil {
		v := *s.SSECustomerAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-algorithm", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKeyMD5 != nil {
		v := *s.SSECustomerKeyMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key-MD5", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSEncryptionContext != nil {
		v := *s.SSEKMSEncryptionContext

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-context", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSKeyId != nil {
		v := *s.SSEKMSKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-aws-kms-key-id", protocol.StringValue(v), metadata)
	}
	if len(s.ServerSideEncryption) > 0 {
		v := s.ServerSideEncryption

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption", v, metadata)
	}
	return nil
}

const opCreateMultipartUpload = "CreateMultipartUpload"

// CreateMultipartUploadRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// This operation initiates a multipart upload and returns an upload ID. This
// upload ID is used to associate all of the parts in the specific multipart
// upload. You specify this upload ID in each of your subsequent upload part
// requests (see UploadPart). You also include this upload ID in the final request
// to either complete or abort the multipart upload request.
//
// For more information about multipart uploads, see Multipart Upload Overview
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html).
//
// If you have configured a lifecycle rule to abort incomplete multipart uploads,
// the upload must complete within the number of days specified in the bucket
// lifecycle configuration. Otherwise, the incomplete multipart upload becomes
// eligible for an abort operation and Amazon S3 aborts the multipart upload.
// For more information, see Aborting Incomplete Multipart Uploads Using a Bucket
// Lifecycle Policy (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config).
//
// For information about the permissions required to use the multipart upload
// API, see Multipart Upload API and Permissions (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html).
//
// For request signing, multipart upload is just a series of regular requests.
// You initiate a multipart upload, send one or more requests to upload parts,
// and then complete the multipart upload process. You sign each request individually.
// There is nothing special about signing multipart upload requests. For more
// information about signing, see Authenticating Requests (AWS Signature Version
// 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html).
//
// After you initiate a multipart upload and upload one or more parts, to stop
// being charged for storing the uploaded parts, you must either complete or
// abort the multipart upload. Amazon S3 frees up the space used to store the
// parts and stop charging you for storing them only after you either complete
// or abort a multipart upload.
//
// You can optionally request server-side encryption. For server-side encryption,
// Amazon S3 encrypts your data as it writes it to disks in its data centers
// and decrypts it when you access it. You can provide your own encryption key,
// or use AWS Key Management Service (AWS KMS) customer master keys (CMKs) or
// Amazon S3-managed encryption keys. If you choose to provide your own encryption
// key, the request headers you provide in UploadPart) and UploadPartCopy) requests
// must match the headers you used in the request to initiate the upload by
// using CreateMultipartUpload.
//
// To perform a multipart upload with encryption using an AWS KMS CMK, the requester
// must have permission to the kms:Encrypt, kms:Decrypt, kms:ReEncrypt*, kms:GenerateDataKey*,
// and kms:DescribeKey actions on the key. These permissions are required because
// Amazon S3 must decrypt and read data from the encrypted file parts before
// it completes the multipart upload.
//
// If your AWS Identity and Access Management (IAM) user or role is in the same
// AWS account as the AWS KMS CMK, then you must have these permissions on the
// key policy. If your IAM user or role belongs to a different account than
// the key, then you must have the permissions on both the key policy and your
// IAM user or role.
//
// For more information, see Protecting Data Using Server-Side Encryption (https://docs.aws.amazon.com/AmazonS3/latest/dev/serv-side-encryption.html).
//
// Access Permissions
//
// When copying an object, you can optionally specify the accounts or groups
// that should be granted specific permissions on the new object. There are
// two ways to grant the permissions using the request headers:
//
//    * Specify a canned ACL with the x-amz-acl request header. For more information,
//    see Canned ACL (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL).
//
//    * Specify access permissions explicitly with the x-amz-grant-read, x-amz-grant-read-acp,
//    x-amz-grant-write-acp, and x-amz-grant-full-control headers. These parameters
//    map to the set of permissions that Amazon S3 supports in an ACL. For more
//    information, see Access Control List (ACL) Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html).
//
// You can use either a canned ACL or specify access permissions explicitly.
// You cannot do both.
//
// Server-Side- Encryption-Specific Request Headers
//
// You can optionally tell Amazon S3 to encrypt data at rest using server-side
// encryption. Server-side encryption is for data encryption at rest. Amazon
// S3 encrypts your data as it writes it to disks in its data centers and decrypts
// it when you access it. The option you use depends on whether you want to
// use AWS-managed encryption keys or provide your own encryption key.
//
//    * Use encryption keys managed by Amazon S3 or customer master keys (CMKs)
//    stored in Amazon Key Management Service (KMS) – If you want AWS to manage
//    the keys used to encrypt data, specify the following headers in the request.
//    x-amz-server-side​-encryption x-amz-server-side-encryption-aws-kms-key-id
//    x-amz-server-side-encryption-context If you specify x-amz-server-side-encryption:aws:kms,
//    but don't provide x-amz-server-side- encryption-aws-kms-key-id, Amazon
//    S3 uses the AWS managed CMK in AWS KMS to protect the data. All GET and
//    PUT requests for an object protected by AWS KMS fail if you don't make
//    them with SSL or by using SigV4. For more information on Server-Side Encryption
//    with CMKs Stored in Amazon KMS (SSE-KMS), see Protecting Data Using Server-Side
//    Encryption with CMKs stored in AWS KMS (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html).
//
//    * Use customer-provided encryption keys – If you want to manage your
//    own encryption keys, provide all the following headers in the request.
//    x-amz-server-side​-encryption​-customer-algorithm x-amz-server-side​-encryption​-customer-key
//    x-amz-server-side​-encryption​-customer-key-MD5 For more information
//    on Server-Side Encryption with CMKs stored in AWS KMS (SSE-KMS), see Protecting
//    Data Using Server-Side Encryption with CMKs stored in AWS KMS (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html).
//
// Access-Control-List (ACL)-Specific Request Headers
//
// You also can use the following access control–related headers with this
// operation. By default, all objects are private. Only the owner has full access
// control. When adding a new object, you can grant permissions to individual
// AWS accounts or to predefined groups defined by Amazon S3. These permissions
// are then added to the Access Control List (ACL) on the object. For more information,
// see Using ACLs (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html).
// With this operation, you can grant access permissions using one of the following
// two methods:
//
//    * Specify a canned ACL (x-amz-acl) — Amazon S3 supports a set of predefined
//    ACLs, known as canned ACLs. Each canned ACL has a predefined set of grantees
//    and permissions. For more information, see Canned ACL (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL).
//
//    * Specify access permissions explicitly — To explicitly grant access
//    permissions to specific AWS accounts or groups, use the following headers.
//    Each header maps to specific permissions that Amazon S3 supports in an
//    ACL. For more information, see Access Control List (ACL) Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html).
//    In the header, you specify a list of grantees who get the specific permission.
//    To grant permissions explicitly use: x-amz-grant-read x-amz-grant-write
//    x-amz-grant-read-acp x-amz-grant-write-acp x-amz-grant-full-control You
//    specify each grantee as a type=value pair, where the type is one of the
//    following: emailAddress – if the value specified is the email address
//    of an AWS account id – if the value specified is the canonical user
//    ID of an AWS account uri – if you are granting permissions to a predefined
//    group For example, the following x-amz-grant-read header grants the AWS
//    accounts identified by email addresses permissions to read object data
//    and its metadata: x-amz-grant-read: emailAddress="xyz@amazon.com", emailAddress="abc@amazon.com"
//
// The following operations are related to CreateMultipartUpload:
//
//    * UploadPart
//
//    * CompleteMultipartUpload
//
//    * AbortMultipartUpload
//
//    * ListParts
//
//    * ListMultipartUploads
//
//    // Example sending a request using CreateMultipartUploadRequest.
//    req := client.CreateMultipartUploadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/CreateMultipartUpload
func (c *Client) CreateMultipartUploadRequest(input *CreateMultipartUploadInput) CreateMultipartUploadRequest {
	op := &aws.Operation{
		Name:       opCreateMultipartUpload,
		HTTPMethod: "POST",
		HTTPPath:   "/{Bucket}/{Key+}?uploads",
	}

	if input == nil {
		input = &CreateMultipartUploadInput{}
	}

	req := c.newRequest(op, input, &CreateMultipartUploadOutput{})
	return CreateMultipartUploadRequest{Request: req, Input: input, Copy: c.CreateMultipartUploadRequest}
}

// CreateMultipartUploadRequest is the request type for the
// CreateMultipartUpload API operation.
type CreateMultipartUploadRequest struct {
	*aws.Request
	Input *CreateMultipartUploadInput
	Copy  func(*CreateMultipartUploadInput) CreateMultipartUploadRequest
}

// Send marshals and sends the CreateMultipartUpload API request.
func (r CreateMultipartUploadRequest) Send(ctx context.Context) (*CreateMultipartUploadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMultipartUploadResponse{
		CreateMultipartUploadOutput: r.Request.Data.(*CreateMultipartUploadOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMultipartUploadResponse is the response type for the
// CreateMultipartUpload API operation.
type CreateMultipartUploadResponse struct {
	*CreateMultipartUploadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMultipartUpload request.
func (r *CreateMultipartUploadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
