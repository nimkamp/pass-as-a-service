package etc_test

import (
	"bytes"
	"testing"

	"github.com/nimkamp/pass-as-a-service/etc"
)

func TestParsePasswd(t *testing.T) {
	wantEntries := []etc.PasswdEntry{
		{
			Name:    "nick",
			UserID:  "123",
			GroupID: "1",
			Comment: "",
			Home:    "/Users/home/nick",
			Shell:   "/bin/bash",
		},
	}

	testPasswdBytes := []byte(`nick:x:123:1::/Users/home/nick:/bin/bash`)
	entries, err := etc.ParsePasswd(bytes.NewReader(testPasswdBytes))
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

		if got.UserID != want.UserID {
			t.Errorf("want uid %s got %s", got.UserID, want.UserID)
		}

		if got.GroupID != want.GroupID {
			t.Errorf("want gid %s got %s", got.GroupID, want.GroupID)
		}

		if got.Comment != want.Comment {
			t.Errorf("want comment %s got %s", got.Comment, want.Comment)
		}

		if got.Home != want.Home {
			t.Errorf("want home %s got %s", got.Home, want.Home)
		}

		if got.Shell != want.Shell {
			t.Errorf("want shell %s got %s", got.Shell, want.Shell)
		}
	}
}

func TestFindPasswordEntryByID(t *testing.T) {
	wantEntries := []etc.PasswdEntry{
		{
			Name:    "nick",
			UserID:  "123",
			GroupID: "1",
			Comment: "",
			Home:    "/Users/home/nick",
			Shell:   "/bin/bash",
		},
	}

	testPasswdBytes := []byte(`nick:x:123:1::/Users/home/nick:/bin/bash
bob:x:456:1::/Users/home/bob:/bin/bash
george:x:789:1::/Users/home/george:/bin/bash`)
	entries, err := etc.ParsePasswd(bytes.NewReader(testPasswdBytes))
	if err != nil {
		t.Fatal(err)
	}

	entry, err := etc.FindPasswordEntryByID("123", entries)
	if err != nil {
		t.Fatal(err)
	}

	want := wantEntries[0]

	if entry != want {
		t.Errorf("want entry %v got %v", wantEntries, entry)
	}

}
