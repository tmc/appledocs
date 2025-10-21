// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DomainStateBiometry] class.
var (
	DomainStateBiometryClass     _DomainStateBiometryClass
	DomainStateBiometryClassOnce sync.Once
)

func getDomainStateBiometryClass() _DomainStateBiometryClass {
	DomainStateBiometryClassOnce.Do(func() {
		DomainStateBiometryClass = _DomainStateBiometryClass{objc.GetClass("LADomainStateBiometry")}
	})
	return DomainStateBiometryClass
}

type _DomainStateBiometryClass struct {
	class objc.Class
}

// An interface definition for the [DomainStateBiometry] class.
type IDomainStateBiometry interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateBiometry
type DomainStateBiometry struct {
	objectivec.Object
}

// DomainStateBiometryFrom constructs a [DomainStateBiometry] from an unsafe.Pointer.
func DomainStateBiometryFrom(ptr unsafe.Pointer) DomainStateBiometry {
	return DomainStateBiometry{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DomainStateBiometryClass) Alloc() DomainStateBiometry {
	rv := objc.Send[DomainStateBiometry](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DomainStateBiometryClass) New() DomainStateBiometry {
	rv := objc.Send[DomainStateBiometry](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DomainStateBiometry) Init() DomainStateBiometry {
	rv := objc.Send[DomainStateBiometry](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DomainStateBiometry) Autorelease() DomainStateBiometry {
	rv := objc.Send[DomainStateBiometry](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDomainStateBiometry creates a new DomainStateBiometry instance.
func NewDomainStateBiometry() DomainStateBiometry {
	return getDomainStateBiometryClass().New()
}


// Indicates biometry type available on the device.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateBiometry/biometryType
func (d_ DomainStateBiometry) BiometryType() BiometryType {
	rv := objc.Send[BiometryType](d_.ID, objc.Sel("biometryType"))
	return rv
}

// Contains state hash data for the available biometry type. Returns if no biometry entities are enrolled.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateBiometry/stateHash
func (d_ DomainStateBiometry) StateHash() foundation.NSData {
	rv := objc.Send[foundation.NSData](d_.ID, objc.Sel("stateHash"))
	return rv
}



