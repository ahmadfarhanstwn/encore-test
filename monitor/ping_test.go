package monitor

import (
	"context"
	"testing"
)

func TestPing(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		url string
		up  bool
	}{
		{"encore.dev", true},
		{"google.com", true},
		// Test both with and without "https://"
		{"httpbin.org/status/200", true},
		{"https://httpbin.org/status/200", true},

		// 4xx and 5xx should considered down.
		{"httpbin.org/status/400", false},
		{"https://httpbin.org/status/500", false},
		// Invalid URLs should be considered down.
		{"invalid://scheme", false},
	}

	for _, test := range tests {
		resp, err := Ping(ctx, test.url)
		if err != nil {
			t.Errorf("url %s: unexpected error: %v", test.url, err)
		} else if resp.Up != test.up {
			t.Errorf("url %s: got up=%v, want %v", test.url, resp.Up, test.up)
		}
	}
}
