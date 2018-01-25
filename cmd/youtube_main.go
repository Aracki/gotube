package main

import (
	"fmt"

	"github.com/aracki/countgo/youtube"
	"github.com/aracki/gotube/service"
)

func main() {

	// init youtube service
	s, _ := youtube.New()

	// getting IvannSerbia channel info
	info, _ := service.ChannelInfo(s, "IvannSerbia")
	fmt.Println(info)

	// getting all the lists info concurrently
	lists, _ := service.GetAllPlaylists(s)
	for _, v := range lists {
		fmt.Printf("%+v", v)
	}
	// getting all the videos of all playlists of mine
	service.GetAllVideos(s)
}
