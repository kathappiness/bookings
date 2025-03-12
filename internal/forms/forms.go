package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if there's no errors, otherwise - false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks if the required field is not empty
func (f *Form) Required(field ...string) {
	for _, fname := range field {
		value := f.Get(fname)
		// check if the form has no value
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(fname, "This field cannot be blank")
		}
	}
}

// Has checks if the form field is in post and not empty
func (f *Form) Has(field string) bool {
	x := f.Get(field)
	return x != ""
}

// MinLenght checks for a form field minimum lenght
func (f *Form) MinLenght(field string, minlen int) bool {
	x := f.Get(field)
	if len(x) < minlen {
		f.Errors.Add(field, fmt.Sprintf("This filed must be at least %d characters long", minlen))
		return false
	}
	return true
}

// IsEmail checks if the email is valid
func (f *Form) IsEmail(field string) bool {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
		return false
	}
	return true
}
