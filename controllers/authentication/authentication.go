package authentication

import (
	"../../helpers/globals"
	"../../models/customer"
	"github.com/ninnemana/web"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Index(ctx *web.Context) {
	r := ctx.Request
	w := ctx.ResponseWriter

	var errors []string

	params := r.URL.Query()

	redirect := params.Get("redirect")

	if err := params.Get("error"); err != "" {
		errors = append(errors, err)
	}
	if session, err := globals.Store.Get(r, globals.SESSION_KEY); err == nil {
		if flashError := session.Flashes(globals.SESSION_ERROR_KEY); len(flashError) > 0 {
			_ = json.Unmarshal([]byte(flashError[0].(string)), &errors)
		}
		session.Save(r, w)
	}

	tmpl := web.NewTemplate(w)

	if len(errors) > 0 {
		tmpl.Bag["Errors"] = errors
	}
	if redirect != "" {
		tmpl.Bag["Redirect"] = redirect
	}
	tmpl.Bag["PageTitle"] = "Login"
	tmpl.Bag["Heading"] = globals.GetGlobal("LOGIN_HEADING")

	tmpl.ParseFile("templates/authentication/index.html", false)

	tmpl.Display(w)
}

func Login(ctx *web.Context) {
	r := ctx.Request
	w := ctx.ResponseWriter

	email := r.FormValue("email")
	password := r.FormValue("password")
	redirect := r.FormValue("redirect")

	user := customer.CustomerUser{
		Email: email,
	}

	cust, err := user.Authenticate(password)

	if err != nil || cust.Id == 0 {
		// We need to populate the error
		// object so we can prompt the user.
		http.Redirect(w, r, "/login?error="+err.Error(), http.StatusFound)
		return
	}

	// Drop the customer ID into session storage
	id_cook := http.Cookie{
		Name:    globals.SESSION_CUSTOMER_KEY,
		Value:   strconv.Itoa(cust.Id),
		Expires: time.Now().AddDate(2, 0, 0),
	}

	// Create a cookie for public and private key
	public_cook := http.Cookie{
		Name:    globals.SESSION_USER_PUBLIC_KEY,
		Expires: time.Now().AddDate(2, 0, 0),
	}
	private_cook := http.Cookie{
		Name:    globals.SESSION_USER_PRIVATE_KEY,
		Expires: time.Now().AddDate(2, 0, 0),
	}

	// Drop the users public and private key into session
	// storage
	for _, key := range user.Keys {
		if strings.ToLower(key.Type) == "private" {
			private_cook.Value = key.Key
		} else if strings.ToLower(key.Type) == "public" {
			public_cook.Value = key.Key
		}
	}

	// Save cookies
	http.SetCookie(w, &id_cook)
	http.SetCookie(w, &public_cook)
	http.SetCookie(w, &private_cook)

	path := "/"

	if redirect != "" {
		path = redirect
	}

	http.Redirect(w, r, path, http.StatusFound)
	return
}

func Logout(ctx *web.Context) {
	r := ctx.Request
	w := ctx.ResponseWriter

	if id_cook, err := r.Cookie(globals.SESSION_CUSTOMER_KEY); err == nil {
		id_cook.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, id_cook)
	}

	if public_cook, err := r.Cookie(globals.SESSION_USER_PUBLIC_KEY); err == nil {
		public_cook.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, public_cook)
	}

	if private_cook, err := r.Cookie(globals.SESSION_USER_PRIVATE_KEY); err == nil {
		private_cook.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, private_cook)
	}

	http.Redirect(w, r, "/", http.StatusFound)
	return
}

func Handler(w http.ResponseWriter, r *http.Request) {
	cook, err := r.Cookie(globals.SESSION_USER_PRIVATE_KEY)
	if err != nil {
		http.Redirect(w, r, "/login?redirect="+r.URL.Path, http.StatusFound)
		return
	}

	if cook == nil || cook.Value == "" {
		http.Redirect(w, r, "/login?redirect="+r.URL.Path, http.StatusFound)
		return
	}
}
