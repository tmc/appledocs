
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [NibOutletConnector] class.
var NibOutletConnectorClass _NibOutletConnectorClass

func init() {
	NibOutletConnectorClass = _NibOutletConnectorClass{objc.GetClass("NSNibOutletConnector")}
}

type _NibOutletConnectorClass struct {
	objc.Class
}

// An interface definition for the [NibOutletConnector] class.
type INibOutletConnector interface {
	ID() objc.ID
}

type NibOutletConnector struct {
	id objc.ID
}

func NibOutletConnectorFrom(ptr unsafe.Pointer) NibOutletConnector {
	return NibOutletConnector{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ NibOutletConnector) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _NibOutletConnectorClass) Alloc() NibOutletConnector {
	rv := objc.Send[NibOutletConnector](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _NibOutletConnectorClass) New() NibOutletConnector {
	rv := objc.Send[NibOutletConnector](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewNibOutletConnector creates and returns a new initialized instance.
func NewNibOutletConnector() NibOutletConnector {
	return NibOutletConnectorClass.New()
}

// Init initializes the instance.
func (n_ NibOutletConnector) Init() NibOutletConnector {
	rv := objc.Send[NibOutletConnector](n_.ID(), selInit)
	return rv
}
