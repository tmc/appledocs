// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextTab */


/* debug [class_header]: Header for NSTextTab */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextTab */
// An interface definition for the [TextTab] class.
type ITextTab interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextTab */
	// properties:
	Alignment() TextAlignment
	Location() float64
	Options() foundation.IDictionary
	TabStopType() TextTabType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextTab */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextTab */
// Alloc allocates a new instance without initialization.
func (tc _TextTabClass) Alloc() TextTab {
	rv := objc.Send[TextTab](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextTab */
// A tab in a paragraph.
//
// A text tab represents a tab in an object, storing an alignment type and location. objects are most frequently used with the TextKit system and with and objects. The text system supports four alignment types: left, center, right, and decimal (based on the decimal separator character of the locale in effect). These alignment types are absolute, not based on the line sweep direction of text. For example, tabbed text is always positioned to the left of a right-aligned tab, whether the line sweep direction is left to right or right to left. A tab’s location, on the other hand, is relative to the back margin. A tab set at 1.5”, for example, is at 1.5” from the right in right to left text.


// A tab in a paragraph.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextTab */

// Initializes a text tab with the specified text alignment, location, and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTab/init(textAlignment:location:options:)
func NewTextTabWithTextAlignmentLocationOptions(alignment TextAlignment, loc float64, options foundation.IDictionary) TextTab {
	instance := getTextTabClass().Alloc()
	rv := objc.Send[TextTab](instance.ID, objc.Sel("initWithTextAlignment:location:options:"), alignment, loc, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextTabWithTextAlignmentLocationOptions */


// Initializes a newly allocated text tab with the specified alignment and location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTab/init(type:location:)
func NewTextTabWithTypeLocation(type_ TextTabType, loc float64) TextTab {
	instance := getTextTabClass().Alloc()
	rv := objc.Send[TextTab](instance.ID, objc.Sel("initWithType:location:"), type_, loc)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextTabWithTypeLocation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextTab */

// Returns the column terminators for the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTab/columnTerminators(for:)
func (tc _TextTabClass) ColumnTerminatorsForLocale(aLocale foundation.Locale) foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](objc.ID(tc.class), objc.Sel("columnTerminatorsForLocale:"), aLocale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ColumnTerminatorsForLocale) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextTab */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextTab */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextTab */

// The text alignment of the text tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTab/alignment
func (t_ TextTab) Alignment() TextAlignment {
	rv := objc.Send[TextAlignment](t_.ID, objc.Sel("alignment"))
	return rv
}/* debug [instance_properties/getter]: alignment */


// The text tab’s ruler location relative to the back margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTab/location
func (t_ TextTab) Location() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("location"))
	return rv
}/* debug [instance_properties/getter]: location */


// The dictionary of attributes for the text tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTab/options
func (t_ TextTab) Options() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// The text tab’s type of tab stop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTab/tabStopType
func (t_ TextTab) TabStopType() TextTabType {
	rv := objc.Send[TextTabType](t_.ID, objc.Sel("tabStopType"))
	return rv
}/* debug [instance_properties/getter]: tabStopType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextTab */


