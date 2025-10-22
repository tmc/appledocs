// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	SupportedRecognitionLanguagesAndReturnError(error_ unsafe.Pointer) []string
	AutomaticallyDetectsLanguage() bool
	SetAutomaticallyDetectsLanguage(value bool)
	CustomWords() []string
	SetCustomWords(value []string)
	RecognitionLanguages() []string
	SetRecognitionLanguages(value []string)
	RecognitionLevel() RequestTextRecognitionLevel
	SetRecognitionLevel(value RequestTextRecognitionLevel)
	Results() []RecognizedTextObservation
	UsesLanguageCorrection() bool
	SetUsesLanguageCorrection(value bool)
	MinimumTextHeight() float32
	SetMinimumTextHeight(value float32)
	VNRecognizeTextRequestRevision1() int
	VNRecognizeTextRequestRevision2() int
	VNRecognizeTextRequestRevision3() int
}

// An image-analysis request that finds and recognizes text in an image.
//
// By default, a text recognition request first locates all possible glyphs or characters in the input image, and then analyzes each string. To specify or limit the languages to find in the request, set the property to an array that contains the names of the languages of text you want to recognize. Vision returns the result of this request in a object.
//
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


// Requests a list of languages that the specified revision recognizes.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/supportedRecognitionLanguages(for:revision:)
func (rc _RecognizeTextRequestClass) SupportedRecognitionLanguagesForTextRecognitionLevelRevisionError(recognitionLevel RequestTextRecognitionLevel, requestRevision uint, error_ unsafe.Pointer) []string {
	rv := objc.Send[[]string](objc.ID(rc.class), objc.Sel("supportedRecognitionLanguagesForTextRecognitionLevel:revision:error:"), recognitionLevel, requestRevision, error_)
	return rv
}

// Returns the identifiers of the languages that the request supports.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/supportedRecognitionLanguages()
func (r_ RecognizeTextRequest) SupportedRecognitionLanguagesAndReturnError(error_ unsafe.Pointer) []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("supportedRecognitionLanguagesAndReturnError:"), error_)
	return rv
}

// A Boolean value that indicates whether to attempt detecting the language to use the appropriate model for recognition and language correction.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/automaticallyDetectsLanguage
func (r_ RecognizeTextRequest) AutomaticallyDetectsLanguage() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("automaticallyDetectsLanguage"))
	return rv
}


// SetAutomaticallyDetectsLanguage sets the value of the automaticallyDetectsLanguage property.
// A Boolean value that indicates whether to attempt detecting the language to use the appropriate model for recognition and language correction.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/automaticallyDetectsLanguage
func (r_ RecognizeTextRequest) SetAutomaticallyDetectsLanguage(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAutomaticallyDetectsLanguage:"), value)
}

// An array of strings to supplement the recognized languages at the word-recognition stage.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/customWords
func (r_ RecognizeTextRequest) CustomWords() []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("customWords"))
	return rv
}


// SetCustomWords sets the value of the customWords property.
// An array of strings to supplement the recognized languages at the word-recognition stage.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/customWords
func (r_ RecognizeTextRequest) SetCustomWords(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](r_.ID, objc.Sel("setCustomWords:"), nsArray)
}

// An array of languages to detect, in priority order.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/recognitionLanguages
func (r_ RecognizeTextRequest) RecognitionLanguages() []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("recognitionLanguages"))
	return rv
}


// SetRecognitionLanguages sets the value of the recognitionLanguages property.
// An array of languages to detect, in priority order.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/recognitionLanguages
func (r_ RecognizeTextRequest) SetRecognitionLanguages(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](r_.ID, objc.Sel("setRecognitionLanguages:"), nsArray)
}

// A value that determines whether the request prioritizes accuracy or speed in text recognition.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/recognitionLevel
func (r_ RecognizeTextRequest) RecognitionLevel() RequestTextRecognitionLevel {
	rv := objc.Send[RequestTextRecognitionLevel](r_.ID, objc.Sel("recognitionLevel"))
	return rv
}


// SetRecognitionLevel sets the value of the recognitionLevel property.
// A value that determines whether the request prioritizes accuracy or speed in text recognition.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/recognitionLevel
func (r_ RecognizeTextRequest) SetRecognitionLevel(value RequestTextRecognitionLevel) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRecognitionLevel:"), value)
}

// The results of the text recognition request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/results
func (r_ RecognizeTextRequest) Results() []RecognizedTextObservation {
	rv := objc.Send[[]RecognizedTextObservation](r_.ID, objc.Sel("results"))
	return rv
}

// A Boolean value that indicates whether the request applies language correction during the recognition process.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/usesLanguageCorrection
func (r_ RecognizeTextRequest) UsesLanguageCorrection() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("usesLanguageCorrection"))
	return rv
}


// SetUsesLanguageCorrection sets the value of the usesLanguageCorrection property.
// A Boolean value that indicates whether the request applies language correction during the recognition process.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/usesLanguageCorrection
func (r_ RecognizeTextRequest) SetUsesLanguageCorrection(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setUsesLanguageCorrection:"), value)
}

// The minimum height, relative to the image height, of the text to recognize.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/minimumtextheight
func (r_ RecognizeTextRequest) MinimumTextHeight() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("minimumTextHeight"))
	return rv
}


// SetMinimumTextHeight sets the value of the minimumTextHeight property.
// The minimum height, relative to the image height, of the text to recognize.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequest/minimumtextheight
func (r_ RecognizeTextRequest) SetMinimumTextHeight(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMinimumTextHeight:"), value)
}

// A constant for specifying revision 1 of the text recognition request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequestrevision1
func (r_ RecognizeTextRequest) VNRecognizeTextRequestRevision1() int {
	rv := objc.Send[int](r_.ID, objc.Sel("VNRecognizeTextRequestRevision1"))
	return rv
}

// A constant for specifying revision 2 of the text recognition request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequestrevision2
func (r_ RecognizeTextRequest) VNRecognizeTextRequestRevision2() int {
	rv := objc.Send[int](r_.ID, objc.Sel("VNRecognizeTextRequestRevision2"))
	return rv
}

// A constant for specifying revision 3 of the text recognition request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequestrevision3
func (r_ RecognizeTextRequest) VNRecognizeTextRequestRevision3() int {
	rv := objc.Send[int](r_.ID, objc.Sel("VNRecognizeTextRequestRevision3"))
	return rv
}



