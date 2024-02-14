package main

import (
	"fluentLog/controls"
	"github.com/bhmy-shm/gofks"
	gofkConf "github.com/bhmy-shm/gofks/core/config"
)

func main() {
	conf := gofkConf.New()

	gofks.Ignite("/v1").
		LoadWatch(conf).
		Mount(controls.NewAccountController()).
		Launch()
}
