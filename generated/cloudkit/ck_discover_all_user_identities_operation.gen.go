// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKDiscoverAllUserIdentitiesOperation */


/* debug [class_header]: Header for CKDiscoverAllUserIdentitiesOperation */
// The class instance for the [CKDiscoverAllUserIdentitiesOperation] class.
var (
	CKDiscoverAllUserIdentitiesOperationClass     _CKDiscoverAllUserIdentitiesOperationClass
	CKDiscoverAllUserIdentitiesOperationClassOnce sync.Once
)

func getCKDiscoverAllUserIdentitiesOperationClass() _CKDiscoverAllUserIdentitiesOperationClass {
	CKDiscoverAllUserIdentitiesOperationClassOnce.Do(func() {
		CKDiscoverAllUserIdentitiesOperationClass = _CKDiscoverAllUserIdentitiesOperationClass{objc.GetClass("CKDiscoverAllUserIdentitiesOperation")}
	})
	return CKDiscoverAllUserIdentitiesOperationClass
}

type _CKDiscoverAllUserIdentitiesOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKDiscoverAllUserIdentitiesOperation */
// An interface definition for the [CKDiscoverAllUserIdentitiesOperation] class.
type ICKDiscoverAllUserIdentitiesOperation interface {
	ICKOperation
	
/* debug [class_interface_properties]: Properties for CKDiscoverAllUserIdentitiesOperation */
	// properties:
	DiscoverAllUserIdentitiesCompletionBlock() unsafe.Pointer
	SetDiscoverAllUserIdentitiesCompletionBlock(value unsafe.Pointer)
	UserIdentityDiscoveredBlock() unsafe.Pointer
	SetUserIdentityDiscoveredBlock(value unsafe.Pointer)
	DiscoverAllUserIdentitiesResultBlock() objectivec.IObject
	SetDiscoverAllUserIdentitiesResultBlock(value objectivec.IObject)
	ContactIdentifiers() objc.IObject /* cross-framework: NSString */
	SetContactIdentifiers(value objc.IObject /* cross-framework: NSString */)
	CompletionBlock() objectivec.IObject
	SetCompletionBlock(value objectivec.IObject)
	QualityOfService() objectivec.IObject
	SetQualityOfService(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKDiscoverAllUserIdentitiesOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKDiscoverAllUserIdentitiesOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKDiscoverAllUserIdentitiesOperationClass) Alloc() CKDiscoverAllUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverAllUserIdentitiesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKDiscoverAllUserIdentitiesOperationClass) New() CKDiscoverAllUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverAllUserIdentitiesOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKDiscoverAllUserIdentitiesOperation) Init() CKDiscoverAllUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverAllUserIdentitiesOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKDiscoverAllUserIdentitiesOperation) Autorelease() CKDiscoverAllUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverAllUserIdentitiesOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKDiscoverAllUserIdentitiesOperation creates a new CKDiscoverAllUserIdentitiesOperation instance.
func NewCKDiscoverAllUserIdentitiesOperation() CKDiscoverAllUserIdentitiesOperation {
	return getCKDiscoverAllUserIdentitiesOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKDiscoverAllUserIdentitiesOperation */
// An operation that uses the device’s contacts to search for discoverable iCloud users.
//
// Use this operation to discover iCloud users that match entries in the device’s Contacts database. CloudKit uses the email addresses and phone numbers in each Contact record to search for a matching iCloud account. Although your app doesn’t need authorization to use the Contacts database to execute this operation, if it has authorization, you can use the property on any returned user identity to fetch the corresponding Contact record from the database. Before CloudKit can return a user’s identity, you must ask for their permission by calling . Do this as part of any onboarding where you can highlight the benefits of being discoverable within the context of your app. The operation executes the handlers you provide on an internal queue it manages. Your handlers must be capable of executing on a background queue. Tasks that need access to the main queue must redirect as appropriate. The operation calls after it executes and returns results. Use the completion handler to perform housekeeping tasks for the operation. It should also manage any failures, whether due to an error or an explicit cancellation. CloudKit operations have a default QoS of . Operations with this service level are discretionary. The system schedules their execution at an optimal time according to battery level and network conditions, among other factors. Use the property to set a more appropriate QoS for the operation. The following example shows how to create the operation, configure its callbacks, and execute it using the default container’s queue:


// An operation that uses the device’s contacts to search for discoverable iCloud users.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverAllUserIdentitiesOperation
type CKDiscoverAllUserIdentitiesOperation struct {
	CKOperation
}

// CKDiscoverAllUserIdentitiesOperationFrom constructs a [CKDiscoverAllUserIdentitiesOperation] from an unsafe.Pointer.
//
// An operation that uses the device’s contacts to search for discoverable iCloud users.
func CKDiscoverAllUserIdentitiesOperationFrom(ptr unsafe.Pointer) CKDiscoverAllUserIdentitiesOperation {
	return CKDiscoverAllUserIdentitiesOperation{
		CKOperation: CKOperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKDiscoverAllUserIdentitiesOperation */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKDiscoverAllUserIdentitiesOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKDiscoverAllUserIdentitiesOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKDiscoverAllUserIdentitiesOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKDiscoverAllUserIdentitiesOperation */

// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverAllUserIdentitiesOperation/discoverAllUserIdentitiesCompletionBlock
func (c_ CKDiscoverAllUserIdentitiesOperation) DiscoverAllUserIdentitiesCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("discoverAllUserIdentitiesCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: discoverAllUserIdentitiesCompletionBlock */


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverAllUserIdentitiesOperation/discoverAllUserIdentitiesCompletionBlock
func (c_ CKDiscoverAllUserIdentitiesOperation) SetDiscoverAllUserIdentitiesCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDiscoverAllUserIdentitiesCompletionBlock:"), value)
}/* debug [instance_properties/setter]: discoverAllUserIdentitiesCompletionBlock */


// The closure to execute for each user identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverAllUserIdentitiesOperation/userIdentityDiscoveredBlock
func (c_ CKDiscoverAllUserIdentitiesOperation) UserIdentityDiscoveredBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("userIdentityDiscoveredBlock"))
	return rv
}/* debug [instance_properties/getter]: userIdentityDiscoveredBlock */


// The closure to execute for each user identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverAllUserIdentitiesOperation/userIdentityDiscoveredBlock
func (c_ CKDiscoverAllUserIdentitiesOperation) SetUserIdentityDiscoveredBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserIdentityDiscoveredBlock:"), value)
}/* debug [instance_properties/setter]: userIdentityDiscoveredBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdiscoveralluseridentitiesoperation/discoveralluseridentitiesresultblock
func (c_ CKDiscoverAllUserIdentitiesOperation) DiscoverAllUserIdentitiesResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("discoverAllUserIdentitiesResultBlock"))
	return rv
}/* debug [instance_properties/getter]: discoverAllUserIdentitiesResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdiscoveralluseridentitiesoperation/discoveralluseridentitiesresultblock
func (c_ CKDiscoverAllUserIdentitiesOperation) SetDiscoverAllUserIdentitiesResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDiscoverAllUserIdentitiesResultBlock:"), value)
}/* debug [instance_properties/setter]: discoverAllUserIdentitiesResultBlock */


// Identifiers that match contacts in the local Contacts database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/contactidentifiers
func (c_ CKDiscoverAllUserIdentitiesOperation) ContactIdentifiers() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("contactIdentifiers"))
	return rv
}/* debug [instance_properties/getter]: contactIdentifiers */


// Identifiers that match contacts in the local Contacts database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/contactidentifiers
func (c_ CKDiscoverAllUserIdentitiesOperation) SetContactIdentifiers(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactIdentifiers:"), value)
}/* debug [instance_properties/setter]: contactIdentifiers */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKDiscoverAllUserIdentitiesOperation) CompletionBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("completionBlock"))
	return rv
}/* debug [instance_properties/getter]: completionBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKDiscoverAllUserIdentitiesOperation) SetCompletionBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}/* debug [instance_properties/setter]: completionBlock */


// The relative amount of importance for granting system resources to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/qualityOfService
func (c_ CKDiscoverAllUserIdentitiesOperation) QualityOfService() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("qualityOfService"))
	return rv
}/* debug [instance_properties/getter]: qualityOfService */


// The relative amount of importance for granting system resources to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/qualityOfService
func (c_ CKDiscoverAllUserIdentitiesOperation) SetQualityOfService(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQualityOfService:"), value)
}/* debug [instance_properties/setter]: qualityOfService */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKDiscoverAllUserIdentitiesOperation */


