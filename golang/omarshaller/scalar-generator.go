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

func RenderScalarDecoders(f io.Writer) {
	var err error
	var Int8ToFloat32Template *template.Template

	Int8ToFloat32Template, err = template.New("Int8ToFloat32Template").Parse(`

func (d *Decoder) get{{.scalar.String}}(number interface{})(value {{.scalar.NativeType}}, err error){
    var fv float64
    var ok bool

	if number==nil{
	    err = fmt.Errorf("Number expected (float64 or {{.scalar.NativeType}}), got nil/null")
		return
	}

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
          HybridType:  hybrids.{{.scalar.String}},
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
          HybridType:  hybrids.{{.scalar.String}},
          OmniqlType:  "{{.scalar.String}}",
          ErrorMsg:    err.Error(),
       }
       return
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
	}
	for _, v := range scalars {
		err = Int8ToFloat32Template.Execute(f, map[string]interface{}{
			"scalar": v,
		})
		die(err)
	}

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
					So(de.HybridType, ShouldEqual, "{{.scalar.String}}")
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
func (d *Decoder) decodeVector{{.scalar.String}}(path string, value interface{}, fn hybrids.FieldNumber, tw hybrids.Vector{{.scalar.String}}WriterAccessor) (err error) {
    var vector hybrids.Vector{{.scalar.String}}Writer
    var item {{.scalar.NativeType}}
    var vi []interface{}
    var ok bool

	if value!=nil{
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
	}

    vector, err = tw.SetVector{{.scalar.String}}(fn)

    if err != nil {
       err = &DecodeError{
           Path:        path,
           Application: d.application,
           HybridType:  hybrids.Vector{{.scalar.String}},
           OmniqlType:  "Vector[{{.scalar.String}}]",
           ErrorMsg:    err.Error(),
       }
       return
    }
    if value==nil{
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
        err = vector.Push{{.scalar.String}}(item)
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
			makeSetVectorFail    bool
			makePushFail            bool
			shouldTryToCreateVector bool
			name                    string
		}{
			{30, nil, nil, false, false, false, true, "null: should create a vector but don't add items (remember vector and tables can have a null state)"},
			{25, "vector", nil, true, false, false, false, "incorrect underlying type"},
			{70, []interface{}{"nan"}, []{{.scalar.NativeType}}{0}, true, false, false, true, "item: incorrect underlying type"},
			{50, []interface{}{1, 2, 3}, []{{.scalar.NativeType}}{1, 2, 3}, true, true, false, true, "Should fails when the vector creation operation fails"},
			{50, []interface{}{ {{.scalar.NativeType}}(1), {{.scalar.NativeType}}(2), {{.scalar.NativeType}}(3)}, []{{.scalar.NativeType}}{1, 2, 3}, true, false, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{ {{.scalar.NativeType}}(1), {{.scalar.NativeType}}(2), {{.scalar.NativeType}}(3)}, []{{.scalar.NativeType}}{1, 2, 3}, false, false, false, true, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				{{.scalar.NativeType}}Mock := &mocks.Vector{{.scalar.String}}WriterBase{}
				for _, item := range ti.mockVector {
					callItem := {{.scalar.NativeType}}Mock.On("Push{{.scalar.String}}", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("Push{{.scalar.String}} failed"))
					} else {
						callItem.Return(nil)
					}
				}

				vector{{.scalar.String}}Mock := &mocks.Vector{{.scalar.String}}WriterAccessor{}

				call := vector{{.scalar.String}}Mock.On("SetVector{{.scalar.String}}", ti.fn)
				if ti.makeSetVectorFail {
					call.Return(nil, fmt.Errorf("SetVector{{.scalar.String}} failed"))
				} else {
					call.Return({{.scalar.NativeType}}Mock, nil)
				}



				err := d.decodeVector{{.scalar.String}}("x.path.vector", ti.vector, ti.fn, vector{{.scalar.String}}Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.Vector{{.scalar.String}})
					So(de.OmniqlType, ShouldEqual, "Vector[{{.scalar.String}}]")

				} else {
					So(err, ShouldBeNil)
					if ti.shouldTryToCreateVector {
						if !ti.makeSetVectorFail {
							vector{{.scalar.String}}Mock.AssertCalled(t, "SetVector{{.scalar.String}}", ti.fn)
						}

						if !ti.makePushFail {
							for _, item := range ti.mockVector {
								{{.scalar.NativeType}}Mock.AssertCalled(t, "Push{{.scalar.String}}", item)
							}
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
	f, err := os.Create("scalar-decoder.gen.go")
	die(err)
	defer f.Close()
	_, err = f.Write([]byte("// Code generated for marshaling/unmarshaling. DO NOT EDIT.\n"))
	die(err)
	_, err = f.Write([]byte(`package omarshaller
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
	RenderVectorScalar_TEST_Decoders(test)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
