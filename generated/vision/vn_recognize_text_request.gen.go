// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [RecognizeTextRequest] class.
var (
	RecognizeTextRequestClass     _RecognizeTextRequestClass
	RecognizeTextRequestClassOnce sync.Once
)

func getRecognizeTextRequestClass() _RecognizeTextRequestClass {
	RecognizeTextRequestClassOnce.Do(func() {
		RecognizeTextRequestClass = _RecognizeTextRequestClass{objc.GetClass("VNRecognizeTextRequest")}
	})
	return RecognizeTextRequestClass
}

type _RecognizeTextRequestClass struct {
	class objc.Class
}

// An interface definition for the [RecognizeTextRequest] class.
type IRecognizeTextRequest interface {
	IImageBasedRequest
	// properties:
	AutomaticallyDetectsLanguage() bool
	SetAutomaticallyDetectsLanguage(value bool)
	CustomWords() objc.IObject /* cross-framework: NSString */
	SetCustomWords(value objc.IObject /* cross-framework: NSString */)
	MinimumTextHeight() float32
	SetMinimumTextHeight(value float32)
	RecognitionLanguages() objc.IObject /* cross-framework: NSString */
	SetRecognitionLanguages(value objc.IObject /* cross-framework: NSString */)
	RecognitionLevel() RequestTextRecognitionLevel /* not a class type */
	SetRecognitionLevel(value RequestTextRecognitionLevel /* not a class type */)
	Results() IVNRecognizedTextObservation
	SetResults(value IVNRecognizedTextObservation)
	UsesLanguageCorrection() bool
	SetUsesLanguageCorrection(value bool)
	VNRecognizeTextRequestRevision1() int
	VNRecognizeTextRequestRevision2() int
	VNRecognizeTextRequestRevision3() int
	// methods:
}

// An image-analysis request that finds and recognizes text in an image.
//
// By default, a text recognition request first locates all possible glyphs or characters in the input image, and then analyzes each string. To specify or limit the languages to find in the request, set the property to an array that contains the names of the languages of text you want to recognize. Vision returns the result of this request in a object.


// An image-analysis request that finds and recognizes text in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest
type RecognizeTextRequest struct {
	ImageBasedRequest
}

// RecognizeTextRequestFrom constructs a [RecognizeTextRequest] from an unsafe.Pointer.
//
// An image-analysis request that finds and recognizes text in an image.
func RecognizeTextRequestFrom(ptr unsafe.Pointer) RecognizeTextRequest {
	return RecognizeTextRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RecognizeTextRequestClass) Alloc() RecognizeTextRequest {
	rv := objc.Send[RecognizeTextRequest](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RecognizeTextRequestClass) New() RecognizeTextRequest {
	rv := objc.Send[RecognizeTextRequest](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizeTextRequest) Init() RecognizeTextRequest {
	rv := objc.Send[RecognizeTextRequest](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizeTextRequest) Autorelease() RecognizeTextRequest {
	rv := objc.Send[RecognizeTextRequest](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizeTextRequest creates a new RecognizeTextRequest instance.
func NewRecognizeTextRequest() RecognizeTextRequest {
	return getRecognizeTextRequestClass().New()
}



// A Boolean value that indicates whether to attempt detecting the language to use the appropriate model for recognition and language correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/automaticallydetectslanguage
func (r_ RecognizeTextRequest) AutomaticallyDetectsLanguage() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("automaticallyDetectsLanguage"))
	return rv
}


// A Boolean value that indicates whether to attempt detecting the language to use the appropriate model for recognition and language correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/automaticallydetectslanguage
func (r_ RecognizeTextRequest) SetAutomaticallyDetectsLanguage(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAutomaticallyDetectsLanguage:"), value)
}


// An array of strings to supplement the recognized languages at the word-recognition stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/customwords
func (r_ RecognizeTextRequest) CustomWords() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("customWords"))
	return rv
}


// An array of strings to supplement the recognized languages at the word-recognition stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/customwords
func (r_ RecognizeTextRequest) SetCustomWords(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCustomWords:"), value)
}


// The minimum height, relative to the image height, of the text to recognize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/minimumtextheight
func (r_ RecognizeTextRequest) MinimumTextHeight() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("minimumTextHeight"))
	return rv
}


// The minimum height, relative to the image height, of the text to recognize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/minimumtextheight
func (r_ RecognizeTextRequest) SetMinimumTextHeight(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMinimumTextHeight:"), value)
}


// An array of languages to detect, in priority order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/recognitionlanguages
func (r_ RecognizeTextRequest) RecognitionLanguages() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("recognitionLanguages"))
	return rv
}


// An array of languages to detect, in priority order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/recognitionlanguages
func (r_ RecognizeTextRequest) SetRecognitionLanguages(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRecognitionLanguages:"), value)
}


// A value that determines whether the request prioritizes accuracy or speed in text recognition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/recognitionlevel
func (r_ RecognizeTextRequest) RecognitionLevel() RequestTextRecognitionLevel /* not a class type */ {
	rv := objc.Send[RequestTextRecognitionLevel](r_.ID, objc.Sel("recognitionLevel"))
	return rv
}


// A value that determines whether the request prioritizes accuracy or speed in text recognition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/recognitionlevel
func (r_ RecognizeTextRequest) SetRecognitionLevel(value RequestTextRecognitionLevel /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRecognitionLevel:"), value)
}


// The results of the text recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/results
func (r_ RecognizeTextRequest) Results() IVNRecognizedTextObservation {
	rv := objc.Send[RecognizedTextObservation](r_.ID, objc.Sel("results"))
	return rv
}


// The results of the text recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/results
func (r_ RecognizeTextRequest) SetResults(value IVNRecognizedTextObservation) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setResults:"), value)
}


// A Boolean value that indicates whether the request applies language correction during the recognition process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/useslanguagecorrection
func (r_ RecognizeTextRequest) UsesLanguageCorrection() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("usesLanguageCorrection"))
	return rv
}


// A Boolean value that indicates whether the request applies language correction during the recognition process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/useslanguagecorrection
func (r_ RecognizeTextRequest) SetUsesLanguageCorrection(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setUsesLanguageCorrection:"), value)
}


// A constant for specifying revision 1 of the text recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequestrevision1
func (r_ RecognizeTextRequest) VNRecognizeTextRequestRevision1() int {
	rv := objc.Send[int](r_.ID, objc.Sel("VNRecognizeTextRequestRevision1"))
	return rv
}


// A constant for specifying revision 2 of the text recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequestrevision2
func (r_ RecognizeTextRequest) VNRecognizeTextRequestRevision2() int {
	rv := objc.Send[int](r_.ID, objc.Sel("VNRecognizeTextRequestRevision2"))
	return rv
}


// A constant for specifying revision 3 of the text recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequestrevision3
func (r_ RecognizeTextRequest) VNRecognizeTextRequestRevision3() int {
	rv := objc.Send[int](r_.ID, objc.Sel("VNRecognizeTextRequestRevision3"))
	return rv
}



