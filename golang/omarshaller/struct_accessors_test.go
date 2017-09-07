// Code generated for marshaling/unmarshaling. DO NOT EDIT.
package omarshaller

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


        Convey("Test Boolean accessor", func() {
	        d := &Decoder{application: "test"}
	        structTest := map[string]interface{}{
	            "Boolean": bool(true),
	        }

	        sw := &mocks.ScalarWriter{}
	        sw.On("SetBoolean", hybrids.FieldNumber(10), bool(true)).Return(nil)

	        fBooleanMock := &omocks.FieldReader{}
	        fBooleanMock.On("Type").Return("Boolean")

	        fm := &oreflection.FieldMap{
	             Camel: map[string]*oreflection.FieldWrapper{
	                 "Boolean":      {fBooleanMock, 10},
	             },
	        }

	        rs := &rmocks.ReflectStore{}
	        d.reflect = rs

	        rs.On("OReflect", "Boolean").Return(&oreflection.OType{HybridType: hybrids.Boolean}, nil)
	        err := d.decodeStruct("struct", structTest, &oreflection.OType{FieldMap: fm}, sw)
	        So(err, ShouldBeNil)
	        sw.AssertCalled(t, "SetBoolean", hybrids.FieldNumber(10), bool(true))
	    })

        Convey("Test Int8 accessor", func() {
	        d := &Decoder{application: "test"}
	        structTest := map[string]interface{}{
	            "Int8": int8(-8),
	        }

	        sw := &mocks.ScalarWriter{}
	        sw.On("SetInt8", hybrids.FieldNumber(10), int8(-8)).Return(nil)

	        fInt8Mock := &omocks.FieldReader{}
	        fInt8Mock.On("Type").Return("Int8")

	        fm := &oreflection.FieldMap{
	             Camel: map[string]*oreflection.FieldWrapper{
	                 "Int8":      {fInt8Mock, 10},
	             },
	        }

	        rs := &rmocks.ReflectStore{}
	        d.reflect = rs

	        rs.On("OReflect", "Int8").Return(&oreflection.OType{HybridType: hybrids.Int8}, nil)
	        err := d.decodeStruct("struct", structTest, &oreflection.OType{FieldMap: fm}, sw)
	        So(err, ShouldBeNil)
	        sw.AssertCalled(t, "SetInt8", hybrids.FieldNumber(10), int8(-8))
	    })

        Convey("Test Uint8 accessor", func() {
	        d := &Decoder{application: "test"}
	        structTest := map[string]interface{}{
	            "Uint8": uint8(8),
	        }

	        sw := &mocks.ScalarWriter{}
	        sw.On("SetUint8", hybrids.FieldNumber(10), uint8(8)).Return(nil)

	        fUint8Mock := &omocks.FieldReader{}
	        fUint8Mock.On("Type").Return("Uint8")

	        fm := &oreflection.FieldMap{
	             Camel: map[string]*oreflection.FieldWrapper{
	                 "Uint8":      {fUint8Mock, 10},
	             },
	        }

	        rs := &rmocks.ReflectStore{}
	        d.reflect = rs

	        rs.On("OReflect", "Uint8").Return(&oreflection.OType{HybridType: hybrids.Uint8}, nil)
	        err := d.decodeStruct("struct", structTest, &oreflection.OType{FieldMap: fm}, sw)
	        So(err, ShouldBeNil)
	        sw.AssertCalled(t, "SetUint8", hybrids.FieldNumber(10), uint8(8))
	    })

        Convey("Test Int16 accessor", func() {
	        d := &Decoder{application: "test"}
	        structTest := map[string]interface{}{
	            "Int16": int16(-16),
	        }

	        sw := &mocks.ScalarWriter{}
	        sw.On("SetInt16", hybrids.FieldNumber(10), int16(-16)).Return(nil)

	        fInt16Mock := &omocks.FieldReader{}
	        fInt16Mock.On("Type").Return("Int16")

	        fm := &oreflection.FieldMap{
	             Camel: map[string]*oreflection.FieldWrapper{
	                 "Int16":      {fInt16Mock, 10},
	             },
	        }

	        rs := &rmocks.ReflectStore{}
	        d.reflect = rs

	        rs.On("OReflect", "Int16").Return(&oreflection.OType{HybridType: hybrids.Int16}, nil)
	        err := d.decodeStruct("struct", structTest, &oreflection.OType{FieldMap: fm}, sw)
	        So(err, ShouldBeNil)
	        sw.AssertCalled(t, "SetInt16", hybrids.FieldNumber(10), int16(-16))
	    })

        Convey("Test Uint16 accessor", func() {
	        d := &Decoder{application: "test"}
	        structTest := map[string]interface{}{
	            "Uint16": uint16(16),
	        }

	        sw := &mocks.ScalarWriter{}
	        sw.On("SetUint16", hybrids.FieldNumber(10), uint16(16)).Return(nil)

	        fUint16Mock := &omocks.FieldReader{}
	        fUint16Mock.On("Type").Return("Uint16")

	        fm := &oreflection.FieldMap{
	             Camel: map[string]*oreflection.FieldWrapper{
	                 "Uint16":      {fUint16Mock, 10},
	             },
	        }

	        rs := &rmocks.ReflectStore{}
	        d.reflect = rs

	        rs.On("OReflect", "Uint16").Return(&oreflection.OType{HybridType: hybrids.Uint16}, nil)
	        err := d.decodeStruct("struct", structTest, &oreflection.OType{FieldMap: fm}, sw)
	        So(err, ShouldBeNil)
	        sw.AssertCalled(t, "SetUint16", hybrids.FieldNumber(10), uint16(16))
	    })

        Convey("Test Int32 accessor", func() {
	        d := &Decoder{application: "test"}
	        structTest := map[string]interface{}{
	            "Int32": int32(-32),
	        }

	        sw := &mocks.ScalarWriter{}
	        sw.On("SetInt32", hybrids.FieldNumber(10), int32(-32)).Return(nil)

	        fInt32Mock := &omocks.FieldReader{}
	        fInt32Mock.On("Type").Return("Int32")

	        fm := &oreflection.FieldMap{
	             Camel: map[string]*oreflection.FieldWrapper{
	                 "Int32":      {fInt32Mock, 10},
	             },
	        }

	        rs := &rmocks.ReflectStore{}
	        d.reflect = rs

	        rs.On("OReflect", "Int32").Return(&oreflection.OType{HybridType: hybrids.Int32}, nil)
	        err := d.decodeStruct("struct", structTest, &oreflection.OType{FieldMap: fm}, sw)
	        So(err, ShouldBeNil)
	        sw.AssertCalled(t, "SetInt32", hybrids.FieldNumber(10), int32(-32))
	    })

        Convey("Test Uint32 accessor", func() {
	        d := &Decoder{application: "test"}
	        structTest := map[string]interface{}{
	            "Uint32": uint32(32),
	        }

	        sw := &mocks.ScalarWriter{}
	        sw.On("SetUint32", hybrids.FieldNumber(10), uint32(32)).Return(nil)

	        fUint32Mock := &omocks.FieldReader{}
	        fUint32Mock.On("Type").Return("Uint32")

	        fm := &oreflection.FieldMap{
	             Camel: map[string]*oreflection.FieldWrapper{
	                 "Uint32":      {fUint32Mock, 10},
	             },
	        }

	        rs := &rmocks.ReflectStore{}
	        d.reflect = rs

	        rs.On("OReflect", "Uint32").Return(&oreflection.OType{HybridType: hybrids.Uint32}, nil)
	        err := d.decodeStruct("struct", structTest, &oreflection.OType{FieldMap: fm}, sw)
	        So(err, ShouldBeNil)
	        sw.AssertCalled(t, "SetUint32", hybrids.FieldNumber(10), uint32(32))
	    })

        Convey("Test Int64 accessor", func() {
	        d := &Decoder{application: "test"}
	        structTest := map[string]interface{}{
	            "Int64": int64(-64),
	        }

	        sw := &mocks.ScalarWriter{}
	        sw.On("SetInt64", hybrids.FieldNumber(10), int64(-64)).Return(nil)

	        fInt64Mock := &omocks.FieldReader{}
	        fInt64Mock.On("Type").Return("Int64")

	        fm := &oreflection.FieldMap{
	             Camel: map[string]*oreflection.FieldWrapper{
	                 "Int64":      {fInt64Mock, 10},
	             },
	        }

	        rs := &rmocks.ReflectStore{}
	        d.reflect = rs

	        rs.On("OReflect", "Int64").Return(&oreflection.OType{HybridType: hybrids.Int64}, nil)
	        err := d.decodeStruct("struct", structTest, &oreflection.OType{FieldMap: fm}, sw)
	        So(err, ShouldBeNil)
	        sw.AssertCalled(t, "SetInt64", hybrids.FieldNumber(10), int64(-64))
	    })

        Convey("Test Uint64 accessor", func() {
	        d := &Decoder{application: "test"}
	        structTest := map[string]interface{}{
	            "Uint64": uint64(64),
	        }

	        sw := &mocks.ScalarWriter{}
	        sw.On("SetUint64", hybrids.FieldNumber(10), uint64(64)).Return(nil)

	        fUint64Mock := &omocks.FieldReader{}
	        fUint64Mock.On("Type").Return("Uint64")

	        fm := &oreflection.FieldMap{
	             Camel: map[string]*oreflection.FieldWrapper{
	                 "Uint64":      {fUint64Mock, 10},
	             },
	        }

	        rs := &rmocks.ReflectStore{}
	        d.reflect = rs

	        rs.On("OReflect", "Uint64").Return(&oreflection.OType{HybridType: hybrids.Uint64}, nil)
	        err := d.decodeStruct("struct", structTest, &oreflection.OType{FieldMap: fm}, sw)
	        So(err, ShouldBeNil)
	        sw.AssertCalled(t, "SetUint64", hybrids.FieldNumber(10), uint64(64))
	    })

        Convey("Test Float32 accessor", func() {
	        d := &Decoder{application: "test"}
	        structTest := map[string]interface{}{
	            "Float32": float32(-32.5),
	        }

	        sw := &mocks.ScalarWriter{}
	        sw.On("SetFloat32", hybrids.FieldNumber(10), float32(-32.5)).Return(nil)

	        fFloat32Mock := &omocks.FieldReader{}
	        fFloat32Mock.On("Type").Return("Float32")

	        fm := &oreflection.FieldMap{
	             Camel: map[string]*oreflection.FieldWrapper{
	                 "Float32":      {fFloat32Mock, 10},
	             },
	        }

	        rs := &rmocks.ReflectStore{}
	        d.reflect = rs

	        rs.On("OReflect", "Float32").Return(&oreflection.OType{HybridType: hybrids.Float32}, nil)
	        err := d.decodeStruct("struct", structTest, &oreflection.OType{FieldMap: fm}, sw)
	        So(err, ShouldBeNil)
	        sw.AssertCalled(t, "SetFloat32", hybrids.FieldNumber(10), float32(-32.5))
	    })

        Convey("Test Float64 accessor", func() {
	        d := &Decoder{application: "test"}
	        structTest := map[string]interface{}{
	            "Float64": float64(64.5),
	        }

	        sw := &mocks.ScalarWriter{}
	        sw.On("SetFloat64", hybrids.FieldNumber(10), float64(64.5)).Return(nil)

	        fFloat64Mock := &omocks.FieldReader{}
	        fFloat64Mock.On("Type").Return("Float64")

	        fm := &oreflection.FieldMap{
	             Camel: map[string]*oreflection.FieldWrapper{
	                 "Float64":      {fFloat64Mock, 10},
	             },
	        }

	        rs := &rmocks.ReflectStore{}
	        d.reflect = rs

	        rs.On("OReflect", "Float64").Return(&oreflection.OType{HybridType: hybrids.Float64}, nil)
	        err := d.decodeStruct("struct", structTest, &oreflection.OType{FieldMap: fm}, sw)
	        So(err, ShouldBeNil)
	        sw.AssertCalled(t, "SetFloat64", hybrids.FieldNumber(10), float64(64.5))
	    })

    })
}
	