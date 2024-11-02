package auth

import (
    "reflect"
    "testing"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
	header := make(http.Header)
	key := "this_is_the_key!"
	keyString := []string{"ApiKey "+key}
	header["Authorization"] = keyString
    ApiKey, err := GetAPIKey(header)
    if !reflect.DeepEqual(key, ApiKey) {
        t.Fatalf("expected: %v, got: %v", key, ApiKey)
    } else if err != nil {
		t.Fatalf("%v", err)
	}
}
