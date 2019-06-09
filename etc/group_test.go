package etc_test

import (
	"bytes"
	"testing"

	"github.com/nimkamp/pass-as-a-service/etc"
)

func TestParseGroup(t *testing.T) {
	wantEntries := []etc.GroupEntry{
		{
			Name:    "nick",
			GroupID: "123",
			Members: "bob",
		},
	}

	testGroupBytes := []byte(`nick:x:123:bob`)
	entries, err := etc.ParseGroup(bytes.NewReader(testGroupBytes))
	if err != nil {
		t.Fatal(err)
	}

	if len(entries) != 1 {
		t.Errorf("want 1 entries got %d", len(entries))
	}

	for i, got := range entries {
		want := wantEntries[i]
		if got.Name != want.Name {
			t.Errorf("want name %s got %s", got.Name, want.Name)
		}

		if got.GroupID != want.GroupID {
			t.Errorf("want uid %s got %s", got.GroupID, want.GroupID)
		}

		if got.Members != want.Members {
			t.Errorf("want members %s got %s", got.Members, want.Members)
		}
	}
}
