package connect

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestConnectMyDql(t *testing.T) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取参数详情
	// 想要正确的处理 time.Time ，您需要带上 parseTime 参数， (更多参数) 要支持完整的 UTF-8 编码，
	// 您需要将 charset=utf8 更改为 charset=utf8mb4
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("open db failed, err: %v", err)
	}

	fmt.Printf("connect %s success\n", db.Name())
}
