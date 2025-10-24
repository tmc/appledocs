// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ScrubberTextItemView] class.
var (
	ScrubberTextItemViewClass     _ScrubberTextItemViewClass
	ScrubberTextItemViewClassOnce sync.Once
)

func getScrubberTextItemViewClass() _ScrubberTextItemViewClass {
	ScrubberTextItemViewClassOnce.Do(func() {
		ScrubberTextItemViewClass = _ScrubberTextItemViewClass{objc.GetClass("NSScrubberTextItemView")}
	})
	return ScrubberTextItemViewClass
}

type _ScrubberTextItemViewClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberTextItemView] class.
type IScrubberTextItemView interface {
	IScrubberItemView
	// properties:
	TextField() ITextField
	SetTextField(value ITextField)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A concrete view subclass for displaying text for an item in a scrubber.
//
// Provide the text you want to display in the scrubber item to the property. If you want finer control over the appearance of the text, you can access the underlying text field using the property.


// A concrete view subclass for displaying text for an item in a scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberTextItemView
type ScrubberTextItemView struct {
	ScrubberItemView
}

// ScrubberTextItemViewFrom constructs a [ScrubberTextItemView] from an unsafe.Pointer.
//
// A concrete view subclass for displaying text for an item in a scrubber.
func ScrubberTextItemViewFrom(ptr unsafe.Pointer) ScrubberTextItemView {
	return ScrubberTextItemView{
		ScrubberItemView: ScrubberItemViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberTextItemViewClass) Alloc() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScrubberTextItemViewClass) New() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberTextItemView) Init() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberTextItemView) Autorelease() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberTextItemView creates a new ScrubberTextItemView instance.
func NewScrubberTextItemView() ScrubberTextItemView {
	return getScrubberTextItemViewClass().New()
}



// The text field that the scrubber item uses to display its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubbertextitemview/textfield
func (s_ ScrubberTextItemView) TextField() ITextField {
	rv := objc.Send[TextField](s_.ID, objc.Sel("textField"))
	return rv
}


// The text field that the scrubber item uses to display its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubbertextitemview/textfield
func (s_ ScrubberTextItemView) SetTextField(value ITextField) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTextField:"), value)
}


// The text displayed for the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubbertextitemview/title
func (s_ ScrubberTextItemView) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("title"))
	return rv
}


// The text displayed for the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubbertextitemview/title
func (s_ ScrubberTextItemView) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitle:"), value)
}



