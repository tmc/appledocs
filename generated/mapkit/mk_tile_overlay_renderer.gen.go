// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKTileOverlayRenderer */


/* debug [class_header]: Header for MKTileOverlayRenderer */
// The class instance for the [MKTileOverlayRenderer] class.
var (
	MKTileOverlayRendererClass     _MKTileOverlayRendererClass
	MKTileOverlayRendererClassOnce sync.Once
)

func getMKTileOverlayRendererClass() _MKTileOverlayRendererClass {
	MKTileOverlayRendererClassOnce.Do(func() {
		MKTileOverlayRendererClass = _MKTileOverlayRendererClass{objc.GetClass("MKTileOverlayRenderer")}
	})
	return MKTileOverlayRendererClass
}

type _MKTileOverlayRendererClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKTileOverlayRenderer */
// An interface definition for the [MKTileOverlayRenderer] class.
type IMKTileOverlayRenderer interface {
	IMKOverlayRenderer
	
/* debug [class_interface_properties]: Properties for MKTileOverlayRenderer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKTileOverlayRenderer */
	// methods:
	ReloadData()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKTileOverlayRenderer */
// Alloc allocates a new instance without initialization.
func (mc _MKTileOverlayRendererClass) Alloc() MKTileOverlayRenderer {
	rv := objc.Send[MKTileOverlayRenderer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKTileOverlayRendererClass) New() MKTileOverlayRenderer {
	rv := objc.Send[MKTileOverlayRenderer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKTileOverlayRenderer) Init() MKTileOverlayRenderer {
	rv := objc.Send[MKTileOverlayRenderer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKTileOverlayRenderer) Autorelease() MKTileOverlayRenderer {
	rv := objc.Send[MKTileOverlayRenderer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKTileOverlayRenderer creates a new MKTileOverlayRenderer instance.
func NewMKTileOverlayRenderer() MKTileOverlayRenderer {
	return getMKTileOverlayRendererClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKTileOverlayRenderer */
// The renderer for a tile overlay that handles the drawing of bitmap images on the map surface.
//
// You create instances of this class when tile overlays become visible on the map view. A renderer works closely with its associated tile overlay object to coordinate the loading and drawing of tiles at appropriate times. For information about how to specify the tiles to display on the map, see .


// The renderer for a tile overlay that handles the drawing of bitmap images on the map surface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlayRenderer
type MKTileOverlayRenderer struct {
	MKOverlayRenderer
}

// MKTileOverlayRendererFrom constructs a [MKTileOverlayRenderer] from an unsafe.Pointer.
//
// The renderer for a tile overlay that handles the drawing of bitmap images on the map surface.
func MKTileOverlayRendererFrom(ptr unsafe.Pointer) MKTileOverlayRenderer {
	return MKTileOverlayRenderer{
		MKOverlayRenderer: MKOverlayRendererFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKTileOverlayRenderer */

// Initializes and returns a tile renderer with the specified overlay object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlayRenderer/init(tileOverlay:)
func NewMKTileOverlayRendererWithTileOverlay(overlay IMKTileOverlay) MKTileOverlayRenderer {
	instance := getMKTileOverlayRendererClass().Alloc()
	rv := objc.Send[MKTileOverlayRenderer](instance.ID, objc.Sel("initWithTileOverlay:"), overlay)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKTileOverlayRendererWithTileOverlay */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKTileOverlayRenderer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKTileOverlayRenderer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKTileOverlayRenderer */

// Forces the tile overlay renderer to reload and redisplay the tiles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlayRenderer/reloadData()
func (m_ MKTileOverlayRenderer) ReloadData() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reloadData"))
}/* debug [instance_methods/method]: ReloadData */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKTileOverlayRenderer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKTileOverlayRenderer */


