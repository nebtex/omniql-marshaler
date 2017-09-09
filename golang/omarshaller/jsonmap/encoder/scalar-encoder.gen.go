// Code generated for marshaling/unmarshaling. DO NOT EDIT.
package encoder

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"strconv"
	"github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
)


func (e *Encoder) encodeBoolean(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.BooleanReaderAccessor) {
	v, ok := tr.Boolean(fn)
	if !ok {
		//field is missing
		return
	}
	out[field_name] = v
}


func (e *Encoder) encodeVectorBoolean(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.VectorBooleanReaderAccessor) {
	var v bool
	vrb, ok := tr.VectorBoolean(fn)

	if !ok {
		return
	}

	if vrb == nil {
		out[field_name] = nil
		return
	}

	r := make([]interface{}, 0, vrb.Len())

	for i := 0; i < vrb.Len(); i++ {
		v = vrb.Get(i)
		r = append(r, v)
	}

	out[field_name] = r

	return
}


func (e *Encoder) encodeInt8(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.Int8ReaderAccessor) {
	v, ok := tr.Int8(fn)
	if !ok {
		//field is missing
		return
	}
	out[field_name] = float64(v)
}


func (e *Encoder) encodeVectorInt8(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.VectorInt8ReaderAccessor) {
	var v int8
	vrb, ok := tr.VectorInt8(fn)

	if !ok {
		return
	}

	if vrb == nil {
		out[field_name] = nil
		return
	}

	r := make([]interface{}, 0, vrb.Len())

	for i := 0; i < vrb.Len(); i++ {
		v = vrb.Get(i)
		r = append(r, float64(v))
	}

	out[field_name] = r

	return
}


func (e *Encoder) encodeUint8(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.Uint8ReaderAccessor) {
	v, ok := tr.Uint8(fn)
	if !ok {
		//field is missing
		return
	}
	out[field_name] = float64(v)
}


func (e *Encoder) encodeVectorUint8(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.VectorUint8ReaderAccessor) {
	var v uint8
	vrb, ok := tr.VectorUint8(fn)

	if !ok {
		return
	}

	if vrb == nil {
		out[field_name] = nil
		return
	}

	r := make([]interface{}, 0, vrb.Len())

	for i := 0; i < vrb.Len(); i++ {
		v = vrb.Get(i)
		r = append(r, float64(v))
	}

	out[field_name] = r

	return
}


func (e *Encoder) encodeInt16(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.Int16ReaderAccessor) {
	v, ok := tr.Int16(fn)
	if !ok {
		//field is missing
		return
	}
	out[field_name] = float64(v)
}


func (e *Encoder) encodeVectorInt16(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.VectorInt16ReaderAccessor) {
	var v int16
	vrb, ok := tr.VectorInt16(fn)

	if !ok {
		return
	}

	if vrb == nil {
		out[field_name] = nil
		return
	}

	r := make([]interface{}, 0, vrb.Len())

	for i := 0; i < vrb.Len(); i++ {
		v = vrb.Get(i)
		r = append(r, float64(v))
	}

	out[field_name] = r

	return
}


func (e *Encoder) encodeUint16(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.Uint16ReaderAccessor) {
	v, ok := tr.Uint16(fn)
	if !ok {
		//field is missing
		return
	}
	out[field_name] = float64(v)
}


func (e *Encoder) encodeVectorUint16(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.VectorUint16ReaderAccessor) {
	var v uint16
	vrb, ok := tr.VectorUint16(fn)

	if !ok {
		return
	}

	if vrb == nil {
		out[field_name] = nil
		return
	}

	r := make([]interface{}, 0, vrb.Len())

	for i := 0; i < vrb.Len(); i++ {
		v = vrb.Get(i)
		r = append(r, float64(v))
	}

	out[field_name] = r

	return
}


func (e *Encoder) encodeInt32(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.Int32ReaderAccessor) {
	v, ok := tr.Int32(fn)
	if !ok {
		//field is missing
		return
	}
	out[field_name] = float64(v)
}


func (e *Encoder) encodeVectorInt32(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.VectorInt32ReaderAccessor) {
	var v int32
	vrb, ok := tr.VectorInt32(fn)

	if !ok {
		return
	}

	if vrb == nil {
		out[field_name] = nil
		return
	}

	r := make([]interface{}, 0, vrb.Len())

	for i := 0; i < vrb.Len(); i++ {
		v = vrb.Get(i)
		r = append(r, float64(v))
	}

	out[field_name] = r

	return
}


func (e *Encoder) encodeUint32(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.Uint32ReaderAccessor) {
	v, ok := tr.Uint32(fn)
	if !ok {
		//field is missing
		return
	}
	out[field_name] = float64(v)
}


func (e *Encoder) encodeVectorUint32(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.VectorUint32ReaderAccessor) {
	var v uint32
	vrb, ok := tr.VectorUint32(fn)

	if !ok {
		return
	}

	if vrb == nil {
		out[field_name] = nil
		return
	}

	r := make([]interface{}, 0, vrb.Len())

	for i := 0; i < vrb.Len(); i++ {
		v = vrb.Get(i)
		r = append(r, float64(v))
	}

	out[field_name] = r

	return
}


func (e *Encoder) encodeInt64(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.Int64ReaderAccessor) {
	v, ok := tr.Int64(fn)
	if !ok {
		//field is missing
		return
	}
	out[field_name] = strconv.FormatInt(v, 10)
}


func (e *Encoder) encodeVectorInt64(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.VectorInt64ReaderAccessor) {
	var v int64
	vrb, ok := tr.VectorInt64(fn)

	if !ok {
		return
	}

	if vrb == nil {
		out[field_name] = nil
		return
	}

	r := make([]interface{}, 0, vrb.Len())

	for i := 0; i < vrb.Len(); i++ {
		v = vrb.Get(i)
		r = append(r, strconv.FormatInt(v, 10))
	}

	out[field_name] = r

	return
}


func (e *Encoder) encodeUint64(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.Uint64ReaderAccessor) {
	v, ok := tr.Uint64(fn)
	if !ok {
		//field is missing
		return
	}
	out[field_name] = strconv.FormatUint(v, 10)
}


func (e *Encoder) encodeVectorUint64(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.VectorUint64ReaderAccessor) {
	var v uint64
	vrb, ok := tr.VectorUint64(fn)

	if !ok {
		return
	}

	if vrb == nil {
		out[field_name] = nil
		return
	}

	r := make([]interface{}, 0, vrb.Len())

	for i := 0; i < vrb.Len(); i++ {
		v = vrb.Get(i)
		r = append(r, strconv.FormatUint(v, 10))
	}

	out[field_name] = r

	return
}


func (e *Encoder) encodeFloat32(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.Float32ReaderAccessor) {
	v, ok := tr.Float32(fn)
	if !ok {
		//field is missing
		return
	}
	out[field_name] = float64(v)
}


func (e *Encoder) encodeVectorFloat32(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.VectorFloat32ReaderAccessor) {
	var v float32
	vrb, ok := tr.VectorFloat32(fn)

	if !ok {
		return
	}

	if vrb == nil {
		out[field_name] = nil
		return
	}

	r := make([]interface{}, 0, vrb.Len())

	for i := 0; i < vrb.Len(); i++ {
		v = vrb.Get(i)
		r = append(r, float64(v))
	}

	out[field_name] = r

	return
}


func (e *Encoder) encodeFloat64(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.Float64ReaderAccessor) {
	v, ok := tr.Float64(fn)
	if !ok {
		//field is missing
		return
	}
	out[field_name] = v
}


func (e *Encoder) encodeVectorFloat64(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tr hybrids.VectorFloat64ReaderAccessor) {
	var v float64
	vrb, ok := tr.VectorFloat64(fn)

	if !ok {
		return
	}

	if vrb == nil {
		out[field_name] = nil
		return
	}

	r := make([]interface{}, 0, vrb.Len())

	for i := 0; i < vrb.Len(); i++ {
		v = vrb.Get(i)
		r = append(r, v)
	}

	out[field_name] = r

	return
}


func (e *Encoder) encodeScalar(out map[string]interface{}, field_name string, fn hybrids.FieldNumber, tp oreflection.OType, tr hybrids.ScalarReader){
	switch tp{

	case hybrids.Boolean:
		e.encodeBoolean(out, field_name, fn, tr)

	case hybrids.Int8:
		e.encodeInt8(out, field_name, fn, tr)

	case hybrids.Uint8:
		e.encodeUint8(out, field_name, fn, tr)

	case hybrids.Int16:
		e.encodeInt16(out, field_name, fn, tr)

	case hybrids.Uint16:
		e.encodeUint16(out, field_name, fn, tr)

	case hybrids.Int32:
		e.encodeInt32(out, field_name, fn, tr)

	case hybrids.Uint32:
		e.encodeUint32(out, field_name, fn, tr)

	case hybrids.Int64:
		e.encodeInt64(out, field_name, fn, tr)

	case hybrids.Uint64:
		e.encodeUint64(out, field_name, fn, tr)

	case hybrids.Float32:
		e.encodeFloat32(out, field_name, fn, tr)

	case hybrids.Float64:
		e.encodeFloat64(out, field_name, fn, tr)

    }
	return
}
