package main

import (
	"net"
	"testing"
)

func TestOpenPageRejectsUnsafeURLs(t *testing.T) {
	for _, value := range []string{"file:///etc/passwd", "https://user:pass@example.com", "not-a-url"} {
		if _, err := openPage(t.Context(), value); err == nil {
			t.Fatalf("accepted %q", value)
		}
	}
}

func TestBlockedIP(t *testing.T) {
	for _, value := range []string{"127.0.0.1", "10.0.0.1", "192.168.1.1", "169.254.1.1", "::1"} {
		if !blockedIP(net.ParseIP(value)) {
			t.Fatalf("did not block %s", value)
		}
	}
	if blockedIP(net.ParseIP("1.1.1.1")) {
		t.Fatal("blocked public address")
	}
}

func TestScreenshotRejectsInvalidName(t *testing.T) {
	if _, err := screenshot(t.Context(), "../file"); err == nil {
		t.Fatal("accepted invalid screenshot name")
	}
}
