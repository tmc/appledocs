// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [RecognizedText] class.
type IRecognizedText interface {
	objectivec.IObject
	

	// properties:
	Confidence() Confidence
	String() foundation.foundation.INSString


	

	// methods:
	BoundingBoxForRangeError(range_ foundation.Range, error_ foundation.foundation.INSError) IRectangleObservation


}





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




















// Calculates the bounding box around the characters in the range of a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedText/boundingBoxForRange:error:
func (r_ RecognizedText) BoundingBoxForRangeError(range_ foundation.Range, error_ foundation.foundation.INSError) IRectangleObservation {
	rv := objc.Send[RectangleObservation](r_.ID, objc.Sel("boundingBoxForRange:error:"), range_, error_)
	return rv
}







// A normalized confidence score for the text recognition result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedText/confidence
func (r_ RecognizedText) Confidence() Confidence {
	rv := objc.Send[Confidence](r_.ID, objc.Sel("confidence"))
	return rv
}


// The top candidate for recognized text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedText/string
func (r_ RecognizedText) String() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("string"))
	return rv
}








