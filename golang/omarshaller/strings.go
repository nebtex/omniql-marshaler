package omarshaller

import (
	b64 "encoding/base64"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
	"reflect"
	"github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
)

func (d *Decoder) getString(value interface{}) (str string, err error) {
	str, _ = value.(string)
	return
}

func (d *Decoder) getByte(value interface{}) (b []byte, err error) {
	var ok bool
	var str string
	b, ok = value.([]byte)

	if ok {
		return
	}

	str, ok = value.(string)
	if !ok {
		return
	}

	b, err = b64.StdEncoding.DecodeString(str)
	return
}

func (d *Decoder) decodeString(path string, value interface{}, fn hybrids.FieldNumber, tw hybrids.StringWriterAccessor) (err error) {
	var str string

	if value == nil {
		//Don't return an error remember that vectors (string is a vector of bytes) can be null in a table
		return
	}

	str, err = d.getString(value)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  "String",
			HybridType:  hybrids.String,
			ErrorMsg:    err.Error(),
		}
		return
	}

	err = tw.SetString(fn, str)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  "String",
			HybridType:  hybrids.String,
			ErrorMsg:    err.Error(),
		}
		return
	}

	return
}

func (d *Decoder) decodeByte(path string, value interface{}, fn hybrids.FieldNumber, tw hybrids.ByteWriterAccessor) (err error) {
	var str []byte

	if value == nil {
		//Don't return an error remember that vectors (string is a vector of bytes) can be null in a table
		return
	}

	str, err = d.getByte(value)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  "Byte",
			HybridType:  hybrids.Byte,
			ErrorMsg:    err.Error(),
		}
		return
	}

	err = tw.SetByte(fn, str)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  "Byte",
			HybridType:  hybrids.Byte,
			ErrorMsg:    err.Error(),
		}
		return
	}

	return
}

func (d *Decoder) decodeResourceID(path string, value interface{}, fn hybrids.FieldNumber, ot *oreflection.OType, tw hybrids.ResourceIDWriterAccessor) (err error) {
	var str []byte

	if value == nil {
		//Don't return an error remember that vectors (string is a vector of bytes) can be null in a table
		return
	}

	str, err = d.getByte(value)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  ot.Id,
			HybridType:  hybrids.ResourceID,
			ErrorMsg:    err.Error(),
		}
		return
	}

	err = tw.SetResourceID(fn, d.ResourceKindType, d.ResourceIDType, str)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  ot.Id,
			HybridType:  hybrids.ResourceID,
			ErrorMsg:    err.Error(),
		}
		return
	}

	return
}

func (d *Decoder) decodeVectorString(path string, value interface{}, tw hybrids.VectorStringWriter) (err error) {
	var vi []interface{}
	var ok bool
	var str string
	var item_path string

	if value == nil {
		//Don't return an error remember that vectors  can be null in a table
		return
	}

	vi, ok = value.([]interface{})

	if !ok {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  "Vector[String]",
			HybridType:  hybrids.VectorString,
			ErrorMsg:    fmt.Sprintf("I expected a vector, got %s", reflect.ValueOf(value).Type().String()),
		}
		return
	}

	for index, item := range vi {
		item_path = path + fmt.Sprintf("[%s]", index)

		str, err = d.getString(item)

		if err != nil {
			err = &DecodeError{
				Path:        item_path,
				Application: d.application,
				OmniqlType:  "Vector[String]",
				HybridType:  hybrids.VectorString,
				ErrorMsg:    err.Error(),
			}
		}

		err = tw.PushString(str)

		if err != nil {
			err = &DecodeError{
				Path:        item_path,
				Application: d.application,
				OmniqlType:  "Vector[String]",
				HybridType:  hybrids.VectorString,
				ErrorMsg:    err.Error(),
			}
			return
		}
	}

	return
}

func (d *Decoder) decodeVectorByte(path string, value interface{}, tw hybrids.VectorByteWriter) (err error) {
	var vi []interface{}
	var ok bool
	var str []byte
	var item_path string

	if value == nil {
		//Don't return an error remember that vectors  can be null in a table
		return
	}

	vi, ok = value.([]interface{})

	if !ok {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  "Vector[Byte]",
			HybridType:  hybrids.VectorByte,
			ErrorMsg:    fmt.Sprintf("I expected a vector, got %s", reflect.ValueOf(value).Type().String()),
		}
		return
	}

	for index, item := range vi {
		item_path = path + fmt.Sprintf("[%s]", index)

		str, err = d.getByte(item)

		if err != nil {
			err = &DecodeError{
				Path:        item_path,
				Application: d.application,
				OmniqlType:  "Vector[Byte]",
				HybridType:  hybrids.VectorByte,
				ErrorMsg:    err.Error(),
			}
		}

		err = tw.PushByte(str)

		if err != nil {
			err = &DecodeError{
				Path:        item_path,
				Application: d.application,
				OmniqlType:  "Vector[Byte]",
				HybridType:  hybrids.VectorByte,
				ErrorMsg:    err.Error(),
			}
			return
		}
	}

	return
}

func (d *Decoder) decodeVectorResourceID(path string, value interface{}, ot *oreflection.OType, tw hybrids.VectorResourceIDWriter) (err error) {
	var vi []interface{}
	var ok bool
	var str []byte
	var item_path string

	if value == nil {
		//Don't return an error remember that vectors  can be null in a table
		return
	}

	vi, ok = value.([]interface{})

	if !ok {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			OmniqlType:  ot.Id,
			HybridType:  hybrids.VectorResourceID,
			ErrorMsg:    fmt.Sprintf("I expected a vector, got %s", reflect.ValueOf(value).Type().String()),
		}
		return
	}

	for index, item := range vi {
		item_path = path + fmt.Sprintf("[%s]", index)

		str, err = d.getByte(item)

		if err != nil {
			err = &DecodeError{
				Path:        item_path,
				Application: d.application,
				OmniqlType:  ot.Items.Id,
				HybridType:  hybrids.VectorResourceID,
				ErrorMsg:    err.Error(),
			}
		}

		err = tw.PushResourceID(d.ResourceKindType, d.ResourceIDType, str)

		if err != nil {
			err = &DecodeError{
				Path:        item_path,
				Application: d.application,
				OmniqlType:  ot.Items.Id,
				HybridType:  hybrids.VectorResourceID,
				ErrorMsg:    err.Error(),
			}
			return
		}
	}

	return
}
