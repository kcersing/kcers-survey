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
	hz update -idl idl_gen/idl/survey.thrift -model_dir idl_gen/model/  --unset_omitempty


.PHONY: gen-ent
gen-ent:
	go generate ./biz/dal/db/mysql/ent
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier ./biz/dal/db/mysql/ent/schema
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/schemaconfig ./biz/dal/db/mysql/ent/schema