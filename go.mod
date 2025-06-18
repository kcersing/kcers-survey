module kcers-survey

go 1.24.1

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	entgo.io/ent v0.14.4
	github.com/casbin/casbin/v2 v2.103.0
	github.com/casbin/ent-adapter v1.0.0
)

require (
	github.com/ArtisanCloud/PowerLibs/v3 v3.3.2
	github.com/ArtisanCloud/PowerWeChat/v3 v3.4.4
	github.com/alibabacloud-go/darabonba-openapi/v2 v2.0.10
	github.com/alibabacloud-go/dysmsapi-20170525/v4 v4.1.2
	github.com/alibabacloud-go/tea v1.3.7
	github.com/alibabacloud-go/tea-utils/v2 v2.0.7
	github.com/apache/thrift v0.13.0
	github.com/bytedance/sonic v1.13.2
	github.com/cloudwego/hertz v0.10.0
	github.com/cockroachdb/errors v1.11.1
	github.com/dgraph-io/ristretto v0.1.1
	github.com/go-sql-driver/mysql v1.9.0
	github.com/hertz-contrib/jwt v1.0.2
	github.com/hertz-contrib/logger/accesslog v0.0.0-20240128134225-6b18af47a115
	github.com/hertz-contrib/monitor-prometheus v0.1.3
	github.com/hertz-contrib/obs-opentelemetry/logging/logrus v0.1.1
	github.com/hertz-contrib/paseto v0.0.0-20230508023022-71af6635a26c
	github.com/hertz-contrib/reverseproxy v1.0.5
	github.com/jackc/pgx/v5 v5.7.2
	github.com/jinzhu/copier v0.4.0
	github.com/kcersing/dy v0.0.0-20250319093200-c849f4153c26
	github.com/lib/pq v1.10.9
	github.com/minio/minio-go/v7 v7.0.69
	github.com/mojocn/base64Captcha v1.3.8
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/olivere/elastic/v7 v7.0.32
	github.com/pkg/errors v0.9.1
	github.com/redis/go-redis/v9 v9.6.1
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cast v1.6.0
	github.com/spf13/viper v1.18.2
	github.com/swaggo/swag v1.16.4
	github.com/wechatpay-apiv3/wechatpay-go v0.2.20
	golang.org/x/crypto v0.35.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	aidanwoods.dev/go-paseto v1.3.0 // indirect
	ariga.io/atlas v0.31.1-0.20250212144724-069be8033e83 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/alibabacloud-go/alibabacloud-gateway-spi v0.0.5 // indirect
	github.com/alibabacloud-go/debug v1.0.1 // indirect
	github.com/alibabacloud-go/endpoint-util v1.1.0 // indirect
	github.com/alibabacloud-go/openapi-util v0.1.1 // indirect
	github.com/alibabacloud-go/tea-xml v1.1.3 // indirect
	github.com/aliyun/credentials-go v1.3.10 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/bmatcuk/doublestar/v4 v4.8.1 // indirect
	github.com/bytedance/gopkg v0.1.1 // indirect
	github.com/bytedance/sonic/loader v0.2.4 // indirect
	github.com/casbin/govaluate v1.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/cloudwego/base64x v0.1.5 // indirect
	github.com/cloudwego/gopkg v0.1.4 // indirect
	github.com/cloudwego/netpoll v0.7.0 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/getsentry/sentry-go v0.18.0 // indirect
	github.com/go-openapi/inflect v0.21.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.1 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/hcl/v2 v2.23.0 // indirect
	github.com/hertz-contrib/websocket v0.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.6 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nyaruka/phonenumbers v1.3.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/prometheus/client_golang v1.17.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/rs/xid v1.5.0 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/tidwall/gjson v1.17.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/zclconf/go-cty v1.16.2 // indirect
	github.com/zclconf/go-cty-yaml v1.1.0 // indirect
	go.opentelemetry.io/otel v1.5.0 // indirect
	go.opentelemetry.io/otel/trace v1.5.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/arch v0.6.0 // indirect
	golang.org/x/exp v0.0.0-20231214170342-aacd6d4b4611 // indirect
	golang.org/x/image v0.23.0 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	golang.org/x/tools v0.30.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
