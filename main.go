package main

import (
	"fmt"
	"log"
	"github.com/georgechieng-sc/interns-2022/folder"
	//"github.com/gofrs/uuid"
)

func main() {
	// Set the orgID to a valid UUID from the sample data
	//orgID := uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")

	// Retrieve all folders
	res := folder.GetAllFolders()

	// Initialize the folder driver
	folderDriver := folder.NewDriver(res)

	// Retrieve folders by organization ID
	//orgFolder := folderDriver.GetFoldersByOrgID(orgID)

	// Print all folders
	// fmt.Println("All Folders:")
	// folder.PrettyPrint(res)
	
	// // Print folders for the specified orgID
	// fmt.Printf("\nFolders for orgID: %s\n", orgID)
	// folder.PrettyPrint(orgFolder)

	// Test case 1: Valid parent folder with children
	// fmt.Println("\nTest Case 1: Valid parent folder 'creative-scalphunter' with children")
	// parentFolderName := "creative-scalphunter"
	// childFolders, err := folderDriver.GetAllChildFolders(orgID, parentFolderName)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Printf("Child folders of '%s' for orgID: %s\n", parentFolderName, orgID)
	// 	folder.PrettyPrint(childFolders)
	// }

	//var numbers []int
	//fmt.Println(numbers)

	// // Test case 2: Valid parent folder with no children
	// fmt.Println("\nTest Case 2: Valid parent folder 'steady-insect' with no children")
	// parentFolderName = "steady-insect"
	// childFolders, err = folderDriver.GetAllChildFolders(orgID, parentFolderName)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Printf("Child folders of '%s' for orgID: %s\n", parentFolderName, orgID)
	// 	folder.PrettyPrint(childFolders)
	// }

	// // Test case 3: Nonexistent folder
	// fmt.Println("\nTest Case 3: Nonexistent folder 'invalid-folder'")
	// parentFolderName = "invalid-folder"
	// childFolders, err = folderDriver.GetAllChildFolders(orgID, parentFolderName)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Printf("Child folders of '%s' for orgID: %s\n", parentFolderName, orgID)
	// 	folder.PrettyPrint(childFolders)
	// }
updatedFolders, err := folderDriver.MoveFolder("creative-scalphunter", "nearby-secret")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Display the updated folder paths after the move
	for _, f := range updatedFolders {
		fmt.Printf("Folder: %s, Path: %s\n", f.Name, f.Paths)
	}

	
}
