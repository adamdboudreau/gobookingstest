package forms

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestValid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()

	if !isValid {
		t.Error("got invalid when it should be valid")
	}
}

type postData struct {
	key   string
	value string
}

func TestForm_RequiredFirstName(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.Required("firstName")

	isValid := form.Valid()

	if isValid {
		t.Error("got valid when it should be invalid")
	} else {
		fmt.Println("form is invalid and should have error cannot be blank")
		if form.Errors.Get("firstName") != "This field cannot be blank" {
			t.Error("first name error not included in errors")
		}
	}

	postedData := url.Values{}
	postedData.Add("firstName", "Bob")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("firstName")

	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}
func TestForm_MinLengthFirstName(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.MinLength("firstName", 3)

	isValid := form.Valid()

	if isValid {
		t.Error("got valid when it should be invalid")
	} else {
		fmt.Println("form is invalid and should have error 3 chars")
		if form.Errors.Get("firstName") != "This field must be at least 3 characters long" {
			t.Error("first name error not included in errors")
		}
	}

	postedData := url.Values{}
	postedData.Add("firstName", "Bobby")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.MinLength("firstName", 3)

	if !form.Valid() {
		t.Error("shows does not have min length of 3 when it does")
	}

	postedData = url.Values{}
	postedData.Add("firstName", "")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.MinLength("firstName", 3)

	if form.Valid() {
		t.Error("no error when name is blank")
	} else {
		err := form.Errors.Get("firstName")
		if err != "This field must be at least 3 characters long" {
			t.Error("min error message not found")
		}
	}
}
func TestForm_HasFirstName(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	hasFirstName := form.Has("firstName")

	if hasFirstName {
		t.Error("got first name when not provided")
	}

	postedData := url.Values{}
	postedData.Add("firstName", "Bobby")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	hasFirstName = form.Has("firstName")

	if !hasFirstName {
		t.Error("doesn't have first name when it is provided")
	}
}
func TestForm_IsEmailAnEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isEmail := form.IsEmail("email")

	if isEmail {
		t.Error("got email when not provided")
	}

	postedData := url.Values{}
	postedData.Add("email", "a@b.com")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	isEmail = form.IsEmail("email")

	if !isEmail {
		t.Error("not email when it is an email")
	}
}
