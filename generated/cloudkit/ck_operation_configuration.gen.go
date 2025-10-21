// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// An object that describes how a CloudKit operation behaves.
//
// All of the properties in have a default value. When determining which properties to apply to a CloudKit operation, consult the operation’s configuration property, as well as the property of the group that the operation belongs to. These properties combine through the following rules:
//
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


// The closure to execute when the server begins to store callbacks for the long-lived operation.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/longlivedoperationwaspersistedblock
func (c_ CKOperationConfiguration) LongLivedOperationWasPersistedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("longLivedOperationWasPersistedBlock"))
	return rv
}


// SetLongLivedOperationWasPersistedBlock sets the value of the longLivedOperationWasPersistedBlock property.
// The closure to execute when the server begins to store callbacks for the long-lived operation.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/longlivedoperationwaspersistedblock
func (c_ CKOperationConfiguration) SetLongLivedOperationWasPersistedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongLivedOperationWasPersistedBlock:"), value)
}

// The operation’s configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.property
func (c_ CKOperationConfiguration) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("configuration"))
	return rv
}


// SetConfiguration sets the value of the configuration property.
// The operation’s configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.property
func (c_ CKOperationConfiguration) SetConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfiguration:"), value)
}

// The operation’s group.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/group
func (c_ CKOperationConfiguration) Group() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("group"))
	return rv
}


// SetGroup sets the value of the group property.
// The operation’s group.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/group
func (c_ CKOperationConfiguration) SetGroup(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroup:"), value)
}

// The default configuration for operations in the group.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/defaultconfiguration
func (c_ CKOperationConfiguration) DefaultConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("defaultConfiguration"))
	return rv
}


// SetDefaultConfiguration sets the value of the defaultConfiguration property.
// The default configuration for operations in the group.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperationgroup/defaultconfiguration
func (c_ CKOperationConfiguration) SetDefaultConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDefaultConfiguration:"), value)
}

// A Boolean value that indicates whether the operations that use this configuration are long-lived.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/islonglived
func (c_ CKOperationConfiguration) IsLongLived() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLongLived"))
	return rv
}


// SetIsLongLived sets the value of the isLongLived property.
// A Boolean value that indicates whether the operations that use this configuration are long-lived.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/islonglived
func (c_ CKOperationConfiguration) SetIsLongLived(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLongLived:"), value)
}

// A Boolean value that indicates whether operations that use this configuration can send data over the cellular network.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/allowsCellularAccess
func (c_ CKOperationConfiguration) AllowsCellularAccess() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}


// SetAllowsCellularAccess sets the value of the allowsCellularAccess property.
// A Boolean value that indicates whether operations that use this configuration can send data over the cellular network.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/allowsCellularAccess
func (c_ CKOperationConfiguration) SetAllowsCellularAccess(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsCellularAccess:"), value)
}

// The configuration’s container.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/container
func (c_ CKOperationConfiguration) Container() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("container"))
	return rv
}


// SetContainer sets the value of the container property.
// The configuration’s container.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/container
func (c_ CKOperationConfiguration) SetContainer(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainer:"), value)
}

// A Boolean value that indicates whether the operations that use this configuration are long-lived.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/isLongLived
func (c_ CKOperationConfiguration) LongLived() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("longLived"))
	return rv
}


// SetLongLived sets the value of the longLived property.
// A Boolean value that indicates whether the operations that use this configuration are long-lived.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/isLongLived
func (c_ CKOperationConfiguration) SetLongLived(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongLived:"), value)
}

// The priority that the system uses when it allocates resources to the operations that use this configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/qualityOfService
func (c_ CKOperationConfiguration) QualityOfService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("qualityOfService"))
	return rv
}


// SetQualityOfService sets the value of the qualityOfService property.
// The priority that the system uses when it allocates resources to the operations that use this configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/qualityOfService
func (c_ CKOperationConfiguration) SetQualityOfService(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQualityOfService:"), value)
}

// The maximum amount of time that a request can take.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/timeoutIntervalForRequest
func (c_ CKOperationConfiguration) TimeoutIntervalForRequest() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](c_.ID, objc.Sel("timeoutIntervalForRequest"))
	return rv
}


// SetTimeoutIntervalForRequest sets the value of the timeoutIntervalForRequest property.
// The maximum amount of time that a request can take.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/timeoutIntervalForRequest
func (c_ CKOperationConfiguration) SetTimeoutIntervalForRequest(value foundation.TimeInterval) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeoutIntervalForRequest:"), value)
}

// The maximum amount of time that a resource request can take.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/timeoutIntervalForResource
func (c_ CKOperationConfiguration) TimeoutIntervalForResource() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](c_.ID, objc.Sel("timeoutIntervalForResource"))
	return rv
}


// SetTimeoutIntervalForResource sets the value of the timeoutIntervalForResource property.
// The maximum amount of time that a resource request can take.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/Configuration-swift.class/timeoutIntervalForResource
func (c_ CKOperationConfiguration) SetTimeoutIntervalForResource(value foundation.TimeInterval) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeoutIntervalForResource:"), value)
}



