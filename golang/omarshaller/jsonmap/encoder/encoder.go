package encoder

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"github.com/nebtex/omnibuff/tools/golang/tools/oreflection"
	"go.uber.org/zap"
)

//go:generate go run scalar-generator.go

type Encoder struct {
	application      string
	ResourceIDType   hybrids.ResourceIDType
	//reflection  hybrids.SimpleStore
	reflect oreflection.ReflectStore
	zap *zap.Logger
}

