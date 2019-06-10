package etc

import (
	"fmt"
)

func FindGroupEntryByUserID(uid string, passwdEntries []PasswdEntry, groupEntries []GroupEntry) (GroupEntry, error) {
	for _, entryPasswd := range passwdEntries {
		if uid == entryPasswd.UserID {
			for _, entryGroup := range groupEntries {
				if entryPasswd.GroupID == entryGroup.GroupID {
					return entryGroup, nil
				}
			}
		}
	}

	return GroupEntry{}, fmt.Errorf("%s is not found", uid)
}
