// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextView] class.
var (
	TextViewClass     _TextViewClass
	TextViewClassOnce sync.Once
)

func getTextViewClass() _TextViewClass {
	TextViewClassOnce.Do(func() {
		TextViewClass = _TextViewClass{objc.GetClass("NSTextView")}
	})
	return TextViewClass
}

type _TextViewClass struct {
	class objc.Class
}

// An interface definition for the [TextView] class.
type ITextView interface {
	IText
}

// A view that draws text and handles user interactions with that text.
//
// The class is the front-end class to the AppKit text system. The class draws the text managed by the back-end components and handles user events to select and modify its text, in addition to supporting rich text, attachments, input management, and key binding, and marked text attributes. is the principal means to obtain a text object that caters to almost all needs for displaying and managing text at the user interface level. While is a subclass of the class — which declares the most general Cocoa interface to the text system — adds major features beyond the capabilities of . You can also do more powerful and more creative text manipulation (such as displaying text in a circle) using , , , and related classes. You’re more likely to use the class than . It’s also important to remember that conforms to a large number of protocols, the methods of which are available to instances of the class. communicates with its delegate through methods declared both by the and by its superclass’s protocol, . All delegation messages come from the first text view. In macOS 12 and later, if you explicitly call the property on a text view or text container, the framework reverts to a compatibility mode that uses . The text view also switches to this compatibility mode when it encounters text content that’s not yet supported, such as .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView
type TextView struct {
	Text
}

// TextViewFrom constructs a [TextView] from an unsafe.Pointer.
//
// A view that draws text and handles user interactions with that text.
func TextViewFrom(ptr unsafe.Pointer) TextView {
	return TextView{
		Text: TextFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextViewClass) Alloc() TextView {
	rv := objc.Send[TextView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextViewClass) New() TextView {
	rv := objc.Send[TextView](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextView) Init() TextView {
	rv := objc.Send[TextView](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextView) Autorelease() TextView {
	rv := objc.Send[TextView](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextView creates a new TextView instance.
func NewTextView() TextView {
	return getTextViewClass().New()
}


// The layout manager that lays out text for the receiver’s text container.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/layoutManager
func (t_ TextView) LayoutManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("layoutManager"))
	return rv
}

// The receiver’s text container.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainer
func (t_ TextView) TextContainer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textContainer"))
	return rv
}


// SetTextContainer sets the value of the textContainer property.
// The receiver’s text container.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainer
func (t_ TextView) SetTextContainer(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainer:"), value)
}
// The receiver’s text storage object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textStorage
func (t_ TextView) TextStorage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textStorage"))
	return rv
}



