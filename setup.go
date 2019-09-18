package certdb

import (
	"github.com/caddyserver/caddy"
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
)

func init() {
	caddy.RegisterPlugin("certdb", caddy.Plugin{
		ServerType: "http",
		Action:     Setup,
	})
}

func Setup(c *caddy.Controller) error {
	return nil
}

