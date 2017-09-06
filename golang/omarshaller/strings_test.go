package omarshaller

import (
	"testing"
	"github.com/nebtex/hybrids/golang/hybrids/mocks"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
	"github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
)

func Test_DecoderString(t *testing.T) {
	Convey("String Tests", t, func() {
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
	})
}

func Test_DecoderByte(t *testing.T) {
	Convey("Byte Tests", t, func() {
		table := []struct {
			fn             hybrids.FieldNumber
			value          interface{}
			mockByte       []byte
			shouldFail     bool
			shouldMockFail bool
			name           string
		}{
			{1, []byte("jupiter"), []byte("jupiter"), false, false, "native_type"},
			{1, "anVwaXRlcg==", []byte("jupiter"), false, false, "base64 string"},
			{25, 54545, nil, true, false, "incorrect underlying type"},
			{10, []byte("sun"), []byte("sun"), true, true, "hybrid set method fail"},
			{10, nil, nil, false, false, "null should not return error, and should call the set method"},
			{10, []byte(""), []byte{}, false, false, "empty string should  call the set method"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				stringMock := &mocks.ByteWriterAccessor{}
				call := stringMock.On("SetByte", ti.fn, ti.mockByte)
				if ti.shouldMockFail {
					call.Return(fmt.Errorf("SetByte failed"))
				} else {
					call.Return(nil)
				}

				err := d.decodeByte("x.path", ti.value, ti.fn, stringMock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.Path, ShouldEqual, "x.path")
					So(de.HybridType, ShouldEqual, hybrids.Byte)
					So(de.OmniqlType, ShouldEqual, "Byte")

				} else {
					So(err, ShouldBeNil)

					if !ti.shouldFail {
						stringMock.AssertCalled(t, "SetByte", ti.fn, ti.mockByte)
					}

				}
			})
		}
	})
}

func Test_DecoderResourceID(t *testing.T) {
	Convey("ResourceID Tests", t, func() {
		table := []struct {
			fn             hybrids.FieldNumber
			value          interface{}
			mockResourceID []byte
			shouldFail     bool
			shouldMockFail bool
			name           string
		}{
			{1, []byte("jupiter"), []byte("jupiter"), false, false, "native_type"},
			{1, "anVwaXRlcg==", []byte("jupiter"), false, false, "base64 string"},
			{25, 54545, nil, true, false, "incorrect underlying type"},
			{10, []byte("sun"), []byte("sun"), true, true, "hybrid set method fail"},
			{10, nil, nil, false, false, "null should not return error, and should call the set method"},
			{10, []byte(""), []byte{}, false, false, "empty string should  call the set method"},

		}
		ot := &oreflection.OType{Id: "Resource/Resource"}

		d := &Decoder{application: "test", ResourceIDType: hybrids.ResourceIDTypeString, ResourceKindType: hybrids.ResourceKindTypeUnsignedShort}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {
				rMock := &mocks.ResourceIDWriterAccessor{}
				call := rMock.On("SetResourceID", ti.fn, hybrids.ResourceKindTypeUnsignedShort, hybrids.ResourceIDTypeString, ti.mockResourceID)
				if ti.shouldMockFail {
					call.Return(fmt.Errorf("SetResourceID failed"))
				} else {
					call.Return(nil)
				}

				err := d.decodeResourceID("x.path", ti.value, ti.fn, ot, rMock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.Path, ShouldEqual, "x.path")
					So(de.HybridType, ShouldEqual, hybrids.ResourceID)
					So(de.OmniqlType, ShouldEqual, "Resource/Resource")

				} else {
					So(err, ShouldBeNil)

					if !ti.shouldFail {
						rMock.AssertCalled(t, "SetResourceID", ti.fn, hybrids.ResourceKindTypeUnsignedShort, hybrids.ResourceIDTypeString, ti.mockResourceID)
					}

				}
			})
		}
	})
}

func Test_Decode_VectorString(t *testing.T) {
	Convey("Decode VectorString test", t, func() {
		table := []struct {
			fn           hybrids.FieldNumber
			vector       interface{}
			mockVector   []string
			shouldFail   bool
			makePushFail bool
			name         string
		}{
			{30, nil, nil, false, false, "null: should not do anything"},
			{25, "vector", nil, true, false, "incorrect underlying type"},
			{70, []interface{}{565}, []string{""}, true, false, "item: incorrect underlying type"},
			{50, []interface{}{"1", "2", "3"}, []string{"1", "2", "3"}, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{"1", "2", "3"}, []string{"1", "2", "3"}, false, false, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				stringMock := &mocks.VectorStringWriter{}
				for _, item := range ti.mockVector {
					callItem := stringMock.On("PushString", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("PushString failed"))
					} else {
						callItem.Return(nil)
					}
				}

				err := d.decodeVectorString("x.path.vector", ti.vector, stringMock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.VectorString)
					So(de.OmniqlType, ShouldEqual, "Vector[String]")

				} else {
					So(err, ShouldBeNil)
					if !ti.makePushFail {
						for _, item := range ti.mockVector {
							stringMock.AssertCalled(t, "PushString", item)
						}
					}
				}
			})
		}

	})

}

func Test_Decode_VectorByte(t *testing.T) {
	Convey("Decode VectorByte test", t, func() {
		table := []struct {
			fn           hybrids.FieldNumber
			vector       interface{}
			mockVector   [][]byte
			shouldFail   bool
			makePushFail bool
			name         string
		}{
			{30, nil, nil, false, false, "null: should not do anything"},
			{25, "vector", nil, true, false, "incorrect underlying type"},
			{70, []interface{}{565}, [][]byte{}, true, false, "item: incorrect underlying type"},
			{50, []interface{}{[]byte("1"), []byte("2"), []byte("3")}, [][]byte{[]byte("1"), []byte("2"), []byte("3")}, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{[]byte("1"), []byte("2"), []byte("3")}, [][]byte{[]byte("1"), []byte("2"), []byte("3")}, false, false, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test"}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				stringMock := &mocks.VectorByteWriter{}
				for _, item := range ti.mockVector {
					callItem := stringMock.On("PushByte", item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("PushByte failed"))
					} else {
						callItem.Return(nil)
					}
				}

				err := d.decodeVectorByte("x.path.vector", ti.vector, stringMock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.VectorByte)
					So(de.OmniqlType, ShouldEqual, "Vector[Byte]")

				} else {
					So(err, ShouldBeNil)
					if !ti.makePushFail {
						for _, item := range ti.mockVector {
							stringMock.AssertCalled(t, "PushByte", item)
						}
					}

				}
			})
		}

	})

}

func Test_Decode_VectorResourceID(t *testing.T) {
	Convey("Decode VectorResourceID test", t, func() {
		table := []struct {
			fn           hybrids.FieldNumber
			vector       interface{}
			mockVector   [][]byte
			shouldFail   bool
			makePushFail bool
			name         string
		}{
			{30, nil, nil, false, false, "null: should not do anything"},
			{25, "vector", nil, true, false, "incorrect underlying type"},
			{70, []interface{}{565}, [][]byte{}, true, false, "item: incorrect underlying type"},
			{50, []interface{}{[]byte("1"), []byte("2"), []byte("3")}, [][]byte{[]byte("1"), []byte("2"), []byte("3")}, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{[]byte("1"), []byte("2"), []byte("3")}, [][]byte{[]byte("1"), []byte("2"), []byte("3")}, false, false, "Valid input, all should be ok"},

		}

		d := &Decoder{application: "test", ResourceIDType: hybrids.ResourceIDTypeString, ResourceKindType: hybrids.ResourceKindTypeUnsignedShort}
		ot := &oreflection.OType{Id: "Resource/Resource", Items: &oreflection.OType{Id: "Resource/Resource"}}

		for _, ti := range table {
			Convey(fmt.Sprintf("Test: %s", ti.name), func() {

				stringMock := &mocks.VectorResourceIDWriter{}
				for _, item := range ti.mockVector {
					callItem := stringMock.On("PushResourceID", hybrids.ResourceKindTypeUnsignedShort, hybrids.ResourceIDTypeString, item)
					if ti.makePushFail {
						callItem.Return(fmt.Errorf("PushResourceID failed"))
					} else {
						callItem.Return(nil)
					}
				}

				err := d.decodeVectorResourceID("x.path.vector", ti.vector, ot, stringMock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, hybrids.VectorResourceID)
					So(de.OmniqlType, ShouldEqual, "Resource/Resource")

				} else {
					So(err, ShouldBeNil)
					if !ti.makePushFail {
						for _, item := range ti.mockVector {
							stringMock.AssertCalled(t, "PushResourceID", hybrids.ResourceKindTypeUnsignedShort, hybrids.ResourceIDTypeString, item)
						}
					}

				}
			})
		}

	})

}
