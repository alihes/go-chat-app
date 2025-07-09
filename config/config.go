package config




type Config struct{
	Port	 string
	CertFile string
	KeyFile	 string
}

func Load() *Config{
	return &Config{
		Port:		":443",
		CertFile:	"certs/cert.pem",
		KeyFile:	"certs/key.pem",
	}
}
