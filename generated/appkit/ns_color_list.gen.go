// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	WriteToURLError(url foundation.IURL, errPtr unsafe.Pointer) bool
	Editable() bool
	AllKeys() unsafe.Pointer
	SetAllKeys(value unsafe.Pointer)
	IsEditable() bool
	SetIsEditable(value bool)
	Name() unsafe.Pointer
	SetName(value unsafe.Pointer)
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
func (c_ ColorList) WriteToURLError(url foundation.IURL, errPtr unsafe.Pointer) bool {
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

// An array of the keys by which the color objects are stored in the color list.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/allkeys
func (c_ ColorList) AllKeys() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("allKeys"))
	return rv
}


// SetAllKeys sets the value of the allKeys property.
// An array of the keys by which the color objects are stored in the color list.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/allkeys
func (c_ ColorList) SetAllKeys(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllKeys:"), value)
}

// A Boolean value that indicates whether the color list can be modified.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/iseditable
func (c_ ColorList) IsEditable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEditable"))
	return rv
}


// SetIsEditable sets the value of the isEditable property.
// A Boolean value that indicates whether the color list can be modified.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/iseditable
func (c_ ColorList) SetIsEditable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEditable:"), value)
}

// The name of the color list.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/name-swift.property
func (c_ ColorList) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the color list.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/name-swift.property
func (c_ ColorList) SetName(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), value)
}



