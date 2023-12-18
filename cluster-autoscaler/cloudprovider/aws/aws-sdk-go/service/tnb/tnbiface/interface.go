// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package tnbiface provides an interface to enable mocking the AWS Telco Network Builder service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package tnbiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/tnb"
)

// TnbAPI provides an interface to enable mocking the
// tnb.Tnb service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Telco Network Builder.
//	func myFunc(svc tnbiface.TnbAPI) bool {
//	    // Make svc.CancelSolNetworkOperation request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := tnb.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockTnbClient struct {
//	    tnbiface.TnbAPI
//	}
//	func (m *mockTnbClient) CancelSolNetworkOperation(input *tnb.CancelSolNetworkOperationInput) (*tnb.CancelSolNetworkOperationOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockTnbClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type TnbAPI interface {
	CancelSolNetworkOperation(*tnb.CancelSolNetworkOperationInput) (*tnb.CancelSolNetworkOperationOutput, error)
	CancelSolNetworkOperationWithContext(aws.Context, *tnb.CancelSolNetworkOperationInput, ...request.Option) (*tnb.CancelSolNetworkOperationOutput, error)
	CancelSolNetworkOperationRequest(*tnb.CancelSolNetworkOperationInput) (*request.Request, *tnb.CancelSolNetworkOperationOutput)

	CreateSolFunctionPackage(*tnb.CreateSolFunctionPackageInput) (*tnb.CreateSolFunctionPackageOutput, error)
	CreateSolFunctionPackageWithContext(aws.Context, *tnb.CreateSolFunctionPackageInput, ...request.Option) (*tnb.CreateSolFunctionPackageOutput, error)
	CreateSolFunctionPackageRequest(*tnb.CreateSolFunctionPackageInput) (*request.Request, *tnb.CreateSolFunctionPackageOutput)

	CreateSolNetworkInstance(*tnb.CreateSolNetworkInstanceInput) (*tnb.CreateSolNetworkInstanceOutput, error)
	CreateSolNetworkInstanceWithContext(aws.Context, *tnb.CreateSolNetworkInstanceInput, ...request.Option) (*tnb.CreateSolNetworkInstanceOutput, error)
	CreateSolNetworkInstanceRequest(*tnb.CreateSolNetworkInstanceInput) (*request.Request, *tnb.CreateSolNetworkInstanceOutput)

	CreateSolNetworkPackage(*tnb.CreateSolNetworkPackageInput) (*tnb.CreateSolNetworkPackageOutput, error)
	CreateSolNetworkPackageWithContext(aws.Context, *tnb.CreateSolNetworkPackageInput, ...request.Option) (*tnb.CreateSolNetworkPackageOutput, error)
	CreateSolNetworkPackageRequest(*tnb.CreateSolNetworkPackageInput) (*request.Request, *tnb.CreateSolNetworkPackageOutput)

	DeleteSolFunctionPackage(*tnb.DeleteSolFunctionPackageInput) (*tnb.DeleteSolFunctionPackageOutput, error)
	DeleteSolFunctionPackageWithContext(aws.Context, *tnb.DeleteSolFunctionPackageInput, ...request.Option) (*tnb.DeleteSolFunctionPackageOutput, error)
	DeleteSolFunctionPackageRequest(*tnb.DeleteSolFunctionPackageInput) (*request.Request, *tnb.DeleteSolFunctionPackageOutput)

	DeleteSolNetworkInstance(*tnb.DeleteSolNetworkInstanceInput) (*tnb.DeleteSolNetworkInstanceOutput, error)
	DeleteSolNetworkInstanceWithContext(aws.Context, *tnb.DeleteSolNetworkInstanceInput, ...request.Option) (*tnb.DeleteSolNetworkInstanceOutput, error)
	DeleteSolNetworkInstanceRequest(*tnb.DeleteSolNetworkInstanceInput) (*request.Request, *tnb.DeleteSolNetworkInstanceOutput)

	DeleteSolNetworkPackage(*tnb.DeleteSolNetworkPackageInput) (*tnb.DeleteSolNetworkPackageOutput, error)
	DeleteSolNetworkPackageWithContext(aws.Context, *tnb.DeleteSolNetworkPackageInput, ...request.Option) (*tnb.DeleteSolNetworkPackageOutput, error)
	DeleteSolNetworkPackageRequest(*tnb.DeleteSolNetworkPackageInput) (*request.Request, *tnb.DeleteSolNetworkPackageOutput)

	GetSolFunctionInstance(*tnb.GetSolFunctionInstanceInput) (*tnb.GetSolFunctionInstanceOutput, error)
	GetSolFunctionInstanceWithContext(aws.Context, *tnb.GetSolFunctionInstanceInput, ...request.Option) (*tnb.GetSolFunctionInstanceOutput, error)
	GetSolFunctionInstanceRequest(*tnb.GetSolFunctionInstanceInput) (*request.Request, *tnb.GetSolFunctionInstanceOutput)

	GetSolFunctionPackage(*tnb.GetSolFunctionPackageInput) (*tnb.GetSolFunctionPackageOutput, error)
	GetSolFunctionPackageWithContext(aws.Context, *tnb.GetSolFunctionPackageInput, ...request.Option) (*tnb.GetSolFunctionPackageOutput, error)
	GetSolFunctionPackageRequest(*tnb.GetSolFunctionPackageInput) (*request.Request, *tnb.GetSolFunctionPackageOutput)

	GetSolFunctionPackageContent(*tnb.GetSolFunctionPackageContentInput) (*tnb.GetSolFunctionPackageContentOutput, error)
	GetSolFunctionPackageContentWithContext(aws.Context, *tnb.GetSolFunctionPackageContentInput, ...request.Option) (*tnb.GetSolFunctionPackageContentOutput, error)
	GetSolFunctionPackageContentRequest(*tnb.GetSolFunctionPackageContentInput) (*request.Request, *tnb.GetSolFunctionPackageContentOutput)

	GetSolFunctionPackageDescriptor(*tnb.GetSolFunctionPackageDescriptorInput) (*tnb.GetSolFunctionPackageDescriptorOutput, error)
	GetSolFunctionPackageDescriptorWithContext(aws.Context, *tnb.GetSolFunctionPackageDescriptorInput, ...request.Option) (*tnb.GetSolFunctionPackageDescriptorOutput, error)
	GetSolFunctionPackageDescriptorRequest(*tnb.GetSolFunctionPackageDescriptorInput) (*request.Request, *tnb.GetSolFunctionPackageDescriptorOutput)

	GetSolNetworkInstance(*tnb.GetSolNetworkInstanceInput) (*tnb.GetSolNetworkInstanceOutput, error)
	GetSolNetworkInstanceWithContext(aws.Context, *tnb.GetSolNetworkInstanceInput, ...request.Option) (*tnb.GetSolNetworkInstanceOutput, error)
	GetSolNetworkInstanceRequest(*tnb.GetSolNetworkInstanceInput) (*request.Request, *tnb.GetSolNetworkInstanceOutput)

	GetSolNetworkOperation(*tnb.GetSolNetworkOperationInput) (*tnb.GetSolNetworkOperationOutput, error)
	GetSolNetworkOperationWithContext(aws.Context, *tnb.GetSolNetworkOperationInput, ...request.Option) (*tnb.GetSolNetworkOperationOutput, error)
	GetSolNetworkOperationRequest(*tnb.GetSolNetworkOperationInput) (*request.Request, *tnb.GetSolNetworkOperationOutput)

	GetSolNetworkPackage(*tnb.GetSolNetworkPackageInput) (*tnb.GetSolNetworkPackageOutput, error)
	GetSolNetworkPackageWithContext(aws.Context, *tnb.GetSolNetworkPackageInput, ...request.Option) (*tnb.GetSolNetworkPackageOutput, error)
	GetSolNetworkPackageRequest(*tnb.GetSolNetworkPackageInput) (*request.Request, *tnb.GetSolNetworkPackageOutput)

	GetSolNetworkPackageContent(*tnb.GetSolNetworkPackageContentInput) (*tnb.GetSolNetworkPackageContentOutput, error)
	GetSolNetworkPackageContentWithContext(aws.Context, *tnb.GetSolNetworkPackageContentInput, ...request.Option) (*tnb.GetSolNetworkPackageContentOutput, error)
	GetSolNetworkPackageContentRequest(*tnb.GetSolNetworkPackageContentInput) (*request.Request, *tnb.GetSolNetworkPackageContentOutput)

	GetSolNetworkPackageDescriptor(*tnb.GetSolNetworkPackageDescriptorInput) (*tnb.GetSolNetworkPackageDescriptorOutput, error)
	GetSolNetworkPackageDescriptorWithContext(aws.Context, *tnb.GetSolNetworkPackageDescriptorInput, ...request.Option) (*tnb.GetSolNetworkPackageDescriptorOutput, error)
	GetSolNetworkPackageDescriptorRequest(*tnb.GetSolNetworkPackageDescriptorInput) (*request.Request, *tnb.GetSolNetworkPackageDescriptorOutput)

	InstantiateSolNetworkInstance(*tnb.InstantiateSolNetworkInstanceInput) (*tnb.InstantiateSolNetworkInstanceOutput, error)
	InstantiateSolNetworkInstanceWithContext(aws.Context, *tnb.InstantiateSolNetworkInstanceInput, ...request.Option) (*tnb.InstantiateSolNetworkInstanceOutput, error)
	InstantiateSolNetworkInstanceRequest(*tnb.InstantiateSolNetworkInstanceInput) (*request.Request, *tnb.InstantiateSolNetworkInstanceOutput)

	ListSolFunctionInstances(*tnb.ListSolFunctionInstancesInput) (*tnb.ListSolFunctionInstancesOutput, error)
	ListSolFunctionInstancesWithContext(aws.Context, *tnb.ListSolFunctionInstancesInput, ...request.Option) (*tnb.ListSolFunctionInstancesOutput, error)
	ListSolFunctionInstancesRequest(*tnb.ListSolFunctionInstancesInput) (*request.Request, *tnb.ListSolFunctionInstancesOutput)

	ListSolFunctionInstancesPages(*tnb.ListSolFunctionInstancesInput, func(*tnb.ListSolFunctionInstancesOutput, bool) bool) error
	ListSolFunctionInstancesPagesWithContext(aws.Context, *tnb.ListSolFunctionInstancesInput, func(*tnb.ListSolFunctionInstancesOutput, bool) bool, ...request.Option) error

	ListSolFunctionPackages(*tnb.ListSolFunctionPackagesInput) (*tnb.ListSolFunctionPackagesOutput, error)
	ListSolFunctionPackagesWithContext(aws.Context, *tnb.ListSolFunctionPackagesInput, ...request.Option) (*tnb.ListSolFunctionPackagesOutput, error)
	ListSolFunctionPackagesRequest(*tnb.ListSolFunctionPackagesInput) (*request.Request, *tnb.ListSolFunctionPackagesOutput)

	ListSolFunctionPackagesPages(*tnb.ListSolFunctionPackagesInput, func(*tnb.ListSolFunctionPackagesOutput, bool) bool) error
	ListSolFunctionPackagesPagesWithContext(aws.Context, *tnb.ListSolFunctionPackagesInput, func(*tnb.ListSolFunctionPackagesOutput, bool) bool, ...request.Option) error

	ListSolNetworkInstances(*tnb.ListSolNetworkInstancesInput) (*tnb.ListSolNetworkInstancesOutput, error)
	ListSolNetworkInstancesWithContext(aws.Context, *tnb.ListSolNetworkInstancesInput, ...request.Option) (*tnb.ListSolNetworkInstancesOutput, error)
	ListSolNetworkInstancesRequest(*tnb.ListSolNetworkInstancesInput) (*request.Request, *tnb.ListSolNetworkInstancesOutput)

	ListSolNetworkInstancesPages(*tnb.ListSolNetworkInstancesInput, func(*tnb.ListSolNetworkInstancesOutput, bool) bool) error
	ListSolNetworkInstancesPagesWithContext(aws.Context, *tnb.ListSolNetworkInstancesInput, func(*tnb.ListSolNetworkInstancesOutput, bool) bool, ...request.Option) error

	ListSolNetworkOperations(*tnb.ListSolNetworkOperationsInput) (*tnb.ListSolNetworkOperationsOutput, error)
	ListSolNetworkOperationsWithContext(aws.Context, *tnb.ListSolNetworkOperationsInput, ...request.Option) (*tnb.ListSolNetworkOperationsOutput, error)
	ListSolNetworkOperationsRequest(*tnb.ListSolNetworkOperationsInput) (*request.Request, *tnb.ListSolNetworkOperationsOutput)

	ListSolNetworkOperationsPages(*tnb.ListSolNetworkOperationsInput, func(*tnb.ListSolNetworkOperationsOutput, bool) bool) error
	ListSolNetworkOperationsPagesWithContext(aws.Context, *tnb.ListSolNetworkOperationsInput, func(*tnb.ListSolNetworkOperationsOutput, bool) bool, ...request.Option) error

	ListSolNetworkPackages(*tnb.ListSolNetworkPackagesInput) (*tnb.ListSolNetworkPackagesOutput, error)
	ListSolNetworkPackagesWithContext(aws.Context, *tnb.ListSolNetworkPackagesInput, ...request.Option) (*tnb.ListSolNetworkPackagesOutput, error)
	ListSolNetworkPackagesRequest(*tnb.ListSolNetworkPackagesInput) (*request.Request, *tnb.ListSolNetworkPackagesOutput)

	ListSolNetworkPackagesPages(*tnb.ListSolNetworkPackagesInput, func(*tnb.ListSolNetworkPackagesOutput, bool) bool) error
	ListSolNetworkPackagesPagesWithContext(aws.Context, *tnb.ListSolNetworkPackagesInput, func(*tnb.ListSolNetworkPackagesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*tnb.ListTagsForResourceInput) (*tnb.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *tnb.ListTagsForResourceInput, ...request.Option) (*tnb.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*tnb.ListTagsForResourceInput) (*request.Request, *tnb.ListTagsForResourceOutput)

	PutSolFunctionPackageContent(*tnb.PutSolFunctionPackageContentInput) (*tnb.PutSolFunctionPackageContentOutput, error)
	PutSolFunctionPackageContentWithContext(aws.Context, *tnb.PutSolFunctionPackageContentInput, ...request.Option) (*tnb.PutSolFunctionPackageContentOutput, error)
	PutSolFunctionPackageContentRequest(*tnb.PutSolFunctionPackageContentInput) (*request.Request, *tnb.PutSolFunctionPackageContentOutput)

	PutSolNetworkPackageContent(*tnb.PutSolNetworkPackageContentInput) (*tnb.PutSolNetworkPackageContentOutput, error)
	PutSolNetworkPackageContentWithContext(aws.Context, *tnb.PutSolNetworkPackageContentInput, ...request.Option) (*tnb.PutSolNetworkPackageContentOutput, error)
	PutSolNetworkPackageContentRequest(*tnb.PutSolNetworkPackageContentInput) (*request.Request, *tnb.PutSolNetworkPackageContentOutput)

	TagResource(*tnb.TagResourceInput) (*tnb.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *tnb.TagResourceInput, ...request.Option) (*tnb.TagResourceOutput, error)
	TagResourceRequest(*tnb.TagResourceInput) (*request.Request, *tnb.TagResourceOutput)

	TerminateSolNetworkInstance(*tnb.TerminateSolNetworkInstanceInput) (*tnb.TerminateSolNetworkInstanceOutput, error)
	TerminateSolNetworkInstanceWithContext(aws.Context, *tnb.TerminateSolNetworkInstanceInput, ...request.Option) (*tnb.TerminateSolNetworkInstanceOutput, error)
	TerminateSolNetworkInstanceRequest(*tnb.TerminateSolNetworkInstanceInput) (*request.Request, *tnb.TerminateSolNetworkInstanceOutput)

	UntagResource(*tnb.UntagResourceInput) (*tnb.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *tnb.UntagResourceInput, ...request.Option) (*tnb.UntagResourceOutput, error)
	UntagResourceRequest(*tnb.UntagResourceInput) (*request.Request, *tnb.UntagResourceOutput)

	UpdateSolFunctionPackage(*tnb.UpdateSolFunctionPackageInput) (*tnb.UpdateSolFunctionPackageOutput, error)
	UpdateSolFunctionPackageWithContext(aws.Context, *tnb.UpdateSolFunctionPackageInput, ...request.Option) (*tnb.UpdateSolFunctionPackageOutput, error)
	UpdateSolFunctionPackageRequest(*tnb.UpdateSolFunctionPackageInput) (*request.Request, *tnb.UpdateSolFunctionPackageOutput)

	UpdateSolNetworkInstance(*tnb.UpdateSolNetworkInstanceInput) (*tnb.UpdateSolNetworkInstanceOutput, error)
	UpdateSolNetworkInstanceWithContext(aws.Context, *tnb.UpdateSolNetworkInstanceInput, ...request.Option) (*tnb.UpdateSolNetworkInstanceOutput, error)
	UpdateSolNetworkInstanceRequest(*tnb.UpdateSolNetworkInstanceInput) (*request.Request, *tnb.UpdateSolNetworkInstanceOutput)

	UpdateSolNetworkPackage(*tnb.UpdateSolNetworkPackageInput) (*tnb.UpdateSolNetworkPackageOutput, error)
	UpdateSolNetworkPackageWithContext(aws.Context, *tnb.UpdateSolNetworkPackageInput, ...request.Option) (*tnb.UpdateSolNetworkPackageOutput, error)
	UpdateSolNetworkPackageRequest(*tnb.UpdateSolNetworkPackageInput) (*request.Request, *tnb.UpdateSolNetworkPackageOutput)

	ValidateSolFunctionPackageContent(*tnb.ValidateSolFunctionPackageContentInput) (*tnb.ValidateSolFunctionPackageContentOutput, error)
	ValidateSolFunctionPackageContentWithContext(aws.Context, *tnb.ValidateSolFunctionPackageContentInput, ...request.Option) (*tnb.ValidateSolFunctionPackageContentOutput, error)
	ValidateSolFunctionPackageContentRequest(*tnb.ValidateSolFunctionPackageContentInput) (*request.Request, *tnb.ValidateSolFunctionPackageContentOutput)

	ValidateSolNetworkPackageContent(*tnb.ValidateSolNetworkPackageContentInput) (*tnb.ValidateSolNetworkPackageContentOutput, error)
	ValidateSolNetworkPackageContentWithContext(aws.Context, *tnb.ValidateSolNetworkPackageContentInput, ...request.Option) (*tnb.ValidateSolNetworkPackageContentOutput, error)
	ValidateSolNetworkPackageContentRequest(*tnb.ValidateSolNetworkPackageContentInput) (*request.Request, *tnb.ValidateSolNetworkPackageContentOutput)
}

var _ TnbAPI = (*tnb.Tnb)(nil)
