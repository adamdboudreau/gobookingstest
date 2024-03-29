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
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"generals", "/generals", "GET", []postData{}, http.StatusOK},
	{"majors", "/majors", "GET", []postData{}, http.StatusOK},
	{"search_availability", "/search_availability", "GET", []postData{}, http.StatusOK},
	{"book_room", "/book_room", "GET", []postData{}, http.StatusOK},
	{"post_search_availability", "/search_availability", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-01"},
	}, http.StatusOK},
	{"post_book_room", "/book_room", "POST", []postData{
		{key: "firstName", value: "J"},
		{key: "lastName", value: "D"},
		{key: "email", value: "jd"},
		{key: "phone", value: "1234567890"},
	}, http.StatusOK},
	{"post_book_room", "/book_room", "POST", []postData{
		{key: "firstName", value: "John"},
		{key: "lastName", value: "Doe"},
		{key: "email", value: "j@d.com"},
		{key: "phone", value: "1234567890"},
	}, http.StatusOK},
	{"post_availability_json", "/availability_json", "POST", []postData{}, http.StatusOK},
	{"get_reservation_summary", "/reservation_summary", "GET", []postData{}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}

		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
