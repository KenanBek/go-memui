package memui

import "testing"

type Dummy1 struct {
	Name string
}

func TestMemUI_Register(t *testing.T) {
	type fields struct {
		memObjects map[string][]interface{}
	}
	type args struct {
		objs []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "add pointer object",
			fields:  fields{memObjects: make(map[string][]interface{})},
			args:    args{objs: []interface{}{&Dummy1{Name: "dummy1"}}},
			wantErr: false,
		},
		{
			name:    "add value object",
			fields:  fields{memObjects: make(map[string][]interface{})},
			args:    args{objs: []interface{}{Dummy1{Name: "dummy1"}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mui := &MemUI{
				memObjects: tt.fields.memObjects,
			}
			if err := mui.Register(tt.args.objs...); (err != nil) != tt.wantErr {
				t.Errorf("MemUI.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
