package middleware

import (
	"../../helpers/globals"
	"../../helpers/recaptcha"
	"../../helpers/sortutil"
	"github.com/ninnemana/web"
	"html/template"
	"log"
	"reflect"
	"strings"
	"time"
	"../../models/customer"
	"net/http"
	"fmt"
)

var (
	BaseTemplateFiles = []string{
		"layout.html",
	}

	MiddlewarePassRoutes = []string{
		"/login",
		"/auth/forgot",
		"/logout",
		"/signup",
	}

	currentPath = ""

	FunctionMap = template.FuncMap{
		"UnderscoreReplacement": func(s string) string {
			return strings.Replace(s, " ", "_", -1)
		},
		"UnderscoreReplacementToLower": func(s string) string {
			return strings.ToLower(strings.Replace(s, " ", "_", -1))
		},
		"LastItem": func(i int, total int) bool {
			return i == (total - 1)
		},
		"DisplayNameExists": func(s string) bool {
			return !(len(s) > 0)
		},
		"CurrentYear": func() int {
			return time.Now().Year()
		},
		"IfEmptyOrNull": func(s string) string {
			s = strings.TrimSpace(s)
			if len(s) <= 2 {
				return " "
			}
			return s
		},
		"IfNotEmptyOrNull": func(s string) bool {
			if len(s) == 0 {
				return false
			}
			return true
		},
		"StringEqual": func(x interface{}, y interface{}) bool {
			if x == nil || y == nil {
				return false
			}
			return strings.EqualFold(strings.TrimSpace(x.(string)), strings.TrimSpace(y.(string)))
		},
		"StringToLower": func(s string) string {
			return strings.ToLower(s)
		},
		"StringToUpper": func(s string) string {
			return strings.ToUpper(s)
		},
		"Increment": func(i int) int {
			return i + 1
		},
		"Decrement": func(i int) int {
			return i - 1
		},
		"IntegerEqual": func(x interface{}, y interface{}) bool {
			if x == nil || y == nil {
				return false
			}

			var xint int = 0
			var yint int = 0

			xtyp := reflect.TypeOf(x)
			switch xtyp.Kind() {
			case reflect.Int:
				xint = int(x.(int))
			case reflect.Int32:
				xint = int(x.(int32))
			case reflect.Int16:
				xint = int(x.(int16))
			case reflect.Int64:
				xint = int(x.(int64))
			}

			ytyp := reflect.TypeOf(y)
			switch ytyp.Kind() {
			case reflect.Int:
				yint = int(y.(int))
			case reflect.Int32:
				yint = int(y.(int32))
			case reflect.Int16:
				yint = int(y.(int16))
			case reflect.Int64:
				yint = int(y.(int64))
			}
			if xint == yint {
				return true
			}
			return false
		},
		"IntegerGreater": func(x interface{}, y interface{}) bool {

			if x == nil || y == nil {
				return false
			}

			var xint int = 0
			var yint int = 0

			xtyp := reflect.TypeOf(x)
			switch xtyp.Kind() {
			case reflect.Int:
				xint = int(x.(int))
			case reflect.Int32:
				xint = int(x.(int32))
			case reflect.Int16:
				xint = int(x.(int16))
			case reflect.Int64:
				xint = int(x.(int64))
			}

			ytyp := reflect.TypeOf(y)
			switch ytyp.Kind() {
			case reflect.Int:
				yint = int(y.(int))
			case reflect.Int32:
				yint = int(y.(int32))
			case reflect.Int16:
				yint = int(y.(int16))
			case reflect.Int64:
				yint = int(y.(int64))
			}

			if xint <= yint {
				return false
			}

			return true
		},
		"HtmlSafe": func(s string) template.HTML {
			return template.HTML(s)
		},
		"ReverseSort": func(i interface{}) interface{} {
			sortutil.Reverse(i)
			return i
		},
		"ImagePath": func(path string) string {
			if path == "" {
				return "/img/noimage.png"
			}
			return path
		},
		"ReviewWidth": func(avg interface{}) float64 {
			var fltAvg float64
			switch avg.(type) {
			case float64:
				fltAvg = avg.(float64)
			case int:
				fltAvg = float64(avg.(int))
			}
			return (fltAvg / 5) * 100
		},
		"ActiveStatus": func(status int) bool {
			if status == 900 || status == 800 {
				return true
			}
			return false
		},
		"GetInstallTime": func(attrs []string) string {
			for _, attr := range attrs {
				if strings.Index(strings.ToLower(attr), "install") != -1 {
					return strings.TrimSpace(strings.Replace(attr, "Install Time: ", "", -1))
				}
			}
			return ""
		},
		"GetVisibility": func(attrs []string) string {
			for _, attr := range attrs {
				if strings.Index(strings.ToLower(attr), "visibility") != -1 {
					strArr := strings.Split(attr, ":")
					if len(strArr) == 2 {
						return strArr[1]
					}
				}
			}
			return ""
		},
		"GetDrilling": func(attrs []string) string {
			for _, attr := range attrs {
				if strings.Index(strings.ToLower(attr), "drilling") != -1 {
					if strings.Index(strings.ToLower(attr), "yes") != -1 {
						return "Yes"
					} else if strings.Index(strings.ToLower(attr), "no") != -1 {
						return "No"
					}
				}
			}
			return ""
		},
		"NotWiringClass": func(class string) bool {
			if class != "" && strings.ToLower(class) != "wiring" {
				return true
			}
			return false
		},
		"GenerateReCaptcha": func(s string) template.HTML {
			return template.HTML(recaptcha.Init(s))
		},
		"NotEmptyArray": func(arr []interface{}) bool {
			return len(arr) > 0
		},
		"CurrentDomain": func() string {
			return globals.FIXED_URL
		},
		"ShortDateString": func(t time.Time) string {
			return t.Format("01/02/2006")
		},
		"HTML5DateString": func(t time.Time) string {
			str := t.Format("2006-01-02")
			if str == "0001-01-01" {
				return ""
			}
			return str
		},
		"ShortenString": func(s string, l int) string {
			arr := []byte(s)
			if len(arr) < l {
				return string(arr)
			}

			return strings.TrimSpace(string(arr[0:l])) + "..."
		},
		"GetGlobalProperty": func(s string) string {
			return globals.GetGlobal(s)
		},
		"FormatDate": func(dt time.Time) string {
			tLayout := "Mon, 01/02/06, 3:04PM MST"
			local, _ := time.LoadLocation("US/Central")
			return dt.In(local).Format(tLayout)
		},
		"SortableDate": func(dt time.Time) string {
			layout := "01022006"
			local, _ := time.LoadLocation("US/Central")
			return dt.In(local).Format(layout)
		},
	}
)

func Base(ctx *web.Context, urlArgs ...string) {

        // Parse out the query string
        params := ctx.Request.URL.Query()
        filter := params.Get("filter")

        path := "/%s?redirect=" + ctx.Request.URL.Path

        //Get our customer object
        //from session storage
        var cust customer.Customer
        if !strings.Contains(ctx.Request.URL.String(), "login") && !strings.Contains(ctx.Request.URL.String(), "logout") {
                if id_cook, err := ctx.Request.Cookie(globals.SESSION_CUSTOMER_KEY); err != nil || id_cook == nil {
                        ctx.Redirect(http.StatusFound, fmt.Sprintf(path, "login"))
                        return
                }

                cook, err := ctx.Request.Cookie(globals.SESSION_USER_PRIVATE_KEY)
                if err == nil && cook != nil {
                        cust, err = customer.GetCustomer(cook.Value)
                }
                if cust.Id == 0 {
                        ctx.Redirect(http.StatusFound, fmt.Sprintf(path, "login"))
                        return
                }
        }

        // Create our template
        tmpl, err := web.CreateTemplate(ctx.ResponseWriter)
        if err != nil {
                log.Println("New Template Error" + err.Error())
                ctx.NotFound(err.Error())
                return
        }

        // Assign all defined template functions
        // to the template's FuncMap
        tmpl.FuncMap = FunctionMap

        // Populate data into the template
        tmpl.Bag["Filtration"] = filter
        tmpl.Bag["Customer"] = cust

        // Parse all the template files
        // that create our layout
        for _, file := range BaseTemplateFiles {
                err = tmpl.ParseFile(file, false)
        }

        // Reassign out template back to
        // the boilerplate
        web.SetTemplate(tmpl)
}
