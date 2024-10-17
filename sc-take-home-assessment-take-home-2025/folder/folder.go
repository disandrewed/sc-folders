package folder

import "github.com/gofrs/uuid"
import (
    "strings"
)

type IDriver interface {
	// GetFoldersByOrgID returns all folders that belong to a specific orgID.
	GetFoldersByOrgID(orgID uuid.UUID) []Folder
	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.
	GetAllChildFolders(orgID uuid.UUID, name string) []Folder

	// component 2
	// Implement the following methods:
	// MoveFolder moves a folder to a new destination.
	MoveFolder(name string, dst string) ([]Folder, error)
}

type driver struct {
	// define attributes here
	// data structure to store folders
	// or preprocessed data

	// example: feel free to change the data structure, if slice is not what you want

	// List of folders
	folders []Folder

	// Map of parent folder index to child folder indices
    children map[int][]int
}

func NewDriver(folders []Folder) IDriver {
	children_map := make(map[int][]int)

	var n int = len(folders)
    for i := 0; i < n; i++ {
        parent_path := folders[i].Paths
		parent_org := folders[i].OrgId
		children_map[i] = []int{}
        for j := 0; j < n; j++ {
            child_path := folders[j].Paths
			child_org := folders[j].OrgId
			
			// if the parent path is a prefix of child path, 
            if strings.HasPrefix(child_path, parent_path + ".") && parent_org == child_org {
				children_map[i] = append(children_map[i], j)
            }
        }
    }

	return &driver{
		// initialize attributes here
		folders: folders,
		children: children_map,
	}
}