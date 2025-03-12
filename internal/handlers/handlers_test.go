package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string     // test name
	url                string     //path that matches the routes
	method             string     // get or post
	params             []postData //the things that are being posted
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"generals-quarters", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors-suite", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"make-res", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"make-res-post", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Smith"},
		{key: "email", value: "john@gmail.com"},
		{key: "phone", value: "555-55-55"},
	}, http.StatusOK},
	{"search-availability-post", "/search-availability", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "start", value: "2020-02-02"},
	}, http.StatusOK},
	{"search-availability-post-js", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "start", value: "2020-02-02"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, test := range theTests {
		if test.method == "GET" {
			response, err := ts.Client().Get(ts.URL + test.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if response.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", test.name, test.expectedStatusCode, response.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range test.params {
				values.Add(x.key, x.value)
			}
			response, err := ts.Client().PostForm(ts.URL+test.url, values)

			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if response.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", test.name, test.expectedStatusCode, response.StatusCode)
			}
		}
	}
}
