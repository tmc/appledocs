// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CKOperation] class.
var (
	CKOperationClass     _CKOperationClass
	CKOperationClassOnce sync.Once
)

func getCKOperationClass() _CKOperationClass {
	CKOperationClassOnce.Do(func() {
		CKOperationClass = _CKOperationClass{objc.GetClass("CKOperation")}
	})
	return CKOperationClass
}

type _CKOperationClass struct {
	class objc.Class
}

// An interface definition for the [CKOperation] class.
type ICKOperation interface {
	IOperation
	AllowsCellularAccess() bool
	SetAllowsCellularAccess(value bool)
	Configuration() CKOperationConfiguration
	SetConfiguration(value ICKOperationConfiguration)
	Container() CKContainer
	SetContainer(value ICKContainer)
	Group() CKOperationGroup
	SetGroup(value ICKOperationGroup)
	LongLived() bool
	SetLongLived(value bool)
	LongLivedOperationWasPersistedBlock() unsafe.Pointer
	SetLongLivedOperationWasPersistedBlock(value unsafe.Pointer)
	OperationID() unsafe.Pointer
	TimeoutIntervalForRequest() foundation.TimeInterval
	SetTimeoutIntervalForRequest(value foundation.ITimeInterval)
	TimeoutIntervalForResource() foundation.TimeInterval
	SetTimeoutIntervalForResource(value foundation.ITimeInterval)
	IsLongLived() bool
	SetIsLongLived(value bool)
	QualityOfService() unsafe.Pointer
	SetQualityOfService(value unsafe.Pointer)
}

// The abstract base class for all operations that execute in a database.
//
// All CloudKit operations descend from , which provides the infrastructure for executing tasks in one of your app’s containers. Don’t subclass or create instances of this class directly. Instead, create instances of one of its concrete subclasses. Use the properties of this class to configure the behavior of the operation before submitting it to a queue or executing it directly. CloudKit operations involve communicating with the iCloud servers to send and receive data. You can use the properties of this class to configure the behavior of those network requests to ensure the best performance for your app.


// The abstract base class for all operations that execute in a database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation
type CKOperation struct {
	Operation
}

// CKOperationFrom constructs a [CKOperation] from an unsafe.Pointer.
//
// The abstract base class for all operations that execute in a database.
func CKOperationFrom(ptr unsafe.Pointer) CKOperation {
	return CKOperation{
		Operation: OperationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKOperationClass) Alloc() CKOperation {
	rv := objc.Send[CKOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKOperationClass) New() CKOperation {
	rv := objc.Send[CKOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKOperation) Init() CKOperation {
	rv := objc.Send[CKOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKOperation) Autorelease() CKOperation {
	rv := objc.Send[CKOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKOperation creates a new CKOperation instance.
func NewCKOperation() CKOperation {
	return getCKOperationClass().New()
}




// A Boolean value that indicates whether the operation can send data over the cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/allowsCellularAccess
func (c_ CKOperation) AllowsCellularAccess() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}


// A Boolean value that indicates whether the operation can send data over the cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/allowsCellularAccess
func (c_ CKOperation) SetAllowsCellularAccess(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsCellularAccess:"), value)
}


// The operation’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/configuration-swift.property
func (c_ CKOperation) Configuration() CKOperationConfiguration {
	rv := objc.Send[CKOperationConfiguration](c_.ID, objc.Sel("configuration"))
	return rv
}


// The operation’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/configuration-swift.property
func (c_ CKOperation) SetConfiguration(value ICKOperationConfiguration) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfiguration:"), value)
}


// The operation’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/container
func (c_ CKOperation) Container() CKContainer {
	rv := objc.Send[CKContainer](c_.ID, objc.Sel("container"))
	return rv
}


// The operation’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/container
func (c_ CKOperation) SetContainer(value ICKContainer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainer:"), value)
}


// The operation’s group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/group
func (c_ CKOperation) Group() CKOperationGroup {
	rv := objc.Send[CKOperationGroup](c_.ID, objc.Sel("group"))
	return rv
}


// The operation’s group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/group
func (c_ CKOperation) SetGroup(value ICKOperationGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroup:"), value)
}


// A Boolean value that indicates whether the operation is long-lived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/isLongLived
func (c_ CKOperation) LongLived() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("longLived"))
	return rv
}


// A Boolean value that indicates whether the operation is long-lived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/isLongLived
func (c_ CKOperation) SetLongLived(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongLived:"), value)
}


// The closure to execute when the server begins to store callbacks for the long-lived operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/longLivedOperationWasPersistedBlock
func (c_ CKOperation) LongLivedOperationWasPersistedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("longLivedOperationWasPersistedBlock"))
	return rv
}


// The closure to execute when the server begins to store callbacks for the long-lived operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/longLivedOperationWasPersistedBlock
func (c_ CKOperation) SetLongLivedOperationWasPersistedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongLivedOperationWasPersistedBlock:"), value)
}


// A unique identifier for a long-lived operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/operationID-3eujz
func (c_ CKOperation) OperationID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("operationID"))
	return rv
}


// The timeout interval when waiting for additional data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/timeoutIntervalForRequest
func (c_ CKOperation) TimeoutIntervalForRequest() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](c_.ID, objc.Sel("timeoutIntervalForRequest"))
	return rv
}


// The timeout interval when waiting for additional data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/timeoutIntervalForRequest
func (c_ CKOperation) SetTimeoutIntervalForRequest(value foundation.ITimeInterval) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeoutIntervalForRequest:"), value)
}


// The maximum amount of time that a resource request can use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/timeoutIntervalForResource
func (c_ CKOperation) TimeoutIntervalForResource() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](c_.ID, objc.Sel("timeoutIntervalForResource"))
	return rv
}


// The maximum amount of time that a resource request can use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperation/timeoutIntervalForResource
func (c_ CKOperation) SetTimeoutIntervalForResource(value foundation.ITimeInterval) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeoutIntervalForResource:"), value)
}


// A Boolean value that indicates whether the operation is long-lived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/islonglived
func (c_ CKOperation) IsLongLived() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLongLived"))
	return rv
}


// A Boolean value that indicates whether the operation is long-lived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/islonglived
func (c_ CKOperation) SetIsLongLived(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLongLived:"), value)
}


// The relative amount of importance for granting system resources to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/qualityOfService
func (c_ CKOperation) QualityOfService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("qualityOfService"))
	return rv
}


// The relative amount of importance for granting system resources to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/qualityOfService
func (c_ CKOperation) SetQualityOfService(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQualityOfService:"), value)
}


