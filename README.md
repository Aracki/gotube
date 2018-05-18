# gotube  [![Go Report Card](https://goreportcard.com/badge/github.com/aracki/gotube)](https://goreportcard.com/report/github.com/aracki/gotube)
Library meant to improve working with youtube api in Go 

## Installation
`go get -u github.com/Aracki/gotube/...`

## Authorization

Document which explains how does OAuth 2.0 works for Mobile & Desktop Apps to access Google APIs: https://developers.google.com/youtube/v3/guides/auth/installed-apps

1. You need a Google Account to access the Google Developers Console, request an API key, and register your application.
2. Create new project on https://console.developers.google.com 
3. Go to Credentials/Create credentials/ OAuth client id
- Applications that use JavaScript to access the YouTube Data API must specify authorized JavaScript origins. The origins identify the domains from which your application can send API requests.
- Applications that use languages and frameworks like PHP, Java, Python, Ruby, and .NET must specify authorized redirect URIs. The redirect URIs are the endpoints to which the OAuth 2.0 server can send responses. (you can put ***http://localhost***)
4. Download secret.json file and save it to /etc/youtube/client_secret.json
5. Enable YouTube Data API v3
6. Start golang app
7. Go to the following link in your browser then type the authorization code  
8. It redirects you to the http://localhost/?state=state-token&code=[code]
9. Paste that *code* from URL into the terminal
10. Saving credential file to: `~/.credentials/youtube-go-quickstart.json`. Gotube initialized!
