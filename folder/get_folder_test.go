package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	// Sample data for testing
	sampleFolders := []folder.Folder{
		{Name: "creative-scalphunter", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter"},
		{Name: "clear-arclight", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter.clear-arclight"},
		{Name: "steady-insect", OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"), Paths: "steady-insect"},
	}

	tests := []struct {
		name    string
		orgID   uuid.UUID
		want    []folder.Folder
	}{
		{
			name:  "Retrieve folders for orgID with folders",
			orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			want: []folder.Folder{
				{Name: "creative-scalphunter", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter"},
				{Name: "clear-arclight", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter.clear-arclight"},
			},
		},
		{
			name:  "Retrieve folders for orgID with no folders",
			orgID: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000000"),
			want:  []folder.Folder{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(sampleFolders)
			got := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetAllChildFolders(t *testing.T) {
	t.Parallel()

	// Sample data for testing
	sampleFolders := []folder.Folder{
		{Name: "creative-scalphunter", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter"},
		{Name: "clear-arclight", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter.clear-arclight"},
		{Name: "topical-micromax", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter.clear-arclight.topical-micromax"},
		{Name: "steady-insect", OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"), Paths: "steady-insect"},
	}

	tests := []struct {
		name          string
		orgID         uuid.UUID
		parentFolder  string
		want          []folder.Folder
		wantErr       bool
		expectedErrMsg string
	}{
		{
			name:         "Valid parent folder with children",
			orgID:        uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			parentFolder: "creative-scalphunter",
			want: []folder.Folder{
				{Name: "clear-arclight", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter.clear-arclight"},
				{Name: "topical-micromax", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter.clear-arclight.topical-micromax"},
			},
			wantErr: false,
		},
		{
			name:          "Parent folder with no children",
			orgID:         uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			parentFolder:  "topical-micromax",
			want:          []folder.Folder(nil),
			wantErr:       false,
		},
		{
			name:          "Nonexistent parent folder",
			orgID:         uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			parentFolder:  "invalid-folder",
			want:          nil,
			wantErr:       true,
			expectedErrMsg: "Folder does not exist",
		},
		{
			name:          "Valid folder but in a different organization",
			orgID:         uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
			parentFolder:  "creative-scalphunter",
			want:          nil,
			wantErr:       true,
			expectedErrMsg: "Folder does not exist in the specified organization",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(sampleFolders)
			got, err := f.GetAllChildFolders(tt.orgID, tt.parentFolder)

			if tt.wantErr {
				assert.Error(t, err)
				// Check that the error message matches the expected message
				assert.EqualError(t, err, tt.expectedErrMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
 