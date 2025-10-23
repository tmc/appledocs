// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CKOperationConfiguration] class.
type ICKOperationConfiguration interface {
	objectivec.IObject
	// properties:
	QualityOfService() QualityOfService /* not a class type */
	SetQualityOfService(value QualityOfService /* not a class type */)
	AllowsCellularAccess() bool /* primitive/slice/pointer. */
	SetAllowsCellularAccess(value bool /* primitive/slice/pointer. */)
	Container() ICKContainer
	SetContainer(value ICKContainer)
	IsLongLived() bool /* primitive/slice/pointer. */
	SetIsLongLived(value bool /* primitive/slice/pointer. */)
	TimeoutIntervalForRequest() unsafe.Pointer
	SetTimeoutIntervalForRequest(value unsafe.Pointer)
	TimeoutIntervalForResource() unsafe.Pointer
	SetTimeoutIntervalForResource(value unsafe.Pointer)
	Configuration() ICKOperationConfiguration
	SetConfiguration(value ICKOperationConfiguration)
	Group() ICKOperationGroup
	SetGroup(value ICKOperationGroup)
	LongLivedOperationWasPersistedBlock() unsafe.Pointer
	SetLongLivedOperationWasPersistedBlock(value unsafe.Pointer)
	DefaultConfiguration() ICKOperationConfiguration
	SetDefaultConfiguration(value ICKOperationConfiguration)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (cc _CKOperationConfigurationClass) Alloc() CKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The priority that the system uses when it allocates resources to the operations that use this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/qualityOfService
func (c_ CKOperationConfiguration) QualityOfService() QualityOfService /* not a class type */ {
	rv := objc.Send[QualityOfService](c_.ID, objc.Sel("qualityOfService"))
	return rv
}


// The priority that the system uses when it allocates resources to the operations that use this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/qualityOfService
func (c_ CKOperationConfiguration) SetQualityOfService(value QualityOfService /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQualityOfService:"), value)
}


// A Boolean value that indicates whether operations that use this configuration can send data over the cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/allowscellularaccess
func (c_ CKOperationConfiguration) AllowsCellularAccess() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}


// A Boolean value that indicates whether operations that use this configuration can send data over the cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/allowscellularaccess
func (c_ CKOperationConfiguration) SetAllowsCellularAccess(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsCellularAccess:"), value)
}


// The configuration’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/container
func (c_ CKOperationConfiguration) Container() ICKContainer {
	rv := objc.Send[CKContainer](c_.ID, objc.Sel("container"))
	return rv
}


// The configuration’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/container
func (c_ CKOperationConfiguration) SetContainer(value ICKContainer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainer:"), value)
}


// A Boolean value that indicates whether the operations that use this configuration are long-lived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/islonglived
func (c_ CKOperationConfiguration) IsLongLived() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLongLived"))
	return rv
}


// A Boolean value that indicates whether the operations that use this configuration are long-lived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/islonglived
func (c_ CKOperationConfiguration) SetIsLongLived(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLongLived:"), value)
}


// The maximum amount of time that a request can take.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/timeoutintervalforrequest
func (c_ CKOperationConfiguration) TimeoutIntervalForRequest() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("timeoutIntervalForRequest"))
	return rv
}


// The maximum amount of time that a request can take.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/timeoutintervalforrequest
func (c_ CKOperationConfiguration) SetTimeoutIntervalForRequest(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeoutIntervalForRequest:"), value)
}


// The maximum amount of time that a resource request can take.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/timeoutintervalforresource
func (c_ CKOperationConfiguration) TimeoutIntervalForResource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("timeoutIntervalForResource"))
	return rv
}


// The maximum amount of time that a resource request can take.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/timeoutintervalforresource
func (c_ CKOperationConfiguration) SetTimeoutIntervalForResource(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeoutIntervalForResource:"), value)
}


// The operation’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.property
func (c_ CKOperationConfiguration) Configuration() ICKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](c_.ID, objc.Sel("configuration"))
	return rv
}


// The operation’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.property
func (c_ CKOperationConfiguration) SetConfiguration(value ICKOperationConfiguration) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfiguration:"), value)
}


// The operation’s group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/group
func (c_ CKOperationConfiguration) Group() ICKOperationGroup {
	rv := objc.Send[CKOperationGroup](c_.ID, objc.Sel("group"))
	return rv
}


// The operation’s group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/group
func (c_ CKOperationConfiguration) SetGroup(value ICKOperationGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroup:"), value)
}


// The closure to execute when the server begins to store callbacks for the long-lived operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/longlivedoperationwaspersistedblock
func (c_ CKOperationConfiguration) LongLivedOperationWasPersistedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("longLivedOperationWasPersistedBlock"))
	return rv
}


// The closure to execute when the server begins to store callbacks for the long-lived operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/longlivedoperationwaspersistedblock
func (c_ CKOperationConfiguration) SetLongLivedOperationWasPersistedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongLivedOperationWasPersistedBlock:"), value)
}


// The default configuration for operations in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/defaultconfiguration
func (c_ CKOperationConfiguration) DefaultConfiguration() ICKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](c_.ID, objc.Sel("defaultConfiguration"))
	return rv
}


// The default configuration for operations in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/defaultconfiguration
func (c_ CKOperationConfiguration) SetDefaultConfiguration(value ICKOperationConfiguration) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDefaultConfiguration:"), value)
}



