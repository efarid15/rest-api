package pluginaccess

import (
	"errors"
	"log"
	"os"
	"plugin"
	"runtime"
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

	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	log.Printf("memory usage: %d", m1.Alloc)

	return payload
}
