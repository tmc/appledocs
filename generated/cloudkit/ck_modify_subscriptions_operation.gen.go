// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKModifySubscriptionsOperation */


/* debug [class_header]: Header for CKModifySubscriptionsOperation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKModifySubscriptionsOperation */
// An interface definition for the [CKModifySubscriptionsOperation] class.
type ICKModifySubscriptionsOperation interface {
	ICKDatabaseOperation
	
/* debug [class_interface_properties]: Properties for CKModifySubscriptionsOperation */
	// properties:
	ModifySubscriptionsCompletionBlock() unsafe.Pointer
	SetModifySubscriptionsCompletionBlock(value unsafe.Pointer)
	PerSubscriptionDeleteBlock() func(unsafe.Pointer, unsafe.Pointer)
	SetPerSubscriptionDeleteBlock(value func(unsafe.Pointer, unsafe.Pointer))
	PerSubscriptionSaveBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	SetPerSubscriptionSaveBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer))
	SubscriptionIDsToDelete() []string
	SetSubscriptionIDsToDelete(value []string)
	SubscriptionsToSave() []CKSubscription
	SetSubscriptionsToSave(value []CKSubscription)
	ModifySubscriptionsResultBlock() objectivec.IObject
	SetModifySubscriptionsResultBlock(value objectivec.IObject)
	CompletionBlock() objectivec.IObject
	SetCompletionBlock(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKModifySubscriptionsOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKModifySubscriptionsOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKModifySubscriptionsOperationClass) Alloc() CKModifySubscriptionsOperation {
	rv := objc.Send[CKModifySubscriptionsOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKModifySubscriptionsOperation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKModifySubscriptionsOperation */

// Creates an operation for saving and deleting the specified subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/initWithSubscriptionsToSave:subscriptionIDsToDelete:
func NewCKModifySubscriptionsOperationWithSubscriptionsToSaveSubscriptionIDsToDelete(subscriptionsToSave []CKSubscription, subscriptionIDsToDelete []string) CKModifySubscriptionsOperation {
	instance := getCKModifySubscriptionsOperationClass().Alloc()
	rv := objc.Send[CKModifySubscriptionsOperation](instance.ID, objc.Sel("initWithSubscriptionsToSave:subscriptionIDsToDelete:"), subscriptionsToSave, subscriptionIDsToDelete)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKModifySubscriptionsOperationWithSubscriptionsToSaveSubscriptionIDsToDelete */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKModifySubscriptionsOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKModifySubscriptionsOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKModifySubscriptionsOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKModifySubscriptionsOperation */

// The block to execute after the operation modifies the subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/modifySubscriptionsCompletionBlock-3v0cp
func (c_ CKModifySubscriptionsOperation) ModifySubscriptionsCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("modifySubscriptionsCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: modifySubscriptionsCompletionBlock */


// The block to execute after the operation modifies the subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/modifySubscriptionsCompletionBlock-3v0cp
func (c_ CKModifySubscriptionsOperation) SetModifySubscriptionsCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModifySubscriptionsCompletionBlock:"), value)
}/* debug [instance_properties/setter]: modifySubscriptionsCompletionBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/perSubscriptionDeleteBlock-55p5p
func (c_ CKModifySubscriptionsOperation) PerSubscriptionDeleteBlock() func(unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer)](c_.ID, objc.Sel("perSubscriptionDeleteBlock"))
	return rv
}/* debug [instance_properties/getter]: perSubscriptionDeleteBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/perSubscriptionDeleteBlock-55p5p
func (c_ CKModifySubscriptionsOperation) SetPerSubscriptionDeleteBlock(value func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerSubscriptionDeleteBlock:"), value)
}/* debug [instance_properties/setter]: perSubscriptionDeleteBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/perSubscriptionSaveBlock-1yn86
func (c_ CKModifySubscriptionsOperation) PerSubscriptionSaveBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)](c_.ID, objc.Sel("perSubscriptionSaveBlock"))
	return rv
}/* debug [instance_properties/getter]: perSubscriptionSaveBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/perSubscriptionSaveBlock-1yn86
func (c_ CKModifySubscriptionsOperation) SetPerSubscriptionSaveBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerSubscriptionSaveBlock:"), value)
}/* debug [instance_properties/setter]: perSubscriptionSaveBlock */


// The IDs of the subscriptions that you want to delete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/subscriptionIDsToDelete-14x82
func (c_ CKModifySubscriptionsOperation) SubscriptionIDsToDelete() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("subscriptionIDsToDelete"))
	return rv
}/* debug [instance_properties/getter]: subscriptionIDsToDelete */


// The IDs of the subscriptions that you want to delete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/subscriptionIDsToDelete-14x82
func (c_ CKModifySubscriptionsOperation) SetSubscriptionIDsToDelete(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubscriptionIDsToDelete:"), nsArray)
}/* debug [instance_properties/setter]: subscriptionIDsToDelete */


// The subscriptions to save to the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/subscriptionsToSave
func (c_ CKModifySubscriptionsOperation) SubscriptionsToSave() []CKSubscription {
	rv := objc.Send[[]CKSubscription](c_.ID, objc.Sel("subscriptionsToSave"))
	return rv
}/* debug [instance_properties/getter]: subscriptionsToSave */


// The subscriptions to save to the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifySubscriptionsOperation/subscriptionsToSave
func (c_ CKModifySubscriptionsOperation) SetSubscriptionsToSave(value []CKSubscription) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubscriptionsToSave:"), nsArray)
}/* debug [instance_properties/setter]: subscriptionsToSave */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/modifysubscriptionsresultblock
func (c_ CKModifySubscriptionsOperation) ModifySubscriptionsResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("modifySubscriptionsResultBlock"))
	return rv
}/* debug [instance_properties/getter]: modifySubscriptionsResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifysubscriptionsoperation/modifysubscriptionsresultblock
func (c_ CKModifySubscriptionsOperation) SetModifySubscriptionsResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModifySubscriptionsResultBlock:"), value)
}/* debug [instance_properties/setter]: modifySubscriptionsResultBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKModifySubscriptionsOperation) CompletionBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("completionBlock"))
	return rv
}/* debug [instance_properties/getter]: completionBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKModifySubscriptionsOperation) SetCompletionBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}/* debug [instance_properties/setter]: completionBlock */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKModifySubscriptionsOperation */


