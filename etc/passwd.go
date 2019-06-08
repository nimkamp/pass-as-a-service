package etc

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ParsePasswd(r io.Reader) ([]PasswdEntry, error) {
	var entries []PasswdEntry

	lr := bufio.NewReader(r)

	for {
		line, _, err := lr.ReadLine()
		if err != nil {
			break
		}

		splitLine := strings.Split(string(line), ":")
		if len(splitLine) != 7 {
			return nil, fmt.Errorf("etcpasswd.Parse: expected 6 semi colons on parsed line but got %d", len(splitLine))
		}

		e := PasswdEntry{
			Name:    splitLine[0],
			UserID:  splitLine[2],
			GroupID: splitLine[3],
			Comment: splitLine[4],
			Home:    splitLine[5],
			Shell:   splitLine[6],
		}

		entries = append(entries, e)
	}

	return entries, nil
}

func FindPasswordEntryByID(uid string, entries []PasswdEntry) (PasswdEntry, error) {
	for _, entry := range entries {
		if uid == entry.UserID {
			return entry, nil
		}
	}

	return PasswdEntry{}, fmt.Errorf("%s is not found", uid)
}

type PasswdEntry struct {
	Name    string
	UserID  string
	GroupID string
	Comment string
	Home    string
	Shell   string
}
