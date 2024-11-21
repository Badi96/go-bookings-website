package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/some_target", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("returned invalid when should have been valid!")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/some_target", nil)
	form := New(r.PostForm)
	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when required field missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/some_target", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}

}

func TestForm_Has(t *testing.T) {
	// case 1: test to get invalid respose from validator
	form := New(url.Values{})

	has := form.Has("something")
	if has {
		t.Error("form shows that it has field when it does not have")
	}

	// case 2: test wirh valid data
	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)
	has = form.Has("a")
	if !has {
		t.Error("Could not find field when it should exitst")
	}
}

func TestForm_MinLength(t *testing.T) {

	// Test non existent value
	postedData := url.Values{}
	form := New(postedData)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	// Test if no error returned
	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have error but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("random_field", "random_value")
	form = New(postedData)

	// test invalid min length
	form.MinLength("random_value", 100)
	if form.Valid() {
		t.Error("gives min length of 100 wheen it's shorter")
	}

	postedData = url.Values{}
	postedData.Add("another_field", "abracadaba123")
	form = New(postedData)

	// test another invalid min length
	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("gives not min length of 1 when it is ")
	}

	// get error when it shoudl not have
	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("Got error when it should not have")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)
	// test incorrect email and see if we get expected error
	form.IsEmail("inavlidemail")
	if form.Valid() {
		t.Error("got valid email for non-existent field")
	}

	// test correct email and see if we get unexpected error
	postedValues = url.Values{}
	postedValues.Add("email", "myemail@gmail.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got unexpected invalid email for a valid email")
	}

	//test existent but invalid email
	postedValues = url.Values{}
	postedValues.Add("email", "invalid_email")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid for invalid email address")
	}

}
