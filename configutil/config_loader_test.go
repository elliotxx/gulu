package configutil

import (
	"testing"
	"time"

	"github.com/spf13/afero"
)

type TestConfiguration struct {
	Network struct {
		RequestTimeout time.Duration `json:"requestTimeout,omitempty" yaml:"requestTimeout,omitempty"`
	} `json:"network,omitempty" yaml:"network,omitempty"`
	Logging struct {
		LogLevel            string `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`
		EnableLoggingToFile bool   `json:"enableLoggingToFile,omitempty" yaml:"enableLoggingToFile,omitempty"`
	} `json:"logging,omitempty" yaml:"logging,omitempty"`
}

func TestIsValidConfigFilename(t *testing.T) {
	type args struct {
		filename string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid config filename: yaml",
			args: args{
				filename: "testdata/config.yaml",
			},
			want: true,
		},
		{
			name: "valid config filename: yml",
			args: args{
				filename: "testdata/config.yml",
			},
			want: true,
		},
		{
			name: "valid config filename: json",
			args: args{
				filename: "testdata/config.json",
			},
			want: true,
		},
		{
			name: "valid config filename: toml",
			args: args{
				filename: "testdata/config.toml",
			},
			want: true,
		},
		{
			name: "invalid config filename: ini",
			args: args{
				filename: "testdata/invalid-config.ini",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidConfigFilename(tt.args.filename); got != tt.want {
				t.Errorf("IsValidConfigFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromFile(t *testing.T) {
	type args struct {
		fs       afero.Fs
		filename string
		data     interface{}
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				fs:       afero.OsFs{},
				filename: "testdata/config.yaml",
				data:     &TestConfiguration{},
			},
			wantErr: false,
		},
		{
			name: "failed",
			args: args{
				fs:       afero.OsFs{},
				filename: "testdata/invalid-config.ini",
				data:     &TestConfiguration{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FromFile(tt.args.fs, tt.args.filename, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("FromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFromConfigString(t *testing.T) {
	type args struct {
		config     string
		configType string
		data       interface{}
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				config: `network:
  requestTimeout: 30s
logging:
  logLevel: "debug"
  enableLoggingToFile: true`,
				configType: "YAML",
				data:       &TestConfiguration{},
			},
			wantErr: false,
		},
		{
			name: "failed",
			args: args{
				config:     `failed format`,
				configType: "YAML",
				data:       &TestConfiguration{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FromConfigString(tt.args.config, tt.args.configType, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("FromConfigString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFileExtension(t *testing.T) {
	type args struct {
		filename string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "successful get file extension for yaml file",
			args: args{
				filename: "testdata/config.yaml",
			},
			want: "yaml",
		},
		{
			name: "successful get file extension for yml file",
			args: args{
				filename: "testdata/config.yml",
			},
			want: "yml",
		},
		{
			name: "successful get file extension for json file",
			args: args{
				filename: "testdata/config.json",
			},
			want: "json",
		},
		{
			name: "successful get file extension for toml file",
			args: args{
				filename: "testdata/config.toml",
			},
			want: "toml",
		},
		{
			name: "successful to get file extension for ini file",
			args: args{
				filename: "testdata/invalid-config.ini",
			},
			want: "ini",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileExtension(tt.args.filename); got != tt.want {
				t.Errorf("GetFileExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}
