// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextTab] class.
var (
	TextTabClass     _TextTabClass
	TextTabClassOnce sync.Once
)

func getTextTabClass() _TextTabClass {
	TextTabClassOnce.Do(func() {
		TextTabClass = _TextTabClass{objc.GetClass("NSTextTab")}
	})
	return TextTabClass
}

type _TextTabClass struct {
	class objc.Class
}

// An interface definition for the [TextTab] class.
type ITextTab interface {
	objectivec.IObject
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	Location() float64
	SetLocation(value float64)
	Options() unsafe.Pointer
	SetOptions(value unsafe.Pointer)
	TabStopType() unsafe.Pointer
	SetTabStopType(value unsafe.Pointer)
}

// A tab in a paragraph.
//
// A text tab represents a tab in an object, storing an alignment type and location. objects are most frequently used with the TextKit system and with and objects. The text system supports four alignment types: left, center, right, and decimal (based on the decimal separator character of the locale in effect). These alignment types are absolute, not based on the line sweep direction of text. For example, tabbed text is always positioned to the left of a right-aligned tab, whether the line sweep direction is left to right or right to left. A tab’s location, on the other hand, is relative to the back margin. A tab set at 1.5”, for example, is at 1.5” from the right in right to left text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTab
type TextTab struct {
	objectivec.Object
}

// TextTabFrom constructs a [TextTab] from an unsafe.Pointer.
//
// A tab in a paragraph.
func TextTabFrom(ptr unsafe.Pointer) TextTab {
	return TextTab{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextTabClass) Alloc() TextTab {
	rv := objc.Send[TextTab](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextTabClass) New() TextTab {
	rv := objc.Send[TextTab](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextTab) Init() TextTab {
	rv := objc.Send[TextTab](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextTab) Autorelease() TextTab {
	rv := objc.Send[TextTab](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextTab creates a new TextTab instance.
func NewTextTab() TextTab {
	return getTextTabClass().New()
}


// The text alignment of the text tab.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttab/alignment
func (t_ TextTab) Alignment() TextAlignment {
	rv := objc.Send[TextAlignment](t_.ID, objc.Sel("alignment"))
	return rv
}


// SetAlignment sets the value of the alignment property.
// The text alignment of the text tab.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttab/alignment
func (t_ TextTab) SetAlignment(value TextAlignment) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlignment:"), value)
}

// The text tab’s ruler location relative to the back margin.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttab/location
func (t_ TextTab) Location() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("location"))
	return rv
}


// SetLocation sets the value of the location property.
// The text tab’s ruler location relative to the back margin.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttab/location
func (t_ TextTab) SetLocation(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLocation:"), value)
}

// The dictionary of attributes for the text tab.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttab/options
func (t_ TextTab) Options() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("options"))
	return rv
}


// SetOptions sets the value of the options property.
// The dictionary of attributes for the text tab.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttab/options
func (t_ TextTab) SetOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setOptions:"), value)
}

// The text tab’s type of tab stop.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttab/tabstoptype
func (t_ TextTab) TabStopType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tabStopType"))
	return rv
}


// SetTabStopType sets the value of the tabStopType property.
// The text tab’s type of tab stop.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttab/tabstoptype
func (t_ TextTab) SetTabStopType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTabStopType:"), value)
}



