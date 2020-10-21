package mts

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

// QuerySubtitleJobList invokes the mts.QuerySubtitleJobList API synchronously
func (client *Client) QuerySubtitleJobList(request *QuerySubtitleJobListRequest) (response *QuerySubtitleJobListResponse, err error) {
	response = CreateQuerySubtitleJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySubtitleJobListWithChan invokes the mts.QuerySubtitleJobList API asynchronously
func (client *Client) QuerySubtitleJobListWithChan(request *QuerySubtitleJobListRequest) (<-chan *QuerySubtitleJobListResponse, <-chan error) {
	responseChan := make(chan *QuerySubtitleJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySubtitleJobList(request)
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

// QuerySubtitleJobListWithCallback invokes the mts.QuerySubtitleJobList API asynchronously
func (client *Client) QuerySubtitleJobListWithCallback(request *QuerySubtitleJobListRequest, callback func(response *QuerySubtitleJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySubtitleJobListResponse
		var err error
		defer close(result)
		response, err = client.QuerySubtitleJobList(request)
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

// QuerySubtitleJobListRequest is the request struct for api QuerySubtitleJobList
type QuerySubtitleJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	JobIds               string           `position:"Query" name:"JobIds"`
}

// QuerySubtitleJobListResponse is the response struct for api QuerySubtitleJobList
type QuerySubtitleJobListResponse struct {
	*responses.BaseResponse
	RequestId      string                               `json:"RequestId" xml:"RequestId"`
	NonExistJobIds NonExistJobIdsInQuerySubtitleJobList `json:"NonExistJobIds" xml:"NonExistJobIds"`
	JobList        JobListInQuerySubtitleJobList        `json:"JobList" xml:"JobList"`
}

// CreateQuerySubtitleJobListRequest creates a request to invoke QuerySubtitleJobList API
func CreateQuerySubtitleJobListRequest() (request *QuerySubtitleJobListRequest) {
	request = &QuerySubtitleJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QuerySubtitleJobList", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySubtitleJobListResponse creates a response to parse from QuerySubtitleJobList response
func CreateQuerySubtitleJobListResponse() (response *QuerySubtitleJobListResponse) {
	response = &QuerySubtitleJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
