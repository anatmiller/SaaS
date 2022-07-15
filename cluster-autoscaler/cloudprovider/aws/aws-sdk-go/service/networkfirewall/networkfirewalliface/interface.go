// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package networkfirewalliface provides an interface to enable mocking the AWS Network Firewall service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package networkfirewalliface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/networkfirewall"
)

// NetworkFirewallAPI provides an interface to enable mocking the
// networkfirewall.NetworkFirewall service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Network Firewall.
//    func myFunc(svc networkfirewalliface.NetworkFirewallAPI) bool {
//        // Make svc.AssociateFirewallPolicy request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := networkfirewall.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockNetworkFirewallClient struct {
//        networkfirewalliface.NetworkFirewallAPI
//    }
//    func (m *mockNetworkFirewallClient) AssociateFirewallPolicy(input *networkfirewall.AssociateFirewallPolicyInput) (*networkfirewall.AssociateFirewallPolicyOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockNetworkFirewallClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type NetworkFirewallAPI interface {
	AssociateFirewallPolicy(*networkfirewall.AssociateFirewallPolicyInput) (*networkfirewall.AssociateFirewallPolicyOutput, error)
	AssociateFirewallPolicyWithContext(aws.Context, *networkfirewall.AssociateFirewallPolicyInput, ...request.Option) (*networkfirewall.AssociateFirewallPolicyOutput, error)
	AssociateFirewallPolicyRequest(*networkfirewall.AssociateFirewallPolicyInput) (*request.Request, *networkfirewall.AssociateFirewallPolicyOutput)

	AssociateSubnets(*networkfirewall.AssociateSubnetsInput) (*networkfirewall.AssociateSubnetsOutput, error)
	AssociateSubnetsWithContext(aws.Context, *networkfirewall.AssociateSubnetsInput, ...request.Option) (*networkfirewall.AssociateSubnetsOutput, error)
	AssociateSubnetsRequest(*networkfirewall.AssociateSubnetsInput) (*request.Request, *networkfirewall.AssociateSubnetsOutput)

	CreateFirewall(*networkfirewall.CreateFirewallInput) (*networkfirewall.CreateFirewallOutput, error)
	CreateFirewallWithContext(aws.Context, *networkfirewall.CreateFirewallInput, ...request.Option) (*networkfirewall.CreateFirewallOutput, error)
	CreateFirewallRequest(*networkfirewall.CreateFirewallInput) (*request.Request, *networkfirewall.CreateFirewallOutput)

	CreateFirewallPolicy(*networkfirewall.CreateFirewallPolicyInput) (*networkfirewall.CreateFirewallPolicyOutput, error)
	CreateFirewallPolicyWithContext(aws.Context, *networkfirewall.CreateFirewallPolicyInput, ...request.Option) (*networkfirewall.CreateFirewallPolicyOutput, error)
	CreateFirewallPolicyRequest(*networkfirewall.CreateFirewallPolicyInput) (*request.Request, *networkfirewall.CreateFirewallPolicyOutput)

	CreateRuleGroup(*networkfirewall.CreateRuleGroupInput) (*networkfirewall.CreateRuleGroupOutput, error)
	CreateRuleGroupWithContext(aws.Context, *networkfirewall.CreateRuleGroupInput, ...request.Option) (*networkfirewall.CreateRuleGroupOutput, error)
	CreateRuleGroupRequest(*networkfirewall.CreateRuleGroupInput) (*request.Request, *networkfirewall.CreateRuleGroupOutput)

	DeleteFirewall(*networkfirewall.DeleteFirewallInput) (*networkfirewall.DeleteFirewallOutput, error)
	DeleteFirewallWithContext(aws.Context, *networkfirewall.DeleteFirewallInput, ...request.Option) (*networkfirewall.DeleteFirewallOutput, error)
	DeleteFirewallRequest(*networkfirewall.DeleteFirewallInput) (*request.Request, *networkfirewall.DeleteFirewallOutput)

	DeleteFirewallPolicy(*networkfirewall.DeleteFirewallPolicyInput) (*networkfirewall.DeleteFirewallPolicyOutput, error)
	DeleteFirewallPolicyWithContext(aws.Context, *networkfirewall.DeleteFirewallPolicyInput, ...request.Option) (*networkfirewall.DeleteFirewallPolicyOutput, error)
	DeleteFirewallPolicyRequest(*networkfirewall.DeleteFirewallPolicyInput) (*request.Request, *networkfirewall.DeleteFirewallPolicyOutput)

	DeleteResourcePolicy(*networkfirewall.DeleteResourcePolicyInput) (*networkfirewall.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyWithContext(aws.Context, *networkfirewall.DeleteResourcePolicyInput, ...request.Option) (*networkfirewall.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyRequest(*networkfirewall.DeleteResourcePolicyInput) (*request.Request, *networkfirewall.DeleteResourcePolicyOutput)

	DeleteRuleGroup(*networkfirewall.DeleteRuleGroupInput) (*networkfirewall.DeleteRuleGroupOutput, error)
	DeleteRuleGroupWithContext(aws.Context, *networkfirewall.DeleteRuleGroupInput, ...request.Option) (*networkfirewall.DeleteRuleGroupOutput, error)
	DeleteRuleGroupRequest(*networkfirewall.DeleteRuleGroupInput) (*request.Request, *networkfirewall.DeleteRuleGroupOutput)

	DescribeFirewall(*networkfirewall.DescribeFirewallInput) (*networkfirewall.DescribeFirewallOutput, error)
	DescribeFirewallWithContext(aws.Context, *networkfirewall.DescribeFirewallInput, ...request.Option) (*networkfirewall.DescribeFirewallOutput, error)
	DescribeFirewallRequest(*networkfirewall.DescribeFirewallInput) (*request.Request, *networkfirewall.DescribeFirewallOutput)

	DescribeFirewallPolicy(*networkfirewall.DescribeFirewallPolicyInput) (*networkfirewall.DescribeFirewallPolicyOutput, error)
	DescribeFirewallPolicyWithContext(aws.Context, *networkfirewall.DescribeFirewallPolicyInput, ...request.Option) (*networkfirewall.DescribeFirewallPolicyOutput, error)
	DescribeFirewallPolicyRequest(*networkfirewall.DescribeFirewallPolicyInput) (*request.Request, *networkfirewall.DescribeFirewallPolicyOutput)

	DescribeLoggingConfiguration(*networkfirewall.DescribeLoggingConfigurationInput) (*networkfirewall.DescribeLoggingConfigurationOutput, error)
	DescribeLoggingConfigurationWithContext(aws.Context, *networkfirewall.DescribeLoggingConfigurationInput, ...request.Option) (*networkfirewall.DescribeLoggingConfigurationOutput, error)
	DescribeLoggingConfigurationRequest(*networkfirewall.DescribeLoggingConfigurationInput) (*request.Request, *networkfirewall.DescribeLoggingConfigurationOutput)

	DescribeResourcePolicy(*networkfirewall.DescribeResourcePolicyInput) (*networkfirewall.DescribeResourcePolicyOutput, error)
	DescribeResourcePolicyWithContext(aws.Context, *networkfirewall.DescribeResourcePolicyInput, ...request.Option) (*networkfirewall.DescribeResourcePolicyOutput, error)
	DescribeResourcePolicyRequest(*networkfirewall.DescribeResourcePolicyInput) (*request.Request, *networkfirewall.DescribeResourcePolicyOutput)

	DescribeRuleGroup(*networkfirewall.DescribeRuleGroupInput) (*networkfirewall.DescribeRuleGroupOutput, error)
	DescribeRuleGroupWithContext(aws.Context, *networkfirewall.DescribeRuleGroupInput, ...request.Option) (*networkfirewall.DescribeRuleGroupOutput, error)
	DescribeRuleGroupRequest(*networkfirewall.DescribeRuleGroupInput) (*request.Request, *networkfirewall.DescribeRuleGroupOutput)

	DescribeRuleGroupMetadata(*networkfirewall.DescribeRuleGroupMetadataInput) (*networkfirewall.DescribeRuleGroupMetadataOutput, error)
	DescribeRuleGroupMetadataWithContext(aws.Context, *networkfirewall.DescribeRuleGroupMetadataInput, ...request.Option) (*networkfirewall.DescribeRuleGroupMetadataOutput, error)
	DescribeRuleGroupMetadataRequest(*networkfirewall.DescribeRuleGroupMetadataInput) (*request.Request, *networkfirewall.DescribeRuleGroupMetadataOutput)

	DisassociateSubnets(*networkfirewall.DisassociateSubnetsInput) (*networkfirewall.DisassociateSubnetsOutput, error)
	DisassociateSubnetsWithContext(aws.Context, *networkfirewall.DisassociateSubnetsInput, ...request.Option) (*networkfirewall.DisassociateSubnetsOutput, error)
	DisassociateSubnetsRequest(*networkfirewall.DisassociateSubnetsInput) (*request.Request, *networkfirewall.DisassociateSubnetsOutput)

	ListFirewallPolicies(*networkfirewall.ListFirewallPoliciesInput) (*networkfirewall.ListFirewallPoliciesOutput, error)
	ListFirewallPoliciesWithContext(aws.Context, *networkfirewall.ListFirewallPoliciesInput, ...request.Option) (*networkfirewall.ListFirewallPoliciesOutput, error)
	ListFirewallPoliciesRequest(*networkfirewall.ListFirewallPoliciesInput) (*request.Request, *networkfirewall.ListFirewallPoliciesOutput)

	ListFirewallPoliciesPages(*networkfirewall.ListFirewallPoliciesInput, func(*networkfirewall.ListFirewallPoliciesOutput, bool) bool) error
	ListFirewallPoliciesPagesWithContext(aws.Context, *networkfirewall.ListFirewallPoliciesInput, func(*networkfirewall.ListFirewallPoliciesOutput, bool) bool, ...request.Option) error

	ListFirewalls(*networkfirewall.ListFirewallsInput) (*networkfirewall.ListFirewallsOutput, error)
	ListFirewallsWithContext(aws.Context, *networkfirewall.ListFirewallsInput, ...request.Option) (*networkfirewall.ListFirewallsOutput, error)
	ListFirewallsRequest(*networkfirewall.ListFirewallsInput) (*request.Request, *networkfirewall.ListFirewallsOutput)

	ListFirewallsPages(*networkfirewall.ListFirewallsInput, func(*networkfirewall.ListFirewallsOutput, bool) bool) error
	ListFirewallsPagesWithContext(aws.Context, *networkfirewall.ListFirewallsInput, func(*networkfirewall.ListFirewallsOutput, bool) bool, ...request.Option) error

	ListRuleGroups(*networkfirewall.ListRuleGroupsInput) (*networkfirewall.ListRuleGroupsOutput, error)
	ListRuleGroupsWithContext(aws.Context, *networkfirewall.ListRuleGroupsInput, ...request.Option) (*networkfirewall.ListRuleGroupsOutput, error)
	ListRuleGroupsRequest(*networkfirewall.ListRuleGroupsInput) (*request.Request, *networkfirewall.ListRuleGroupsOutput)

	ListRuleGroupsPages(*networkfirewall.ListRuleGroupsInput, func(*networkfirewall.ListRuleGroupsOutput, bool) bool) error
	ListRuleGroupsPagesWithContext(aws.Context, *networkfirewall.ListRuleGroupsInput, func(*networkfirewall.ListRuleGroupsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*networkfirewall.ListTagsForResourceInput) (*networkfirewall.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *networkfirewall.ListTagsForResourceInput, ...request.Option) (*networkfirewall.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*networkfirewall.ListTagsForResourceInput) (*request.Request, *networkfirewall.ListTagsForResourceOutput)

	ListTagsForResourcePages(*networkfirewall.ListTagsForResourceInput, func(*networkfirewall.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *networkfirewall.ListTagsForResourceInput, func(*networkfirewall.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	PutResourcePolicy(*networkfirewall.PutResourcePolicyInput) (*networkfirewall.PutResourcePolicyOutput, error)
	PutResourcePolicyWithContext(aws.Context, *networkfirewall.PutResourcePolicyInput, ...request.Option) (*networkfirewall.PutResourcePolicyOutput, error)
	PutResourcePolicyRequest(*networkfirewall.PutResourcePolicyInput) (*request.Request, *networkfirewall.PutResourcePolicyOutput)

	TagResource(*networkfirewall.TagResourceInput) (*networkfirewall.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *networkfirewall.TagResourceInput, ...request.Option) (*networkfirewall.TagResourceOutput, error)
	TagResourceRequest(*networkfirewall.TagResourceInput) (*request.Request, *networkfirewall.TagResourceOutput)

	UntagResource(*networkfirewall.UntagResourceInput) (*networkfirewall.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *networkfirewall.UntagResourceInput, ...request.Option) (*networkfirewall.UntagResourceOutput, error)
	UntagResourceRequest(*networkfirewall.UntagResourceInput) (*request.Request, *networkfirewall.UntagResourceOutput)

	UpdateFirewallDeleteProtection(*networkfirewall.UpdateFirewallDeleteProtectionInput) (*networkfirewall.UpdateFirewallDeleteProtectionOutput, error)
	UpdateFirewallDeleteProtectionWithContext(aws.Context, *networkfirewall.UpdateFirewallDeleteProtectionInput, ...request.Option) (*networkfirewall.UpdateFirewallDeleteProtectionOutput, error)
	UpdateFirewallDeleteProtectionRequest(*networkfirewall.UpdateFirewallDeleteProtectionInput) (*request.Request, *networkfirewall.UpdateFirewallDeleteProtectionOutput)

	UpdateFirewallDescription(*networkfirewall.UpdateFirewallDescriptionInput) (*networkfirewall.UpdateFirewallDescriptionOutput, error)
	UpdateFirewallDescriptionWithContext(aws.Context, *networkfirewall.UpdateFirewallDescriptionInput, ...request.Option) (*networkfirewall.UpdateFirewallDescriptionOutput, error)
	UpdateFirewallDescriptionRequest(*networkfirewall.UpdateFirewallDescriptionInput) (*request.Request, *networkfirewall.UpdateFirewallDescriptionOutput)

	UpdateFirewallEncryptionConfiguration(*networkfirewall.UpdateFirewallEncryptionConfigurationInput) (*networkfirewall.UpdateFirewallEncryptionConfigurationOutput, error)
	UpdateFirewallEncryptionConfigurationWithContext(aws.Context, *networkfirewall.UpdateFirewallEncryptionConfigurationInput, ...request.Option) (*networkfirewall.UpdateFirewallEncryptionConfigurationOutput, error)
	UpdateFirewallEncryptionConfigurationRequest(*networkfirewall.UpdateFirewallEncryptionConfigurationInput) (*request.Request, *networkfirewall.UpdateFirewallEncryptionConfigurationOutput)

	UpdateFirewallPolicy(*networkfirewall.UpdateFirewallPolicyInput) (*networkfirewall.UpdateFirewallPolicyOutput, error)
	UpdateFirewallPolicyWithContext(aws.Context, *networkfirewall.UpdateFirewallPolicyInput, ...request.Option) (*networkfirewall.UpdateFirewallPolicyOutput, error)
	UpdateFirewallPolicyRequest(*networkfirewall.UpdateFirewallPolicyInput) (*request.Request, *networkfirewall.UpdateFirewallPolicyOutput)

	UpdateFirewallPolicyChangeProtection(*networkfirewall.UpdateFirewallPolicyChangeProtectionInput) (*networkfirewall.UpdateFirewallPolicyChangeProtectionOutput, error)
	UpdateFirewallPolicyChangeProtectionWithContext(aws.Context, *networkfirewall.UpdateFirewallPolicyChangeProtectionInput, ...request.Option) (*networkfirewall.UpdateFirewallPolicyChangeProtectionOutput, error)
	UpdateFirewallPolicyChangeProtectionRequest(*networkfirewall.UpdateFirewallPolicyChangeProtectionInput) (*request.Request, *networkfirewall.UpdateFirewallPolicyChangeProtectionOutput)

	UpdateLoggingConfiguration(*networkfirewall.UpdateLoggingConfigurationInput) (*networkfirewall.UpdateLoggingConfigurationOutput, error)
	UpdateLoggingConfigurationWithContext(aws.Context, *networkfirewall.UpdateLoggingConfigurationInput, ...request.Option) (*networkfirewall.UpdateLoggingConfigurationOutput, error)
	UpdateLoggingConfigurationRequest(*networkfirewall.UpdateLoggingConfigurationInput) (*request.Request, *networkfirewall.UpdateLoggingConfigurationOutput)

	UpdateRuleGroup(*networkfirewall.UpdateRuleGroupInput) (*networkfirewall.UpdateRuleGroupOutput, error)
	UpdateRuleGroupWithContext(aws.Context, *networkfirewall.UpdateRuleGroupInput, ...request.Option) (*networkfirewall.UpdateRuleGroupOutput, error)
	UpdateRuleGroupRequest(*networkfirewall.UpdateRuleGroupInput) (*request.Request, *networkfirewall.UpdateRuleGroupOutput)

	UpdateSubnetChangeProtection(*networkfirewall.UpdateSubnetChangeProtectionInput) (*networkfirewall.UpdateSubnetChangeProtectionOutput, error)
	UpdateSubnetChangeProtectionWithContext(aws.Context, *networkfirewall.UpdateSubnetChangeProtectionInput, ...request.Option) (*networkfirewall.UpdateSubnetChangeProtectionOutput, error)
	UpdateSubnetChangeProtectionRequest(*networkfirewall.UpdateSubnetChangeProtectionInput) (*request.Request, *networkfirewall.UpdateSubnetChangeProtectionOutput)
}

var _ NetworkFirewallAPI = (*networkfirewall.NetworkFirewall)(nil)
