// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CKSyncEngineSendChangesContext] class.
var (
	CKSyncEngineSendChangesContextClass     _CKSyncEngineSendChangesContextClass
	CKSyncEngineSendChangesContextClassOnce sync.Once
)

func getCKSyncEngineSendChangesContextClass() _CKSyncEngineSendChangesContextClass {
	CKSyncEngineSendChangesContextClassOnce.Do(func() {
		CKSyncEngineSendChangesContextClass = _CKSyncEngineSendChangesContextClass{objc.GetClass("CKSyncEngineSendChangesContext")}
	})
	return CKSyncEngineSendChangesContextClass
}

type _CKSyncEngineSendChangesContextClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineSendChangesContext] class.
type ICKSyncEngineSendChangesContext interface {
	objectivec.IObject
}

// An object that describes a single attempt to send changes to the iCloud servers.
//
// A sync engine has two ways to send changes to iCloud — periodically, in cooperation with the system scheduler, and manually, whenever your app invokes the method. This object provides information about a single attempt to send changes that includes both the reason for the attempt and any additional options in use by the attempt.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesContext
type CKSyncEngineSendChangesContext struct {
	objectivec.Object
}

// CKSyncEngineSendChangesContextFrom constructs a [CKSyncEngineSendChangesContext] from an unsafe.Pointer.
//
// An object that describes a single attempt to send changes to the iCloud servers.
func CKSyncEngineSendChangesContextFrom(ptr unsafe.Pointer) CKSyncEngineSendChangesContext {
	return CKSyncEngineSendChangesContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineSendChangesContextClass) Alloc() CKSyncEngineSendChangesContext {
	rv := objc.Send[CKSyncEngineSendChangesContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineSendChangesContextClass) New() CKSyncEngineSendChangesContext {
	rv := objc.Send[CKSyncEngineSendChangesContext](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineSendChangesContext) Init() CKSyncEngineSendChangesContext {
	rv := objc.Send[CKSyncEngineSendChangesContext](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineSendChangesContext) Autorelease() CKSyncEngineSendChangesContext {
	rv := objc.Send[CKSyncEngineSendChangesContext](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineSendChangesContext creates a new CKSyncEngineSendChangesContext instance.
func NewCKSyncEngineSendChangesContext() CKSyncEngineSendChangesContext {
	return getCKSyncEngineSendChangesContextClass().New()
}


// The additional options for the send operation.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesContext/options
func (c_ CKSyncEngineSendChangesContext) Options() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("options"))
	return rv
}

// The reason for the send operation.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesContext/reason
func (c_ CKSyncEngineSendChangesContext) Reason() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("reason"))
	return rv
}



