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
			Members: splitLine[3],
		}

		entries = append(entries, e)
	}

	return entries, nil
}

type GroupEntry struct {
	Name    string
	GroupID string
	Members string
}
