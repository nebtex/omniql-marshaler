package marshaller

//go:generate easyjson -all errors.go

//DecodeError contains the data to fully identify  an error produced by the decoder
type DecodeError struct {
	// field full path
	Path string         `json:"path,omitempty"`
	// field underlying type
	HybridType  string  `json:"hybrid_type,omitempty"`
	OmniqlType  string  `json:"omniql_type,omitempty"`
	OmniqlItems string  `json:"omniql_items,omitempty"`
	Application string  `json:"application,omitempty"`
	ErrorMsg    string  `json:"error,omitempty"`
}

func (d *DecodeError) Error() string {
	json, _ := d.MarshalJSON()
	return string(json)
}
