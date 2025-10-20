// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKAllowedSharingOptions] class.
var (
	CKAllowedSharingOptionsClass     _CKAllowedSharingOptionsClass
	CKAllowedSharingOptionsClassOnce sync.Once
)

func getCKAllowedSharingOptionsClass() _CKAllowedSharingOptionsClass {
	CKAllowedSharingOptionsClassOnce.Do(func() {
		CKAllowedSharingOptionsClass = _CKAllowedSharingOptionsClass{objc.GetClass("CKAllowedSharingOptions")}
	})
	return CKAllowedSharingOptionsClass
}

type _CKAllowedSharingOptionsClass struct {
	class objc.Class
}

// An interface definition for the [CKAllowedSharingOptions] class.
type ICKAllowedSharingOptions interface {
	objectivec.IObject
}

// An object that controls participant access and permission options.
//
// Register an instance of this class with an or when preparing a before your app invokes the share sheet. The share sheet uses the registered   object to let the user choose between the allowed options when sharing.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions
type CKAllowedSharingOptions struct {
	objectivec.Object
}

// CKAllowedSharingOptionsFrom constructs a [CKAllowedSharingOptions] from an unsafe.Pointer.
//
// An object that controls participant access and permission options.
func CKAllowedSharingOptionsFrom(ptr unsafe.Pointer) CKAllowedSharingOptions {
	return CKAllowedSharingOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKAllowedSharingOptionsClass) Alloc() CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKAllowedSharingOptionsClass) New() CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKAllowedSharingOptions) Init() CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKAllowedSharingOptions) Autorelease() CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKAllowedSharingOptions creates a new CKAllowedSharingOptions instance.
func NewCKAllowedSharingOptions() CKAllowedSharingOptions {
	return getCKAllowedSharingOptionsClass().New()
}




