package folder

import (
	"strings"
	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			res = append(res, folder)
		}
	}

	return res
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	folders := f.GetFoldersByOrgID(orgID)  // Get all folders for the specified orgID
	var res []Folder

	for _, folder := range folders {
		// Check if the folder path starts with the pattern "<name>."
		if folder.Paths != name && strings.HasPrefix(folder.Paths, name+".") {
			res = append(res, folder)
		}
	}

	return res
}
