package omarshaller

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"go.uber.org/zap"
	"fmt"
	"strconv"
	"reflect"
)

type Decoder struct {
	application string
	//reflection  hybrids.SimpleStore
	zap *zap.Logger
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

func (d *Decoder) decodeScalar(path string, number float64, tw hybrids.TableWriter) (err error) {
	return
}

//go:generate go run scalar-generator.go

func (d *Decoder) getFloat64(number interface{}) (value float64, err error) {
	var ok bool
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
			HybridType:  "Float64",
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
			HybridType:  "Float64",
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
			HybridType:  "Int64",
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
			HybridType:  "Int64",
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
			HybridType:  "Uint64",
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
			HybridType:  "Uint64",
			OmniqlType:  "Uint64",
			ErrorMsg:    err.Error(),
		}
		return
	}

	return
}
