package rds

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

// EvaluateFailover invokes the rds.EvaluateFailover API synchronously
// api document: https://help.aliyun.com/api/rds/evaluatefailover.html
func (client *Client) EvaluateFailover(request *EvaluateFailoverRequest) (response *EvaluateFailoverResponse, err error) {
	response = CreateEvaluateFailoverResponse()
	err = client.DoAction(request, response)
	return
}

// EvaluateFailoverWithChan invokes the rds.EvaluateFailover API asynchronously
// api document: https://help.aliyun.com/api/rds/evaluatefailover.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EvaluateFailoverWithChan(request *EvaluateFailoverRequest) (<-chan *EvaluateFailoverResponse, <-chan error) {
	responseChan := make(chan *EvaluateFailoverResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EvaluateFailover(request)
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

// EvaluateFailoverWithCallback invokes the rds.EvaluateFailover API asynchronously
// api document: https://help.aliyun.com/api/rds/evaluatefailover.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EvaluateFailoverWithCallback(request *EvaluateFailoverRequest, callback func(response *EvaluateFailoverResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EvaluateFailoverResponse
		var err error
		defer close(result)
		response, err = client.EvaluateFailover(request)
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

// EvaluateFailoverRequest is the request struct for api EvaluateFailover
type EvaluateFailoverRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ReplicaId            string           `position:"Query" name:"ReplicaId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// EvaluateFailoverResponse is the response struct for api EvaluateFailover
type EvaluateFailoverResponse struct {
	*responses.BaseResponse
	RequestId      string        `json:"RequestId" xml:"RequestId"`
	ReplicaId      string        `json:"ReplicaId" xml:"ReplicaId"`
	EvaluateResult string        `json:"EvaluateResult" xml:"EvaluateResult"`
	Items          []ItemsItem   `json:"Items" xml:"Items"`
	Reports        []ReportsItem `json:"Reports" xml:"Reports"`
}

// CreateEvaluateFailoverRequest creates a request to invoke EvaluateFailover API
func CreateEvaluateFailoverRequest() (request *EvaluateFailoverRequest) {
	request = &EvaluateFailoverRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "EvaluateFailover", "rds", "openAPI")
	return
}

// CreateEvaluateFailoverResponse creates a response to parse from EvaluateFailover response
func CreateEvaluateFailoverResponse() (response *EvaluateFailoverResponse) {
	response = &EvaluateFailoverResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
