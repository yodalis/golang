package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

// Configurando notation pra conseguir ler as letras maiusculas da conf e entender qual é a chave correspondente da struct
type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JwtExpiresIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config") //nome da config
	viper.SetConfigType("env")        //tipo da config, pode ser yaml, env
	viper.AddConfigPath(path)         // path da config
	viper.SetConfigFile(".env")       // file
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	// vai pegar os valores e colocar na struct
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	// Configurando JWT
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil) //Instância pra poder gerar token JWT

	return cfg, err
}
