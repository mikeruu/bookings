package forms

import (
	"fmt"
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
		t.Error("Got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form show valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("a")
	if has {
		t.Error("form show valid when it does not have the field")
	}

	postedData := url.Values{}
	postedData.Add("a", "b")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	has = form.Has("a")
	if !has {
		t.Error("form show invalid when it does have the field")
	}
}

func TestForm_MinLenght(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.MinLength("a", 20)
	if form.Valid() {
		t.Error("form shows min lenght for non-existent field")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("should have an error but didng get one")

	}

	postedData = url.Values{}
	postedData.Add("a", "abcd")
	form = New(postedData)

	fmt.Print(form)
	minlen := form.MinLength("a", 200)
	if minlen {
		t.Error("Form shows  min lenght of 200 when data is shorter")
	}

	postedData = url.Values{}
	postedData.Add("another", "abc123")

	form = New(postedData)
	form.MinLength("another", 1)
	if !form.Valid() {
		t.Error("shows min lenght of 1 is not met when ti is")
	}
	isError = form.Errors.Get("another")
	if isError != "" {
		t.Error("should not have an error but got one")

	}

}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("email")

	if form.Valid() {
		t.Error("Showing valid email when it is not")
	}
	postedData = url.Values{}
	postedData.Add("email", "me@here.com")
	form = New(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email add when should not have")
	}

}
