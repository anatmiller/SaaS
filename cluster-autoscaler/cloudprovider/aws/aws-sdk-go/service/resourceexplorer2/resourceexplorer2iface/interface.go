// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package resourceexplorer2iface provides an interface to enable mocking the AWS Resource Explorer service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package resourceexplorer2iface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/resourceexplorer2"
)

// ResourceExplorer2API provides an interface to enable mocking the
// resourceexplorer2.ResourceExplorer2 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Resource Explorer.
//	func myFunc(svc resourceexplorer2iface.ResourceExplorer2API) bool {
//	    // Make svc.AssociateDefaultView request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := resourceexplorer2.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockResourceExplorer2Client struct {
//	    resourceexplorer2iface.ResourceExplorer2API
//	}
//	func (m *mockResourceExplorer2Client) AssociateDefaultView(input *resourceexplorer2.AssociateDefaultViewInput) (*resourceexplorer2.AssociateDefaultViewOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockResourceExplorer2Client{}
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
type ResourceExplorer2API interface {
	AssociateDefaultView(*resourceexplorer2.AssociateDefaultViewInput) (*resourceexplorer2.AssociateDefaultViewOutput, error)
	AssociateDefaultViewWithContext(aws.Context, *resourceexplorer2.AssociateDefaultViewInput, ...request.Option) (*resourceexplorer2.AssociateDefaultViewOutput, error)
	AssociateDefaultViewRequest(*resourceexplorer2.AssociateDefaultViewInput) (*request.Request, *resourceexplorer2.AssociateDefaultViewOutput)

	BatchGetView(*resourceexplorer2.BatchGetViewInput) (*resourceexplorer2.BatchGetViewOutput, error)
	BatchGetViewWithContext(aws.Context, *resourceexplorer2.BatchGetViewInput, ...request.Option) (*resourceexplorer2.BatchGetViewOutput, error)
	BatchGetViewRequest(*resourceexplorer2.BatchGetViewInput) (*request.Request, *resourceexplorer2.BatchGetViewOutput)

	CreateIndex(*resourceexplorer2.CreateIndexInput) (*resourceexplorer2.CreateIndexOutput, error)
	CreateIndexWithContext(aws.Context, *resourceexplorer2.CreateIndexInput, ...request.Option) (*resourceexplorer2.CreateIndexOutput, error)
	CreateIndexRequest(*resourceexplorer2.CreateIndexInput) (*request.Request, *resourceexplorer2.CreateIndexOutput)

	CreateView(*resourceexplorer2.CreateViewInput) (*resourceexplorer2.CreateViewOutput, error)
	CreateViewWithContext(aws.Context, *resourceexplorer2.CreateViewInput, ...request.Option) (*resourceexplorer2.CreateViewOutput, error)
	CreateViewRequest(*resourceexplorer2.CreateViewInput) (*request.Request, *resourceexplorer2.CreateViewOutput)

	DeleteIndex(*resourceexplorer2.DeleteIndexInput) (*resourceexplorer2.DeleteIndexOutput, error)
	DeleteIndexWithContext(aws.Context, *resourceexplorer2.DeleteIndexInput, ...request.Option) (*resourceexplorer2.DeleteIndexOutput, error)
	DeleteIndexRequest(*resourceexplorer2.DeleteIndexInput) (*request.Request, *resourceexplorer2.DeleteIndexOutput)

	DeleteView(*resourceexplorer2.DeleteViewInput) (*resourceexplorer2.DeleteViewOutput, error)
	DeleteViewWithContext(aws.Context, *resourceexplorer2.DeleteViewInput, ...request.Option) (*resourceexplorer2.DeleteViewOutput, error)
	DeleteViewRequest(*resourceexplorer2.DeleteViewInput) (*request.Request, *resourceexplorer2.DeleteViewOutput)

	DisassociateDefaultView(*resourceexplorer2.DisassociateDefaultViewInput) (*resourceexplorer2.DisassociateDefaultViewOutput, error)
	DisassociateDefaultViewWithContext(aws.Context, *resourceexplorer2.DisassociateDefaultViewInput, ...request.Option) (*resourceexplorer2.DisassociateDefaultViewOutput, error)
	DisassociateDefaultViewRequest(*resourceexplorer2.DisassociateDefaultViewInput) (*request.Request, *resourceexplorer2.DisassociateDefaultViewOutput)

	GetAccountLevelServiceConfiguration(*resourceexplorer2.GetAccountLevelServiceConfigurationInput) (*resourceexplorer2.GetAccountLevelServiceConfigurationOutput, error)
	GetAccountLevelServiceConfigurationWithContext(aws.Context, *resourceexplorer2.GetAccountLevelServiceConfigurationInput, ...request.Option) (*resourceexplorer2.GetAccountLevelServiceConfigurationOutput, error)
	GetAccountLevelServiceConfigurationRequest(*resourceexplorer2.GetAccountLevelServiceConfigurationInput) (*request.Request, *resourceexplorer2.GetAccountLevelServiceConfigurationOutput)

	GetDefaultView(*resourceexplorer2.GetDefaultViewInput) (*resourceexplorer2.GetDefaultViewOutput, error)
	GetDefaultViewWithContext(aws.Context, *resourceexplorer2.GetDefaultViewInput, ...request.Option) (*resourceexplorer2.GetDefaultViewOutput, error)
	GetDefaultViewRequest(*resourceexplorer2.GetDefaultViewInput) (*request.Request, *resourceexplorer2.GetDefaultViewOutput)

	GetIndex(*resourceexplorer2.GetIndexInput) (*resourceexplorer2.GetIndexOutput, error)
	GetIndexWithContext(aws.Context, *resourceexplorer2.GetIndexInput, ...request.Option) (*resourceexplorer2.GetIndexOutput, error)
	GetIndexRequest(*resourceexplorer2.GetIndexInput) (*request.Request, *resourceexplorer2.GetIndexOutput)

	GetView(*resourceexplorer2.GetViewInput) (*resourceexplorer2.GetViewOutput, error)
	GetViewWithContext(aws.Context, *resourceexplorer2.GetViewInput, ...request.Option) (*resourceexplorer2.GetViewOutput, error)
	GetViewRequest(*resourceexplorer2.GetViewInput) (*request.Request, *resourceexplorer2.GetViewOutput)

	ListIndexes(*resourceexplorer2.ListIndexesInput) (*resourceexplorer2.ListIndexesOutput, error)
	ListIndexesWithContext(aws.Context, *resourceexplorer2.ListIndexesInput, ...request.Option) (*resourceexplorer2.ListIndexesOutput, error)
	ListIndexesRequest(*resourceexplorer2.ListIndexesInput) (*request.Request, *resourceexplorer2.ListIndexesOutput)

	ListIndexesPages(*resourceexplorer2.ListIndexesInput, func(*resourceexplorer2.ListIndexesOutput, bool) bool) error
	ListIndexesPagesWithContext(aws.Context, *resourceexplorer2.ListIndexesInput, func(*resourceexplorer2.ListIndexesOutput, bool) bool, ...request.Option) error

	ListIndexesForMembers(*resourceexplorer2.ListIndexesForMembersInput) (*resourceexplorer2.ListIndexesForMembersOutput, error)
	ListIndexesForMembersWithContext(aws.Context, *resourceexplorer2.ListIndexesForMembersInput, ...request.Option) (*resourceexplorer2.ListIndexesForMembersOutput, error)
	ListIndexesForMembersRequest(*resourceexplorer2.ListIndexesForMembersInput) (*request.Request, *resourceexplorer2.ListIndexesForMembersOutput)

	ListIndexesForMembersPages(*resourceexplorer2.ListIndexesForMembersInput, func(*resourceexplorer2.ListIndexesForMembersOutput, bool) bool) error
	ListIndexesForMembersPagesWithContext(aws.Context, *resourceexplorer2.ListIndexesForMembersInput, func(*resourceexplorer2.ListIndexesForMembersOutput, bool) bool, ...request.Option) error

	ListSupportedResourceTypes(*resourceexplorer2.ListSupportedResourceTypesInput) (*resourceexplorer2.ListSupportedResourceTypesOutput, error)
	ListSupportedResourceTypesWithContext(aws.Context, *resourceexplorer2.ListSupportedResourceTypesInput, ...request.Option) (*resourceexplorer2.ListSupportedResourceTypesOutput, error)
	ListSupportedResourceTypesRequest(*resourceexplorer2.ListSupportedResourceTypesInput) (*request.Request, *resourceexplorer2.ListSupportedResourceTypesOutput)

	ListSupportedResourceTypesPages(*resourceexplorer2.ListSupportedResourceTypesInput, func(*resourceexplorer2.ListSupportedResourceTypesOutput, bool) bool) error
	ListSupportedResourceTypesPagesWithContext(aws.Context, *resourceexplorer2.ListSupportedResourceTypesInput, func(*resourceexplorer2.ListSupportedResourceTypesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*resourceexplorer2.ListTagsForResourceInput) (*resourceexplorer2.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *resourceexplorer2.ListTagsForResourceInput, ...request.Option) (*resourceexplorer2.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*resourceexplorer2.ListTagsForResourceInput) (*request.Request, *resourceexplorer2.ListTagsForResourceOutput)

	ListViews(*resourceexplorer2.ListViewsInput) (*resourceexplorer2.ListViewsOutput, error)
	ListViewsWithContext(aws.Context, *resourceexplorer2.ListViewsInput, ...request.Option) (*resourceexplorer2.ListViewsOutput, error)
	ListViewsRequest(*resourceexplorer2.ListViewsInput) (*request.Request, *resourceexplorer2.ListViewsOutput)

	ListViewsPages(*resourceexplorer2.ListViewsInput, func(*resourceexplorer2.ListViewsOutput, bool) bool) error
	ListViewsPagesWithContext(aws.Context, *resourceexplorer2.ListViewsInput, func(*resourceexplorer2.ListViewsOutput, bool) bool, ...request.Option) error

	Search(*resourceexplorer2.SearchInput) (*resourceexplorer2.SearchOutput, error)
	SearchWithContext(aws.Context, *resourceexplorer2.SearchInput, ...request.Option) (*resourceexplorer2.SearchOutput, error)
	SearchRequest(*resourceexplorer2.SearchInput) (*request.Request, *resourceexplorer2.SearchOutput)

	SearchPages(*resourceexplorer2.SearchInput, func(*resourceexplorer2.SearchOutput, bool) bool) error
	SearchPagesWithContext(aws.Context, *resourceexplorer2.SearchInput, func(*resourceexplorer2.SearchOutput, bool) bool, ...request.Option) error

	TagResource(*resourceexplorer2.TagResourceInput) (*resourceexplorer2.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *resourceexplorer2.TagResourceInput, ...request.Option) (*resourceexplorer2.TagResourceOutput, error)
	TagResourceRequest(*resourceexplorer2.TagResourceInput) (*request.Request, *resourceexplorer2.TagResourceOutput)

	UntagResource(*resourceexplorer2.UntagResourceInput) (*resourceexplorer2.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *resourceexplorer2.UntagResourceInput, ...request.Option) (*resourceexplorer2.UntagResourceOutput, error)
	UntagResourceRequest(*resourceexplorer2.UntagResourceInput) (*request.Request, *resourceexplorer2.UntagResourceOutput)

	UpdateIndexType(*resourceexplorer2.UpdateIndexTypeInput) (*resourceexplorer2.UpdateIndexTypeOutput, error)
	UpdateIndexTypeWithContext(aws.Context, *resourceexplorer2.UpdateIndexTypeInput, ...request.Option) (*resourceexplorer2.UpdateIndexTypeOutput, error)
	UpdateIndexTypeRequest(*resourceexplorer2.UpdateIndexTypeInput) (*request.Request, *resourceexplorer2.UpdateIndexTypeOutput)

	UpdateView(*resourceexplorer2.UpdateViewInput) (*resourceexplorer2.UpdateViewOutput, error)
	UpdateViewWithContext(aws.Context, *resourceexplorer2.UpdateViewInput, ...request.Option) (*resourceexplorer2.UpdateViewOutput, error)
	UpdateViewRequest(*resourceexplorer2.UpdateViewInput) (*request.Request, *resourceexplorer2.UpdateViewOutput)
}

var _ ResourceExplorer2API = (*resourceexplorer2.ResourceExplorer2)(nil)
