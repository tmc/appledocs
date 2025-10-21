// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKModifySubscriptionsOperation] class.
var (
	CKModifySubscriptionsOperationClass     _CKModifySubscriptionsOperationClass
	CKModifySubscriptionsOperationClassOnce sync.Once
)

func getCKModifySubscriptionsOperationClass() _CKModifySubscriptionsOperationClass {
	CKModifySubscriptionsOperationClassOnce.Do(func() {
		CKModifySubscriptionsOperationClass = _CKModifySubscriptionsOperationClass{objc.GetClass("CKModifySubscriptionsOperation")}
	})
	return CKModifySubscriptionsOperationClass
}

type _CKModifySubscriptionsOperationClass struct {
	class objc.Class
}

// An interface definition for the [CKModifySubscriptionsOperation] class.
type ICKModifySubscriptionsOperation interface {
	ICKDatabaseOperation
}

// An operation for modifying one or more subscriptions.
//
// After you create or change the configuration of a subscription, use this operation to save those changes to the server. You can also use this operation to permanently delete subscriptions. If you assign a handler to the property, the operation calls it after it executes and passes it the results. Use the handler to perform any housekeeping tasks for the operation. The handler you specify should manage any failures, whether due to an error or an explicit cancellation.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation
type CKModifySubscriptionsOperation struct {
	CKDatabaseOperation
}

// CKModifySubscriptionsOperationFrom constructs a [CKModifySubscriptionsOperation] from an unsafe.Pointer.
//
// An operation for modifying one or more subscriptions.
func CKModifySubscriptionsOperationFrom(ptr unsafe.Pointer) CKModifySubscriptionsOperation {
	return CKModifySubscriptionsOperation{
		CKDatabaseOperation: CKDatabaseOperationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKModifySubscriptionsOperationClass) Alloc() CKModifySubscriptionsOperation {
	rv := objc.Send[CKModifySubscriptionsOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKModifySubscriptionsOperationClass) New() CKModifySubscriptionsOperation {
	rv := objc.Send[CKModifySubscriptionsOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKModifySubscriptionsOperation) Init() CKModifySubscriptionsOperation {
	rv := objc.Send[CKModifySubscriptionsOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKModifySubscriptionsOperation) Autorelease() CKModifySubscriptionsOperation {
	rv := objc.Send[CKModifySubscriptionsOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKModifySubscriptionsOperation creates a new CKModifySubscriptionsOperation instance.
func NewCKModifySubscriptionsOperation() CKModifySubscriptionsOperation {
	return getCKModifySubscriptionsOperationClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/perSubscriptionSaveBlock-1yn86
func (c_ CKModifySubscriptionsOperation) PerSubscriptionSaveBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perSubscriptionSaveBlock"))
	return rv
}


// SetPerSubscriptionSaveBlock sets the value of the perSubscriptionSaveBlock property.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/perSubscriptionSaveBlock-1yn86
func (c_ CKModifySubscriptionsOperation) SetPerSubscriptionSaveBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerSubscriptionSaveBlock:"), value)
}



