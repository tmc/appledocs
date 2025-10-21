// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CKSyncEngineSendChangesOptions] class.
var (
	CKSyncEngineSendChangesOptionsClass     _CKSyncEngineSendChangesOptionsClass
	CKSyncEngineSendChangesOptionsClassOnce sync.Once
)

func getCKSyncEngineSendChangesOptionsClass() _CKSyncEngineSendChangesOptionsClass {
	CKSyncEngineSendChangesOptionsClassOnce.Do(func() {
		CKSyncEngineSendChangesOptionsClass = _CKSyncEngineSendChangesOptionsClass{objc.GetClass("CKSyncEngineSendChangesOptions")}
	})
	return CKSyncEngineSendChangesOptionsClass
}

type _CKSyncEngineSendChangesOptionsClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineSendChangesOptions] class.
type ICKSyncEngineSendChangesOptions interface {
	objectivec.IObject
}

// A set of options to use with a send operation.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions
type CKSyncEngineSendChangesOptions struct {
	objectivec.Object
}

// CKSyncEngineSendChangesOptionsFrom constructs a [CKSyncEngineSendChangesOptions] from an unsafe.Pointer.
//
// A set of options to use with a send operation.
func CKSyncEngineSendChangesOptionsFrom(ptr unsafe.Pointer) CKSyncEngineSendChangesOptions {
	return CKSyncEngineSendChangesOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineSendChangesOptionsClass) Alloc() CKSyncEngineSendChangesOptions {
	rv := objc.Send[CKSyncEngineSendChangesOptions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineSendChangesOptionsClass) New() CKSyncEngineSendChangesOptions {
	rv := objc.Send[CKSyncEngineSendChangesOptions](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineSendChangesOptions) Init() CKSyncEngineSendChangesOptions {
	rv := objc.Send[CKSyncEngineSendChangesOptions](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineSendChangesOptions) Autorelease() CKSyncEngineSendChangesOptions {
	rv := objc.Send[CKSyncEngineSendChangesOptions](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineSendChangesOptions creates a new CKSyncEngineSendChangesOptions instance.
func NewCKSyncEngineSendChangesOptions() CKSyncEngineSendChangesOptions {
	return getCKSyncEngineSendChangesOptionsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions/initWithScope:
func NewCKSyncEngineSendChangesOptionsWithScope(scope unsafe.Pointer) CKSyncEngineSendChangesOptions {
	instance := getCKSyncEngineSendChangesOptionsClass().Alloc()
	rv := objc.Send[CKSyncEngineSendChangesOptions](instance.ID, objc.Sel("initWithScope:"), scope)
	rv.Autorelease()
	return rv
}


// The operation group to use for the underlying CloudKit operations.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions/operationGroup
func (c_ CKSyncEngineSendChangesOptions) OperationGroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("operationGroup"))
	return rv
}


// SetOperationGroup sets the value of the operationGroup property.
// The operation group to use for the underlying CloudKit operations.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions/operationGroup
func (c_ CKSyncEngineSendChangesOptions) SetOperationGroup(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOperationGroup:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions/scope
func (c_ CKSyncEngineSendChangesOptions) Scope() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("scope"))
	return rv
}


// SetScope sets the value of the scope property.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions/scope
func (c_ CKSyncEngineSendChangesOptions) SetScope(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScope:"), value)
}


