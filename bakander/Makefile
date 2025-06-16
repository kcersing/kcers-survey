# start the environment of FreeCar

.PHONY: start
start:
	docker-compose up -d --build 

# stop the environment of FreeCar

.PHONY: stop
stop:
	docker-compose down

# run the user
.PHONY: admin
admin:
	go run ./app/admin

# run the api
.PHONY: api
api:
	go run ./app/api

.PHONY: idl-gen
idl-gen:
#	@cd ./app/admin && go hz update -idl idl_gen/admin/schedule.thrift -model_dir idl_gen/model/  --unset_omitempty
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



.PHONY: gen-ent
gen-ent:
	go generate ./biz/dal/db/mysql/ent
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier ./biz/dal/db/mysql/ent/schema
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/schemaconfig ./biz/dal/db/mysql/ent/schema