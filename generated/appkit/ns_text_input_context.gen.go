// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)
	AllowedInputSourceLocales() objc.IObject /* cross-framework: NSString */
	SetAllowedInputSourceLocales(value objc.IObject /* cross-framework: NSString */)
	Client() TextInputClient /* not a class type */
	SetClient(value TextInputClient /* not a class type */)
	KeyboardInputSources() TextInputSourceIdentifier /* not a class type */
	SetKeyboardInputSources(value TextInputSourceIdentifier /* not a class type */)
	SelectedKeyboardInputSource() TextInputSourceIdentifier /* not a class type */
	SetSelectedKeyboardInputSource(value TextInputSourceIdentifier /* not a class type */)
	// methods:
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
func NewTextInputContextWithClient(client objc.IObject) TextInputContext {
	instance := getTextInputContextClass().Alloc()
	rv := objc.Send[TextInputContext](instance.ID, objc.Sel("initWithClient:"), client)
	rv.Autorelease()
	return rv
}



// A Boolean value that indicates whether the client handles
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/acceptsglyphinfo
func (t_ TextInputContext) AcceptsGlyphInfo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("acceptsGlyphInfo"))
	return rv
}


// A Boolean value that indicates whether the client handles
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/acceptsglyphinfo
func (t_ TextInputContext) SetAcceptsGlyphInfo(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAcceptsGlyphInfo:"), value)
}


// The set of keyboard input source locales allowed when this input context is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/allowedinputsourcelocales
func (t_ TextInputContext) AllowedInputSourceLocales() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("allowedInputSourceLocales"))
	return rv
}


// The set of keyboard input source locales allowed when this input context is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/allowedinputsourcelocales
func (t_ TextInputContext) SetAllowedInputSourceLocales(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedInputSourceLocales:"), value)
}


// The owner of this input context. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/client
func (t_ TextInputContext) Client() TextInputClient /* not a class type */ {
	rv := objc.Send[TextInputClient](t_.ID, objc.Sel("client"))
	return rv
}


// The owner of this input context. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/client
func (t_ TextInputContext) SetClient(value TextInputClient /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setClient:"), value)
}


// The array of keyboard text input source identifier strings available to the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/keyboardinputsources
func (t_ TextInputContext) KeyboardInputSources() TextInputSourceIdentifier /* not a class type */ {
	rv := objc.Send[TextInputSourceIdentifier](t_.ID, objc.Sel("keyboardInputSources"))
	return rv
}


// The array of keyboard text input source identifier strings available to the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/keyboardinputsources
func (t_ TextInputContext) SetKeyboardInputSources(value TextInputSourceIdentifier /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setKeyboardInputSources:"), value)
}


// The identifier string for the selected keyboard text input source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/selectedkeyboardinputsource
func (t_ TextInputContext) SelectedKeyboardInputSource() TextInputSourceIdentifier /* not a class type */ {
	rv := objc.Send[TextInputSourceIdentifier](t_.ID, objc.Sel("selectedKeyboardInputSource"))
	return rv
}


// The identifier string for the selected keyboard text input source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputcontext/selectedkeyboardinputsource
func (t_ TextInputContext) SetSelectedKeyboardInputSource(value TextInputSourceIdentifier /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedKeyboardInputSource:"), value)
}


