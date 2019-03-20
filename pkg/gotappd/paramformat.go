package gotappd

import (
	"strconv"
)

// Parameter interface for UserInfoParams
func (m *UserInfoParams) urlformat() string {
	fmt := ""

	if m == nil {
		return fmt
	}

	if m.Compact {
		fmt += "compact=true"
	} else {
		fmt += "compact=false"
	}

	return fmt
}

// Parameter interface for UserWishListParams
func(m *UserWishListParams) urlformat() string {
	fmt := ""

	if m == nil {
		return fmt
	}

	if m.Offset != 0 {
		fmt += "offset=" + strconv.FormatUint(m.Offset, 10) 
	}

	if m.Limit != 0 {
		if len(fmt) != 0 {
			fmt += "&"
		}
		fmt += "limit=" + strconv.FormatUint(m.Limit, 10)
	}

	if m.Sort != SORT_NONE {
		if len(fmt) != 0 {
			fmt += "&"
		} 
		fmt += "sort="

		switch v := m.Sort; v {
		default:
			fmt += "date"
		case SORT_CHECKIN:
			fmt += "checkin"
		case SORT_HIGHEST_RATED:
			fmt += "highest_rated"
		case SORT_LOWEST_RATED:
			fmt += "lowest_rated"
		case SORT_HIGHEST_ABV:
			fmt += "highest_abv"
		case SORT_LOWEST_ABV:
			fmt += "lowest_abv"
		}
	}

	return fmt
}

// Parameter interface for UserFriendsParams
func(m *UserFriendsParams) urlformat() string {
	fmt := ""

	if m == nil {
		return fmt
	}

	if m.Offset != 0 {
		fmt += "offset=" + strconv.FormatUint(m.Offset, 10) 
	}

	if m.Limit != 0 {
		if len(fmt) != 0 {
			fmt += "&"
		}
		fmt += "limit=" + strconv.FormatUint(m.Limit, 10)
	}

	return fmt
}

// Parameter interface for UserBadgesParams
func(m *UserBadgesParams) urlformat() string {
	fmt := ""

	if m == nil {
		return fmt
	}

	if m.Offset != 0 {
		fmt += "offset=" + strconv.FormatUint(m.Offset, 10) 
	}

	if m.Limit != 0 {
		if len(fmt) != 0 {
			fmt += "&"
		}
		fmt += "limit=" + strconv.FormatUint(m.Limit, 10)
	}

	return fmt
}

// Parameter interface for UserBeersParams
func(m *UserBeersParams) urlformat() string {
	fmt := ""

	if m == nil {
		return fmt
	}

	if m.Offset != 0 {
		fmt += "offset=" + strconv.FormatUint(m.Offset, 10) 
	}

	if m.Limit != 0 {
		if len(fmt) != 0 {
			fmt += "&"
		}
		fmt += "limit=" + strconv.FormatUint(m.Limit, 10)
	}
	
	if m.Sort != SORT_NONE {
		if len(fmt) != 0 {
			fmt += "&"
		} 
		fmt += "sort="

		switch v := m.Sort; v {
		default:
			fmt += "date"
		case SORT_CHECKIN:
			fmt += "checkin"
		case SORT_HIGHEST_RATED:
			fmt += "highest_rated"
		case SORT_LOWEST_RATED:
			fmt += "lowest_rated"
		case SORT_HIGHEST_RATED_YOU:
			fmt += "highest_rated_you"
		case SORT_LOWEST_RATED_YOU:
			fmt += "lowest_rated_you"
		}
	}

	return fmt
}

// Parameter interface for BreweryInfo
func (m *BreweryInfoParams) urlformat() string {
	fmt := ""

	if m == nil {
		return fmt
	}

	if m.Compact {
		fmt += "compact=true"
	} else {
		fmt += "compact=false"
	}
	return fmt
}

// Parameter interface for BeerInfoParams
func (m *BeerInfoParams) urlformat() string {
	fmt := ""

	if m == nil {
		return fmt
	}

	if m.Compact {
		fmt += "compact=true"
	} else {
		fmt += "compact=false"
	}
	return fmt
}

// Parameter interface for VenueInfoParams
func (m *VenueInfoParams) urlformat() string {
	fmt := ""

	if m == nil {
		return fmt
	}

	if m.Compact {
		fmt += "compact=true"
	} else {
		fmt += "compact=false"
	}
	return fmt
}

// Parameter interface for BeerSearchParams
func(m *BeerSearchParams) urlformat() string {
	fmt := ""

	if m == nil {
		return fmt
	}

	if len(m.query) != 0 {
		fmt += "q=" + m.query
	}

	if m.Offset != 0 {
		if len(fmt) != 0 {
			fmt += "&"
		}
		fmt += "offset=" + strconv.FormatUint(m.Offset, 10) 
	}

	if m.Limit != 0 {
		if len(fmt) != 0 {
			fmt += "&"
		}
		fmt += "limit=" + strconv.FormatUint(m.Limit, 10)
	}
	
	if m.Sort != SORT_NONE {
		if len(fmt) != 0 {
			fmt += "&"
		} 
		fmt += "sort="

		switch v := m.Sort; v {
		default:
			fmt += "checkin"
		case SORT_NAME:
			fmt += "name"
		}
	}

	return fmt
}

// Parameter interface for BrewerySearchParams
func(m *BrewerySearchParams) urlformat() string {
	fmt := ""

	if m == nil {
		return fmt
	}

	if len(m.query) != 0 {
		fmt += "q=" + m.query
	}

	if m.Offset != 0 {
		if len(fmt) != 0 {
			fmt += "&"
		}
		fmt += "offset=" + strconv.FormatUint(m.Offset, 10) 
	}

	if m.Limit != 0 {
		if len(fmt) != 0 {
			fmt += "&"
		}
		fmt += "limit=" + strconv.FormatUint(m.Limit, 10)
	}
	
	return fmt
}