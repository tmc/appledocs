// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVSpeechSynthesisMarker */


/* debug [class_header]: Header for AVSpeechSynthesisMarker */
// The class instance for the [SpeechSynthesisMarker] class.
var (
	SpeechSynthesisMarkerClass     _SpeechSynthesisMarkerClass
	SpeechSynthesisMarkerClassOnce sync.Once
)

func getSpeechSynthesisMarkerClass() _SpeechSynthesisMarkerClass {
	SpeechSynthesisMarkerClassOnce.Do(func() {
		SpeechSynthesisMarkerClass = _SpeechSynthesisMarkerClass{objc.GetClass("AVSpeechSynthesisMarker")}
	})
	return SpeechSynthesisMarkerClass
}

type _SpeechSynthesisMarkerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SpeechSynthesisMarker */
// An interface definition for the [SpeechSynthesisMarker] class.
type ISpeechSynthesisMarker interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SpeechSynthesisMarker */
	// properties:
	BookmarkName() objc.IObject /* cross-framework: NSString */
	SetBookmarkName(value objc.IObject /* cross-framework: NSString */)
	ByteSampleOffset() uint
	SetByteSampleOffset(value uint)
	Mark() SpeechSynthesisMarkerMark
	SetMark(value SpeechSynthesisMarkerMark)
	Phoneme() objc.IObject /* cross-framework: NSString */
	SetPhoneme(value objc.IObject /* cross-framework: NSString */)
	TextRange() corefoundation.Range
	SetTextRange(value corefoundation.Range)
	SpeechSynthesisOutputMetadataBlock() SpeechSynthesisProviderOutputBlock /* not a class type */
	SetSpeechSynthesisOutputMetadataBlock(value SpeechSynthesisProviderOutputBlock /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SpeechSynthesisMarker */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SpeechSynthesisMarker */
// Alloc allocates a new instance without initialization.
func (sc _SpeechSynthesisMarkerClass) Alloc() SpeechSynthesisMarker {
	rv := objc.Send[SpeechSynthesisMarker](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SpeechSynthesisMarkerClass) New() SpeechSynthesisMarker {
	rv := objc.Send[SpeechSynthesisMarker](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpeechSynthesisMarker) Init() SpeechSynthesisMarker {
	rv := objc.Send[SpeechSynthesisMarker](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpeechSynthesisMarker) Autorelease() SpeechSynthesisMarker {
	rv := objc.Send[SpeechSynthesisMarker](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpeechSynthesisMarker creates a new SpeechSynthesisMarker instance.
func NewSpeechSynthesisMarker() SpeechSynthesisMarker {
	return getSpeechSynthesisMarkerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SpeechSynthesisMarker */
// An object that contains information about the synthesized audio.


// An object that contains information about the synthesized audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker
type SpeechSynthesisMarker struct {
	objectivec.Object
}

// SpeechSynthesisMarkerFrom constructs a [SpeechSynthesisMarker] from an unsafe.Pointer.
//
// An object that contains information about the synthesized audio.
func SpeechSynthesisMarkerFrom(ptr unsafe.Pointer) SpeechSynthesisMarker {
	return SpeechSynthesisMarker{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SpeechSynthesisMarker */

// Creates a bookmark marker with a name and offset into the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(bookmarkName:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithBookmarkNameAtByteSampleOffset(mark objc.IObject /* cross-framework: NSString */, byteSampleOffset int) SpeechSynthesisMarker {
	instance := getSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[SpeechSynthesisMarker](instance.ID, objc.Sel("initWithBookmarkName:atByteSampleOffset:"), mark, byteSampleOffset)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSpeechSynthesisMarkerWithBookmarkNameAtByteSampleOffset */


// Creates a marker with a type and location of the request’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(markerType:forTextRange:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithMarkerTypeForTextRangeAtByteSampleOffset(type_ SpeechSynthesisMarkerMark, range_ corefoundation.Range, byteSampleOffset uint) SpeechSynthesisMarker {
	instance := getSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[SpeechSynthesisMarker](instance.ID, objc.Sel("initWithMarkerType:forTextRange:atByteSampleOffset:"), type_, range_, byteSampleOffset)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSpeechSynthesisMarkerWithMarkerTypeForTextRangeAtByteSampleOffset */


// Creates a paragraph marker with a range of the paragraph and offset into the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(paragraphRange:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithParagraphRangeAtByteSampleOffset(range_ corefoundation.Range, byteSampleOffset int) SpeechSynthesisMarker {
	instance := getSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[SpeechSynthesisMarker](instance.ID, objc.Sel("initWithParagraphRange:atByteSampleOffset:"), range_, byteSampleOffset)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSpeechSynthesisMarkerWithParagraphRangeAtByteSampleOffset */


// Creates a phoneme marker with a range of the phoneme and offset into the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(phonemeString:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithPhonemeStringAtByteSampleOffset(phoneme objc.IObject /* cross-framework: NSString */, byteSampleOffset int) SpeechSynthesisMarker {
	instance := getSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[SpeechSynthesisMarker](instance.ID, objc.Sel("initWithPhonemeString:atByteSampleOffset:"), phoneme, byteSampleOffset)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSpeechSynthesisMarkerWithPhonemeStringAtByteSampleOffset */


// Creates a sentence marker with a range of the sentence and offset into the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(sentenceRange:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithSentenceRangeAtByteSampleOffset(range_ corefoundation.Range, byteSampleOffset int) SpeechSynthesisMarker {
	instance := getSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[SpeechSynthesisMarker](instance.ID, objc.Sel("initWithSentenceRange:atByteSampleOffset:"), range_, byteSampleOffset)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSpeechSynthesisMarkerWithSentenceRangeAtByteSampleOffset */


// Creates a word marker with a range of the word and offset into the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(wordRange:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithWordRangeAtByteSampleOffset(range_ corefoundation.Range, byteSampleOffset int) SpeechSynthesisMarker {
	instance := getSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[SpeechSynthesisMarker](instance.ID, objc.Sel("initWithWordRange:atByteSampleOffset:"), range_, byteSampleOffset)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSpeechSynthesisMarkerWithWordRangeAtByteSampleOffset */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SpeechSynthesisMarker */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SpeechSynthesisMarker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SpeechSynthesisMarker */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SpeechSynthesisMarker */

// A string that represents the name of a bookmark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/bookmarkName
func (s_ SpeechSynthesisMarker) BookmarkName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("bookmarkName"))
	return rv
}/* debug [instance_properties/getter]: bookmarkName */


// A string that represents the name of a bookmark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/bookmarkName
func (s_ SpeechSynthesisMarker) SetBookmarkName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBookmarkName:"), value)
}/* debug [instance_properties/setter]: bookmarkName */


// The byte offset into the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/byteSampleOffset
func (s_ SpeechSynthesisMarker) ByteSampleOffset() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("byteSampleOffset"))
	return rv
}/* debug [instance_properties/getter]: byteSampleOffset */


// The byte offset into the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/byteSampleOffset
func (s_ SpeechSynthesisMarker) SetByteSampleOffset(value uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setByteSampleOffset:"), value)
}/* debug [instance_properties/setter]: byteSampleOffset */


// The type that describes the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/mark-swift.property
func (s_ SpeechSynthesisMarker) Mark() SpeechSynthesisMarkerMark {
	rv := objc.Send[SpeechSynthesisMarkerMark](s_.ID, objc.Sel("mark"))
	return rv
}/* debug [instance_properties/getter]: mark */


// The type that describes the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/mark-swift.property
func (s_ SpeechSynthesisMarker) SetMark(value SpeechSynthesisMarkerMark) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMark:"), value)
}/* debug [instance_properties/setter]: mark */


// A string that represents a distinct sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/phoneme
func (s_ SpeechSynthesisMarker) Phoneme() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("phoneme"))
	return rv
}/* debug [instance_properties/getter]: phoneme */


// A string that represents a distinct sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/phoneme
func (s_ SpeechSynthesisMarker) SetPhoneme(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPhoneme:"), value)
}/* debug [instance_properties/setter]: phoneme */


// The location and length of the request’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/textRange
func (s_ SpeechSynthesisMarker) TextRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](s_.ID, objc.Sel("textRange"))
	return rv
}/* debug [instance_properties/getter]: textRange */


// The location and length of the request’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/textRange
func (s_ SpeechSynthesisMarker) SetTextRange(value corefoundation.Range) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTextRange:"), value)
}/* debug [instance_properties/setter]: textRange */


// A block that subclasses use to send marker information to the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovideraudiounit/speechsynthesisoutputmetadatablock
func (s_ SpeechSynthesisMarker) SpeechSynthesisOutputMetadataBlock() SpeechSynthesisProviderOutputBlock /* not a class type */ {
	rv := objc.Send[SpeechSynthesisProviderOutputBlock](s_.ID, objc.Sel("speechSynthesisOutputMetadataBlock"))
	return rv
}/* debug [instance_properties/getter]: speechSynthesisOutputMetadataBlock */


// A block that subclasses use to send marker information to the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovideraudiounit/speechsynthesisoutputmetadatablock
func (s_ SpeechSynthesisMarker) SetSpeechSynthesisOutputMetadataBlock(value SpeechSynthesisProviderOutputBlock /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpeechSynthesisOutputMetadataBlock:"), value)
}/* debug [instance_properties/setter]: speechSynthesisOutputMetadataBlock */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVSpeechSynthesisMarker */


