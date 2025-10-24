//go:build darwin && ios

// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// iOS-only methods for ToolPickerCustomItem


// Requests a new image for the custom tool item from the image provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/reloadImage()
func (t_ ToolPickerCustomItem) ReloadImage() {
	objc.Send[objc.ID](t_.ID, objc.Sel("reloadImage"))
}

// iOS-only properties

// Present color selection UI to the user. Defaults to the value set in .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/allowsColorSelection
func (t_ ToolPickerCustomItem) AllowsColorSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsColorSelection"))
	return rv
}
func (t_ ToolPickerCustomItem) SetAllowsColorSelection(value bool) {
	t_.ID.Send(objc.RegisterName("setAllowsColorSelection:"), value)
}

// The current color of the custom tool item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/color
func (t_ ToolPickerCustomItem) Color() appkit.Color {
	rv := objc.Send[appkit.Color](t_.ID, objc.Sel("color"))
	return rv
}
func (t_ ToolPickerCustomItem) SetColor(value appkit.Color) {
	t_.ID.Send(objc.RegisterName("setColor:"), value)
}

// The configuration of the custom tool item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/configuration-v7e5
func (t_ ToolPickerCustomItem) Configuration() IPKToolPickerCustomItemConfiguration {
	rv := objc.Send[ToolPickerCustomItemConfiguration](t_.ID, objc.Sel("configuration"))
	return rv
}

// The current width of the custom tool item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/width
func (t_ ToolPickerCustomItem) Width() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("width"))
	return rv
}
func (t_ ToolPickerCustomItem) SetWidth(value float64) {
	t_.ID.Send(objc.RegisterName("setWidth:"), value)
}




