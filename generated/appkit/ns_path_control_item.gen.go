// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PathControlItem] class.
var (
	PathControlItemClass     _PathControlItemClass
	PathControlItemClassOnce sync.Once
)

func getPathControlItemClass() _PathControlItemClass {
	PathControlItemClassOnce.Do(func() {
		PathControlItemClass = _PathControlItemClass{objc.GetClass("NSPathControlItem")}
	})
	return PathControlItemClass
}

type _PathControlItemClass struct {
	class objc.Class
}

// An interface definition for the [PathControlItem] class.
type IPathControlItem interface {
	objectivec.IObject
	// properties:
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.AttributedString)
	Image() IImage
	SetImage(value IImage)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	URL() objc.IObject /* cross-framework: NSURL */
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem
type PathControlItem struct {
	objectivec.Object
}

// PathControlItemFrom constructs a [PathControlItem] from an unsafe.Pointer.
func PathControlItemFrom(ptr unsafe.Pointer) PathControlItem {
	return PathControlItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PathControlItemClass) Alloc() PathControlItem {
	rv := objc.Send[PathControlItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PathControlItemClass) New() PathControlItem {
	rv := objc.Send[PathControlItem](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PathControlItem) Init() PathControlItem {
	rv := objc.Send[PathControlItem](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PathControlItem) Autorelease() PathControlItem {
	rv := objc.Send[PathControlItem](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPathControlItem creates a new PathControlItem instance.
func NewPathControlItem() PathControlItem {
	return getPathControlItemClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem/attributedTitle
func (p_ PathControlItem) AttributedTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](p_.ID, objc.Sel("attributedTitle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem/attributedTitle
func (p_ PathControlItem) SetAttributedTitle(value foundation.AttributedString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAttributedTitle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem/image
func (p_ PathControlItem) Image() IImage {
	rv := objc.Send[Image](p_.ID, objc.Sel("image"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem/image
func (p_ PathControlItem) SetImage(value IImage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImage:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem/title
func (p_ PathControlItem) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("title"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem/title
func (p_ PathControlItem) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem/url
func (p_ PathControlItem) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](p_.ID, objc.Sel("URL"))
	return rv
}



