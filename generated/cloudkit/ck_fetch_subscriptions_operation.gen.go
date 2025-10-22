// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKFetchSubscriptionsOperation] class.
var (
	CKFetchSubscriptionsOperationClass     _CKFetchSubscriptionsOperationClass
	CKFetchSubscriptionsOperationClassOnce sync.Once
)

func getCKFetchSubscriptionsOperationClass() _CKFetchSubscriptionsOperationClass {
	CKFetchSubscriptionsOperationClassOnce.Do(func() {
		CKFetchSubscriptionsOperationClass = _CKFetchSubscriptionsOperationClass{objc.GetClass("CKFetchSubscriptionsOperation")}
	})
	return CKFetchSubscriptionsOperationClass
}

type _CKFetchSubscriptionsOperationClass struct {
	class objc.Class
}

// An interface definition for the [CKFetchSubscriptionsOperation] class.
type ICKFetchSubscriptionsOperation interface {
	ICKDatabaseOperation
	FetchSubscriptionCompletionBlock() unsafe.Pointer
	SetFetchSubscriptionCompletionBlock(value unsafe.Pointer)
	FetchSubscriptionsResultBlock() unsafe.Pointer
	SetFetchSubscriptionsResultBlock(value unsafe.Pointer)
	PerSubscriptionResultBlock() unsafe.Pointer
	SetPerSubscriptionResultBlock(value unsafe.Pointer)
	SubscriptionIDs() unsafe.Pointer
	SetSubscriptionIDs(value unsafe.Pointer)
	CompletionBlock() unsafe.Pointer
	SetCompletionBlock(value unsafe.Pointer)
}

// An operation for fetching subscriptions.
//
// A fetch subscriptions operation retrieves subscriptions (with IDs you already know) from iCloud and can fetch all subscriptions for the current user. You might fetch subscriptions so you can examine or modify their parameters — for example, to adjust the delivery options for push notifications that the subscription generates. If you assign a handler to the property, the operation calls it after it executes and passes it the results. Use the handler to perform any housekeeping tasks for the operation. The handler you specify should manage any failures, whether due to an error or an explicit cancellation.


// An operation for fetching subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation

type CKFetchSubscriptionsOperation struct {
	CKDatabaseOperation
}

// CKFetchSubscriptionsOperationFrom constructs a [CKFetchSubscriptionsOperation] from an unsafe.Pointer.
//
// An operation for fetching subscriptions.
func CKFetchSubscriptionsOperationFrom(ptr unsafe.Pointer) CKFetchSubscriptionsOperation {
	return CKFetchSubscriptionsOperation{
		CKDatabaseOperation: CKDatabaseOperationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKFetchSubscriptionsOperationClass) Alloc() CKFetchSubscriptionsOperation {
	rv := objc.Send[CKFetchSubscriptionsOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKFetchSubscriptionsOperationClass) New() CKFetchSubscriptionsOperation {
	rv := objc.Send[CKFetchSubscriptionsOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKFetchSubscriptionsOperation) Init() CKFetchSubscriptionsOperation {
	rv := objc.Send[CKFetchSubscriptionsOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKFetchSubscriptionsOperation) Autorelease() CKFetchSubscriptionsOperation {
	rv := objc.Send[CKFetchSubscriptionsOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchSubscriptionsOperation creates a new CKFetchSubscriptionsOperation instance.
func NewCKFetchSubscriptionsOperation() CKFetchSubscriptionsOperation {
	return getCKFetchSubscriptionsOperationClass().New()
}



// The block to execute with the fetch results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/fetchsubscriptioncompletionblock-6hhpi

func (c_ CKFetchSubscriptionsOperation) FetchSubscriptionCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchSubscriptionCompletionBlock"))
	return rv
}


// The block to execute with the fetch results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/fetchsubscriptioncompletionblock-6hhpi

func (c_ CKFetchSubscriptionsOperation) SetFetchSubscriptionCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchSubscriptionCompletionBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/fetchsubscriptionsresultblock

func (c_ CKFetchSubscriptionsOperation) FetchSubscriptionsResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchSubscriptionsResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/fetchsubscriptionsresultblock

func (c_ CKFetchSubscriptionsOperation) SetFetchSubscriptionsResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchSubscriptionsResultBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/persubscriptionresultblock

func (c_ CKFetchSubscriptionsOperation) PerSubscriptionResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perSubscriptionResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/persubscriptionresultblock

func (c_ CKFetchSubscriptionsOperation) SetPerSubscriptionResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerSubscriptionResultBlock:"), value)
}


// The IDs of the subscriptions to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/subscriptionids-17f4q

func (c_ CKFetchSubscriptionsOperation) SubscriptionIDs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("subscriptionIDs"))
	return rv
}


// The IDs of the subscriptions to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/subscriptionids-17f4q

func (c_ CKFetchSubscriptionsOperation) SetSubscriptionIDs(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubscriptionIDs:"), value)
}


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock

func (c_ CKFetchSubscriptionsOperation) CompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("completionBlock"))
	return rv
}


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock

func (c_ CKFetchSubscriptionsOperation) SetCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}



