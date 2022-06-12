package models

import (
	//"cloud-disk/core/internal/config"
	//"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

/*func Init(dataSource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}
	return engine
}*/

var Engine = Init();
func Init() *xorm.Engine {
	//todo replace with personal mysql server
	engine, err := xorm.NewEngine("mysql", "root:vgy87dgy16rhd61fAGHbcgA@tcp(127.0.0.1:3306)/cloud_disk?charset=utf8mb4")
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}
	return engine
}

/*func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}*/
