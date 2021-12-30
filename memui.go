package memui

import (
	"fmt"
	"reflect"
)

type MemUI struct {
	memObjects map[string][]interface{}
}

func New() *MemUI {
	return &MemUI{
		memObjects: make(map[string][]interface{}),
	}
}

func (mui *MemUI) addValue(typeName string, obj interface{}) error {
	if _, ok := mui.memObjects[typeName]; !ok {
		mui.memObjects[typeName] = make([]interface{}, 0)
	}

	mui.memObjects[typeName] = append(mui.memObjects[typeName], obj)
	return nil
}

func (mui *MemUI) registerValue(obj interface{}) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("object must be a pointer")
	}

	typeName := v.Type().String()
	return mui.addValue(typeName, obj)
}

func (mui *MemUI) Register(objs ...interface{}) error {
	for _, obj := range objs {
		err := mui.registerValue(obj)
		if err != nil {
			return err
		}
	}

	return nil
}

func (mui *MemUI) ListValues(typeName string) []interface{} {
	if _, ok := mui.memObjects[typeName]; !ok {
		return []interface{}{}
	}

	return mui.memObjects[typeName]
}
