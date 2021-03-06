package baas

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

// DescribeOrdererLogs invokes the baas.DescribeOrdererLogs API synchronously
// api document: https://help.aliyun.com/api/baas/describeordererlogs.html
func (client *Client) DescribeOrdererLogs(request *DescribeOrdererLogsRequest) (response *DescribeOrdererLogsResponse, err error) {
	response = CreateDescribeOrdererLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOrdererLogsWithChan invokes the baas.DescribeOrdererLogs API asynchronously
// api document: https://help.aliyun.com/api/baas/describeordererlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOrdererLogsWithChan(request *DescribeOrdererLogsRequest) (<-chan *DescribeOrdererLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeOrdererLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOrdererLogs(request)
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

// DescribeOrdererLogsWithCallback invokes the baas.DescribeOrdererLogs API asynchronously
// api document: https://help.aliyun.com/api/baas/describeordererlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOrdererLogsWithCallback(request *DescribeOrdererLogsRequest, callback func(response *DescribeOrdererLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOrdererLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeOrdererLogs(request)
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

// DescribeOrdererLogsRequest is the request struct for api DescribeOrdererLogs
type DescribeOrdererLogsRequest struct {
	*requests.RpcRequest
	Lines        string `position:"Query" name:"Lines"`
	ConsortiumId string `position:"Query" name:"ConsortiumId"`
	OrdererName  string `position:"Query" name:"OrdererName"`
}

// DescribeOrdererLogsResponse is the response struct for api DescribeOrdererLogs
type DescribeOrdererLogsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateDescribeOrdererLogsRequest creates a request to invoke DescribeOrdererLogs API
func CreateDescribeOrdererLogsRequest() (request *DescribeOrdererLogsRequest) {
	request = &DescribeOrdererLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-07-31", "DescribeOrdererLogs", "", "")
	return
}

// CreateDescribeOrdererLogsResponse creates a response to parse from DescribeOrdererLogs response
func CreateDescribeOrdererLogsResponse() (response *DescribeOrdererLogsResponse) {
	response = &DescribeOrdererLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
