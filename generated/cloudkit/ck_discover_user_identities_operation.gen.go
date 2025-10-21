// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CKDiscoverUserIdentitiesOperation] class.
type ICKDiscoverUserIdentitiesOperation interface {
	ICKOperation
}

// An operation that uses the provided criteria to search for discoverable iCloud users.
//
// Use this operation to discover one or more iCloud users that match identity information you provide, such as email addresses and phone numbers. Before CloudKit can return a user’s identity, you must ask for their permission by calling . Do this as part of any onboarding where you can highlight the benefits of being discoverable within the context of your app. The operation executes the handlers you provide on an internal queue it manages. Your handlers must be capable of executing on a background queue. Tasks that need access to the main queue must redirect as appropriate. The operation calls after it executes and returns results. Use the completion handler to perform housekeeping tasks for the operation. It should also manage any failures, whether due to an error or an explicit cancellation. CloudKit operations have a default QoS of . Operations with this service level are discretionary. The system schedules their execution at an optimal time according to battery level and network conditions, among other factors. Use the property to set a more appropriate QoS for the operation. The following example shows how to create the operation, configure its callbacks, and execute it using the default container’s queue:
//
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

// Alloc allocates a new instance without initialization.
func (cc _CKDiscoverUserIdentitiesOperationClass) Alloc() CKDiscoverUserIdentitiesOperation {
	rv := objc.Send[CKDiscoverUserIdentitiesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates an operation for discovering the user identities of the specified lookup infos.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/init(userIdentityLookupInfos:)
func NewCKDiscoverUserIdentitiesOperationWithUserIdentityLookupInfos(userIdentityLookupInfos unsafe.Pointer) CKDiscoverUserIdentitiesOperation {
	instance := getCKDiscoverUserIdentitiesOperationClass().Alloc()
	rv := objc.Send[CKDiscoverUserIdentitiesOperation](instance.ID, objc.Sel("initWithUserIdentityLookupInfos:"), userIdentityLookupInfos)
	rv.Autorelease()
	return rv
}


// The lookup info for discovering user identities.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/userIdentityLookupInfos
func (c_ CKDiscoverUserIdentitiesOperation) UserIdentityLookupInfos() []CKUserIdentityLookupInfo {
	rv := objc.Send[[]CKUserIdentityLookupInfo](c_.ID, objc.Sel("userIdentityLookupInfos"))
	return rv
}


// SetUserIdentityLookupInfos sets the value of the userIdentityLookupInfos property.
// The lookup info for discovering user identities.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDiscoverUserIdentitiesOperation/userIdentityLookupInfos
func (c_ CKDiscoverUserIdentitiesOperation) SetUserIdentityLookupInfos(value []CKUserIdentityLookupInfo) {
	// Convert Go slice to NSArray
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
}

