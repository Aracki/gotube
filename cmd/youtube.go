package main

import (
	"fmt"

	"github.com/aracki/gotube"
)

func main() {

	// init youtube service
	yt, _ := gotube.New()

	// getting IvannSerbia channel info
	info, _ := yt.ChannelInfo("IvannSerbia")
	fmt.Println(info)

	// getting all the lists info concurrently
	lists, _ := yt.GetAllPlaylists()
	for _, v := range lists {
		fmt.Printf("%+v", v)
	}
	// getting all the videos of all playlists of mine
	yt.GetAllVideos()
}
