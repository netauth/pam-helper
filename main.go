package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/netauth/pam-helper/internal/module"
	"github.com/spf13/viper"

	"github.com/netauth/netauth/pkg/netauth"
)

var (
	appLogger hclog.Logger
)

func main() {
	llevel := os.Getenv("NETAUTH_LOGLEVEL")
	if llevel != "" {
		appLogger = hclog.New(&hclog.LoggerOptions{
			Name:  "pam-helper",
			Level: hclog.LevelFromString(llevel),
		})
	} else {
		appLogger = hclog.NewNullLogger()
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.netauth/")
	viper.AddConfigPath("/etc/netauth/")
	if err := viper.ReadInConfig(); err != nil {
		appLogger.Error("Error loading config", "error", err)
		os.Exit(5)
	}

	nacl, err := netauth.NewWithLog(appLogger.Named("netauth"))
	if err != nil {

		os.Exit(2)
	}

	os.Exit(module.Exec(appLogger, nacl))
}
