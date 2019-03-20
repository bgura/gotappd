package gotappd

import (
	"time"
	"fmt"
	"strconv"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

// Client for the untappd API
type Gotappd struct {
	clientid     string
	clientsecret string
}

// Struct containing the result of a ping request to the api
type PingResult struct {
	Duration  time.Duration
}

// Create an instance of the Gotappd client
func newClient(clientid string, clientsecret string) *Gotappd {
	api := Gotappd{}

	api.clientid     = clientid
	api.clientsecret = clientsecret
	
	return &api
}

// Build the authenticated resource path for the api
func (api Gotappd) auth(resource string) string {
	auth := "?client_id=" + api.clientid + "&client_secret=" + api.clientsecret
	return resource + auth
}

func (api Gotappd) buildRequest(resource string) (*http.Request, error) {
	fmt.Println(BaseUrl + resource)
	req, err := http.NewRequest("GET", BaseUrl + resource, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", UserAgent)

	return req, nil
}

func (api Gotappd) doRequest(request *http.Request) (*http.Response, error) {
	client := &http.Client {
		Timeout: DefaultTimeout,
	}	
	return client.Do(request)
}

// Function takes an internal params object and appends it to the resource path
func (api Gotappd) get(resource string, params parameter) (*http.Response, error) {
	// Get the authenticated url with parameters
	// Build the request
	// submit the request
	path := api.auth(resource) + "&" + params.urlformat()
	if req, err := api.buildRequest(path); err != nil {
		return nil, err
	} else if resp, err := api.doRequest(req); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

func(api Gotappd) getJson(path string, params parameter) (*json.RawMessage, error) {
	resp, err := api.get(path, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *json.RawMessage
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	return data, err
}

// Send a 'ping' to the server and test its connection.
func (api Gotappd) Ping() (PingResult, error) {
	result := PingResult{}

	start := time.Now()

	resp, err := api.get(PingEndpoint, nil)
	if err == nil {
		defer resp.Body.Close()
	}

	result.Duration = time.Since(start)

	return result, err
}

// Get information about a particular user
func (api Gotappd) GetUserInfo(username *string, params *UserInfoParams) (*json.RawMessage, error) {
	path := UserInfoEndpoint

	// Optionally take a username on the path
	if username != nil {
		path += "/"
		path += *username
	}

	return api.getJson(path, params)
}

// Get information about a users wish list
func (api Gotappd) GetUserWishList(username *string, params *UserWishListParams) (*json.RawMessage, error) {
	path := UserWishListEndpoint

	// Optionally take a username on the path
	if username != nil {
		path += "/"
		path += *username
	}

	return api.getJson(path, params)
}

// Get information about a users friends
func (api Gotappd) GetUserFriends(username *string, params *UserFriendsParams) (*json.RawMessage, error) {
	path := UserFriendsEndpoint

	// Optionally take a user name on the path
	if username != nil {
		path += "/"
		path += *username
	}

	return api.getJson(path, params)
}

// Get information about a users badgers
func (api Gotappd) GetUserBadges(username *string, params *UserBadgesParams) (*json.RawMessage, error) {
	path := UserBadgesEndpoint

	// Optionally take a user name on the path
	if username != nil {
		path += "/"
		path += *username
	}

	return api.getJson(path, params)
}

// Get information about a users distinct beers
func (api Gotappd) GetUserBeers(username *string, params *UserBeersParams) (*json.RawMessage, error) {
	path := UserBeersEndpoint

	// Optionally take a user name on the path
	if username != nil {
		path += "/"
		path += *username
	}

	return api.getJson(path, params)
}

// Get infomration about a brewery
func (api Gotappd) GetBreweryInfo(id uint64, params *BreweryInfoParams) (*json.RawMessage, error) {
	path := BeerInfoEndpoint
	path += "/"
	path += strconv.FormatUint(id, 10)

	return api.getJson(path, params)
}

// Get information about a particular beer
func (api Gotappd) GetBeerInfo(id uint64, params *BeerInfoParams) (*json.RawMessage, error) {
	path := BeerInfoEndpoint
	path += "/"
	path += strconv.FormatUint(id, 10)

	return api.getJson(path, params)
}

// Get information about a particular venue
func (api Gotappd) GetVenueInfo(id uint64, params *VenueInfoParams) (*json.RawMessage, error) {
	path := VenueInfoEndpoint
	path += "/"
	path += strconv.FormatUint(id, 10)

	return api.getJson(path, params)
}

// Search for a beer
func (api Gotappd) GetSearchBeer(query string, params *BeerSearchParams) (*json.RawMessage, error) {
	path := BeerSearchEndpoint

	// Query is a param which is required. We enforce the requirement of it by making
	// it a parameter of this function, vs a field on the params object thats exposed
	// to the public world
	if params == nil {
		params = &BeerSearchParams{}
		params.query = query
	} else {
		params.query = query
	}

	return api.getJson(path, params)
}

// Search for a brewery
func (api Gotappd) GetSearchBrewery(query string, params *BrewerySearchParams) (*json.RawMessage, error) {
	path := BeerSearchEndpoint

	// Query is a param which is required. We enforce the requirement of it by making
	// it a parameter of this function, vs a field on the params object thats exposed
	// to the public world
	if params == nil {
		params = &BrewerySearchParams{}
		params.query = query
	} else {
		params.query = query
	}

	return api.getJson(path, params)
}