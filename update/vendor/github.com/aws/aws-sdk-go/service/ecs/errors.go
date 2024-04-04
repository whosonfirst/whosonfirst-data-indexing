// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You don't have authorization to perform the requested action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeAttributeLimitExceededException for service response error code
	// "AttributeLimitExceededException".
	//
	// You can apply up to 10 custom attributes for each resource. You can view
	// the attributes of a resource with ListAttributes. You can remove existing
	// attributes on a resource with DeleteAttributes.
	ErrCodeAttributeLimitExceededException = "AttributeLimitExceededException"

	// ErrCodeBlockedException for service response error code
	// "BlockedException".
	//
	// Your Amazon Web Services account was blocked. For more information, contact
	// Amazon Web Services Support (http://aws.amazon.com/contact-us/).
	ErrCodeBlockedException = "BlockedException"

	// ErrCodeClientException for service response error code
	// "ClientException".
	//
	// These errors are usually caused by a client action. This client action might
	// be using an action or resource on behalf of a user that doesn't have permissions
	// to use the action or resource. Or, it might be specifying an identifier that
	// isn't valid.
	ErrCodeClientException = "ClientException"

	// ErrCodeClusterContainsContainerInstancesException for service response error code
	// "ClusterContainsContainerInstancesException".
	//
	// You can't delete a cluster that has registered container instances. First,
	// deregister the container instances before you can delete the cluster. For
	// more information, see DeregisterContainerInstance.
	ErrCodeClusterContainsContainerInstancesException = "ClusterContainsContainerInstancesException"

	// ErrCodeClusterContainsServicesException for service response error code
	// "ClusterContainsServicesException".
	//
	// You can't delete a cluster that contains services. First, update the service
	// to reduce its desired task count to 0, and then delete the service. For more
	// information, see UpdateService and DeleteService.
	ErrCodeClusterContainsServicesException = "ClusterContainsServicesException"

	// ErrCodeClusterContainsTasksException for service response error code
	// "ClusterContainsTasksException".
	//
	// You can't delete a cluster that has active tasks.
	ErrCodeClusterContainsTasksException = "ClusterContainsTasksException"

	// ErrCodeClusterNotFoundException for service response error code
	// "ClusterNotFoundException".
	//
	// The specified cluster wasn't found. You can view your available clusters
	// with ListClusters. Amazon ECS clusters are Region specific.
	ErrCodeClusterNotFoundException = "ClusterNotFoundException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The RunTask request could not be processed due to conflicts. The provided
	// clientToken is already in use with a different RunTask request. The resourceIds
	// are the existing task ARNs which are already associated with the clientToken.
	//
	// To fix this issue:
	//
	//    * Run RunTask with a unique clientToken.
	//
	//    * Run RunTask with the clientToken and the original set of parameters
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// The specified parameter isn't valid. Review the available parameters for
	// the API request.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The limit for the resource was exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMissingVersionException for service response error code
	// "MissingVersionException".
	//
	// Amazon ECS can't determine the current version of the Amazon ECS container
	// agent on the container instance and doesn't have enough information to proceed
	// with an update. This could be because the agent running on the container
	// instance is a previous or custom version that doesn't use our version information.
	ErrCodeMissingVersionException = "MissingVersionException"

	// ErrCodeNamespaceNotFoundException for service response error code
	// "NamespaceNotFoundException".
	//
	// The specified namespace wasn't found.
	ErrCodeNamespaceNotFoundException = "NamespaceNotFoundException"

	// ErrCodeNoUpdateAvailableException for service response error code
	// "NoUpdateAvailableException".
	//
	// There's no update available for this Amazon ECS container agent. This might
	// be because the agent is already running the latest version or because it's
	// so old that there's no update path to the current version.
	ErrCodeNoUpdateAvailableException = "NoUpdateAvailableException"

	// ErrCodePlatformTaskDefinitionIncompatibilityException for service response error code
	// "PlatformTaskDefinitionIncompatibilityException".
	//
	// The specified platform version doesn't satisfy the required capabilities
	// of the task definition.
	ErrCodePlatformTaskDefinitionIncompatibilityException = "PlatformTaskDefinitionIncompatibilityException"

	// ErrCodePlatformUnknownException for service response error code
	// "PlatformUnknownException".
	//
	// The specified platform version doesn't exist.
	ErrCodePlatformUnknownException = "PlatformUnknownException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The specified resource is in-use and can't be removed.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource wasn't found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServerException for service response error code
	// "ServerException".
	//
	// These errors are usually caused by a server issue.
	ErrCodeServerException = "ServerException"

	// ErrCodeServiceNotActiveException for service response error code
	// "ServiceNotActiveException".
	//
	// The specified service isn't active. You can't update a service that's inactive.
	// If you have previously deleted a service, you can re-create it with CreateService.
	ErrCodeServiceNotActiveException = "ServiceNotActiveException"

	// ErrCodeServiceNotFoundException for service response error code
	// "ServiceNotFoundException".
	//
	// The specified service wasn't found. You can view your available services
	// with ListServices. Amazon ECS services are cluster specific and Region specific.
	ErrCodeServiceNotFoundException = "ServiceNotFoundException"

	// ErrCodeTargetNotConnectedException for service response error code
	// "TargetNotConnectedException".
	//
	// The execute command cannot run. This error can be caused by any of the following
	// configuration issues:
	//
	//    * Incorrect IAM permissions
	//
	//    * The SSM agent is not installed or is not running
	//
	//    * There is an interface Amazon VPC endpoint for Amazon ECS, but there
	//    is not one for Systems Manager Session Manager
	//
	// For information about how to troubleshoot the issues, see Troubleshooting
	// issues with ECS Exec (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-exec.html)
	// in the Amazon Elastic Container Service Developer Guide.
	ErrCodeTargetNotConnectedException = "TargetNotConnectedException"

	// ErrCodeTargetNotFoundException for service response error code
	// "TargetNotFoundException".
	//
	// The specified target wasn't found. You can view your available container
	// instances with ListContainerInstances. Amazon ECS container instances are
	// cluster-specific and Region-specific.
	ErrCodeTargetNotFoundException = "TargetNotFoundException"

	// ErrCodeTaskSetNotFoundException for service response error code
	// "TaskSetNotFoundException".
	//
	// The specified task set wasn't found. You can view your available task sets
	// with DescribeTaskSets. Task sets are specific to each cluster, service and
	// Region.
	ErrCodeTaskSetNotFoundException = "TaskSetNotFoundException"

	// ErrCodeUnsupportedFeatureException for service response error code
	// "UnsupportedFeatureException".
	//
	// The specified task isn't supported in this Region.
	ErrCodeUnsupportedFeatureException = "UnsupportedFeatureException"

	// ErrCodeUpdateInProgressException for service response error code
	// "UpdateInProgressException".
	//
	// There's already a current Amazon ECS container agent update in progress on
	// the container instance that's specified. If the container agent becomes disconnected
	// while it's in a transitional stage, such as PENDING or STAGING, the update
	// process can get stuck in that state. However, when the agent reconnects,
	// it resumes where it stopped previously.
	ErrCodeUpdateInProgressException = "UpdateInProgressException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                          newErrorAccessDeniedException,
	"AttributeLimitExceededException":                newErrorAttributeLimitExceededException,
	"BlockedException":                               newErrorBlockedException,
	"ClientException":                                newErrorClientException,
	"ClusterContainsContainerInstancesException":     newErrorClusterContainsContainerInstancesException,
	"ClusterContainsServicesException":               newErrorClusterContainsServicesException,
	"ClusterContainsTasksException":                  newErrorClusterContainsTasksException,
	"ClusterNotFoundException":                       newErrorClusterNotFoundException,
	"ConflictException":                              newErrorConflictException,
	"InvalidParameterException":                      newErrorInvalidParameterException,
	"LimitExceededException":                         newErrorLimitExceededException,
	"MissingVersionException":                        newErrorMissingVersionException,
	"NamespaceNotFoundException":                     newErrorNamespaceNotFoundException,
	"NoUpdateAvailableException":                     newErrorNoUpdateAvailableException,
	"PlatformTaskDefinitionIncompatibilityException": newErrorPlatformTaskDefinitionIncompatibilityException,
	"PlatformUnknownException":                       newErrorPlatformUnknownException,
	"ResourceInUseException":                         newErrorResourceInUseException,
	"ResourceNotFoundException":                      newErrorResourceNotFoundException,
	"ServerException":                                newErrorServerException,
	"ServiceNotActiveException":                      newErrorServiceNotActiveException,
	"ServiceNotFoundException":                       newErrorServiceNotFoundException,
	"TargetNotConnectedException":                    newErrorTargetNotConnectedException,
	"TargetNotFoundException":                        newErrorTargetNotFoundException,
	"TaskSetNotFoundException":                       newErrorTaskSetNotFoundException,
	"UnsupportedFeatureException":                    newErrorUnsupportedFeatureException,
	"UpdateInProgressException":                      newErrorUpdateInProgressException,
}
