package pluginaccess

import (
	"errors"
	"os"
	"plugin"
)

func PluginAccess(path string, symbolname string, data interface{}) (payload interface{}) {
	pluginOpen, err := plugin.Open(path)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	symName, err := pluginOpen.Lookup(symbolname)
	if err != nil {
		println("Salah simbol")
		os.Exit(1)
	}

	pluginPointer, ok := symName.(func() interface{})
	if !ok {
		panic(errors.New("gagal binding ke plugin interface"))
	}

	data = pluginPointer()
	payload = data

	return payload
}
