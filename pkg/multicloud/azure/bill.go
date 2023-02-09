// Copyright 2019 Yunion
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

package azure

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/httputils"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

func (self *SAzureClient) GetBilling(ctx context.Context, opts *cloudprovider.SBillingOptions) (*cloudprovider.SBillingInfo, error) {
	if self.GetAccessEnv() == api.CLOUD_ACCESS_ENV_AZURE_CHINA {
		return self.getChinaBilling(ctx, opts)
	}
	return self.getBilling(ctx, opts)
}

func (self *SAzureClient) downloadBilling(ctx context.Context, url string, params url.Values) (jsonutils.JSONObject, error) {
	for i := 0; i < 3; i++ {
		resp, err := self._downloadBilling(ctx, url, params)
		if err == nil {
			return resp, nil
		}
		if errors.Cause(err) == cloudprovider.ErrTimeout {
			time.Sleep(time.Second * 10)
			continue
		}
		return nil, err
	}
	return self._downloadBilling(ctx, url, params)
}

func (self *SAzureClient) _downloadBilling(ctx context.Context, url string, params url.Values) (jsonutils.JSONObject, error) {
	client := self.getHttpClient(time.Minute * 30)
	header := http.Header{}
	header.Set("authorization", fmt.Sprintf("bearer %s", self.billAccessKey))
	uri := fmt.Sprintf("%s?%s", url, params.Encode())
	_, resp, err := httputils.JSONRequest(client, ctx, "GET", uri, header, nil, self.debug)
	if err != nil {
		return nil, err
	}
	errMsg := struct {
		Error struct {
			Code string
		}
	}{}
	resp.Unmarshal(&errMsg)
	if len(errMsg.Error.Code) > 0 {
		return nil, errors.Wrapf(cloudprovider.ErrTimeout, resp.String())
	}
	return resp, nil
}

func (self *SAzureClient) getChinaBilling(ctx context.Context, opts *cloudprovider.SBillingOptions) (*cloudprovider.SBillingInfo, error) {
	uri := fmt.Sprintf("https://ea.azure.cn/rest/%s/usage-report", self.enrollmentNumber)
	params := url.Values{}
	params.Set("month", opts.StartDate.Format("2006-01"))
	params.Set("type", "Detail")
	params.Set("fmt", "json")
	resp, err := self.downloadBilling(ctx, uri, params)
	if err != nil {
		return nil, errors.Wrapf(err, "downloadBilling")
	}
	ret := []cloudprovider.AzureBillInfo{}
	err = resp.Unmarshal(&ret)
	if err != nil {
		return nil, errors.Wrapf(err, "resp.Unmarshal")
	}
	return &cloudprovider.SBillingInfo{Azure: ret}, nil
}

func (self *SAzureClient) getBilling(ctx context.Context, opts *cloudprovider.SBillingOptions) (*cloudprovider.SBillingInfo, error) {
	return nil, cloudprovider.ErrNotImplemented
}
