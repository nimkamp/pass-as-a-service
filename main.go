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

	fmt.Println(passwdFile)
	fmt.Println(groupFile)

	var pTestBytes = []byte(`nick:x:1:1::/Users/home/nick:/bin/bash
george:x:3:2::/Users/home/george:/bin/sh
docker:x:321:1002::/Users/home/george:/bin/sh
`)
	var gTestBytes = []byte(`nick:x:1:bob
george:x:3:321
docker:x:1002:dwoodlins
`)

	pEntries, err := etc.ParsePasswd(bytes.NewReader(pTestBytes))
	if err != nil {
		log.Fatal(err)
	}

	gEntries, err := etc.ParseGroup(bytes.NewReader(gTestBytes))
	if err != nil {
		log.Fatal(err)
	}

	pEntriesJson, err := json.Marshal(pEntries)
	if err != nil {
		fmt.Println(err)
	}

	gEntriesJson, err := json.Marshal(gEntries)
	if err != nil {
		fmt.Println(err)
	}
	// log.Println("\nFrom byte array")
	// for _, entry := range fromByteArrayExample() {
	// 	log.Printf("%v", entry)
	// }

	// log.Println("\nFrom file")
	// for _, entry := range fileExample() {
	// 	log.Printf("%v", entry)
	// }

	// log.Println("\n Group From byte array")
	// for _, entry := range fromByteArrayExampleGroup() {
	// 	log.Printf("%v", entry)
	// }

	// log.Println("\n Group From file")
	// for _, entry := range fileExampleGroup() {
	// 	log.Printf("%v", entry)
	// }

	// log.Println("\n Return a user based off uid (GET /users/<uid>) byte array")
	// pEntry, err := etc.FindPasswdEntryByID("321", fromByteArrayExample())
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// log.Printf("%v\n", pEntry)

	// log.Println("\n Return all groups for a given user (GET /groups/gid)")
	// gEntry, err := etc.FindGroupEntryByID("3", fromByteArrayExampleGroup())
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// log.Printf("%v\n", gEntry)

	// log.Println("\n Return all groups for a given user (GET /users/<uid>/groups)")
	// gEntry, err := etc.FindGroupEntryByUserID("321", fromByteArrayExample(), fromByteArrayExampleGroup())
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// log.Printf("%v\n", gEntry)

	// log.Println("\n Users with group id=1002")
	// pEntry, err := etc.GetPasswdByQuery("docker", "", "1002", "", "", "", fromByteArrayExample())
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// log.Printf("%v\n", pEntry)

	// log.Println("\n Groups with group id=1002")
	// pEntry, err := etc.GetGroupByQuery("george", "", "", fromByteArrayExampleGroup())
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// log.Printf("%v\n", pEntry)

	r := chi.NewRouter()
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(pEntriesJson))
	})

	r.Get("/users/{uid}", func(w http.ResponseWriter, r *http.Request) {
		uid := chi.URLParam(r, "uid")
		entry, err := etc.FindPasswdEntryByID(uid, pEntries)
		if err != nil {
			http.Error(w, "User ID does not exist", 404)
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

	r.Get("/groups", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(gEntriesJson))
	})

	r.Get("/groups/{gid}", func(w http.ResponseWriter, r *http.Request) {
		gid := chi.URLParam(r, "gid")
		entry, err := etc.FindGroupEntryByID(gid, gEntries)
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
		entry, err := etc.GetPasswdByQuery(queryValues.Get("name"), queryValues.Get("uid"), queryValues.Get("gid"), queryValues.Get("comment"), queryValues.Get("home"), queryValues.Get("shell"), pEntries)
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
		entry, err := etc.GetGroupByQuery(queryValues.Get("name"), queryValues.Get("gid"), queryValues.Get("member"), gEntries)
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

func fileExample() []etc.PasswdEntry {
	b, err := ioutil.ReadFile("/home/nimkamp/go/src/github.com/pass-as-a-service/fakeetcpasswd")
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

func fileExampleGroup() []etc.GroupEntry {
	b, err := ioutil.ReadFile("/home/nimkamp/go/src/github.com/pass-as-a-service/fakeetcgroup")
	if err != nil {
		log.Fatal(err)
	}

	etcGroupEntries, err := etc.ParseGroup(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	return etcGroupEntries
}
