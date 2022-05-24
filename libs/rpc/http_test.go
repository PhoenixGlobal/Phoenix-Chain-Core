package rpc

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func confirmStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got == want {
		return
	}
	if gotName := http.StatusText(got); len(gotName) > 0 {
		if wantName := http.StatusText(want); len(wantName) > 0 {
			t.Fatalf("response status code: got %d (%s), want %d (%s)", got, gotName, want, wantName)
		}
	}
	t.Fatalf("response status code: got %d, want %d", got, want)
}

func confirmRequestValidationCode(t *testing.T, method, contentType, body string, expectedStatusCode int) {
	t.Helper()
	request := httptest.NewRequest(method, "http://url.com", strings.NewReader(body))
	if len(contentType) > 0 {
		request.Header.Set("Content-Type", contentType)
	}
	code, err := validateRequest(request)
	if code == 0 {
		if err != nil {
			t.Errorf("validation: got error %v, expected nil", err)
		}
	} else if err == nil {
		t.Errorf("validation: code %d: got nil, expected error", code)
	}
	confirmStatusCode(t, code, expectedStatusCode)
}

func TestHTTPErrorResponseWithDelete(t *testing.T) {
	confirmRequestValidationCode(t, http.MethodDelete, contentType, "", http.StatusMethodNotAllowed)
}

func TestHTTPErrorResponseWithPut(t *testing.T) {
	confirmRequestValidationCode(t, http.MethodPut, contentType, "", http.StatusMethodNotAllowed)
}

func TestHTTPErrorResponseWithMaxContentLength(t *testing.T) {
	body := make([]rune, maxRequestContentLength+1)
	confirmRequestValidationCode(t,
		http.MethodPost, contentType, string(body), http.StatusRequestEntityTooLarge)
}

func TestHTTPErrorResponseWithEmptyContentType(t *testing.T) {
	confirmRequestValidationCode(t, http.MethodPost, "", "", http.StatusUnsupportedMediaType)
}

func TestHTTPErrorResponseWithValidRequest(t *testing.T) {
	confirmRequestValidationCode(t, http.MethodPost, contentType, "", 0)
}

func confirmHTTPRequestYieldsStatusCode(t *testing.T, method, contentType, body string, expectedStatusCode int) {
	t.Helper()
	s := Server{}
	ts := httptest.NewServer(&s)
	defer ts.Close()

	request, err := http.NewRequest(method, ts.URL, strings.NewReader(body))
	if err != nil {
		t.Fatalf("failed to create a valid HTTP request: %v", err)
	}
	if len(contentType) > 0 {
		request.Header.Set("Content-Type", contentType)
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	confirmStatusCode(t, resp.StatusCode, expectedStatusCode)
}

func TestHTTPResponseWithEmptyGet(t *testing.T) {
	confirmHTTPRequestYieldsStatusCode(t, http.MethodGet, "", "", http.StatusOK)
}
