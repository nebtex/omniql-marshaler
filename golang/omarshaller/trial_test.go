package omarshaller

import (
	"testing"
	"github.com/nebtex/hybrids/golang/hybrids/mocks"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
)

func Test_Vector_Trial(t *testing.T) {
	Convey("Vector String native", t, func() {
		table := []struct {
			fn                      hybrids.FieldNumber
			vector                  interface{}
			mockVector              []float32
			shouldFail              bool
			makeUpsertVectorFail    bool
			makePushFail            bool
			shouldTryToCreateVector bool
			name                    string
		}{
			{30, nil, nil, false, false, false, true, "null: should create a vector but don't add items (remember vector and tables can have a null state)"},
			{25, "vector", nil, true, false, false, false, "incorrect underlying type"},
			{70, []interface{}{"nan"}, []float32{0}, true, false, false, true, "item: incorrect underlying type"},
			{50, []interface{}{1, 2, 3}, []float32{1, 2, 3}, true, true, false, true, "Should fails when the vector creation operation fails"},
			{50, []interface{}{float32(1), float32(2), float32(3)}, []float32{1, 2, 3}, true, false, true, true, "Should fails when the push operation fails"},
			{50, []interface{}{float32(1), float32(2), float32(3)}, []float32{1, 2, 3}, false, false, false, true, "Valid input, all should be ok"},


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

				call := vectorFloat32Mock.On("UpsertVectorFloat32", ti.fn)
				if ti.makeUpsertVectorFail {
					call.Return(nil, fmt.Errorf("UpsertVectorFloat32 failed"))
				} else {
					call.Return(float32Mock, nil)
				}



				err := d.decodeVectorFloat32("x.path.vector", ti.vector, ti.fn, vectorFloat32Mock)
				if ti.shouldFail {
					t.Log(err)
					So(err, ShouldNotBeNil)
					de, _ := err.(*DecodeError)
					So(de.Application, ShouldEqual, "test")
					So(de.HybridType, ShouldEqual, "VectorFloat32")
					So(de.OmniqlType, ShouldEqual, "Vector")
					So(de.OmniqlItems, ShouldEqual, "Float32")

				} else {
					So(err, ShouldBeNil)
					if ti.shouldTryToCreateVector {
						if !ti.makeUpsertVectorFail {
							vectorFloat32Mock.AssertCalled(t, "UpsertVectorFloat32", ti.fn)
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
