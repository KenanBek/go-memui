package memui

import "reflect"

type MemUI struct {
	memObjects map[string][]interface{}
}

func New() *MemUI {
	return &MemUI{
		memObjects: make(map[string][]interface{}),
	}
}

func (mui *MemUI) Register(objs ...interface{}) error {
	for _, obj := range objs {
		typeName := reflect.TypeOf(obj).String()

		if _, ok := mui.memObjects[typeName]; !ok {
			mui.memObjects[typeName] = make([]interface{}, 0)
		}

		mui.memObjects[typeName] = append(mui.memObjects[typeName], obj)
	}

	return nil
}

func (mui *MemUI) ListValues(typeName string) []interface{} {
	if _, ok := mui.memObjects[typeName]; !ok {
		return []interface{}{}
	}

	return mui.memObjects[typeName]
}
