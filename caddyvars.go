package caddyvarsdirective

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(VarsDirective{})
	httpcaddyfile.RegisterHandlerDirective("set_var", parseSetVar)
}

type VarsDirective struct{}

func (VarsDirective) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "edudobay.varsdirective",
		New: func() caddy.Module { return new(VarsDirective) },
	}
}

func parseSetVar(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var name, value string
	for h.Next() {
		if !h.NextArg() {
			return nil, h.ArgErr()
		}
		name = h.Val()
		if !h.NextArg() {
			return nil, h.ArgErr()
		}
		value = h.Val()
		if h.NextArg() {
			return nil, h.ArgErr()
		}
	}
	return caddyhttp.VarsMiddleware{name: value}, nil
}
