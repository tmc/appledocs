// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [SpeechSynthesisMarker] class.
type ISpeechSynthesisMarker interface {
	objectivec.IObject
	

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


	

	// methods:


}





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






// Creates a bookmark marker with a name and offset into the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(bookmarkName:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithBookmarkNameAtByteSampleOffset(mark objc.IObject /* cross-framework: NSString */, byteSampleOffset int) SpeechSynthesisMarker {
	instance := getSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[SpeechSynthesisMarker](instance.ID, objc.Sel("initWithBookmarkName:atByteSampleOffset:"), mark, byteSampleOffset)
	rv.Autorelease()
	return rv
}


// Creates a marker with a type and location of the request’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(markerType:forTextRange:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithMarkerTypeForTextRangeAtByteSampleOffset(type_ SpeechSynthesisMarkerMark, range_ corefoundation.Range, byteSampleOffset uint) SpeechSynthesisMarker {
	instance := getSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[SpeechSynthesisMarker](instance.ID, objc.Sel("initWithMarkerType:forTextRange:atByteSampleOffset:"), type_, range_, byteSampleOffset)
	rv.Autorelease()
	return rv
}


// Creates a paragraph marker with a range of the paragraph and offset into the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(paragraphRange:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithParagraphRangeAtByteSampleOffset(range_ corefoundation.Range, byteSampleOffset int) SpeechSynthesisMarker {
	instance := getSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[SpeechSynthesisMarker](instance.ID, objc.Sel("initWithParagraphRange:atByteSampleOffset:"), range_, byteSampleOffset)
	rv.Autorelease()
	return rv
}


// Creates a phoneme marker with a range of the phoneme and offset into the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(phonemeString:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithPhonemeStringAtByteSampleOffset(phoneme objc.IObject /* cross-framework: NSString */, byteSampleOffset int) SpeechSynthesisMarker {
	instance := getSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[SpeechSynthesisMarker](instance.ID, objc.Sel("initWithPhonemeString:atByteSampleOffset:"), phoneme, byteSampleOffset)
	rv.Autorelease()
	return rv
}


// Creates a sentence marker with a range of the sentence and offset into the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(sentenceRange:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithSentenceRangeAtByteSampleOffset(range_ corefoundation.Range, byteSampleOffset int) SpeechSynthesisMarker {
	instance := getSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[SpeechSynthesisMarker](instance.ID, objc.Sel("initWithSentenceRange:atByteSampleOffset:"), range_, byteSampleOffset)
	rv.Autorelease()
	return rv
}


// Creates a word marker with a range of the word and offset into the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(wordRange:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithWordRangeAtByteSampleOffset(range_ corefoundation.Range, byteSampleOffset int) SpeechSynthesisMarker {
	instance := getSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[SpeechSynthesisMarker](instance.ID, objc.Sel("initWithWordRange:atByteSampleOffset:"), range_, byteSampleOffset)
	rv.Autorelease()
	return rv
}






















// A string that represents the name of a bookmark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/bookmarkName
func (s_ SpeechSynthesisMarker) BookmarkName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("bookmarkName"))
	return rv
}


// A string that represents the name of a bookmark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/bookmarkName
func (s_ SpeechSynthesisMarker) SetBookmarkName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBookmarkName:"), value)
}


// The byte offset into the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/byteSampleOffset
func (s_ SpeechSynthesisMarker) ByteSampleOffset() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("byteSampleOffset"))
	return rv
}


// The byte offset into the audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/byteSampleOffset
func (s_ SpeechSynthesisMarker) SetByteSampleOffset(value uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setByteSampleOffset:"), value)
}


// The type that describes the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/mark-swift.property
func (s_ SpeechSynthesisMarker) Mark() SpeechSynthesisMarkerMark {
	rv := objc.Send[SpeechSynthesisMarkerMark](s_.ID, objc.Sel("mark"))
	return rv
}


// The type that describes the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/mark-swift.property
func (s_ SpeechSynthesisMarker) SetMark(value SpeechSynthesisMarkerMark) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMark:"), value)
}


// A string that represents a distinct sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/phoneme
func (s_ SpeechSynthesisMarker) Phoneme() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("phoneme"))
	return rv
}


// A string that represents a distinct sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/phoneme
func (s_ SpeechSynthesisMarker) SetPhoneme(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPhoneme:"), value)
}


// The location and length of the request’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/textRange
func (s_ SpeechSynthesisMarker) TextRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](s_.ID, objc.Sel("textRange"))
	return rv
}


// The location and length of the request’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/textRange
func (s_ SpeechSynthesisMarker) SetTextRange(value corefoundation.Range) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTextRange:"), value)
}


// A block that subclasses use to send marker information to the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovideraudiounit/speechsynthesisoutputmetadatablock
func (s_ SpeechSynthesisMarker) SpeechSynthesisOutputMetadataBlock() SpeechSynthesisProviderOutputBlock /* not a class type */ {
	rv := objc.Send[SpeechSynthesisProviderOutputBlock](s_.ID, objc.Sel("speechSynthesisOutputMetadataBlock"))
	return rv
}


// A block that subclasses use to send marker information to the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovideraudiounit/speechsynthesisoutputmetadatablock
func (s_ SpeechSynthesisMarker) SetSpeechSynthesisOutputMetadataBlock(value SpeechSynthesisProviderOutputBlock /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpeechSynthesisOutputMetadataBlock:"), value)
}







