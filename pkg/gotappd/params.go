package gotappd

// Interface for parameter objects
type parameter interface {
	// Get the parameters formatter as url parameters
	urlformat()  string
}

// Parameters for querying user information
type UserInfoParams struct {
	// compact (string, optional) - You can pass "true" here only show the user information, 
	// and remove the "checkins", "media", "recent_brews", etc attributes
	Compact   bool
}

type SortType int

const (
	SORT_NONE          SortType = iota
	SORT_DATE          
	SORT_CHECKIN       
	SORT_HIGHEST_RATED 
	SORT_LOWEST_RATED  
	SORT_HIGHEST_ABV   
	SORT_LOWEST_ABV    

	SORT_HIGHEST_RATED_YOU
	SORT_LOWEST_RATED_YOU

	SORT_NAME
)

// Parameters for querying user wish list information
type UserWishListParams struct {
	// offset (int, optional) - The numeric offset that you what results to start
	Offset  uint64
	// limit (int, optional) - The number of results to return, max of 50, default is 25
	Limit   uint64
	// sort (string, optional) - You can sort the results using these values: date - 
	// sorts by date (default), checkin - sorted by highest checkin, highest_rated - sorts
	// by global rating descending order, lowest_rated - sorts by global rating ascending 
	// order, highest_abv - highest ABV from the wishlist, lowest_abv - lowest ABV from the
	// wishlist
	Sort    SortType
}


// Parameters for querying user friends list information
type UserFriendsParams struct {
	// offset (int, optional) - The numeric offset that you what results to start
	Offset  uint64
	// limit (int, optional) - The number of results to return, max of 25, default is 25
	Limit   uint64
}


// Parameters for querying user badgers list information
type UserBadgesParams struct {
	// offset (int, optional) - The numeric offset that you what results to start
	Offset  uint64
	// limit (int, optional) - The number of badges to return in your result set
	Limit   uint64
}

// Parameters for querying user unique beers infromation
type UserBeersParams struct {
	// offset (int, optional) - The numeric offset that you what results to start
	Offset  uint64
	// limit (int, optional) - The number of results to return, max of 50, default is 25
	Limit   uint64
	// sort (string, optional) - Your can sort the results using these values: 
	// date - sorts by date (default), checkin - sorted by highest checkin,
	// highest_rated - sorts by global rating descending order, lowest_rated -
	// sorts by global rating ascending order, highest_rated_you - the user's
	// highest rated beer, lowest_rated_you - the user's lowest rated beer
	Sort     SortType
}

// Parameters when querying brewery info
type BreweryInfoParams struct {
	// compact (string, optional) - You can pass "true" here only show the user information, 
	// and remove the "checkins", "media", "recent_brews", etc attributes
	Compact  bool
}

// Parameters when querying beer info
type BeerInfoParams struct {
	// compact (string, optional) - You can pass "true" here only show the user information, 
	// and remove the "checkins", "media", "recent_brews", etc attributes
	Compact  bool
}

// Parameters when querying venue info
type VenueInfoParams struct {
	// compact (string, optional) - You can pass "true" here only show the user information, 
	// and remove the "checkins", "media", "recent_brews", etc attributes
	Compact  bool
}

// Parameters when searching for a beer
type BeerSearchParams struct {
	// query (string, required) - The search term that you want to search.
	query   string
	// offset (int, optional) - The numeric offset that you what results to start
	Offset  uint64
	// limit (int, optional) - The number of results to return, max of 50, default is 25
	Limit   uint64
	// sort (string, optional) - Your can sort the results using these values: 
	// checkin - sorts by checkin count (default), name - sorted by alphabetic beer name
	Sort    SortType
}

// Parameters when searching for a brewery
type BrewerySearchParams struct {
	// query (string, required) - The search term that you want to search.
	query   string
	// offset (int, optional) - The numeric offset that you what results to start
	Offset  uint64
	// limit (int, optional) - The number of results to return, max of 50, default is 25
	Limit   uint64
}