package main

import (
	"event/core"
	"event/initialize"
)

func main() {
	// TODO: implement
	initialize.InitGlobal()
	initialize.InitLog()
	core.SyncEvent()
}
