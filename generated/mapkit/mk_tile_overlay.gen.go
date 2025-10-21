// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKTileOverlay] class.
var (
	MKTileOverlayClass     _MKTileOverlayClass
	MKTileOverlayClassOnce sync.Once
)

func getMKTileOverlayClass() _MKTileOverlayClass {
	MKTileOverlayClassOnce.Do(func() {
		MKTileOverlayClass = _MKTileOverlayClass{objc.GetClass("MKTileOverlay")}
	})
	return MKTileOverlayClass
}

type _MKTileOverlayClass struct {
	class objc.Class
}

// An interface definition for the [MKTileOverlay] class.
type IMKTileOverlay interface {
	objectivec.IObject
}

// An overlay that covers an area of the map with tiles of bitmap images.
//
// You use tile overlay objects to represent your own tile-based content and to coordinate the display of that content in a map view. Your tiles can supplement the underlying map content or replace it completely. A tile overlay object coordinates the loading and management of the tiles, and a corresponding object handles the actual drawing of the tiles on the map. You can use a single tile overlay object to represent all of the tiles at one or more zoom levels of the map. The default tile overlay object uses a template string to build URLs so that it can locate the map tiles it needs. Each URL incorporates the x and y index of the map tile, the zoom level it’s intended for, and the scale factor corresponding to the screen resolution on which to display the tile. The default class lets you specify map tiles with indexes that start in either the upper-left corner or lower-left corner of the map. If you use a different indexing scheme for your tiles, you can also subclass and override the or methods to map between the requested tile and your custom indexing scheme.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay
type MKTileOverlay struct {
	objectivec.Object
}

// MKTileOverlayFrom constructs a [MKTileOverlay] from an unsafe.Pointer.
//
// An overlay that covers an area of the map with tiles of bitmap images.
func MKTileOverlayFrom(ptr unsafe.Pointer) MKTileOverlay {
	return MKTileOverlay{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKTileOverlayClass) Alloc() MKTileOverlay {
	rv := objc.Send[MKTileOverlay](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKTileOverlayClass) New() MKTileOverlay {
	rv := objc.Send[MKTileOverlay](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKTileOverlay) Init() MKTileOverlay {
	rv := objc.Send[MKTileOverlay](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKTileOverlay) Autorelease() MKTileOverlay {
	rv := objc.Send[MKTileOverlay](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKTileOverlay creates a new MKTileOverlay instance.
func NewMKTileOverlay() MKTileOverlay {
	return getMKTileOverlayClass().New()
}


// A Boolean value that indicates whether the tile content is fully opaque.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mktileoverlay/canreplacemapcontent
func (m_ MKTileOverlay) CanReplaceMapContent() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("canReplaceMapContent"))
	return rv
}


// SetCanReplaceMapContent sets the value of the canReplaceMapContent property.
// A Boolean value that indicates whether the tile content is fully opaque.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mktileoverlay/canreplacemapcontent
func (m_ MKTileOverlay) SetCanReplaceMapContent(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCanReplaceMapContent:"), value)
}

// A Boolean value that indicates the orientation of tile indexes along the y-axis.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mktileoverlay/isgeometryflipped
func (m_ MKTileOverlay) IsGeometryFlipped() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isGeometryFlipped"))
	return rv
}


// SetIsGeometryFlipped sets the value of the isGeometryFlipped property.
// A Boolean value that indicates the orientation of tile indexes along the y-axis.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mktileoverlay/isgeometryflipped
func (m_ MKTileOverlay) SetIsGeometryFlipped(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsGeometryFlipped:"), value)
}

// The maximum zoom level that the tiles of this overlay object support.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mktileoverlay/maximumz
func (m_ MKTileOverlay) MaximumZ() int {
	rv := objc.Send[int](m_.ID, objc.Sel("maximumZ"))
	return rv
}


// SetMaximumZ sets the value of the maximumZ property.
// The maximum zoom level that the tiles of this overlay object support.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mktileoverlay/maximumz
func (m_ MKTileOverlay) SetMaximumZ(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumZ:"), value)
}

// The minimum zoom level that the tiles of this overlay object support.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mktileoverlay/minimumz
func (m_ MKTileOverlay) MinimumZ() int {
	rv := objc.Send[int](m_.ID, objc.Sel("minimumZ"))
	return rv
}


// SetMinimumZ sets the value of the minimumZ property.
// The minimum zoom level that the tiles of this overlay object support.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mktileoverlay/minimumz
func (m_ MKTileOverlay) SetMinimumZ(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimumZ:"), value)
}

// The size (in pixels) of your tile images.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mktileoverlay/tilesize
func (m_ MKTileOverlay) TileSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](m_.ID, objc.Sel("tileSize"))
	return rv
}


// SetTileSize sets the value of the tileSize property.
// The size (in pixels) of your tile images.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mktileoverlay/tilesize
func (m_ MKTileOverlay) SetTileSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTileSize:"), value)
}

// The template for generating tile image URLs.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mktileoverlay/urltemplate
func (m_ MKTileOverlay) UrlTemplate() string {
	rv := objc.Send[string](m_.ID, objc.Sel("urlTemplate"))
	return rv
}


// SetUrlTemplate sets the value of the urlTemplate property.
// The template for generating tile image URLs.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mktileoverlay/urltemplate
func (m_ MKTileOverlay) SetUrlTemplate(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUrlTemplate:"), objc.String(value))
}



