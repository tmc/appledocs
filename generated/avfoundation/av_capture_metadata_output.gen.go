// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	AvailableMetadataObjectTypes() []string
	MetadataObjectsCallbackQueue() objectivec.IObject
	MetadataObjectsDelegate() unsafe.Pointer
	MetadataObjectTypes() []string
	SetMetadataObjectTypes(value []string)
	RectOfInterest() corefoundation.CGRect
	SetRectOfInterest(value corefoundation.CGRect)
	RequiredMetadataObjectTypesForCinematicVideoCapture() []string


	

	// methods:
	SetMetadataObjectsDelegateQueue(objectsDelegate unsafe.Pointer, objectsCallbackQueue objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureMetadataOutputClass) Alloc() CaptureMetadataOutput {
	rv := objc.Send[CaptureMetadataOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A capture output for processing timed metadata produced by a capture session.
//
// An object intercepts metadata objects emitted by its associated capture connection and forwards them to a delegate object for processing. You can use instances of this class to process specific types of metadata included with the input data. You use this class the way you do other output objects, typically by adding it as an output to an object.


// A capture output for processing timed metadata produced by a capture session.
//
// [Full Topic]
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





















// Sets the delegate and dispatch queue to use handle callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/setMetadataObjectsDelegate(_:queue:)
func (c_ CaptureMetadataOutput) SetMetadataObjectsDelegateQueue(objectsDelegate unsafe.Pointer, objectsCallbackQueue objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadataObjectsDelegate:queue:"), objectsDelegate, objectsCallbackQueue)
}







// An array of strings identifying the types of metadata objects that can be captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/availableMetadataObjectTypes
func (c_ CaptureMetadataOutput) AvailableMetadataObjectTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableMetadataObjectTypes"))
	return rv
}


// The dispatch queue on which to execute the delegate’s methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/metadataObjectsCallbackQueue
func (c_ CaptureMetadataOutput) MetadataObjectsCallbackQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("metadataObjectsCallbackQueue"))
	return rv
}


// The delegate of the capture metadata output object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/metadataObjectsDelegate
func (c_ CaptureMetadataOutput) MetadataObjectsDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("metadataObjectsDelegate"))
	return rv
}


// An array of strings identifying the types of metadata objects to process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/metadataObjectTypes
func (c_ CaptureMetadataOutput) MetadataObjectTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("metadataObjectTypes"))
	return rv
}


// An array of strings identifying the types of metadata objects to process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/metadataObjectTypes
func (c_ CaptureMetadataOutput) SetMetadataObjectTypes(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadataObjectTypes:"), nsArray)
}


// A rectangle of interest for limiting the search area for visual metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/rectOfInterest
func (c_ CaptureMetadataOutput) RectOfInterest() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("rectOfInterest"))
	return rv
}


// A rectangle of interest for limiting the search area for visual metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/rectOfInterest
func (c_ CaptureMetadataOutput) SetRectOfInterest(value corefoundation.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRectOfInterest:"), value)
}


// The required metadata object types when Cinematic Video capture is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/requiredMetadataObjectTypesForCinematicVideoCapture
func (c_ CaptureMetadataOutput) RequiredMetadataObjectTypesForCinematicVideoCapture() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("requiredMetadataObjectTypesForCinematicVideoCapture"))
	return rv
}







