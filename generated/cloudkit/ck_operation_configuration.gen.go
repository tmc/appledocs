// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKOperationConfiguration */


/* debug [class_header]: Header for CKOperationConfiguration */
// The class instance for the [CKOperationConfiguration] class.
var (
	CKOperationConfigurationClass     _CKOperationConfigurationClass
	CKOperationConfigurationClassOnce sync.Once
)

func getCKOperationConfigurationClass() _CKOperationConfigurationClass {
	CKOperationConfigurationClassOnce.Do(func() {
		CKOperationConfigurationClass = _CKOperationConfigurationClass{objc.GetClass("CKOperationConfiguration")}
	})
	return CKOperationConfigurationClass
}

type _CKOperationConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKOperationConfiguration */
// An interface definition for the [CKOperationConfiguration] class.
type ICKOperationConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKOperationConfiguration */
	// properties:
	AllowsCellularAccess() bool
	SetAllowsCellularAccess(value bool)
	Container() ICKContainer
	SetContainer(value ICKContainer)
	LongLived() bool
	SetLongLived(value bool)
	QualityOfService() QualityOfService /* not a class type */
	SetQualityOfService(value QualityOfService /* not a class type */)
	TimeoutIntervalForRequest() float64
	SetTimeoutIntervalForRequest(value float64)
	TimeoutIntervalForResource() float64
	SetTimeoutIntervalForResource(value float64)
	IsLongLived() bool
	SetIsLongLived(value bool)
	Configuration() ICKOperationConfiguration
	SetConfiguration(value ICKOperationConfiguration)
	Group() ICKOperationGroup
	SetGroup(value ICKOperationGroup)
	LongLivedOperationWasPersistedBlock() objectivec.IObject
	SetLongLivedOperationWasPersistedBlock(value objectivec.IObject)
	DefaultConfiguration() ICKOperationConfiguration
	SetDefaultConfiguration(value ICKOperationConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKOperationConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKOperationConfiguration */
// Alloc allocates a new instance without initialization.
func (cc _CKOperationConfigurationClass) Alloc() CKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKOperationConfigurationClass) New() CKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKOperationConfiguration) Init() CKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKOperationConfiguration) Autorelease() CKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKOperationConfiguration creates a new CKOperationConfiguration instance.
func NewCKOperationConfiguration() CKOperationConfiguration {
	return getCKOperationConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKOperationConfiguration */
// An object that describes how a CloudKit operation behaves.
//
// All of the properties in have a default value. When determining which properties to apply to a CloudKit operation, consult the operation’s configuration property, as well as the property of the group that the operation belongs to. These properties combine through the following rules:


// An object that describes how a CloudKit operation behaves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class
type CKOperationConfiguration struct {
	objectivec.Object
}

// CKOperationConfigurationFrom constructs a [CKOperationConfiguration] from an unsafe.Pointer.
//
// An object that describes how a CloudKit operation behaves.
func CKOperationConfigurationFrom(ptr unsafe.Pointer) CKOperationConfiguration {
	return CKOperationConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKOperationConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKOperationConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKOperationConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKOperationConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKOperationConfiguration */

// A Boolean value that indicates whether operations that use this configuration can send data over the cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/allowsCellularAccess
func (c_ CKOperationConfiguration) AllowsCellularAccess() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}/* debug [instance_properties/getter]: allowsCellularAccess */


// A Boolean value that indicates whether operations that use this configuration can send data over the cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/allowsCellularAccess
func (c_ CKOperationConfiguration) SetAllowsCellularAccess(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsCellularAccess:"), value)
}/* debug [instance_properties/setter]: allowsCellularAccess */


// The configuration’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/container
func (c_ CKOperationConfiguration) Container() ICKContainer {
	rv := objc.Send[CKContainer](c_.ID, objc.Sel("container"))
	return rv
}/* debug [instance_properties/getter]: container */


// The configuration’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/container
func (c_ CKOperationConfiguration) SetContainer(value ICKContainer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainer:"), value)
}/* debug [instance_properties/setter]: container */


// A Boolean value that indicates whether the operations that use this configuration are long-lived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/isLongLived
func (c_ CKOperationConfiguration) LongLived() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("longLived"))
	return rv
}/* debug [instance_properties/getter]: longLived */


// A Boolean value that indicates whether the operations that use this configuration are long-lived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/isLongLived
func (c_ CKOperationConfiguration) SetLongLived(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongLived:"), value)
}/* debug [instance_properties/setter]: longLived */


// The priority that the system uses when it allocates resources to the operations that use this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/qualityOfService
func (c_ CKOperationConfiguration) QualityOfService() QualityOfService /* not a class type */ {
	rv := objc.Send[QualityOfService](c_.ID, objc.Sel("qualityOfService"))
	return rv
}/* debug [instance_properties/getter]: qualityOfService */


// The priority that the system uses when it allocates resources to the operations that use this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/qualityOfService
func (c_ CKOperationConfiguration) SetQualityOfService(value QualityOfService /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQualityOfService:"), value)
}/* debug [instance_properties/setter]: qualityOfService */


// The maximum amount of time that a request can take.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/timeoutIntervalForRequest
func (c_ CKOperationConfiguration) TimeoutIntervalForRequest() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("timeoutIntervalForRequest"))
	return rv
}/* debug [instance_properties/getter]: timeoutIntervalForRequest */


// The maximum amount of time that a request can take.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/timeoutIntervalForRequest
func (c_ CKOperationConfiguration) SetTimeoutIntervalForRequest(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeoutIntervalForRequest:"), value)
}/* debug [instance_properties/setter]: timeoutIntervalForRequest */


// The maximum amount of time that a resource request can take.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/timeoutIntervalForResource
func (c_ CKOperationConfiguration) TimeoutIntervalForResource() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("timeoutIntervalForResource"))
	return rv
}/* debug [instance_properties/getter]: timeoutIntervalForResource */


// The maximum amount of time that a resource request can take.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/timeoutIntervalForResource
func (c_ CKOperationConfiguration) SetTimeoutIntervalForResource(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeoutIntervalForResource:"), value)
}/* debug [instance_properties/setter]: timeoutIntervalForResource */


// A Boolean value that indicates whether the operations that use this configuration are long-lived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/islonglived
func (c_ CKOperationConfiguration) IsLongLived() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLongLived"))
	return rv
}/* debug [instance_properties/getter]: isLongLived */


// A Boolean value that indicates whether the operations that use this configuration are long-lived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/islonglived
func (c_ CKOperationConfiguration) SetIsLongLived(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLongLived:"), value)
}/* debug [instance_properties/setter]: isLongLived */


// The operation’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.property
func (c_ CKOperationConfiguration) Configuration() ICKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](c_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// The operation’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.property
func (c_ CKOperationConfiguration) SetConfiguration(value ICKOperationConfiguration) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfiguration:"), value)
}/* debug [instance_properties/setter]: configuration */


// The operation’s group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/group
func (c_ CKOperationConfiguration) Group() ICKOperationGroup {
	rv := objc.Send[CKOperationGroup](c_.ID, objc.Sel("group"))
	return rv
}/* debug [instance_properties/getter]: group */


// The operation’s group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/group
func (c_ CKOperationConfiguration) SetGroup(value ICKOperationGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroup:"), value)
}/* debug [instance_properties/setter]: group */


// The closure to execute when the server begins to store callbacks for the long-lived operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/longlivedoperationwaspersistedblock
func (c_ CKOperationConfiguration) LongLivedOperationWasPersistedBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("longLivedOperationWasPersistedBlock"))
	return rv
}/* debug [instance_properties/getter]: longLivedOperationWasPersistedBlock */


// The closure to execute when the server begins to store callbacks for the long-lived operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/longlivedoperationwaspersistedblock
func (c_ CKOperationConfiguration) SetLongLivedOperationWasPersistedBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongLivedOperationWasPersistedBlock:"), value)
}/* debug [instance_properties/setter]: longLivedOperationWasPersistedBlock */


// The default configuration for operations in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/defaultconfiguration
func (c_ CKOperationConfiguration) DefaultConfiguration() ICKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](c_.ID, objc.Sel("defaultConfiguration"))
	return rv
}/* debug [instance_properties/getter]: defaultConfiguration */


// The default configuration for operations in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/defaultconfiguration
func (c_ CKOperationConfiguration) SetDefaultConfiguration(value ICKOperationConfiguration) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDefaultConfiguration:"), value)
}/* debug [instance_properties/setter]: defaultConfiguration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKOperationConfiguration */



