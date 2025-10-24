// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNRecognizedText */


/* debug [class_header]: Header for VNRecognizedText */
// The class instance for the [RecognizedText] class.
var (
	RecognizedTextClass     _RecognizedTextClass
	RecognizedTextClassOnce sync.Once
)

func getRecognizedTextClass() _RecognizedTextClass {
	RecognizedTextClassOnce.Do(func() {
		RecognizedTextClass = _RecognizedTextClass{objc.GetClass("VNRecognizedText")}
	})
	return RecognizedTextClass
}

type _RecognizedTextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RecognizedText */
// An interface definition for the [RecognizedText] class.
type IRecognizedText interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RecognizedText */
	// properties:
	Confidence() Confidence /* typedef */
	String() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RecognizedText */
	// methods:
	BoundingBoxForRangeError(range_ corefoundation.Range, error_ objectivec.IObject) IRectangleObservation
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RecognizedText */
// Alloc allocates a new instance without initialization.
func (rc _RecognizedTextClass) Alloc() RecognizedText {
	rv := objc.Send[RecognizedText](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RecognizedTextClass) New() RecognizedText {
	rv := objc.Send[RecognizedText](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizedText) Init() RecognizedText {
	rv := objc.Send[RecognizedText](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizedText) Autorelease() RecognizedText {
	rv := objc.Send[RecognizedText](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizedText creates a new RecognizedText instance.
func NewRecognizedText() RecognizedText {
	return getRecognizedTextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RecognizedText */
// Text recognized in an image through a text recognition request.
//
// A single can contain multiple recognized text objects—one for each candidate.


// Text recognized in an image through a text recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedText
type RecognizedText struct {
	objectivec.Object
}

// RecognizedTextFrom constructs a [RecognizedText] from an unsafe.Pointer.
//
// Text recognized in an image through a text recognition request.
func RecognizedTextFrom(ptr unsafe.Pointer) RecognizedText {
	return RecognizedText{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RecognizedText *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RecognizedText */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RecognizedText */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RecognizedText */

// Calculates the bounding box around the characters in the range of a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedText/boundingBoxForRange:error:
func (r_ RecognizedText) BoundingBoxForRangeError(range_ corefoundation.Range, error_ objectivec.IObject) IRectangleObservation {
	rv := objc.Send[RectangleObservation](r_.ID, objc.Sel("boundingBoxForRange:error:"), range_, error_)
	return rv
}/* debug [instance_methods/method]: BoundingBoxForRangeError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RecognizedText */

// A normalized confidence score for the text recognition result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedText/confidence
func (r_ RecognizedText) Confidence() Confidence /* typedef */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("confidence"))
	return rv
}/* debug [instance_properties/getter]: confidence */


// The top candidate for recognized text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedText/string
func (r_ RecognizedText) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNRecognizedText */



