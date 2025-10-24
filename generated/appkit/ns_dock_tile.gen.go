// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	ApplicationIconImage() IImage
	SetApplicationIconImage(value IImage)
	BadgeLabel() objc.IObject /* cross-framework: NSString */
	SetBadgeLabel(value objc.IObject /* cross-framework: NSString */)
	ContentView() IView
	SetContentView(value IView)
	ShowsApplicationBadge() bool
	SetShowsApplicationBadge(value bool)
	Size() objc.IObject /* cross-framework: Size */
	SetSize(value objc.IObject /* cross-framework: Size */)
	DockTile() IDockTile
	SetDockTile(value IDockTile)
	// methods:
	Display()
}

// The visual representation of your app’s miniaturized windows and app icon as they appear in the Dock.
//
// You do not create Dock tile objects explicitly in your app. Instead, you retrieve the Dock tile for an existing window or for the app by calling that object’s method. Also, you do not subclass the class; instead, you use the methods of the class to make the following customizations: Badge the tile with a custom string. Remove or show the application icon badge. Draw the tile content yourself. If you decide to draw the tile content yourself, you must provide a custom content view to handle the drawing.


// The visual representation of your app’s miniaturized windows and app icon as they appear in the Dock.
//
// [Full Topic]
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



// Redraws the dock tile’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDockTile/display()
func (d_ DockTile) Display() {
	objc.Send[objc.ID](d_.ID, objc.Sel("display"))
}


// The image used for the app’s icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/applicationiconimage
func (d_ DockTile) ApplicationIconImage() IImage {
	rv := objc.Send[Image](d_.ID, objc.Sel("applicationIconImage"))
	return rv
}


// The image used for the app’s icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/applicationiconimage
func (d_ DockTile) SetApplicationIconImage(value IImage) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setApplicationIconImage:"), value)
}


// The string to be displayed in the tile’s badging area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/badgelabel
func (d_ DockTile) BadgeLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("badgeLabel"))
	return rv
}


// The string to be displayed in the tile’s badging area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/badgelabel
func (d_ DockTile) SetBadgeLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBadgeLabel:"), value)
}


// The view to use for drawing the dock tile contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/contentview
func (d_ DockTile) ContentView() IView {
	rv := objc.Send[View](d_.ID, objc.Sel("contentView"))
	return rv
}


// The view to use for drawing the dock tile contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/contentview
func (d_ DockTile) SetContentView(value IView) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContentView:"), value)
}


// A Boolean showing whether the tile is badged with the application’s icon
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/showsapplicationbadge
func (d_ DockTile) ShowsApplicationBadge() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("showsApplicationBadge"))
	return rv
}


// A Boolean showing whether the tile is badged with the application’s icon
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/showsapplicationbadge
func (d_ DockTile) SetShowsApplicationBadge(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShowsApplicationBadge:"), value)
}


// The size of the tile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/size
func (d_ DockTile) Size() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](d_.ID, objc.Sel("size"))
	return rv
}


// The size of the tile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocktile/size
func (d_ DockTile) SetSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSize:"), value)
}


// The application’s Dock tile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/docktile
func (d_ DockTile) DockTile() IDockTile {
	rv := objc.Send[DockTile](d_.ID, objc.Sel("dockTile"))
	return rv
}


// The application’s Dock tile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/docktile
func (d_ DockTile) SetDockTile(value IDockTile) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDockTile:"), value)
}



