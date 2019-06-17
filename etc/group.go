package etc

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ParseGroup(r io.Reader) ([]GroupEntry, error) {
	var entries []GroupEntry

	lr := bufio.NewReader(r)

	for {
		line, _, err := lr.ReadLine()
		if err != nil {
			break
		}

		splitLine := strings.Split(string(line), ":")
		if len(splitLine) != 4 {
			return nil, fmt.Errorf("etcgroup.Parse: expected 3 semi colons on parsed line but got %d", len(splitLine))
		}

		e := GroupEntry{
			Name:    splitLine[0],
			GroupID: splitLine[2],
			Member:  splitLine[3],
		}

		entries = append(entries, e)
	}

	return entries, nil
}

func FindGroupEntryByID(gid string, groupEntries []GroupEntry) (GroupEntry, error) {
	for _, entryGroup := range groupEntries {
		if entryGroup.GroupID == gid {
			return entryGroup, nil
		}
	}

	return GroupEntry{}, fmt.Errorf("%s is not found", gid)
}

func GetGroupByQuery(name string, gid string, member string, entries []GroupEntry) ([]GroupEntry, error) {
	var matchedEntries []GroupEntry

	for _, entry := range entries {
		if name != "" && entry.Name != name {
			continue
		}

		if gid != "" && entry.GroupID != gid {
			continue
		}

		if member != "" && entry.Member != member {
			continue
		}

		matchedEntries = append(matchedEntries, entry)
	}

	return matchedEntries, nil
}

type GroupEntry struct {
	Name    string
	GroupID string
	Member  string
}
