package certdbb

import (
	"github.com/caddyserver/caddy"
)

func init() {
	caddy.RegisterPlugin("certdbb", caddy.Plugin{
		ServerType: "http",
		Action:     Setup,
	})
}

func Setup(c *caddy.Controller) error {
	return nil
}
