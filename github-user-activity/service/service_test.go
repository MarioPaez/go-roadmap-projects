package service

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

// IntegrationTest
func TestGetGithubActivity(t *testing.T) {

}

func TestManageResponse_OK(t *testing.T) {
	body := "[]"
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
	err := manageResponse(resp, "mario")
	if err != nil {
		t.Errorf("received the following error: %s", err)
	}

}

func TestManageResponse_StatusCodeNotSupported(t *testing.T) {
	body := "[]"
	resp := &http.Response{
		StatusCode: http.StatusAccepted,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
	err := manageResponse(resp, "mario")
	if err == nil {
		t.Errorf("the next error must be thrown: %s", err)
	}

}

func TestManageResponse_UsernameNotFound(t *testing.T) {
	body := "[]"
	resp := &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
	err := manageResponse(resp, "mario")
	if err == nil {
		t.Errorf("the next error must be thrown: %s", err)
	}

}
