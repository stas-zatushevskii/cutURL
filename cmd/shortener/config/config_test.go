package config

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestParseFlags(t *testing.T) {
	resetFlags()

	os.Args = []string{
		"cmd",
		"-a=192.168.0.1:9000",
		"-b=http://mydomain.com/",
	}

	cfg := &Config{}
	cfg.ParseFlags()

	want := &Config{
		ServerURL: "192.168.0.1:9000",
		BaseURL:   "http://mydomain.com/",
	}

	if !reflect.DeepEqual(cfg, want) {
		t.Errorf("ParseFlags() = %+v, want %+v", cfg, want)
	}
}

func TestParseEnv(t *testing.T) {
	_ = os.Setenv("SERVER_ADDRESS", "10.0.0.1:3000")
	_ = os.Setenv("BASE_URL", "http://envdomain.com/")

	cfg := &Config{}
	cfg.ParseEnv()

	want := &Config{
		ServerURL: "10.0.0.1:3000",
		BaseURL:   "http://envdomain.com/",
	}

	if !reflect.DeepEqual(cfg, want) {
		t.Errorf("ParseEnv() = %+v, want %+v", cfg, want)
	}

	_ = os.Unsetenv("SERVER_ADDRESS")
	_ = os.Unsetenv("BASE_URL")
}

func TestLoadConfig(t *testing.T) {
	resetFlags()

	os.Args = []string{
		"cmd",
		"-a=192.168.1.1:8081",
		"-b=http://flagurl.com/",
	}

	_ = os.Setenv("SERVER_ADDRESS", "192.168.1.2:9090")
	_ = os.Setenv("BASE_URL", "http://envurl.com/")

	cfg := LoadConfig()

	want := &Config{
		ServerURL: "192.168.1.2:9090",
		BaseURL:   "http://envurl.com/",
	}

	if !reflect.DeepEqual(cfg, want) {
		t.Errorf("LoadConfig() = %+v, want %+v", cfg, want)
	}

	_ = os.Unsetenv("SERVER_ADDRESS")
	_ = os.Unsetenv("BASE_URL")
}
