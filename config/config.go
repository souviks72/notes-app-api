package config

type Properties struct {
	Port            string `env:"NOTES_PORT" env-default:"8080"`
	Host            string `env:"NOTES_HOST" env-default:"localhost"`
	DBPort          string `env:"DB_PORT" env-default:"27017"`
	DBHost          string `env:"DB_HOST" env-defaul:"localhost"`
	DBName          string `env:"DB_NAME" env-default:"notes-api-db"`
	NotesCollection string `env:"NOTES_COLLECTION_NAME" env-default:"notes"`
}
