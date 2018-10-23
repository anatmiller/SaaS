/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ess

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/alicloud/alibaba-cloud-sdk-go/sdk/requests"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/alicloud/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeScalingRules invokes the ess.DescribeScalingRules API synchronously
// api document: https://help.aliyun.com/api/ess/describescalingrules.html
func (client *Client) DescribeScalingRules(request *DescribeScalingRulesRequest) (response *DescribeScalingRulesResponse, err error) {
	response = CreateDescribeScalingRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScalingRulesWithChan invokes the ess.DescribeScalingRules API asynchronously
// api document: https://help.aliyun.com/api/ess/describescalingrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScalingRulesWithChan(request *DescribeScalingRulesRequest) (<-chan *DescribeScalingRulesResponse, <-chan error) {
	responseChan := make(chan *DescribeScalingRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScalingRules(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeScalingRulesWithCallback invokes the ess.DescribeScalingRules API asynchronously
// api document: https://help.aliyun.com/api/ess/describescalingrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScalingRulesWithCallback(request *DescribeScalingRulesRequest, callback func(response *DescribeScalingRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScalingRulesResponse
		var err error
		defer close(result)
		response, err = client.DescribeScalingRules(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeScalingRulesRequest is the request struct for api DescribeScalingRules
type DescribeScalingRulesRequest struct {
	*requests.RpcRequest
	ScalingRuleName1     string           `position:"Query" name:"ScalingRuleName.1"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ScalingRuleName2     string           `position:"Query" name:"ScalingRuleName.2"`
	ScalingRuleName3     string           `position:"Query" name:"ScalingRuleName.3"`
	ScalingRuleName4     string           `position:"Query" name:"ScalingRuleName.4"`
	ScalingRuleName5     string           `position:"Query" name:"ScalingRuleName.5"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	ScalingRuleName6     string           `position:"Query" name:"ScalingRuleName.6"`
	ScalingRuleName7     string           `position:"Query" name:"ScalingRuleName.7"`
	ScalingRuleName8     string           `position:"Query" name:"ScalingRuleName.8"`
	ScalingRuleAri9      string           `position:"Query" name:"ScalingRuleAri.9"`
	ScalingRuleName9     string           `position:"Query" name:"ScalingRuleName.9"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ScalingRuleId10      string           `position:"Query" name:"ScalingRuleId.10"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ScalingRuleAri1      string           `position:"Query" name:"ScalingRuleAri.1"`
	ScalingRuleAri2      string           `position:"Query" name:"ScalingRuleAri.2"`
	ScalingRuleName10    string           `position:"Query" name:"ScalingRuleName.10"`
	ScalingRuleAri3      string           `position:"Query" name:"ScalingRuleAri.3"`
	ScalingRuleAri4      string           `position:"Query" name:"ScalingRuleAri.4"`
	ScalingRuleId8       string           `position:"Query" name:"ScalingRuleId.8"`
	ScalingRuleAri5      string           `position:"Query" name:"ScalingRuleAri.5"`
	ScalingRuleId9       string           `position:"Query" name:"ScalingRuleId.9"`
	ScalingRuleAri6      string           `position:"Query" name:"ScalingRuleAri.6"`
	ScalingRuleAri7      string           `position:"Query" name:"ScalingRuleAri.7"`
	ScalingRuleAri10     string           `position:"Query" name:"ScalingRuleAri.10"`
	ScalingRuleAri8      string           `position:"Query" name:"ScalingRuleAri.8"`
	ScalingRuleId4       string           `position:"Query" name:"ScalingRuleId.4"`
	ScalingRuleId5       string           `position:"Query" name:"ScalingRuleId.5"`
	ScalingRuleId6       string           `position:"Query" name:"ScalingRuleId.6"`
	ScalingRuleId7       string           `position:"Query" name:"ScalingRuleId.7"`
	ScalingRuleId1       string           `position:"Query" name:"ScalingRuleId.1"`
	ScalingRuleId2       string           `position:"Query" name:"ScalingRuleId.2"`
	ScalingRuleId3       string           `position:"Query" name:"ScalingRuleId.3"`
}

// DescribeScalingRulesResponse is the response struct for api DescribeScalingRules
type DescribeScalingRulesResponse struct {
	*responses.BaseResponse
	TotalCount   int          `json:"TotalCount" xml:"TotalCount"`
	PageNumber   int          `json:"PageNumber" xml:"PageNumber"`
	PageSize     int          `json:"PageSize" xml:"PageSize"`
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	ScalingRules ScalingRules `json:"ScalingRules" xml:"ScalingRules"`
}

// CreateDescribeScalingRulesRequest creates a request to invoke DescribeScalingRules API
func CreateDescribeScalingRulesRequest() (request *DescribeScalingRulesRequest) {
	request = &DescribeScalingRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DescribeScalingRules", "ess", "openAPI")
	return
}

// CreateDescribeScalingRulesResponse creates a response to parse from DescribeScalingRules response
func CreateDescribeScalingRulesResponse() (response *DescribeScalingRulesResponse) {
	response = &DescribeScalingRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

// ScalingRules is a nested struct in ess response
type ScalingRules struct {
	ScalingRule []ScalingRule `json:"ScalingRule" xml:"ScalingRule"`
}

// ScalingRule is a nested struct in ess response
type ScalingRule struct {
	ScalingRuleId   string `json:"ScalingRuleId" xml:"ScalingRuleId"`
	ScalingGroupId  string `json:"ScalingGroupId" xml:"ScalingGroupId"`
	ScalingRuleName string `json:"ScalingRuleName" xml:"ScalingRuleName"`
	Cooldown        int    `json:"Cooldown" xml:"Cooldown"`
	AdjustmentType  string `json:"AdjustmentType" xml:"AdjustmentType"`
	AdjustmentValue int    `json:"AdjustmentValue" xml:"AdjustmentValue"`
	MinSize         int    `json:"MinSize" xml:"MinSize"`
	MaxSize         int    `json:"MaxSize" xml:"MaxSize"`
	ScalingRuleAri  string `json:"ScalingRuleAri" xml:"ScalingRuleAri"`
}
