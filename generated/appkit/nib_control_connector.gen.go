
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [NibControlConnector] class.
var NibControlConnectorClass _NibControlConnectorClass

func init() {
	NibControlConnectorClass = _NibControlConnectorClass{objc.GetClass("NSNibControlConnector")}
}

type _NibControlConnectorClass struct {
	objc.Class
}

// An interface definition for the [NibControlConnector] class.
type INibControlConnector interface {
	ID() objc.ID
}

type NibControlConnector struct {
	id objc.ID
}

func NibControlConnectorFrom(ptr unsafe.Pointer) NibControlConnector {
	return NibControlConnector{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ NibControlConnector) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _NibControlConnectorClass) Alloc() NibControlConnector {
	rv := objc.Send[NibControlConnector](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _NibControlConnectorClass) New() NibControlConnector {
	rv := objc.Send[NibControlConnector](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewNibControlConnector creates and returns a new initialized instance.
func NewNibControlConnector() NibControlConnector {
	return NibControlConnectorClass.New()
}

// Init initializes the instance.
func (n_ NibControlConnector) Init() NibControlConnector {
	rv := objc.Send[NibControlConnector](n_.ID(), selInit)
	return rv
}
