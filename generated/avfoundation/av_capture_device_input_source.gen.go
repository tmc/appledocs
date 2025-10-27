// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptureDeviceInputSource] class.
var (
	CaptureDeviceInputSourceClass     _CaptureDeviceInputSourceClass
	CaptureDeviceInputSourceClassOnce sync.Once
)

func getCaptureDeviceInputSourceClass() _CaptureDeviceInputSourceClass {
	CaptureDeviceInputSourceClassOnce.Do(func() {
		CaptureDeviceInputSourceClass = _CaptureDeviceInputSourceClass{objc.GetClass("AVCaptureDeviceInputSource")}
	})
	return CaptureDeviceInputSourceClass
}

type _CaptureDeviceInputSourceClass struct {
	class objc.Class
}





// An interface definition for the [CaptureDeviceInputSource] class.
type ICaptureDeviceInputSource interface {
	objectivec.IObject
	

	// properties:
	InputSourceID() foundation.foundation.INSString
	LocalizedName() foundation.foundation.INSString
	ActiveInputSource() IAVCaptureDeviceInputSource
	SetActiveInputSource(value IAVCaptureDeviceInputSource)
	InputSources() IAVCaptureDeviceInputSource
	SetInputSources(value IAVCaptureDeviceInputSource)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureDeviceInputSourceClass) Alloc() CaptureDeviceInputSource {
	rv := objc.Send[CaptureDeviceInputSource](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureDeviceInputSourceClass) New() CaptureDeviceInputSource {
	rv := objc.Send[CaptureDeviceInputSource](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDeviceInputSource) Init() CaptureDeviceInputSource {
	rv := objc.Send[CaptureDeviceInputSource](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDeviceInputSource) Autorelease() CaptureDeviceInputSource {
	rv := objc.Send[CaptureDeviceInputSource](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDeviceInputSource creates a new CaptureDeviceInputSource instance.
func NewCaptureDeviceInputSource() CaptureDeviceInputSource {
	return getCaptureDeviceInputSourceClass().New()
}





// A distinct input source on a capture device.
//
// A capture device may optionally present an array of input sources that represent distinct mutually exclusive inputs to the device. For example, an audio capture device might have ADAT optical and analog input sources; a video capture device might have an HDMI or component input source.


// A distinct input source on a capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/InputSource
type CaptureDeviceInputSource struct {
	objectivec.Object
}

// CaptureDeviceInputSourceFrom constructs a [CaptureDeviceInputSource] from an unsafe.Pointer.
//
// A distinct input source on a capture device.
func CaptureDeviceInputSourceFrom(ptr unsafe.Pointer) CaptureDeviceInputSource {
	return CaptureDeviceInputSource{objectivec.Object{objc.ID(ptr)}}
}

























// An identifier for an input source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/InputSource/inputSourceID
func (c_ CaptureDeviceInputSource) InputSourceID() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("inputSourceID"))
	return rv
}


// A localized, human-readable name for the input source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/InputSource/localizedName
func (c_ CaptureDeviceInputSource) LocalizedName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedName"))
	return rv
}


// The currently active input source of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeinputsource
func (c_ CaptureDeviceInputSource) ActiveInputSource() IAVCaptureDeviceInputSource {
	rv := objc.Send[CaptureDeviceInputSource](c_.ID, objc.Sel("activeInputSource"))
	return rv
}


// The currently active input source of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeinputsource
func (c_ CaptureDeviceInputSource) SetActiveInputSource(value IAVCaptureDeviceInputSource) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveInputSource:"), value)
}


// An array of input sources that the device supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/inputsources
func (c_ CaptureDeviceInputSource) InputSources() IAVCaptureDeviceInputSource {
	rv := objc.Send[CaptureDeviceInputSource](c_.ID, objc.Sel("inputSources"))
	return rv
}


// An array of input sources that the device supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/inputsources
func (c_ CaptureDeviceInputSource) SetInputSources(value IAVCaptureDeviceInputSource) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInputSources:"), value)
}








