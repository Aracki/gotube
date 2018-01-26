package gotube

import (
	"fmt"
	"os"
)

const (
	TempFileName = "all_songs"
)

// WriteAllSongsToFile goes through all playlists and write all songs to file
// it returns pointer to file
func (yt Youtube) WriteAllSongsToFile() error {

	file, err := os.Create(TempFileName)
	if err != nil {
		return err
	}

	// get all playlists of mine
	call := yt.Service.Playlists.List(snippetContentDetails)
	call = call.MaxResults(50).Mine(true)
	response, err := call.Do()
	if err != nil {
		return err
	}

	// goes through all playlists
	for _, pl := range response.Items {
		file.WriteString(fmt.Sprintf("___%s\n", pl.Snippet.Title))

		pageToken := ""
		for {
			call := yt.Service.PlaylistItems.List(snippetContentDetails)
			call = call.PlaylistId(pl.Id).MaxResults(50)
			response, err := call.PageToken(pageToken).Do()
			if err != nil {
				return err
			}

			// move pageToken to another page
			pageToken = response.NextPageToken

			for _, item := range response.Items {
				file.WriteString(fmt.Sprintf("%s\n", item.Snippet.Title))
			}
			// if there are no pages
			if pageToken == "" {
				break
			}
		}
	}

	return nil
}
