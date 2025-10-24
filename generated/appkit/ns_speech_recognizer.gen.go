// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SpeechRecognizer] class.
var (
	SpeechRecognizerClass     _SpeechRecognizerClass
	SpeechRecognizerClassOnce sync.Once
)

func getSpeechRecognizerClass() _SpeechRecognizerClass {
	SpeechRecognizerClassOnce.Do(func() {
		SpeechRecognizerClass = _SpeechRecognizerClass{objc.GetClass("NSSpeechRecognizer")}
	})
	return SpeechRecognizerClass
}

type _SpeechRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [SpeechRecognizer] class.
type ISpeechRecognizer interface {
	objectivec.IObject
	// properties:
	BlocksOtherRecognizers() bool
	SetBlocksOtherRecognizers(value bool)
	Commands() objc.IObject /* cross-framework: NSString */
	SetCommands(value objc.IObject /* cross-framework: NSString */)
	Delegate() SpeechRecognizerDelegate /* not a class type */
	SetDelegate(value SpeechRecognizerDelegate /* not a class type */)
	DisplayedCommandsTitle() objc.IObject /* cross-framework: NSString */
	SetDisplayedCommandsTitle(value objc.IObject /* cross-framework: NSString */)
	ListensInForegroundOnly() bool
	SetListensInForegroundOnly(value bool)
	// methods:
}

// The Cocoa interface to speech recognition in macOS.
//
// provides a “command and control” style of voice recognition system, where the command phrases must be defined prior to listening, in contrast to a dictation system where the recognized text is unconstrained. Through an instance, Cocoa apps can use the speech recognition engine built into macOS to recognize spoken commands. With speech recognition, users can accomplish complex tasks with spoken commands—for example, “Move pawn B2 to B4” and “Take back move.” The class has a property that lets you specify which spoken words should be recognized as commands ( ) and methods that let you start and stop listening ( and ). When the speech recognition facility recognizes one of the designated commands, invokes the delegation method , allowing the delegate to perform the command. Speech recognition is just one of the macOS speech technologies. The speech synthesis technology allows applications to “pronounce” written text in U.S. English and over 25 other languages, with a number of different voices and dialects for each language ( is the Cocoa interface to this technology). Both speech technologies provide benefits for all users, and are particularly useful to those users who have difficulties seeing the screen or using the mouse and keyboard. By incorporating speech into your application, you can provide a concurrent mode of interaction for your users: In macOS, your software can accept input and provide output without requiring users to change their working context.


// The Cocoa interface to speech recognition in macOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer
type SpeechRecognizer struct {
	objectivec.Object
}

// SpeechRecognizerFrom constructs a [SpeechRecognizer] from an unsafe.Pointer.
//
// The Cocoa interface to speech recognition in macOS.
func SpeechRecognizerFrom(ptr unsafe.Pointer) SpeechRecognizer {
	return SpeechRecognizer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SpeechRecognizerClass) Alloc() SpeechRecognizer {
	rv := objc.Send[SpeechRecognizer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SpeechRecognizerClass) New() SpeechRecognizer {
	rv := objc.Send[SpeechRecognizer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpeechRecognizer) Init() SpeechRecognizer {
	rv := objc.Send[SpeechRecognizer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpeechRecognizer) Autorelease() SpeechRecognizer {
	rv := objc.Send[SpeechRecognizer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpeechRecognizer creates a new SpeechRecognizer instance.
func NewSpeechRecognizer() SpeechRecognizer {
	return getSpeechRecognizerClass().New()
}



// A Boolean value that indicates whether the speech recognizer object should block all other recognizers (that is, other applications attempting to understand spoken commands) when listening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/blocksotherrecognizers
func (s_ SpeechRecognizer) BlocksOtherRecognizers() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("blocksOtherRecognizers"))
	return rv
}


// A Boolean value that indicates whether the speech recognizer object should block all other recognizers (that is, other applications attempting to understand spoken commands) when listening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/blocksotherrecognizers
func (s_ SpeechRecognizer) SetBlocksOtherRecognizers(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBlocksOtherRecognizers:"), value)
}


// An array of strings defining the commands for which the speech recognizer object should listen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/commands
func (s_ SpeechRecognizer) Commands() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("commands"))
	return rv
}


// An array of strings defining the commands for which the speech recognizer object should listen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/commands
func (s_ SpeechRecognizer) SetCommands(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCommands:"), value)
}


// The delegate for the speech recognizer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/delegate
func (s_ SpeechRecognizer) Delegate() SpeechRecognizerDelegate /* not a class type */ {
	rv := objc.Send[SpeechRecognizerDelegate](s_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for the speech recognizer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/delegate
func (s_ SpeechRecognizer) SetDelegate(value SpeechRecognizerDelegate /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}


// The title of the commands section in the Speech Commands window or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/displayedcommandstitle
func (s_ SpeechRecognizer) DisplayedCommandsTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("displayedCommandsTitle"))
	return rv
}


// The title of the commands section in the Speech Commands window or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/displayedcommandstitle
func (s_ SpeechRecognizer) SetDisplayedCommandsTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDisplayedCommandsTitle:"), value)
}


// A Boolean value that indicates whether the speech recognizer object should only enable its commands when its application is the frontmost one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/listensinforegroundonly
func (s_ SpeechRecognizer) ListensInForegroundOnly() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("listensInForegroundOnly"))
	return rv
}


// A Boolean value that indicates whether the speech recognizer object should only enable its commands when its application is the frontmost one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsspeechrecognizer/listensinforegroundonly
func (s_ SpeechRecognizer) SetListensInForegroundOnly(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setListensInForegroundOnly:"), value)
}



