package folder
import (
	"fmt"
	"strings"
)


func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Your code here...

	var source_idx, dest_idx int
    var source_found, dest_found bool

    for i, folder := range f.folders {
        if folder.Name == name {
            source_idx = i
            source_found = true
        }
        if folder.Name == dst {
            dest_idx = i
            dest_found = true
        }
    }

    if !source_found {
        return nil, fmt.Errorf("Source folder '%s' does not exist", name)
    }
    if !dest_found {
        return nil, fmt.Errorf("Destination folder '%s' does not exist", dst)
    }
    if f.folders[source_idx].OrgId != f.folders[dest_idx].OrgId {
        return nil, fmt.Errorf("Cannot move folder between different organizations")
    }
    if source_idx == dest_idx || strings.HasPrefix(f.folders[dest_idx].Paths, f.folders[source_idx].Paths + ".") {
        return nil, fmt.Errorf("Cannot move a folder into itself or its descendants")
    }

    res := make([]Folder, len(f.folders))
    copy(res, f.folders)

	// keep old path to replace for descendants
	old_path := res[source_idx].Paths

	// move old folder
	split_path := strings.Split(res[source_idx].Paths, ".")
	last := split_path[len(split_path) - 1]
	res[source_idx].Paths = res[dest_idx].Paths + "." + last

	new_path := res[source_idx].Paths
	for _, descendantIdx := range f.children[source_idx] {
		to_replace := res[descendantIdx].Paths
		updatedPath := strings.Replace(to_replace, old_path, new_path, 1)
		res[descendantIdx].Paths = updatedPath
	}

    return res, nil

}
