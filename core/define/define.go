package define

import (
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

//匿名导入并加载

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"

// MailPassword 获取环境变量中的密码
var MailPassword = os.Getenv("MailPassword")

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpire 验证码过期时间
var CodeExpire = 300

// Bucket todo xxx
var Bucket = "mvpwfb"
var Endpoint = "oss-cn-hangzhou.aliyuncs.com"
var AccessKeyId = os.Getenv("AccessKeyId")
var AccessKeySecret = os.Getenv("AccessKeySecret")
var OssPrefix = "https://" + Bucket + "." + Endpoint + "/"

// Authorization auth
var Authorization = "Authorization"

// PageSize 分页的默认参数
var PageSize = 20

var Datetime = "2006-01-02 15:04:05"

var TokenExpire = 3600
var RefreshTokenExpire = 7200
