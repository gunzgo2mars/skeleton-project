package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gunzgo2mars/skeleton-project/src/pkg/configurer"
)

type mysqlDatabase struct {
	dburi string
}

func NewMysql(conf *configurer.AppConfig) *mysqlDatabase {
	mysqlConf := mysqlDatabase{}

	mysqlConf.dburi = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Secrets.MySQLUser,
		conf.Secrets.MySQLPass,
		conf.MySQL.Address,
		conf.MySQL.Port,
		conf.Secrets.MySQLDBName,
	)

	return &mysqlConf
}

func (p *mysqlDatabase) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(p.dburi), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Close(db *gorm.DB) error {
	dbIns, err := db.DB()
	if err != nil {
		return err
	}

	if closeErr := dbIns.Close(); closeErr != nil {
		return fmt.Errorf("DB close error: %s \n", closeErr.Error())
	}

	return nil
}
