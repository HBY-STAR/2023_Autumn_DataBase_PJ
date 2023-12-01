package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() error {
	//price_comparator
	dsn := "root:root@tcp(localhost:3306)/price_comparator?charset=utf8mb4&parseTime=True&loc=Local"
	// 根据你的 MySQL 配置进行修改
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `Login` would be `user` with this option enabled
		}})
	if err != nil {
		return err
	}

	// 迁移数据库，确保 Login 表存在
	err = DB.AutoMigrate(&User{}, &Seller{}, &Admin{}, &Commodity{}, &Platform{}, &UserJwtSecret{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&CommodityItem{}, &Favorite{}, &Message{}, &PriceChange{})

	return err
}

//func setupDatabase() (*gorm.DB, error) {
//	dsn := "root:root@tcp(localhost:3306)/price_comparator"
//	// 根据你的 MySQL 配置进行修改
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true, // use singular table name, table for `Login` would be `user` with this option enabled
//		}})
//	if err != nil {
//		return nil, err
//	}
//
//	// 迁移数据库，确保 Login 表存在
//	err = db.AutoMigrate(&Login{})
//	if err != nil {
//		return nil, err
//	}
//
//	return db, nil
//}
