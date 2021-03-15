package adminConf

import (
	"gdemost/config"
	"gdemost/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jteeuwen/go-bindata"
	"github.com/qor/admin"
	"net/http"

)

func InitAdmin() *http.ServeMux {
	println("Hello bridges admin!")

	dbs := config.DbString
	println(dbs)
	DB, err := gorm.Open("mysql", dbs)
	if err != nil {
		println(err)
	}
	DB.AutoMigrate(&models.Bridge{})
	DB.AutoMigrate(&models.Divorce{})

	Admin := admin.New(&admin.AdminConfig{DB: DB})

	Admin.SetAssetFS(bindatafs.AssetFS.NameSpace("admin"))
	
	Admin.AddResource(&models.Bridge{}, &admin.Config{
		Menu: []string{"Мосты"},
		Name: "Мосты",
	})

	m := http.NewServeMux()
	Admin.MountTo("/admin", m)
	return m
}
