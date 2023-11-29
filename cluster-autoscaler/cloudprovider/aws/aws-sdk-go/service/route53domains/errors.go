// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeDnssecLimitExceeded for service response error code
	// "DnssecLimitExceeded".
	//
	// This error is returned if you call AssociateDelegationSignerToDomain when
	// the specified domain has reached the maximum number of DS records. You can't
	// add any additional DS records unless you delete an existing one first.
	ErrCodeDnssecLimitExceeded = "DnssecLimitExceeded"

	// ErrCodeDomainLimitExceeded for service response error code
	// "DomainLimitExceeded".
	//
	// The number of domains has exceeded the allowed threshold for the account.
	ErrCodeDomainLimitExceeded = "DomainLimitExceeded"

	// ErrCodeDuplicateRequest for service response error code
	// "DuplicateRequest".
	//
	// The request is already in progress for the domain.
	ErrCodeDuplicateRequest = "DuplicateRequest"

	// ErrCodeInvalidInput for service response error code
	// "InvalidInput".
	//
	// The requested item is not acceptable. For example, for APIs that accept a
	// domain name, the request might specify a domain name that doesn't belong
	// to the account that submitted the request. For AcceptDomainTransferFromAnotherAwsAccount,
	// the password might be invalid.
	ErrCodeInvalidInput = "InvalidInput"

	// ErrCodeOperationLimitExceeded for service response error code
	// "OperationLimitExceeded".
	//
	// The number of operations or jobs running exceeded the allowed threshold for
	// the account.
	ErrCodeOperationLimitExceeded = "OperationLimitExceeded"

	// ErrCodeTLDRulesViolation for service response error code
	// "TLDRulesViolation".
	//
	// The top-level domain does not support this operation.
	ErrCodeTLDRulesViolation = "TLDRulesViolation"

	// ErrCodeUnsupportedTLD for service response error code
	// "UnsupportedTLD".
	//
	// Amazon Route 53 does not support this top-level domain (TLD).
	ErrCodeUnsupportedTLD = "UnsupportedTLD"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"DnssecLimitExceeded":    newErrorDnssecLimitExceeded,
	"DomainLimitExceeded":    newErrorDomainLimitExceeded,
	"DuplicateRequest":       newErrorDuplicateRequest,
	"InvalidInput":           newErrorInvalidInput,
	"OperationLimitExceeded": newErrorOperationLimitExceeded,
	"TLDRulesViolation":      newErrorTLDRulesViolation,
	"UnsupportedTLD":         newErrorUnsupportedTLD,
}
