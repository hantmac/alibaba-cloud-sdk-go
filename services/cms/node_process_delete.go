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

func (client *Client) NodeProcessDelete(request *NodeProcessDeleteRequest) (response *NodeProcessDeleteResponse, err error) {
	response = CreateNodeProcessDeleteResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) NodeProcessDeleteWithChan(request *NodeProcessDeleteRequest) (<-chan *NodeProcessDeleteResponse, <-chan error) {
	responseChan := make(chan *NodeProcessDeleteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.NodeProcessDelete(request)
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

func (client *Client) NodeProcessDeleteWithCallback(request *NodeProcessDeleteRequest, callback func(response *NodeProcessDeleteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *NodeProcessDeleteResponse
		var err error
		defer close(result)
		response, err = client.NodeProcessDelete(request)
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

type NodeProcessDeleteRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Name       string `position:"Query" name:"Name"`
	Id         string `position:"Query" name:"Id"`
}

type NodeProcessDeleteResponse struct {
	*responses.BaseResponse
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

func CreateNodeProcessDeleteRequest() (request *NodeProcessDeleteRequest) {
	request = &NodeProcessDeleteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "NodeProcessDelete", "cms", "openAPI")
	return
}

func CreateNodeProcessDeleteResponse() (response *NodeProcessDeleteResponse) {
	response = &NodeProcessDeleteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
