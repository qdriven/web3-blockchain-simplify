module chains-gotest-backend

go 1.16

require (
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.4.0
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/aliyun/aliyun-oss-go-sdk v2.1.8+incompatible
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/casbin/casbin/v2 v2.30.2
	github.com/casbin/gorm-adapter/v3 v3.2.12
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ethereum/go-ethereum v1.10.17
	github.com/fsnotify/fsnotify v1.4.9
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.7.1
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-resty/resty/v2 v2.4.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gookit/color v1.3.6
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/mojocn/base64Captcha v1.3.4
	github.com/onsi/gomega v1.12.0 // indirect
	github.com/polynetwork/poly v1.3.1-0.20210115104304-aa006115a87d
	github.com/qiniu/api.v7/v7 v7.8.2
	github.com/robfig/cron/v3 v3.0.1
	github.com/satori/go.uuid v1.2.0
	github.com/shirou/gopsutil v3.21.4+incompatible
	github.com/spf13/viper v1.7.1
	github.com/swaggo/gin-swagger v1.3.0
	github.com/tencentyun/cos-go-sdk-v5 v0.7.25
	github.com/tklauser/go-sysconf v0.3.5 // indirect
	github.com/unrolled/secure v1.0.9
	go.uber.org/zap v1.16.0
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.9
)

replace (
	github.com/ethereum/go-ethereum v1.9.15 => /Users/Patrick/go/src/github.com/palettechain/palette
	github.com/polynetwork/poly v0.0.1 => /Users/Patrick/go/src/github.com/palettechain/poly
)
