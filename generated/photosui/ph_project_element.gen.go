// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHProjectElement] class.
var (
	PHProjectElementClass     _PHProjectElementClass
	PHProjectElementClassOnce sync.Once
)

func getPHProjectElementClass() _PHProjectElementClass {
	PHProjectElementClassOnce.Do(func() {
		PHProjectElementClass = _PHProjectElementClass{objc.GetClass("PHProjectElement")}
	})
	return PHProjectElementClass
}

type _PHProjectElementClass struct {
	class objc.Class
}

// An interface definition for the [PHProjectElement] class.
type IPHProjectElement interface {
	objectivec.IObject
}

// The superclass for all element objects.
//
// You should never use this class directly; opt instead for one of its subclasses. It defines the shared properties of any element in an instance of .
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectElement
type PHProjectElement struct {
	objectivec.Object
}

// PHProjectElementFrom constructs a [PHProjectElement] from an unsafe.Pointer.
//
// The superclass for all element objects.
func PHProjectElementFrom(ptr unsafe.Pointer) PHProjectElement {
	return PHProjectElement{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHProjectElementClass) Alloc() PHProjectElement {
	rv := objc.Send[PHProjectElement](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHProjectElementClass) New() PHProjectElement {
	rv := objc.Send[PHProjectElement](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHProjectElement) Init() PHProjectElement {
	rv := objc.Send[PHProjectElement](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHProjectElement) Autorelease() PHProjectElement {
	rv := objc.Send[PHProjectElement](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHProjectElement creates a new PHProjectElement instance.
func NewPHProjectElement() PHProjectElement {
	return getPHProjectElementClass().New()
}


// A rectangle defining where an element is placed in grid space coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectElement/placement
func (p_ PHProjectElement) Placement() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("placement"))
	return rv
}

// A value between 0 and 1 representing relative significance of the element in its section.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectElement/weight
func (p_ PHProjectElement) Weight() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("weight"))
	return rv
}



