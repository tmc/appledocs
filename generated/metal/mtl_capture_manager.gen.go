// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptureManager] class.
var (
	CaptureManagerClass     _CaptureManagerClass
	CaptureManagerClassOnce sync.Once
)

func getCaptureManagerClass() _CaptureManagerClass {
	CaptureManagerClassOnce.Do(func() {
		CaptureManagerClass = _CaptureManagerClass{objc.GetClass("MTLCaptureManager")}
	})
	return CaptureManagerClass
}

type _CaptureManagerClass struct {
	class objc.Class
}





// An interface definition for the [CaptureManager] class.
type ICaptureManager interface {
	objectivec.IObject
	

	// properties:
	DefaultCaptureScope() unsafe.Pointer
	SetDefaultCaptureScope(value unsafe.Pointer)
	IsCapturing() bool


	

	// methods:
	NewCaptureScopeWithCommandQueue(commandQueue unsafe.Pointer) unsafe.Pointer
	NewCaptureScopeWithMTL4CommandQueue(commandQueue unsafe.Pointer) unsafe.Pointer
	NewCaptureScopeWithDevice(device unsafe.Pointer) unsafe.Pointer
	StartCaptureWithDescriptorError(descriptor IMTLCaptureDescriptor, error_ foundation.foundation.INSError) bool
	StopCapture()
	SupportsDestination(destination CaptureDestination) bool


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureManagerClass) Alloc() CaptureManager {
	rv := objc.Send[CaptureManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureManagerClass) New() CaptureManager {
	rv := objc.Send[CaptureManager](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureManager) Init() CaptureManager {
	rv := objc.Send[CaptureManager](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureManager) Autorelease() CaptureManager {
	rv := objc.Send[CaptureManager](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureManager creates a new CaptureManager instance.
func NewCaptureManager() CaptureManager {
	return getCaptureManagerClass().New()
}





// An instance you use to capture Metal command data in your app.
//
// A capture manager works with the frame capture feature to: Capture data about Metal commands programmatically. See . Only capture commands that apply to a specific , command queue, or instance. Assign a default instance for captures you create in Xcode by clicking the Capture GPU workload button in the debug bar, which has an icon with the Metal logo. The Metal debugger requires you to enable GPU Frame Capture in your project settings; see . For more information about Metal frame capture, see .


// An instance you use to capture Metal command data in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureManager
type CaptureManager struct {
	objectivec.Object
}

// CaptureManagerFrom constructs a [CaptureManager] from an unsafe.Pointer.
//
// An instance you use to capture Metal command data in your app.
func CaptureManagerFrom(ptr unsafe.Pointer) CaptureManager {
	return CaptureManager{objectivec.Object{objc.ID(ptr)}}
}











// Provides the shared capture manager for your Metal app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureManager/shared()
func (cc _CaptureManagerClass) SharedCaptureManager() ICaptureManager {
	rv := objc.Send[CaptureManager](objc.ID(cc.class), objc.Sel("sharedCaptureManager"))
	return rv
}












// Creates a capture scope for commands submitted to a specific command queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureManager/makeCaptureScope(commandQueue:)-1rozd
func (c_ CaptureManager) NewCaptureScopeWithCommandQueue(commandQueue unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("newCaptureScopeWithCommandQueue:"), commandQueue)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureManager/makeCaptureScope(commandQueue:)-9wie3
func (c_ CaptureManager) NewCaptureScopeWithMTL4CommandQueue(commandQueue unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("newCaptureScopeWithMTL4CommandQueue:"), commandQueue)
	return rv
}


// Creates a capture scope for commands submitted to a specific device object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureManager/makeCaptureScope(device:)
func (c_ CaptureManager) NewCaptureScopeWithDevice(device unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("newCaptureScopeWithDevice:"), device)
	return rv
}


// Starts capturing any of your app’s Metal commands, with the capture session defined by a descriptor object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureManager/startCapture(with:)
func (c_ CaptureManager) StartCaptureWithDescriptorError(descriptor IMTLCaptureDescriptor, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("startCaptureWithDescriptor:error:"), descriptor, error_)
	return rv
}


// Stops capturing Metal commands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureManager/stopCapture()
func (c_ CaptureManager) StopCapture() {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopCapture"))
}


// Checks to see whether a particular capture destination is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureManager/supportsDestination(_:)
func (c_ CaptureManager) SupportsDestination(destination CaptureDestination) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsDestination:"), destination)
	return rv
}







// The capture scope to use when a capture is initiated in Xcode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureManager/defaultCaptureScope
func (c_ CaptureManager) DefaultCaptureScope() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("defaultCaptureScope"))
	return rv
}


// The capture scope to use when a capture is initiated in Xcode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureManager/defaultCaptureScope
func (c_ CaptureManager) SetDefaultCaptureScope(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDefaultCaptureScope:"), value)
}


// A Boolean value that indicates whether Metal commands are being captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureManager/isCapturing
func (c_ CaptureManager) IsCapturing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCapturing"))
	return rv
}







