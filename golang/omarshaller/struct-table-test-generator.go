// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates scalar-decoder.gen.go It can be invoked by running
// go generate

package main

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"io"
	"text/template"
	"os"
	"log"
)

func generate_struct_tests(f io.Writer) {

	var err error
	var structTestTemplate *template.Template

	structTestTemplate, err = template.New("structTestTemplate").Parse(`
        Convey("Test {{.scalar.String}} accessor", func() {
	        d := &Decoder{application: "test"}
	        structTest := map[string]interface{}{
	            "{{.scalar.String}}": {{.scalar.NativeType}}({{.value}}),
	        }

	        sw := &mocks.ScalarWriter{}
	        sw.On("Set{{.scalar.String}}", hybrids.FieldNumber(10), {{.scalar.NativeType}}({{.value}})).Return(nil)

	        f{{.scalar.String}}Mock := &omocks.FieldReader{}
	        f{{.scalar.String}}Mock.On("Type").Return("{{.scalar.String}}")

	        fm := &oreflection.FieldMap{
	             Camel: map[string]*oreflection.FieldWrapper{
	                 "{{.scalar.String}}":      {f{{.scalar.String}}Mock, 10},
	             },
	        }

	        rs := &rmocks.ReflectStore{}
	        d.reflect = rs

	        rs.On("OReflect", "{{.scalar.String}}").Return(&oreflection.OType{HybridType: hybrids.{{.scalar.String}}}, nil)
	        err := d.decodeStruct("struct", structTest, &oreflection.OType{FieldMap: fm}, sw)
	        So(err, ShouldBeNil)
	        sw.AssertCalled(t, "Set{{.scalar.String}}", hybrids.FieldNumber(10), {{.scalar.NativeType}}({{.value}}))
	    })
`)

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

	scalarsValues := []string{
		"true",
		"-8",
		"8",
		"-16",
		"16",
		"-32",
		"32",
		"-64",
		"64",
		"-32.5",
		"64.5",
	}

	for index, v := range scalars {
		err = structTestTemplate.Execute(f, map[string]interface{}{
			"scalar": v,
			"value":  scalarsValues[index],
		})
		die(err)
	}

}

func main() {
	f, err := os.Create("struct_accessors_test.go")
	die(err)
	defer f.Close()
	_, err = f.Write([]byte("// Code generated for marshaling/unmarshaling. DO NOT EDIT.\n"))
	die(err)
	_, err = f.Write([]byte(`package omarshaller
`))
	die(err)
	_, err = f.Write([]byte(`
import (
	"github.com/nebtex/hybrids/golang/hybrids"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/nebtex/hybrids/golang/hybrids/mocks"
	omocks "github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1/mocks"
	rmocks "github.com/nebtex/omnibuff/tools/golang/tools/oreflection/mocks"
	"github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
)

func Test_StructScalarAccessors(t *testing.T) {
    Convey("Test Struct Scalar Accessors", t, func() {

`))
	die(err)
	generate_struct_tests(f)
	_, err = f.Write([]byte(`
    })
}
	`))
	die(err)

}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
