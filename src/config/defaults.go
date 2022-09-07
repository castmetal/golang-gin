package config

const (
	PORT             = "8000"
	ENV              = "production"
	AUTH_HEADER      = "X-Auth"
	SERVER_TYPE      = "http"
	SALT_CHARSET     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	DB_HOST          = "localhost"
	DB_PORT          = "5432"
	DB_USER          = "myusername"
	DB_PASSWORD      = "mypassword"
	DB_TIME_ZONE     = "America/Sao_Paulo"
	DB_DATABASE_NAME = "postgres"
	DB_SCHEMA        = "public"
	REDIS_HOST       = "localhost"
	REDIS_PORT       = "6379"
	REDIS_PASSWORD
	REDIS_DB = "0"
)

type Config struct {
	PORT             string
	ENV              string
	AUTH_HEADER      string
	SERVER_TYPE      string
	SALT_CHARSET     string
	DB_HOST          string
	DB_PORT          string
	DB_USER          string
	DB_PASSWORD      string
	DB_TIME_ZONE     string
	DB_DATABASE_NAME string
	DB_SCHEMA        string
	REDIS_HOST       string
	REDIS_PORT       string
	REDIS_PASSWORD   string
	REDIS_DB         string
}

var cfg = Config{
	PORT,
	ENV,
	AUTH_HEADER,
	SERVER_TYPE,
	SALT_CHARSET,
	DB_HOST,
	DB_PORT,
	DB_USER,
	DB_PASSWORD,
	DB_TIME_ZONE,
	DB_DATABASE_NAME,
	DB_SCHEMA,
	REDIS_HOST,
	REDIS_PORT,
	REDIS_PASSWORD,
	REDIS_DB,
}

var portEnvVar = GetEnvVariable("PORT")
var envVar = GetEnvVariable("ENV")
var authHeaderEnvVar = GetEnvVariable("AUTH_HEADER")
var serverTypeEnvVar = GetEnvVariable("SERVER_TYPE")
var saltCharsetEnvVar = GetEnvVariable("SERVER_TYPE")
var dbHostEnvVar = GetEnvVariable("DB_HOST")
var dbUserEnvVar = GetEnvVariable("DB_USER")
var dbPortEnvVar = GetEnvVariable("DB_PORT")
var dbPasswordEnvVar = GetEnvVariable("DB_PASSWORD")
var dbTimeZoneEnvVar = GetEnvVariable("DB_TIME_ZONE")
var dbDatabaseEnvVar = GetEnvVariable("DB_DATABASE_NAME")
var dbSchemaVar = GetEnvVariable("DB_SCHEMA")
var redisHostVar = GetEnvVariable("REDIS_HOST")
var redisPortVar = GetEnvVariable("REDIS_PORT")
var redisPasswordVar = GetEnvVariable("REDIS_PASSWORD")
var redisDbVar = GetEnvVariable("REDIS_DB")

var SystemParams = Config{
	PORT:             *CoalesceString(&portEnvVar, &cfg.PORT),
	ENV:              *CoalesceString(&envVar, &cfg.ENV),
	AUTH_HEADER:      *CoalesceString(&authHeaderEnvVar, &cfg.AUTH_HEADER),
	SERVER_TYPE:      *CoalesceString(&serverTypeEnvVar, &cfg.SERVER_TYPE),
	SALT_CHARSET:     *CoalesceString(&saltCharsetEnvVar, &cfg.SALT_CHARSET),
	DB_HOST:          *CoalesceString(&dbHostEnvVar, &cfg.DB_HOST),
	DB_USER:          *CoalesceString(&dbUserEnvVar, &cfg.DB_USER),
	DB_PORT:          *CoalesceString(&dbPortEnvVar, &cfg.DB_PORT),
	DB_PASSWORD:      *CoalesceString(&dbPasswordEnvVar, &cfg.DB_PASSWORD),
	DB_TIME_ZONE:     *CoalesceString(&dbTimeZoneEnvVar, &cfg.DB_TIME_ZONE),
	DB_DATABASE_NAME: *CoalesceString(&dbDatabaseEnvVar, &cfg.DB_DATABASE_NAME),
	DB_SCHEMA:        *CoalesceString(&dbSchemaVar, &cfg.DB_SCHEMA),
	REDIS_HOST:       *CoalesceString(&redisHostVar, &cfg.REDIS_HOST),
	REDIS_PORT:       *CoalesceString(&redisPortVar, &cfg.REDIS_PORT),
	REDIS_PASSWORD:   *CoalesceString(&redisPasswordVar, &cfg.REDIS_PASSWORD),
	REDIS_DB:         *CoalesceString(&redisDbVar, &cfg.REDIS_DB),
}
