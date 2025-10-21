// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterContentAppObserver] class.
var (
	MTRClusterContentAppObserverClass     _MTRClusterContentAppObserverClass
	MTRClusterContentAppObserverClassOnce sync.Once
)

func getMTRClusterContentAppObserverClass() _MTRClusterContentAppObserverClass {
	MTRClusterContentAppObserverClassOnce.Do(func() {
		MTRClusterContentAppObserverClass = _MTRClusterContentAppObserverClass{objc.GetClass("MTRClusterContentAppObserver")}
	})
	return MTRClusterContentAppObserverClass
}

type _MTRClusterContentAppObserverClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterContentAppObserver] class.
type IMTRClusterContentAppObserver interface {
	IMTRGenericCluster
	ContentAppMessageWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRContentAppObserverClusterContentAppMessageParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
}

// Cluster Content App Observer This cluster provides an interface for sending targeted commands to an Observer of a Content App on a Video Player device such as a Streaming Media Player, Smart TV or Smart Screen. The cluster server for Content App Observer is implemented by an endpoint that communicates with a Content App, such as a Casting Video Client. The cluster client for Content App Observer is implemented by a Content App endpoint. A Content App is informed of the NodeId of an Observer when a binding is set on the Content App. The Content App can then send the ContentAppMessage to the Observer (server cluster), and the Observer responds with a ContentAppMessageResponse.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterContentAppObserver
type MTRClusterContentAppObserver struct {
	MTRGenericCluster
}

// MTRClusterContentAppObserverFrom constructs a [MTRClusterContentAppObserver] from an unsafe.Pointer.
//
// Cluster Content App Observer This cluster provides an interface for sending targeted commands to an Observer of a Content App on a Video Player device such as a Streaming Media Player, Smart TV or Smart Screen. The cluster server for Content App Observer is implemented by an endpoint that communicates with a Content App, such as a Casting Video Client. The cluster client for Content App Observer is implemented by a Content App endpoint. A Content App is informed of the NodeId of an Observer when a binding is set on the Content App. The Content App can then send the ContentAppMessage to the Observer (server cluster), and the Observer responds with a ContentAppMessageResponse.
func MTRClusterContentAppObserverFrom(ptr unsafe.Pointer) MTRClusterContentAppObserver {
	return MTRClusterContentAppObserver{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterContentAppObserverClass) Alloc() MTRClusterContentAppObserver {
	rv := objc.Send[MTRClusterContentAppObserver](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterContentAppObserverClass) New() MTRClusterContentAppObserver {
	rv := objc.Send[MTRClusterContentAppObserver](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterContentAppObserver) Init() MTRClusterContentAppObserver {
	rv := objc.Send[MTRClusterContentAppObserver](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterContentAppObserver) Autorelease() MTRClusterContentAppObserver {
	rv := objc.Send[MTRClusterContentAppObserver](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterContentAppObserver creates a new MTRClusterContentAppObserver instance.
func NewMTRClusterContentAppObserver() MTRClusterContentAppObserver {
	return getMTRClusterContentAppObserverClass().New()
}




// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterContentAppObserver/init(device:endpointID:queue:)
func NewMTRClusterContentAppObserverWithDeviceEndpointIDQueue(device IMTRDevice, endpointID foundation.INumber, queue unsafe.Pointer) MTRClusterContentAppObserver {
	instance := getMTRClusterContentAppObserverClass().Alloc()
	rv := objc.Send[MTRClusterContentAppObserver](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterContentAppObserver/contentAppMessage(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterContentAppObserver) ContentAppMessageWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRContentAppObserverClusterContentAppMessageParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("contentAppMessageWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterContentAppObserver/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterContentAppObserver) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterContentAppObserver/readAttributeAttributeList(with:)
func (m_ MTRClusterContentAppObserver) ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterContentAppObserver/readAttributeClusterRevision(with:)
func (m_ MTRClusterContentAppObserver) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterContentAppObserver/readAttributeFeatureMap(with:)
func (m_ MTRClusterContentAppObserver) ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterContentAppObserver/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterContentAppObserver) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}


