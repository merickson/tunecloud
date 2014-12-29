package main

import (
	"flag"
	"fmt"
	"github.com/merickson/tunecloud/tunecloud"
)

var musicpath = flag.String("path", "", "Filesystem path to scan for music")

func main() {
	flag.Parse()

	md := tunecloud.NewMusicDirectory(*musicpath)

	res, _ := md.Scan()

	for _, v := range res {
		fmt.Printf("%#v\n", v)
	}
}
