// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NWTLSParameters] class.
var (
	NWTLSParametersClass     _NWTLSParametersClass
	NWTLSParametersClassOnce sync.Once
)

func getNWTLSParametersClass() _NWTLSParametersClass {
	NWTLSParametersClassOnce.Do(func() {
		NWTLSParametersClass = _NWTLSParametersClass{objc.GetClass("NWTLSParameters")}
	})
	return NWTLSParametersClass
}

type _NWTLSParametersClass struct {
	class objc.Class
}

// An interface definition for the [NWTLSParameters] class.
type INWTLSParameters interface {
	objectivec.IObject
}

// TLS properties for creating a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters
type NWTLSParameters struct {
	objectivec.Object
}

// NWTLSParametersFrom constructs a [NWTLSParameters] from an unsafe.Pointer.
//
// TLS properties for creating a connection.
func NWTLSParametersFrom(ptr unsafe.Pointer) NWTLSParameters {
	return NWTLSParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NWTLSParametersClass) Alloc() NWTLSParameters {
	rv := objc.Send[NWTLSParameters](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NWTLSParametersClass) New() NWTLSParameters {
	rv := objc.Send[NWTLSParameters](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NWTLSParameters) Init() NWTLSParameters {
	rv := objc.Send[NWTLSParameters](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NWTLSParameters) Autorelease() NWTLSParameters {
	rv := objc.Send[NWTLSParameters](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWTLSParameters creates a new NWTLSParameters instance.
func NewNWTLSParameters() NWTLSParameters {
	return getNWTLSParametersClass().New()
}




