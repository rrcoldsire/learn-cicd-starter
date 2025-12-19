package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headers        http.Header
		expectingError bool
		expectedKey    string
	}{
		{
			name:           "no authorization header",
			headers:        http.Header{},
			expectingError: true,
		},
		{
			name: "malformed authorization header",
			headers: http.Header{
				"Authorization": []string{"Bearer key"},
			},
			expectingError: true,
		},
		{
			name: "valid api key",
			headers: http.Header{
				"Authorization": []string{"ApiKey key"},
			},
			expectingError: false,
			expectedKey:    "key",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			key, err := GetAPIKey(test.headers)

			if test.expectingError {
				if err == nil {
					t.Errorf("Test '%s' failed: expecting error, but got nil", test.name)
					return
				}
			} else {
				if err != nil {
					t.Errorf("Test '%s' failed: unexpected error %v", test.name, err)
					return
				}
			}

			if key != test.expectedKey {
				t.Errorf("Test '%s' failed: expected key '%s', got '%s'", test.name, test.expectedKey, key)
				return
			}
		})
	}
}
