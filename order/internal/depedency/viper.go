package dependency

type Config struct {
  DB_HOST string
}

func LoadConfig() *Config {
  return &Config{
  }
}
