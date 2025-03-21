package configurer

type AppConfig struct {
	App     App   `mapstructure:"app"`
	Log     Log   `mapstructure:"log"`
	MySQL   MySQL `mapstructure:"mysql"`
	Secrets Secrets
}

type App struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
}

type Log struct {
	Env string `mapstructure:"env"`
}

type MySQL struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
}

type Secrets struct {
	MySQLUser   string `envconfig:"SECRET_MYSQL_USER"`
	MySQLPass   string `envconfig:"SECRET_MYSQL_PASS"`
	MySQLDBName string `envconfig:"SECRET_MYSQLDB"`
}
