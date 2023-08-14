// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a copy of an object that is already stored in Amazon S3. You can store
// individual objects of up to 5 TB in Amazon S3. You create a copy of your object
// up to 5 GB in size in a single atomic action using this API. However, to copy an
// object greater than 5 GB, you must use the multipart upload Upload Part - Copy
// (UploadPartCopy) API. For more information, see Copy Object Using the REST
// Multipart Upload API (https://docs.aws.amazon.com/AmazonS3/latest/dev/CopyingObjctsUsingRESTMPUapi.html)
// . All copy requests must be authenticated. Additionally, you must have read
// access to the source object and write access to the destination bucket. For more
// information, see REST Authentication (https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html)
// . Both the Region that you want to copy the object from and the Region that you
// want to copy the object to must be enabled for your account. A copy request
// might return an error when Amazon S3 receives the copy request or while Amazon
// S3 is copying the files. If the error occurs before the copy action starts, you
// receive a standard Amazon S3 error. If the error occurs during the copy
// operation, the error response is embedded in the 200 OK response. This means
// that a 200 OK response can contain either a success or an error. If you call
// the S3 API directly, make sure to design your application to parse the contents
// of the response and handle it appropriately. If you use Amazon Web Services
// SDKs, SDKs handle this condition. The SDKs detect the embedded error and apply
// error handling per your configuration settings (including automatically retrying
// the request as appropriate). If the condition persists, the SDKs throws an
// exception (or, for the SDKs that don't use exceptions, they return the error).
// If the copy is successful, you receive a response with information about the
// copied object. If the request is an HTTP 1.1 request, the response is chunk
// encoded. If it were not, it would not contain the content-length, and you would
// need to read the entire body. The copy request charge is based on the storage
// class and Region that you specify for the destination object. For pricing
// information, see Amazon S3 pricing (http://aws.amazon.com/s3/pricing/) . Amazon
// S3 transfer acceleration does not support cross-Region copies. If you request a
// cross-Region copy using a transfer acceleration endpoint, you get a 400 Bad
// Request error. For more information, see Transfer Acceleration (https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html)
// . Metadata When copying an object, you can preserve all metadata (the default)
// or specify new metadata. However, the access control list (ACL) is not preserved
// and is set to private for the user making the request. To override the default
// ACL setting, specify a new ACL when generating a copy request. For more
// information, see Using ACLs (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html)
// . To specify whether you want the object metadata copied from the source object
// or replaced with metadata provided in the request, you can optionally add the
// x-amz-metadata-directive header. When you grant permissions, you can use the
// s3:x-amz-metadata-directive condition key to enforce certain metadata behavior
// when objects are uploaded. For more information, see Specifying Conditions in a
// Policy (https://docs.aws.amazon.com/AmazonS3/latest/dev/amazon-s3-policy-keys.html)
// in the Amazon S3 User Guide. For a complete list of Amazon S3-specific condition
// keys, see Actions, Resources, and Condition Keys for Amazon S3 (https://docs.aws.amazon.com/AmazonS3/latest/dev/list_amazons3.html)
// . x-amz-website-redirect-location is unique to each object and must be
// specified in the request headers to copy the value. x-amz-copy-source-if Headers
// To only copy an object under certain conditions, such as whether the Etag
// matches or whether the object was modified before or after a specified date, use
// the following request parameters:
//   - x-amz-copy-source-if-match
//   - x-amz-copy-source-if-none-match
//   - x-amz-copy-source-if-unmodified-since
//   - x-amz-copy-source-if-modified-since
//
// If both the x-amz-copy-source-if-match and x-amz-copy-source-if-unmodified-since
// headers are present in the request and evaluate as follows, Amazon S3 returns
// 200 OK and copies the data:
//   - x-amz-copy-source-if-match condition evaluates to true
//   - x-amz-copy-source-if-unmodified-since condition evaluates to false
//
// If both the x-amz-copy-source-if-none-match and
// x-amz-copy-source-if-modified-since headers are present in the request and
// evaluate as follows, Amazon S3 returns the 412 Precondition Failed response
// code:
//   - x-amz-copy-source-if-none-match condition evaluates to false
//   - x-amz-copy-source-if-modified-since condition evaluates to true
//
// All headers with the x-amz- prefix, including x-amz-copy-source , must be
// signed. Server-side encryption Amazon S3 automatically encrypts all new objects
// that are copied to an S3 bucket. When copying an object, if you don't specify
// encryption information in your copy request, the encryption setting of the
// target object is set to the default encryption configuration of the destination
// bucket. By default, all buckets have a base level of encryption configuration
// that uses server-side encryption with Amazon S3 managed keys (SSE-S3). If the
// destination bucket has a default encryption configuration that uses server-side
// encryption with Key Management Service (KMS) keys (SSE-KMS), dual-layer
// server-side encryption with Amazon Web Services KMS keys (DSSE-KMS), or
// server-side encryption with customer-provided encryption keys (SSE-C), Amazon S3
// uses the corresponding KMS key, or a customer-provided key to encrypt the target
// object copy. When you perform a CopyObject operation, if you want to use a
// different type of encryption setting for the target object, you can use other
// appropriate encryption-related headers to encrypt the target object with a KMS
// key, an Amazon S3 managed key, or a customer-provided key. With server-side
// encryption, Amazon S3 encrypts your data as it writes your data to disks in its
// data centers and decrypts the data when you access it. If the encryption setting
// in your request is different from the default encryption configuration of the
// destination bucket, the encryption setting in your request takes precedence. If
// the source object for the copy is stored in Amazon S3 using SSE-C, you must
// provide the necessary encryption information in your request so that Amazon S3
// can decrypt the object for copying. For more information about server-side
// encryption, see Using Server-Side Encryption (https://docs.aws.amazon.com/AmazonS3/latest/dev/serv-side-encryption.html)
// . If a target object uses SSE-KMS, you can enable an S3 Bucket Key for the
// object. For more information, see Amazon S3 Bucket Keys (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html)
// in the Amazon S3 User Guide. Access Control List (ACL)-Specific Request Headers
// When copying an object, you can optionally use headers to grant ACL-based
// permissions. By default, all objects are private. Only the owner has full access
// control. When adding a new object, you can grant permissions to individual
// Amazon Web Services accounts or to predefined groups that are defined by Amazon
// S3. These permissions are then added to the ACL on the object. For more
// information, see Access Control List (ACL) Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html)
// and Managing ACLs Using the REST API (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-using-rest-api.html)
// . If the bucket that you're copying objects to uses the bucket owner enforced
// setting for S3 Object Ownership, ACLs are disabled and no longer affect
// permissions. Buckets that use this setting only accept PUT requests that don't
// specify an ACL or PUT requests that specify bucket owner full control ACLs,
// such as the bucket-owner-full-control canned ACL or an equivalent form of this
// ACL expressed in the XML format. For more information, see Controlling
// ownership of objects and disabling ACLs (https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html)
// in the Amazon S3 User Guide. If your bucket uses the bucket owner enforced
// setting for Object Ownership, all objects written to the bucket by any account
// will be owned by the bucket owner. Checksums When copying an object, if it has a
// checksum, that checksum will be copied to the new object by default. When you
// copy the object over, you can optionally specify a different checksum algorithm
// to use with the x-amz-checksum-algorithm header. Storage Class Options You can
// use the CopyObject action to change the storage class of an object that is
// already stored in Amazon S3 by using the StorageClass parameter. For more
// information, see Storage Classes (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
// in the Amazon S3 User Guide. If the source object's storage class is GLACIER,
// you must restore a copy of this object before you can use it as a source object
// for the copy operation. For more information, see RestoreObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_RestoreObject.html)
// . For more information, see Copying Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/CopyingObjectsExamples.html)
// . Versioning By default, x-amz-copy-source header identifies the current
// version of an object to copy. If the current version is a delete marker, Amazon
// S3 behaves as if the object was deleted. To copy a different version, use the
// versionId subresource. If you enable versioning on the target bucket, Amazon S3
// generates a unique version ID for the object being copied. This version ID is
// different from the version ID of the source object. Amazon S3 returns the
// version ID of the copied object in the x-amz-version-id response header in the
// response. If you do not enable versioning or suspend it on the target bucket,
// the version ID that Amazon S3 generates is always null. The following operations
// are related to CopyObject :
//   - PutObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
//   - GetObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)
func (c *Client) CopyObject(ctx context.Context, params *CopyObjectInput, optFns ...func(*Options)) (*CopyObjectOutput, error) {
	if params == nil {
		params = &CopyObjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyObject", params, optFns, c.addOperationCopyObjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyObjectInput struct {

	// The name of the destination bucket. When using this action with an access
	// point, you must direct requests to the access point hostname. The access point
	// hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// action with an access point through the Amazon Web Services SDKs, you provide
	// the access point ARN in place of the bucket name. For more information about
	// access point ARNs, see Using access points (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html)
	// in the Amazon S3 User Guide. When you use this action with Amazon S3 on
	// Outposts, you must direct requests to the S3 on Outposts hostname. The S3 on
	// Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com . When you
	// use this action with S3 on Outposts through the Amazon Web Services SDKs, you
	// provide the Outposts access point ARN in place of the bucket name. For more
	// information about S3 on Outposts ARNs, see What is S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html)
	// in the Amazon S3 User Guide.
	//
	// This member is required.
	Bucket *string

	// Specifies the source object for the copy operation. You specify the value in
	// one of two formats, depending on whether you want to access the source object
	// through an access point (https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points.html)
	// :
	//   - For objects not accessed through an access point, specify the name of the
	//   source bucket and the key of the source object, separated by a slash (/). For
	//   example, to copy the object reports/january.pdf from the bucket
	//   awsexamplebucket , use awsexamplebucket/reports/january.pdf . The value must
	//   be URL-encoded.
	//   - For objects accessed through access points, specify the Amazon Resource
	//   Name (ARN) of the object as accessed through the access point, in the format
	//   arn:aws:s3:::accesspoint//object/ . For example, to copy the object
	//   reports/january.pdf through access point my-access-point owned by account
	//   123456789012 in Region us-west-2 , use the URL encoding of
	//   arn:aws:s3:us-west-2:123456789012:accesspoint/my-access-point/object/reports/january.pdf
	//   . The value must be URL encoded. Amazon S3 supports copy operations using access
	//   points only when the source and destination buckets are in the same Amazon Web
	//   Services Region. Alternatively, for objects accessed through Amazon S3 on
	//   Outposts, specify the ARN of the object as accessed in the format
	//   arn:aws:s3-outposts:::outpost//object/ . For example, to copy the object
	//   reports/january.pdf through outpost my-outpost owned by account 123456789012
	//   in Region us-west-2 , use the URL encoding of
	//   arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/object/reports/january.pdf
	//   . The value must be URL-encoded.
	// To copy a specific version of an object, append ?versionId= to the value (for
	// example,
	// awsexamplebucket/reports/january.pdf?versionId=QUpfdndhfd8438MNFDN93jdnJFkdmqnh893
	// ). If you don't specify a version ID, Amazon S3 copies the latest version of the
	// source object.
	//
	// This member is required.
	CopySource *string

	// The key of the destination object.
	//
	// This member is required.
	Key *string

	// The canned ACL to apply to the object. This action is not supported by Amazon
	// S3 on Outposts.
	ACL types.ObjectCannedACL

	// Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption
	// with server-side encryption using Key Management Service (KMS) keys (SSE-KMS).
	// Setting this header to true causes Amazon S3 to use an S3 Bucket Key for object
	// encryption with SSE-KMS. Specifying this header with a COPY action doesn’t
	// affect bucket-level settings for S3 Bucket Key.
	BucketKeyEnabled bool

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string

	// Indicates the algorithm you want Amazon S3 to use to create the checksum for
	// the object. For more information, see Checking object integrity (https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html)
	// in the Amazon S3 User Guide.
	ChecksumAlgorithm types.ChecksumAlgorithm

	// Specifies presentational information for the object.
	ContentDisposition *string

	// Specifies what content encodings have been applied to the object and thus what
	// decoding mechanisms must be applied to obtain the media-type referenced by the
	// Content-Type header field.
	ContentEncoding *string

	// The language the content is in.
	ContentLanguage *string

	// A standard MIME type describing the format of the object data.
	ContentType *string

	// Copies the object if its entity tag (ETag) matches the specified tag.
	CopySourceIfMatch *string

	// Copies the object if it has been modified since the specified time.
	CopySourceIfModifiedSince *time.Time

	// Copies the object if its entity tag (ETag) is different than the specified ETag.
	CopySourceIfNoneMatch *string

	// Copies the object if it hasn't been modified since the specified time.
	CopySourceIfUnmodifiedSince *time.Time

	// Specifies the algorithm to use when decrypting the source object (for example,
	// AES256).
	CopySourceSSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use to decrypt
	// the source object. The encryption key provided in this header must be one that
	// was used when the source object was created.
	CopySourceSSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	CopySourceSSECustomerKeyMD5 *string

	// The account ID of the expected destination bucket owner. If the destination
	// bucket is owned by a different account, the request fails with the HTTP status
	// code 403 Forbidden (access denied).
	ExpectedBucketOwner *string

	// The account ID of the expected source bucket owner. If the source bucket is
	// owned by a different account, the request fails with the HTTP status code 403
	// Forbidden (access denied).
	ExpectedSourceBucketOwner *string

	// The date and time at which the object is no longer cacheable.
	Expires *time.Time

	// Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object. This
	// action is not supported by Amazon S3 on Outposts.
	GrantFullControl *string

	// Allows grantee to read the object data and its metadata. This action is not
	// supported by Amazon S3 on Outposts.
	GrantRead *string

	// Allows grantee to read the object ACL. This action is not supported by Amazon
	// S3 on Outposts.
	GrantReadACP *string

	// Allows grantee to write the ACL for the applicable object. This action is not
	// supported by Amazon S3 on Outposts.
	GrantWriteACP *string

	// A map of metadata to store with the object in S3.
	Metadata map[string]string

	// Specifies whether the metadata is copied from the source object or replaced
	// with metadata provided in the request.
	MetadataDirective types.MetadataDirective

	// Specifies whether you want to apply a legal hold to the copied object.
	ObjectLockLegalHoldStatus types.ObjectLockLegalHoldStatus

	// The Object Lock mode that you want to apply to the copied object.
	ObjectLockMode types.ObjectLockMode

	// The date and time when you want the copied object's Object Lock to expire.
	ObjectLockRetainUntilDate *time.Time

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from Requester Pays buckets, see Downloading Objects
	// in Requester Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 User Guide.
	RequestPayer types.RequestPayer

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use in
	// encrypting data. This value is used to store the object and then it is
	// discarded; Amazon S3 does not store the encryption key. The key must be
	// appropriate for use with the algorithm specified in the
	// x-amz-server-side-encryption-customer-algorithm header.
	SSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string

	// Specifies the Amazon Web Services KMS Encryption Context to use for object
	// encryption. The value of this header is a base64-encoded UTF-8 string holding
	// JSON with the encryption context key-value pairs.
	SSEKMSEncryptionContext *string

	// Specifies the KMS key ID to use for object encryption. All GET and PUT requests
	// for an object protected by KMS will fail if they're not made via SSL or using
	// SigV4. For information about configuring any of the officially supported Amazon
	// Web Services SDKs and Amazon Web Services CLI, see Specifying the Signature
	// Version in Request Authentication (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingAWSSDK.html#specify-signature-version)
	// in the Amazon S3 User Guide.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when storing this object in Amazon S3
	// (for example, AES256 , aws:kms , aws:kms:dsse ).
	ServerSideEncryption types.ServerSideEncryption

	// By default, Amazon S3 uses the STANDARD Storage Class to store newly created
	// objects. The STANDARD storage class provides high durability and high
	// availability. Depending on performance needs, you can specify a different
	// Storage Class. Amazon S3 on Outposts only uses the OUTPOSTS Storage Class. For
	// more information, see Storage Classes (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
	// in the Amazon S3 User Guide.
	StorageClass types.StorageClass

	// The tag-set for the object destination object this value must be used in
	// conjunction with the TaggingDirective . The tag-set must be encoded as URL Query
	// parameters.
	Tagging *string

	// Specifies whether the object tag-set are copied from the source object or
	// replaced with tag-set provided in the request.
	TaggingDirective types.TaggingDirective

	// If the bucket is configured as a website, redirects requests for this object to
	// another object in the same bucket or to an external URL. Amazon S3 stores the
	// value of this header in the object metadata. This value is unique to each object
	// and is not copied when using the x-amz-metadata-directive header. Instead, you
	// may opt to provide this header in combination with the directive.
	WebsiteRedirectLocation *string

	noSmithyDocumentSerde
}

type CopyObjectOutput struct {

	// Indicates whether the copied object uses an S3 Bucket Key for server-side
	// encryption with Key Management Service (KMS) keys (SSE-KMS).
	BucketKeyEnabled bool

	// Container for all response elements.
	CopyObjectResult *types.CopyObjectResult

	// Version of the copied object in the destination bucket.
	CopySourceVersionId *string

	// If the object expiration is configured, the response includes this header.
	Expiration *string

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// If server-side encryption with a customer-provided encryption key was
	// requested, the response will include this header confirming the encryption
	// algorithm used.
	SSECustomerAlgorithm *string

	// If server-side encryption with a customer-provided encryption key was
	// requested, the response will include this header to provide round-trip message
	// integrity verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string

	// If present, specifies the Amazon Web Services KMS Encryption Context to use for
	// object encryption. The value of this header is a base64-encoded UTF-8 string
	// holding JSON with the encryption context key-value pairs.
	SSEKMSEncryptionContext *string

	// If present, specifies the ID of the Key Management Service (KMS) symmetric
	// encryption customer managed key that was used for the object.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when storing this object in Amazon S3
	// (for example, AES256 , aws:kms , aws:kms:dsse ).
	ServerSideEncryption types.ServerSideEncryption

	// Version ID of the newly created copy.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyObjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCopyObject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCopyObject{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCopyObjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyObject(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addCopyObjectUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = s3cust.HandleResponseErrorWith200Status(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCopyObject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "CopyObject",
	}
}

// getCopyObjectBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getCopyObjectBucketMember(input interface{}) (*string, bool) {
	in := input.(*CopyObjectInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addCopyObjectUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getCopyObjectBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
