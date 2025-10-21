// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterMessages] class.
var (
	MTRClusterMessagesClass     _MTRClusterMessagesClass
	MTRClusterMessagesClassOnce sync.Once
)

func getMTRClusterMessagesClass() _MTRClusterMessagesClass {
	MTRClusterMessagesClassOnce.Do(func() {
		MTRClusterMessagesClass = _MTRClusterMessagesClass{objc.GetClass("MTRClusterMessages")}
	})
	return MTRClusterMessagesClass
}

type _MTRClusterMessagesClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterMessages] class.
type IMTRClusterMessages interface {
	IMTRGenericCluster
	CancelMessagesRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	PresentMessagesRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeActiveMessageIDsWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeMessagesWithParams(params unsafe.Pointer) unsafe.Pointer
}

// Cluster Messages This cluster provides an interface for passing messages to be presented by a device.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMessages
type MTRClusterMessages struct {
	MTRGenericCluster
}

// MTRClusterMessagesFrom constructs a [MTRClusterMessages] from an unsafe.Pointer.
//
// Cluster Messages This cluster provides an interface for passing messages to be presented by a device.
func MTRClusterMessagesFrom(ptr unsafe.Pointer) MTRClusterMessages {
	return MTRClusterMessages{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterMessagesClass) Alloc() MTRClusterMessages {
	rv := objc.Send[MTRClusterMessages](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterMessagesClass) New() MTRClusterMessages {
	rv := objc.Send[MTRClusterMessages](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterMessages) Init() MTRClusterMessages {
	rv := objc.Send[MTRClusterMessages](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterMessages) Autorelease() MTRClusterMessages {
	rv := objc.Send[MTRClusterMessages](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterMessages creates a new MTRClusterMessages instance.
func NewMTRClusterMessages() MTRClusterMessages {
	return getMTRClusterMessagesClass().New()
}


// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMessages/init(device:endpointID:queue:)
func NewMTRClusterMessagesWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRClusterMessages {
	instance := getMTRClusterMessagesClass().Alloc()
	rv := objc.Send[MTRClusterMessages](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMessages/cancelRequest(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterMessages) CancelMessagesRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelMessagesRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMessages/presentRequest(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterMessages) PresentMessagesRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("presentMessagesRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMessages/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterMessages) ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMessages/readAttributeActiveMessageIDs(with:)
func (m_ MTRClusterMessages) ReadAttributeActiveMessageIDsWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeActiveMessageIDsWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMessages/readAttributeAttributeList(with:)
func (m_ MTRClusterMessages) ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMessages/readAttributeClusterRevision(with:)
func (m_ MTRClusterMessages) ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMessages/readAttributeFeatureMap(with:)
func (m_ MTRClusterMessages) ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMessages/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterMessages) ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMessages/readAttributeMessages(with:)
func (m_ MTRClusterMessages) ReadAttributeMessagesWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeMessagesWithParams:"), params)
	return rv
}


