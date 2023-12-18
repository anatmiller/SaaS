// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kendra

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You don't have sufficient access to perform this action. Please ensure you
	// have the required permission policies and user accounts and try again.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// A conflict occurred with the request. Please fix any inconsistences with
	// your resources and try again.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeFeaturedResultsConflictException for service response error code
	// "FeaturedResultsConflictException".
	//
	// An error message with a list of conflicting queries used across different
	// sets of featured results. This occurred with the request for a new featured
	// results set. Check that the queries you specified for featured results are
	// unique per featured results set for each index.
	ErrCodeFeaturedResultsConflictException = "FeaturedResultsConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// An issue occurred with the internal server used for your Amazon Kendra service.
	// Please wait a few minutes and try again, or contact Support (http://aws.amazon.com/contact-us/)
	// for help.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The input to the request is not valid. Please provide the correct input and
	// try again.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeResourceAlreadyExistException for service response error code
	// "ResourceAlreadyExistException".
	//
	// The resource you want to use already exists. Please check you have provided
	// the correct resource and try again.
	ErrCodeResourceAlreadyExistException = "ResourceAlreadyExistException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The resource you want to use is currently in use. Please check you have provided
	// the correct resource and try again.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource you want to use doesn’t exist. Please check you have provided
	// the correct resource and try again.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourceUnavailableException for service response error code
	// "ResourceUnavailableException".
	//
	// The resource you want to use isn't available. Please check you have provided
	// the correct resource and try again.
	ErrCodeResourceUnavailableException = "ResourceUnavailableException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// You have exceeded the set limits for your Amazon Kendra service. Please see
	// Quotas (https://docs.aws.amazon.com/kendra/latest/dg/quotas.html) for more
	// information, or contact Support (http://aws.amazon.com/contact-us/) to inquire
	// about an increase of limits.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling. Please reduce the number
	// of requests and try again.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The input fails to satisfy the constraints set by the Amazon Kendra service.
	// Please provide the correct input and try again.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":            newErrorAccessDeniedException,
	"ConflictException":                newErrorConflictException,
	"FeaturedResultsConflictException": newErrorFeaturedResultsConflictException,
	"InternalServerException":          newErrorInternalServerException,
	"InvalidRequestException":          newErrorInvalidRequestException,
	"ResourceAlreadyExistException":    newErrorResourceAlreadyExistException,
	"ResourceInUseException":           newErrorResourceInUseException,
	"ResourceNotFoundException":        newErrorResourceNotFoundException,
	"ResourceUnavailableException":     newErrorResourceUnavailableException,
	"ServiceQuotaExceededException":    newErrorServiceQuotaExceededException,
	"ThrottlingException":              newErrorThrottlingException,
	"ValidationException":              newErrorValidationException,
}
