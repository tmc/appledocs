// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [PathComponentCell] class.
var (
	PathComponentCellClass     _PathComponentCellClass
	PathComponentCellClassOnce sync.Once
)

func getPathComponentCellClass() _PathComponentCellClass {
	PathComponentCellClassOnce.Do(func() {
		PathComponentCellClass = _PathComponentCellClass{objc.GetClass("NSPathComponentCell")}
	})
	return PathComponentCellClass
}

type _PathComponentCellClass struct {
	class objc.Class
}





// An interface definition for the [PathComponentCell] class.
type IPathComponentCell interface {
	ITextFieldCell
	

	// properties:
	Image() IImage
	SetImage(value IImage)
	URL() foundation.foundation.INSURL
	SetURL(value foundation.foundation.INSURL)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PathComponentCellClass) Alloc() PathComponentCell {
	rv := objc.Send[PathComponentCell](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PathComponentCellClass) New() PathComponentCell {
	rv := objc.Send[PathComponentCell](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PathComponentCell) Init() PathComponentCell {
	rv := objc.Send[PathComponentCell](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PathComponentCell) Autorelease() PathComponentCell {
	rv := objc.Send[PathComponentCell](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPathComponentCell creates a new PathComponentCell instance.
func NewPathComponentCell() PathComponentCell {
	return getPathComponentCellClass().New()
}





// A component of a path.
//
// An object manages a collection of objects, in conjunction with an object, to represent a path.


// A component of a path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathComponentCell
type PathComponentCell struct {
	TextFieldCell
}

// PathComponentCellFrom constructs a [PathComponentCell] from an unsafe.Pointer.
//
// A component of a path.
func PathComponentCellFrom(ptr unsafe.Pointer) PathComponentCell {
	return PathComponentCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}

























// The image displayed for this component cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathComponentCell/image
func (p_ PathComponentCell) Image() IImage {
	rv := objc.Send[Image](p_.ID, objc.Sel("image"))
	return rv
}


// The image displayed for this component cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathComponentCell/image
func (p_ PathComponentCell) SetImage(value IImage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImage:"), value)
}


// The portion of the path from the root through the component represented by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathComponentCell/url
func (p_ PathComponentCell) URL() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](p_.ID, objc.Sel("URL"))
	return rv
}


// The portion of the path from the root through the component represented by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathComponentCell/url
func (p_ PathComponentCell) SetURL(value foundation.foundation.INSURL) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setURL:"), value)
}








