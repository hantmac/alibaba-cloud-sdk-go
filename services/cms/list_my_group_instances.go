package cms

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

func (client *Client) ListMyGroupInstances(request *ListMyGroupInstancesRequest) (response *ListMyGroupInstancesResponse, err error) {
	response = CreateListMyGroupInstancesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListMyGroupInstancesWithChan(request *ListMyGroupInstancesRequest) (<-chan *ListMyGroupInstancesResponse, <-chan error) {
	responseChan := make(chan *ListMyGroupInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMyGroupInstances(request)
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

func (client *Client) ListMyGroupInstancesWithCallback(request *ListMyGroupInstancesRequest, callback func(response *ListMyGroupInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMyGroupInstancesResponse
		var err error
		defer close(result)
		response, err = client.ListMyGroupInstances(request)
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

type ListMyGroupInstancesRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Total      requests.Boolean `position:"Query" name:"Total"`
	GroupId    requests.Integer `position:"Query" name:"GroupId"`
	Category   string           `position:"Query" name:"Category"`
}

type ListMyGroupInstancesResponse struct {
	*responses.BaseResponse
	RequestId    string                          `json:"RequestId" xml:"RequestId"`
	Success      bool                            `json:"Success" xml:"Success"`
	ErrorCode    int                             `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string                          `json:"ErrorMessage" xml:"ErrorMessage"`
	PageNumber   int                             `json:"PageNumber" xml:"PageNumber"`
	PageSize     int                             `json:"PageSize" xml:"PageSize"`
	Total        int                             `json:"Total" xml:"Total"`
	Resources    ResourcesInListMyGroupInstances `json:"Resources" xml:"Resources"`
}

func CreateListMyGroupInstancesRequest() (request *ListMyGroupInstancesRequest) {
	request = &ListMyGroupInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "ListMyGroupInstances", "cms", "openAPI")
	return
}

func CreateListMyGroupInstancesResponse() (response *ListMyGroupInstancesResponse) {
	response = &ListMyGroupInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
