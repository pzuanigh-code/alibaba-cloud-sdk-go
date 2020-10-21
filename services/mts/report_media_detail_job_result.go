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

// ReportMediaDetailJobResult invokes the mts.ReportMediaDetailJobResult API synchronously
func (client *Client) ReportMediaDetailJobResult(request *ReportMediaDetailJobResultRequest) (response *ReportMediaDetailJobResultResponse, err error) {
	response = CreateReportMediaDetailJobResultResponse()
	err = client.DoAction(request, response)
	return
}

// ReportMediaDetailJobResultWithChan invokes the mts.ReportMediaDetailJobResult API asynchronously
func (client *Client) ReportMediaDetailJobResultWithChan(request *ReportMediaDetailJobResultRequest) (<-chan *ReportMediaDetailJobResultResponse, <-chan error) {
	responseChan := make(chan *ReportMediaDetailJobResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportMediaDetailJobResult(request)
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

// ReportMediaDetailJobResultWithCallback invokes the mts.ReportMediaDetailJobResult API asynchronously
func (client *Client) ReportMediaDetailJobResultWithCallback(request *ReportMediaDetailJobResultRequest, callback func(response *ReportMediaDetailJobResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportMediaDetailJobResultResponse
		var err error
		defer close(result)
		response, err = client.ReportMediaDetailJobResult(request)
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

// ReportMediaDetailJobResultRequest is the request struct for api ReportMediaDetailJobResult
type ReportMediaDetailJobResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JobId                string           `position:"Query" name:"JobId"`
	Tag                  string           `position:"Query" name:"Tag"`
	Results              string           `position:"Query" name:"Results"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ReportMediaDetailJobResultResponse is the response struct for api ReportMediaDetailJobResult
type ReportMediaDetailJobResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateReportMediaDetailJobResultRequest creates a request to invoke ReportMediaDetailJobResult API
func CreateReportMediaDetailJobResultRequest() (request *ReportMediaDetailJobResultRequest) {
	request = &ReportMediaDetailJobResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ReportMediaDetailJobResult", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReportMediaDetailJobResultResponse creates a response to parse from ReportMediaDetailJobResult response
func CreateReportMediaDetailJobResultResponse() (response *ReportMediaDetailJobResultResponse) {
	response = &ReportMediaDetailJobResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
