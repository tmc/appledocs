// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKFetchSubscriptionsOperation */


/* debug [class_header]: Header for CKFetchSubscriptionsOperation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKFetchSubscriptionsOperation */
// An interface definition for the [CKFetchSubscriptionsOperation] class.
type ICKFetchSubscriptionsOperation interface {
	ICKDatabaseOperation
	
/* debug [class_interface_properties]: Properties for CKFetchSubscriptionsOperation */
	// properties:
	FetchSubscriptionCompletionBlock() unsafe.Pointer
	SetFetchSubscriptionCompletionBlock(value unsafe.Pointer)
	PerSubscriptionCompletionBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	SetPerSubscriptionCompletionBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer))
	SubscriptionIDs() []string
	SetSubscriptionIDs(value []string)
	FetchSubscriptionsResultBlock() objectivec.IObject
	SetFetchSubscriptionsResultBlock(value objectivec.IObject)
	PerSubscriptionResultBlock() objectivec.IObject
	SetPerSubscriptionResultBlock(value objectivec.IObject)
	CompletionBlock() objectivec.IObject
	SetCompletionBlock(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKFetchSubscriptionsOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKFetchSubscriptionsOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKFetchSubscriptionsOperationClass) Alloc() CKFetchSubscriptionsOperation {
	rv := objc.Send[CKFetchSubscriptionsOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKFetchSubscriptionsOperation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKFetchSubscriptionsOperation */

// Creates an operation for fetching the specified subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation/initWithSubscriptionIDs:
func NewCKFetchSubscriptionsOperationWithSubscriptionIDs(subscriptionIDs []string) CKFetchSubscriptionsOperation {
	instance := getCKFetchSubscriptionsOperationClass().Alloc()
	rv := objc.Send[CKFetchSubscriptionsOperation](instance.ID, objc.Sel("initWithSubscriptionIDs:"), subscriptionIDs)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKFetchSubscriptionsOperationWithSubscriptionIDs */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKFetchSubscriptionsOperation */

// Returns an operation that fetches all of the user’s subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation/fetchAllSubscriptionsOperation()
func (cc _CKFetchSubscriptionsOperationClass) FetchAllSubscriptionsOperation() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("fetchAllSubscriptionsOperation"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FetchAllSubscriptionsOperation) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKFetchSubscriptionsOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKFetchSubscriptionsOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKFetchSubscriptionsOperation */

// The block to execute after the operation fetches the subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation/fetchSubscriptionCompletionBlock-207ep
func (c_ CKFetchSubscriptionsOperation) FetchSubscriptionCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchSubscriptionCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchSubscriptionCompletionBlock */


// The block to execute after the operation fetches the subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation/fetchSubscriptionCompletionBlock-207ep
func (c_ CKFetchSubscriptionsOperation) SetFetchSubscriptionCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchSubscriptionCompletionBlock:"), value)
}/* debug [instance_properties/setter]: fetchSubscriptionCompletionBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation/perSubscriptionCompletionBlock
func (c_ CKFetchSubscriptionsOperation) PerSubscriptionCompletionBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)](c_.ID, objc.Sel("perSubscriptionCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: perSubscriptionCompletionBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation/perSubscriptionCompletionBlock
func (c_ CKFetchSubscriptionsOperation) SetPerSubscriptionCompletionBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerSubscriptionCompletionBlock:"), value)
}/* debug [instance_properties/setter]: perSubscriptionCompletionBlock */


// The IDs of the subscriptions to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation/subscriptionIDs-714ct
func (c_ CKFetchSubscriptionsOperation) SubscriptionIDs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("subscriptionIDs"))
	return rv
}/* debug [instance_properties/getter]: subscriptionIDs */


// The IDs of the subscriptions to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchSubscriptionsOperation/subscriptionIDs-714ct
func (c_ CKFetchSubscriptionsOperation) SetSubscriptionIDs(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubscriptionIDs:"), nsArray)
}/* debug [instance_properties/setter]: subscriptionIDs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/fetchsubscriptionsresultblock
func (c_ CKFetchSubscriptionsOperation) FetchSubscriptionsResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("fetchSubscriptionsResultBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchSubscriptionsResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/fetchsubscriptionsresultblock
func (c_ CKFetchSubscriptionsOperation) SetFetchSubscriptionsResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchSubscriptionsResultBlock:"), value)
}/* debug [instance_properties/setter]: fetchSubscriptionsResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/persubscriptionresultblock
func (c_ CKFetchSubscriptionsOperation) PerSubscriptionResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("perSubscriptionResultBlock"))
	return rv
}/* debug [instance_properties/getter]: perSubscriptionResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsubscriptionsoperation/persubscriptionresultblock
func (c_ CKFetchSubscriptionsOperation) SetPerSubscriptionResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerSubscriptionResultBlock:"), value)
}/* debug [instance_properties/setter]: perSubscriptionResultBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKFetchSubscriptionsOperation) CompletionBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("completionBlock"))
	return rv
}/* debug [instance_properties/getter]: completionBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKFetchSubscriptionsOperation) SetCompletionBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}/* debug [instance_properties/setter]: completionBlock */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKFetchSubscriptionsOperation */


