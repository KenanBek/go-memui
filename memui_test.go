package memui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DummtType struct {
	Name string
}

func TestMemUI_Register(t *testing.T) {
	type checks func(t *testing.T, mui *MemUI) bool
	tests := []struct {
		name    string
		args    []interface{}
		wantErr bool
		checks  checks
	}{
		{"add a pointer object", []interface{}{&DummtType{"Test"}}, false, func(t *testing.T, mui *MemUI) bool {
			tn := "*memui.DummtType"
			mo := mui.memObjects
			return assert.Len(t, mo, 1) && assert.Contains(t, mo, tn) && assert.Len(t, mo[tn], 1)
		}},
		{"add a value object", []interface{}{DummtType{"Test"}}, true, func(t *testing.T, mui *MemUI) bool {
			return assert.Len(t, mui.memObjects, 0)
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mui := New()
			if err := mui.Register(tt.args...); (err != nil) != tt.wantErr {
				t.Errorf("MemUI.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.checks != nil {
				tt.checks(t, mui)
			}
		})
	}
}
