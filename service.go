package gotube

import (
	"fmt"
	"sync"

	"github.com/aracki/countgo/model"
	"google.golang.org/api/youtube/v3"
)

const (
	snippetContentDetailsStatistics = "snippet,contentDetails,statistics"
	snippetContentDetails           = "snippet,contentDetails"
)

type Youtube struct {
	Service *youtube.Service
}

// The getChannelInfo uses forUsername
// to get info (id, tittle, totalViews and description)
func (yt Youtube) ChannelInfo(forUsername string) (string, error) {

	call := yt.Service.Channels.List(snippetContentDetailsStatistics)
	call = call.ForUsername(forUsername)
	response, err := call.Do()
	if err != nil {
		return "", err
	}

	var info string

	info = fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
		"and it has %d views. \n",
		response.Items[0].Id,
		response.Items[0].Snippet.Title,
		response.Items[0].Statistics.ViewCount)
	info += fmt.Sprintf(response.Items[0].Snippet.Description)

	return info, nil
}

// Gets all playlists of current user - maxResult is set to 50 (default is 5)
// returns array of all playlists (id, name, count)
func (yt Youtube) GetAllPlaylists() ([]model.Playlist, error) {

	// get all playlists
	call := yt.Service.Playlists.List(snippetContentDetails)
	call = call.MaxResults(50).Mine(true)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	var pls []model.Playlist
	for _, pl := range response.Items {
		pls = append(pls, model.Playlist{
			Id:          pl.Id,
			Title:       pl.Snippet.Title,
			VideosCount: int(pl.ContentDetails.ItemCount),
		})
	}

	return pls, nil
}

// Gets all the videos of all playlists of mine
// Different goroutines are appending the same vds slice;
// WaitGroup waits for all goroutines to finish
func (yt Youtube) GetAllVideos() (vds []model.Video, err error) {

	// get all playlists of mine
	call := yt.Service.Playlists.List(snippetContentDetails)
	call = call.MaxResults(50).Mine(true)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(response.Items))

	for _, pl := range response.Items {
		go func(p *youtube.Playlist) {
			v, _ := getAllVideosByPlaylist(yt.Service, p)
			vds = append(vds, v...)
			wg.Done()
		}(pl)
	}
	wg.Wait()

	return vds, nil
}

// Gets all the videos of specific youtube.Playlist
func getAllVideosByPlaylist(service *youtube.Service, pl *youtube.Playlist) ([]model.Video, error) {

	var vds []model.Video
	pageToken := ""

	for {
		call := service.PlaylistItems.List(snippetContentDetails)
		call = call.PlaylistId(pl.Id).MaxResults(50)
		response, err := call.PageToken(pageToken).Do()
		if err != nil {
			return nil, err
		}

		// move pageToken to another page
		pageToken = response.NextPageToken

		for _, item := range response.Items {
			t := ""
			if item.Snippet.Thumbnails != nil && item.Snippet.Thumbnails.Medium != nil {
				t = item.Snippet.Thumbnails.Medium.Url
			}
			vds = append(vds, model.Video{
				Title:       item.Snippet.Title,
				Description: item.Snippet.Description,
				PublishedAt: item.Snippet.PublishedAt,
				ResourceId:  item.Snippet.ResourceId.VideoId,
				Thumbnail:   t,
			})
		}
		// if there are no pages
		if pageToken == "" {
			break
		}
	}
	return vds, nil
}
