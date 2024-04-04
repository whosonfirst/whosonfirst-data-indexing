// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the roles for an identity pool. You must use AWS Developer credentials to
// call this API.
func (c *Client) GetIdentityPoolRoles(ctx context.Context, params *GetIdentityPoolRolesInput, optFns ...func(*Options)) (*GetIdentityPoolRolesOutput, error) {
	if params == nil {
		params = &GetIdentityPoolRolesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIdentityPoolRoles", params, optFns, c.addOperationGetIdentityPoolRolesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIdentityPoolRolesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the GetIdentityPoolRoles action.
type GetIdentityPoolRolesInput struct {

	// An identity pool ID in the format REGION:GUID.
	//
	// This member is required.
	IdentityPoolId *string

	noSmithyDocumentSerde
}

// Returned in response to a successful GetIdentityPoolRoles operation.
type GetIdentityPoolRolesOutput struct {

	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId *string

	// How users for a specific identity provider are to mapped to roles. This is a
	// String-to- RoleMapping object map. The string identifies the identity provider,
	// for example, "graph.facebook.com" or
	// "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id".
	RoleMappings map[string]types.RoleMapping

	// The map of roles associated with this pool. Currently only authenticated and
	// unauthenticated roles are supported.
	Roles map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIdentityPoolRolesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetIdentityPoolRoles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetIdentityPoolRoles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetIdentityPoolRoles"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetIdentityPoolRolesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdentityPoolRoles(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetIdentityPoolRoles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetIdentityPoolRoles",
	}
}
