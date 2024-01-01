package services

import (
	"main/models"
	"reflect"
	"testing"
)

func TestDetectOutliers(t *testing.T) {
	testCases := []struct {
		name         string
		machines     []models.Machine
		wantOutliers []models.Machine
		wantErr      bool
	}{
		{
			name: "NoOutliers",
			machines: []models.Machine{
				{ID: "1", Age: "2 years"},
				{ID: "2", Age: "1 year"},
				{ID: "3", Age: "10 months"},
				{ID: "4", Age: "16 months"},
				{ID: "5", Age: "1 months"},
				{ID: "6", Age: "1 day"},
				{ID: "7", Age: "57 days"},
			},
			wantOutliers: []models.Machine{},
			wantErr:      false,
		},
		{
			name: "WithOutliers",
			machines: []models.Machine{
				{ID: "1", Age: "2 years"},
				{ID: "2", Age: "1 year"},
				{ID: "3", Age: "10 months"},
				{ID: "4", Age: "16 months"},
				{ID: "5", Age: "1000 months"},
				{ID: "6", Age: "100 years"},
				{ID: "7", Age: "57 days"},
				{ID: "8", Age: "3 month"},
				{ID: "9", Age: "10 weeks"},
			},
			wantOutliers: []models.Machine{
				{ID: "5", Age: "1000 months"},
				{ID: "6", Age: "100 years"},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			outliers, err := DetectOutliers(tc.machines)
			if (err != nil) != tc.wantErr {
				t.Errorf("DetectOutliers() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(outliers, tc.wantOutliers) {
				t.Errorf("DetectOutliers() = %v, want %v", outliers, tc.wantOutliers)
			}
		})
	}
}
