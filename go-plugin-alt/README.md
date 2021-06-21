# Go Plugin Alternative

Quick and dirty alternative to using Go plugins inspired by [Caddy's plugin system](https://caddyserver.com/docs/extending-caddy).

```
# run app without any plugins enabled
go run ./cmd

# enable Plugin A
cat << EOF > cmd/plugin_a.go
package main

import (
	_ "github.com/felixge/dump/go-plugin-alt/pluga"
)
EOF

# enable Plugin B
cat << EOF > cmd/plugin_b.go
package main

import (
	_ "github.com/felixge/dump/go-plugin-alt/plugb"
)
EOF


# run app with the above plugins enabled
go run ./cmd
```
