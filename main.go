package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	// Retrieve all folders
	res := folder.GetAllFolders()

	// Initialize the folder driver
	folderDriver := folder.NewDriver(res)

	// Retrieve folders by organization ID
	orgFolder := folderDriver.GetFoldersByOrgID(orgID)

	// Print all folders
	fmt.Println("All Folders:")
	folder.PrettyPrint(res)
	
	// Print folders for the specified orgID
	fmt.Printf("\nFolders for orgID: %s\n", orgID)
	folder.PrettyPrint(orgFolder)

	// Example usage of GetAllChildFolders
	parentFolderName := "stunning-horridus" // Specify the name of the parent folder
	childFolders := folderDriver.GetAllChildFolders(orgID, parentFolderName)
	
	// Print child folders for the specified parent folder
	fmt.Printf("\nChild folders of '%s' for orgID: %s\n", parentFolderName, orgID)
	folder.PrettyPrint(childFolders)
}
