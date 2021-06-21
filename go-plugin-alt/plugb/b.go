package plugb

import (
	plugalt "github.com/felixge/dump/go-plugin-alt"
)

func init() {
	plugalt.RegisterPlugin(plugalt.Plugin{Name: "Plugin B"})
}
