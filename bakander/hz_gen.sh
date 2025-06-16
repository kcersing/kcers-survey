#!/bin/sh

#if [ "$1" = "init" ]; then
#    hz new -mod kcers
#    hz update -api api/init.proto
#fi

#hz 代码生成
# /api  saas-basics/app/api 下
#hz new -idl_gen ./../idl_gen/http/user.thrift





#本项目使用 `hz` 生成代码. `hz` 详细使用说明参考 [hz](https://www.cloudwego.io/docs/hertz/tutorials/toolkit/toolkit/).
#- hz install.
#```bash
#go install github.com/cloudwego/hertz/cmd/hz@latest
#```
#- hz new: 新建一个 Hertz 项目
#```bash
#hz new -I api -idl_gen ./../idl_gen/admin/token.thrift -model_dir ./../app/biz/model/ --unset_omitempty
#```
#- hz update: 当你的IDL文件更新，使用该指令进行项目代码更新
#- api.proto 与 base.proto是不需要更新与生成的，因为它们是由导入它们的proto文件生成的
#```bash

	hz update -idl idl_gen/idl/auth.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/banner.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/captcha.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/contract.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/dictionary.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/entry.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/logs.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/member.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/menu.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/order.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/payment.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/product.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/pub.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/schedule.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/service.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/sms.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/sys.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/token.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/user.thrift -model_dir idl_gen/model/  --unset_omitempty
	hz update -idl idl_gen/idl/venue.thrift -model_dir idl_gen/model/  --unset_omitempty