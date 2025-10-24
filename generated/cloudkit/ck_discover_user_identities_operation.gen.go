// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKDiscoverUserIdentitiesOperation */


/* debug [class_header]: Header for CKDiscoverUserIdentitiesOperation */
// The class instance for the [CKDiscoverUserIdentitiesOperation] class.
var (
	CKDiscoverUserIdentitiesOperationClass     _CKDiscoverUserIdentitiesOperationClass
	CKDiscoverUserIdentitiesOperationClassOnce sync.Once
)

func getCKDiscoverUserIdentitiesOperationClass() _CKDiscoverUserIdentitiesOperationClass {
	CKDiscoverUserIdentitiesOperationClassOnce.Do(func() {
		CKDiscoverUserIdentitiesOperationClass = _CKDiscoverUserIdentitiesOperationClass{objc.GetClass("CKDiscoverUserIdentitiesOperation")}
	})
	return CKDiscoverUserIdentitiesOperationClass
}

type _CKDiscoverUserIdentitiesOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKDiscoverUserIdentitiesOperation */
// An interface definition for the [CKDiscoverUserIdentitiesOperation] class.
type ICKDiscoverUserIdentitiesOperation interface {
	ICKOperation
	
/* debug [class_interface_properties]: Properties for CKDiscoverUserIdentitiesOperation */
	// properties:
	DiscoverUserIdentitiesCompletionBlock() unsafe.Pointer
	SetDiscoverUserIdentitiesCompletionBlock(value unsafe.Pointer)
	UserIdentityDiscoveredBlock() unsafe.Pointer
	SetUserIdentityDiscoveredBlock(value unsafe.Pointer)
	UserIdentityLookupInfos() []CKUserIdentityLookupInfo
	SetUserIdentityLookupInfos(value []CKUserIdentityLookupInfo)
	DiscoverUserIdentitiesResultBlock() objectivec.IObject
	SetDiscoverUserIdentitiesResultBlock(value objectivec.IObject)
	CompletionBlock() objectivec.IObject
	SetCompletionBlock(value objectivec.IObject)
	QualityOfService() objectivec.IObject
	SetQualityOfService(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKDiscoverUserIdentitiesOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKDiscoverUserIdentitiesOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKDiscoverUserIdentitiesOperationClass) Alloc() CKDiscoverUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverUserIdentitiesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKDiscoverUserIdentitiesOperationClass) New() CKDiscoverUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverUserIdentitiesOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKDiscoverUserIdentitiesOperation) Init() CKDiscoverUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverUserIdentitiesOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKDiscoverUserIdentitiesOperation) Autorelease() CKDiscoverUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverUserIdentitiesOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKDiscoverUserIdentitiesOperation creates a new CKDiscoverUserIdentitiesOperation instance.
func NewCKDiscoverUserIdentitiesOperation() CKDiscoverUserIdentitiesOperation {
	return getCKDiscoverUserIdentitiesOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKDiscoverUserIdentitiesOperation */
// An operation that uses the provided criteria to search for discoverable iCloud users.
//
// Use this operation to discover one or more iCloud users that match identity information you provide, such as email addresses and phone numbers. Before CloudKit can return a user’s identity, you must ask for their permission by calling . Do this as part of any onboarding where you can highlight the benefits of being discoverable within the context of your app. The operation executes the handlers you provide on an internal queue it manages. Your handlers must be capable of executing on a background queue. Tasks that need access to the main queue must redirect as appropriate. The operation calls after it executes and returns results. Use the completion handler to perform housekeeping tasks for the operation. It should also manage any failures, whether due to an error or an explicit cancellation. CloudKit operations have a default QoS of . Operations with this service level are discretionary. The system schedules their execution at an optimal time according to battery level and network conditions, among other factors. Use the property to set a more appropriate QoS for the operation. The following example shows how to create the operation, configure its callbacks, and execute it using the default container’s queue:


// An operation that uses the provided criteria to search for discoverable iCloud users.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation
type CKDiscoverUserIdentitiesOperation struct {
	CKOperation
}

// CKDiscoverUserIdentitiesOperationFrom constructs a [CKDiscoverUserIdentitiesOperation] from an unsafe.Pointer.
//
// An operation that uses the provided criteria to search for discoverable iCloud users.
func CKDiscoverUserIdentitiesOperationFrom(ptr unsafe.Pointer) CKDiscoverUserIdentitiesOperation {
	return CKDiscoverUserIdentitiesOperation{
		CKOperation: CKOperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKDiscoverUserIdentitiesOperation */

// Creates an operation for discovering the user identities of the specified lookup infos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/init(userIdentityLookupInfos:)
func NewCKDiscoverUserIdentitiesOperationWithUserIdentityLookupInfos(userIdentityLookupInfos []CKUserIdentityLookupInfo) CKDiscoverUserIdentitiesOperation {
	instance := getCKDiscoverUserIdentitiesOperationClass().Alloc()
	rv := objc.Send[CKDiscoverUserIdentitiesOperation](instance.ID, objc.Sel("initWithUserIdentityLookupInfos:"), userIdentityLookupInfos)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKDiscoverUserIdentitiesOperationWithUserIdentityLookupInfos */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKDiscoverUserIdentitiesOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKDiscoverUserIdentitiesOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKDiscoverUserIdentitiesOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKDiscoverUserIdentitiesOperation */

// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/discoverUserIdentitiesCompletionBlock
func (c_ CKDiscoverUserIdentitiesOperation) DiscoverUserIdentitiesCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("discoverUserIdentitiesCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: discoverUserIdentitiesCompletionBlock */


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/discoverUserIdentitiesCompletionBlock
func (c_ CKDiscoverUserIdentitiesOperation) SetDiscoverUserIdentitiesCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDiscoverUserIdentitiesCompletionBlock:"), value)
}/* debug [instance_properties/setter]: discoverUserIdentitiesCompletionBlock */


// The closure to execute for each user identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/userIdentityDiscoveredBlock
func (c_ CKDiscoverUserIdentitiesOperation) UserIdentityDiscoveredBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("userIdentityDiscoveredBlock"))
	return rv
}/* debug [instance_properties/getter]: userIdentityDiscoveredBlock */


// The closure to execute for each user identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/userIdentityDiscoveredBlock
func (c_ CKDiscoverUserIdentitiesOperation) SetUserIdentityDiscoveredBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserIdentityDiscoveredBlock:"), value)
}/* debug [instance_properties/setter]: userIdentityDiscoveredBlock */


// The lookup info for discovering user identities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/userIdentityLookupInfos
func (c_ CKDiscoverUserIdentitiesOperation) UserIdentityLookupInfos() []CKUserIdentityLookupInfo {
	rv := objc.Send[[]CKUserIdentityLookupInfo](c_.ID, objc.Sel("userIdentityLookupInfos"))
	return rv
}/* debug [instance_properties/getter]: userIdentityLookupInfos */


// The lookup info for discovering user identities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/userIdentityLookupInfos
func (c_ CKDiscoverUserIdentitiesOperation) SetUserIdentityLookupInfos(value []CKUserIdentityLookupInfo) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserIdentityLookupInfos:"), nsArray)
}/* debug [instance_properties/setter]: userIdentityLookupInfos */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdiscoveruseridentitiesoperation/discoveruseridentitiesresultblock
func (c_ CKDiscoverUserIdentitiesOperation) DiscoverUserIdentitiesResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("discoverUserIdentitiesResultBlock"))
	return rv
}/* debug [instance_properties/getter]: discoverUserIdentitiesResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdiscoveruseridentitiesoperation/discoveruseridentitiesresultblock
func (c_ CKDiscoverUserIdentitiesOperation) SetDiscoverUserIdentitiesResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDiscoverUserIdentitiesResultBlock:"), value)
}/* debug [instance_properties/setter]: discoverUserIdentitiesResultBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKDiscoverUserIdentitiesOperation) CompletionBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("completionBlock"))
	return rv
}/* debug [instance_properties/getter]: completionBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKDiscoverUserIdentitiesOperation) SetCompletionBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}/* debug [instance_properties/setter]: completionBlock */


// The relative amount of importance for granting system resources to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/qualityOfService
func (c_ CKDiscoverUserIdentitiesOperation) QualityOfService() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("qualityOfService"))
	return rv
}/* debug [instance_properties/getter]: qualityOfService */


// The relative amount of importance for granting system resources to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/qualityOfService
func (c_ CKDiscoverUserIdentitiesOperation) SetQualityOfService(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQualityOfService:"), value)
}/* debug [instance_properties/setter]: qualityOfService */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKDiscoverUserIdentitiesOperation */


