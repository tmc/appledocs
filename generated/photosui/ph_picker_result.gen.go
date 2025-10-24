// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHPickerResult] class.
var (
	PHPickerResultClass     _PHPickerResultClass
	PHPickerResultClassOnce sync.Once
)

func getPHPickerResultClass() _PHPickerResultClass {
	PHPickerResultClassOnce.Do(func() {
		PHPickerResultClass = _PHPickerResultClass{objc.GetClass("PHPickerResult")}
	})
	return PHPickerResultClass
}

type _PHPickerResultClass struct {
	class objc.Class
}

// An interface definition for the [PHPickerResult] class.
type IPHPickerResult interface {
	objectivec.IObject
	// properties:
	// methods:
}

// Types that represent a selected asset from the user’s photo library.

// Types that represent a selected asset from the user’s photo library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerResult-c.class
type PHPickerResult struct {
	objectivec.Object
}

// PHPickerResultFrom constructs a [PHPickerResult] from an unsafe.Pointer.
//
// Types that represent a selected asset from the user’s photo library.
func PHPickerResultFrom(ptr unsafe.Pointer) PHPickerResult {
	return PHPickerResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHPickerResultClass) Alloc() PHPickerResult {
	rv := objc.Send[PHPickerResult](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHPickerResultClass) New() PHPickerResult {
	rv := objc.Send[PHPickerResult](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHPickerResult) Init() PHPickerResult {
	rv := objc.Send[PHPickerResult](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHPickerResult) Autorelease() PHPickerResult {
	rv := objc.Send[PHPickerResult](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHPickerResult creates a new PHPickerResult instance.
func NewPHPickerResult() PHPickerResult {
	return getPHPickerResultClass().New()
}
