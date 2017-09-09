package encoder

import "github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
import "github.com/nebtex/hybrids/golang/hybrids"

func (e *Encoder) encodeTable(path string, out map[string]interface{}, fieldName string, ot oreflection.OType, tr hybrids.TableReader) (err error) {
	var i hybrids.FieldNumber

	if tr == nil {
		return
	}

	out = map[string]interface{}{}
	table := map[string]interface{}{}
	for i = 0; i < ot.FieldCount(); i++ {
		fieldName, fieldType, _ := ot.LookupFields().ByNumber(i)
		if fieldType.HybridType().IsScalar() {
			e.encodeScalar(table, fieldName, i, fieldType, tr)

		} else if fieldType.HybridType().IsVectorScalar() {
			e.encodeVectorScalar(table, fieldName, i, fieldType, tr)
		}

		switch fieldType.HybridType() {
		case hybrids.Struct:
			e.encodeStruct(out, fieldName, i, fieldType, tr)
		case hybrids.VectorStruct:
			e.encodeVectorStruct(out, fieldName, i, fieldType, tr)
		case hybrids.Table:
			e.encodeTable(out, fieldName, i, fieldType, tr)
		case hybrids.VectorTable:
			e.encodeVectorTable(out, fieldName, i, fieldType, tr)

		}
	}

	return
}
