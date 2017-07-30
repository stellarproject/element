package proxy

const configTemplate = ` # element router configuration
{{ range $frontend := .Frontends }}
# {{ $frontend.Name }}
{{ range $host := $frontend.Hosts }}{{ $host }} {
    proxy {{ $frontend.Backend.Path }}{{ range $upstream := $frontend.Backend.Upstreams }} {{ $upstream }} {{ end }} {
	transparent
    }
} {{ end }} {{ end }}
`
