package databases

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func getConnect() *gorm.DB {
	var (
		_user, _pass, _host, _port, _dbname string
	)
	// Change user password here
	_user = "root"
	_pass = "root#123PD"
	_host = "127.0.0.1"
	_port = "3306"
	_dbname = "demo_springct_app"

	_dns := _user + ":" + _pass + "@tcp(" + _host + ":" + _port + ")/" + _dbname + "?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", _dns)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("In db connection")

	return db
}

func InitDB() *gorm.DB {
	return getConnect()
}

func GetDbConnect(c *gin.Context) {
	fmt.Println("In GetDbConnect test")
	c.JSON(200, gin.H{"message": "In GetDbConnect test"})
}
