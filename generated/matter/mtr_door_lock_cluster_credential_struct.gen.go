// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterCredentialStruct] class.
var (
	MTRDoorLockClusterCredentialStructClass     _MTRDoorLockClusterCredentialStructClass
	MTRDoorLockClusterCredentialStructClassOnce sync.Once
)

func getMTRDoorLockClusterCredentialStructClass() _MTRDoorLockClusterCredentialStructClass {
	MTRDoorLockClusterCredentialStructClassOnce.Do(func() {
		MTRDoorLockClusterCredentialStructClass = _MTRDoorLockClusterCredentialStructClass{objc.GetClass("MTRDoorLockClusterCredentialStruct")}
	})
	return MTRDoorLockClusterCredentialStructClass
}

type _MTRDoorLockClusterCredentialStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterCredentialStruct] class.
type IMTRDoorLockClusterCredentialStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterCredentialStruct
type MTRDoorLockClusterCredentialStruct struct {
	objectivec.Object
}

// MTRDoorLockClusterCredentialStructFrom constructs a [MTRDoorLockClusterCredentialStruct] from an unsafe.Pointer.
func MTRDoorLockClusterCredentialStructFrom(ptr unsafe.Pointer) MTRDoorLockClusterCredentialStruct {
	return MTRDoorLockClusterCredentialStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterCredentialStructClass) Alloc() MTRDoorLockClusterCredentialStruct {
	rv := objc.Send[MTRDoorLockClusterCredentialStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterCredentialStructClass) New() MTRDoorLockClusterCredentialStruct {
	rv := objc.Send[MTRDoorLockClusterCredentialStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterCredentialStruct) Init() MTRDoorLockClusterCredentialStruct {
	rv := objc.Send[MTRDoorLockClusterCredentialStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterCredentialStruct) Autorelease() MTRDoorLockClusterCredentialStruct {
	rv := objc.Send[MTRDoorLockClusterCredentialStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterCredentialStruct creates a new MTRDoorLockClusterCredentialStruct instance.
func NewMTRDoorLockClusterCredentialStruct() MTRDoorLockClusterCredentialStruct {
	return getMTRDoorLockClusterCredentialStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustercredentialstruct/credentialindex
func (m_ MTRDoorLockClusterCredentialStruct) CredentialIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("credentialIndex"))
	return rv
}


// SetCredentialIndex sets the value of the credentialIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustercredentialstruct/credentialindex
func (m_ MTRDoorLockClusterCredentialStruct) SetCredentialIndex(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustercredentialstruct/credentialtype
func (m_ MTRDoorLockClusterCredentialStruct) CredentialType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("credentialType"))
	return rv
}


// SetCredentialType sets the value of the credentialType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustercredentialstruct/credentialtype
func (m_ MTRDoorLockClusterCredentialStruct) SetCredentialType(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialType:"), value)
}



