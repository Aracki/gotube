# gotube  [![Go Report Card](https://goreportcard.com/badge/github.com/aracki/gotube)](https://goreportcard.com/report/github.com/aracki/gotube)
Library meant to improve working with youtube api in Go 

## Installation
`go get -u github.com/Aracki/gotube/...`

## Authorization

Document which explains how does OAuth 2.0 works for Mobile & Desktop Apps to access Google APIs: https://developers.google.com/youtube/v3/guides/auth/installed-apps

Steps for enabling Youtube API service:

1. You need a Google Account to access the Google Developers Console, request an API key, and register your application.
2. Create new project on https://console.developers.google.com
3. Go to _Credentials_ -> _Create credentials_ -> _OAuth client ID_
    - Select Web Application
    - Authorized redirect URIs: _http://localhost_. (it is the endpoint to which the OAuth 2.0 server can send responses)
4. Download `secret.json` file and save it to `/etc/youtube/client_secret.json`.
5. Make sure YouTube Data API v3 is enabled.
6. Start Go app.
7. Stdout will print link. Open it up in your browser and select proper Youtube account  
8. It redirects you to the http://localhost/?state=state-token&code=_[code]_
9. Paste the *code* from the URL to the shell.
10. It will save credential file to: `~/.credentials/youtube-go-quickstart.json`. 
11. Gotube initialized!
  
