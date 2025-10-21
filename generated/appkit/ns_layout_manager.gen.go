// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [LayoutManager] class.
var (
	LayoutManagerClass     _LayoutManagerClass
	LayoutManagerClassOnce sync.Once
)

func getLayoutManagerClass() _LayoutManagerClass {
	LayoutManagerClassOnce.Do(func() {
		LayoutManagerClass = _LayoutManagerClass{objc.GetClass("NSLayoutManager")}
	})
	return LayoutManagerClass
}

type _LayoutManagerClass struct {
	class objc.Class
}

// An interface definition for the [LayoutManager] class.
type ILayoutManager interface {
	objectivec.IObject
	AddTemporaryAttributeValueForCharacterRange(attrName unsafe.Pointer, value objc.ID, charRange foundation.Range)
	AddTemporaryAttributesForCharacterRange(attrs unsafe.Pointer, charRange foundation.Range)
	RemoveTemporaryAttributeForCharacterRange(attrName unsafe.Pointer, charRange foundation.Range)
}

// An object that coordinates the layout and display of text characters.
//
// maps Unicode character codes to glyphs, sets the glyphs in a series of objects, and displays them in a series of objects. In addition to its core function of laying out text, a layout manager object coordinates its text view objects, provides services to those text views to support instances for editing paragraph styles, and handles the layout and display of text attributes not inherent in glyphs (such as underline or strikethrough). You can create a subclass of to handle additional text attributes, whether inherent or not.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager
type LayoutManager struct {
	objectivec.Object
}

// LayoutManagerFrom constructs a [LayoutManager] from an unsafe.Pointer.
//
// An object that coordinates the layout and display of text characters.
func LayoutManagerFrom(ptr unsafe.Pointer) LayoutManager {
	return LayoutManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LayoutManagerClass) Alloc() LayoutManager {
	rv := objc.Send[LayoutManager](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LayoutManagerClass) New() LayoutManager {
	rv := objc.Send[LayoutManager](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LayoutManager) Init() LayoutManager {
	rv := objc.Send[LayoutManager](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LayoutManager) Autorelease() LayoutManager {
	rv := objc.Send[LayoutManager](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayoutManager creates a new LayoutManager instance.
func NewLayoutManager() LayoutManager {
	return getLayoutManagerClass().New()
}


// Adds a temporary attribute to the characters in the specified range.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/addTemporaryAttribute(_:value:forCharacterRange:)
func (l_ LayoutManager) AddTemporaryAttributeValueForCharacterRange(attrName unsafe.Pointer, value objc.ID, charRange foundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("addTemporaryAttribute:value:forCharacterRange:"), attrName, value, charRange)
}

// Appends one or more temporary attributes to the attributes dictionary of the specified character range.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/addTemporaryAttributes(_:forCharacterRange:)
func (l_ LayoutManager) AddTemporaryAttributesForCharacterRange(attrs unsafe.Pointer, charRange foundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("addTemporaryAttributes:forCharacterRange:"), attrs, charRange)
}

// Removes a temporary attribute from the list of attributes for the specified character range.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/removeTemporaryAttribute(_:forCharacterRange:)
func (l_ LayoutManager) RemoveTemporaryAttributeForCharacterRange(attrName unsafe.Pointer, charRange foundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("removeTemporaryAttribute:forCharacterRange:"), attrName, charRange)
}

// The default typesetter behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/typesetterBehavior-swift.property
func (l_ LayoutManager) TypesetterBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("typesetterBehavior"))
	return rv
}


// SetTypesetterBehavior sets the value of the typesetterBehavior property.
// The default typesetter behavior.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/typesetterBehavior-swift.property
func (l_ LayoutManager) SetTypesetterBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTypesetterBehavior:"), value)
}



