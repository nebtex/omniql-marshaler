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
					"Boolean": {fBooleanMock, 10},
				},
			}
			err := d.decodeStruct("struct", []interface{}{1, 2, 3}, &oreflection.OType{FieldMap: fm, Id: "Struct/All"}, sw)
			So(err, ShouldNotBeNil)
			fmt.Println(err)
		})

		Convey("Ignore field if is not present in the reflection", func() {
			isw := &mocks.ScalarWriter{}
			fBooleanMock := &omocks.FieldReader{}
			fBooleanMock.On("Type").Return("Boolean")
			fm := &oreflection.FieldMap{
				Camel: map[string]*oreflection.FieldWrapper{
					"Boolean": {fBooleanMock, 10},
				},
			}
			err := d.decodeStruct("struct", map[string]interface{}{"no_field": 30}, &oreflection.OType{FieldMap: fm, Id: "Struct/All"}, isw)
			So(err, ShouldBeNil)
			fmt.Println(err)
		})

		Convey("test when oreflection returns an error", func() {
			fBooleanMock := &omocks.FieldReader{}
			fBooleanMock.On("Type").Return("Boolean")
			fm := &oreflection.FieldMap{
				Camel: map[string]*oreflection.FieldWrapper{
					"Boolean": {fBooleanMock, 10},
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
					"Boolean": {fBooleanMock, 10},
				},
			}
			frs := &rmocks.ReflectStore{}
			d.reflect = frs
			frs.On("OReflect", "Boolean").Return(&oreflection.OType{HybridType: hybrids.VectorString}, nil)
			err := d.decodeStruct("struct", map[string]interface{}{"Boolean": true}, &oreflection.OType{FieldMap: fm, Id: "Struct/All"}, sw)
			So(err, ShouldNotBeNil)
		})

	})

}

func Test_Decode_VectorStruct(t *testing.T) {
	Convey("Decode VectorStruct test", t, func() {
		table := []struct {
			fn                     hybrids.FieldNumber
			vector                 interface{}
			shouldDecodeStructFail bool
			shouldFail             bool
			makePushFail           bool
			name                   string
		}{
			{30, nil, false, false, false, "null: should do nothing (remember vector and tables can have a null state)"},
			{25, "vector", false, true, false, "incorrect underlying type"},
			{50, []interface{}{uint64(1), uint64(2), uint64(3)}, false, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{uint64(1), uint64(2), uint64(3)}, true, true, false, "Should fails when decode struct fails"},
			{50, []interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}, false, false, false, "Valid input, all should be ok"},
		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				swMock := &mocks.ScalarWriter{}

				vsMock := &mocks.VectorStructWriter{}
				callItem := vsMock.On("PushStruct")
				if ti.makePushFail {
					callItem.Return(nil, fmt.Errorf("PushStruct failed"))
				} else {
					callItem.Return(swMock, nil)
				}

				err := d.decodeVectorStruct("x.path.vector", ti.vector,
					&oreflection.OType{
						HybridType: hybrids.VectorStruct,
						Id:         "Vector[Struct/Struct]",
						Items:      &oreflection.OType{Id: "Struct/Struct", HybridType: hybrids.Struct}},
					vsMock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					if ti.shouldDecodeStructFail {
						So(de.HybridType, ShouldEqual, hybrids.Struct)
						So(de.OmniqlType, ShouldEqual, "Struct/Struct")
					} else {
						So(de.HybridType, ShouldEqual, hybrids.VectorStruct)
						So(de.OmniqlType, ShouldEqual, "Vector[Struct/Struct]")
					}

				} else {
					So(err, ShouldBeNil)
					if !ti.makePushFail {
						if ti.vector != nil {
							vsMock.AssertCalled(t, "PushStruct")
						}
					}
				}
			})
		}

	})

}
