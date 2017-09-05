package omarshaller

import (
	"fmt"
	"github.com/nebtex/hybrids/golang/hybrids"
)

//DecodeError contains the data to fully identify  an error produced by the decoder
type DecodeError struct {
	// field full path
	Path string         `json:"path,omitempty"`
	// field underlying type
	HybridType  hybrids.Types  `json:"hybrid_type,omitempty"`
	OmniqlType  string  `json:"omniql_type,omitempty"`
	Application string  `json:"application,omitempty"`
	ErrorMsg    string  `json:"error,omitempty"`
}

func (d *DecodeError) Error() string {
	return fmt.Sprintf("#Application: %s, #Path: %s, #HybridType: %s, #OmniqlType: %s, Error: %s", d.Application, d.Path, d.HybridType.String(), d.OmniqlType, d.ErrorMsg)
}
