package aegis

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

// Process is a nested struct in aegis response
type Process struct {
	Id          int    `json:"Id" xml:"Id"`
	ProcessId   int    `json:"ProcessId" xml:"ProcessId"`
	ProcessName string `json:"ProcessName" xml:"ProcessName"`
	FilePath    string `json:"FilePath" xml:"FilePath"`
	Md5         string `json:"Md5" xml:"Md5"`
	Level       int    `json:"Level" xml:"Level"`
	ProcessType int    `json:"ProcessType" xml:"ProcessType"`
	Status      int    `json:"Status" xml:"Status"`
}
