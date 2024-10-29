package folder

import (
	"errors"
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

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	allFolders := f.folders
	folders := f.GetFoldersByOrgID(orgID)

	// Find the full path of the specified parent folder
	var parentFolderPath string
	for _, folder := range folders {
		if folder.Name == name {
			parentFolderPath = folder.Paths
			break
		}
	}

	// If the parent folder is not found, return an error
	if parentFolderPath == "" {
		var exists bool
		for _, folder := range allFolders {
			if folder.Name == name {
				exists = true
				break
			}
		}
		if exists == true {
			return nil, errors.New("Folder does not exist in the specified organization")
		} else {
			return nil, errors.New("Folder does not exist")
		}
		
	}

	// Initialize a slice to store all descendant folders (children, grandchildren, etc.)
	var res []Folder

	// Match all folders whose path starts with the full parent path
	for _, folder := range folders {
		// Check that folder is not the parent itself and that it's a descendant of the parent path
		if folder.Paths != parentFolderPath && strings.HasPrefix(folder.Paths, parentFolderPath+".") {
			res = append(res, folder)
		}
	}

	// Return the list of all descendant folders
	return res, nil
}
