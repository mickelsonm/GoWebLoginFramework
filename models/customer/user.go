package customer

import (
	"../../helpers/globals"
	"../../helpers/rest"
	"github.com/ninnemana/web"
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"time"
)

type CustomerUser struct {
	Id                    string
	Name, Email           string
	DateAdded             time.Time
	Active, Sudo, Current bool
	Location              *CustomerLocation
	Keys                  []ApiCredentials
}

type ApiCredentials struct {
	Key, Type string
	DateAdded time.Time
}

func (user *CustomerUser) Authenticate(pass string) (customer Customer, err error) {

	path := globals.API_DOMAIN + "/customer/auth"
	values := url.Values{}
	values.Set("email", user.Email)
	values.Set("password", pass)

	buf, err := rest.Post(path, values)
	if err != nil {
		return
	}

	if err = json.Unmarshal(buf, &customer); err != nil {
		err = errors.New(string(buf))
	}

	for _, u := range customer.Users {
		if u.Current {
			user.Id = u.Id
			user.Name = u.Name
			user.Email = u.Email
			user.DateAdded = u.DateAdded
			user.Active = u.Active
			user.Sudo = u.Sudo
			user.Current = u.Current
			user.Location = u.Location
			user.Keys = u.Keys
		}
	}

	return
}

func GetUser(ctx *web.Context) (user CustomerUser, err error) {
	var key string
	private_cook, e := ctx.Request.Cookie(globals.SESSION_USER_PRIVATE_KEY)
	if e != nil {
		return user, e
	} else {
		key = private_cook.Value
	}

	if key == "" {
		return user, errors.New("Missing private key cookie")
	}

	path := globals.API_DOMAIN + "/customer/user"
	values := url.Values{}
	values.Set("key", key)

	buf, err := rest.Post(path, values)
	if err != nil {
		return
	}

	err = json.Unmarshal(buf, &user)
	return

}

func (user CustomerUser) GetPrivateKey() string {
	for _, k := range user.Keys {
		if strings.ToLower(k.Type) == "private" {
			return k.Key
		}
	}
	return ""
}

func (user CustomerUser) GetPublicKey() string {
	for _, k := range user.Keys {
		if strings.ToLower(k.Type) == "public" {
			return k.Key
		}
	}
	return ""
}
