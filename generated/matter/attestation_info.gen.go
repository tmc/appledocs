// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AttestationInfo] class.
var (
	AttestationInfoClass     _AttestationInfoClass
	AttestationInfoClassOnce sync.Once
)

func getAttestationInfoClass() _AttestationInfoClass {
	AttestationInfoClassOnce.Do(func() {
		AttestationInfoClass = _AttestationInfoClass{objc.GetClass("AttestationInfo")}
	})
	return AttestationInfoClass
}

type _AttestationInfoClass struct {
	class objc.Class
}

// An interface definition for the [AttestationInfo] class.
type IAttestationInfo interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/AttestationInfo
type AttestationInfo struct {
	objectivec.Object
}

// AttestationInfoFrom constructs a [AttestationInfo] from an unsafe.Pointer.
func AttestationInfoFrom(ptr unsafe.Pointer) AttestationInfo {
	return AttestationInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AttestationInfoClass) Alloc() AttestationInfo {
	rv := objc.Send[AttestationInfo](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AttestationInfoClass) New() AttestationInfo {
	rv := objc.Send[AttestationInfo](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AttestationInfo) Init() AttestationInfo {
	rv := objc.Send[AttestationInfo](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AttestationInfo) Autorelease() AttestationInfo {
	rv := objc.Send[AttestationInfo](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAttestationInfo creates a new AttestationInfo instance.
func NewAttestationInfo() AttestationInfo {
	return getAttestationInfoClass().New()
}




