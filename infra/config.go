package infra

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ConnectionString string
	HttpPort         int
	HttpHost         string
}

const (
	DefaultPort         = 8080
	DefaultHost         = "http://localhost"
	EnvConnectionString = "CONNECTION_STRING"
	EnvPort             = "HTTP_PORT"
	EnvHost             = "HTTP_HOST"
)

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func logError(format string, args ...interface{}) error {
	log.Printf(format, args...)
	return fmt.Errorf(format, args...)
}

func getConnectionString() (string, error) {
	connectionString := GetEnv(EnvConnectionString, "")
	if len(connectionString) == 0 {
		return "", logError("%s is not set", EnvConnectionString)
	}
	return connectionString, nil
}

func getHttpPort() (int, error) {
	port := GetEnv(EnvPort, strconv.Itoa(DefaultPort))
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return 0, logError("Error converting %s to int: %s - Using %d", EnvPort, err, DefaultPort)
	}
	return portInt, nil
}

func getHttpHost() (string, error) {
	host := GetEnv(EnvHost, DefaultHost)
	hostUrl, err := url.Parse(host)
	if err != nil {
		return "", logError("Error parsing %s (%s) - %v", EnvHost, host, err)
	} else {
		if len(hostUrl.Host) == 0 || len(hostUrl.Scheme) == 0 {
			return "", logError("%s is not a valid URL", host)
		}
		if len(strings.Split(hostUrl.Host, ":")) > 1 {
			return "", logError("%s is not a valid URL - DO NOT USE PORT", host)
		}
		host = fmt.Sprintf("%s://%s", hostUrl.Scheme, hostUrl.Host)
	}
	return host, nil
}

func GetConfig() (*Config, error) {
	connectionString, err := getConnectionString()
	if err != nil {
		return nil, err
	}
	port, err := getHttpPort()
	if err != nil {
		return nil, err
	}
	host, err := getHttpHost()
	if err != nil {
		return nil, err
	}

	return &Config{
		ConnectionString: connectionString,
		HttpPort:         port,
		HttpHost:         host,
	}, nil
}
