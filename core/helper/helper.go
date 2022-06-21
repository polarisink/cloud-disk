package helper

import (
	"cloud-disk/core/define"
	"crypto/tls"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	"github.com/satori/go.uuid"
	"math/rand"
	"net/smtp"
	"os"
	"path"
	"time"
)

func GenerateToken(id int, identity, name string) (string, error) {
	//id
	//identity
	//name
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	ttoken, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return ttoken, nil
}

// MailCodeSend 邮箱验证码发送
func MailCodeSend(mail, code string) error {
	e := email.NewEmail()
	e.From = "Get <1952482944@qq.com>"
	e.To = []string{"lqswuhan1999@gmail.com"}
	e.Subject = "Awesome Subject"
	e.HTML = []byte("rx,你的验证码为: <h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "1952482944@qq.com", define.MailPassword, "smtp.qq.com"), &tls.Config{
		InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		return err
	}
	return nil
}

func RandCode() string {
	s := "1234567890"
	code := ""
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

func UUID() string {
	return uuid.NewV4().String()
}

// GetOssBucket todo use like mysql client
func GetOssBucket() (*oss.Bucket, error) {
	client, err := oss.New(define.Endpoint, define.AccessKeyId, define.AccessKeySecret)
	if err != nil {
		return nil, errors.New("oss Client Init Error")
	}
	bucket, err := client.Bucket(define.Bucket)
	if err != nil {
		return nil, errors.New("oss Bucket Init Error")
	}
	return bucket, nil
}

func OssUpload(r *os.File) (string, error) {
	key := "cloud-disk" + "/" + time.Now().Format("2006-01-02") + "/" + path.Base(r.Name())
	bucket, err := GetOssBucket()
	if err != nil {
		return "", errors.New("get oss bucket error")
	}
	err2 := bucket.PutObject(key, r)
	if err2 != nil {
		return "", err2
	}
	//拼接地址
	return define.OssPrefix + key, nil
}
