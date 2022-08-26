// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package billingconductoriface provides an interface to enable mocking the AWSBillingConductor service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package billingconductoriface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/billingconductor"
)

// BillingConductorAPI provides an interface to enable mocking the
// billingconductor.BillingConductor service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWSBillingConductor.
//	func myFunc(svc billingconductoriface.BillingConductorAPI) bool {
//	    // Make svc.AssociateAccounts request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := billingconductor.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockBillingConductorClient struct {
//	    billingconductoriface.BillingConductorAPI
//	}
//	func (m *mockBillingConductorClient) AssociateAccounts(input *billingconductor.AssociateAccountsInput) (*billingconductor.AssociateAccountsOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockBillingConductorClient{}
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
type BillingConductorAPI interface {
	AssociateAccounts(*billingconductor.AssociateAccountsInput) (*billingconductor.AssociateAccountsOutput, error)
	AssociateAccountsWithContext(aws.Context, *billingconductor.AssociateAccountsInput, ...request.Option) (*billingconductor.AssociateAccountsOutput, error)
	AssociateAccountsRequest(*billingconductor.AssociateAccountsInput) (*request.Request, *billingconductor.AssociateAccountsOutput)

	AssociatePricingRules(*billingconductor.AssociatePricingRulesInput) (*billingconductor.AssociatePricingRulesOutput, error)
	AssociatePricingRulesWithContext(aws.Context, *billingconductor.AssociatePricingRulesInput, ...request.Option) (*billingconductor.AssociatePricingRulesOutput, error)
	AssociatePricingRulesRequest(*billingconductor.AssociatePricingRulesInput) (*request.Request, *billingconductor.AssociatePricingRulesOutput)

	BatchAssociateResourcesToCustomLineItem(*billingconductor.BatchAssociateResourcesToCustomLineItemInput) (*billingconductor.BatchAssociateResourcesToCustomLineItemOutput, error)
	BatchAssociateResourcesToCustomLineItemWithContext(aws.Context, *billingconductor.BatchAssociateResourcesToCustomLineItemInput, ...request.Option) (*billingconductor.BatchAssociateResourcesToCustomLineItemOutput, error)
	BatchAssociateResourcesToCustomLineItemRequest(*billingconductor.BatchAssociateResourcesToCustomLineItemInput) (*request.Request, *billingconductor.BatchAssociateResourcesToCustomLineItemOutput)

	BatchDisassociateResourcesFromCustomLineItem(*billingconductor.BatchDisassociateResourcesFromCustomLineItemInput) (*billingconductor.BatchDisassociateResourcesFromCustomLineItemOutput, error)
	BatchDisassociateResourcesFromCustomLineItemWithContext(aws.Context, *billingconductor.BatchDisassociateResourcesFromCustomLineItemInput, ...request.Option) (*billingconductor.BatchDisassociateResourcesFromCustomLineItemOutput, error)
	BatchDisassociateResourcesFromCustomLineItemRequest(*billingconductor.BatchDisassociateResourcesFromCustomLineItemInput) (*request.Request, *billingconductor.BatchDisassociateResourcesFromCustomLineItemOutput)

	CreateBillingGroup(*billingconductor.CreateBillingGroupInput) (*billingconductor.CreateBillingGroupOutput, error)
	CreateBillingGroupWithContext(aws.Context, *billingconductor.CreateBillingGroupInput, ...request.Option) (*billingconductor.CreateBillingGroupOutput, error)
	CreateBillingGroupRequest(*billingconductor.CreateBillingGroupInput) (*request.Request, *billingconductor.CreateBillingGroupOutput)

	CreateCustomLineItem(*billingconductor.CreateCustomLineItemInput) (*billingconductor.CreateCustomLineItemOutput, error)
	CreateCustomLineItemWithContext(aws.Context, *billingconductor.CreateCustomLineItemInput, ...request.Option) (*billingconductor.CreateCustomLineItemOutput, error)
	CreateCustomLineItemRequest(*billingconductor.CreateCustomLineItemInput) (*request.Request, *billingconductor.CreateCustomLineItemOutput)

	CreatePricingPlan(*billingconductor.CreatePricingPlanInput) (*billingconductor.CreatePricingPlanOutput, error)
	CreatePricingPlanWithContext(aws.Context, *billingconductor.CreatePricingPlanInput, ...request.Option) (*billingconductor.CreatePricingPlanOutput, error)
	CreatePricingPlanRequest(*billingconductor.CreatePricingPlanInput) (*request.Request, *billingconductor.CreatePricingPlanOutput)

	CreatePricingRule(*billingconductor.CreatePricingRuleInput) (*billingconductor.CreatePricingRuleOutput, error)
	CreatePricingRuleWithContext(aws.Context, *billingconductor.CreatePricingRuleInput, ...request.Option) (*billingconductor.CreatePricingRuleOutput, error)
	CreatePricingRuleRequest(*billingconductor.CreatePricingRuleInput) (*request.Request, *billingconductor.CreatePricingRuleOutput)

	DeleteBillingGroup(*billingconductor.DeleteBillingGroupInput) (*billingconductor.DeleteBillingGroupOutput, error)
	DeleteBillingGroupWithContext(aws.Context, *billingconductor.DeleteBillingGroupInput, ...request.Option) (*billingconductor.DeleteBillingGroupOutput, error)
	DeleteBillingGroupRequest(*billingconductor.DeleteBillingGroupInput) (*request.Request, *billingconductor.DeleteBillingGroupOutput)

	DeleteCustomLineItem(*billingconductor.DeleteCustomLineItemInput) (*billingconductor.DeleteCustomLineItemOutput, error)
	DeleteCustomLineItemWithContext(aws.Context, *billingconductor.DeleteCustomLineItemInput, ...request.Option) (*billingconductor.DeleteCustomLineItemOutput, error)
	DeleteCustomLineItemRequest(*billingconductor.DeleteCustomLineItemInput) (*request.Request, *billingconductor.DeleteCustomLineItemOutput)

	DeletePricingPlan(*billingconductor.DeletePricingPlanInput) (*billingconductor.DeletePricingPlanOutput, error)
	DeletePricingPlanWithContext(aws.Context, *billingconductor.DeletePricingPlanInput, ...request.Option) (*billingconductor.DeletePricingPlanOutput, error)
	DeletePricingPlanRequest(*billingconductor.DeletePricingPlanInput) (*request.Request, *billingconductor.DeletePricingPlanOutput)

	DeletePricingRule(*billingconductor.DeletePricingRuleInput) (*billingconductor.DeletePricingRuleOutput, error)
	DeletePricingRuleWithContext(aws.Context, *billingconductor.DeletePricingRuleInput, ...request.Option) (*billingconductor.DeletePricingRuleOutput, error)
	DeletePricingRuleRequest(*billingconductor.DeletePricingRuleInput) (*request.Request, *billingconductor.DeletePricingRuleOutput)

	DisassociateAccounts(*billingconductor.DisassociateAccountsInput) (*billingconductor.DisassociateAccountsOutput, error)
	DisassociateAccountsWithContext(aws.Context, *billingconductor.DisassociateAccountsInput, ...request.Option) (*billingconductor.DisassociateAccountsOutput, error)
	DisassociateAccountsRequest(*billingconductor.DisassociateAccountsInput) (*request.Request, *billingconductor.DisassociateAccountsOutput)

	DisassociatePricingRules(*billingconductor.DisassociatePricingRulesInput) (*billingconductor.DisassociatePricingRulesOutput, error)
	DisassociatePricingRulesWithContext(aws.Context, *billingconductor.DisassociatePricingRulesInput, ...request.Option) (*billingconductor.DisassociatePricingRulesOutput, error)
	DisassociatePricingRulesRequest(*billingconductor.DisassociatePricingRulesInput) (*request.Request, *billingconductor.DisassociatePricingRulesOutput)

	ListAccountAssociations(*billingconductor.ListAccountAssociationsInput) (*billingconductor.ListAccountAssociationsOutput, error)
	ListAccountAssociationsWithContext(aws.Context, *billingconductor.ListAccountAssociationsInput, ...request.Option) (*billingconductor.ListAccountAssociationsOutput, error)
	ListAccountAssociationsRequest(*billingconductor.ListAccountAssociationsInput) (*request.Request, *billingconductor.ListAccountAssociationsOutput)

	ListAccountAssociationsPages(*billingconductor.ListAccountAssociationsInput, func(*billingconductor.ListAccountAssociationsOutput, bool) bool) error
	ListAccountAssociationsPagesWithContext(aws.Context, *billingconductor.ListAccountAssociationsInput, func(*billingconductor.ListAccountAssociationsOutput, bool) bool, ...request.Option) error

	ListBillingGroupCostReports(*billingconductor.ListBillingGroupCostReportsInput) (*billingconductor.ListBillingGroupCostReportsOutput, error)
	ListBillingGroupCostReportsWithContext(aws.Context, *billingconductor.ListBillingGroupCostReportsInput, ...request.Option) (*billingconductor.ListBillingGroupCostReportsOutput, error)
	ListBillingGroupCostReportsRequest(*billingconductor.ListBillingGroupCostReportsInput) (*request.Request, *billingconductor.ListBillingGroupCostReportsOutput)

	ListBillingGroupCostReportsPages(*billingconductor.ListBillingGroupCostReportsInput, func(*billingconductor.ListBillingGroupCostReportsOutput, bool) bool) error
	ListBillingGroupCostReportsPagesWithContext(aws.Context, *billingconductor.ListBillingGroupCostReportsInput, func(*billingconductor.ListBillingGroupCostReportsOutput, bool) bool, ...request.Option) error

	ListBillingGroups(*billingconductor.ListBillingGroupsInput) (*billingconductor.ListBillingGroupsOutput, error)
	ListBillingGroupsWithContext(aws.Context, *billingconductor.ListBillingGroupsInput, ...request.Option) (*billingconductor.ListBillingGroupsOutput, error)
	ListBillingGroupsRequest(*billingconductor.ListBillingGroupsInput) (*request.Request, *billingconductor.ListBillingGroupsOutput)

	ListBillingGroupsPages(*billingconductor.ListBillingGroupsInput, func(*billingconductor.ListBillingGroupsOutput, bool) bool) error
	ListBillingGroupsPagesWithContext(aws.Context, *billingconductor.ListBillingGroupsInput, func(*billingconductor.ListBillingGroupsOutput, bool) bool, ...request.Option) error

	ListCustomLineItems(*billingconductor.ListCustomLineItemsInput) (*billingconductor.ListCustomLineItemsOutput, error)
	ListCustomLineItemsWithContext(aws.Context, *billingconductor.ListCustomLineItemsInput, ...request.Option) (*billingconductor.ListCustomLineItemsOutput, error)
	ListCustomLineItemsRequest(*billingconductor.ListCustomLineItemsInput) (*request.Request, *billingconductor.ListCustomLineItemsOutput)

	ListCustomLineItemsPages(*billingconductor.ListCustomLineItemsInput, func(*billingconductor.ListCustomLineItemsOutput, bool) bool) error
	ListCustomLineItemsPagesWithContext(aws.Context, *billingconductor.ListCustomLineItemsInput, func(*billingconductor.ListCustomLineItemsOutput, bool) bool, ...request.Option) error

	ListPricingPlans(*billingconductor.ListPricingPlansInput) (*billingconductor.ListPricingPlansOutput, error)
	ListPricingPlansWithContext(aws.Context, *billingconductor.ListPricingPlansInput, ...request.Option) (*billingconductor.ListPricingPlansOutput, error)
	ListPricingPlansRequest(*billingconductor.ListPricingPlansInput) (*request.Request, *billingconductor.ListPricingPlansOutput)

	ListPricingPlansPages(*billingconductor.ListPricingPlansInput, func(*billingconductor.ListPricingPlansOutput, bool) bool) error
	ListPricingPlansPagesWithContext(aws.Context, *billingconductor.ListPricingPlansInput, func(*billingconductor.ListPricingPlansOutput, bool) bool, ...request.Option) error

	ListPricingPlansAssociatedWithPricingRule(*billingconductor.ListPricingPlansAssociatedWithPricingRuleInput) (*billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput, error)
	ListPricingPlansAssociatedWithPricingRuleWithContext(aws.Context, *billingconductor.ListPricingPlansAssociatedWithPricingRuleInput, ...request.Option) (*billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput, error)
	ListPricingPlansAssociatedWithPricingRuleRequest(*billingconductor.ListPricingPlansAssociatedWithPricingRuleInput) (*request.Request, *billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput)

	ListPricingPlansAssociatedWithPricingRulePages(*billingconductor.ListPricingPlansAssociatedWithPricingRuleInput, func(*billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput, bool) bool) error
	ListPricingPlansAssociatedWithPricingRulePagesWithContext(aws.Context, *billingconductor.ListPricingPlansAssociatedWithPricingRuleInput, func(*billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput, bool) bool, ...request.Option) error

	ListPricingRules(*billingconductor.ListPricingRulesInput) (*billingconductor.ListPricingRulesOutput, error)
	ListPricingRulesWithContext(aws.Context, *billingconductor.ListPricingRulesInput, ...request.Option) (*billingconductor.ListPricingRulesOutput, error)
	ListPricingRulesRequest(*billingconductor.ListPricingRulesInput) (*request.Request, *billingconductor.ListPricingRulesOutput)

	ListPricingRulesPages(*billingconductor.ListPricingRulesInput, func(*billingconductor.ListPricingRulesOutput, bool) bool) error
	ListPricingRulesPagesWithContext(aws.Context, *billingconductor.ListPricingRulesInput, func(*billingconductor.ListPricingRulesOutput, bool) bool, ...request.Option) error

	ListPricingRulesAssociatedToPricingPlan(*billingconductor.ListPricingRulesAssociatedToPricingPlanInput) (*billingconductor.ListPricingRulesAssociatedToPricingPlanOutput, error)
	ListPricingRulesAssociatedToPricingPlanWithContext(aws.Context, *billingconductor.ListPricingRulesAssociatedToPricingPlanInput, ...request.Option) (*billingconductor.ListPricingRulesAssociatedToPricingPlanOutput, error)
	ListPricingRulesAssociatedToPricingPlanRequest(*billingconductor.ListPricingRulesAssociatedToPricingPlanInput) (*request.Request, *billingconductor.ListPricingRulesAssociatedToPricingPlanOutput)

	ListPricingRulesAssociatedToPricingPlanPages(*billingconductor.ListPricingRulesAssociatedToPricingPlanInput, func(*billingconductor.ListPricingRulesAssociatedToPricingPlanOutput, bool) bool) error
	ListPricingRulesAssociatedToPricingPlanPagesWithContext(aws.Context, *billingconductor.ListPricingRulesAssociatedToPricingPlanInput, func(*billingconductor.ListPricingRulesAssociatedToPricingPlanOutput, bool) bool, ...request.Option) error

	ListResourcesAssociatedToCustomLineItem(*billingconductor.ListResourcesAssociatedToCustomLineItemInput) (*billingconductor.ListResourcesAssociatedToCustomLineItemOutput, error)
	ListResourcesAssociatedToCustomLineItemWithContext(aws.Context, *billingconductor.ListResourcesAssociatedToCustomLineItemInput, ...request.Option) (*billingconductor.ListResourcesAssociatedToCustomLineItemOutput, error)
	ListResourcesAssociatedToCustomLineItemRequest(*billingconductor.ListResourcesAssociatedToCustomLineItemInput) (*request.Request, *billingconductor.ListResourcesAssociatedToCustomLineItemOutput)

	ListResourcesAssociatedToCustomLineItemPages(*billingconductor.ListResourcesAssociatedToCustomLineItemInput, func(*billingconductor.ListResourcesAssociatedToCustomLineItemOutput, bool) bool) error
	ListResourcesAssociatedToCustomLineItemPagesWithContext(aws.Context, *billingconductor.ListResourcesAssociatedToCustomLineItemInput, func(*billingconductor.ListResourcesAssociatedToCustomLineItemOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*billingconductor.ListTagsForResourceInput) (*billingconductor.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *billingconductor.ListTagsForResourceInput, ...request.Option) (*billingconductor.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*billingconductor.ListTagsForResourceInput) (*request.Request, *billingconductor.ListTagsForResourceOutput)

	TagResource(*billingconductor.TagResourceInput) (*billingconductor.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *billingconductor.TagResourceInput, ...request.Option) (*billingconductor.TagResourceOutput, error)
	TagResourceRequest(*billingconductor.TagResourceInput) (*request.Request, *billingconductor.TagResourceOutput)

	UntagResource(*billingconductor.UntagResourceInput) (*billingconductor.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *billingconductor.UntagResourceInput, ...request.Option) (*billingconductor.UntagResourceOutput, error)
	UntagResourceRequest(*billingconductor.UntagResourceInput) (*request.Request, *billingconductor.UntagResourceOutput)

	UpdateBillingGroup(*billingconductor.UpdateBillingGroupInput) (*billingconductor.UpdateBillingGroupOutput, error)
	UpdateBillingGroupWithContext(aws.Context, *billingconductor.UpdateBillingGroupInput, ...request.Option) (*billingconductor.UpdateBillingGroupOutput, error)
	UpdateBillingGroupRequest(*billingconductor.UpdateBillingGroupInput) (*request.Request, *billingconductor.UpdateBillingGroupOutput)

	UpdateCustomLineItem(*billingconductor.UpdateCustomLineItemInput) (*billingconductor.UpdateCustomLineItemOutput, error)
	UpdateCustomLineItemWithContext(aws.Context, *billingconductor.UpdateCustomLineItemInput, ...request.Option) (*billingconductor.UpdateCustomLineItemOutput, error)
	UpdateCustomLineItemRequest(*billingconductor.UpdateCustomLineItemInput) (*request.Request, *billingconductor.UpdateCustomLineItemOutput)

	UpdatePricingPlan(*billingconductor.UpdatePricingPlanInput) (*billingconductor.UpdatePricingPlanOutput, error)
	UpdatePricingPlanWithContext(aws.Context, *billingconductor.UpdatePricingPlanInput, ...request.Option) (*billingconductor.UpdatePricingPlanOutput, error)
	UpdatePricingPlanRequest(*billingconductor.UpdatePricingPlanInput) (*request.Request, *billingconductor.UpdatePricingPlanOutput)

	UpdatePricingRule(*billingconductor.UpdatePricingRuleInput) (*billingconductor.UpdatePricingRuleOutput, error)
	UpdatePricingRuleWithContext(aws.Context, *billingconductor.UpdatePricingRuleInput, ...request.Option) (*billingconductor.UpdatePricingRuleOutput, error)
	UpdatePricingRuleRequest(*billingconductor.UpdatePricingRuleInput) (*request.Request, *billingconductor.UpdatePricingRuleOutput)
}

var _ BillingConductorAPI = (*billingconductor.BillingConductor)(nil)
