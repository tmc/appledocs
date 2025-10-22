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
	DefaultCaptureScope() unsafe.Pointer
	SetDefaultCaptureScope(value unsafe.Pointer)
	IsCapturing() bool
	SetIsCapturing(value bool)
}

// An instance you use to capture Metal command data in your app.
//
// A capture manager works with the frame capture feature to: Capture data about Metal commands programmatically. See . Only capture commands that apply to a specific , command queue, or instance. Assign a default instance for captures you create in Xcode by clicking the Capture GPU workload button in the debug bar, which has an icon with the Metal logo. The Metal debugger requires you to enable GPU Frame Capture in your project settings; see . For more information about Metal frame capture, see .
//
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

// Alloc allocates a new instance without initialization.
func (cc _CaptureManagerClass) Alloc() CaptureManager {
	rv := objc.Send[CaptureManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The capture scope to use when a capture is initiated in Xcode.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/defaultcapturescope
func (c_ CaptureManager) DefaultCaptureScope() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("defaultCaptureScope"))
	return rv
}


// SetDefaultCaptureScope sets the value of the defaultCaptureScope property.
// The capture scope to use when a capture is initiated in Xcode.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/defaultcapturescope
func (c_ CaptureManager) SetDefaultCaptureScope(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDefaultCaptureScope:"), value)
}

// A Boolean value that indicates whether Metal commands are being captured.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/iscapturing
func (c_ CaptureManager) IsCapturing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCapturing"))
	return rv
}


// SetIsCapturing sets the value of the isCapturing property.
// A Boolean value that indicates whether Metal commands are being captured.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturemanager/iscapturing
func (c_ CaptureManager) SetIsCapturing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCapturing:"), value)
}


