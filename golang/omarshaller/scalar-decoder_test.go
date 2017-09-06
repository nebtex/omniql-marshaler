// Code generated. DO NOT EDIT.
package omarshaller

import (
	"testing"
	"github.com/nebtex/hybrids/golang/hybrids/mocks"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
)


func Test_DecoderInt8(t *testing.T) {
	Convey("Int8 Tests", t, func() {
		table := []struct {
			fn             hybrids.FieldNumber
			number         interface{}
			mockNumber     int8
			shouldFail     bool
			shouldMockFail bool
			name           string
		}{
			{30, float64(35), 35, false, false, "underlying type: float64"},
			{1, int8(25), 25, false, false, "native_type"},
			{25, "2656", 0, true, false, "incorrect underlying type"},
			{70, float64(hybrids.MaxInt8) + 1e39, 0, true, false, "max value"},
			{50, float64(hybrids.MinInt8) - 1e39, 0, true, false, "min value"},
			{10, int8(44), 44, true, true, "hybrid set method fail"},
			{10, nil, 0, true, false, "null should return error in any scalar"},


		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				int8Mock := &mocks.Int8WriterAccessor{}

				call := int8Mock.On("SetInt8", ti.fn, ti.mockNumber)
				if ti.shouldMockFail {
					call.Return(fmt.Errorf("SetInt8 failed"))
				} else {
					call.Return(nil)
				}

				err := d.decodeInt8("x.path", ti.number, ti.fn, int8Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _:= err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.Path, ShouldEqual, "x.path")
					So(de.HybridType, ShouldEqual, hybrids.Int8)
					So(de.OmniqlType, ShouldEqual, "Int8")

				} else {
					So(err, ShouldBeNil)
					if !ti.shouldFail {
						int8Mock.AssertCalled(t, "SetInt8", ti.fn, ti.mockNumber)
					}
				}


			})
		}

	})

}


func Test_DecoderUint8(t *testing.T) {
	Convey("Uint8 Tests", t, func() {
		table := []struct {
			fn             hybrids.FieldNumber
			number         interface{}
			mockNumber     uint8
			shouldFail     bool
			shouldMockFail bool
			name           string
		}{
			{30, float64(35), 35, false, false, "underlying type: float64"},
			{1, uint8(25), 25, false, false, "native_type"},
			{25, "2656", 0, true, false, "incorrect underlying type"},
			{70, float64(hybrids.MaxUint8) + 1e39, 0, true, false, "max value"},
			{50, float64(hybrids.MinUint8) - 1e39, 0, true, false, "min value"},
			{10, uint8(44), 44, true, true, "hybrid set method fail"},
			{10, nil, 0, true, false, "null should return error in any scalar"},


		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				uint8Mock := &mocks.Uint8WriterAccessor{}

				call := uint8Mock.On("SetUint8", ti.fn, ti.mockNumber)
				if ti.shouldMockFail {
					call.Return(fmt.Errorf("SetUint8 failed"))
				} else {
					call.Return(nil)
				}

				err := d.decodeUint8("x.path", ti.number, ti.fn, uint8Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _:= err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.Path, ShouldEqual, "x.path")
					So(de.HybridType, ShouldEqual, hybrids.Uint8)
					So(de.OmniqlType, ShouldEqual, "Uint8")

				} else {
					So(err, ShouldBeNil)
					if !ti.shouldFail {
						uint8Mock.AssertCalled(t, "SetUint8", ti.fn, ti.mockNumber)
					}
				}


			})
		}

	})

}


func Test_DecoderInt16(t *testing.T) {
	Convey("Int16 Tests", t, func() {
		table := []struct {
			fn             hybrids.FieldNumber
			number         interface{}
			mockNumber     int16
			shouldFail     bool
			shouldMockFail bool
			name           string
		}{
			{30, float64(35), 35, false, false, "underlying type: float64"},
			{1, int16(25), 25, false, false, "native_type"},
			{25, "2656", 0, true, false, "incorrect underlying type"},
			{70, float64(hybrids.MaxInt16) + 1e39, 0, true, false, "max value"},
			{50, float64(hybrids.MinInt16) - 1e39, 0, true, false, "min value"},
			{10, int16(44), 44, true, true, "hybrid set method fail"},
			{10, nil, 0, true, false, "null should return error in any scalar"},


		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				int16Mock := &mocks.Int16WriterAccessor{}

				call := int16Mock.On("SetInt16", ti.fn, ti.mockNumber)
				if ti.shouldMockFail {
					call.Return(fmt.Errorf("SetInt16 failed"))
				} else {
					call.Return(nil)
				}

				err := d.decodeInt16("x.path", ti.number, ti.fn, int16Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _:= err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.Path, ShouldEqual, "x.path")
					So(de.HybridType, ShouldEqual, hybrids.Int16)
					So(de.OmniqlType, ShouldEqual, "Int16")

				} else {
					So(err, ShouldBeNil)
					if !ti.shouldFail {
						int16Mock.AssertCalled(t, "SetInt16", ti.fn, ti.mockNumber)
					}
				}


			})
		}

	})

}


func Test_DecoderUint16(t *testing.T) {
	Convey("Uint16 Tests", t, func() {
		table := []struct {
			fn             hybrids.FieldNumber
			number         interface{}
			mockNumber     uint16
			shouldFail     bool
			shouldMockFail bool
			name           string
		}{
			{30, float64(35), 35, false, false, "underlying type: float64"},
			{1, uint16(25), 25, false, false, "native_type"},
			{25, "2656", 0, true, false, "incorrect underlying type"},
			{70, float64(hybrids.MaxUint16) + 1e39, 0, true, false, "max value"},
			{50, float64(hybrids.MinUint16) - 1e39, 0, true, false, "min value"},
			{10, uint16(44), 44, true, true, "hybrid set method fail"},
			{10, nil, 0, true, false, "null should return error in any scalar"},


		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				uint16Mock := &mocks.Uint16WriterAccessor{}

				call := uint16Mock.On("SetUint16", ti.fn, ti.mockNumber)
				if ti.shouldMockFail {
					call.Return(fmt.Errorf("SetUint16 failed"))
				} else {
					call.Return(nil)
				}

				err := d.decodeUint16("x.path", ti.number, ti.fn, uint16Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _:= err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.Path, ShouldEqual, "x.path")
					So(de.HybridType, ShouldEqual, hybrids.Uint16)
					So(de.OmniqlType, ShouldEqual, "Uint16")

				} else {
					So(err, ShouldBeNil)
					if !ti.shouldFail {
						uint16Mock.AssertCalled(t, "SetUint16", ti.fn, ti.mockNumber)
					}
				}


			})
		}

	})

}


func Test_DecoderInt32(t *testing.T) {
	Convey("Int32 Tests", t, func() {
		table := []struct {
			fn             hybrids.FieldNumber
			number         interface{}
			mockNumber     int32
			shouldFail     bool
			shouldMockFail bool
			name           string
		}{
			{30, float64(35), 35, false, false, "underlying type: float64"},
			{1, int32(25), 25, false, false, "native_type"},
			{25, "2656", 0, true, false, "incorrect underlying type"},
			{70, float64(hybrids.MaxInt32) + 1e39, 0, true, false, "max value"},
			{50, float64(hybrids.MinInt32) - 1e39, 0, true, false, "min value"},
			{10, int32(44), 44, true, true, "hybrid set method fail"},
			{10, nil, 0, true, false, "null should return error in any scalar"},


		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				int32Mock := &mocks.Int32WriterAccessor{}

				call := int32Mock.On("SetInt32", ti.fn, ti.mockNumber)
				if ti.shouldMockFail {
					call.Return(fmt.Errorf("SetInt32 failed"))
				} else {
					call.Return(nil)
				}

				err := d.decodeInt32("x.path", ti.number, ti.fn, int32Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _:= err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.Path, ShouldEqual, "x.path")
					So(de.HybridType, ShouldEqual, hybrids.Int32)
					So(de.OmniqlType, ShouldEqual, "Int32")

				} else {
					So(err, ShouldBeNil)
					if !ti.shouldFail {
						int32Mock.AssertCalled(t, "SetInt32", ti.fn, ti.mockNumber)
					}
				}


			})
		}

	})

}


func Test_DecoderUint32(t *testing.T) {
	Convey("Uint32 Tests", t, func() {
		table := []struct {
			fn             hybrids.FieldNumber
			number         interface{}
			mockNumber     uint32
			shouldFail     bool
			shouldMockFail bool
			name           string
		}{
			{30, float64(35), 35, false, false, "underlying type: float64"},
			{1, uint32(25), 25, false, false, "native_type"},
			{25, "2656", 0, true, false, "incorrect underlying type"},
			{70, float64(hybrids.MaxUint32) + 1e39, 0, true, false, "max value"},
			{50, float64(hybrids.MinUint32) - 1e39, 0, true, false, "min value"},
			{10, uint32(44), 44, true, true, "hybrid set method fail"},
			{10, nil, 0, true, false, "null should return error in any scalar"},


		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				uint32Mock := &mocks.Uint32WriterAccessor{}

				call := uint32Mock.On("SetUint32", ti.fn, ti.mockNumber)
				if ti.shouldMockFail {
					call.Return(fmt.Errorf("SetUint32 failed"))
				} else {
					call.Return(nil)
				}

				err := d.decodeUint32("x.path", ti.number, ti.fn, uint32Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _:= err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.Path, ShouldEqual, "x.path")
					So(de.HybridType, ShouldEqual, hybrids.Uint32)
					So(de.OmniqlType, ShouldEqual, "Uint32")

				} else {
					So(err, ShouldBeNil)
					if !ti.shouldFail {
						uint32Mock.AssertCalled(t, "SetUint32", ti.fn, ti.mockNumber)
					}
				}


			})
		}

	})

}


func Test_DecoderFloat32(t *testing.T) {
	Convey("Float32 Tests", t, func() {
		table := []struct {
			fn             hybrids.FieldNumber
			number         interface{}
			mockNumber     float32
			shouldFail     bool
			shouldMockFail bool
			name           string
		}{
			{30, float64(35), 35, false, false, "underlying type: float64"},
			{1, float32(25), 25, false, false, "native_type"},
			{25, "2656", 0, true, false, "incorrect underlying type"},
			{70, float64(hybrids.MaxFloat32) + 1e39, 0, true, false, "max value"},
			{50, float64(hybrids.MinFloat32) - 1e39, 0, true, false, "min value"},
			{10, float32(44), 44, true, true, "hybrid set method fail"},
			{10, nil, 0, true, false, "null should return error in any scalar"},


		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				float32Mock := &mocks.Float32WriterAccessor{}

				call := float32Mock.On("SetFloat32", ti.fn, ti.mockNumber)
				if ti.shouldMockFail {
					call.Return(fmt.Errorf("SetFloat32 failed"))
				} else {
					call.Return(nil)
				}

				err := d.decodeFloat32("x.path", ti.number, ti.fn, float32Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _:= err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.Path, ShouldEqual, "x.path")
					So(de.HybridType, ShouldEqual, hybrids.Float32)
					So(de.OmniqlType, ShouldEqual, "Float32")

				} else {
					So(err, ShouldBeNil)
					if !ti.shouldFail {
						float32Mock.AssertCalled(t, "SetFloat32", ti.fn, ti.mockNumber)
					}
				}


			})
		}

	})

}

func Test_Decode_VectorInt8(t *testing.T) {
	Convey("Decode VectorInt8 test", t, func() {
		table := []struct {
			fn                      hybrids.FieldNumber
			vector                  interface{}
			mockVector              []int8
			shouldFail              bool
			makeSetVectorFail    bool
			makePushFail            bool
			shouldTryToCreateVector bool
			name                    string
		}{
			{30, nil, nil, false, false, false, true, "null: should create a vector but don't add items (remember vector and tables can have a null state)"},
			{25, "vector", nil, true, false, false, false, "incorrect underlying type"},
			{70, []interface{}{"nan"}, []int8{0}, true, false, false, true, "item: incorrect underlying type"},
			{50, []interface{}{1, 2, 3}, []int8{1, 2, 3}, true, true, false, true, "Should fails when the vector creation operation fails"},
			{50, []interface{}{ int8(1), int8(2), int8(3)}, []int8{1, 2, 3}, true, false, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{ int8(1), int8(2), int8(3)}, []int8{1, 2, 3}, false, false, false, true, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				int8Mock := &mocks.VectorInt8WriterBase{}
				for _, item := range ti.mockVector {
					callItem := int8Mock.On("PushInt8", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("PushInt8 failed"))
					} else {
						callItem.Return(nil)
					}
				}

				vectorInt8Mock := &mocks.VectorInt8WriterAccessor{}

				call := vectorInt8Mock.On("SetVectorInt8", ti.fn)
				if ti.makeSetVectorFail {
					call.Return(nil, fmt.Errorf("SetVectorInt8 failed"))
				} else {
					call.Return(int8Mock, nil)
				}



				err := d.decodeVectorInt8("x.path.vector", ti.vector, ti.fn, vectorInt8Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.VectorInt8)
					So(de.OmniqlType, ShouldEqual, "Vector[Int8]")

				} else {
					So(err, ShouldBeNil)
					if ti.shouldTryToCreateVector {
						if !ti.makeSetVectorFail {
							vectorInt8Mock.AssertCalled(t, "SetVectorInt8", ti.fn)
						}

						if !ti.makePushFail {
							for _, item := range ti.mockVector {
								int8Mock.AssertCalled(t, "PushInt8", item)
							}
						}
					}
				}
			})
		}

	})

}


func Test_Decode_VectorUint8(t *testing.T) {
	Convey("Decode VectorUint8 test", t, func() {
		table := []struct {
			fn                      hybrids.FieldNumber
			vector                  interface{}
			mockVector              []uint8
			shouldFail              bool
			makeSetVectorFail    bool
			makePushFail            bool
			shouldTryToCreateVector bool
			name                    string
		}{
			{30, nil, nil, false, false, false, true, "null: should create a vector but don't add items (remember vector and tables can have a null state)"},
			{25, "vector", nil, true, false, false, false, "incorrect underlying type"},
			{70, []interface{}{"nan"}, []uint8{0}, true, false, false, true, "item: incorrect underlying type"},
			{50, []interface{}{1, 2, 3}, []uint8{1, 2, 3}, true, true, false, true, "Should fails when the vector creation operation fails"},
			{50, []interface{}{ uint8(1), uint8(2), uint8(3)}, []uint8{1, 2, 3}, true, false, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{ uint8(1), uint8(2), uint8(3)}, []uint8{1, 2, 3}, false, false, false, true, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				uint8Mock := &mocks.VectorUint8WriterBase{}
				for _, item := range ti.mockVector {
					callItem := uint8Mock.On("PushUint8", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("PushUint8 failed"))
					} else {
						callItem.Return(nil)
					}
				}

				vectorUint8Mock := &mocks.VectorUint8WriterAccessor{}

				call := vectorUint8Mock.On("SetVectorUint8", ti.fn)
				if ti.makeSetVectorFail {
					call.Return(nil, fmt.Errorf("SetVectorUint8 failed"))
				} else {
					call.Return(uint8Mock, nil)
				}



				err := d.decodeVectorUint8("x.path.vector", ti.vector, ti.fn, vectorUint8Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.VectorUint8)
					So(de.OmniqlType, ShouldEqual, "Vector[Uint8]")

				} else {
					So(err, ShouldBeNil)
					if ti.shouldTryToCreateVector {
						if !ti.makeSetVectorFail {
							vectorUint8Mock.AssertCalled(t, "SetVectorUint8", ti.fn)
						}

						if !ti.makePushFail {
							for _, item := range ti.mockVector {
								uint8Mock.AssertCalled(t, "PushUint8", item)
							}
						}
					}
				}
			})
		}

	})

}


func Test_Decode_VectorInt16(t *testing.T) {
	Convey("Decode VectorInt16 test", t, func() {
		table := []struct {
			fn                      hybrids.FieldNumber
			vector                  interface{}
			mockVector              []int16
			shouldFail              bool
			makeSetVectorFail    bool
			makePushFail            bool
			shouldTryToCreateVector bool
			name                    string
		}{
			{30, nil, nil, false, false, false, true, "null: should create a vector but don't add items (remember vector and tables can have a null state)"},
			{25, "vector", nil, true, false, false, false, "incorrect underlying type"},
			{70, []interface{}{"nan"}, []int16{0}, true, false, false, true, "item: incorrect underlying type"},
			{50, []interface{}{1, 2, 3}, []int16{1, 2, 3}, true, true, false, true, "Should fails when the vector creation operation fails"},
			{50, []interface{}{ int16(1), int16(2), int16(3)}, []int16{1, 2, 3}, true, false, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{ int16(1), int16(2), int16(3)}, []int16{1, 2, 3}, false, false, false, true, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				int16Mock := &mocks.VectorInt16WriterBase{}
				for _, item := range ti.mockVector {
					callItem := int16Mock.On("PushInt16", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("PushInt16 failed"))
					} else {
						callItem.Return(nil)
					}
				}

				vectorInt16Mock := &mocks.VectorInt16WriterAccessor{}

				call := vectorInt16Mock.On("SetVectorInt16", ti.fn)
				if ti.makeSetVectorFail {
					call.Return(nil, fmt.Errorf("SetVectorInt16 failed"))
				} else {
					call.Return(int16Mock, nil)
				}



				err := d.decodeVectorInt16("x.path.vector", ti.vector, ti.fn, vectorInt16Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.VectorInt16)
					So(de.OmniqlType, ShouldEqual, "Vector[Int16]")

				} else {
					So(err, ShouldBeNil)
					if ti.shouldTryToCreateVector {
						if !ti.makeSetVectorFail {
							vectorInt16Mock.AssertCalled(t, "SetVectorInt16", ti.fn)
						}

						if !ti.makePushFail {
							for _, item := range ti.mockVector {
								int16Mock.AssertCalled(t, "PushInt16", item)
							}
						}
					}
				}
			})
		}

	})

}


func Test_Decode_VectorUint16(t *testing.T) {
	Convey("Decode VectorUint16 test", t, func() {
		table := []struct {
			fn                      hybrids.FieldNumber
			vector                  interface{}
			mockVector              []uint16
			shouldFail              bool
			makeSetVectorFail    bool
			makePushFail            bool
			shouldTryToCreateVector bool
			name                    string
		}{
			{30, nil, nil, false, false, false, true, "null: should create a vector but don't add items (remember vector and tables can have a null state)"},
			{25, "vector", nil, true, false, false, false, "incorrect underlying type"},
			{70, []interface{}{"nan"}, []uint16{0}, true, false, false, true, "item: incorrect underlying type"},
			{50, []interface{}{1, 2, 3}, []uint16{1, 2, 3}, true, true, false, true, "Should fails when the vector creation operation fails"},
			{50, []interface{}{ uint16(1), uint16(2), uint16(3)}, []uint16{1, 2, 3}, true, false, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{ uint16(1), uint16(2), uint16(3)}, []uint16{1, 2, 3}, false, false, false, true, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				uint16Mock := &mocks.VectorUint16WriterBase{}
				for _, item := range ti.mockVector {
					callItem := uint16Mock.On("PushUint16", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("PushUint16 failed"))
					} else {
						callItem.Return(nil)
					}
				}

				vectorUint16Mock := &mocks.VectorUint16WriterAccessor{}

				call := vectorUint16Mock.On("SetVectorUint16", ti.fn)
				if ti.makeSetVectorFail {
					call.Return(nil, fmt.Errorf("SetVectorUint16 failed"))
				} else {
					call.Return(uint16Mock, nil)
				}



				err := d.decodeVectorUint16("x.path.vector", ti.vector, ti.fn, vectorUint16Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.VectorUint16)
					So(de.OmniqlType, ShouldEqual, "Vector[Uint16]")

				} else {
					So(err, ShouldBeNil)
					if ti.shouldTryToCreateVector {
						if !ti.makeSetVectorFail {
							vectorUint16Mock.AssertCalled(t, "SetVectorUint16", ti.fn)
						}

						if !ti.makePushFail {
							for _, item := range ti.mockVector {
								uint16Mock.AssertCalled(t, "PushUint16", item)
							}
						}
					}
				}
			})
		}

	})

}


func Test_Decode_VectorInt32(t *testing.T) {
	Convey("Decode VectorInt32 test", t, func() {
		table := []struct {
			fn                      hybrids.FieldNumber
			vector                  interface{}
			mockVector              []int32
			shouldFail              bool
			makeSetVectorFail    bool
			makePushFail            bool
			shouldTryToCreateVector bool
			name                    string
		}{
			{30, nil, nil, false, false, false, true, "null: should create a vector but don't add items (remember vector and tables can have a null state)"},
			{25, "vector", nil, true, false, false, false, "incorrect underlying type"},
			{70, []interface{}{"nan"}, []int32{0}, true, false, false, true, "item: incorrect underlying type"},
			{50, []interface{}{1, 2, 3}, []int32{1, 2, 3}, true, true, false, true, "Should fails when the vector creation operation fails"},
			{50, []interface{}{ int32(1), int32(2), int32(3)}, []int32{1, 2, 3}, true, false, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{ int32(1), int32(2), int32(3)}, []int32{1, 2, 3}, false, false, false, true, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				int32Mock := &mocks.VectorInt32WriterBase{}
				for _, item := range ti.mockVector {
					callItem := int32Mock.On("PushInt32", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("PushInt32 failed"))
					} else {
						callItem.Return(nil)
					}
				}

				vectorInt32Mock := &mocks.VectorInt32WriterAccessor{}

				call := vectorInt32Mock.On("SetVectorInt32", ti.fn)
				if ti.makeSetVectorFail {
					call.Return(nil, fmt.Errorf("SetVectorInt32 failed"))
				} else {
					call.Return(int32Mock, nil)
				}



				err := d.decodeVectorInt32("x.path.vector", ti.vector, ti.fn, vectorInt32Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.VectorInt32)
					So(de.OmniqlType, ShouldEqual, "Vector[Int32]")

				} else {
					So(err, ShouldBeNil)
					if ti.shouldTryToCreateVector {
						if !ti.makeSetVectorFail {
							vectorInt32Mock.AssertCalled(t, "SetVectorInt32", ti.fn)
						}

						if !ti.makePushFail {
							for _, item := range ti.mockVector {
								int32Mock.AssertCalled(t, "PushInt32", item)
							}
						}
					}
				}
			})
		}

	})

}


func Test_Decode_VectorUint32(t *testing.T) {
	Convey("Decode VectorUint32 test", t, func() {
		table := []struct {
			fn                      hybrids.FieldNumber
			vector                  interface{}
			mockVector              []uint32
			shouldFail              bool
			makeSetVectorFail    bool
			makePushFail            bool
			shouldTryToCreateVector bool
			name                    string
		}{
			{30, nil, nil, false, false, false, true, "null: should create a vector but don't add items (remember vector and tables can have a null state)"},
			{25, "vector", nil, true, false, false, false, "incorrect underlying type"},
			{70, []interface{}{"nan"}, []uint32{0}, true, false, false, true, "item: incorrect underlying type"},
			{50, []interface{}{1, 2, 3}, []uint32{1, 2, 3}, true, true, false, true, "Should fails when the vector creation operation fails"},
			{50, []interface{}{ uint32(1), uint32(2), uint32(3)}, []uint32{1, 2, 3}, true, false, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{ uint32(1), uint32(2), uint32(3)}, []uint32{1, 2, 3}, false, false, false, true, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				uint32Mock := &mocks.VectorUint32WriterBase{}
				for _, item := range ti.mockVector {
					callItem := uint32Mock.On("PushUint32", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("PushUint32 failed"))
					} else {
						callItem.Return(nil)
					}
				}

				vectorUint32Mock := &mocks.VectorUint32WriterAccessor{}

				call := vectorUint32Mock.On("SetVectorUint32", ti.fn)
				if ti.makeSetVectorFail {
					call.Return(nil, fmt.Errorf("SetVectorUint32 failed"))
				} else {
					call.Return(uint32Mock, nil)
				}



				err := d.decodeVectorUint32("x.path.vector", ti.vector, ti.fn, vectorUint32Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.VectorUint32)
					So(de.OmniqlType, ShouldEqual, "Vector[Uint32]")

				} else {
					So(err, ShouldBeNil)
					if ti.shouldTryToCreateVector {
						if !ti.makeSetVectorFail {
							vectorUint32Mock.AssertCalled(t, "SetVectorUint32", ti.fn)
						}

						if !ti.makePushFail {
							for _, item := range ti.mockVector {
								uint32Mock.AssertCalled(t, "PushUint32", item)
							}
						}
					}
				}
			})
		}

	})

}


func Test_Decode_VectorFloat32(t *testing.T) {
	Convey("Decode VectorFloat32 test", t, func() {
		table := []struct {
			fn                      hybrids.FieldNumber
			vector                  interface{}
			mockVector              []float32
			shouldFail              bool
			makeSetVectorFail    bool
			makePushFail            bool
			shouldTryToCreateVector bool
			name                    string
		}{
			{30, nil, nil, false, false, false, true, "null: should create a vector but don't add items (remember vector and tables can have a null state)"},
			{25, "vector", nil, true, false, false, false, "incorrect underlying type"},
			{70, []interface{}{"nan"}, []float32{0}, true, false, false, true, "item: incorrect underlying type"},
			{50, []interface{}{1, 2, 3}, []float32{1, 2, 3}, true, true, false, true, "Should fails when the vector creation operation fails"},
			{50, []interface{}{ float32(1), float32(2), float32(3)}, []float32{1, 2, 3}, true, false, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{ float32(1), float32(2), float32(3)}, []float32{1, 2, 3}, false, false, false, true, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				float32Mock := &mocks.VectorFloat32WriterBase{}
				for _, item := range ti.mockVector {
					callItem := float32Mock.On("PushFloat32", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("PushFloat32 failed"))
					} else {
						callItem.Return(nil)
					}
				}

				vectorFloat32Mock := &mocks.VectorFloat32WriterAccessor{}

				call := vectorFloat32Mock.On("SetVectorFloat32", ti.fn)
				if ti.makeSetVectorFail {
					call.Return(nil, fmt.Errorf("SetVectorFloat32 failed"))
				} else {
					call.Return(float32Mock, nil)
				}



				err := d.decodeVectorFloat32("x.path.vector", ti.vector, ti.fn, vectorFloat32Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.VectorFloat32)
					So(de.OmniqlType, ShouldEqual, "Vector[Float32]")

				} else {
					So(err, ShouldBeNil)
					if ti.shouldTryToCreateVector {
						if !ti.makeSetVectorFail {
							vectorFloat32Mock.AssertCalled(t, "SetVectorFloat32", ti.fn)
						}

						if !ti.makePushFail {
							for _, item := range ti.mockVector {
								float32Mock.AssertCalled(t, "PushFloat32", item)
							}
						}
					}
				}
			})
		}

	})

}


func Test_Decode_VectorFloat64(t *testing.T) {
	Convey("Decode VectorFloat64 test", t, func() {
		table := []struct {
			fn                      hybrids.FieldNumber
			vector                  interface{}
			mockVector              []float64
			shouldFail              bool
			makeSetVectorFail    bool
			makePushFail            bool
			shouldTryToCreateVector bool
			name                    string
		}{
			{30, nil, nil, false, false, false, true, "null: should create a vector but don't add items (remember vector and tables can have a null state)"},
			{25, "vector", nil, true, false, false, false, "incorrect underlying type"},
			{70, []interface{}{"nan"}, []float64{0}, true, false, false, true, "item: incorrect underlying type"},
			{50, []interface{}{1, 2, 3}, []float64{1, 2, 3}, true, true, false, true, "Should fails when the vector creation operation fails"},
			{50, []interface{}{ float64(1), float64(2), float64(3)}, []float64{1, 2, 3}, true, false, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{ float64(1), float64(2), float64(3)}, []float64{1, 2, 3}, false, false, false, true, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				float64Mock := &mocks.VectorFloat64WriterBase{}
				for _, item := range ti.mockVector {
					callItem := float64Mock.On("PushFloat64", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("PushFloat64 failed"))
					} else {
						callItem.Return(nil)
					}
				}

				vectorFloat64Mock := &mocks.VectorFloat64WriterAccessor{}

				call := vectorFloat64Mock.On("SetVectorFloat64", ti.fn)
				if ti.makeSetVectorFail {
					call.Return(nil, fmt.Errorf("SetVectorFloat64 failed"))
				} else {
					call.Return(float64Mock, nil)
				}



				err := d.decodeVectorFloat64("x.path.vector", ti.vector, ti.fn, vectorFloat64Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.VectorFloat64)
					So(de.OmniqlType, ShouldEqual, "Vector[Float64]")

				} else {
					So(err, ShouldBeNil)
					if ti.shouldTryToCreateVector {
						if !ti.makeSetVectorFail {
							vectorFloat64Mock.AssertCalled(t, "SetVectorFloat64", ti.fn)
						}

						if !ti.makePushFail {
							for _, item := range ti.mockVector {
								float64Mock.AssertCalled(t, "PushFloat64", item)
							}
						}
					}
				}
			})
		}

	})

}


func Test_Decode_VectorInt64(t *testing.T) {
	Convey("Decode VectorInt64 test", t, func() {
		table := []struct {
			fn                      hybrids.FieldNumber
			vector                  interface{}
			mockVector              []int64
			shouldFail              bool
			makeSetVectorFail    bool
			makePushFail            bool
			shouldTryToCreateVector bool
			name                    string
		}{
			{30, nil, nil, false, false, false, true, "null: should create a vector but don't add items (remember vector and tables can have a null state)"},
			{25, "vector", nil, true, false, false, false, "incorrect underlying type"},
			{70, []interface{}{"nan"}, []int64{0}, true, false, false, true, "item: incorrect underlying type"},
			{50, []interface{}{1, 2, 3}, []int64{1, 2, 3}, true, true, false, true, "Should fails when the vector creation operation fails"},
			{50, []interface{}{ int64(1), int64(2), int64(3)}, []int64{1, 2, 3}, true, false, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{ int64(1), int64(2), int64(3)}, []int64{1, 2, 3}, false, false, false, true, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				int64Mock := &mocks.VectorInt64WriterBase{}
				for _, item := range ti.mockVector {
					callItem := int64Mock.On("PushInt64", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("PushInt64 failed"))
					} else {
						callItem.Return(nil)
					}
				}

				vectorInt64Mock := &mocks.VectorInt64WriterAccessor{}

				call := vectorInt64Mock.On("SetVectorInt64", ti.fn)
				if ti.makeSetVectorFail {
					call.Return(nil, fmt.Errorf("SetVectorInt64 failed"))
				} else {
					call.Return(int64Mock, nil)
				}



				err := d.decodeVectorInt64("x.path.vector", ti.vector, ti.fn, vectorInt64Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.VectorInt64)
					So(de.OmniqlType, ShouldEqual, "Vector[Int64]")

				} else {
					So(err, ShouldBeNil)
					if ti.shouldTryToCreateVector {
						if !ti.makeSetVectorFail {
							vectorInt64Mock.AssertCalled(t, "SetVectorInt64", ti.fn)
						}

						if !ti.makePushFail {
							for _, item := range ti.mockVector {
								int64Mock.AssertCalled(t, "PushInt64", item)
							}
						}
					}
				}
			})
		}

	})

}


func Test_Decode_VectorUint64(t *testing.T) {
	Convey("Decode VectorUint64 test", t, func() {
		table := []struct {
			fn                      hybrids.FieldNumber
			vector                  interface{}
			mockVector              []uint64
			shouldFail              bool
			makeSetVectorFail    bool
			makePushFail            bool
			shouldTryToCreateVector bool
			name                    string
		}{
			{30, nil, nil, false, false, false, true, "null: should create a vector but don't add items (remember vector and tables can have a null state)"},
			{25, "vector", nil, true, false, false, false, "incorrect underlying type"},
			{70, []interface{}{"nan"}, []uint64{0}, true, false, false, true, "item: incorrect underlying type"},
			{50, []interface{}{1, 2, 3}, []uint64{1, 2, 3}, true, true, false, true, "Should fails when the vector creation operation fails"},
			{50, []interface{}{ uint64(1), uint64(2), uint64(3)}, []uint64{1, 2, 3}, true, false, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{ uint64(1), uint64(2), uint64(3)}, []uint64{1, 2, 3}, false, false, false, true, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				uint64Mock := &mocks.VectorUint64WriterBase{}
				for _, item := range ti.mockVector {
					callItem := uint64Mock.On("PushUint64", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("PushUint64 failed"))
					} else {
						callItem.Return(nil)
					}
				}

				vectorUint64Mock := &mocks.VectorUint64WriterAccessor{}

				call := vectorUint64Mock.On("SetVectorUint64", ti.fn)
				if ti.makeSetVectorFail {
					call.Return(nil, fmt.Errorf("SetVectorUint64 failed"))
				} else {
					call.Return(uint64Mock, nil)
				}



				err := d.decodeVectorUint64("x.path.vector", ti.vector, ti.fn, vectorUint64Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.VectorUint64)
					So(de.OmniqlType, ShouldEqual, "Vector[Uint64]")

				} else {
					So(err, ShouldBeNil)
					if ti.shouldTryToCreateVector {
						if !ti.makeSetVectorFail {
							vectorUint64Mock.AssertCalled(t, "SetVectorUint64", ti.fn)
						}

						if !ti.makePushFail {
							for _, item := range ti.mockVector {
								uint64Mock.AssertCalled(t, "PushUint64", item)
							}
						}
					}
				}
			})
		}

	})

}

