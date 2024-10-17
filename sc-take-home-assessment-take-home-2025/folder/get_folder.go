package folder

import "github.com/gofrs/uuid"
import(
	"fmt"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Your code here...

	var parent_idx int
	var found bool
	for idx, folder := range f.folders {
		if folder.Name == name && folder.OrgId == orgID {
			parent_idx = idx
			found = true
			break
		}
	}
	
	if !found {
		panic(fmt.Sprintf("folder does not exist in the specified organization"))
	}
	res := []Folder{}

	for _, elem := range f.children[parent_idx] {
		res = append(res, f.folders[elem])
	}
	if res == nil {
        return []Folder{}
    }
	return res
}
