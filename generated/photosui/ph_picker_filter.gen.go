// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHPickerFilter] class.
var (
	PHPickerFilterClass     _PHPickerFilterClass
	PHPickerFilterClassOnce sync.Once
)

func getPHPickerFilterClass() _PHPickerFilterClass {
	PHPickerFilterClassOnce.Do(func() {
		PHPickerFilterClass = _PHPickerFilterClass{objc.GetClass("PHPickerFilter")}
	})
	return PHPickerFilterClass
}

type _PHPickerFilterClass struct {
	class objc.Class
}

// An interface definition for the [PHPickerFilter] class.
type IPHPickerFilter interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A type that defines the filter to apply to the photo library.

// A type that defines the filter to apply to the photo library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class
type PHPickerFilter struct {
	objectivec.Object
}

// PHPickerFilterFrom constructs a [PHPickerFilter] from an unsafe.Pointer.
//
// A type that defines the filter to apply to the photo library.
func PHPickerFilterFrom(ptr unsafe.Pointer) PHPickerFilter {
	return PHPickerFilter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHPickerFilterClass) Alloc() PHPickerFilter {
	rv := objc.Send[PHPickerFilter](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHPickerFilterClass) New() PHPickerFilter {
	rv := objc.Send[PHPickerFilter](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHPickerFilter) Init() PHPickerFilter {
	rv := objc.Send[PHPickerFilter](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHPickerFilter) Autorelease() PHPickerFilter {
	rv := objc.Send[PHPickerFilter](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHPickerFilter creates a new PHPickerFilter instance.
func NewPHPickerFilter() PHPickerFilter {
	return getPHPickerFilterClass().New()
}
