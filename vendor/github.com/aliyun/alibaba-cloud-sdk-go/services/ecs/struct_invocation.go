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

// Invocation is a nested struct in ecs response
type Invocation struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
>>>>>>> 03397665 (update api)
	Name              string                                   `json:"Name" xml:"Name"`
	PageSize          int64                                    `json:"PageSize" xml:"PageSize"`
	Timed             bool                                     `json:"Timed" xml:"Timed"`
	Frequency         string                                   `json:"Frequency" xml:"Frequency"`
	Content           string                                   `json:"Content" xml:"Content"`
	CommandContent    string                                   `json:"CommandContent" xml:"CommandContent"`
	InvocationStatus  string                                   `json:"InvocationStatus" xml:"InvocationStatus"`
	FileGroup         string                                   `json:"FileGroup" xml:"FileGroup"`
	Description       string                                   `json:"Description" xml:"Description"`
	Overwrite         string                                   `json:"Overwrite" xml:"Overwrite"`
	PageNumber        int64                                    `json:"PageNumber" xml:"PageNumber"`
	CommandId         string                                   `json:"CommandId" xml:"CommandId"`
	TargetDir         string                                   `json:"TargetDir" xml:"TargetDir"`
	FileMode          string                                   `json:"FileMode" xml:"FileMode"`
	TotalCount        int64                                    `json:"TotalCount" xml:"TotalCount"`
	Username          string                                   `json:"Username" xml:"Username"`
	ContentType       string                                   `json:"ContentType" xml:"ContentType"`
	CreationTime      string                                   `json:"CreationTime" xml:"CreationTime"`
	Parameters        string                                   `json:"Parameters" xml:"Parameters"`
	CommandName       string                                   `json:"CommandName" xml:"CommandName"`
	VmCount           int                                      `json:"VmCount" xml:"VmCount"`
	InvokeId          string                                   `json:"InvokeId" xml:"InvokeId"`
	RepeatMode        string                                   `json:"RepeatMode" xml:"RepeatMode"`
	InvokeStatus      string                                   `json:"InvokeStatus" xml:"InvokeStatus"`
	FileOwner         string                                   `json:"FileOwner" xml:"FileOwner"`
	CommandType       string                                   `json:"CommandType" xml:"CommandType"`
	InvocationResults InvocationResults                        `json:"InvocationResults" xml:"InvocationResults"`
	InvokeInstances   InvokeInstancesInDescribeSendFileResults `json:"InvokeInstances" xml:"InvokeInstances"`
<<<<<<< HEAD
<<<<<<< HEAD
=======
	CommandId         string            `json:"CommandId" xml:"CommandId"`
	PageNumber        int64             `json:"PageNumber" xml:"PageNumber"`
	TotalCount        int64             `json:"TotalCount" xml:"TotalCount"`
	PageSize          int64             `json:"PageSize" xml:"PageSize"`
	Timed             bool              `json:"Timed" xml:"Timed"`
	Frequency         string            `json:"Frequency" xml:"Frequency"`
	CommandName       string            `json:"CommandName" xml:"CommandName"`
	Parameters        string            `json:"Parameters" xml:"Parameters"`
	InvokeId          string            `json:"InvokeId" xml:"InvokeId"`
	InvokeStatus      string            `json:"InvokeStatus" xml:"InvokeStatus"`
	CommandContent    string            `json:"CommandContent" xml:"CommandContent"`
	CommandType       string            `json:"CommandType" xml:"CommandType"`
	InvocationResults InvocationResults `json:"InvocationResults" xml:"InvocationResults"`
	InvokeInstances   InvokeInstances   `json:"InvokeInstances" xml:"InvokeInstances"`
>>>>>>> 79bfea2d (update vendor)
=======
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
>>>>>>> 03397665 (update api)
}
