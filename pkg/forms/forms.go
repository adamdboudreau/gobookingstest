package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates custom form struct; embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Required checks fields are not blank
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MinLength checks for min length
// func (f *Form) MinLength(field string, length int, r *http.Request) bool {
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field) //r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form field is empty
// func (f *Form) Has(field string, r *http.Request) bool {
func (f *Form) Has(field string) bool {
	x := f.Get(field) //r.Form.Get(field)
	if x == "" {
		// f.Errors.Add(field, "This field cannot be blank") // might want to check checkbox field with different error
		return false
	}
	return true
}

// IsEmail checks for valid email address
func (f *Form) IsEmail(field string) bool {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
		return false
	}
	return true
}
