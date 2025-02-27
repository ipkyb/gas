package internal

import "github.com/ipkyb/gas/api"

type Config struct {
	Verbose    bool             `json:"verbose"`
	Http       HttpConfig       `json:"http"`
	Middleware MiddlewareConfig `json:"middleware"`
}

type HttpConfig struct {
	ListenAddress       string `json:"listen_address"`
	ListenEnablePrefork bool   `json:"listen_enable_prefork"`
}

type MiddlewareConfig struct {
	Global []string `json:"global"`
}

var middlewareKVs = []kv[api.Handler]{
	{"recovery", HandlerRecovery()},
}

func ConfigDefault() Config {
	return Config{
		Verbose: true,
		Http: HttpConfig{
			ListenAddress:       ":8080",
			ListenEnablePrefork: false,
		},
		Middleware: MiddlewareConfig{
			Global: []string{"recovery"},
		},
	}
}

func configMiddlewareTranslate(key string) api.Handler {
	for _, e := range middlewareKVs {
		if e.Key == key {
			return e.Value
		}
	}
	return nil
}
