package gotube

import (
	"log"
	"testing"
)

func TestNewGotube(t *testing.T) {

	// init youtube service
	yt, err := New()
	if err != nil {
		log.Fatal(err)
	}

	// getting IvannSerbia channel info
	_, err = yt.ChannelInfo("IvannSerbia")
	if err != nil {
		t.Error(err)
	}

	// getting all the lists info concurrently
	_, err = yt.GetAllPlaylists()
	if err != nil {
		t.Error(err)
	}

	// getting all the videos of all playlists of mine
	_, err = yt.GetAllVideos()
	if err != nil {
		t.Error(err)
	}
}
