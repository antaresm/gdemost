package adminConf

import (
	"gdemost/config"
	"gdemost/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"net/http"
)

func InitAdmin() *http.ServeMux {
	DB, _ := gorm.Open("mysql", config.DbString)
	DB.AutoMigrate(&models.Bridge{})
	DB.AutoMigrate(&models.Divorce{})

	Admin := admin.New(&admin.AdminConfig{DB: DB})
	Admin.AddResource(&models.Bridge{}, &admin.Config{
		Menu: []string{"Мосты"},
		Name: "Мосты",
	})

	m := http.NewServeMux()
	Admin.MountTo("/admin", m)
	return m
}
