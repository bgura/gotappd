package gotappd

import (
	"time"
)

// Constants used in the gotappd api
const (
	Version        string        = "0.0.1"
	UserAgent      string        = "Gotappd/v" + Version + " (+https://github.com/bgura/gotappd)"
	ApiVersion     string        = "v4"
	BaseUrl	       string        = "http://api.untappd.com/" + ApiVersion
	DefaultTimeout time.Duration = time.Second * 10

	PingEndpoint          string = "/ping"

	UserInfoEndpoint      string = "/user/info"
	UserWishListEndpoint  string = "/user/wishlist"
	UserFriendsEndpoint   string = "/user/friends"
	UserBadgesEndpoint    string = "/user/badges"
	UserBeersEndpoint     string = "/user/beers"
	BreweryInfoEndpoint   string = "/brewery/info"
	BeerInfoEndpoint      string = "/beer/info"
	VenueInfoEndpoint     string = "/venue/info"
	BeerSearchEndpoint    string = "/search/beer"
	BrewerySearchEndpoint string = "/search/brewery"
)