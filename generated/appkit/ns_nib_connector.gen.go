// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NibConnector] class.
var (
	NibConnectorClass     _NibConnectorClass
	NibConnectorClassOnce sync.Once
)

func getNibConnectorClass() _NibConnectorClass {
	NibConnectorClassOnce.Do(func() {
		NibConnectorClass = _NibConnectorClass{objc.GetClass("NSNibConnector")}
	})
	return NibConnectorClass
}

type _NibConnectorClass struct {
	class objc.Class
}

// An interface definition for the [NibConnector] class.
type INibConnector interface {
	objectivec.IObject
	ReplaceObjectWithObject(oldObject objc.ID, newObject objc.ID)
}

// A connection between two nibs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibConnector
type NibConnector struct {
	objectivec.Object
}

// NibConnectorFrom constructs a [NibConnector] from an unsafe.Pointer.
//
// A connection between two nibs.
func NibConnectorFrom(ptr unsafe.Pointer) NibConnector {
	return NibConnector{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NibConnectorClass) Alloc() NibConnector {
	rv := objc.Send[NibConnector](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NibConnectorClass) New() NibConnector {
	rv := objc.Send[NibConnector](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NibConnector) Init() NibConnector {
	rv := objc.Send[NibConnector](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NibConnector) Autorelease() NibConnector {
	rv := objc.Send[NibConnector](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNibConnector creates a new NibConnector instance.
func NewNibConnector() NibConnector {
	return getNibConnectorClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibConnector/replaceObject:withObject:
func (n_ NibConnector) ReplaceObjectWithObject(oldObject objc.ID, newObject objc.ID) {
	objc.Send[objc.ID](n_.ID, objc.Sel("replaceObject:withObject:"), oldObject, newObject)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibConnector/source
func (n_ NibConnector) Source() objc.ID {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("source"))
	return rv
}


// SetSource sets the value of the source property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibConnector/source
func (n_ NibConnector) SetSource(value objc.ID) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSource:"), value)
}



