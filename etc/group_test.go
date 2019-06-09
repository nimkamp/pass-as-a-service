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
			GroupID: "1",
			Member:  "bob",
		},
	}

	testGroupBytes := []byte(`nick:x:1:bob`)
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
			t.Errorf("want gid %s got %s", got.GroupID, want.GroupID)
		}

		if got.Member != want.Member {
			t.Errorf("want Member %s got %s", got.Member, want.Member)
		}
	}
}

func TestFindGroupEntryByID(t *testing.T) {
	wantEntries := []etc.GroupEntry{
		{
			Name:    "nick",
			GroupID: "1",
			Member:  "bob",
		},
	}

	testGroupBytes := []byte(`nick:x:1:bob
James:x:5:world
Bob:x:7:jones`)
	entries, err := etc.ParseGroup(bytes.NewReader(testGroupBytes))
	if err != nil {
		t.Fatal(err)
	}

	entry, err := etc.FindGroupEntryByID("1", entries)
	if err != nil {
		t.Fatal(err)
	}

	want := wantEntries[0]

	if entry != want {
		t.Errorf("want entry %v got %v", wantEntries, entry)
	}

}
