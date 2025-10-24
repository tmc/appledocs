// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNRecognizeTextRequest */


/* debug [class_header]: Header for VNRecognizeTextRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RecognizeTextRequest */
// An interface definition for the [RecognizeTextRequest] class.
type IRecognizeTextRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for RecognizeTextRequest */
	// properties:
	AutomaticallyDetectsLanguage() bool
	SetAutomaticallyDetectsLanguage(value bool)
	CustomWords() []string
	SetCustomWords(value []string)
	MinimumTextHeight() float32
	SetMinimumTextHeight(value float32)
	RecognitionLanguages() []string
	SetRecognitionLanguages(value []string)
	RecognitionLevel() RequestTextRecognitionLevel
	SetRecognitionLevel(value RequestTextRecognitionLevel)
	Results() []RecognizedTextObservation
	UsesLanguageCorrection() bool
	SetUsesLanguageCorrection(value bool)
	VNRecognizeTextRequestRevision1() int
	VNRecognizeTextRequestRevision2() int
	VNRecognizeTextRequestRevision3() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RecognizeTextRequest */
	// methods:
	SupportedRecognitionLanguagesAndReturnError(error_ objectivec.IObject) []string
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RecognizeTextRequest */
// Alloc allocates a new instance without initialization.
func (rc _RecognizeTextRequestClass) Alloc() RecognizeTextRequest {
	rv := objc.Send[RecognizeTextRequest](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RecognizeTextRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RecognizeTextRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RecognizeTextRequest */

// Requests a list of languages that the specified revision recognizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/supportedRecognitionLanguages(for:revision:)
func (rc _RecognizeTextRequestClass) SupportedRecognitionLanguagesForTextRecognitionLevelRevisionError(recognitionLevel RequestTextRecognitionLevel, requestRevision uint, error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](objc.ID(rc.class), objc.Sel("supportedRecognitionLanguagesForTextRecognitionLevel:revision:error:"), recognitionLevel, requestRevision, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedRecognitionLanguagesForTextRecognitionLevelRevisionError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RecognizeTextRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RecognizeTextRequest */

// Returns the identifiers of the languages that the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/supportedRecognitionLanguages()
func (r_ RecognizeTextRequest) SupportedRecognitionLanguagesAndReturnError(error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("supportedRecognitionLanguagesAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: SupportedRecognitionLanguagesAndReturnError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RecognizeTextRequest */

// A Boolean value that indicates whether to attempt detecting the language to use the appropriate model for recognition and language correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/automaticallyDetectsLanguage
func (r_ RecognizeTextRequest) AutomaticallyDetectsLanguage() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("automaticallyDetectsLanguage"))
	return rv
}/* debug [instance_properties/getter]: automaticallyDetectsLanguage */


// A Boolean value that indicates whether to attempt detecting the language to use the appropriate model for recognition and language correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/automaticallyDetectsLanguage
func (r_ RecognizeTextRequest) SetAutomaticallyDetectsLanguage(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAutomaticallyDetectsLanguage:"), value)
}/* debug [instance_properties/setter]: automaticallyDetectsLanguage */


// An array of strings to supplement the recognized languages at the word-recognition stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/customWords
func (r_ RecognizeTextRequest) CustomWords() []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("customWords"))
	return rv
}/* debug [instance_properties/getter]: customWords */


// An array of strings to supplement the recognized languages at the word-recognition stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/customWords
func (r_ RecognizeTextRequest) SetCustomWords(value []string) {
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
}/* debug [instance_properties/setter]: customWords */


// The minimum height, relative to the image height, of the text to recognize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/minimumTextHeight
func (r_ RecognizeTextRequest) MinimumTextHeight() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("minimumTextHeight"))
	return rv
}/* debug [instance_properties/getter]: minimumTextHeight */


// The minimum height, relative to the image height, of the text to recognize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/minimumTextHeight
func (r_ RecognizeTextRequest) SetMinimumTextHeight(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMinimumTextHeight:"), value)
}/* debug [instance_properties/setter]: minimumTextHeight */


// An array of languages to detect, in priority order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/recognitionLanguages
func (r_ RecognizeTextRequest) RecognitionLanguages() []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("recognitionLanguages"))
	return rv
}/* debug [instance_properties/getter]: recognitionLanguages */


// An array of languages to detect, in priority order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/recognitionLanguages
func (r_ RecognizeTextRequest) SetRecognitionLanguages(value []string) {
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
}/* debug [instance_properties/setter]: recognitionLanguages */


// A value that determines whether the request prioritizes accuracy or speed in text recognition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/recognitionLevel
func (r_ RecognizeTextRequest) RecognitionLevel() RequestTextRecognitionLevel {
	rv := objc.Send[RequestTextRecognitionLevel](r_.ID, objc.Sel("recognitionLevel"))
	return rv
}/* debug [instance_properties/getter]: recognitionLevel */


// A value that determines whether the request prioritizes accuracy or speed in text recognition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/recognitionLevel
func (r_ RecognizeTextRequest) SetRecognitionLevel(value RequestTextRecognitionLevel) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRecognitionLevel:"), value)
}/* debug [instance_properties/setter]: recognitionLevel */


// The results of the text recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/results
func (r_ RecognizeTextRequest) Results() []RecognizedTextObservation {
	rv := objc.Send[[]RecognizedTextObservation](r_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A Boolean value that indicates whether the request applies language correction during the recognition process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/usesLanguageCorrection
func (r_ RecognizeTextRequest) UsesLanguageCorrection() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("usesLanguageCorrection"))
	return rv
}/* debug [instance_properties/getter]: usesLanguageCorrection */


// A Boolean value that indicates whether the request applies language correction during the recognition process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/usesLanguageCorrection
func (r_ RecognizeTextRequest) SetUsesLanguageCorrection(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setUsesLanguageCorrection:"), value)
}/* debug [instance_properties/setter]: usesLanguageCorrection */


// A constant for specifying revision 1 of the text recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequestrevision1
func (r_ RecognizeTextRequest) VNRecognizeTextRequestRevision1() int {
	rv := objc.Send[int](r_.ID, objc.Sel("VNRecognizeTextRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNRecognizeTextRequestRevision1 */


// A constant for specifying revision 2 of the text recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequestrevision2
func (r_ RecognizeTextRequest) VNRecognizeTextRequestRevision2() int {
	rv := objc.Send[int](r_.ID, objc.Sel("VNRecognizeTextRequestRevision2"))
	return rv
}/* debug [instance_properties/getter]: VNRecognizeTextRequestRevision2 */


// A constant for specifying revision 3 of the text recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizetextrequestrevision3
func (r_ RecognizeTextRequest) VNRecognizeTextRequestRevision3() int {
	rv := objc.Send[int](r_.ID, objc.Sel("VNRecognizeTextRequestRevision3"))
	return rv
}/* debug [instance_properties/getter]: VNRecognizeTextRequestRevision3 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNRecognizeTextRequest */



