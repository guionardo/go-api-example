package infra

import (
	"os"
	"reflect"
	"testing"
)

func TestGetEnv(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
	}
	os.Setenv("USER", "api-user")
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Unexistent key",
			args: args{
				key:          "Unexistent key",
				defaultValue: "default",
			},
			want: "default",
		}, {
			name: "Existent key",
			args: args{
				key:          "USER",
				defaultValue: "default",
			},
			want: os.Getenv("USER"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnv(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetConfig(t *testing.T) {
	type envs struct {
		conn_string string
		port        string
		host        string
	}
	tests := []struct {
		name    string
		envs    envs
		want    *Config
		wantErr bool
	}{
		{
			name:    "Empty connection string",
			envs:    envs{},
			want:    &Config{},
			wantErr: true,
		},
		{
			name: "Invalid HTTP Port",
			envs: envs{
				conn_string: "feiras_test.db",
				port:        "invalid",
			},
			want:    &Config{},
			wantErr: true,
		},
		{
			name: "Invalid HTTP Host",
			envs: envs{
				conn_string: "feiras_test.db",
				port:        "8081",
				host:        "invalid",
			},
			want:    &Config{},
			wantErr: true,
		},
		{
			name: "HTTP Host with port",
			envs: envs{
				conn_string: "feiras_test.db",
				port:        "8081",
				host:        "http://localhost:8081",
			},
			want:    &Config{},
			wantErr: true,
		},
		{
			name: "GetConfig",
			envs: envs{
				conn_string: "feiras_test.db",
				port:        "8081",
				host:        "http://localhost",
			},
			want: &Config{
				ConnectionString: "feiras_test.db",
				HttpPort:         8081,
				HttpHost:         "http://localhost",
			},
		},
		{
			name: "Connection string",
			envs: envs{
				conn_string: "feiras_test.db",
			},
			want: &Config{
				ConnectionString: "feiras_test.db",
				HttpPort:         DefaultPort,
				HttpHost:         DefaultHost,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(EnvConnectionString, tt.envs.conn_string)
			os.Setenv(EnvPort, tt.envs.port)
			os.Setenv(EnvHost, tt.envs.host)
			got, err := GetConfig()
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
