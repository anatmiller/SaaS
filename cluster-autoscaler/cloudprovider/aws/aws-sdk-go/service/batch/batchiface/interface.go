// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package batchiface provides an interface to enable mocking the AWS Batch service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package batchiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/batch"
)

// BatchAPI provides an interface to enable mocking the
// batch.Batch service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Batch.
//	func myFunc(svc batchiface.BatchAPI) bool {
//	    // Make svc.CancelJob request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := batch.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockBatchClient struct {
//	    batchiface.BatchAPI
//	}
//	func (m *mockBatchClient) CancelJob(input *batch.CancelJobInput) (*batch.CancelJobOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockBatchClient{}
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
type BatchAPI interface {
	CancelJob(*batch.CancelJobInput) (*batch.CancelJobOutput, error)
	CancelJobWithContext(aws.Context, *batch.CancelJobInput, ...request.Option) (*batch.CancelJobOutput, error)
	CancelJobRequest(*batch.CancelJobInput) (*request.Request, *batch.CancelJobOutput)

	CreateComputeEnvironment(*batch.CreateComputeEnvironmentInput) (*batch.CreateComputeEnvironmentOutput, error)
	CreateComputeEnvironmentWithContext(aws.Context, *batch.CreateComputeEnvironmentInput, ...request.Option) (*batch.CreateComputeEnvironmentOutput, error)
	CreateComputeEnvironmentRequest(*batch.CreateComputeEnvironmentInput) (*request.Request, *batch.CreateComputeEnvironmentOutput)

	CreateJobQueue(*batch.CreateJobQueueInput) (*batch.CreateJobQueueOutput, error)
	CreateJobQueueWithContext(aws.Context, *batch.CreateJobQueueInput, ...request.Option) (*batch.CreateJobQueueOutput, error)
	CreateJobQueueRequest(*batch.CreateJobQueueInput) (*request.Request, *batch.CreateJobQueueOutput)

	CreateSchedulingPolicy(*batch.CreateSchedulingPolicyInput) (*batch.CreateSchedulingPolicyOutput, error)
	CreateSchedulingPolicyWithContext(aws.Context, *batch.CreateSchedulingPolicyInput, ...request.Option) (*batch.CreateSchedulingPolicyOutput, error)
	CreateSchedulingPolicyRequest(*batch.CreateSchedulingPolicyInput) (*request.Request, *batch.CreateSchedulingPolicyOutput)

	DeleteComputeEnvironment(*batch.DeleteComputeEnvironmentInput) (*batch.DeleteComputeEnvironmentOutput, error)
	DeleteComputeEnvironmentWithContext(aws.Context, *batch.DeleteComputeEnvironmentInput, ...request.Option) (*batch.DeleteComputeEnvironmentOutput, error)
	DeleteComputeEnvironmentRequest(*batch.DeleteComputeEnvironmentInput) (*request.Request, *batch.DeleteComputeEnvironmentOutput)

	DeleteJobQueue(*batch.DeleteJobQueueInput) (*batch.DeleteJobQueueOutput, error)
	DeleteJobQueueWithContext(aws.Context, *batch.DeleteJobQueueInput, ...request.Option) (*batch.DeleteJobQueueOutput, error)
	DeleteJobQueueRequest(*batch.DeleteJobQueueInput) (*request.Request, *batch.DeleteJobQueueOutput)

	DeleteSchedulingPolicy(*batch.DeleteSchedulingPolicyInput) (*batch.DeleteSchedulingPolicyOutput, error)
	DeleteSchedulingPolicyWithContext(aws.Context, *batch.DeleteSchedulingPolicyInput, ...request.Option) (*batch.DeleteSchedulingPolicyOutput, error)
	DeleteSchedulingPolicyRequest(*batch.DeleteSchedulingPolicyInput) (*request.Request, *batch.DeleteSchedulingPolicyOutput)

	DeregisterJobDefinition(*batch.DeregisterJobDefinitionInput) (*batch.DeregisterJobDefinitionOutput, error)
	DeregisterJobDefinitionWithContext(aws.Context, *batch.DeregisterJobDefinitionInput, ...request.Option) (*batch.DeregisterJobDefinitionOutput, error)
	DeregisterJobDefinitionRequest(*batch.DeregisterJobDefinitionInput) (*request.Request, *batch.DeregisterJobDefinitionOutput)

	DescribeComputeEnvironments(*batch.DescribeComputeEnvironmentsInput) (*batch.DescribeComputeEnvironmentsOutput, error)
	DescribeComputeEnvironmentsWithContext(aws.Context, *batch.DescribeComputeEnvironmentsInput, ...request.Option) (*batch.DescribeComputeEnvironmentsOutput, error)
	DescribeComputeEnvironmentsRequest(*batch.DescribeComputeEnvironmentsInput) (*request.Request, *batch.DescribeComputeEnvironmentsOutput)

	DescribeComputeEnvironmentsPages(*batch.DescribeComputeEnvironmentsInput, func(*batch.DescribeComputeEnvironmentsOutput, bool) bool) error
	DescribeComputeEnvironmentsPagesWithContext(aws.Context, *batch.DescribeComputeEnvironmentsInput, func(*batch.DescribeComputeEnvironmentsOutput, bool) bool, ...request.Option) error

	DescribeJobDefinitions(*batch.DescribeJobDefinitionsInput) (*batch.DescribeJobDefinitionsOutput, error)
	DescribeJobDefinitionsWithContext(aws.Context, *batch.DescribeJobDefinitionsInput, ...request.Option) (*batch.DescribeJobDefinitionsOutput, error)
	DescribeJobDefinitionsRequest(*batch.DescribeJobDefinitionsInput) (*request.Request, *batch.DescribeJobDefinitionsOutput)

	DescribeJobDefinitionsPages(*batch.DescribeJobDefinitionsInput, func(*batch.DescribeJobDefinitionsOutput, bool) bool) error
	DescribeJobDefinitionsPagesWithContext(aws.Context, *batch.DescribeJobDefinitionsInput, func(*batch.DescribeJobDefinitionsOutput, bool) bool, ...request.Option) error

	DescribeJobQueues(*batch.DescribeJobQueuesInput) (*batch.DescribeJobQueuesOutput, error)
	DescribeJobQueuesWithContext(aws.Context, *batch.DescribeJobQueuesInput, ...request.Option) (*batch.DescribeJobQueuesOutput, error)
	DescribeJobQueuesRequest(*batch.DescribeJobQueuesInput) (*request.Request, *batch.DescribeJobQueuesOutput)

	DescribeJobQueuesPages(*batch.DescribeJobQueuesInput, func(*batch.DescribeJobQueuesOutput, bool) bool) error
	DescribeJobQueuesPagesWithContext(aws.Context, *batch.DescribeJobQueuesInput, func(*batch.DescribeJobQueuesOutput, bool) bool, ...request.Option) error

	DescribeJobs(*batch.DescribeJobsInput) (*batch.DescribeJobsOutput, error)
	DescribeJobsWithContext(aws.Context, *batch.DescribeJobsInput, ...request.Option) (*batch.DescribeJobsOutput, error)
	DescribeJobsRequest(*batch.DescribeJobsInput) (*request.Request, *batch.DescribeJobsOutput)

	DescribeSchedulingPolicies(*batch.DescribeSchedulingPoliciesInput) (*batch.DescribeSchedulingPoliciesOutput, error)
	DescribeSchedulingPoliciesWithContext(aws.Context, *batch.DescribeSchedulingPoliciesInput, ...request.Option) (*batch.DescribeSchedulingPoliciesOutput, error)
	DescribeSchedulingPoliciesRequest(*batch.DescribeSchedulingPoliciesInput) (*request.Request, *batch.DescribeSchedulingPoliciesOutput)

	ListJobs(*batch.ListJobsInput) (*batch.ListJobsOutput, error)
	ListJobsWithContext(aws.Context, *batch.ListJobsInput, ...request.Option) (*batch.ListJobsOutput, error)
	ListJobsRequest(*batch.ListJobsInput) (*request.Request, *batch.ListJobsOutput)

	ListJobsPages(*batch.ListJobsInput, func(*batch.ListJobsOutput, bool) bool) error
	ListJobsPagesWithContext(aws.Context, *batch.ListJobsInput, func(*batch.ListJobsOutput, bool) bool, ...request.Option) error

	ListSchedulingPolicies(*batch.ListSchedulingPoliciesInput) (*batch.ListSchedulingPoliciesOutput, error)
	ListSchedulingPoliciesWithContext(aws.Context, *batch.ListSchedulingPoliciesInput, ...request.Option) (*batch.ListSchedulingPoliciesOutput, error)
	ListSchedulingPoliciesRequest(*batch.ListSchedulingPoliciesInput) (*request.Request, *batch.ListSchedulingPoliciesOutput)

	ListSchedulingPoliciesPages(*batch.ListSchedulingPoliciesInput, func(*batch.ListSchedulingPoliciesOutput, bool) bool) error
	ListSchedulingPoliciesPagesWithContext(aws.Context, *batch.ListSchedulingPoliciesInput, func(*batch.ListSchedulingPoliciesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*batch.ListTagsForResourceInput) (*batch.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *batch.ListTagsForResourceInput, ...request.Option) (*batch.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*batch.ListTagsForResourceInput) (*request.Request, *batch.ListTagsForResourceOutput)

	RegisterJobDefinition(*batch.RegisterJobDefinitionInput) (*batch.RegisterJobDefinitionOutput, error)
	RegisterJobDefinitionWithContext(aws.Context, *batch.RegisterJobDefinitionInput, ...request.Option) (*batch.RegisterJobDefinitionOutput, error)
	RegisterJobDefinitionRequest(*batch.RegisterJobDefinitionInput) (*request.Request, *batch.RegisterJobDefinitionOutput)

	SubmitJob(*batch.SubmitJobInput) (*batch.SubmitJobOutput, error)
	SubmitJobWithContext(aws.Context, *batch.SubmitJobInput, ...request.Option) (*batch.SubmitJobOutput, error)
	SubmitJobRequest(*batch.SubmitJobInput) (*request.Request, *batch.SubmitJobOutput)

	TagResource(*batch.TagResourceInput) (*batch.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *batch.TagResourceInput, ...request.Option) (*batch.TagResourceOutput, error)
	TagResourceRequest(*batch.TagResourceInput) (*request.Request, *batch.TagResourceOutput)

	TerminateJob(*batch.TerminateJobInput) (*batch.TerminateJobOutput, error)
	TerminateJobWithContext(aws.Context, *batch.TerminateJobInput, ...request.Option) (*batch.TerminateJobOutput, error)
	TerminateJobRequest(*batch.TerminateJobInput) (*request.Request, *batch.TerminateJobOutput)

	UntagResource(*batch.UntagResourceInput) (*batch.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *batch.UntagResourceInput, ...request.Option) (*batch.UntagResourceOutput, error)
	UntagResourceRequest(*batch.UntagResourceInput) (*request.Request, *batch.UntagResourceOutput)

	UpdateComputeEnvironment(*batch.UpdateComputeEnvironmentInput) (*batch.UpdateComputeEnvironmentOutput, error)
	UpdateComputeEnvironmentWithContext(aws.Context, *batch.UpdateComputeEnvironmentInput, ...request.Option) (*batch.UpdateComputeEnvironmentOutput, error)
	UpdateComputeEnvironmentRequest(*batch.UpdateComputeEnvironmentInput) (*request.Request, *batch.UpdateComputeEnvironmentOutput)

	UpdateJobQueue(*batch.UpdateJobQueueInput) (*batch.UpdateJobQueueOutput, error)
	UpdateJobQueueWithContext(aws.Context, *batch.UpdateJobQueueInput, ...request.Option) (*batch.UpdateJobQueueOutput, error)
	UpdateJobQueueRequest(*batch.UpdateJobQueueInput) (*request.Request, *batch.UpdateJobQueueOutput)

	UpdateSchedulingPolicy(*batch.UpdateSchedulingPolicyInput) (*batch.UpdateSchedulingPolicyOutput, error)
	UpdateSchedulingPolicyWithContext(aws.Context, *batch.UpdateSchedulingPolicyInput, ...request.Option) (*batch.UpdateSchedulingPolicyOutput, error)
	UpdateSchedulingPolicyRequest(*batch.UpdateSchedulingPolicyInput) (*request.Request, *batch.UpdateSchedulingPolicyOutput)
}

var _ BatchAPI = (*batch.Batch)(nil)
