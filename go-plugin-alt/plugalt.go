package plugalt

type Plugin struct {
	Name string
}

var plugins []Plugin

func RegisterPlugin(p Plugin) {
	plugins = append(plugins, p)
}

func Plugins() []Plugin {
	pluginsCopy := make([]Plugin, len(plugins))
	copy(pluginsCopy, plugins)
	return pluginsCopy
}
