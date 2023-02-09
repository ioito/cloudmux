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

package shell

import (
	"context"
	"time"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/azure"
)

func init() {
	type BillingListOptions struct {
		BillingType string `choices:"month|dayly" default:"dayly"`
		StartDate   time.Time
		EndDate     time.Time
	}
	shellutils.R(&BillingListOptions{}, "bill-list", "List billing", func(cli *azure.SRegion, args *BillingListOptions) error {
		if args.StartDate.IsZero() {
			args.StartDate = time.Now().Add(time.Hour * -24)
		}
		opts := &cloudprovider.SBillingOptions{
			BillingType: cloudprovider.TBillType(args.BillingType),
			StartDate:   args.StartDate,
		}
		bill, err := cli.GetClient().GetBilling(context.Background(), opts)
		if err != nil {
			return err
		}
		printList(bill.Azure, 0, 0, 0, nil)
		return nil
	})
}
