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
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type LbListOptions struct {
		Id     string
		Marker string
	}
	shellutils.R(&LbListOptions{}, "elb-list", "List loadbalancer", func(cli *aws.SRegion, args *LbListOptions) error {
		ret, _, e := cli.GetLoadbalancers(args.Id, args.Marker)
		if e != nil {
			return e
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

	type LbIdOptions struct {
		ID string
	}
	shellutils.R(&LbIdOptions{}, "elb-attr-show", "Show loadbalancer attribute", func(cli *aws.SRegion, args *LbIdOptions) error {
		ret, err := cli.GetElbAttributes(args.ID)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

	shellutils.R(&LbIdOptions{}, "elb-delete", "Delete loadbalancer", func(cli *aws.SRegion, args *LbIdOptions) error {
		return cli.DeleteElb(args.ID)
	})

	shellutils.R(&LbIdOptions{}, "elb-tag-list", "Show loadbalancer tags", func(cli *aws.SRegion, args *LbIdOptions) error {
		ret, err := cli.DescribeElbTags(args.ID)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

	type LbBackendGroupListOptions struct {
		ElbId  string
		Id     string
		Marker string
	}

	shellutils.R(&LbBackendGroupListOptions{}, "elb-backend-group-list", "List loadbalancer backend groups", func(cli *aws.SRegion, args *LbBackendGroupListOptions) error {
		ret, _, err := cli.GetElbBackendgroups(args.ElbId, args.Id, args.Marker)
		if err != nil {
			return err
		}
		printList(ret, 0, 0, 0, []string{})
		return nil
	})

}
