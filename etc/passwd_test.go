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
			t.Errorf("want name %s got %s", want.Name, got.Name)
		}

		if got.UserID != want.UserID {
			t.Errorf("want uid %s got %s", want.UserID, got.UserID)
		}

		if got.GroupID != want.GroupID {
			t.Errorf("want gid %s got %s", want.GroupID, got.GroupID)
		}

		if got.Comment != want.Comment {
			t.Errorf("want comment %s got %s", want.Comment, got.Comment)
		}

		if got.Home != want.Home {
			t.Errorf("want home %s got %s", want.Home, got.Home)
		}

		if got.Shell != want.Shell {
			t.Errorf("want shell %s got %s", want.Shell, got.Shell)
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

	entry, err := etc.FindPasswdEntryByID("123", entries)
	if err != nil {
		t.Fatal(err)
	}

	want := wantEntries[0]

	if entry != want {
		t.Errorf("want entry %v got %v", wantEntries, entry)
	}
}

func TestFindPasswdEntryByWrongID(t *testing.T) {
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
jacob:x:111:2::/Users/home/jacob:/bin/bash
george:x:789:1::/Users/home/george:/bin/bash`)
	entries, err := etc.ParsePasswd(bytes.NewReader(testPasswdBytes))
	if err != nil {
		t.Fatal(err)
	}

	entry, err := etc.FindPasswdEntryByID("21", entries)
	if err == nil {
		t.Fatal(err)
	}

	want := wantEntries[0]

	if entry == want {
		t.Errorf("want entry %v got %v", want, entry)
	}
}

func TestFindPasswdEntryByEmptyID(t *testing.T) {
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
jacob:x:111:2::/Users/home/jacob:/bin/bash
george:x:789:1::/Users/home/george:/bin/bash`)
	entries, err := etc.ParsePasswd(bytes.NewReader(testPasswdBytes))
	if err != nil {
		t.Fatal(err)
	}

	entry, err := etc.FindPasswdEntryByID("", entries)
	if err == nil {
		t.Fatal(err)
	}

	want := wantEntries[0]

	if entry == want {
		t.Errorf("want entry %v got %v", want, entry)
	}
}

func TestGetPasswdByQuery(t *testing.T) {
	wantEntries := []etc.PasswdEntry{
		{
			Name:    "chris",
			UserID:  "999",
			GroupID: "7",
			Comment: "Hello",
			Home:    "/Users/home/chris",
			Shell:   "/bin/bash",
		},
	}

	testPasswdBytes := []byte(`nick:x:123:1:Hello:/Users/home/nick:/bin/bash
chris:x:999:7:Hello:/Users/home/chris:/bin/bash
george:x:444:9:Hello:/Users/home/george:/bin/bash`)
	entries, err := etc.ParsePasswd(bytes.NewReader(testPasswdBytes))
	if err != nil {
		t.Fatal(err)
	}

	passwdQueryEntry, err := etc.GetPasswdByQuery("", "999", "", "", "", "", entries)

	for i, got := range passwdQueryEntry {
		want := wantEntries[i]
		if got.Name != want.Name {
			t.Errorf("want name %s got %s", want.Name, got.Name)
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

func TestGetPasswdByInvalidQuery(t *testing.T) {
	wantEntries := []etc.PasswdEntry{
		{
			Name:    "",
			UserID:  "",
			GroupID: "",
			Comment: "",
			Home:    "",
			Shell:   "",
		},
	}

	testPasswdBytes := []byte(`nick:x:123:1:Hello:/Users/home/nick:/bin/bash
chris:x:999:7:Hello:/Users/home/chris:/bin/bash
george:x:444:9:Hello:/Users/home/george:/bin/bash`)
	entries, err := etc.ParsePasswd(bytes.NewReader(testPasswdBytes))
	if err != nil {
		t.Fatal(err)
	}

	passwdQueryEntry, err := etc.GetPasswdByQuery("bob", "", "", "", "", "", entries)

	for i, got := range passwdQueryEntry {
		want := wantEntries[i]
		if got.Name == want.Name {
			t.Errorf("want name %s got %s", want.Name, got.Name)
		}

		if got.UserID == want.UserID {
			t.Errorf("want uid %s got %s", want.UserID, got.UserID)
		}

		if got.GroupID == want.GroupID {
			t.Errorf("want gid %s got %s", want.GroupID, got.GroupID)
		}

		if got.Comment == want.Comment {
			t.Errorf("want comment %s got %s", want.Comment, got.Comment)
		}

		if got.Home == want.Home {
			t.Errorf("want home %s got %s", want.Home, got.Home)
		}

		if got.Shell == want.Shell {
			t.Errorf("want shell %s got %s", want.Shell, got.Shell)
		}
	}
}

