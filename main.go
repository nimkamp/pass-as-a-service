package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/nimkamp/pass-as-a-service/etc"
)

func main() {
	passwdFile := os.Args[1]
	groupFile := os.Args[2]

	passwdEntries := fileExamplePasswd(passwdFile)
	groupEntries := fileExampleGroup(groupFile)

	passwdEntriesJson, err := json.Marshal(passwdEntries)
	if err != nil {
		fmt.Println(err)
	}

	groupEntriesJson, err := json.Marshal(groupEntries)
	if err != nil {
		fmt.Println(err)
	}

	r := chi.NewRouter()
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(passwdEntriesJson))
	})

	r.Get("/users/{uid}", func(w http.ResponseWriter, r *http.Request) {
		uid := chi.URLParam(r, "uid")
		entry, err := etc.FindPasswdEntryByID(uid, passwdEntries)
		if err != nil {
			http.Error(w, "User ID does not exist", 404)
			return
		}

		jsonEntry, err := json.Marshal(&entry)
		if err != nil {
			http.Error(w, "Unable to parse json", 406)
			return

		}
		w.Write(jsonEntry)
		return
	})

	r.Get("/groups", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(groupEntriesJson))
	})

	r.Get("/groups/{gid}", func(w http.ResponseWriter, r *http.Request) {
		gid := chi.URLParam(r, "gid")
		entry, err := etc.FindGroupEntryByID(gid, groupEntries)
		if err != nil {
			http.Error(w, "Group ID does not exist", 404)
			return
		}

		jsonEntry, err := json.Marshal(&entry)
		if err != nil {
			// http.Error(w, err.Error(), err.Status())
			return

		}
		w.Write(jsonEntry)
		return
	})

	r.Get("/users/query", func(w http.ResponseWriter, r *http.Request) {
		queryValues := r.URL.Query()
		fmt.Fprintf(w, "hello, %s!\n", queryValues.Get("name"))
		entry, err := etc.GetPasswdByQuery(queryValues.Get("name"), queryValues.Get("uid"), queryValues.Get("gid"), queryValues.Get("comment"), queryValues.Get("home"), queryValues.Get("shell"), passwdEntries)
		if err != nil {
			http.Error(w, "User does not exist", 404)
			return
		}

		jsonEntry, err := json.Marshal(entry)
		if err != nil {
			fmt.Println(err)
		}

		w.Write(jsonEntry)
	})

	r.Get("/groups/query", func(w http.ResponseWriter, r *http.Request) {
		queryValues := r.URL.Query()
		fmt.Fprintf(w, "hello, %s!\n", queryValues.Get("name"))
		entry, err := etc.GetGroupByQuery(queryValues.Get("name"), queryValues.Get("gid"), queryValues.Get("member"), groupEntries)
		if err != nil {
			http.Error(w, "Group does not exist", 404)
			return
		}

		jsonEntry, err := json.Marshal(entry)
		if err != nil {
			fmt.Println(err)
		}

		w.Write(jsonEntry)
	})

	http.ListenAndServe(":3000", r)
}

func fromByteArrayExample() []etc.PasswdEntry {
	var testBytes = []byte(`nick:x:1:1::/Users/home/nick:/bin/bash
george:x:3:2::/Users/home/george:/bin/sh
docker:x:321:1002::/Users/home/george:/bin/sh
`)

	etcPasswdEntries, err := etc.ParsePasswd(bytes.NewReader(testBytes))
	if err != nil {
		log.Fatal(err)
	}

	return etcPasswdEntries
}

func fileExamplePasswd(file string) []etc.PasswdEntry {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	etcPasswdEntries, err := etc.ParsePasswd(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	return etcPasswdEntries
}

func fromByteArrayExampleGroup() []etc.GroupEntry {
	var testBytes = []byte(`nick:x:1:bob
george:x:3:321
docker:x:1002:dwoodlins
`)

	etcGroupEntries, err := etc.ParseGroup(bytes.NewReader(testBytes))
	if err != nil {
		log.Fatal(err)
	}

	return etcGroupEntries
}

func fileExampleGroup(file string) []etc.GroupEntry {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	etcGroupEntries, err := etc.ParseGroup(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	return etcGroupEntries
}
