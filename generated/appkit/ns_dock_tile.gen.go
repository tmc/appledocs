// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DockTile] class.
var (
	DockTileClass     _DockTileClass
	DockTileClassOnce sync.Once
)

func getDockTileClass() _DockTileClass {
	DockTileClassOnce.Do(func() {
		DockTileClass = _DockTileClass{objc.GetClass("NSDockTile")}
	})
	return DockTileClass
}

type _DockTileClass struct {
	class objc.Class
}

// An interface definition for the [DockTile] class.
type IDockTile interface {
	objectivec.IObject
}

// The visual representation of your app’s miniaturized windows and app icon as they appear in the Dock.
//
// You do not create Dock tile objects explicitly in your app. Instead, you retrieve the Dock tile for an existing window or for the app by calling that object’s method. Also, you do not subclass the class; instead, you use the methods of the class to make the following customizations: Badge the tile with a custom string. Remove or show the application icon badge. Draw the tile content yourself. If you decide to draw the tile content yourself, you must provide a custom content view to handle the drawing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDockTile
type DockTile struct {
	objectivec.Object
}

// DockTileFrom constructs a [DockTile] from an unsafe.Pointer.
//
// The visual representation of your app’s miniaturized windows and app icon as they appear in the Dock.
func DockTileFrom(ptr unsafe.Pointer) DockTile {
	return DockTile{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DockTileClass) Alloc() DockTile {
	rv := objc.Send[DockTile](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DockTileClass) New() DockTile {
	rv := objc.Send[DockTile](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DockTile) Init() DockTile {
	rv := objc.Send[DockTile](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DockTile) Autorelease() DockTile {
	rv := objc.Send[DockTile](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDockTile creates a new DockTile instance.
func NewDockTile() DockTile {
	return getDockTileClass().New()
}


// The image used for the app’s icon.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/applicationiconimage
func (d_ DockTile) ApplicationIconImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("applicationIconImage"))
	return rv
}


// SetApplicationIconImage sets the value of the applicationIconImage property.
// The image used for the app’s icon.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/applicationiconimage
func (d_ DockTile) SetApplicationIconImage(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setApplicationIconImage:"), value)
}

// The string to be displayed in the tile’s badging area.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/badgelabel
func (d_ DockTile) BadgeLabel() string {
	rv := objc.Send[string](d_.ID, objc.Sel("badgeLabel"))
	return rv
}


// SetBadgeLabel sets the value of the badgeLabel property.
// The string to be displayed in the tile’s badging area.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/badgelabel
func (d_ DockTile) SetBadgeLabel(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBadgeLabel:"), objc.String(value))
}

// The view to use for drawing the dock tile contents.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/contentview
func (d_ DockTile) ContentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("contentView"))
	return rv
}


// SetContentView sets the value of the contentView property.
// The view to use for drawing the dock tile contents.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/contentview
func (d_ DockTile) SetContentView(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContentView:"), value)
}

// The object represented by the dock tile.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/owner
func (d_ DockTile) Owner() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("owner"))
	return rv
}


// SetOwner sets the value of the owner property.
// The object represented by the dock tile.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/owner
func (d_ DockTile) SetOwner(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setOwner:"), value)
}

// A Boolean showing whether the tile is badged with the application’s icon
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/showsapplicationbadge
func (d_ DockTile) ShowsApplicationBadge() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("showsApplicationBadge"))
	return rv
}


// SetShowsApplicationBadge sets the value of the showsApplicationBadge property.
// A Boolean showing whether the tile is badged with the application’s icon

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/showsapplicationbadge
func (d_ DockTile) SetShowsApplicationBadge(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShowsApplicationBadge:"), value)
}

// The size of the tile.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/size
func (d_ DockTile) Size() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](d_.ID, objc.Sel("size"))
	return rv
}


// SetSize sets the value of the size property.
// The size of the tile.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/size
func (d_ DockTile) SetSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSize:"), value)
}

// The application’s Dock tile.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/docktile
func (d_ DockTile) DockTile() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("dockTile"))
	return rv
}


// SetDockTile sets the value of the dockTile property.
// The application’s Dock tile.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/docktile
func (d_ DockTile) SetDockTile(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDockTile:"), value)
}



