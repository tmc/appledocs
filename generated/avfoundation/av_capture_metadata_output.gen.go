// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureMetadataOutput */


/* debug [class_header]: Header for AVCaptureMetadataOutput */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureMetadataOutput */
// An interface definition for the [CaptureMetadataOutput] class.
type ICaptureMetadataOutput interface {
	ICaptureOutput
	
/* debug [class_interface_properties]: Properties for CaptureMetadataOutput */
	// properties:
	AvailableMetadataObjectTypes() []string
	MetadataObjectsCallbackQueue() objectivec.IObject
	MetadataObjectsDelegate() unsafe.Pointer
	MetadataObjectTypes() []string
	SetMetadataObjectTypes(value []string)
	RectOfInterest() corefoundation.CGRect
	SetRectOfInterest(value corefoundation.CGRect)
	RequiredMetadataObjectTypesForCinematicVideoCapture() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureMetadataOutput */
	// methods:
	SetMetadataObjectsDelegateQueue(objectsDelegate unsafe.Pointer, objectsCallbackQueue objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureMetadataOutput */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureMetadataOutput */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureMetadataOutput */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureMetadataOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureMetadataOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureMetadataOutput */

// Sets the delegate and dispatch queue to use handle callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/setMetadataObjectsDelegate(_:queue:)
func (c_ CaptureMetadataOutput) SetMetadataObjectsDelegateQueue(objectsDelegate unsafe.Pointer, objectsCallbackQueue objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadataObjectsDelegate:queue:"), objectsDelegate, objectsCallbackQueue)
}/* debug [instance_methods/method]: SetMetadataObjectsDelegateQueue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureMetadataOutput */

// An array of strings identifying the types of metadata objects that can be captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/availableMetadataObjectTypes
func (c_ CaptureMetadataOutput) AvailableMetadataObjectTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableMetadataObjectTypes"))
	return rv
}/* debug [instance_properties/getter]: availableMetadataObjectTypes */


// The dispatch queue on which to execute the delegate’s methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/metadataObjectsCallbackQueue
func (c_ CaptureMetadataOutput) MetadataObjectsCallbackQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("metadataObjectsCallbackQueue"))
	return rv
}/* debug [instance_properties/getter]: metadataObjectsCallbackQueue */


// The delegate of the capture metadata output object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/metadataObjectsDelegate
func (c_ CaptureMetadataOutput) MetadataObjectsDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("metadataObjectsDelegate"))
	return rv
}/* debug [instance_properties/getter]: metadataObjectsDelegate */


// An array of strings identifying the types of metadata objects to process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/metadataObjectTypes
func (c_ CaptureMetadataOutput) MetadataObjectTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("metadataObjectTypes"))
	return rv
}/* debug [instance_properties/getter]: metadataObjectTypes */


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
}/* debug [instance_properties/setter]: metadataObjectTypes */


// A rectangle of interest for limiting the search area for visual metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/rectOfInterest
func (c_ CaptureMetadataOutput) RectOfInterest() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("rectOfInterest"))
	return rv
}/* debug [instance_properties/getter]: rectOfInterest */


// A rectangle of interest for limiting the search area for visual metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/rectOfInterest
func (c_ CaptureMetadataOutput) SetRectOfInterest(value corefoundation.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRectOfInterest:"), value)
}/* debug [instance_properties/setter]: rectOfInterest */


// The required metadata object types when Cinematic Video capture is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/requiredMetadataObjectTypesForCinematicVideoCapture
func (c_ CaptureMetadataOutput) RequiredMetadataObjectTypesForCinematicVideoCapture() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("requiredMetadataObjectTypesForCinematicVideoCapture"))
	return rv
}/* debug [instance_properties/getter]: requiredMetadataObjectTypesForCinematicVideoCapture */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureMetadataOutput */


