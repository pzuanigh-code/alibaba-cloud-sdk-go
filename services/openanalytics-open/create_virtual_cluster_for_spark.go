package openanalytics_open

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

// CreateVirtualClusterForSpark invokes the openanalytics_open.CreateVirtualClusterForSpark API synchronously
func (client *Client) CreateVirtualClusterForSpark(request *CreateVirtualClusterForSparkRequest) (response *CreateVirtualClusterForSparkResponse, err error) {
	response = CreateCreateVirtualClusterForSparkResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVirtualClusterForSparkWithChan invokes the openanalytics_open.CreateVirtualClusterForSpark API asynchronously
func (client *Client) CreateVirtualClusterForSparkWithChan(request *CreateVirtualClusterForSparkRequest) (<-chan *CreateVirtualClusterForSparkResponse, <-chan error) {
	responseChan := make(chan *CreateVirtualClusterForSparkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVirtualClusterForSpark(request)
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

// CreateVirtualClusterForSparkWithCallback invokes the openanalytics_open.CreateVirtualClusterForSpark API asynchronously
func (client *Client) CreateVirtualClusterForSparkWithCallback(request *CreateVirtualClusterForSparkRequest, callback func(response *CreateVirtualClusterForSparkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVirtualClusterForSparkResponse
		var err error
		defer close(result)
		response, err = client.CreateVirtualClusterForSpark(request)
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

// CreateVirtualClusterForSparkRequest is the request struct for api CreateVirtualClusterForSpark
type CreateVirtualClusterForSparkRequest struct {
	*requests.RpcRequest
	DefaultExecutorSpecName string           `position:"Body" name:"DefaultExecutorSpecName"`
	MaxMemory               requests.Float   `position:"Body" name:"MaxMemory"`
	SparkModuleReleaseName  string           `position:"Body" name:"SparkModuleReleaseName"`
	Description             string           `position:"Body" name:"Description"`
	DefaultExecutorNumbers  requests.Integer `position:"Body" name:"DefaultExecutorNumbers"`
	MaxCpu                  requests.Float   `position:"Body" name:"MaxCpu"`
	Name                    string           `position:"Body" name:"Name"`
	DefaultDriverSpecName   string           `position:"Body" name:"DefaultDriverSpecName"`
}

// CreateVirtualClusterForSparkResponse is the response struct for api CreateVirtualClusterForSpark
type CreateVirtualClusterForSparkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateVirtualClusterForSparkRequest creates a request to invoke CreateVirtualClusterForSpark API
func CreateCreateVirtualClusterForSparkRequest() (request *CreateVirtualClusterForSparkRequest) {
	request = &CreateVirtualClusterForSparkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "CreateVirtualClusterForSpark", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVirtualClusterForSparkResponse creates a response to parse from CreateVirtualClusterForSpark response
func CreateCreateVirtualClusterForSparkResponse() (response *CreateVirtualClusterForSparkResponse) {
	response = &CreateVirtualClusterForSparkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}