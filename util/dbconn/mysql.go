package dbconn

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Config struct {
	// SSLMode to enable/disable SSL connection
	SSLMode bool `envconfig:"MYSQL_SSL_MODE" default:"true"`
	// MaxIdleConnection to set max idle connection pooling
	MaxIdleConnection int `envconfig:"MYSQL_MAX_IDLE_CONNECTION" default:"5"`
	// MaxOpenConnection to set max open connection pooling
	MaxOpenConnection int `envconfig:"MYSQL_MAX_OPEN_CONNECTION" default:"10"`
	// MaxLifetimeConnectionn to set max lifetime of pooling | minutes unit
	MaxLifetimeConnection int `envconfig:"MYSQL_MAX_LIFETIME_CONNECTION" default:"10"`
	// Host is host of mysql service
	Host string `envconfig:"MYSQL_HOST" required:"true"`
	// Port is port of mysql service
	Port string `envconfig:"MYSQL_PORT" required:"true" default:"3306"`
	// Username is name of registered user in mysql service
	Username string `envconfig:"MYSQL_USERNAME" required:"true"`
	// DBName is name of registered database in mysql service
	DBName string `envconfig:"MYSQL_DB_NAME" required:"true"`
	// Password is password of used Username in mysql service
	Password string `envconfig:"MYSQL_PASSWORD" default:""`
	// LogMode is toggle to enable/disable log query in your service by default false
	LogMode bool `envconfig:"MYSQL_LOG_MODE" default:"true"`
	// SingularTable to activate singular table if you are using eloquent query
	SingularTable bool `envconfig:"MYSQL_SINGULAR_TABLE" default:"true"`
	// ParseTime to parse to local time
	ParseTime bool `envconfig:"MYSQL_PARSE_TIME" default:"true"`
	// Charset to define charset of database
	Charset string `envconfig:"MYSQL_CHARSET" default:"utf8mb4"`
	// Charset to define charset of database
	Loc string `envconfig:"MYSQL_LOC" default:"Local"`
}

// InitGorm is helper function to init gorm database from envar values
// it will panic it cannot find required keys or failed to open database connection
func InitGorm(dbName string) *gorm.DB {
	var cfg Config
	envconfig.MustProcess(dbName, &cfg)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%+v&loc=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Charset, cfg.ParseTime, cfg.Loc)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("cannot open mysql connection with dsn: %s: err:%v", dsn, err)
	}
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnection)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLifetimeConnection) * time.Minute)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnection)

	logMode := logger.Error
	if cfg.LogMode {
		logMode = logger.Info
	}
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: false},
		Logger:         logger.Default.LogMode(logMode),
	})
	if err != nil {
		log.Fatalf("cannot initialize gorm database: %v", err)
	}

	return db
}
