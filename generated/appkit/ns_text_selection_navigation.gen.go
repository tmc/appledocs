// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextSelectionNavigation] class.
var (
	TextSelectionNavigationClass     _TextSelectionNavigationClass
	TextSelectionNavigationClassOnce sync.Once
)

func getTextSelectionNavigationClass() _TextSelectionNavigationClass {
	TextSelectionNavigationClassOnce.Do(func() {
		TextSelectionNavigationClass = _TextSelectionNavigationClass{objc.GetClass("NSTextSelectionNavigation")}
	})
	return TextSelectionNavigationClass
}

type _TextSelectionNavigationClass struct {
	class objc.Class
}

// An interface definition for the [TextSelectionNavigation] class.
type ITextSelectionNavigation interface {
	objectivec.IObject
}

// An interface you use to expose methods for obtaining results from actions performed on text selections.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation
type TextSelectionNavigation struct {
	objectivec.Object
}

// TextSelectionNavigationFrom constructs a [TextSelectionNavigation] from an unsafe.Pointer.
//
// An interface you use to expose methods for obtaining results from actions performed on text selections.
func TextSelectionNavigationFrom(ptr unsafe.Pointer) TextSelectionNavigation {
	return TextSelectionNavigation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextSelectionNavigationClass) Alloc() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextSelectionNavigationClass) New() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextSelectionNavigation) Init() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextSelectionNavigation) Autorelease() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextSelectionNavigation creates a new TextSelectionNavigation instance.
func NewTextSelectionNavigation() TextSelectionNavigation {
	return getTextSelectionNavigationClass().New()
}


// Determines if the instance could produce selections with multiple noncontiguous selections.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselectionnavigation/allowsnoncontiguousranges
func (t_ TextSelectionNavigation) AllowsNonContiguousRanges() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsNonContiguousRanges"))
	return rv
}


// SetAllowsNonContiguousRanges sets the value of the allowsNonContiguousRanges property.
// Determines if the instance could produce selections with multiple noncontiguous selections.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselectionnavigation/allowsnoncontiguousranges
func (t_ TextSelectionNavigation) SetAllowsNonContiguousRanges(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsNonContiguousRanges:"), value)
}

// Determines if the framework rotates the coordinate system to match the layout orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselectionnavigation/rotatescoordinatesystemforlayoutorientation
func (t_ TextSelectionNavigation) RotatesCoordinateSystemForLayoutOrientation() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rotatesCoordinateSystemForLayoutOrientation"))
	return rv
}


// SetRotatesCoordinateSystemForLayoutOrientation sets the value of the rotatesCoordinateSystemForLayoutOrientation property.
// Determines if the framework rotates the coordinate system to match the layout orientation.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselectionnavigation/rotatescoordinatesystemforlayoutorientation
func (t_ TextSelectionNavigation) SetRotatesCoordinateSystemForLayoutOrientation(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRotatesCoordinateSystemForLayoutOrientation:"), value)
}

// The data source associated with this selection navigation.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselectionnavigation/textselectiondatasource
func (t_ TextSelectionNavigation) TextSelectionDataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textSelectionDataSource"))
	return rv
}


// SetTextSelectionDataSource sets the value of the textSelectionDataSource property.
// The data source associated with this selection navigation.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextselectionnavigation/textselectiondatasource
func (t_ TextSelectionNavigation) SetTextSelectionDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextSelectionDataSource:"), value)
}



