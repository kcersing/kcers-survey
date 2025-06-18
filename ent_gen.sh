#!/bin/sh

#ent 代码生成
# /biz/del/db/ent/schema/
#go generate ./ent

go generate ./pkg/db/ent


#--feature sql/versioned-migration
#--feature sql/modifier



 go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier ./pkg/db/ent/schema
  go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier  ./ent/schema




