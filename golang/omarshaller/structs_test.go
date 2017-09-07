package omarshaller

import (
	"testing"
	"github.com/nebtex/hybrids/golang/hybrids/mocks"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
	"fmt"
	omocks "github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1/mocks"
	rmocks "github.com/nebtex/omnibuff/tools/golang/tools/oreflection/mocks"

	"github.com/nebtex/hybrids/golang/hybrids"
)

func Test_DecoderStruct(t *testing.T) {
	Convey("Test with every scalar accessor", t, func() {
		d := &Decoder{application: "test"}
		sw := &mocks.ScalarWriter{}

		Convey("Fail if input have a invalid data type", func() {
			fBooleanMock := &omocks.FieldReader{}
			fBooleanMock.On("Type").Return("Boolean")

			fm := &oreflection.FieldMap{
				Camel: map[string]*oreflection.FieldWrapper{
					"Boolean":      {fBooleanMock, 10},
				},
			}
			err := d.decodeStruct("struct", []interface{}{1, 2, 3}, &oreflection.OType{FieldMap: fm, Id: "Struct/All"}, sw)
			So(err, ShouldNotBeNil)
			fmt.Println(err)
		})
/*
		Convey("Ignore field if is not present in the reflection", func() {
			isw := &mocks.ScalarWriter{}
			err := d.decodeStruct("struct", map[string]interface{}{"no_field": 30}, &oreflection.OType{FieldMap: fm, Id: "Struct/All"}, isw)
			So(err, ShouldBeNil)
			fmt.Println(err)
		})*/

			Convey("test when oreflection returns an error", func() {
				fBooleanMock := &omocks.FieldReader{}
				fBooleanMock.On("Type").Return("Boolean")
				fm := &oreflection.FieldMap{
					Camel: map[string]*oreflection.FieldWrapper{
						"Boolean":      {fBooleanMock, 10},
					},
				}
				frs := &rmocks.ReflectStore{}
				d.reflect = frs
				frs.On("OReflect", "Boolean").Return(nil, fmt.Errorf("reflection failed"))
				err := d.decodeStruct("struct", map[string]interface{}{"Boolean": true}, &oreflection.OType{FieldMap: fm, Id: "Struct/All"}, sw)
				So(err, ShouldNotBeNil)
			})

		Convey("test when oreflection returns an invalid type", func() {
			fBooleanMock := &omocks.FieldReader{}
			fBooleanMock.On("Type").Return("Boolean")
			fm := &oreflection.FieldMap{
				Camel: map[string]*oreflection.FieldWrapper{
					"Boolean":      {fBooleanMock, 10},
				},
			}
			frs := &rmocks.ReflectStore{}
			d.reflect = frs
			frs.On("OReflect", "Boolean").Return(&oreflection.OType{HybridType:hybrids.VectorString}, nil)
			err := d.decodeStruct("struct", map[string]interface{}{"Boolean": true}, &oreflection.OType{FieldMap: fm, Id: "Struct/All"}, sw)
			So(err, ShouldNotBeNil)
		})

	})

}
