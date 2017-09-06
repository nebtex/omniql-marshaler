package omarshaller

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"go.uber.org/zap"
	"fmt"
	"strconv"
	"reflect"
	"github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
)

type Decoder struct {
	application      string
	ResourceKindType hybrids.ResourceKindType
	ResourceIDType   hybrids.ResourceIDType

	//reflection  hybrids.SimpleStore
	reflect oreflection.ReflectStore
}

/*
	path is the full path of the field ex: Spec.Fields[0].Name
	data is the json underlying struct can be []interface or map[string]interface{}
	TableID
 */
func (d *Decoder) decode(path string, data interface{}, fieldType string, items string) (out []byte, err error) {
	//check if fieldType is table, union or resource

	//if is a vector check if item is a table, union or resource

	switch v := data.(type) {
	case []interface{}:
		fmt.Printf("Twice %v is %v\n", v, 2)
	case map[string]interface{}:

		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		err = fmt.Errorf("only map[string]interface{} or []interface{} can be decoded")
		zap.Error(err)
		return
	}
	return

}

func (d *Decoder) decodeScalar(path string, value interface{}, fieldType hybrids.Types, fn hybrids.FieldNumber, sw hybrids.ScalarWriter) (err error) {
	switch  fieldType {
	case hybrids.Boolean:
		err = d.decodeBoolean(path, value, fn, sw)
	case hybrids.Int8:
		err = d.decodeInt8(path, value, fn, sw)
	case hybrids.Uint8:
		err = d.decodeUint8(path, value, fn, sw)
	case hybrids.Int16:
		err = d.decodeInt16(path, value, fn, sw)
	case hybrids.Uint16:
		err = d.decodeUint16(path, value, fn, sw)
	case hybrids.Int32:
		err = d.decodeInt32(path, value, fn, sw)
	case hybrids.Uint32:
		err = d.decodeUint32(path, value, fn, sw)
	case hybrids.Int64:
		err = d.decodeInt64(path, value, fn, sw)
	case hybrids.Uint64:
		err = d.decodeUint64(path, value, fn, sw)
	case hybrids.Float32:
		err = d.decodeFloat32(path, value, fn, sw)
	case hybrids.Float64:
		err = d.decodeFloat64(path, value, fn, sw)
	default:
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  "Scalar",
			ErrorMsg:    fmt.Sprintf("%s not recognized as scalar, this error should not happend :(", fieldType),
		}
	}

	return
}

//go:generate go run scalar-generator.go

func (d *Decoder) getFloat64(number interface{}) (value float64, err error) {
	var ok bool

	if number == nil {
		err = fmt.Errorf("Number (float64) expected, got null/nil")
		return
	}
	value, ok = number.(float64)

	if !ok {
		err = fmt.Errorf("Number (float64) expected, got %s", reflect.ValueOf(number).Type().String())
		return
	}
	return
}

func (d *Decoder) decodeFloat64(path string, number interface{}, fn hybrids.FieldNumber, tw hybrids.Float64WriterAccessor) (err error) {
	var fv float64

	fv, err = d.getFloat64(number)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.Float64,
			OmniqlType:  "Float64",
			ErrorMsg:    err.Error(),
		}
		return
	}

	err = tw.SetFloat64(fn, fv)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.Float64,
			OmniqlType:  "Float64",
			ErrorMsg:    err.Error(),
		}
		return
	}

	return
}

func (d *Decoder) getInt64(number interface{}) (value int64, err error) {
	var ok bool
	var sv string
	var fv float64

	if number == nil {
		err = fmt.Errorf("Number (float64 or string) expected, got nil/null")
		return
	}

	value, ok = number.(int64)

	if ok {
		return
	}

	fv, ok = number.(float64)

	if !ok {

		sv, ok = number.(string)
		if !ok {
			err = fmt.Errorf("Number (float64 or string) expected, got %s", reflect.ValueOf(number).Type().String())
			return
		}

		value, err = strconv.ParseInt(sv, 10, 64)

		if err != nil {
			err = fmt.Errorf("(failed to convert string to Int64) %s", err.Error())
			return
		}

		return
	}

	value = int64(fv)
	return
}

func (d *Decoder) decodeInt64(path string, number interface{}, fn hybrids.FieldNumber, tw hybrids.Int64WriterAccessor) (err error) {
	var v int64

	v, err = d.getInt64(number)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.Int64,
			OmniqlType:  "Int64",
			ErrorMsg:    err.Error(),
		}
		return
	}

	err = tw.SetInt64(fn, v)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.Int64,
			OmniqlType:  "Int64",
			ErrorMsg:    err.Error(),
		}
		return
	}

	return
}

func (d *Decoder) getUint64(number interface{}) (value uint64, err error) {
	var ok bool
	var sv string
	var fv float64

	if number == nil {
		err = fmt.Errorf("Number (float64 or string) expected, got nil/null")
		return
	}

	value, ok = number.(uint64)

	if ok {
		return
	}

	fv, ok = number.(float64)

	if !ok {

		sv, ok = number.(string)
		if !ok {
			err = fmt.Errorf("Number (float64 or string) expected, got %s", reflect.ValueOf(number).Type().String())
			return
		}

		value, err = strconv.ParseUint(sv, 10, 64)

		if err != nil {
			err = fmt.Errorf("(failed to convert string to Int64) %s", err.Error())
			return
		}

		return
	}

	value = uint64(fv)
	return
}

func (d *Decoder) decodeUint64(path string, number interface{}, fn hybrids.FieldNumber, tw hybrids.Uint64WriterAccessor) (err error) {
	var v uint64

	v, err = d.getUint64(number)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.Uint64,
			OmniqlType:  "Uint64",
			ErrorMsg:    err.Error(),
		}
		return
	}

	err = tw.SetUint64(fn, v)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.Uint64,
			OmniqlType:  "Uint64",
			ErrorMsg:    err.Error(),
		}
		return
	}

	return
}

func (d *Decoder) getBoolean(number interface{}) (value bool, err error) {
	var ok bool

	if number == nil {
		err = fmt.Errorf("Boolean expected, got nil/null")
		return
	}
	value, ok = number.(bool)

	if !ok {
		err = fmt.Errorf("Boolean expected, got %s", reflect.ValueOf(number).Type().String())
		return
	}

	return
}

func (d *Decoder) decodeBoolean(path string, number interface{}, fn hybrids.FieldNumber, tw hybrids.BooleanWriterAccessor) (err error) {
	var v bool

	v, err = d.getBoolean(number)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.Boolean,
			OmniqlType:  "Boolean",
			ErrorMsg:    err.Error(),
		}
		return
	}

	err = tw.SetBoolean(fn, v)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.Boolean,
			OmniqlType:  "Boolean",
			ErrorMsg:    err.Error(),
		}
		return
	}

	return
}

func (d *Decoder) decodeVectorBoolean(path string, value interface{}, fn hybrids.FieldNumber, tw hybrids.VectorBooleanWriterAccessor) (err error) {
	var vector hybrids.VectorBooleanWriter
	var item bool
	var vi []interface{}
	var ok bool

	if value != nil {
		vi, ok = value.([]interface{})

		if !ok {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				HybridType:  hybrids.VectorBoolean,
				OmniqlType:  "Vector[Boolean]",
				ErrorMsg:    fmt.Sprintf("vector [] expected, got %s", reflect.ValueOf(value).Type().String()),
			}
			return
		}
	}

	vector, err = tw.SetVectorBoolean(fn)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.VectorBoolean,
			OmniqlType:  "Vector[Boolean]",
			ErrorMsg:    err.Error(),
		}
		return
	}
	if value == nil {
		return
	}

	for index, v := range vi {
		item, err = d.getBoolean(v)
		if err != nil {
			err = &DecodeError{
				Path:        fmt.Sprintf("%s[%d]", path, index),
				Application: d.application,
				HybridType:  hybrids.VectorBoolean,
				OmniqlType:  "Vector[Boolean]",
				ErrorMsg:    err.Error(),
			}
			return
		}
		err = vector.PushBoolean(item)
		if err != nil {
			err = &DecodeError{
				Path:        fmt.Sprintf("%s[%d]", path, index),
				Application: d.application,
				HybridType:  hybrids.VectorBoolean,
				OmniqlType:  "Vector",
				ErrorMsg:    err.Error(),
			}
			return
		}
	}
	return
}
