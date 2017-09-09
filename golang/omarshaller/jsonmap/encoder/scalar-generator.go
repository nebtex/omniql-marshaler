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

func RenderScalarEncoders(f io.Writer) {
	var err error
	var tmp *template.Template

	tmp, err = template.New("ScalarEncoders").Parse(`

func (e *Encoder) encode{{.scalar.String}}(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.{{.scalar.String}}ReaderAccessor) {
	v, ok := tr.{{.scalar.String}}(fn)
	if !ok {
		//field is missing
		return
	}
	out[field_name] = {{.jsTransform}}
}


func (e *Encoder) encodeVector{{.scalar.String}}(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.Vector{{.scalar.String}}ReaderAccessor) {
	var v {{.scalar.NativeType}}
	vrb, ok := tr.Vector{{.scalar.String}}(fn)

	if !ok {
		return
	}

	if vrb == nil {
		out[field_name] = nil
		return
	}

	r := make([]interface{}, 0, vrb.Len())

	for i := 0; i < vrb.Len(); i++ {
		v = vrb.Get(i)
		r = append(r, {{.jsTransform}})
	}

	out[field_name] = r

	return
}
`)
	die(err)

	scalars := []hybrids.Types{
		hybrids.Boolean,
		hybrids.Int8,
		hybrids.Uint8,
		hybrids.Int16,
		hybrids.Uint16,
		hybrids.Int32,
		hybrids.Uint32,
		hybrids.Int64,
		hybrids.Uint64,
		hybrids.Float32,
		hybrids.Float64,
	}
	jsonFunction := []string{
		"v",
		"float64(v)",
		"float64(v)",
		"float64(v)",
		"float64(v)",
		"float64(v)",
		"float64(v)",
		"strconv.FormatInt(v, 10)",
		"strconv.FormatUint(v, 10)",
		"float64(v)",
		"v",
	}

	jsonType := []string{
		"bool",
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"string",
		"string",
		"float64",
		"float64",
	}
	for index, v := range scalars {
		err = tmp.Execute(f, map[string]interface{}{
			"scalar":     v,
			"jsTransform": jsonFunction[index],
			"jsType": jsonType[index],

		})
		die(err)
	}

	decodeScalarTemplate, err := template.New("decodeScalarTemplate").Parse(`

func (e *Encoder) encodeScalar(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tp hybrids.Types, tr hybrids.ScalarReader){
	switch tp{
{{ range $index, $value := .scalars }}
	case hybrids.{{$value.String}}:
		e.encode{{$value.String}}(out, field_name, fn, tr)
{{ end }}
    }
	return
}
`)
	die(err)

	err = decodeScalarTemplate.Execute(f, map[string]interface{}{
		"scalars":     scalars,
	})
	die(err)



}

func RenderTestScalarDecoders(f io.Writer) {
	var err error
	var Int8ToFloat32_Test_Template *template.Template
	Int8ToFloat32_Test_Template, err = template.New("Int8ToFloat32_Test_Template").Parse(`

func Test_Decoder{{.scalar.String}}(t *testing.T) {
	Convey("{{.scalar.String}} Tests", t, func() {
		table := []struct {
			fn             hybrids.FieldNumber
			number         interface{}
			mockNumber     {{.scalar.NativeType}}
			shouldFail     bool
			shouldMockFail bool
			name           string
		}{
			{30, float64(35), 35, false, false, "underlying type: float64"},
			{1, {{.scalar.NativeType}}(25), 25, false, false, "native_type"},
			{25, "2656", 0, true, false, "incorrect underlying type"},
			{70, float64(hybrids.Max{{.scalar.String}}) + 1e39, 0, true, false, "max value"},
			{50, float64(hybrids.Min{{.scalar.String}}) - 1e39, 0, true, false, "min value"},
			{10, {{.scalar.NativeType}}(44), 44, true, true, "hybrid set method fail"},
			{10, nil, 0, true, false, "null should return error in any scalar"},


		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				{{.scalar.NativeType}}Mock := &mocks.{{.scalar.String}}WriterAccessor{}

				call := {{.scalar.NativeType}}Mock.On("Set{{.scalar.String}}", ti.fn, ti.mockNumber)
				if ti.shouldMockFail {
					call.Return(fmt.Errorf("Set{{.scalar.String}} failed"))
				} else {
					call.Return(nil)
				}

				err := d.decode{{.scalar.String}}("x.path", ti.number, ti.fn, {{.scalar.NativeType}}Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _:= err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.Path, ShouldEqual, "x.path")
					So(de.HybridType, ShouldEqual, hybrids.{{.scalar.String}})
					So(de.OmniqlType, ShouldEqual, "{{.scalar.String}}")

				} else {
					So(err, ShouldBeNil)
					if !ti.shouldFail {
						{{.scalar.NativeType}}Mock.AssertCalled(t, "Set{{.scalar.String}}", ti.fn, ti.mockNumber)
					}
				}


			})
		}

	})

}
`)
	die(err)

	scalars := []hybrids.Types{
		hybrids.Int8,
		hybrids.Uint8,
		hybrids.Int16,
		hybrids.Uint16,
		hybrids.Int32,
		hybrids.Uint32,
		hybrids.Float32,
	}
	for _, v := range scalars {
		err = Int8ToFloat32_Test_Template.Execute(f, map[string]interface{}{
			"scalar": v,
		})
		die(err)
	}

}
func RenderVectorScalarDecoders(f io.Writer) {

	var err error
	var VectorScalarTemplate *template.Template

	VectorScalarTemplate, err = template.New("VectorScalarTemplate").Parse(`
func (d *Decoder) decodeVector{{.scalar.String}}(path string, value interface{}, vw hybrids.Vector{{.scalar.String}}Writer) (err error) {
    var item {{.scalar.NativeType}}
    var vi []interface{}
    var ok bool

    if value == nil {
		//Don't return an error remember that vector  can be null in a table
		return
	}

	vi, ok = value.([]interface{})

	if !ok{
		err = &DecodeError{
			Path:        path,
			Application: d.application,
			HybridType:  hybrids.Vector{{.scalar.String}},
			OmniqlType:  "Vector[{{.scalar.String}}]",
			ErrorMsg:    fmt.Sprintf("vector [] expected, got %s", reflect.ValueOf(value).Type().String()),
        }
     return
    }


    for index, v := range vi {
        item, err = d.get{{.scalar.String}}(v)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:  hybrids.Vector{{.scalar.String}},
               OmniqlType:  "Vector[{{.scalar.String}}]",
               ErrorMsg:    err.Error(),
            }
            return
        }
        err = vw.Push{{.scalar.String}}(item)
        if err!=nil{
            err = &DecodeError{
               Path:        fmt.Sprintf("%s[%d]", path, index),
               Application: d.application,
               HybridType:   hybrids.Vector{{.scalar.String}},
               OmniqlType:  "Vector[{{.scalar.String}}]",
               ErrorMsg:    err.Error(),
            }
            return
        }
    }
    return
}
`)

	die(err)
	scalars := []hybrids.Types{
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

func RenderVectorScalar_TEST_Decoders(f io.Writer) {

	var err error
	var VectorScalar_TEST_Template *template.Template

	VectorScalar_TEST_Template, err = template.New("VectorScalar_TEST_Template").Parse(`
func Test_Decode_Vector{{.scalar.String}}(t *testing.T) {
	Convey("Decode Vector{{.scalar.String}} test", t, func() {
		table := []struct {
			fn                      hybrids.FieldNumber
			vector                  interface{}
			mockVector              []{{.scalar.NativeType}}
			shouldFail              bool
			makePushFail            bool
			name                    string
		}{
			{30, nil, nil, false, false, "null: should create a vector but don't add items (remember vector and tables can have a null state)"},
			{25, "vector", nil, true, false, "incorrect underlying type"},
			{70, []interface{}{"nan"}, []{{.scalar.NativeType}}{0}, true, false, "item: incorrect underlying type"},
			{50, []interface{}{ {{.scalar.NativeType}}(1), {{.scalar.NativeType}}(2), {{.scalar.NativeType}}(3)}, []{{.scalar.NativeType}}{1, 2, 3}, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{ {{.scalar.NativeType}}(1), {{.scalar.NativeType}}(2), {{.scalar.NativeType}}(3)}, []{{.scalar.NativeType}}{1, 2, 3}, false, false, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				{{.scalar.NativeType}}Mock := &mocks.Vector{{.scalar.String}}Writer{}
				for _, item := range ti.mockVector {
					callItem := {{.scalar.NativeType}}Mock.On("Push{{.scalar.String}}", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("Push{{.scalar.String}} failed"))
					} else {
						callItem.Return(nil)
					}
				}

				err := d.decodeVector{{.scalar.String}}("x.path.vector", ti.vector, {{.scalar.NativeType}}Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.Vector{{.scalar.String}})
					So(de.OmniqlType, ShouldEqual, "Vector[{{.scalar.String}}]")

				} else {
					So(err, ShouldBeNil)
						if !ti.makePushFail {
							for _, item := range ti.mockVector {
								{{.scalar.NativeType}}Mock.AssertCalled(t, "Push{{.scalar.String}}", item)
							}
					}
				}
			})
		}

	})

}

`)

	die(err)
	scalars := []hybrids.Types{
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
		err = VectorScalar_TEST_Template.Execute(f, map[string]interface{}{
			"scalar": v,
		})
		die(err)
	}
}

func main() {
	f, err := os.Create("scalar-encoder.gen.go")
	die(err)
	defer f.Close()
	_, err = f.Write([]byte("// Code generated for marshaling/unmarshaling. DO NOT EDIT.\n"))
	die(err)
	_, err = f.Write([]byte(`package encoder
`))
	die(err)
	_, err = f.Write([]byte(`
import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"strconv"
)
`))
	die(err)
	RenderScalarEncoders(f)
	/*RenderVectorScalarDecoders(f)

	test, err := os.Create("scalar-decoder_test.go")
	die(err)
	defer test.Close()
	_, err = test.Write([]byte("// Code generated. DO NOT EDIT.\n"))
	die(err)
	_, err = test.Write([]byte(`package omarshaller
`))
	die(err)
	_, err = test.Write([]byte(`
import (
	"testing"
	"github.com/nebtex/hybrids/golang/hybrids/mocks"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
)
`))
	die(err)
	RenderTestScalarDecoders(test)
	RenderVectorScalar_TEST_Decoders(test)*/
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
