// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextInputContext] class.
var (
	TextInputContextClass     _TextInputContextClass
	TextInputContextClassOnce sync.Once
)

func getTextInputContextClass() _TextInputContextClass {
	TextInputContextClassOnce.Do(func() {
		TextInputContextClass = _TextInputContextClass{objc.GetClass("NSTextInputContext")}
	})
	return TextInputContextClass
}

type _TextInputContextClass struct {
	class objc.Class
}

// An interface definition for the [TextInputContext] class.
type ITextInputContext interface {
	objectivec.IObject
	// properties:
	AllowedInputSourceLocales() []string /* primitive/slice/pointer. */
	SetAllowedInputSourceLocales(value []string /* primitive/slice/pointer. */)
	Client() objc.ID
	KeyboardInputSources() []string /* primitive/slice/pointer. */
	SelectedKeyboardInputSource() TextInputSourceIdentifier /* not a class type */
	SetSelectedKeyboardInputSource(value TextInputSourceIdentifier /* not a class type */)
	AcceptsGlyphInfo() bool /* primitive/slice/pointer. */
	SetAcceptsGlyphInfo(value bool /* primitive/slice/pointer. */)
	// methods:
	Deactivate()
	DiscardMarkedText()
	HandleEvent(event IEvent) bool /* primitive/slice/pointer. */
	InvalidateCharacterCoordinates()
	TextInputClientDidScroll()
	TextInputClientDidUpdateSelection()
	TextInputClientWillStartScrollingOrZooming()
}

// An object that represents the Cocoa text input system.
//
// The text input system communicates primarily with the client of the activated input context via the protocol.


// An object that represents the Cocoa text input system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext
type TextInputContext struct {
	objectivec.Object
}

// TextInputContextFrom constructs a [TextInputContext] from an unsafe.Pointer.
//
// An object that represents the Cocoa text input system.
func TextInputContextFrom(ptr unsafe.Pointer) TextInputContext {
	return TextInputContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextInputContextClass) Alloc() TextInputContext {
	rv := objc.Send[TextInputContext](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextInputContextClass) New() TextInputContext {
	rv := objc.Send[TextInputContext](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextInputContext) Init() TextInputContext {
	rv := objc.Send[TextInputContext](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextInputContext) Autorelease() TextInputContext {
	rv := objc.Send[TextInputContext](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextInputContext creates a new TextInputContext instance.
func NewTextInputContext() TextInputContext {
	return getTextInputContextClass().New()
}



// The designated initializer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/init(client:)
func NewTextInputContextWithClient(client objectivec.IObject) TextInputContext {
	instance := getTextInputContextClass().Alloc()
	rv := objc.Send[TextInputContext](instance.ID, objc.Sel("initWithClient:"), client)
	rv.Autorelease()
	return rv
}



// Returns the display name for the given text input source identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/localizedName(forInputSource:)
func (tc _TextInputContextClass) LocalizedNameForInputSource(inputSourceIdentifier TextInputSourceIdentifier /* not a class type */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](objc.ID(tc.class), objc.Sel("localizedNameForInputSource:"), inputSourceIdentifier)
	return rv
}


// Returns the current, activated, text input context object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/current
func (tc _TextInputContextClass) CurrentInputContext() TextInputContext {
	rv := objc.Send[TextInputContext](objc.ID(tc.class), objc.Sel("currentInputContext"))
	return rv
}

// Deactivates the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/deactivate()
func (t_ TextInputContext) Deactivate() {
	objc.Send[objc.ID](t_.ID, objc.Sel("deactivate"))
}


// Tells the Cocoa text input system to discard the current conversion session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/discardMarkedText()
func (t_ TextInputContext) DiscardMarkedText() {
	objc.Send[objc.ID](t_.ID, objc.Sel("discardMarkedText"))
}


// Tells the Cocoa text input system to handle mouse or key events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/handleEvent(_:)
func (t_ TextInputContext) HandleEvent(event IEvent) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("handleEvent:"), event)
	return rv
}


// Notifies the Cocoa text input system that the position information previously queried via methods like needs to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/invalidateCharacterCoordinates()
func (t_ TextInputContext) InvalidateCharacterCoordinates() {
	objc.Send[objc.ID](t_.ID, objc.Sel("invalidateCharacterCoordinates"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/textInputClientDidScroll()
func (t_ TextInputContext) TextInputClientDidScroll() {
	objc.Send[objc.ID](t_.ID, objc.Sel("textInputClientDidScroll"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/textInputClientDidUpdateSelection()
func (t_ TextInputContext) TextInputClientDidUpdateSelection() {
	objc.Send[objc.ID](t_.ID, objc.Sel("textInputClientDidUpdateSelection"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/textInputClientWillStartScrollingOrZooming()
func (t_ TextInputContext) TextInputClientWillStartScrollingOrZooming() {
	objc.Send[objc.ID](t_.ID, objc.Sel("textInputClientWillStartScrollingOrZooming"))
}


// The set of keyboard input source locales allowed when this input context is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/allowedInputSourceLocales
func (t_ TextInputContext) AllowedInputSourceLocales() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](t_.ID, objc.Sel("allowedInputSourceLocales"))
	return rv
}


// The set of keyboard input source locales allowed when this input context is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/allowedInputSourceLocales
func (t_ TextInputContext) SetAllowedInputSourceLocales(value []string /* primitive/slice/pointer. */) {
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
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedInputSourceLocales:"), nsArray)
}


// The owner of this input context. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/client
func (t_ TextInputContext) Client() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("client"))
	return rv
}


// Returns the current, activated, text input context object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/current
func (t_ TextInputContext) CurrentInputContext() ITextInputContext {
	rv := objc.Send[TextInputContext](t_.ID, objc.Sel("currentInputContext"))
	return rv
}


// The array of keyboard text input source identifier strings available to the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/keyboardInputSources
func (t_ TextInputContext) KeyboardInputSources() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](t_.ID, objc.Sel("keyboardInputSources"))
	return rv
}


// The identifier string for the selected keyboard text input source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/selectedKeyboardInputSource
func (t_ TextInputContext) SelectedKeyboardInputSource() TextInputSourceIdentifier /* not a class type */ {
	rv := objc.Send[TextInputSourceIdentifier](t_.ID, objc.Sel("selectedKeyboardInputSource"))
	return rv
}


// The identifier string for the selected keyboard text input source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext/selectedKeyboardInputSource
func (t_ TextInputContext) SetSelectedKeyboardInputSource(value TextInputSourceIdentifier /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedKeyboardInputSource:"), value)
}


// A Boolean value that indicates whether the client handles
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/acceptsglyphinfo
func (t_ TextInputContext) AcceptsGlyphInfo() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("acceptsGlyphInfo"))
	return rv
}


// A Boolean value that indicates whether the client handles
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/acceptsglyphinfo
func (t_ TextInputContext) SetAcceptsGlyphInfo(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAcceptsGlyphInfo:"), value)
}


