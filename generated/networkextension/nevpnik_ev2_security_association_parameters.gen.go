// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [NEVPNIKEv2SecurityAssociationParameters] class.
var (
	NEVPNIKEv2SecurityAssociationParametersClass     _NEVPNIKEv2SecurityAssociationParametersClass
	NEVPNIKEv2SecurityAssociationParametersClassOnce sync.Once
)

func getNEVPNIKEv2SecurityAssociationParametersClass() _NEVPNIKEv2SecurityAssociationParametersClass {
	NEVPNIKEv2SecurityAssociationParametersClassOnce.Do(func() {
		NEVPNIKEv2SecurityAssociationParametersClass = _NEVPNIKEv2SecurityAssociationParametersClass{objc.GetClass("NEVPNIKEv2SecurityAssociationParameters")}
	})
	return NEVPNIKEv2SecurityAssociationParametersClass
}

type _NEVPNIKEv2SecurityAssociationParametersClass struct {
	class objc.Class
}

// An interface definition for the [NEVPNIKEv2SecurityAssociationParameters] class.
type INEVPNIKEv2SecurityAssociationParameters interface {
	objectivec.IObject
}

// Parameters for an IKEv2 Security Association.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters
type NEVPNIKEv2SecurityAssociationParameters struct {
	objectivec.Object
}

// NEVPNIKEv2SecurityAssociationParametersFrom constructs a [NEVPNIKEv2SecurityAssociationParameters] from an unsafe.Pointer.
//
// Parameters for an IKEv2 Security Association.
func NEVPNIKEv2SecurityAssociationParametersFrom(ptr unsafe.Pointer) NEVPNIKEv2SecurityAssociationParameters {
	return NEVPNIKEv2SecurityAssociationParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEVPNIKEv2SecurityAssociationParametersClass) Alloc() NEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEVPNIKEv2SecurityAssociationParametersClass) New() NEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEVPNIKEv2SecurityAssociationParameters) Init() NEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEVPNIKEv2SecurityAssociationParameters) Autorelease() NEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNIKEv2SecurityAssociationParameters creates a new NEVPNIKEv2SecurityAssociationParameters instance.
func NewNEVPNIKEv2SecurityAssociationParameters() NEVPNIKEv2SecurityAssociationParameters {
	return getNEVPNIKEv2SecurityAssociationParametersClass().New()
}




