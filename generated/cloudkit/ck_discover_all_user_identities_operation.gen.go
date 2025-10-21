// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CKDiscoverAllUserIdentitiesOperation] class.
type ICKDiscoverAllUserIdentitiesOperation interface {
	ICKOperation
}

// An operation that uses the device’s contacts to search for discoverable iCloud users.
//
// Use this operation to discover iCloud users that match entries in the device’s Contacts database. CloudKit uses the email addresses and phone numbers in each Contact record to search for a matching iCloud account. Although your app doesn’t need authorization to use the Contacts database to execute this operation, if it has authorization, you can use the property on any returned user identity to fetch the corresponding Contact record from the database. Before CloudKit can return a user’s identity, you must ask for their permission by calling . Do this as part of any onboarding where you can highlight the benefits of being discoverable within the context of your app. The operation executes the handlers you provide on an internal queue it manages. Your handlers must be capable of executing on a background queue. Tasks that need access to the main queue must redirect as appropriate. The operation calls after it executes and returns results. Use the completion handler to perform housekeeping tasks for the operation. It should also manage any failures, whether due to an error or an explicit cancellation. CloudKit operations have a default QoS of . Operations with this service level are discretionary. The system schedules their execution at an optimal time according to battery level and network conditions, among other factors. Use the property to set a more appropriate QoS for the operation. The following example shows how to create the operation, configure its callbacks, and execute it using the default container’s queue:
//
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

// Alloc allocates a new instance without initialization.
func (cc _CKDiscoverAllUserIdentitiesOperationClass) Alloc() CKDiscoverAllUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverAllUserIdentitiesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdiscoveralluseridentitiesoperation/discoveralluseridentitiesresultblock
func (c_ CKDiscoverAllUserIdentitiesOperation) DiscoverAllUserIdentitiesResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("discoverAllUserIdentitiesResultBlock"))
	return rv
}


// SetDiscoverAllUserIdentitiesResultBlock sets the value of the discoverAllUserIdentitiesResultBlock property.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdiscoveralluseridentitiesoperation/discoveralluseridentitiesresultblock
func (c_ CKDiscoverAllUserIdentitiesOperation) SetDiscoverAllUserIdentitiesResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDiscoverAllUserIdentitiesResultBlock:"), value)
}

// The relative amount of importance for granting system resources to the operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/qualityOfService
func (c_ CKDiscoverAllUserIdentitiesOperation) QualityOfService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("qualityOfService"))
	return rv
}


// SetQualityOfService sets the value of the qualityOfService property.
// The relative amount of importance for granting system resources to the operation.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/qualityOfService
func (c_ CKDiscoverAllUserIdentitiesOperation) SetQualityOfService(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQualityOfService:"), value)
}

// The block to execute after the operation’s main task is completed.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKDiscoverAllUserIdentitiesOperation) CompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("completionBlock"))
	return rv
}


// SetCompletionBlock sets the value of the completionBlock property.
// The block to execute after the operation’s main task is completed.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKDiscoverAllUserIdentitiesOperation) SetCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}

// Identifiers that match contacts in the local Contacts database.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/contactidentifiers
func (c_ CKDiscoverAllUserIdentitiesOperation) ContactIdentifiers() string {
	rv := objc.Send[string](c_.ID, objc.Sel("contactIdentifiers"))
	return rv
}


// SetContactIdentifiers sets the value of the contactIdentifiers property.
// Identifiers that match contacts in the local Contacts database.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/contactidentifiers
func (c_ CKDiscoverAllUserIdentitiesOperation) SetContactIdentifiers(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactIdentifiers:"), objc.String(value))
}


