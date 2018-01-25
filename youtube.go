package gotube

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/aracki/gotube/client"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

const clientSecretPath = "/etc/youtube/client_secret.json"

// readConfigFile will return oauth2 config
// based on client_secret.json which is located in project root
func readConfigFile(path string) (*oauth2.Config, error) {

	filePath, _ := filepath.Abs(path)

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return nil, err
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/youtube-go-quickstart.json
	config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}

	return config, nil
}

// New will read from config file
// than make youtube client and service according to that client
// returns pointer to youtube.Service
func New() (*youtube.Service, error) {
	ctx := context.Background()

	// reads from config file
	config, err := readConfigFile(clientSecretPath)
	if err != nil {
		fmt.Println("Unable to read/parse client secret file", err)
	}

	// making new client
	c := client.GetClient(ctx, config)

	// making new service based on client
	s, err := youtube.New(c)
	if err != nil {
		fmt.Println("Cannot make youtube client", err)
		return nil, err
	}

	return s, err
}
