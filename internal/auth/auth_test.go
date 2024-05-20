package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input      func() http.Header
		wantAPIKey string
		wantErr    bool
	}{
		"simple": {
			input: func() http.Header {
				header := make(http.Header)
				key := "ApiKey GFFqRp+bANBFqGX2aBqhfV+PndbvGZWy0Vge9ozoJyUema4ROEPp2yygbhhcOaFpqM9/iq6dQe3Z3F6QyaGfFw=="
				header.Add("Authorization", key)
				return header
			},
			wantAPIKey: "GFFqRp+bANBFqGX2aBqhfV+PndbvGZWy0Vge9ozoJyUema4ROEPp2yygbhhcOaFpqM9/iq6dQe3Z3F6QyaGfFw==",
			wantErr:    false,
		},
		"no key": {
			input: func() http.Header {
				header := make(http.Header)
				key := "ApiKey"
				header.Add("Authorization", key)
				return header
			},
			wantAPIKey: "",
			wantErr:    true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			APIKey, err := GetAPIKey(tc.input())
			if (err != nil) != tc.wantErr {
				t.Fatalf("unexpected error %v", err)
			}
			if APIKey != tc.wantAPIKey {
				t.Fatalf("want: %v got: %v", tc.wantAPIKey, APIKey)
			}

		})
	}

}
