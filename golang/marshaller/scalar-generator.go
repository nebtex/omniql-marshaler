// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates scalar-decoder.gen.go It can be invoked by running
// go generate

package main

import "os"
import "text/template"
import "log"

import (
	"io"
	"github.com/nebtex/hybrids/golang/hybrids"
)

var Int8ToFloat32Template *template.Template
var VectorScalarTemplate *template.Template

func init() {
	var err error
	Int8ToFloat32Template, err = template.New("Int8ToFloat32TableTemplate").Parse(`

func (d *Decoder) get{{.scalar.String}}(number interface{})(value {{.scalar.NativeType}}, err error){
    var fv float64
    var ok bool

    value, ok = number.({{.scalar.NativeType}})
	if ok {
		return
	}

    fv, ok = number.(float64)

    if !ok{
        err = fmt.Errorf("Number expected (float64 or {{.scalar.NativeType}}), got %s", reflect.ValueOf(number).Type().String())
        return
    }

    if !(hybrids.Min{{.scalar.String}} <= fv && fv <= hybrids.Max{{.scalar.String}}){
        err = fmt.Errorf("Number is out of bound, got %d", fv)
		return
	}

	value = {{.scalar.NativeType}}(fv)
	return
}

func (d *Decoder) decode{{.scalar.String}}(path string, number interface{}, fn hybrids.FieldNumber, tw hybrids.{{.scalar.String}}WriterAccessor) (err error) {
	var value {{.scalar.NativeType}}

	value, err = d.get{{.scalar.String}}(number)

	if err!=nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  "{{.scalar.String}}",
			OmniqlType:  "{{.scalar.String}}",
			ErrorMsg:    err.Error(),
		}
		return
	}


	err = tw.Set{{.scalar.String}}(fn, value)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  "{{.scalar.String}}",
			OmniqlType:  "{{.scalar.String}}",
			ErrorMsg:    err.Error(),
		}
		return
	}

	return
}
`)
	die(err)

	VectorScalarTemplate, err = template.New("VectorScalarTemplate").Parse(`
func (d *Decoder) decodeVector{{.scalar.String}}(path string, value interface{}, fn hybrids.FieldNumber, tw hybrids.Vector{{.scalar.String}}WriterAccessor) (err error) {
	var vector hybrids.Vector{{.scalar.String}}Writer
	var item {{.scalar.NativeType}}
	var vi []interface{}
	var ok bool
	vi, ok = value.([]interface{})

    if !ok{
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  "Vector{{.scalar.String}}",
			OmniqlType:  "Vector",
			OmniqlItems:  "{{.scalar.String}}",
			ErrorMsg:    fmt.Sprintf("vector [] expected, got %s", reflect.ValueOf(value).Type().String()),
		}
		return
	}
	vector, err = tw.UpsertVector{{.scalar.String}}(fn)

	if err != nil {
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  "Vector{{.scalar.String}}",
			OmniqlType:  "Vector",
			OmniqlItems:  "{{.scalar.String}}",
			ErrorMsg:    err.Error(),
		}
		return
	}

	for _, v := range vi {
		item, err = d.get{{.scalar.String}}(v)
		if err!=nil{
		    err = &DecodeError{
			   Path:        path,
			   Application: d.application,
			   HybridType:  "{{.scalar.String}}",
			   OmniqlType:  "{{.scalar.String}}",
			   ErrorMsg:    err.Error(),
		    }
		    return
		}
		err = vector.Insert{{.scalar.String}}(-1, item)
		if err!=nil{
		    err = &DecodeError{
			   Path:        path,
			   Application: d.application,
			   HybridType:  "{{.scalar.String}}",
			   OmniqlType:  "{{.scalar.String}}",
			   ErrorMsg:    err.Error(),
		    }
		    return
		}
	}
	return
}
`)
	die(err)

}

func RenderScalarDecoders(f io.Writer) {
	var err error
	scalars := []hybrids.BasicTypes{
		hybrids.Int8,
		hybrids.Uint8,
		hybrids.Int16,
		hybrids.Uint16,
		hybrids.Int32,
		hybrids.Uint32,
		hybrids.Float32,
	}
	for _, v := range scalars {
		err = Int8ToFloat32Template.Execute(f, map[string]interface{}{
			"scalar": v,
		})
		die(err)
	}

}

func RenderVectorScalarDecoders(f io.Writer) {
	var err error
	scalars := []hybrids.BasicTypes{
		hybrids.Int8,
		hybrids.Uint8,
		hybrids.Int16,
		hybrids.Uint16,
		hybrids.Int32,
		hybrids.Uint32,
		hybrids.Float32,
		hybrids.Float64,
		hybrids.Int64,
		hybrids.Uint64,
	}
	for _, v := range scalars {
		err = VectorScalarTemplate.Execute(f, map[string]interface{}{
			"scalar": v,
		})
		die(err)
	}

}
func main() {
	f, err := os.Create("scalar-decoder.gen.go")
	die(err)
	defer f.Close()
	_, err = f.Write([]byte("// Code generated for marshaling/unmarshaling. DO NOT EDIT.\n"))
	die(err)
	_, err = f.Write([]byte(`package marshaller
`))
	die(err)
	_, err = f.Write([]byte(`
import (
	"fmt"
	"github.com/nebtex/hybrids/golang/hybrids"
	"reflect"
)
`))
	die(err)
	RenderScalarDecoders(f)
	RenderVectorScalarDecoders(f)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
