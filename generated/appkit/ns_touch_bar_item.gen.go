// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [TouchBarItem] class.
var (
	TouchBarItemClass     _TouchBarItemClass
	TouchBarItemClassOnce sync.Once
)

func getTouchBarItemClass() _TouchBarItemClass {
	TouchBarItemClassOnce.Do(func() {
		TouchBarItemClass = _TouchBarItemClass{objc.GetClass("NSTouchBarItem")}
	})
	return TouchBarItemClass
}

type _TouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [TouchBarItem] class.
type ITouchBarItem interface {
	objectivec.IObject
}

// A UI control shown in the Touch Bar on supported models of MacBook Pro.
//
// An instance of the class is called an . It appears to the user on the Touch Bar, typically along with other items, within the (invisible) bounds of the view for an object, called a . You use an item by adding it or its identifier to one or another of a bar’s arrays, depending on your app’s architecture and on the user customization you want to support. Because of the close interaction between bars and items, be sure you have read the overview for the class before continuing here to learn about items. AppKit provides a rich set of subclasses of , each of which is described in the corresponding class reference document: An object (a ), along with its delegate, provides a list of textual suggestions for the current text view An object (a ) provides a system-defined color picker An object (a ) contains a responder of your choice, such as a view, a button, or a scrubber (an instance of the class) An object (a ) provides a bar to contain other items An object (a ) provides a two-state control that, when touched or pressed, expands into its second state, showing the contents of a bar it owns An object (a ), along with its delegate, provides a list of objects eligible for sharing An object (a ) provides a slider control for choosing a value in a range The two most commonly-used item classes are and . Refer to the following sample code projects which demonstrate how to use and related classes:
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem
type TouchBarItem struct {
	objectivec.Object
}

// TouchBarItemFrom constructs a [TouchBarItem] from an unsafe.Pointer.
//
// A UI control shown in the Touch Bar on supported models of MacBook Pro.
func TouchBarItemFrom(ptr unsafe.Pointer) TouchBarItem {
	return TouchBarItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TouchBarItemClass) Alloc() TouchBarItem {
	rv := objc.Send[TouchBarItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TouchBarItemClass) New() TouchBarItem {
	rv := objc.Send[TouchBarItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TouchBarItem) Init() TouchBarItem {
	rv := objc.Send[TouchBarItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TouchBarItem) Autorelease() TouchBarItem {
	rv := objc.Send[TouchBarItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTouchBarItem creates a new TouchBarItem instance.
func NewTouchBarItem() TouchBarItem {
	return getTouchBarItemClass().New()
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/customizationLabel
func (t_ TouchBarItem) CustomizationLabel() string {
	rv := objc.Send[string](t_.ID, objc.Sel("customizationLabel"))
	return rv
}

// The view associated with this item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/view
func (t_ TouchBarItem) View() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("view"))
	return rv
}

// The view controller associated with this item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/viewController
func (t_ TouchBarItem) ViewController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("viewController"))
	return rv
}

// Determines which items are shown in a bar when space is limited.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/visibilityPriority
func (t_ TouchBarItem) VisibilityPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("visibilityPriority"))
	return rv
}


// SetVisibilityPriority sets the value of the visibilityPriority property.
// Determines which items are shown in a bar when space is limited.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/visibilityPriority
func (t_ TouchBarItem) SetVisibilityPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisibilityPriority:"), value)
}



