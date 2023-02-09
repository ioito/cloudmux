// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudprovider

import "time"

type TBillType string

const (
	DailyBill TBillType = TBillType("daily")
	MonthBill TBillType = TBillType("month")
)

type SBillingOptions struct {
	BillingType TBillType
	StartDate   time.Time
	EndData     time.Time
}

type AzureBillInfo struct {
	SubscriptionGuid string
	AccountOwnerId   string
	AccountName      string
	SubscriptionName string
	Date             string
	Month            int
	Day              int
	Year             int
	Product          string
	MeterId          string
	MeterCategory    string
	MeterSubCategory string
	MeterRegion      string
	MeterName        string
	ConsumedQuantity float64
	ResourceRate     float64
	ExtendedCost     float64
	ResourceLocation string
	ConsumedService  string
	InstanceId       string
	Tags             map[string]string
	UnitOfMeasure    string
	ResourceGroup    string
}

type SBillingInfo struct {
	// Azure
	Azure []AzureBillInfo
}
