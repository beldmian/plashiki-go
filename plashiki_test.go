package plashikigo

import "testing"

func TestGetAnimeID(t *testing.T) {
	_, err := GetAnimeID("38000")
	if err != nil {
		t.Error(
			"For 38000",
			"got", "Request error: "+err.Error(),
		)
	}
}
