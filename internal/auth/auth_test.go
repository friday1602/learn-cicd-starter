package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := make(http.Header)
	key := "GFFqRp+bANBFqGX2aBqhfV+PndbvGZWy0Vge9ozoJyUema4ROEPp2yygbhhcOaFpqM9/iq6dQe3Z3F6QyaGfFw=="
	header.Add("APIKey", key)

	tests := map[string]struct {
		input http.Header
		want  struct {
			APIKey string
			err    error
		}
	}{
		"simple": {input: header, want: struct {
			APIKey string
			err    error
		}{
			APIKey: key,
		}},
		"no space": {input: header, want: struct {
			APIKey string
			err    error
		}{
			APIKey: key,
		}},
		"more than two parts": {input: header, want: struct {
			APIKey string
			err    error
		}{
			APIKey: key,
		}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			APIKey, err := GetAPIKey(header)
			got := struct {
				gotAPIKey string
				gotErr    error
			}{
				gotAPIKey: APIKey,
				gotErr:    err,
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expect: %v, got: %v", tc.want, got)
			}
		})
	}

}
