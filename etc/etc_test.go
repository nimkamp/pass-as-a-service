package etc_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/nimkamp/pass-as-a-service/etc"
)

func TestFindGroupEntryByUserID(t *testing.T) {
	wantEntries := []etc.GroupEntry{
		{
			Name:    "docker",
			GroupID: "1002",
			Member:  "dwoodlins",
		},
	}

	var testPasswdBytes = []byte(`nick:x:1:1::/Users/home/nick:/bin/bash
george:x:321:2::/Users/home/george:/bin/sh
docker:x:321:1002::/Users/home/george:/bin/sh
`)

	var testGroupBytes = []byte(`nick:x:1:bob
george:x:3:321
docker:x:1002:dwoodlins
`)

	passwdEntries, err := etc.ParsePasswd(bytes.NewReader(testPasswdBytes))
	if err != nil {
		t.Fatal(err)
	}

	groupEntries, err := etc.ParseGroup(bytes.NewReader(testGroupBytes))
	if err != nil {
		t.Fatal(err)
	}

	got, err := etc.FindGroupEntryByUserID("321", passwdEntries, groupEntries)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, wantEntries[0]) {
		t.Errorf("want entry %v got %v", wantEntries, got)
	}
}
