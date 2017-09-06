package omarshaller

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
	"github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
	"reflect"
)

func (d *Decoder) decodeStruct(path string, value interface{}, oType *oreflection.OType, sw hybrids.ScalarWriter) (err error) {
	var vi map[string]interface{}
	var ok bool
	var fw *oreflection.FieldWrapper
	var fo *oreflection.OType

	vi, ok = value.(map[string]interface{})

	if !ok {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  oType.Id,
			HybridType:  hybrids.Struct,
			ErrorMsg:    fmt.Sprintf("I expected a struct, got %s", reflect.ValueOf(value).Type().String()),
		}
		return
	}

	for field_name, value := range vi {
		//try to match Camelcase (default)
		fw, ok = oType.FieldMap.Camel[field_name]
		if !ok {
			//try to match snake_case
			fw, ok = oType.FieldMap.Snake[field_name]
			if !ok {
				//ignore field
				continue
			}
		}

		fo, err = d.reflect.OReflect(fw.Field.Type())
		if err != nil {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				OmniqlType:  oType.Id,
				HybridType:  hybrids.Struct,
				ErrorMsg:    err.Error(),
			}
			return
		}

		err = d.decodeScalar(path+"."+field_name, value, fo.HybridType, fw.Position, sw)
		if err != nil {
			return
		}
	}
	return
}

func (d *Decoder) decodeVectorStruct(path string, value interface{}, oType *oreflection.OType, tw hybrids.VectorStructWriter) (err error) {
	var vi []interface{}
	var ok bool
	var scalarWriter hybrids.ScalarWriter

	if value == nil {
		//Don't return an error remember that vector struct can be null in a table
		return
	}
	vi, ok = value.([]interface{})

	if !ok {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  oType.Id,
			HybridType:  hybrids.VectorStruct,
			ErrorMsg:    fmt.Sprintf("I expected a vector, got %s", reflect.ValueOf(value).Type().String()),
		}
		return
	}

	for index, item := range vi {
		scalarWriter, err = tw.PushStruct()
		if err != nil {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				OmniqlType:  oType.Id,
				HybridType:  hybrids.VectorStruct,
				ErrorMsg:    err.Error(),
			}
		}
		err = d.decodeStruct(path+fmt.Sprintf("[%d]", index), item, oType.Items, scalarWriter)
		if err != nil {
			return err
		}

	}
	return
}
