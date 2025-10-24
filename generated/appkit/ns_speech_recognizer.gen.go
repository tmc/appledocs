// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSSpeechRecognizer */


/* debug [class_header]: Header for NSSpeechRecognizer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SpeechRecognizer */
// An interface definition for the [SpeechRecognizer] class.
type ISpeechRecognizer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SpeechRecognizer */
	// properties:
	BlocksOtherRecognizers() bool
	SetBlocksOtherRecognizers(value bool)
	Commands() []string
	SetCommands(value []string)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DisplayedCommandsTitle() objc.IObject /* cross-framework: NSString */
	SetDisplayedCommandsTitle(value objc.IObject /* cross-framework: NSString */)
	ListensInForegroundOnly() bool
	SetListensInForegroundOnly(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SpeechRecognizer */
	// methods:
	StartListening()
	StopListening()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SpeechRecognizer */
// Alloc allocates a new instance without initialization.
func (sc _SpeechRecognizerClass) Alloc() SpeechRecognizer {
	rv := objc.Send[SpeechRecognizer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SpeechRecognizer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SpeechRecognizer */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SpeechRecognizer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SpeechRecognizer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SpeechRecognizer */

// Tells the speech recognition engine to begin listening for commands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/startListening()
func (s_ SpeechRecognizer) StartListening() {
	objc.Send[objc.ID](s_.ID, objc.Sel("startListening"))
}/* debug [instance_methods/method]: StartListening */


// Tells the speech recognition engine to suspend listening for commands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/stopListening()
func (s_ SpeechRecognizer) StopListening() {
	objc.Send[objc.ID](s_.ID, objc.Sel("stopListening"))
}/* debug [instance_methods/method]: StopListening */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SpeechRecognizer */

// A Boolean value that indicates whether the speech recognizer object should block all other recognizers (that is, other applications attempting to understand spoken commands) when listening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/blocksOtherRecognizers
func (s_ SpeechRecognizer) BlocksOtherRecognizers() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("blocksOtherRecognizers"))
	return rv
}/* debug [instance_properties/getter]: blocksOtherRecognizers */


// A Boolean value that indicates whether the speech recognizer object should block all other recognizers (that is, other applications attempting to understand spoken commands) when listening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/blocksOtherRecognizers
func (s_ SpeechRecognizer) SetBlocksOtherRecognizers(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBlocksOtherRecognizers:"), value)
}/* debug [instance_properties/setter]: blocksOtherRecognizers */


// An array of strings defining the commands for which the speech recognizer object should listen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/commands
func (s_ SpeechRecognizer) Commands() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("commands"))
	return rv
}/* debug [instance_properties/getter]: commands */


// An array of strings defining the commands for which the speech recognizer object should listen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/commands
func (s_ SpeechRecognizer) SetCommands(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setCommands:"), nsArray)
}/* debug [instance_properties/setter]: commands */


// The delegate for the speech recognizer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/delegate
func (s_ SpeechRecognizer) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the speech recognizer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/delegate
func (s_ SpeechRecognizer) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The title of the commands section in the Speech Commands window or if there is no title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/displayedCommandsTitle
func (s_ SpeechRecognizer) DisplayedCommandsTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("displayedCommandsTitle"))
	return rv
}/* debug [instance_properties/getter]: displayedCommandsTitle */


// The title of the commands section in the Speech Commands window or if there is no title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/displayedCommandsTitle
func (s_ SpeechRecognizer) SetDisplayedCommandsTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDisplayedCommandsTitle:"), value)
}/* debug [instance_properties/setter]: displayedCommandsTitle */


// A Boolean value that indicates whether the speech recognizer object should only enable its commands when its application is the frontmost one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/listensInForegroundOnly
func (s_ SpeechRecognizer) ListensInForegroundOnly() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("listensInForegroundOnly"))
	return rv
}/* debug [instance_properties/getter]: listensInForegroundOnly */


// A Boolean value that indicates whether the speech recognizer object should only enable its commands when its application is the frontmost one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/listensInForegroundOnly
func (s_ SpeechRecognizer) SetListensInForegroundOnly(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setListensInForegroundOnly:"), value)
}/* debug [instance_properties/setter]: listensInForegroundOnly */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSpeechRecognizer */


