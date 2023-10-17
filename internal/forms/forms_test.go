package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postData := url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "a")
	postData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}

}

func TestForm_Has(t *testing.T) {
	postValues := url.Values{}
	form := New(postValues)

	has := form.Has("whatever")

	if has {
		t.Error("got field when should have been empty")
	}

	postData := url.Values{}
	postData.Add("first_name", "John")
	form = New(postData)

	has = form.Has("first_name")

	if !has {
		t.Errorf("shows form does not have field when it should")
	}

}

func TestForm_MinLength(t *testing.T) {
	postValues := url.Values{}
	form := New(postValues)

	form.MinLength("x", 3)
	if form.Valid() {
		t.Error("Form shows min length for non-existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("Should have an error, but did not get one")
	}

	postValues = url.Values{}
	postValues.Add("first_name", "John")
	form = New(postValues)

	form.MinLength("first_name", 100)

	if form.Valid() {
		t.Errorf("shows min length is 100 met when data is shorter")
	}

	postValues = url.Values{}
	postValues.Add("another_field", "abc123")
	form = New(postValues)

	form.MinLength("another_field", 1)

	if !form.Valid() {
		t.Error("shows minlength of 1 is not met when it is")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("Should have an error, but got one")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postValues := url.Values{}
	form := New(postValues)
	form.IsEmail("x")

	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postValues = url.Values{}
	postValues.Add("email", "john@gmail.com")
	form = New(postValues)
	form.IsEmail("email")

	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}

	postValues = url.Values{}
	postValues.Add("email", "x")
	form = New(postValues)
	form.IsEmail("email")

	if form.Valid() {
		t.Error("got valid for invalid email address")
	}

}
