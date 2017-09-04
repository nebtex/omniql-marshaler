package omarshaller

import "fmt"

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
	return fmt.Sprintf("#Application: %s, #Path: %s, #HybridType: %s, #OmniqlType: %s, #OmniqlItems:[%s], Error: %s", d.Application, d.Path, d.HybridType, d.OmniqlType, d.OmniqlItems, d.ErrorMsg)
}
