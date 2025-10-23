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
	ModifySubscriptionsCompletionBlock() unsafe.Pointer
	SetModifySubscriptionsCompletionBlock(value unsafe.Pointer)
	ModifySubscriptionsResultBlock() unsafe.Pointer
	SetModifySubscriptionsResultBlock(value unsafe.Pointer)
	PerSubscriptionDeleteBlock() unsafe.Pointer
	SetPerSubscriptionDeleteBlock(value unsafe.Pointer)
	PerSubscriptionSaveBlock() unsafe.Pointer
	SetPerSubscriptionSaveBlock(value unsafe.Pointer)
	SubscriptionIDsToDelete() unsafe.Pointer
	SetSubscriptionIDsToDelete(value unsafe.Pointer)
	SubscriptionsToSave() ICKSubscription
	SetSubscriptionsToSave(value ICKSubscription)
	CompletionBlock() unsafe.Pointer
	SetCompletionBlock(value unsafe.Pointer)
}

// An operation for modifying one or more subscriptions.
//
// After you create or change the configuration of a subscription, use this operation to save those changes to the server. You can also use this operation to permanently delete subscriptions. If you assign a handler to the property, the operation calls it after it executes and passes it the results. Use the handler to perform any housekeeping tasks for the operation. The handler you specify should manage any failures, whether due to an error or an explicit cancellation.


// An operation for modifying one or more subscriptions.
//
// [Full Topic]
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



// The closure to execute after the operation modifies the subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/modifysubscriptionscompletionblock-7l56
func (c_ CKModifySubscriptionsOperation) ModifySubscriptionsCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("modifySubscriptionsCompletionBlock"))
	return rv
}


// The closure to execute after the operation modifies the subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/modifysubscriptionscompletionblock-7l56
func (c_ CKModifySubscriptionsOperation) SetModifySubscriptionsCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModifySubscriptionsCompletionBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/modifysubscriptionsresultblock
func (c_ CKModifySubscriptionsOperation) ModifySubscriptionsResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("modifySubscriptionsResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/modifysubscriptionsresultblock
func (c_ CKModifySubscriptionsOperation) SetModifySubscriptionsResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModifySubscriptionsResultBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/persubscriptiondeleteblock-5ke2l
func (c_ CKModifySubscriptionsOperation) PerSubscriptionDeleteBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perSubscriptionDeleteBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/persubscriptiondeleteblock-5ke2l
func (c_ CKModifySubscriptionsOperation) SetPerSubscriptionDeleteBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerSubscriptionDeleteBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/persubscriptionsaveblock-8y9zn
func (c_ CKModifySubscriptionsOperation) PerSubscriptionSaveBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perSubscriptionSaveBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/persubscriptionsaveblock-8y9zn
func (c_ CKModifySubscriptionsOperation) SetPerSubscriptionSaveBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerSubscriptionSaveBlock:"), value)
}


// The IDs of the subscriptions that you want to delete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/subscriptionidstodelete-3534e
func (c_ CKModifySubscriptionsOperation) SubscriptionIDsToDelete() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("subscriptionIDsToDelete"))
	return rv
}


// The IDs of the subscriptions that you want to delete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/subscriptionidstodelete-3534e
func (c_ CKModifySubscriptionsOperation) SetSubscriptionIDsToDelete(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubscriptionIDsToDelete:"), value)
}


// The subscriptions to save to the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/subscriptionstosave
func (c_ CKModifySubscriptionsOperation) SubscriptionsToSave() ICKSubscription {
	rv := objc.Send[CKSubscription](c_.ID, objc.Sel("subscriptionsToSave"))
	return rv
}


// The subscriptions to save to the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/subscriptionstosave
func (c_ CKModifySubscriptionsOperation) SetSubscriptionsToSave(value ICKSubscription) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubscriptionsToSave:"), value)
}


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKModifySubscriptionsOperation) CompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("completionBlock"))
	return rv
}


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKModifySubscriptionsOperation) SetCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}



