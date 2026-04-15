package config

import (
	"fmt"
	"strings"
)

var _ GatewayConfig = (*Gateway)(nil)

type GatewayConfig interface {
	Validator
}

// Gateway is the configuration for the API gateway service.
// It embeds Base and adds upstream routing, timeouts, and auth settings.
type Gateway struct {
	Base

	// Routes maps URL path prefixes to backend service base URLs.
	// Example: { "/inventory/": "http://inventory:8081", "/shipment/": "http://shipment:8082" }
	Routes map[string]string `json:"routes"`
}

// Validate checks that the Gateway config is complete and sensible.
func (g *Gateway) Validate() error {
	if err := g.Base.validate(); err != nil {
		return fmt.Errorf("base: %w", err)
	}
	if len(g.Routes) == 0 {
		return fmt.Errorf("routes: at least one route must be defined")
	}
	for prefix, target := range g.Routes {
		if strings.TrimSpace(target) == "" {
			return fmt.Errorf("routes: target for prefix %q is empty", prefix)
		}
	}
	return nil
}
