package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllChildFolders(t *testing.T) {
	// Define sample folders

	orgID, _ := uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")
	orgID2, _ := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")

	sampleFolders := []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: orgID},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: orgID},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgID},
		{Name: "delta", Paths: "alpha.delta", OrgId: orgID},
		{Name: "echo", Paths: "echo", OrgId: orgID},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: orgID2},
	}

	driver := folder.NewDriver(sampleFolders)

	tests := []struct {
		orgID   uuid.UUID
		name    string
		want    []folder.Folder
		wantErr bool
	}{
		{
			orgID: orgID,
			name:  "alpha",
			want: []folder.Folder{
				{Name: "bravo", Paths: "alpha.bravo", OrgId: orgID},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgID},
				{Name: "delta", Paths: "alpha.delta", OrgId: orgID},
			},
			wantErr: false,
		},
		{
			orgID: orgID,
			name:  "bravo",
			want: []folder.Folder{
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgID},
			},
			wantErr: false,
		},
		{
			orgID: orgID,
			name:  "charlie",
			want:  []folder.Folder{}, // No children for "charlie"
			wantErr: false,
		},
		{
			orgID: orgID,
			name:  "invalid_folder",
			want:  []folder.Folder{}, // Folder does not exist
			wantErr: true,
		},
		{
			orgID: orgID,
			name:  "foxtrot",
			want:  []folder.Folder{}, // Folder does not exist in this org
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				assert.Panics(t, func() {
					driver.GetAllChildFolders(tt.orgID, tt.name)
				}, "Expected panic, but none occurred")
			} else {
				got := driver.GetAllChildFolders(tt.orgID, tt.name)
				assert.Equal(t, tt.want, got, "Expected and actual result do not match")
			}
		})
	}
}

