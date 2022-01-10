package memui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Dummy1 struct {
	Name string
}

type Dummy2 struct {
	Count int
}

func TestMemUI_Register(t *testing.T) {
	type checks func(t *testing.T, mui *MemUI)
	tests := []struct {
		name    string
		args    []interface{}
		wantErr bool
		checks  checks
	}{
		{
			"add a pointer",
			[]interface{}{&Dummy1{"Test"}},
			false,
			func(t *testing.T, mui *MemUI) {
				assert.Len(t, mui.memObjects, 1)
				assert.Contains(t, mui.memObjects, "*memui.Dummy1")
			},
		},
		{
			"add a value",
			[]interface{}{Dummy1{"Test"}},
			true,
			func(t *testing.T, mui *MemUI) {
				assert.Len(t, mui.memObjects, 0)
			},
		},
		{
			"add multiple objects",
			[]interface{}{&Dummy1{"dummy2"}, &Dummy2{2}, Dummy2{2}},
			true,
			func(t *testing.T, mui *MemUI) {
				assert.Len(t, mui.memObjects, 2)
				assert.Len(t, mui.memObjects["*memui.Dummy1"], 1)
				assert.Len(t, mui.memObjects["*memui.Dummy2"], 1)
			},
		},
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
