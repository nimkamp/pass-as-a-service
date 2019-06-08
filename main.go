package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/pass-as-a-service/etc"
)

func main() {
	log.Println("\nFrom byte array")
	for _, entry := range fromByteArrayExample() {
		log.Printf("%v", entry)
	}

	log.Println("\nFrom file")
	for _, entry := range fileExample() {
		log.Printf("%v", entry)
	}

}

func fromByteArrayExample() []etc.PasswdEntry {
	var testBytes = []byte(`nick:x:123:1::/Users/home/nick:/bin/bash
george:x:321:2::/Users/home/george:/bin/sh
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
