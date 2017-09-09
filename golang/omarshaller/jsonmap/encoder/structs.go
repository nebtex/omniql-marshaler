package encoder

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
)

func (e *Encoder) encodeStruct(path string, out map[string]interface{}, fieldName string, ot oreflection.OType, fn hybrids.FieldNumber, tr hybrids.StructReaderAccessor) {
	var sr hybrids.ScalarReader
	var i hybrids.FieldNumber

	sr = tr.Struct(fn)
	if sr == nil {
		return
	}

	structObject := map[string]interface{}{}

	for i = 0; i < ot.FieldCount(); i++ {
		structFieldName, structFieldType, _ := ot.LookupFields().ByNumber(i)
		e.encodeScalar(structObject, structFieldName, i, structFieldType, tr)
		_, ok := structObject[structFieldName]
		if !ok {
			structObject[structFieldName] = float64(0)
		}
	}

	out[fieldName] = structObject
	return
}

func (e *Encoder) encodeVectorStruct(path string, out map[string]interface{}, fieldName string, ot oreflection.OType, fn hybrids.FieldNumber, vsra hybrids.VectorStructReaderAccessor) {
	var vs hybrids.VectorStructReader
	var field hybrids.FieldNumber
	var ok bool
	var sr hybrids.ScalarReader
	vs, ok = vsra.VectorStruct(fn)

	if !ok {
		return
	}

	if vs.Len() == 0 {
		out[fieldName] = nil
		return
	}

	itemsType := ot.Items()
	r := []map[string]interface{}{}

	for i := 0; i < vs.Len(); i++ {
		sr = vs.Get(i)
		structObject := map[string]interface{}{}

		for field = 0; field < itemsType.FieldCount(); field++ {
			structFieldName, structFieldType, _ := ot.LookupFields().ByNumber(field)
			e.encodeScalar(structObject, structFieldName, field, structFieldType, sr)
			_, ok := structObject[structFieldName]
			if !ok {
				structObject[structFieldName] = float64(0)
			}
		}

		r = append(r, structObject)
	}

	out[fieldName] = r
	return
}
