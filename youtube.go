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

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// readConfigFile will return oauth2 config
// based on client_secret.json which is located under clientSecretPath
func readConfigFile() (*oauth2.Config, error) {

	filePath, _ := filepath.Abs(clientSecretPath)

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/youtube-go-quickstart.json
	config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// New will read from config file
// than make youtube client and service according to that client
// returns pointer to youtube.Service
func New() (Youtube, error) {
	ctx := context.Background()

	log.Println("initializing new gotube...")

	// reads from config file
	config, err := readConfigFile()
	if err != nil {
		fmt.Println("unable to read/parse client secret file: ", err)
		return Youtube{}, err
	} else {
		log.Println("successfully parsed config from client secret json file")
	}

	// making new client
	c, err := client.GetClient(ctx, config)
	if err != nil {
		return Youtube{}, err
	} else {
		log.Println("successfully created youtube http client")
	}

	// making new service based on client
	s, err := youtube.New(c)
	if err != nil {
		fmt.Println("cannot make youtube client: ", err)
		return Youtube{}, err
	} else {
		log.Println("successfully created youtube service based on http client")
	}

	return Youtube{s}, nil
}