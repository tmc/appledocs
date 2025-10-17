
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [NibConnector] class.
var NibConnectorClass _NibConnectorClass

func init() {
	NibConnectorClass = _NibConnectorClass{objc.GetClass("NSNibConnector")}
}

type _NibConnectorClass struct {
	objc.Class
}

// An interface definition for the [NibConnector] class.
type INibConnector interface {
	ID() objc.ID
}

type NibConnector struct {
	id objc.ID
}

func NibConnectorFrom(ptr unsafe.Pointer) NibConnector {
	return NibConnector{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ NibConnector) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _NibConnectorClass) Alloc() NibConnector {
	rv := objc.Send[NibConnector](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _NibConnectorClass) New() NibConnector {
	rv := objc.Send[NibConnector](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewNibConnector creates and returns a new initialized instance.
func NewNibConnector() NibConnector {
	return NibConnectorClass.New()
}

// Init initializes the instance.
func (n_ NibConnector) Init() NibConnector {
	rv := objc.Send[NibConnector](n_.ID(), selInit)
	return rv
}
