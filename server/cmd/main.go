package main

import (
	"github.com/ArkamFahry/uploadnexus/server/bootstrap"
)

func main() {
	bootstrap.Init()
	bootstrap.Serve()
}
