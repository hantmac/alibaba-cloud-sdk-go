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

func (client *Client) DeleteAlarm(request *DeleteAlarmRequest) (response *DeleteAlarmResponse, err error) {
	response = CreateDeleteAlarmResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteAlarmWithChan(request *DeleteAlarmRequest) (<-chan *DeleteAlarmResponse, <-chan error) {
	responseChan := make(chan *DeleteAlarmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAlarm(request)
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

func (client *Client) DeleteAlarmWithCallback(request *DeleteAlarmRequest, callback func(response *DeleteAlarmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAlarmResponse
		var err error
		defer close(result)
		response, err = client.DeleteAlarm(request)
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

type DeleteAlarmRequest struct {
	*requests.RpcRequest
	CallbyCmsOwner string `position:"Query" name:"callby_cms_owner"`
	Id             string `position:"Query" name:"Id"`
}

type DeleteAlarmResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteAlarmRequest() (request *DeleteAlarmRequest) {
	request = &DeleteAlarmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "DeleteAlarm", "cms", "openAPI")
	return
}

func CreateDeleteAlarmResponse() (response *DeleteAlarmResponse) {
	response = &DeleteAlarmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
