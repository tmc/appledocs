// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKTileOverlay */


/* debug [class_header]: Header for MKTileOverlay */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKTileOverlay */
// An interface definition for the [MKTileOverlay] class.
type IMKTileOverlay interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKTileOverlay */
	// properties:
	CanReplaceMapContent() bool
	SetCanReplaceMapContent(value bool)
	GeometryFlipped() bool
	SetGeometryFlipped(value bool)
	MaximumZ() int
	SetMaximumZ(value int)
	MinimumZ() int
	SetMinimumZ(value int)
	TileSize() corefoundation.CGSize
	SetTileSize(value corefoundation.CGSize)
	URLTemplate() objc.IObject /* cross-framework: NSString */
	IsGeometryFlipped() bool
	SetIsGeometryFlipped(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKTileOverlay */
	// methods:
	LoadTileAtPathResult(path objc.IObject /* cross-framework: MKTileOverlayPath */, result unsafe.Pointer)
	URLForTilePath(path objc.IObject /* cross-framework: MKTileOverlayPath */) foundation.URL
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKTileOverlay */
// Alloc allocates a new instance without initialization.
func (mc _MKTileOverlayClass) Alloc() MKTileOverlay {
	rv := objc.Send[MKTileOverlay](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKTileOverlay */
// An overlay that covers an area of the map with tiles of bitmap images.
//
// You use tile overlay objects to represent your own tile-based content and to coordinate the display of that content in a map view. Your tiles can supplement the underlying map content or replace it completely. A tile overlay object coordinates the loading and management of the tiles, and a corresponding object handles the actual drawing of the tiles on the map. You can use a single tile overlay object to represent all of the tiles at one or more zoom levels of the map. The default tile overlay object uses a template string to build URLs so that it can locate the map tiles it needs. Each URL incorporates the x and y index of the map tile, the zoom level it’s intended for, and the scale factor corresponding to the screen resolution on which to display the tile. The default class lets you specify map tiles with indexes that start in either the upper-left corner or lower-left corner of the map. If you use a different indexing scheme for your tiles, you can also subclass and override the or methods to map between the requested tile and your custom indexing scheme.


// An overlay that covers an area of the map with tiles of bitmap images.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKTileOverlay */

// Creates and returns a tile overlay object using the specified tile-access template.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay/init(urlTemplate:)
func NewMKTileOverlayWithURLTemplate(URLTemplate objc.IObject /* cross-framework: NSString */) MKTileOverlay {
	instance := getMKTileOverlayClass().Alloc()
	rv := objc.Send[MKTileOverlay](instance.ID, objc.Sel("initWithURLTemplate:"), URLTemplate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKTileOverlayWithURLTemplate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKTileOverlay */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKTileOverlay */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKTileOverlay */

// Loads the specified tile asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay/loadTile(at:result:)
func (m_ MKTileOverlay) LoadTileAtPathResult(path objc.IObject /* cross-framework: MKTileOverlayPath */, result unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("loadTileAtPath:result:"), path, result)
}/* debug [instance_methods/method]: LoadTileAtPathResult */


// Returns the URL to use to access the specified tile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay/url(forTilePath:)
func (m_ MKTileOverlay) URLForTilePath(path objc.IObject /* cross-framework: MKTileOverlayPath */) foundation.URL {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("URLForTilePath:"), path)
	return rv
}/* debug [instance_methods/method]: URLForTilePath */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKTileOverlay */

// A Boolean value that indicates whether the tile content is fully opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay/canReplaceMapContent
func (m_ MKTileOverlay) CanReplaceMapContent() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("canReplaceMapContent"))
	return rv
}/* debug [instance_properties/getter]: canReplaceMapContent */


// A Boolean value that indicates whether the tile content is fully opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay/canReplaceMapContent
func (m_ MKTileOverlay) SetCanReplaceMapContent(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCanReplaceMapContent:"), value)
}/* debug [instance_properties/setter]: canReplaceMapContent */


// A Boolean value that indicates the orientation of tile indexes along the y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay/isGeometryFlipped
func (m_ MKTileOverlay) GeometryFlipped() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("geometryFlipped"))
	return rv
}/* debug [instance_properties/getter]: geometryFlipped */


// A Boolean value that indicates the orientation of tile indexes along the y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay/isGeometryFlipped
func (m_ MKTileOverlay) SetGeometryFlipped(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGeometryFlipped:"), value)
}/* debug [instance_properties/setter]: geometryFlipped */


// The maximum zoom level that the tiles of this overlay object support.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay/maximumZ
func (m_ MKTileOverlay) MaximumZ() int {
	rv := objc.Send[int](m_.ID, objc.Sel("maximumZ"))
	return rv
}/* debug [instance_properties/getter]: maximumZ */


// The maximum zoom level that the tiles of this overlay object support.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay/maximumZ
func (m_ MKTileOverlay) SetMaximumZ(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumZ:"), value)
}/* debug [instance_properties/setter]: maximumZ */


// The minimum zoom level that the tiles of this overlay object support.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay/minimumZ
func (m_ MKTileOverlay) MinimumZ() int {
	rv := objc.Send[int](m_.ID, objc.Sel("minimumZ"))
	return rv
}/* debug [instance_properties/getter]: minimumZ */


// The minimum zoom level that the tiles of this overlay object support.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay/minimumZ
func (m_ MKTileOverlay) SetMinimumZ(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimumZ:"), value)
}/* debug [instance_properties/setter]: minimumZ */


// The size (in pixels) of your tile images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay/tileSize
func (m_ MKTileOverlay) TileSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m_.ID, objc.Sel("tileSize"))
	return rv
}/* debug [instance_properties/getter]: tileSize */


// The size (in pixels) of your tile images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay/tileSize
func (m_ MKTileOverlay) SetTileSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTileSize:"), value)
}/* debug [instance_properties/setter]: tileSize */


// The template for generating tile image URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlay/urlTemplate
func (m_ MKTileOverlay) URLTemplate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("URLTemplate"))
	return rv
}/* debug [instance_properties/getter]: URLTemplate */


// A Boolean value that indicates the orientation of tile indexes along the y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mktileoverlay/isgeometryflipped
func (m_ MKTileOverlay) IsGeometryFlipped() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isGeometryFlipped"))
	return rv
}/* debug [instance_properties/getter]: isGeometryFlipped */


// A Boolean value that indicates the orientation of tile indexes along the y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mktileoverlay/isgeometryflipped
func (m_ MKTileOverlay) SetIsGeometryFlipped(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsGeometryFlipped:"), value)
}/* debug [instance_properties/setter]: isGeometryFlipped */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKTileOverlay */


