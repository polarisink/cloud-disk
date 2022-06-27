package helper

import (
	"bytes"
	"cloud-disk/core/define"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	"github.com/satori/go.uuid"
	"math/rand"
	"mime/multipart"
	"net/smtp"
	"os"
	"path"
	"time"
)

//func GenerateToken(id int, identity, name string) (string, error) {
//	//id
//	//identity
//	//name
//	uc := define.UserClaim{
//		Id:       id,
//		Identity: identity,
//		Name:     name,
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
//	ttoken, err := token.SignedString([]byte(define.JwtKey))
//	if err != nil {
//		return "", err
//	}
//	return ttoken, nil
//}

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

// OssLocalFile oss上传本地文件
func OssLocalFile(r *os.File) (string, error) {
	key := GetFileKey(path.Base(r.Name()))
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

// UploadFromByte 上传文件的byte数组
func UploadFromByte(objectKey string, b []byte) (string, error) {
	bucket, err := GetOssBucket()
	if err != nil {
		return "", errors.New("get oss bucket error")
	}
	err2 := bucket.PutObject(objectKey, bytes.NewReader(b))
	if err2 != nil {
		return "", err2
	}
	//拼接地址
	return define.OssPrefix + objectKey, nil
}

// UploadFile 上传文件,文件命名策略在本地
func UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	fileName := fileHeader.Filename
	key := GetFileKey(fileName)
	b := make([]byte, fileHeader.Size)
	_, err := file.Read(b)
	if err != nil {
		return "", err
	}
	//拼接地址
	return UploadFromByte(key, b)
}

func GetFileKey(name string) string {
	return "cloud-disk" + "/" + time.Now().Format("2006-01-02") + "/" + name
}

// AnalyzeToken
// Token 解析
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}

func GenerateToken(id int, identity, name string, second int) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
