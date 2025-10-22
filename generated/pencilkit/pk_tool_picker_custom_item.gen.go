// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [ToolPickerCustomItem] class.
var (
	ToolPickerCustomItemClass     _ToolPickerCustomItemClass
	ToolPickerCustomItemClassOnce sync.Once
)

func getToolPickerCustomItemClass() _ToolPickerCustomItemClass {
	ToolPickerCustomItemClassOnce.Do(func() {
		ToolPickerCustomItemClass = _ToolPickerCustomItemClass{objc.GetClass("PKToolPickerCustomItem")}
	})
	return ToolPickerCustomItemClass
}

type _ToolPickerCustomItemClass struct {
	class objc.Class
}

// An interface definition for the [ToolPickerCustomItem] class.
type IToolPickerCustomItem interface {
	IToolPickerItem
	ReloadImage()
	AllowsColorSelection() bool
	SetAllowsColorSelection(value bool)
	Color() appkit.Color
	SetColor(value appkit.IColor)
	Configuration() PKToolPickerCustomItemConfiguration
	Width() float64
	SetWidth(value float64)
	ImageProvider() appkit.Image
	SetImageProvider(value appkit.IImage)
}

// An item that represents a custom tool in the tool picker.
//
// A custom tool item represents a tool that isn’t one of the system tools. You configure details about a custom tool item yourself using , including providing custom images to draw the body of the tool. The following code shows how to create a tool picker with a custom tool item. This basic implementation of retrieves an image for the tool body from an asset catalog. A full app might use a more advanced drawing implementation for the image provider, such as using . For a more complete example of creating a custom tool item, see .
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem
type ToolPickerCustomItem struct {
	ToolPickerItem
}

// ToolPickerCustomItemFrom constructs a [ToolPickerCustomItem] from an unsafe.Pointer.
//
// An item that represents a custom tool in the tool picker.
func ToolPickerCustomItemFrom(ptr unsafe.Pointer) ToolPickerCustomItem {
	return ToolPickerCustomItem{
		ToolPickerItem: ToolPickerItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _ToolPickerCustomItemClass) Alloc() ToolPickerCustomItem {
	rv := objc.Send[ToolPickerCustomItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ToolPickerCustomItemClass) New() ToolPickerCustomItem {
	rv := objc.Send[ToolPickerCustomItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ToolPickerCustomItem) Init() ToolPickerCustomItem {
	rv := objc.Send[ToolPickerCustomItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ToolPickerCustomItem) Autorelease() ToolPickerCustomItem {
	rv := objc.Send[ToolPickerCustomItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewToolPickerCustomItem creates a new ToolPickerCustomItem instance.
func NewToolPickerCustomItem() ToolPickerCustomItem {
	return getToolPickerCustomItemClass().New()
}




// Creates a new custom item with the specified configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/initWithConfiguration:
func NewToolPickerCustomItemWithConfiguration(configuration IPKToolPickerCustomItemConfiguration) ToolPickerCustomItem {
	instance := getToolPickerCustomItemClass().Alloc()
	rv := objc.Send[ToolPickerCustomItem](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
}


// Requests a new image for the custom tool item from the image provider.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/reloadImage()
func (t_ ToolPickerCustomItem) ReloadImage() {
	objc.Send[objc.ID](t_.ID, objc.Sel("reloadImage"))
}

// Present color selection UI to the user. Defaults to the value set in .
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/allowsColorSelection
func (t_ ToolPickerCustomItem) AllowsColorSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsColorSelection"))
	return rv
}


// SetAllowsColorSelection sets the value of the allowsColorSelection property.
// Present color selection UI to the user. Defaults to the value set in .

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/allowsColorSelection
func (t_ ToolPickerCustomItem) SetAllowsColorSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsColorSelection:"), value)
}

// The current color of the custom tool item.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/color
func (t_ ToolPickerCustomItem) Color() appkit.Color {
	rv := objc.Send[appkit.Color](t_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
// The current color of the custom tool item.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/color
func (t_ ToolPickerCustomItem) SetColor(value appkit.IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColor:"), value)
}

// The configuration of the custom tool item.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/configuration-v7e5
func (t_ ToolPickerCustomItem) Configuration() PKToolPickerCustomItemConfiguration {
	rv := objc.Send[PKToolPickerCustomItemConfiguration](t_.ID, objc.Sel("configuration"))
	return rv
}

// The current width of the custom tool item.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/width
func (t_ ToolPickerCustomItem) Width() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("width"))
	return rv
}


// SetWidth sets the value of the width property.
// The current width of the custom tool item.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/width
func (t_ ToolPickerCustomItem) SetWidth(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWidth:"), value)
}

// A closure to provide an image that represents the custom tool item.
//
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pktoolpickercustomitem/configuration-swift.struct/imageprovider
func (t_ ToolPickerCustomItem) ImageProvider() appkit.Image {
	rv := objc.Send[appkit.Image](t_.ID, objc.Sel("imageProvider"))
	return rv
}


// SetImageProvider sets the value of the imageProvider property.
// A closure to provide an image that represents the custom tool item.

//
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pktoolpickercustomitem/configuration-swift.struct/imageprovider
func (t_ ToolPickerCustomItem) SetImageProvider(value appkit.IImage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImageProvider:"), value)
}


