# lolchest.win
Entry to Riot's [API challenge 2016](http://na.leagueoflegends.com/en/news/community/contests/riot-games-api-challenge-2016), to be entered under "usable/practical".

There's no way during the current champion select phase to distinguish which champions can earn the player a hextech chest. This information is viewable on the player's profile view, but there's no access to this view during champion select, and even if it were there's no view to sort by "who can I play to earn a chest?"

A live version of the app is available at [lolchest.win](http://lolchest.win).

## Getting Started

To start a local instance of lolchest, after setting up [Go](https://golang.org/doc/install) and [bower](http://bower.io):

```bash
$gopath> go get github.com/nmonterroso/lolchest.win
$gopath> cd src/github.com/nmonterroso/lolchest.win
lolchest.win> PORT=8080 RIOT_API_KEY=<api key> go run cmd/lolchest-win-server/main.go
```

This will start the bridge to Riot's API authenticated with the environment var `RIOT_API_KEY`. There's also a server that can be used during development to serve the client app, whose code lives in `www`:

```bash
www> bower install
lolchest.win>PORT=8081 WWW_ROOT=/path/to//www go run cmd/www-dev-server/main.go
```

The app is now viewable on `localhost:8081`. 

*Note* - Though the service ports are driven by environment variables the client is currently hardcoded to talk to the server over `8080` in [`www/app/modules/profile/profile.service.js`](https://github.com/nmonterroso/lolchest.win/blob/master/www/app/modules/profile/profile.service.js)

## Development

The underlying http handling for both `client <-> server` and `server <-> riot api` leverages auto-generated server and client code for APIs annotated using [Swagger](http://swagger.io/). To change these interfaces you must have [`go-swagger`](https://github.com/go-swagger/go-swagger) installed:

```bash
$gopath> go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

### lolchest Spec
The spec for the lolchest server lives in `resources/swagger.yaml`. To regenerate the server code after editing the spec, run:

```bash
lolchest.win> swagger generate server -f resources/swagger.yaml -A lolchest.win
```

After making changes, the entry point for making edits to the interaction from the lolchest side is in `restapi/configure_lolchest_win.go`.

### Riot API Spec
The spec for the bridge from lolchest and the Riot API lives in `resources/riot_swagger.yaml`. To regenerate the code after making edits:

```bash
lolchest.win> swagger generate client -f ./resources/riot_swagger.yaml -t riotapi
```

All the code for interacting with the Riot API is in the `riotapi/bridge.go`.

### Client Development
The client application is an [AngularJS](https://angularjs.org/) app. The code lives in `www` and consists of two main views: `home` and `profile`.
