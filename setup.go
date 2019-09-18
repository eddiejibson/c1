package certdb

import (
	"github.com/caddyserver/caddy"
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

