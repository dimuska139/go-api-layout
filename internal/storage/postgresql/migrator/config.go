package migrator

type Config struct {
	DSN                 string `yaml:"dsn"`
	MigrationsTableName string `yaml:"db_version"`
}
