package templatesCache

import "testing"

func TestCreate(t *testing.T) {
	templatesCache, err := Create()
	if err != nil {
		t.Error("error by cache creating")
	}

	if templatesCache == nil {
		t.Error("templates cache is nil")
	}

	for _, value := range templatesCache {
		if value == nil {
			t.Error("one of the items in templates cache is nil")
		}
	}
}
