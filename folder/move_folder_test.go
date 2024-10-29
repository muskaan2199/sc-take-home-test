package folder_test

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/georgechieng-sc/interns-2022/folder"
)

func TestMoveFolder(t *testing.T) {
	t.Parallel()

	// Sample UUID to represent the organization
	orgID := uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")

	// Initialize sample folders for testing
	initialFolders := []folder.Folder{
		{Name: "root", OrgId: orgID, Paths: "root"},
		{Name: "sub1", OrgId: orgID, Paths: "root.sub1"},
		{Name: "sub2", OrgId: orgID, Paths: "root.sub2"},
		{Name: "sub3", OrgId: orgID, Paths: "root.sub1.sub3"},
	}

	// Create the driver with these folders
	d := folder.NewDriver(initialFolders)

	// Define test cases
	tests := []struct {
		name           string
		source         string
		destination    string
		expectError    bool
		expectedPath   string
		expectedErrors string
	}{
		// Test successful move of sub1 under sub2
		{
			name:         "Move sub1 under sub2",
			source:       "sub1",
			destination:  "sub2",
			expectError:  false,
			expectedPath: "root.sub2.sub1",
		},
		// Test invalid move (sub1 into its own descendant sub3)
		{
			name:           "Move sub1 under sub3 (invalid descendant)",
			source:         "sub1",
			destination:    "sub3",
			expectError:    true,
			expectedErrors: "cannot move a folder to a child of itself",
		},

		// Test invalid move (sub1 into itself)
		{
			name:           "Move sub1 under sub1",
			source:         "sub1",
			destination:    "sub1",
			expectError:    true,
			expectedErrors: "cannot move a folder to itself",
		},
		// Test move to a nonexistent destination
		{
			name:           "Move sub1 under non-existing folder",
			source:         "sub1",
			destination:    "nonexistent",
			expectError:    true,
			expectedErrors: "destination folder does not exist",
		},
		// Test move for nonexistent source folder
		{
			name:           "Move non-existing source to root",
			source:         "nonexistent",
			destination:    "root",
			expectError:    true,
			expectedErrors: "source folder does not exist",
		},
	}

	// Execute test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updatedFolders, err := d.MoveFolder(tt.source, tt.destination)

			if tt.expectError {
				// Assert an error was expected
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErrors)
				return
			}

			// Assert no error was expected
			assert.NoError(t, err)

			// Verify the source folder's path was updated correctly in the result
			found := false
			for _, f := range updatedFolders {
				if f.Name == tt.source && f.Paths == tt.expectedPath {
					found = true
					break
				}
			}

			// Check if expected path was found for the source folder
			assert.True(t, found, "expected path of %s to be updated to %s, but it was not found", tt.source, tt.expectedPath)
		})
	}
}
