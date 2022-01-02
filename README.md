# Caddy vars directive

[Caddy][] provides the [http.handlers.vars][] module, but it isn't available for use in a `Caddyfile`. Sometimes setting a variable might be useful but not worth abandoning the simplicity of the Caddyfile configuration format.


## Setup

Build with [xcaddy][]:

```
xcaddy build --with github.com/edudobay/caddy-varsdirective
```


## Motivation and example usage

Rewriting the URL in a PHP/FastCGI app is not so obvious (see [this thread](https://caddy.community/t/symfony-with-uri-prefix/10148/3) for an explanation). In this example, we remove the `/_api` prefix, if present, and also remove trailing slashes.

```
route {
	uri strip_prefix /_api
	uri strip_suffix /
	set_var resolved_uri {path}
	php_fastcgi localhost:9000 {
		env REQUEST_URI {http.vars.resolved_uri}
	}
}
```


[Caddy]: https://caddyserver.com/
[http.handlers.vars]: https://caddyserver.com/docs/modules/http.handlers.vars
[xcaddy]: https://github.com/caddyserver/xcaddy
