package vul

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVulnerabilities_Check(t *testing.T) {
	tests := []struct {
		name            string
		vulnerabilities Vulnerabilities
		wantErr         bool
	}{
		{
			name: "test",
			vulnerabilities: Vulnerabilities{
				&BaseVulnerability{
					Name:        "test vul",
					Description: "description",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.vulnerabilities.Check(context.TODO()); (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
			for _, v := range tt.vulnerabilities {
				assert.True(t, v.(*BaseVulnerability).CheckSecHaveRan)
			}
		})
	}
}
