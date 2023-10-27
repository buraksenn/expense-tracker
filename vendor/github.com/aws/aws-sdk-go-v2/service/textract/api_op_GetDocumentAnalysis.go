// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the results for an Amazon Textract asynchronous operation that analyzes
// text in a document. You start asynchronous text analysis by calling
// StartDocumentAnalysis , which returns a job identifier ( JobId ). When the text
// analysis operation finishes, Amazon Textract publishes a completion status to
// the Amazon Simple Notification Service (Amazon SNS) topic that's registered in
// the initial call to StartDocumentAnalysis . To get the results of the
// text-detection operation, first check that the status value published to the
// Amazon SNS topic is SUCCEEDED . If so, call GetDocumentAnalysis , and pass the
// job identifier ( JobId ) from the initial call to StartDocumentAnalysis .
// GetDocumentAnalysis returns an array of Block objects. The following types of
// information are returned:
//   - Form data (key-value pairs). The related information is returned in two
//     Block objects, each of type KEY_VALUE_SET : a KEY Block object and a VALUE
//     Block object. For example, Name: Ana Silva Carolina contains a key and value.
//     Name: is the key. Ana Silva Carolina is the value.
//   - Table and table cell data. A TABLE Block object contains information about a
//     detected table. A CELL Block object is returned for each cell in a table.
//   - Lines and words of text. A LINE Block object contains one or more WORD Block
//     objects. All lines and words that are detected in the document are returned
//     (including text that doesn't have a relationship with the value of the
//     StartDocumentAnalysis FeatureTypes input parameter).
//   - Query. A QUERY Block object contains the query text, alias and link to the
//     associated Query results block object.
//   - Query Results. A QUERY_RESULT Block object contains the answer to the query
//     and an ID that connects it to the query asked. This Block also contains a
//     confidence score.
//
// While processing a document with queries, look out for
// INVALID_REQUEST_PARAMETERS output. This indicates that either the per page query
// limit has been exceeded or that the operation is trying to query a page in the
// document which doesn’t exist. Selection elements such as check boxes and option
// buttons (radio buttons) can be detected in form data and in tables. A
// SELECTION_ELEMENT Block object contains information about a selection element,
// including the selection status. Use the MaxResults parameter to limit the
// number of blocks that are returned. If there are more results than specified in
// MaxResults , the value of NextToken in the operation response contains a
// pagination token for getting the next set of results. To get the next page of
// results, call GetDocumentAnalysis , and populate the NextToken request
// parameter with the token value that's returned from the previous call to
// GetDocumentAnalysis . For more information, see Document Text Analysis (https://docs.aws.amazon.com/textract/latest/dg/how-it-works-analyzing.html)
// .
func (c *Client) GetDocumentAnalysis(ctx context.Context, params *GetDocumentAnalysisInput, optFns ...func(*Options)) (*GetDocumentAnalysisOutput, error) {
	if params == nil {
		params = &GetDocumentAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDocumentAnalysis", params, optFns, c.addOperationGetDocumentAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDocumentAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDocumentAnalysisInput struct {

	// A unique identifier for the text-detection job. The JobId is returned from
	// StartDocumentAnalysis . A JobId value is only valid for 7 days.
	//
	// This member is required.
	JobId *string

	// The maximum number of results to return per paginated call. The largest value
	// that you can specify is 1,000. If you specify a value greater than 1,000, a
	// maximum of 1,000 results is returned. The default value is 1,000.
	MaxResults *int32

	// If the previous response was incomplete (because there are more blocks to
	// retrieve), Amazon Textract returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of blocks.
	NextToken *string

	noSmithyDocumentSerde
}

type GetDocumentAnalysisOutput struct {

	//
	AnalyzeDocumentModelVersion *string

	// The results of the text-analysis operation.
	Blocks []types.Block

	// Information about a document that Amazon Textract processed. DocumentMetadata
	// is returned in every page of paginated responses from an Amazon Textract video
	// operation.
	DocumentMetadata *types.DocumentMetadata

	// The current status of the text detection job.
	JobStatus types.JobStatus

	// If the response is truncated, Amazon Textract returns this token. You can use
	// this token in the subsequent request to retrieve the next set of text detection
	// results.
	NextToken *string

	// Returns if the detection job could not be completed. Contains explanation for
	// what error occured.
	StatusMessage *string

	// A list of warnings that occurred during the document-analysis operation.
	Warnings []types.Warning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDocumentAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDocumentAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDocumentAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addGetDocumentAnalysisResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetDocumentAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDocumentAnalysis(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetDocumentAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "textract",
		OperationName: "GetDocumentAnalysis",
	}
}

type opGetDocumentAnalysisResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetDocumentAnalysisResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetDocumentAnalysisResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "textract"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "textract"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("textract")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addGetDocumentAnalysisResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetDocumentAnalysisResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}