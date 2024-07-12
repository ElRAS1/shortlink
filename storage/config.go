package storage

type data struct {
	id      uint8
	oldlink string
	newlink string
}
type configDB struct {
	Port     string `env:"DB_PORT" env-default:"5432"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Name     string `env:"DB_NAME" env-default:"postgres"`
	User     string `env:"DB_USER" env-default:"user"`
	Password string `env:"DB_PASSWORD"`
}
