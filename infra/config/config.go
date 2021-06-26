package config

var config Config

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func setDefaultConfig() {
	config.Db = &DBConfig{
		Host:     "host.docker.internal",
		Port:     "3309",
		User:     "tonu",
		DBName:   "first_go",
		Password: "tonu",
	}
}

type Config struct {
	Db *DBConfig
}

func Db() *DBConfig {
	return config.Db
}

func LoadConfig() {
	setDefaultConfig()
	// _ = viper.BindEnv("ENV_PATH")
	// envPath := viper.GetString("ENV_PATH")
	// if envPath != "" {
	// 	viper.AddConfigPath(envPath)
	// 	viper.SetConfigName("config.local")
	// 	viper.SetConfigType("json")
	// 	err := viper.ReadInConfig()

	// 	if err != nil {
	// 		log.Println(fmt.Sprintf("%s named \"%s\"", err.Error(), envPath))
	// 	}

	// 	config = Config{}

	// 	if err := viper.Unmarshal(&config); err != nil {
	// 		panic(err)
	// 	}

	// 	if r, err := json.MarshalIndent(&config, "", "  "); err == nil {
	// 		fmt.Println(string(r))
	// 	}
	// } else {
	// 	log.Println("ENV_PATH is missing! Serving with default config...")
	// }
}
