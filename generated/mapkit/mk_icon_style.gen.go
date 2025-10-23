// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKIconStyle] class.
var (
	MKIconStyleClass     _MKIconStyleClass
	MKIconStyleClassOnce sync.Once
)

func getMKIconStyleClass() _MKIconStyleClass {
	MKIconStyleClassOnce.Do(func() {
		MKIconStyleClass = _MKIconStyleClass{objc.GetClass("MKIconStyle")}
	})
	return MKIconStyleClass
}

type _MKIconStyleClass struct {
	class objc.Class
}

// An interface definition for the [MKIconStyle] class.
type IMKIconStyle interface {
	objectivec.IObject
	// properties:
	BackgroundColor() objc.IObject /* cross-framework: Color */
	SetBackgroundColor(value objc.IObject /* cross-framework: Color */)
	Image() objc.IObject /* cross-framework: Image */
	SetImage(value objc.IObject /* cross-framework: Image */)
	// methods:
}

// A class you use to customize the annotation view icon of a point of interest (POI) on the map.


// A class you use to customize the annotation view icon of a point of interest (POI) on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKIconStyle
type MKIconStyle struct {
	objectivec.Object
}

// MKIconStyleFrom constructs a [MKIconStyle] from an unsafe.Pointer.
//
// A class you use to customize the annotation view icon of a point of interest (POI) on the map.
func MKIconStyleFrom(ptr unsafe.Pointer) MKIconStyle {
	return MKIconStyle{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKIconStyleClass) Alloc() MKIconStyle {
	rv := objc.Send[MKIconStyle](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKIconStyleClass) New() MKIconStyle {
	rv := objc.Send[MKIconStyle](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKIconStyle) Init() MKIconStyle {
	rv := objc.Send[MKIconStyle](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKIconStyle) Autorelease() MKIconStyle {
	rv := objc.Send[MKIconStyle](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKIconStyle creates a new MKIconStyle instance.
func NewMKIconStyle() MKIconStyle {
	return getMKIconStyleClass().New()
}



// The background color of the icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkiconstyle/backgroundcolor
func (m_ MKIconStyle) BackgroundColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[Color](m_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The background color of the icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkiconstyle/backgroundcolor
func (m_ MKIconStyle) SetBackgroundColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The icon image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkiconstyle/image
func (m_ MKIconStyle) Image() objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[Image](m_.ID, objc.Sel("image"))
	return rv
}


// The icon image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkiconstyle/image
func (m_ MKIconStyle) SetImage(value objc.IObject /* cross-framework: Image */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImage:"), value)
}



