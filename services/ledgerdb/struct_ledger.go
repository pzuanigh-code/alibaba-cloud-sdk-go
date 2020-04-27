package ledgerdb

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

// Ledger is a nested struct in ledgerdb response
type Ledger struct {
	LedgerName        string         `json:"LedgerName" xml:"LedgerName"`
	LedgerId          string         `json:"LedgerId" xml:"LedgerId"`
	UpdateTime        string         `json:"UpdateTime" xml:"UpdateTime"`
	LedgerStatus      string         `json:"LedgerStatus" xml:"LedgerStatus"`
	CreateTime        string         `json:"CreateTime" xml:"CreateTime"`
	OwnerAliUid       string         `json:"OwnerAliUid" xml:"OwnerAliUid"`
	TimeAnchorCount   int64          `json:"TimeAnchorCount" xml:"TimeAnchorCount"`
	RegionId          string         `json:"RegionId" xml:"RegionId"`
	ZoneId            string         `json:"ZoneId" xml:"ZoneId"`
	JournalCount      int64          `json:"JournalCount" xml:"JournalCount"`
	LedgerType        string         `json:"LedgerType" xml:"LedgerType"`
	LedgerDescription string         `json:"LedgerDescription" xml:"LedgerDescription"`
	MemberCount       int64          `json:"MemberCount" xml:"MemberCount"`
	LastTimeAnchor    LastTimeAnchor `json:"LastTimeAnchor" xml:"LastTimeAnchor"`
}