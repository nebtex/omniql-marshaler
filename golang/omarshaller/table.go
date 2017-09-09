package omarshaller

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
	"github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
	"reflect"
)

func (d *Decoder) decodeTable(path string, value interface{}, oType *oreflection.OType, tw hybrids.TableWriter) (err error) {
	var vi map[string]interface{}
	var ok bool
	var fw *oreflection.FieldWrapper
	var fo *oreflection.OType

	if value == nil {
		//Don't return an error remember that tables  can be null in a parent table
		return
	}

	vi, ok = value.(map[string]interface{})

	if !ok {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  oType.Id,
			HybridType:  hybrids.Table,
			ErrorMsg:    fmt.Sprintf("I expected a table, got %s", reflect.ValueOf(value).Type().String()),
		}
		return
	}

	for field_name, value := range vi {
		//try to match Camelcase (default)
		fw, ok = oType.FieldMap.Camel[field_name]
		if !ok {
			//try to match snake_case
			fw, ok = oType.FieldMap.Snake[field_name]
			if ok {
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
				HybridType:  hybrids.Table,
				ErrorMsg:    err.Error(),
			}
			return
		}

		if fo.HybridType.IsScalar() {
			err = d.decodeScalar(path+"."+field_name, value, fo.HybridType, fw.Position, tw)
			if err != nil {
				return
			}
			continue
		} else if fo.HybridType.IsVectorScalar() {
			err = d.decodeVectorScalar(path+"."+field_name, value, fo.HybridType, fw.Position, tw)
			if err != nil {
				return
			}
			continue
		}

		switch fo.HybridType {
		case hybrids.Struct:
			if value == nil {
				//Remember a struct is a superset of scalars, so it should never have a nil value
				err = &DecodeError{
					Path:        path,
					Application: d.application,
					OmniqlType:  fo.Id,
					HybridType:  hybrids.Struct,
					ErrorMsg:    fmt.Sprintf("I expected a Struct, got nil/null"),
				}
			} else {
				str, err := tw.SetStruct(fw.Position)
				if err != nil {
					err = &DecodeError{
						Path:        path,
						Application: d.application,
						OmniqlType:  fo.Id,
						HybridType:  hybrids.Struct,
						ErrorMsg:    err.Error(),
					}
				} else {
					err = d.decodeStruct(path+"."+field_name, value, fo, str)
				}
			}
		case hybrids.String:
			err = d.decodeString(path+"."+field_name, value, fw.Position, tw)
		case hybrids.Byte:
			err = d.decodeByte(path+"."+field_name, value, fw.Position, tw)
		case hybrids.ResourceID:
			err = d.decodeResourceID(path+"."+field_name, value, fw.Position, fo, tw)
		case hybrids.VectorStruct:
			vw, err := tw.SetVectorStruct(fw.Position)
			if err != nil {
				err = &DecodeError{
					Path:        path,
					Application: d.application,
					OmniqlType:  fo.Id,
					HybridType:  hybrids.VectorStruct,
					ErrorMsg:    err.Error(),
				}
			} else {
				err = d.decodeVectorStruct(path+"."+field_name, value, fo.Items, vw)
			}
		case hybrids.VectorString:
			vw, err := tw.SetVectorString(fw.Position)
			if err != nil {
				err = &DecodeError{
					Path:        path,
					Application: d.application,
					OmniqlType:  "Vector[String]",
					HybridType:  hybrids.VectorString,
					ErrorMsg:    err.Error(),
				}
			} else {
				err = d.decodeVectorString(path+"."+field_name, value, vw)

			}
		case hybrids.VectorByte:
			vw, err := tw.SetVectorByte(fw.Position)
			if err != nil {
				err = &DecodeError{
					Path:        path,
					Application: d.application,
					OmniqlType:  "Vector[Byte]",
					HybridType:  hybrids.VectorByte,
					ErrorMsg:    err.Error(),
				}
			} else {
				err = d.decodeVectorByte(path+"."+field_name, value, vw)
			}
		case hybrids.VectorResourceID:
			vw, err := tw.SetVectorResourceID(fw.Position)
			if err != nil {
				err = &DecodeError{
					Path:        path,
					Application: d.application,
					OmniqlType:  fo.Id,
					HybridType:  hybrids.VectorResourceID,
					ErrorMsg:    err.Error(),
				}
			} else {
				err = d.decodeVectorResourceID(path+"."+field_name, value, fo, vw)
			}
		case hybrids.Table:
			tw, err := tw.SetTable(fw.Position)
			if err != nil {
				err = &DecodeError{
					Path:        path,
					Application: d.application,
					OmniqlType:  fo.Id,
					HybridType:  hybrids.Table,
					ErrorMsg:    err.Error(),
				}
			} else {
				err = d.decodeTable(path+"."+field_name, value, fo, tw)
			}
		case hybrids.VectorTable:
			vtw, err := tw.SetVectorTable(fw.Position)
			if err != nil {
				err = &DecodeError{
					Path:        path,
					Application: d.application,
					OmniqlType:  fo.Id,
					HybridType:  hybrids.VectorTable,
					ErrorMsg:    err.Error(),
				}
			} else {
				err = d.decodeVectorTable(path+"."+field_name, value, fo, vtw)
			}
		}

		if err != nil {
			return
		}

	}
	return
}

func (d *Decoder) decodeVectorTable(path string, value interface{}, oType *oreflection.OType, vtw hybrids.VectorTableWriter) (err error) {
	var vi []interface{}
	var ok bool
	var tableWriter hybrids.TableWriter

	if value == nil {
		//Don't return an error remember that a table can be null in a  parent table
		return
	}
	vi, ok = value.([]interface{})

	if !ok {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  oType.Id,
			HybridType:  hybrids.VectorTable,
			ErrorMsg:    fmt.Sprintf("I expected a vector, got %s", reflect.ValueOf(value).Type().String()),
		}
		return
	}

	for index, item := range vi {
		tableWriter, err = vtw.PushTable()
		if err != nil {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				OmniqlType:  oType.Id,
				HybridType:  hybrids.VectorTable,
				ErrorMsg:    err.Error(),
			}
		}
		err = d.decodeTable(path+fmt.Sprintf("[%d]", index), item, oType.Items, tableWriter)
		if err != nil {
			return err
		}
	}
	return
}
