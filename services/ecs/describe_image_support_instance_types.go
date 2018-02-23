package ecs

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

func (client *Client) DescribeImageSupportInstanceTypes(request *DescribeImageSupportInstanceTypesRequest) (response *DescribeImageSupportInstanceTypesResponse, err error) {
	response = CreateDescribeImageSupportInstanceTypesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeImageSupportInstanceTypesWithChan(request *DescribeImageSupportInstanceTypesRequest) (<-chan *DescribeImageSupportInstanceTypesResponse, <-chan error) {
	responseChan := make(chan *DescribeImageSupportInstanceTypesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImageSupportInstanceTypes(request)
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

func (client *Client) DescribeImageSupportInstanceTypesWithCallback(request *DescribeImageSupportInstanceTypesRequest, callback func(response *DescribeImageSupportInstanceTypesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImageSupportInstanceTypesResponse
		var err error
		defer close(result)
		response, err = client.DescribeImageSupportInstanceTypes(request)
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

type DescribeImageSupportInstanceTypesRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer                           `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string                                     `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer                           `position:"Query" name:"ResourceOwnerId"`
	ImageId              string                                     `position:"Query" name:"ImageId"`
	ActionType           string                                     `position:"Query" name:"ActionType"`
	Filter               *[]DescribeImageSupportInstanceTypesFilter `position:"Query" name:"Filter"  type:"Repeated"`
}

type DescribeImageSupportInstanceTypesFilter struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type DescribeImageSupportInstanceTypesResponse struct {
	*responses.BaseResponse
	RequestId     string                                           `json:"RequestId" xml:"RequestId"`
	RegionId      string                                           `json:"RegionId" xml:"RegionId"`
	ImageId       string                                           `json:"ImageId" xml:"ImageId"`
	InstanceTypes InstanceTypesInDescribeImageSupportInstanceTypes `json:"InstanceTypes" xml:"InstanceTypes"`
}

func CreateDescribeImageSupportInstanceTypesRequest() (request *DescribeImageSupportInstanceTypesRequest) {
	request = &DescribeImageSupportInstanceTypesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeImageSupportInstanceTypes", "ecs", "openAPI")
	return
}

func CreateDescribeImageSupportInstanceTypesResponse() (response *DescribeImageSupportInstanceTypesResponse) {
	response = &DescribeImageSupportInstanceTypesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
