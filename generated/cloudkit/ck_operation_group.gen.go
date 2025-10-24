// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKOperationGroup */


/* debug [class_header]: Header for CKOperationGroup */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKOperationGroup */
// An interface definition for the [CKOperationGroup] class.
type ICKOperationGroup interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKOperationGroup */
	// properties:
	DefaultConfiguration() ICKOperationConfiguration
	SetDefaultConfiguration(value ICKOperationConfiguration)
	ExpectedReceiveSize() CKOperationGroupTransferSize
	SetExpectedReceiveSize(value CKOperationGroupTransferSize)
	ExpectedSendSize() CKOperationGroupTransferSize
	SetExpectedSendSize(value CKOperationGroupTransferSize)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	OperationGroupID() objc.IObject /* cross-framework: NSString */
	Quantity() uint
	SetQuantity(value uint)
	Group() ICKOperationGroup
	SetGroup(value ICKOperationGroup)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKOperationGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKOperationGroup */
// Alloc allocates a new instance without initialization.
func (cc _CKOperationGroupClass) Alloc() CKOperationGroup {
	rv := objc.Send[CKOperationGroup](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKOperationGroup */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKOperationGroup */

// Creates an operation group from a serialized instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/init(coder:)
func NewCKOperationGroupWithCoder(aDecoder foundation.Coder) CKOperationGroup {
	instance := getCKOperationGroupClass().Alloc()
	rv := objc.Send[CKOperationGroup](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKOperationGroupWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKOperationGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKOperationGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKOperationGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKOperationGroup */

// The default configuration for operations in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/defaultConfiguration
func (c_ CKOperationGroup) DefaultConfiguration() ICKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](c_.ID, objc.Sel("defaultConfiguration"))
	return rv
}/* debug [instance_properties/getter]: defaultConfiguration */


// The default configuration for operations in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/defaultConfiguration
func (c_ CKOperationGroup) SetDefaultConfiguration(value ICKOperationConfiguration) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDefaultConfiguration:"), value)
}/* debug [instance_properties/setter]: defaultConfiguration */


// The estimated size of traffic to download from CloudKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/expectedReceiveSize
func (c_ CKOperationGroup) ExpectedReceiveSize() CKOperationGroupTransferSize {
	rv := objc.Send[CKOperationGroupTransferSize](c_.ID, objc.Sel("expectedReceiveSize"))
	return rv
}/* debug [instance_properties/getter]: expectedReceiveSize */


// The estimated size of traffic to download from CloudKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/expectedReceiveSize
func (c_ CKOperationGroup) SetExpectedReceiveSize(value CKOperationGroupTransferSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExpectedReceiveSize:"), value)
}/* debug [instance_properties/setter]: expectedReceiveSize */


// The estimated size of traffic to upload to CloudKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/expectedSendSize
func (c_ CKOperationGroup) ExpectedSendSize() CKOperationGroupTransferSize {
	rv := objc.Send[CKOperationGroupTransferSize](c_.ID, objc.Sel("expectedSendSize"))
	return rv
}/* debug [instance_properties/getter]: expectedSendSize */


// The estimated size of traffic to upload to CloudKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/expectedSendSize
func (c_ CKOperationGroup) SetExpectedSendSize(value CKOperationGroupTransferSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExpectedSendSize:"), value)
}/* debug [instance_properties/setter]: expectedSendSize */


// The operation group’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/name
func (c_ CKOperationGroup) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The operation group’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/name
func (c_ CKOperationGroup) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// The operation group’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/operationGroupID
func (c_ CKOperationGroup) OperationGroupID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("operationGroupID"))
	return rv
}/* debug [instance_properties/getter]: operationGroupID */


// The number of operations in the operation group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/quantity
func (c_ CKOperationGroup) Quantity() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("quantity"))
	return rv
}/* debug [instance_properties/getter]: quantity */


// The number of operations in the operation group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/quantity
func (c_ CKOperationGroup) SetQuantity(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQuantity:"), value)
}/* debug [instance_properties/setter]: quantity */


// The operation’s group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/group
func (c_ CKOperationGroup) Group() ICKOperationGroup {
	rv := objc.Send[CKOperationGroup](c_.ID, objc.Sel("group"))
	return rv
}/* debug [instance_properties/getter]: group */


// The operation’s group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/group
func (c_ CKOperationGroup) SetGroup(value ICKOperationGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroup:"), value)
}/* debug [instance_properties/setter]: group */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKOperationGroup */


