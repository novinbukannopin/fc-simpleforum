package configs

type (
	Config struct {
		Service Service `mapstructure:"SERVICE"`
	}

	Service struct {
		Port string `mapstructure:"PORT"`
	}
)
