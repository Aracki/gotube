package main

import (
	"fmt"
	"log"

	"github.com/aracki/gotube"
)

func main() {

	// init youtube service
	yt, err := gotube.New()
	if err != nil {
		log.Fatal(err)
	}

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
