package database

import (
	"gdemost/config"
	"gdemost/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

type dataProvider struct {
	db  *gorm.DB
}

var once sync.Once
var DataInstance *dataProvider

func init() {
	newDB := getDB(config.DbString)

	once.Do(func() {
		DataInstance = &dataProvider{
			db: newDB,
		}
	})
}

func (dbp dataProvider) Close() {
	dbp.db.Close()
	dbp.db = nil
}

func getDB(dbString string) *gorm.DB {
	log.Println(dbString)

	newDb, err := gorm.Open("mysql", dbString)
	if err != nil {
		log.Println(err)
		return nil
	}
	return newDb
}

func (dbp dataProvider) GetBridge(id uint) models.Bridge {
	var bridge models.Bridge
	dbp.db.Table("bridges").Where("id like ?", id).First(&bridge)
	return bridge
}

func (dbp dataProvider) GetBridges() []models.Bridge {
	var bridges []models.Bridge
	dbp.db.Preload("Divorces").Find(&bridges)
	return bridges
}
