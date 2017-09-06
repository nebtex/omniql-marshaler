package omarshaller

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"go.uber.org/zap"
	"fmt"
	"strconv"
	"reflect"
	"github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
)

//go:generate go run scalar-generator.go

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

func (d *Decoder) decodeVectorScalar(path string, value interface{}, fieldType hybrids.Types, fn hybrids.FieldNumber, tw hybrids.TableWriter) (err error) {
	switch  fieldType {
	case hybrids.VectorBoolean:
		sw, err := tw.SetVectorBoolean(fn)
		if err != nil {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				OmniqlType:  "Vector[Boolean]",
				HybridType:  hybrids.VectorBoolean,
				ErrorMsg:    err.Error(),
			}
		}
		err = d.decodeVectorBoolean(path, value, fn, sw)
	case hybrids.VectorInt8:
		sw, err := tw.SetVectorInt8(fn)
		if err != nil {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				OmniqlType:  "Vector[Int8]",
				HybridType:  hybrids.VectorInt8,
				ErrorMsg:    err.Error(),
			}
		}
		err = d.decodeVectorInt8(path, value, fn, sw)
	case hybrids.VectorUint8:
		sw, err := tw.SetVectorUint8(fn)
		if err != nil {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				OmniqlType:  "Vector[Uint8]",
				HybridType:  hybrids.VectorUint8,
				ErrorMsg:    err.Error(),
			}
		}
		err = d.decodeVectorUint8(path, value, fn, sw)
	case hybrids.VectorInt16:
		sw, err := tw.SetVectorInt16(fn)
		if err != nil {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				OmniqlType:  "Vector[Int16]",
				HybridType:  hybrids.VectorInt16,
				ErrorMsg:    err.Error(),
			}
		}
		err = d.decodeVectorInt16(path, value, fn, sw)
	case hybrids.VectorUint16:
		sw, err := tw.SetVectorUint16(fn)
		if err != nil {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				OmniqlType:  "Vector[Uint16]",
				HybridType:  hybrids.VectorUint16,
				ErrorMsg:    err.Error(),
			}
		}
		err = d.decodeVectorUint16(path, value, fn, sw)
	case hybrids.VectorInt32:
		sw, err := tw.SetVectorInt32(fn)
		if err != nil {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				OmniqlType:  "Vector[Int32]",
				HybridType:  hybrids.VectorInt32,
				ErrorMsg:    err.Error(),
			}
		}
		err = d.decodeVectorInt32(path, value, fn, sw)
	case hybrids.VectorUint32:
		sw, err := tw.SetVectorUint32(fn)
		if err != nil {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				OmniqlType:  "Vector[Uint32]",
				HybridType:  hybrids.VectorUint32,
				ErrorMsg:    err.Error(),
			}
		}
		err = d.decodeVectorUint32(path, value, fn, sw)
	case hybrids.VectorInt64:
		sw, err := tw.SetVectorInt64(fn)
		if err != nil {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				OmniqlType:  "Vector[Int64]",
				HybridType:  hybrids.VectorInt64,
				ErrorMsg:    err.Error(),
			}
		}
		err = d.decodeVectorInt64(path, value, fn, sw)
	case hybrids.VectorUint64:
		sw, err := tw.SetVectorUint64(fn)
		if err != nil {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				OmniqlType:  "Vector[Uint64]",
				HybridType:  hybrids.VectorUint64,
				ErrorMsg:    err.Error(),
			}
		}
		err = d.decodeVectorUint64(path, value, fn, sw)
	case hybrids.VectorFloat32:
		sw, err := tw.SetVectorFloat32(fn)
		if err != nil {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				OmniqlType:  "Vector[Float32]",
				HybridType:  hybrids.VectorFloat32,
				ErrorMsg:    err.Error(),
			}
		}
		err = d.decodeVectorFloat32(path, value, fn, sw)
	case hybrids.VectorFloat64:
		sw, err := tw.SetVectorFloat64(fn)
		if err != nil {
			err = &DecodeError{
				Path:        path,
				Application: d.application,
				OmniqlType:  "Vector[Float64]",
				HybridType:  hybrids.VectorFloat64,
				ErrorMsg:    err.Error(),
			}
		}
		err = d.decodeVectorFloat64(path, value, fn, sw)
	default:
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  "VectorScalar",
			ErrorMsg:    fmt.Sprintf("%s not recognized as VectorScalar, this error should not happend :(", fieldType),
		}
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

func (d *Decoder) decodeVectorBoolean(path string, value interface{}, fn hybrids.FieldNumber, vw hybrids.VectorBooleanWriter) (err error) {
	var item bool
	var vi []interface{}
	var ok bool
	if value == nil {
		//Don't return an error remember that vector  can be null in a table
		return
	}

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
		err = vw.PushBoolean(item)
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
	}
	return
}
