// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFTranscription */


/* debug [class_header]: Header for SFTranscription */
// The class instance for the [SFTranscription] class.
var (
	SFTranscriptionClass     _SFTranscriptionClass
	SFTranscriptionClassOnce sync.Once
)

func getSFTranscriptionClass() _SFTranscriptionClass {
	SFTranscriptionClassOnce.Do(func() {
		SFTranscriptionClass = _SFTranscriptionClass{objc.GetClass("SFTranscription")}
	})
	return SFTranscriptionClass
}

type _SFTranscriptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFTranscription */
// An interface definition for the [SFTranscription] class.
type ISFTranscription interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFTranscription */
	// properties:
	AveragePauseDuration() float64
	FormattedString() objc.IObject /* cross-framework: NSString */
	Segments() []SFTranscriptionSegment
	SpeakingRate() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFTranscription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFTranscription */
// Alloc allocates a new instance without initialization.
func (sc _SFTranscriptionClass) Alloc() SFTranscription {
	rv := objc.Send[SFTranscription](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFTranscriptionClass) New() SFTranscription {
	rv := objc.Send[SFTranscription](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFTranscription) Init() SFTranscription {
	rv := objc.Send[SFTranscription](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFTranscription) Autorelease() SFTranscription {
	rv := objc.Send[SFTranscription](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFTranscription creates a new SFTranscription instance.
func NewSFTranscription() SFTranscription {
	return getSFTranscriptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFTranscription */
// A textual representation of the specified speech in its entirety, as recognized by the speech recognizer.
//
// Use to obtain all the recognized utterances from your audio content. An is a vocalized word or group of words that represent a single meaning to the speech recognizer ( ). Use the property to retrieve the entire transcription of utterances, or use the property to retrieve an individual utterance ( ). You don’t create an directly. Instead, you retrieve it from an instance. The speech recognizer sends a speech recognition result to your app in one of two ways, depending on how your app started a speech recognition task. You can start a speech recognition task by using the speech recognizer’s method. When the task is complete, the speech recognizer sends an instance to your closure. Alternatively, you can use the speech recognizer’s method to start a speech recognition task. When the task is complete, the speech recognizer uses your to send an by using the delegate’s method. An represents only a potential version of the speech. It might not be an accurate representation of the utterances.


// A textual representation of the specified speech in its entirety, as recognized by the speech recognizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFTranscription
type SFTranscription struct {
	objectivec.Object
}

// SFTranscriptionFrom constructs a [SFTranscription] from an unsafe.Pointer.
//
// A textual representation of the specified speech in its entirety, as recognized by the speech recognizer.
func SFTranscriptionFrom(ptr unsafe.Pointer) SFTranscription {
	return SFTranscription{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFTranscription *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFTranscription */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFTranscription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFTranscription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFTranscription */

// The average pause duration between words, measured in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFTranscription/averagePauseDuration
func (s_ SFTranscription) AveragePauseDuration() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("averagePauseDuration"))
	return rv
}/* debug [instance_properties/getter]: averagePauseDuration */


// The entire transcription of utterances, formatted into a single, user-displayable string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFTranscription/formattedString
func (s_ SFTranscription) FormattedString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("formattedString"))
	return rv
}/* debug [instance_properties/getter]: formattedString */


// An array of transcription segments that represent the parts of the transcription, as identified by the speech recognizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFTranscription/segments
func (s_ SFTranscription) Segments() []SFTranscriptionSegment {
	rv := objc.Send[[]SFTranscriptionSegment](s_.ID, objc.Sel("segments"))
	return rv
}/* debug [instance_properties/getter]: segments */


// The number of words spoken per minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFTranscription/speakingRate
func (s_ SFTranscription) SpeakingRate() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("speakingRate"))
	return rv
}/* debug [instance_properties/getter]: speakingRate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFTranscription */



