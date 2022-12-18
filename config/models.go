package config

import "time"

// Config of application
type Config struct {
	Server     Server     `mapstructure:"server,omitempty"`
	Swagger    Swagger    `mapstructure:"swagger,omitempty"`
	Http       Http       `mapstructure:"http,omitempty"`
	Grpc       Grpc       `mapstructure:"grpc,omitempty"`
	Logger     Logger     `mapstructure:"logger,omitempty"`
	Postgresql Postgresql `mapstructure:"postgresql,omitempty"`
	Mysql      Mysql      `mapstructure:"mysql,omitempty"`
	Mssql      Mssql      `mapstructure:"mssql,omitempty"`
	MongoDB    MongoDB    `mapstructure:"mongodb,omitempty"`
	Redis      Redis      `mapstructure:"redis,omitempty"`
	Clickhouse Clickhouse `mapstructure:"clickhouse,omitempty"`
	Firestore  Firestore  `mapstructure:"firestore,omitempty"`
	Jobs       Jobs       `mapstructure:"jobs,omitempty"`
	Nats       Nats       `mapstructure:"nats,omitempty"`
	RabbitMq   RabbitMq   `mapstructure:"rabbitmq,omitempty"`
}

// Swagger config
type Swagger struct {
	SWAGGER_BASIC_AUTH_USERNAME string `mapstructure:"SWAGGER_BASIC_AUTH_USERNAME,omitempty"`
	SWAGGER_BASIC_AUTH_PASSWORD string `mapstructure:"SWAGGER_BASIC_AUTH_PASSWORD,omitempty"`
}

// Server config
type Server struct {
	PROJECT_NAME          string        `mapstructure:"PROJECT_NAME,omitempty"`
	SERVICE_NAME          string        `mapstructure:"SERVICE_NAME,omitempty"`
	APP_ENV               string        `mapstructure:"APP_ENV,omitempty"`
	APP_DEBUG             bool          `mapstructure:"APP_DEBUG,omitempty"`
	TIMEOUT               int           `mapstructure:"TIMEOUT,omitempty"`
	APP_SECRET            string        `mapstructure:"APP_SECRET,omitempty"`
	JWT_TOKEN_EXPIRE_TIME int           `mapstructure:"JWT_TOKEN_EXPIRE_TIME,omitempty"`
	APP_VERSION           string        `mapstructure:"APP_VERSION,omitempty"`
	READ_TIMEOUT          time.Duration `mapstructure:"READ_TIMEOUT,omitempty"`
	WRITE_TIMEOUT         time.Duration `mapstructure:"WRITE_TIMEOUT,omitempty"`
	MAX_CONN_IDLE         time.Duration `mapstructure:"MAX_CONN_IDLE,omitempty"`
	MAX_CONN_AGE          time.Duration `mapstructure:"MAX_CONN_AGE,omitempty"`
}

// Http config
type Http struct {
	PORT                string        `mapstructure:"PORT,omitempty"`
	TIMEOUT             time.Duration `mapstructure:"TIMEOUT,omitempty"`
	READ_TIMEOUT        time.Duration `mapstructure:"READ_TIMEOUT,omitempty"`
	WRITE_TIMEOUT       time.Duration `mapstructure:"WRITE_TIMEOUT,omitempty"`
	COOKIE_LIFE_TIME    int           `mapstructure:"COOKIE_LIFE_TIME,omitempty"`
	SESSION_COOKIE_NAME string        `mapstructure:"SESSION_COOKIE_NAME,omitempty"`
	SSL_CERT_PATH       string        `mapstructure:"SSL_CERT_PATH,omitempty"`
	SSL_CERT_KEY        string        `mapstructure:"SSL_CERT_KEY,omitempty"`
}

// Http config
type Grpc struct {
	PORT                string        `mapstructure:"PORT,omitempty"`
	TIMEOUT             time.Duration `mapstructure:"TIMEOUT,omitempty"`
	READ_TIMEOUT        time.Duration `mapstructure:"READ_TIMEOUT,omitempty"`
	WRITE_TIMEOUT       time.Duration `mapstructure:"WRITE_TIMEOUT,omitempty"`
	COOKIE_LIFE_TIME    int           `mapstructure:"COOKIE_LIFE_TIME,omitempty"`
	SESSION_COOKIE_NAME string        `mapstructure:"SESSION_COOKIE_NAME,omitempty"`
	SSL_CERT_PATH       string        `mapstructure:"SSL_CERT_PATH,omitempty"`
	SSL_CERT_KEY        string        `mapstructure:"SSL_CERT_KEY,omitempty"`
}

// Logger config
type Logger struct {
	DISABLE_CALLER     bool   `mapstructure:"DISABLE_CALLER,omitempty"`
	DISABLE_STACKTRACE bool   `mapstructure:"DISABLE_STACKTRACE,omitempty"`
	ENCODING           string `mapstructure:"ENCODING,omitempty"`
	LEVEL              string `mapstructure:"LEVEL,omitempty"`
}

// Postgresql config
type Postgresql struct {
	HOST       string `mapstructure:"HOST,omitempty"`
	PORT       int    `mapstructure:"PORT,omitempty"`
	USER       string `mapstructure:"USER,omitempty"`
	PASS       string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB string `mapstructure:"DEFAULT_DB,omitempty"`
	MAX_CONN   int    `mapstructure:"MAX_CONN,omitempty"`
	DRIVER     string `mapstructure:"DRIVER,omitempty"`
}

// Mysql config
type Mysql struct {
	HOST       string `mapstructure:"HOST,omitempty"`
	PORT       int    `mapstructure:"PORT,omitempty"`
	USER       string `mapstructure:"USER,omitempty"`
	PASS       string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB string `mapstructure:"DEFAULT_DB,omitempty"`
	MAX_CONN   int    `mapstructure:"MAX_CONN,omitempty"`
}

// Mssql config
type Mssql struct {
	HOST       string `mapstructure:"HOST,omitempty"`
	PORT       int    `mapstructure:"PORT,omitempty"`
	USER       string `mapstructure:"USER,omitempty"`
	PASS       string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB string `mapstructure:"DEFAULT_DB,omitempty"`
	MAX_CONN   int    `mapstructure:"MAX_CONN,omitempty"`
}

// MongoDB config
type MongoDB struct {
	HOST           string `mapstructure:"HOST,omitempty"`
	PORT           int    `mapstructure:"PORT,omitempty"`
	USER           string `mapstructure:"USER,omitempty"`
	PASS           string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB     string `mapstructure:"DEFAULT_DB,omitempty"`
	MONGO_DB_ATLAS string `mapstructure:"MONGO_DB_ATLAS,omitempty"`
}

// Redis config
type Redis struct {
	HOST          string `mapstructure:"HOST,omitempty"`
	PORT          int    `mapstructure:"PORT,omitempty"`
	USER          string `mapstructure:"USER,omitempty"`
	PASS          string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB    int    `mapstructure:"DEFAULT_DB,omitempty"`
	MIN_IDLE_CONN int    `mapstructure:"MIN_IDLE_CONN,omitempty"`
	POOL_SIZE     int    `mapstructure:"POOL_SIZE,omitempty"`
	POOL_TIMEOUT  int    `mapstructure:"POOL_TIMEOUT,omitempty"`
}

// Clickhouse config
type Clickhouse struct {
	HOST       string `mapstructure:"HOST,omitempty"`
	PORT       int    `mapstructure:"PORT,omitempty"`
	USER       string `mapstructure:"USER,omitempty"`
	PASS       string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB string `mapstructure:"DEFAULT_DB,omitempty"`
}

// Firestore config
type Firestore struct {
	PROJECT_ID        string `json:"PROJECT_ID,omitempty"`
	DEFULT_COLLECTION string `json:"DEFULT_COLLECTION,omitempty"`
	CREDENTIALS_PATH  string `json:"CREDENTIALS_PATH,omitempty"`
}

// Jobs run intervals
type Jobs struct {
	INTERVAL_INDICATORS   time.Duration `json:"INTERVAL_INDICATORS,omitempty"`
	INTERVAL_FEARGREED    time.Duration `json:"INTERVAL_FEARGREED,omitempty"`
	INTERVAL_LONGSHORT    time.Duration `json:"INTERVAL_LONGSHORT,omitempty"`
	INTERVAL_SEASONINDEX  time.Duration `json:"INTERVAL_SEASONINDEX,omitempty"`
	INTERVAL_WORLDINDICES time.Duration `json:"INTERVAL_WORLDINDICES,omitempty"`
}

// Nats run intervals
type Nats struct {
	SERVER_HOST string `json:"SERVER_HOST,omitempty"`
	SERVER_PORT string `json:"SERVER_PORT,omitempty"`
}

// Nats run intervals
type RabbitMq struct {
	URI string `json:"URI,omitempty"`
}
