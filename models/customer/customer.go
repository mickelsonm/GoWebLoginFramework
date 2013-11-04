package customer

import (
	"../../helpers/globals"
	"../../helpers/rest"
	"encoding/json"
	"errors"
	"net/url"
)

type State struct {
	Id                  int
	State, Abbreviation string
	Country             *Country
}

type Country struct {
	Id                    int
	Country, Abbreviation string
}

type Customer struct {
	Id                                   int
	Name, Email, Address, Address2, City string
	State                                *State
	PostalCode                           string
	Phone, Fax                           string
	ContactPerson                        string
	Latitude, Longitude                  float64
	Website                              *url.URL
	Parent                               *Customer
	SearchUrl, Logo                      *url.URL
	DealerType                           DealerType
	DealerTier                           DealerTier
	SalesRepresentative                  string
	SalesRepresentativeCode              string
	MapixCode, MapixDescription          string
	Locations                            *[]CustomerLocation
	Users                                []CustomerUser
}

type CustomerLocation struct {
	Id                                     int
	Name, Email, Address, City, PostalCode string
	State                                  *State
	Phone, Fax                             string
	Latitude, Longitude                    float64
	CustomerId                             int
	ContactPerson                          string
	IsPrimary, ShippingDefault             bool
}

type DealerType struct {
	Id           int
	Type, Label  string
	Online, Show bool
	MapIcon      MapIcon
}

type DealerTier struct {
	Id   int
	Tier string
	Sort int
}

type MapIcon struct {
	Id, TierId             int
	MapIcon, MapIconShadow *url.URL
}

type MapGraphics struct {
	DealerTier DealerTier
	DealerType DealerType
	MapIcon    MapIcon
}

type GeoLocation struct {
	Latitude, Longitude float64
}

type DealerLocation struct {
	Id, LocationId                       int
	Name, Email, Address, Address2, City string
	State                                *State
	PostalCode                           string
	Phone, Fax                           string
	ContactPerson                        string
	Latitude, Longitude, Distance        float64
	Website                              *url.URL
	Parent                               *Customer
	SearchUrl, Logo                      *url.URL
	DealerType                           DealerType
	DealerTier                           DealerTier
	SalesRepresentative                  string
	SalesRepresentativeCode              string
	MapixCode, MapixDescription          string
}

func GetCustomer(private_key string) (cust Customer, err error) {
	path := globals.API_DOMAIN + "/customer"
	values := url.Values{
		"key": {private_key},
	}

	buf, err := rest.Post(path, values)
	if err != nil {
		return
	}
	if err = json.Unmarshal(buf, &cust); err != nil {
		err = errors.New(string(buf))
	}

	return
}
