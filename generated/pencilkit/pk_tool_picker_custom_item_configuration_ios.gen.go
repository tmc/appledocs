//go:build darwin && ios

// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ToolPickerCustomItemConfiguration


// iOS-only properties

// A Boolean value that determines whether to show the color selection UI for the tool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/allowsColorSelection
func (t_ ToolPickerCustomItemConfiguration) AllowsColorSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsColorSelection"))
	return rv
}
func (t_ ToolPickerCustomItemConfiguration) SetAllowsColorSelection(value bool) {
	t_.ID.Send(objc.RegisterName("setAllowsColorSelection:"), value)
}

// The default color for the tool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/defaultColor
func (t_ ToolPickerCustomItemConfiguration) DefaultColor() appkit.Color {
	rv := objc.Send[appkit.Color](t_.ID, objc.Sel("defaultColor"))
	return rv
}
func (t_ ToolPickerCustomItemConfiguration) SetDefaultColor(value appkit.Color) {
	t_.ID.Send(objc.RegisterName("setDefaultColor:"), value)
}

// The default width for the tool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/defaultWidth
func (t_ ToolPickerCustomItemConfiguration) DefaultWidth() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("defaultWidth"))
	return rv
}
func (t_ ToolPickerCustomItemConfiguration) SetDefaultWidth(value float64) {
	t_.ID.Send(objc.RegisterName("setDefaultWidth:"), value)
}

// A string that uniquely identifies the tool in the picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/identifier
func (t_ ToolPickerCustomItemConfiguration) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("identifier"))
	return rv
}
func (t_ ToolPickerCustomItemConfiguration) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	t_.ID.Send(objc.RegisterName("setIdentifier:"), value)
}

// A closure that provides an image for the tool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/imageProvider
func (t_ ToolPickerCustomItemConfiguration) ImageProvider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("imageProvider"))
	return rv
}
func (t_ ToolPickerCustomItemConfiguration) SetImageProvider(value unsafe.Pointer) {
	t_.ID.Send(objc.RegisterName("setImageProvider:"), value)
}

// A short string to show as the name of the tool in the UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/name
func (t_ ToolPickerCustomItemConfiguration) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("name"))
	return rv
}
func (t_ ToolPickerCustomItemConfiguration) SetName(value objc.IObject /* cross-framework: NSString */) {
	t_.ID.Send(objc.RegisterName("setName:"), value)
}

// Defines which attribute controls are available to be presented in UI such as the tool attributes popover, or inline in the picker presented from a pencil squeeze. Controls for properties which the tool item does not support will not be presented. Excluding a control here does not hide all UI for adjusting that value. For example, excluding the opacity control here will not remove it from the color picker, if the color picker is otherwise available. Defaults to all controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/toolAttributeControls
func (t_ ToolPickerCustomItemConfiguration) ToolAttributeControls() ToolPickerCustomItemControlOptions {
	rv := objc.Send[ToolPickerCustomItemControlOptions](t_.ID, objc.Sel("toolAttributeControls"))
	return rv
}
func (t_ ToolPickerCustomItemConfiguration) SetToolAttributeControls(value ToolPickerCustomItemControlOptions) {
	t_.ID.Send(objc.RegisterName("setToolAttributeControls:"), value)
}

// A closure to provide a view controller above the system controls in the tool attributes popover.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/viewControllerProvider
func (t_ ToolPickerCustomItemConfiguration) ViewControllerProvider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("viewControllerProvider"))
	return rv
}
func (t_ ToolPickerCustomItemConfiguration) SetViewControllerProvider(value unsafe.Pointer) {
	t_.ID.Send(objc.RegisterName("setViewControllerProvider:"), value)
}

// A dictionary with UI options for selecting width, with each element containing a width value and its corresponding image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/widthVariants
func (t_ ToolPickerCustomItemConfiguration) WidthVariants() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("widthVariants"))
	return rv
}
func (t_ ToolPickerCustomItemConfiguration) SetWidthVariants(value foundation.IDictionary) {
	t_.ID.Send(objc.RegisterName("setWidthVariants:"), value)
}




