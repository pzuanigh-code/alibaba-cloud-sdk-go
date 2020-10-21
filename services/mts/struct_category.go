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

// Category is a nested struct in mts response
type Category struct {
	CategoryName        string  `json:"CategoryName" xml:"CategoryName"`
	Score               string  `json:"Score" xml:"Score"`
	CateId              string  `json:"CateId" xml:"CateId"`
	CategoryId          string  `json:"CategoryId" xml:"CategoryId"`
	ParentId            string  `json:"ParentId" xml:"ParentId"`
	Label               string  `json:"Label" xml:"Label"`
	Level               string  `json:"Level" xml:"Level"`
	CateName            string  `json:"CateName" xml:"CateName"`
	CategoryDescription string  `json:"CategoryDescription" xml:"CategoryDescription"`
	Persons             Persons `json:"Persons" xml:"Persons"`
}
