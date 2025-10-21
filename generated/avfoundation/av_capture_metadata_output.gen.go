// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [CaptureMetadataOutput] class.
var (
	CaptureMetadataOutputClass     _CaptureMetadataOutputClass
	CaptureMetadataOutputClassOnce sync.Once
)

func getCaptureMetadataOutputClass() _CaptureMetadataOutputClass {
	CaptureMetadataOutputClassOnce.Do(func() {
		CaptureMetadataOutputClass = _CaptureMetadataOutputClass{objc.GetClass("AVCaptureMetadataOutput")}
	})
	return CaptureMetadataOutputClass
}

type _CaptureMetadataOutputClass struct {
	class objc.Class
}

// An interface definition for the [CaptureMetadataOutput] class.
type ICaptureMetadataOutput interface {
	ICaptureOutput
}

// A capture output for processing timed metadata produced by a capture session.
//
// An object intercepts metadata objects emitted by its associated capture connection and forwards them to a delegate object for processing. You can use instances of this class to process specific types of metadata included with the input data. You use this class the way you do other output objects, typically by adding it as an output to an object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput
type CaptureMetadataOutput struct {
	CaptureOutput
}

// CaptureMetadataOutputFrom constructs a [CaptureMetadataOutput] from an unsafe.Pointer.
//
// A capture output for processing timed metadata produced by a capture session.
func CaptureMetadataOutputFrom(ptr unsafe.Pointer) CaptureMetadataOutput {
	return CaptureMetadataOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureMetadataOutputClass) Alloc() CaptureMetadataOutput {
	rv := objc.Send[CaptureMetadataOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureMetadataOutputClass) New() CaptureMetadataOutput {
	rv := objc.Send[CaptureMetadataOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureMetadataOutput) Init() CaptureMetadataOutput {
	rv := objc.Send[CaptureMetadataOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureMetadataOutput) Autorelease() CaptureMetadataOutput {
	rv := objc.Send[CaptureMetadataOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureMetadataOutput creates a new CaptureMetadataOutput instance.
func NewCaptureMetadataOutput() CaptureMetadataOutput {
	return getCaptureMetadataOutputClass().New()
}


// An array of strings identifying the types of metadata objects that can be captured.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemetadataoutput/availablemetadataobjecttypes
func (c_ CaptureMetadataOutput) AvailableMetadataObjectTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableMetadataObjectTypes"))
	return rv
}


// SetAvailableMetadataObjectTypes sets the value of the availableMetadataObjectTypes property.
// An array of strings identifying the types of metadata objects that can be captured.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemetadataoutput/availablemetadataobjecttypes
func (c_ CaptureMetadataOutput) SetAvailableMetadataObjectTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableMetadataObjectTypes:"), value)
}

// An array of strings identifying the types of metadata objects to process.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemetadataoutput/metadataobjecttypes
func (c_ CaptureMetadataOutput) MetadataObjectTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("metadataObjectTypes"))
	return rv
}


// SetMetadataObjectTypes sets the value of the metadataObjectTypes property.
// An array of strings identifying the types of metadata objects to process.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemetadataoutput/metadataobjecttypes
func (c_ CaptureMetadataOutput) SetMetadataObjectTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadataObjectTypes:"), value)
}

// The dispatch queue on which to execute the delegate’s methods.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemetadataoutput/metadataobjectscallbackqueue
func (c_ CaptureMetadataOutput) MetadataObjectsCallbackQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("metadataObjectsCallbackQueue"))
	return rv
}


// SetMetadataObjectsCallbackQueue sets the value of the metadataObjectsCallbackQueue property.
// The dispatch queue on which to execute the delegate’s methods.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemetadataoutput/metadataobjectscallbackqueue
func (c_ CaptureMetadataOutput) SetMetadataObjectsCallbackQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadataObjectsCallbackQueue:"), value)
}

// The delegate of the capture metadata output object.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemetadataoutput/metadataobjectsdelegate
func (c_ CaptureMetadataOutput) MetadataObjectsDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("metadataObjectsDelegate"))
	return rv
}


// SetMetadataObjectsDelegate sets the value of the metadataObjectsDelegate property.
// The delegate of the capture metadata output object.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemetadataoutput/metadataobjectsdelegate
func (c_ CaptureMetadataOutput) SetMetadataObjectsDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadataObjectsDelegate:"), value)
}

// A rectangle of interest for limiting the search area for visual metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemetadataoutput/rectofinterest
func (c_ CaptureMetadataOutput) RectOfInterest() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("rectOfInterest"))
	return rv
}


// SetRectOfInterest sets the value of the rectOfInterest property.
// A rectangle of interest for limiting the search area for visual metadata.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemetadataoutput/rectofinterest
func (c_ CaptureMetadataOutput) SetRectOfInterest(value coregraphics.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRectOfInterest:"), value)
}

// The required metadata object types when Cinematic Video capture is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemetadataoutput/requiredmetadataobjecttypesforcinematicvideocapture
func (c_ CaptureMetadataOutput) RequiredMetadataObjectTypesForCinematicVideoCapture() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("requiredMetadataObjectTypesForCinematicVideoCapture"))
	return rv
}


// SetRequiredMetadataObjectTypesForCinematicVideoCapture sets the value of the requiredMetadataObjectTypesForCinematicVideoCapture property.
// The required metadata object types when Cinematic Video capture is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemetadataoutput/requiredmetadataobjecttypesforcinematicvideocapture
func (c_ CaptureMetadataOutput) SetRequiredMetadataObjectTypesForCinematicVideoCapture(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequiredMetadataObjectTypesForCinematicVideoCapture:"), value)
}



