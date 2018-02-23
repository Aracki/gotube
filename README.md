# gotube  [![Go Report Card](https://goreportcard.com/badge/github.com/aracki/gotube)](https://goreportcard.com/report/github.com/aracki/gotube)
Library meant to improve working with youtube api in Go 

## Installation
`go get -u github.com/Aracki/gotube/...`

## Authorization

Using OAuth 2.0 to Access Google APIs -> https://developers.google.com/identity/protocols/OAuth2

https://developers.google.com/youtube/v3/guides/auth/installed-apps

1. You need a Google Account to access the Google Developers Console, request an API key, and register your application.
2. Create new project on https://console.developers.google.com 
3. Go to Credentials/Create credentials/ OAuth client id
- Applications that use JavaScript to access the YouTube Data API must specify authorized JavaScript origins. The origins identify the domains from which your application can send API requests.
- Applications that use languages and frameworks like PHP, Java, Python, Ruby, and .NET must specify authorized redirect URIs. The redirect URIs are the endpoints to which the OAuth 2.0 server can send responses. (you can put ***http://localhost***)
- Download secret.json file and save it to /etc/youtube/client_secret.json
- Enable YouTube Data API v3
- Start the app
- Go to address -? http://localhost/?state=state-token&code=$CODE
- paste $CODE into terminal
- Saving credential file to: ~/.credentials/youtube-go-quickstart.json
