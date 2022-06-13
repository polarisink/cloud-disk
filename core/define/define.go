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
	jwt.RegisteredClaims
}

var JwtKey = "cloud-disk-key"

// MailPassword 获取环境变量中的密码
var MailPassword = os.Getenv("MailPassword")
