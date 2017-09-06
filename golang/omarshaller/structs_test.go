package omarshaller

import (
	"testing"
	"github.com/nebtex/hybrids/golang/hybrids/mocks"
	omocks "github.com/nebtex/omnibuff/pkg/io/omniql/Nextcorev1/mocks"
	rmocks "github.com/nebtex/omnibuff/tools/golang/tools/oreflection/mocks"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
	"github.com/nebtex/hybrids/golang/hybrids"
)

func Test_DecoderStruct(t *testing.T) {
	Convey("Test with every scalar accessor", t, func() {

		structTest := map[string]interface{}{
			"Boolean":       true,
			"Int8":          int8(-8),
			"Uint8":         uint8(8),
			"Int16":         int16(-16),
			"Uint16":        uint16(16),
			"Int32":         int32(-32),
			"Uint32":        uint32(32),
			"Int64":         int64(-64),
			"Uint64":        uint64(64),
			"Float32":       float32(-64.5),
			"Float64":       float64(64.5),
			"snake_float64": float64(1256897),

		}
		d := &Decoder{application: "test"}

		//create mocks
		sw := &mocks.ScalarWriter{}
		sw.On("SetBoolean", hybrids.FieldNumber(10), true).Return(nil)
		sw.On("SetInt8", hybrids.FieldNumber(9), int8(-8)).Return(nil)
		sw.On("SetUint8", hybrids.FieldNumber(8), uint8(8)).Return(nil)
		sw.On("SetInt16", hybrids.FieldNumber(7), int16(-16)).Return(nil)
		sw.On("SetUint16", hybrids.FieldNumber(6), uint16(16)).Return(nil)
		sw.On("SetInt32", hybrids.FieldNumber(5), int32(-32)).Return(nil)
		sw.On("SetUint32", hybrids.FieldNumber(4), uint32(32)).Return(nil)
		sw.On("SetInt64", hybrids.FieldNumber(3), int64(-64)).Return(nil)
		sw.On("SetUint64", hybrids.FieldNumber(2), uint64(64)).Return(nil)
		sw.On("SetFloat32", hybrids.FieldNumber(1), float32(-64.5)).Return(nil)
		sw.On("SetFloat64", hybrids.FieldNumber(0), float64(64.5)).Return(nil)
		sw.On("SetFloat64", hybrids.FieldNumber(11), float64(1256897)).Return(nil)

		fBooleanMock := &omocks.FieldReader{}
		fBooleanMock.On("Type").Return("Boolean")
		fInt8Mock := &omocks.FieldReader{}
		fInt8Mock.On("Type").Return("Int8")
		fUint8Mock := &omocks.FieldReader{}
		fUint8Mock.On("Type").Return("Uint8")
		fInt16Mock := &omocks.FieldReader{}
		fInt16Mock.On("Type").Return("Int16")
		fUint16Mock := &omocks.FieldReader{}
		fUint16Mock.On("Type").Return("Uint16")
		fInt32Mock := &omocks.FieldReader{}
		fInt32Mock.On("Type").Return("Int32")
		fUint32Mock := &omocks.FieldReader{}
		fUint32Mock.On("Type").Return("Uint32")
		fInt64Mock := &omocks.FieldReader{}
		fInt64Mock.On("Type").Return("Int64")
		fUint64Mock := &omocks.FieldReader{}
		fUint64Mock.On("Type").Return("Uint64")
		fFloat32Mock := &omocks.FieldReader{}
		fFloat32Mock.On("Type").Return("Float32")
		fFloat64Mock := &omocks.FieldReader{}
		fFloat64Mock.On("Type").Return("Float64")
		fSnakeFloat64Mock := &omocks.FieldReader{}
		fSnakeFloat64Mock.On("Type").Return("Float64")

		fm := &oreflection.FieldMap{
			Snake: map[string]*oreflection.FieldWrapper{
				"boolean":       {fBooleanMock, 10},
				"int8":          {fInt8Mock, 9},
				"uint8":         {fUint8Mock, 8},
				"int16":         {fInt16Mock, 7},
				"uint16":        {fUint16Mock, 6},
				"int32":         {fInt32Mock, 5},
				"uint32":        {fUint32Mock, 4},
				"int64":         {fInt64Mock, 3},
				"uint64":        {fUint64Mock, 2},
				"float32":       {fFloat32Mock, 1},
				"float64":       {fFloat64Mock, 0},
				"snake_float64": {fSnakeFloat64Mock, 11},

			},
			Camel: map[string]*oreflection.FieldWrapper{
				"Boolean":      {fBooleanMock, 10},
				"Int8":         {fInt8Mock, 9},
				"Uint8":        {fUint8Mock, 8},
				"Int16":        {fInt16Mock, 7},
				"Uint16":       {fUint16Mock, 6},
				"Int32":        {fInt32Mock, 5},
				"Uint32":       {fUint32Mock, 4},
				"Int64":        {fInt64Mock, 3},
				"Uint64":       {fUint64Mock, 2},
				"Float32":      {fFloat32Mock, 1},
				"Float64":      {fFloat64Mock, 0},
				"SnakeFloat64": {fSnakeFloat64Mock, 11},
			},
		}
		rs := &rmocks.ReflectStore{}
		d.reflect = rs
		rs.On("OReflect", "Boolean").Return(&oreflection.OType{HybridType: hybrids.Boolean}, nil)
		rs.On("OReflect", "Int8").Return(&oreflection.OType{HybridType: hybrids.Int8}, nil)
		rs.On("OReflect", "Uint8").Return(&oreflection.OType{HybridType: hybrids.Uint8}, nil)
		rs.On("OReflect", "Int16").Return(&oreflection.OType{HybridType: hybrids.Int16}, nil)
		rs.On("OReflect", "Uint16").Return(&oreflection.OType{HybridType: hybrids.Uint16}, nil)
		rs.On("OReflect", "Int32").Return(&oreflection.OType{HybridType: hybrids.Int32}, nil)
		rs.On("OReflect", "Uint32").Return(&oreflection.OType{HybridType: hybrids.Uint32}, nil)
		rs.On("OReflect", "Int64").Return(&oreflection.OType{HybridType: hybrids.Int64}, nil)
		rs.On("OReflect", "Uint64").Return(&oreflection.OType{HybridType: hybrids.Uint64}, nil)
		rs.On("OReflect", "Float32").Return(&oreflection.OType{HybridType: hybrids.Float32}, nil)
		rs.On("OReflect", "Float64").Return(&oreflection.OType{HybridType: hybrids.Float64}, nil)

		err := d.decodeStruct("struct", structTest, &oreflection.OType{FieldMap: fm}, sw)
		sw.AssertCalled(t, "SetBoolean", hybrids.FieldNumber(10), true)
		sw.AssertCalled(t, "SetInt8", hybrids.FieldNumber(9), int8(-8))
		sw.AssertCalled(t, "SetUint8", hybrids.FieldNumber(8), uint8(8))
		sw.AssertCalled(t, "SetInt16", hybrids.FieldNumber(7), int16(-16))
		sw.AssertCalled(t, "SetUint16", hybrids.FieldNumber(6), uint16(16))
		sw.AssertCalled(t, "SetInt32", hybrids.FieldNumber(5), int32(-32))
		sw.AssertCalled(t, "SetUint32", hybrids.FieldNumber(4), uint32(32))
		sw.AssertCalled(t, "SetInt64", hybrids.FieldNumber(3), int64(-64))
		sw.AssertCalled(t, "SetUint64", hybrids.FieldNumber(2), uint64(64))
		sw.AssertCalled(t, "SetFloat32", hybrids.FieldNumber(1), float32(-64.5))
		sw.AssertCalled(t, "SetFloat64", hybrids.FieldNumber(0), float64(64.5))
		sw.AssertCalled(t, "SetFloat64", hybrids.FieldNumber(11), float64(1256897))
		So(err, ShouldBeNil)

	})
	/*
		Convey("Struct Tests", t, func() {
			table := []struct {
				fn             hybrids.FieldNumber
				value          interface{}
				mockString     string
				shouldFail     bool
				shouldMockFail bool
				name           string
			}{
				{1, "jupiter", "jupiter", false, false, "native_type"},
				{25, 54545, "", true, false, "incorrect underlying type"},
				{10, "sun", "sun", true, true, "hybrid set method fail"},
				{10, nil, "", false, false, "null should not return error, and should call the set method"},
				{10, "", "", false, false, "empty string should  call the set method"},

			}

			d := &Decoder{application: "test"}

			for _, ti := range table {
				Convey(fmt.Sprintf("Test: %s", ti.name), func() {
					stringMock := &mocks.StringWriterAccessor{}
					call := stringMock.On("SetString", ti.fn, ti.mockString)
					if ti.shouldMockFail {
						call.Return(fmt.Errorf("SetString failed"))
					} else {
						call.Return(nil)
					}

					err := d.decodeString("x.path", ti.value, ti.fn, stringMock)
					if ti.shouldFail {
						t.Log(err)
						So(err, ShouldNotBeNil)
						de, _ := err.(*DecodeError)
						So(de.Application, ShouldEqual, "test")
						So(de.Path, ShouldEqual, "x.path")
						So(de.HybridType, ShouldEqual, hybrids.String)
						So(de.OmniqlType, ShouldEqual, "String")

					} else {
						So(err, ShouldBeNil)

						if !ti.shouldFail {
							stringMock.AssertCalled(t, "SetString", ti.fn, ti.mockString)
						}

					}
				})
			}
		})*/
}
