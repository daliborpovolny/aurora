package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type CreateUserParams struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Hash      string   `json:"hash"`
	Scores    []int    `json:"scores"`
	Tags      []string `json:"tags"`
}

func TestDecodeFormSimpleFields(t *testing.T) {
	form := url.Values{}
	form.Set("first_name", "Alice")
	form.Set("last_name", "Smith")
	form.Set("email", "alice@example.com")
	form.Set("hash", "abc123")

	req := &http.Request{
		Method: "POST",
		Header: http.Header{"Content-Type": []string{"application/x-www-form-urlencoded"}},
		Body:   io.NopCloser(strings.NewReader(form.Encode())),
	}
	req.ParseForm()

	var params CreateUserParams
	err := DecodeForm(req, &params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	fmt.Println(params)

	if params.FirstName != "Alice" {
		t.Errorf("expected FirstName to be Alice, got %s", params.FirstName)
	}
	if params.Email != "alice@example.com" {
		t.Errorf("expected Email to be alice@example.com, got %s", params.Email)
	}
}

func TestDecodeFormSliceFields(t *testing.T) {
	form := url.Values{}
	form.Add("tags", "go")
	form.Add("tags", "web")
	form.Add("scores", "10")
	form.Add("scores", "20")

	req := &http.Request{
		Method: "POST",
		Header: http.Header{"Content-Type": []string{"application/x-www-form-urlencoded"}},
		Body:   io.NopCloser(strings.NewReader(form.Encode())),
	}
	req.ParseForm()

	var params CreateUserParams
	err := DecodeForm(req, &params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedTags := []string{"go", "web"}
	for i, tag := range expectedTags {
		if params.Tags[i] != tag {
			t.Errorf("expected Tags[%d] = %s, got %s", i, tag, params.Tags[i])
		}
	}

	expectedScores := []int{10, 20}
	for i, score := range expectedScores {
		if params.Scores[i] != score {
			t.Errorf("expected Scores[%d] = %d, got %d", i, score, params.Scores[i])
		}
	}
}

func TestDecodeFormInvalidInt(t *testing.T) {
	form := url.Values{}
	form.Add("scores", "badvalue")

	req := &http.Request{
		Method: "POST",
		Header: http.Header{"Content-Type": []string{"application/x-www-form-urlencoded"}},
		Body:   io.NopCloser(strings.NewReader(form.Encode())),
	}
	req.ParseForm()

	var params CreateUserParams
	err := DecodeForm(req, &params)
	if err == nil {
		t.Fatal("expected error for invalid int, got nil")
	}
}
