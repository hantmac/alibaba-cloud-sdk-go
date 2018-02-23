package ess

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) CreateScalingGroup(request *CreateScalingGroupRequest) (response *CreateScalingGroupResponse, err error) {
	response = CreateCreateScalingGroupResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateScalingGroupWithChan(request *CreateScalingGroupRequest) (<-chan *CreateScalingGroupResponse, <-chan error) {
	responseChan := make(chan *CreateScalingGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateScalingGroup(request)
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

func (client *Client) CreateScalingGroupWithCallback(request *CreateScalingGroupRequest, callback func(response *CreateScalingGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateScalingGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateScalingGroup(request)
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

type CreateScalingGroupRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupName     string           `position:"Query" name:"ScalingGroupName"`
	MinSize              requests.Integer `position:"Query" name:"MinSize"`
	MaxSize              requests.Integer `position:"Query" name:"MaxSize"`
	DefaultCooldown      requests.Integer `position:"Query" name:"DefaultCooldown"`
	LoadBalancerIds      string           `position:"Query" name:"LoadBalancerIds"`
	DBInstanceIds        string           `position:"Query" name:"DBInstanceIds"`
	RemovalPolicy1       string           `position:"Query" name:"RemovalPolicy.1"`
	RemovalPolicy2       string           `position:"Query" name:"RemovalPolicy.2"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	VSwitchIds           *[]string        `position:"Query" name:"VSwitchIds"  type:"Repeated"`
}

type CreateScalingGroupResponse struct {
	*responses.BaseResponse
	ScalingGroupId string `json:"ScalingGroupId" xml:"ScalingGroupId"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

func CreateCreateScalingGroupRequest() (request *CreateScalingGroupRequest) {
	request = &CreateScalingGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "CreateScalingGroup", "ess", "openAPI")
	return
}

func CreateCreateScalingGroupResponse() (response *CreateScalingGroupResponse) {
	response = &CreateScalingGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
