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
}

// An operation for fetching subscriptions.
//
// A fetch subscriptions operation retrieves subscriptions (with IDs you already know) from iCloud and can fetch all subscriptions for the current user. You might fetch subscriptions so you can examine or modify their parameters — for example, to adjust the delivery options for push notifications that the subscription generates. If you assign a handler to the property, the operation calls it after it executes and passes it the results. Use the handler to perform any housekeeping tasks for the operation. The handler you specify should manage any failures, whether due to an error or an explicit cancellation.
//
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




