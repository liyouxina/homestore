package config

type ServerConfig struct {
	Port int
}

var (
	defaultPath = []string{
		"./homestore.config",
		"/etc/homestore/homestore.config",
	}

	serverConfig *ServerConfig
)

func GetConfig(fullPath *string) (ServerConfig, error) {
	if serverConfig != nil {
		return *serverConfig, nil
	}
	return ServerConfig{}, nil
}
