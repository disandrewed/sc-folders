package folder_test

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/georgechieng-sc/interns-2022/folder"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()

	orgID, _ := uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")
	orgID2, _ := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")

	initialFolders := []folder.Folder{
		{Name: "alpha", OrgId: orgID, Paths: "alpha"},
		{Name: "bravo", OrgId: orgID, Paths: "alpha.bravo"},
		{Name: "charlie", OrgId: orgID, Paths: "alpha.bravo.charlie"},
		{Name: "delta", OrgId: orgID, Paths: "alpha.delta"},
		{Name: "echo", OrgId: orgID, Paths: "alpha.delta.echo"},
		{Name: "foxtrot", OrgId: orgID2, Paths: "foxtrot"},
		{Name: "golf", OrgId: orgID, Paths: "golf"},
	}

	tests := []struct {
		name        string
		srcFolder   string
		destFolder  string
		expected    []folder.Folder
		expectError bool
	}{
		{
			name:       "Move bravo under delta",
			srcFolder:  "bravo",
			destFolder: "delta",
			expected: []folder.Folder{
				{Name: "alpha", OrgId: orgID, Paths: "alpha"},
				{Name: "bravo", OrgId: orgID, Paths: "alpha.delta.bravo"},
				{Name: "charlie", OrgId: orgID, Paths: "alpha.delta.bravo.charlie"},
				{Name: "delta", OrgId: orgID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: orgID, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: orgID2, Paths: "foxtrot"},
				{Name: "golf", OrgId: orgID, Paths: "golf"},
			},
			expectError: false,
		},
		{
			name:       "Move bravo under golf",
			srcFolder:  "bravo",
			destFolder: "golf",
			expected: []folder.Folder{
				{Name: "alpha", OrgId: orgID, Paths: "alpha"},
				{Name: "bravo", OrgId: orgID, Paths: "golf.bravo"},
				{Name: "charlie", OrgId: orgID, Paths: "golf.bravo.charlie"},
				{Name: "delta", OrgId: orgID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: orgID, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: orgID2, Paths: "foxtrot"},
				{Name: "golf", OrgId: orgID, Paths: "golf"},
			},
			expectError: false,
		},
		{
			name:        "Move bravo under charlie (invalid)",
			srcFolder:   "bravo",
			destFolder:  "charlie",
			expectError: true,
		},
		{
			name:        "Move bravo under foxtrot (different org)",
			srcFolder:   "bravo",
			destFolder:  "foxtrot",
			expectError: true,
		},
		{
			name:        "Move invalid_folder under delta",
			srcFolder:   "invalid_folder",
			destFolder:  "delta",
			expectError: true,
		},
		{
			name:        "Move bravo under invalid_folder",
			srcFolder:   "bravo",
			destFolder:  "invalid_folder",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(initialFolders)
			result, err := f.MoveFolder(tt.srcFolder, tt.destFolder)
			if tt.expectError {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.ElementsMatch(t, tt.expected, result, "Expected and actual folder structures do not match")
			}
		})
	}
}