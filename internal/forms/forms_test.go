package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	// creating new request
	r := httptest.NewRequest("POST", "/whatever", nil)
	// creating a new form
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	// create a request
	r := httptest.NewRequest("POST", "/whatever", nil)
	// initiating a new form object
	form := New(r.PostForm)

	// defining which fields are required for a form
	form.Required("a", "b", "c")
	// if it's valid return an error
	if form.Valid() {
		t.Error("form shows valid when required field is missing")
	}

	// put values in the form
	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	// call the request again
	r, _ = http.NewRequest("POST", "/whatever", nil)
	// set PostForm field in the request object to postedData
	r.PostForm = postedData
	// create a new Form again
	form = New(r.PostForm)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("shows doesn't have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	if form.Has("username") {
		t.Error("passed validation when the field doesn't exist")
	}

	postedData = url.Values{}
	postedData.Add("aaa", "a")
	form = New(postedData)
	if !form.Has("aaa") {
		t.Error("shows form doesn't have a field when it does")
	}
}

func TestForm_MinLenght(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.MinLenght("x", 3)
	if form.Valid() {
		t.Error("form shows min lenght for non-existing field")
	}
	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("doesn't return an error when it should")
	}

	postedData = url.Values{}
	postedData.Add("username", "a") // the value is too short
	form = New(postedData)
	form.MinLenght("username", 3)
	if form.Valid() {
		t.Error("passed validation when the field is too short")
	}

	postedData = url.Values{}
	postedData.Add("name", "tim") // the value is too short
	form = New(postedData)

	form.MinLenght("name", 3)
	if !form.Valid() {
		t.Error("got an error when the value is of the right size")
	}
	isError = form.Errors.Get("name")
	if isError != "" {
		t.Error("returns an error when there's no errors")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("passed email validation when the field doesn't exist")
	}

	// put values in the form
	postedData = url.Values{}
	postedData.Add("email", "mailgmail.com")
	form = New(postedData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("passed email validation when the value is not a valid email")
	}

	postedData = url.Values{}
	postedData.Add("mail", "mail@gmail.com")
	form = New(postedData)
	form.IsEmail("mail")
	if !form.Valid() {
		t.Error("didn't pass validation when the email is valid")
	}
}
