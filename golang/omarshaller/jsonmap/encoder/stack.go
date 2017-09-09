package encoder

import (
	"github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
	"github.com/nebtex/hybrids/golang/hybrids"
)

func (e *Encoder) encodeStack(application string, out map[string]interface{}, vtr hybrids.StackReader) (err error) {
	var table hybrids.TableReader
	var rid []byte
	var sRid string
	var ot oreflection.OType
	var ok bool

	vr := vtr.VectorResource()

	for i := 0; i < vr.Len(); i++ {
		rid, table, err = vr.Get(i, e.ResourceIDType)
		if err != nil {
			return
		}

		if table == nil {
			continue
		}

		ot, ok = e.reflect.LookupResourceByID(application, rid)
		if !ok {



		}
		sRid = string(rid)
		err = e.encodeTable(sRid, out, sRid, ot, table)
		if err != nil {
			return
		}
	}
	return
}
