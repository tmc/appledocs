// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ColorList] class.
var (
	ColorListClass     _ColorListClass
	ColorListClassOnce sync.Once
)

func getColorListClass() _ColorListClass {
	ColorListClassOnce.Do(func() {
		ColorListClass = _ColorListClass{objc.GetClass("NSColorList")}
	})
	return ColorListClass
}

type _ColorListClass struct {
	class objc.Class
}

// An interface definition for the [ColorList] class.
type IColorList interface {
	objectivec.IObject
	WriteToURLError(url unsafe.Pointer, errPtr unsafe.Pointer) bool
}

// An ordered list of color objects, identified by keys.
//
// A color list manages a list of objects, each of which has an associated name. The list mode color picker uses instances of to represent any lists of colors that come with the system, as well as any lists the user creates. An app can use a color list to manage document-specific color lists.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorList
type ColorList struct {
	objectivec.Object
}

// ColorListFrom constructs a [ColorList] from an unsafe.Pointer.
//
// An ordered list of color objects, identified by keys.
func ColorListFrom(ptr unsafe.Pointer) ColorList {
	return ColorList{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ColorListClass) Alloc() ColorList {
	rv := objc.Send[ColorList](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ColorListClass) New() ColorList {
	rv := objc.Send[ColorList](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ColorList) Init() ColorList {
	rv := objc.Send[ColorList](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ColorList) Autorelease() ColorList {
	rv := objc.Send[ColorList](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewColorList creates a new ColorList instance.
func NewColorList() ColorList {
	return getColorListClass().New()
}

// Saves the color list to the file at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorList/write(to:)
func (c_ ColorList) WriteToURLError(url unsafe.Pointer, errPtr unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("writeToURL:error:"), url, errPtr)
	return rv
}

// A Boolean value that indicates whether the color list can be modified.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorList/isEditable
func (c_ ColorList) Editable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("editable"))
	return rv
}
