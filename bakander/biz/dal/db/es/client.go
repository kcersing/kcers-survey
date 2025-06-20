// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package es

import (
	//"saas/conf"
	"sync"

	//"github.com/cloudwego/biz-demo/book-shop/app/item/common/entity"
	//"github.com/cloudwego/biz-demo/book-shop/pkg/conf"
	"github.com/olivere/elastic/v7"
)

// ES client
var (
	esOnce sync.Once
	esCli  *elastic.Client
)

// GetESClient get ES client
func GetESClient() *elastic.Client {
	if esCli != nil {
		return esCli
	}

	esOnce.Do(func() {
		cli, err := elastic.NewSimpleClient(
		//elastic.SetURL(conf.Conf().ESAddress.Host + ":" + conf.Conf().ESAddress.Port),
		)
		if err != nil {
			panic("new es client failed, err=" + err.Error())
		}
		esCli = cli
	})
	return esCli
}
