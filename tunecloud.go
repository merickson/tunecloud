package main

import (
	"fmt"
	"github.com/merickson/tunecloud/tunecloud"
)

func main() {
	musicpath := "/Users/matt/Music/iTunes/iTunes Media/Music"

	md := tunecloud.NewMusicDirectory(musicpath)

	res, _ := md.Scan()

	for _, v := range res {
		fmt.Printf("%#v\n", v)
	}
}
