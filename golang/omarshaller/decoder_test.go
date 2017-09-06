// Code generated. DO NOT EDIT.
package omarshaller

import (
	"testing"
	"github.com/nebtex/hybrids/golang/hybrids/mocks"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
)

func Test_DecoderFloat64(t *testing.T) {
	Convey("Float64 Tests", t, func() {
		table := []struct {
			fn             hybrids.FieldNumber
			number         interface{}
			mockNumber     float64
			shouldFail     bool
			shouldMockFail bool
			name           string
		}{
			{1, float64(25), 25, false, false, "native_type"},
			{25, "2656", 0, true, false, "incorrect underlying type"},
			{10, float64(44), 44, true, true, "hybrid set method fail"},
			{10, nil, 0, true, false, "null should return error in any scalar"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				float64Mock := &mocks.Float64WriterAccessor{}

				call := float64Mock.On("SetFloat64", ti.fn, ti.mockNumber)
				if ti.shouldMockFail {
					call.Return(fmt.Errorf("SetFloat64 failed"))
				} else {
					call.Return(nil)
				}

				err := d.decodeFloat64("x.path", ti.number, ti.fn, float64Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.Path, ShouldEqual, "x.path")
					So(de.HybridType, ShouldEqual, hybrids.Float64)
					So(de.OmniqlType, ShouldEqual, "Float64")

				} else {
					So(err, ShouldBeNil)
					if !ti.shouldFail {
						float64Mock.AssertCalled(t, "SetFloat64", ti.fn, ti.mockNumber)
					}
				}
			})
		}
	})

}

func Test_DecoderUint64(t *testing.T) {
	Convey("Uint64 Tests", t, func() {
		table := []struct {
			fn             hybrids.FieldNumber
			number         interface{}
			mockNumber     uint64
			shouldFail     bool
			shouldMockFail bool
			name           string
		}{
			{1, float64(77), 77, false, false, "float64 as underlying type"},
			{1, uint64(25), 25, false, false, "native_type"},
			{25, false, 0, true, false, "incorrect underlying type"},
			{25, "Nan", 0, true, false, "incorrect string"},
			{25, "2300", 2300, false, false, "using string as underlying type"},
			{10, uint64(44), 44, true, true, "hybrid set method fail"},
			{10, nil, 0, true, false, "null should return error in any scalar"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				uint64Mock := &mocks.Uint64WriterAccessor{}

				call := uint64Mock.On("SetUint64", ti.fn, ti.mockNumber)
				if ti.shouldMockFail {
					call.Return(fmt.Errorf("SetUint64 failed"))
				} else {
					call.Return(nil)
				}

				err := d.decodeUint64("x.path", ti.number, ti.fn, uint64Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.Path, ShouldEqual, "x.path")
					So(de.HybridType, ShouldEqual, hybrids.Uint64)
					So(de.OmniqlType, ShouldEqual, "Uint64")

				} else {
					So(err, ShouldBeNil)
					if !ti.shouldFail {
						uint64Mock.AssertCalled(t, "SetUint64", ti.fn, ti.mockNumber)
					}
				}
			})
		}
	})

}

func Test_DecoderInt64(t *testing.T) {
	Convey("Int64 Tests", t, func() {
		table := []struct {
			fn             hybrids.FieldNumber
			number         interface{}
			mockNumber     int64
			shouldFail     bool
			shouldMockFail bool
			name           string
		}{
			{1, float64(77), 77, false, false, "float64 as underlying type"},
			{1, int64(25), 25, false, false, "native_type"},
			{25, false, 0, true, false, "incorrect underlying type"},
			{25, "Nan", 0, true, false, "incorrect string"},
			{25, "2300", 2300, false, false, "using string as underlying type"},
			{10, int64(44), 44, true, true, "hybrid set method fail"},
			{10, nil, 0, true, false, "null should return error in any scalar"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				int64Mock := &mocks.Int64WriterAccessor{}

				call := int64Mock.On("SetInt64", ti.fn, ti.mockNumber)
				if ti.shouldMockFail {
					call.Return(fmt.Errorf("SetInt64 failed"))
				} else {
					call.Return(nil)
				}

				err := d.decodeInt64("x.path", ti.number, ti.fn, int64Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.Path, ShouldEqual, "x.path")
					So(de.HybridType, ShouldEqual, hybrids.Int64)
					So(de.OmniqlType, ShouldEqual, "Int64")

				} else {
					So(err, ShouldBeNil)
					if !ti.shouldFail {
						int64Mock.AssertCalled(t, "SetInt64", ti.fn, ti.mockNumber)
					}
				}
			})
		}
	})

}

func Test_DecoderBoolean(t *testing.T) {
	Convey("Boolean Tests", t, func() {
		table := []struct {
			fn             hybrids.FieldNumber
			number         interface{}
			mockNumber     bool
			shouldFail     bool
			shouldMockFail bool
			name           string
		}{
			{1, true, true, false, false, "native_type - true"},
			{1, false, false, false, false, "native_type - false"},
			{25, "2656", false, true, false, "incorrect underlying type"},
			{10, true, true, true, true, "hybrid set method fail"},
			{10, nil, false, true, false, "null should return error in any scalar"},
		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				boolMock := &mocks.BooleanWriterAccessor{}

				call := boolMock.On("SetBoolean", ti.fn, ti.mockNumber)
				if ti.shouldMockFail {
					call.Return(fmt.Errorf("SetBoolean failed"))
				} else {
					call.Return(nil)
				}

				err := d.decodeBoolean("x.path", ti.number, ti.fn, boolMock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.Path, ShouldEqual, "x.path")
					So(de.HybridType, ShouldEqual, hybrids.Boolean)
					So(de.OmniqlType, ShouldEqual, "Boolean")

				} else {
					So(err, ShouldBeNil)
					if !ti.shouldFail {
						boolMock.AssertCalled(t, "SetBoolean", ti.fn, ti.mockNumber)
					}
				}
			})
		}
	})

}

func Test_Decode_VectorBoolean(t *testing.T) {
	Convey("Decode VectorBoolean test", t, func() {
		table := []struct {
			fn                      hybrids.FieldNumber
			vector                  interface{}
			mockVector              []bool
			shouldFail              bool
			makePushFail            bool
			shouldTryToCreateVector bool
			name                    string
		}{
			{30, nil, nil, false, false, true, "null: should not fail"},
			{25, "vector", nil, true, false, false, "incorrect underlying type"},
			{70, []interface{}{nil}, []bool{}, true, false, true, "item: incorrect underlying type"},
			{50, []interface{}{false, true, false}, []bool{false, true, false}, true, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{false, true, false}, []bool{false, true, false}, false, false, true, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				boolMock := &mocks.VectorBooleanWriter{}
				for _, item := range ti.mockVector {
					callItem := boolMock.On("PushBoolean", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("PushBoolean failed"))
					} else {
						callItem.Return(nil)
					}
				}

				err := d.decodeVectorBoolean("x.path.vector", ti.vector, ti.fn, boolMock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.VectorBoolean)
					So(de.OmniqlType, ShouldEqual, "Vector[Boolean]")

				} else {
					So(err, ShouldBeNil)
					if ti.shouldTryToCreateVector {

						if !ti.makePushFail {
							for _, item := range ti.mockVector {
								boolMock.AssertCalled(t, "PushBoolean", item)
							}
						}
					}
				}
			})
		}
	})
}
