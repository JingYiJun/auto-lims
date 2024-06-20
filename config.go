package main

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"log/slog"
	"os"
	"strings"
)

var config struct {
	Token      string `env:"TOKEN,required"`
	DefaultUID string `env:"DEFAULT_UID,required"`
}

var keys = make(map[string]string)

func init() {
	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}

	slog.SetLogLoggerLevel(slog.LevelInfo)

	for _, envString := range os.Environ() {
		if strings.HasPrefix(envString, "KEY_") {
			rawString := strings.TrimPrefix(envString, "KEY_")
			keyValueList := strings.Split(rawString, "=")
			if len(keyValueList) < 2 {
				panic("invalid env " + envString)
			}
			keys[keyValueList[0]] = keyValueList[1]
		}
	}

	fmt.Printf("%+v", config)
	fmt.Printf("%+v", keys)
}
