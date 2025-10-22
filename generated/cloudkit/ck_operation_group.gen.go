// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKOperationGroup] class.
var (
	CKOperationGroupClass     _CKOperationGroupClass
	CKOperationGroupClassOnce sync.Once
)

func getCKOperationGroupClass() _CKOperationGroupClass {
	CKOperationGroupClassOnce.Do(func() {
		CKOperationGroupClass = _CKOperationGroupClass{objc.GetClass("CKOperationGroup")}
	})
	return CKOperationGroupClass
}

type _CKOperationGroupClass struct {
	class objc.Class
}

// An interface definition for the [CKOperationGroup] class.
type ICKOperationGroup interface {
	objectivec.IObject
	DefaultConfiguration() CKOperationConfiguration
	SetDefaultConfiguration(value ICKOperationConfiguration)
	ExpectedReceiveSize() CKOperationGroupTransferSize
	SetExpectedReceiveSize(value ICKOperationGroupTransferSize)
	ExpectedSendSize() CKOperationGroupTransferSize
	SetExpectedSendSize(value ICKOperationGroupTransferSize)
	Name() string
	SetName(value string)
	OperationGroupID() string
	Quantity() uint
	SetQuantity(value uint)
	Group() CKOperationGroup
	SetGroup(value ICKOperationGroup)
}

// An explicit association between two or more operations.
//
// In certain situations, you might want to perform several CloudKit operations together. Grouping operations in CloudKit doesn’t ensure atomicity. For example, when building a Calendar app, you group the following actions: Fetch records from CloudKit, which consists of numerous queries that fetch both new records and records with changes. Perform incremental fetches of records in response to a push notification. Update several records when the user saves a calendar event. Associate operation groups with operations by setting their property. Create a new operation group for each distinct user interaction.


// An explicit association between two or more operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup

type CKOperationGroup struct {
	objectivec.Object
}

// CKOperationGroupFrom constructs a [CKOperationGroup] from an unsafe.Pointer.
//
// An explicit association between two or more operations.
func CKOperationGroupFrom(ptr unsafe.Pointer) CKOperationGroup {
	return CKOperationGroup{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKOperationGroupClass) Alloc() CKOperationGroup {
	rv := objc.Send[CKOperationGroup](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKOperationGroupClass) New() CKOperationGroup {
	rv := objc.Send[CKOperationGroup](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKOperationGroup) Init() CKOperationGroup {
	rv := objc.Send[CKOperationGroup](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKOperationGroup) Autorelease() CKOperationGroup {
	rv := objc.Send[CKOperationGroup](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKOperationGroup creates a new CKOperationGroup instance.
func NewCKOperationGroup() CKOperationGroup {
	return getCKOperationGroupClass().New()
}




// Creates an operation group from a serialized instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/init(coder:)

func NewCKOperationGroupWithCoder(aDecoder foundation.ICoder) CKOperationGroup {
	instance := getCKOperationGroupClass().Alloc()
	rv := objc.Send[CKOperationGroup](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}



// The default configuration for operations in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/defaultConfiguration

func (c_ CKOperationGroup) DefaultConfiguration() CKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](c_.ID, objc.Sel("defaultConfiguration"))
	return rv
}


// The default configuration for operations in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/defaultConfiguration

func (c_ CKOperationGroup) SetDefaultConfiguration(value ICKOperationConfiguration) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDefaultConfiguration:"), value)
}


// The estimated size of traffic to download from CloudKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/expectedReceiveSize

func (c_ CKOperationGroup) ExpectedReceiveSize() CKOperationGroupTransferSize {
	rv := objc.Send[CKOperationGroupTransferSize](c_.ID, objc.Sel("expectedReceiveSize"))
	return rv
}


// The estimated size of traffic to download from CloudKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/expectedReceiveSize

func (c_ CKOperationGroup) SetExpectedReceiveSize(value ICKOperationGroupTransferSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExpectedReceiveSize:"), value)
}


// The estimated size of traffic to upload to CloudKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/expectedSendSize

func (c_ CKOperationGroup) ExpectedSendSize() CKOperationGroupTransferSize {
	rv := objc.Send[CKOperationGroupTransferSize](c_.ID, objc.Sel("expectedSendSize"))
	return rv
}


// The estimated size of traffic to upload to CloudKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/expectedSendSize

func (c_ CKOperationGroup) SetExpectedSendSize(value ICKOperationGroupTransferSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExpectedSendSize:"), value)
}


// The operation group’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/name

func (c_ CKOperationGroup) Name() string {
	rv := objc.Send[string](c_.ID, objc.Sel("name"))
	return rv
}


// The operation group’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/name

func (c_ CKOperationGroup) SetName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), objc.String(value))
}


// The operation group’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/operationGroupID

func (c_ CKOperationGroup) OperationGroupID() string {
	rv := objc.Send[string](c_.ID, objc.Sel("operationGroupID"))
	return rv
}


// The number of operations in the operation group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/quantity

func (c_ CKOperationGroup) Quantity() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("quantity"))
	return rv
}


// The number of operations in the operation group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/quantity

func (c_ CKOperationGroup) SetQuantity(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQuantity:"), value)
}


// The operation’s group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/group

func (c_ CKOperationGroup) Group() CKOperationGroup {
	rv := objc.Send[CKOperationGroup](c_.ID, objc.Sel("group"))
	return rv
}


// The operation’s group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/group

func (c_ CKOperationGroup) SetGroup(value ICKOperationGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroup:"), value)
}


