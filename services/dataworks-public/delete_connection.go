package dataworks_public

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

// DeleteConnection invokes the dataworks_public.DeleteConnection API synchronously
func (client *Client) DeleteConnection(request *DeleteConnectionRequest) (response *DeleteConnectionResponse, err error) {
	response = CreateDeleteConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteConnectionWithChan invokes the dataworks_public.DeleteConnection API asynchronously
func (client *Client) DeleteConnectionWithChan(request *DeleteConnectionRequest) (<-chan *DeleteConnectionResponse, <-chan error) {
	responseChan := make(chan *DeleteConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteConnection(request)
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

// DeleteConnectionWithCallback invokes the dataworks_public.DeleteConnection API asynchronously
func (client *Client) DeleteConnectionWithCallback(request *DeleteConnectionRequest, callback func(response *DeleteConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteConnectionResponse
		var err error
		defer close(result)
		response, err = client.DeleteConnection(request)
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

// DeleteConnectionRequest is the request struct for api DeleteConnection
type DeleteConnectionRequest struct {
	*requests.RpcRequest
	ConnectionId requests.Integer `position:"Query" name:"ConnectionId"`
}

// DeleteConnectionResponse is the response struct for api DeleteConnection
type DeleteConnectionResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	HttpStatusCode string `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           bool   `json:"Data" xml:"Data"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteConnectionRequest creates a request to invoke DeleteConnection API
func CreateDeleteConnectionRequest() (request *DeleteConnectionRequest) {
	request = &DeleteConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "DeleteConnection", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteConnectionResponse creates a response to parse from DeleteConnection response
func CreateDeleteConnectionResponse() (response *DeleteConnectionResponse) {
	response = &DeleteConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}