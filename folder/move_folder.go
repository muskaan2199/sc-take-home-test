package folder

import (
	"errors"
	"strings"
)

// MoveFolder moves a folder from its current location to a new parent folder within the same organization, based on folder names only.
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Locate the source and destination folders by name only
	var sourceFolder, destFolder *Folder
	for i := range f.folders {
		if f.folders[i].Name == name {
			sourceFolder = &f.folders[i]
		}
		if f.folders[i].Name == dst {
			destFolder = &f.folders[i]
		}
	}

	// Error handling for missing source or destination
	if sourceFolder == nil {
		return nil, errors.New("source folder does not exist")
	}
	if destFolder == nil {
		return nil, errors.New("destination folder does not exist")
	}

	// Ensure the source and destination are within the same organization
	if sourceFolder.OrgId != destFolder.OrgId {
		return nil, errors.New("cannot move a folder to a different organization")
	}

	// Prevent moving a folder to itself or one of its descendants
	if sourceFolder.Paths == destFolder.Paths  {
		return nil, errors.New("cannot move a folder to itself")
	}

	if strings.HasPrefix(destFolder.Paths, sourceFolder.Paths+".") {
		return nil, errors.New("cannot move a folder to a child of itself")
	}


	// Compute the new path for the source folder
	newPath := destFolder.Paths + "." + sourceFolder.Name

	// Update paths of source folder and all its descendants
	updatedFolders := make([]Folder, len(f.folders))
	copy(updatedFolders, f.folders)

	for i := range updatedFolders {
		// Only update folders within the subtree rooted at the source folder
		if updatedFolders[i].OrgId == sourceFolder.OrgId && strings.HasPrefix(updatedFolders[i].Paths, sourceFolder.Paths) {
			// Calculate new path by replacing the old root path with the new path
			updatedFolders[i].Paths = strings.Replace(updatedFolders[i].Paths, sourceFolder.Paths, newPath, 1)
		}
	}

	return updatedFolders, nil
}
