//go:build darwin && ios

// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/mapkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ToolPicker


// Adds the specified object to the list of objects to notify when the picker configuration changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/addObserver(_:)
func (t_ ToolPicker) AddObserver(observer unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addObserver:"), observer)
}

// Returns the portion of the specified view that the tool picker obscures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/frameObscured(in:)
func (t_ ToolPicker) FrameObscuredInView(view appkit.View) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("frameObscuredInView:"), view)
	return rv
}

// Removes the specified object from the list of objects to notify when the picker configuration changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/removeObserver(_:)
func (t_ ToolPicker) RemoveObserver(observer unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeObserver:"), observer)
}

// Sets the visibility for the tool picker, based on when the specified responder object becomes active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/setVisible(_:forFirstResponder:)
func (t_ ToolPicker) SetVisibleForFirstResponder(visible bool, responder appkit.Responder) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisible:forFirstResponder:"), visible, responder)
}

// iOS-only properties

// An optional button that appears at the trailing edge of the tool picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/accessoryItem
func (t_ ToolPicker) AccessoryItem() mapkit.BarButtonItem {
	rv := objc.Send[mapkit.BarButtonItem](t_.ID, objc.Sel("accessoryItem"))
	return rv
}
func (t_ ToolPicker) SetAccessoryItem(value mapkit.BarButtonItem) {
	t_.ID.Send(objc.RegisterName("setAccessoryItem:"), value)
}

// Maximum linear exposure for the color picker used by the tool picker. Can be used to enable picking HDR colors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/colorMaximumLinearExposure
func (t_ ToolPicker) ColorMaximumLinearExposure() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("colorMaximumLinearExposure"))
	return rv
}
func (t_ ToolPicker) SetColorMaximumLinearExposure(value float64) {
	t_.ID.Send(objc.RegisterName("setColorMaximumLinearExposure:"), value)
}

// The user interface style for the tool picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/colorUserInterfaceStyle
func (t_ ToolPicker) ColorUserInterfaceStyle() UserInterfaceStyle /* not a class type */ {
	rv := objc.Send[UserInterfaceStyle](t_.ID, objc.Sel("colorUserInterfaceStyle"))
	return rv
}
func (t_ ToolPicker) SetColorUserInterfaceStyle(value UserInterfaceStyle /* not a class type */) {
	t_.ID.Send(objc.RegisterName("setColorUserInterfaceStyle:"), value)
}

// The delegate for the tool picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/delegate-swift.property
func (t_ ToolPicker) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}
func (t_ ToolPicker) SetDelegate(value unsafe.Pointer) {
	t_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// A Boolean value that indicates whether the ruler is visible on the canvas.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/isRulerActive
func (t_ ToolPicker) RulerActive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rulerActive"))
	return rv
}
func (t_ ToolPicker) SetRulerActive(value bool) {
	t_.ID.Send(objc.RegisterName("setRulerActive:"), value)
}

// A Boolean value that indicates whether the tool picker is currently visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/isVisible
func (t_ ToolPicker) IsVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isVisible"))
	return rv
}

// The maximum version of PencilKit to support.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/maximumSupportedContentVersion
func (t_ ToolPicker) MaximumSupportedContentVersion() ContentVersion {
	rv := objc.Send[ContentVersion](t_.ID, objc.Sel("maximumSupportedContentVersion"))
	return rv
}
func (t_ ToolPicker) SetMaximumSupportedContentVersion(value ContentVersion) {
	t_.ID.Send(objc.RegisterName("setMaximumSupportedContentVersion:"), value)
}

// The specific user interface style to apply to the tool picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/overrideUserInterfaceStyle
func (t_ ToolPicker) OverrideUserInterfaceStyle() UserInterfaceStyle /* not a class type */ {
	rv := objc.Send[UserInterfaceStyle](t_.ID, objc.Sel("overrideUserInterfaceStyle"))
	return rv
}
func (t_ ToolPicker) SetOverrideUserInterfaceStyle(value UserInterfaceStyle /* not a class type */) {
	t_.ID.Send(objc.RegisterName("setOverrideUserInterfaceStyle:"), value)
}

// If this is true the tool picker may show UI that allows dismissing it. If this is false the tool picker will not show this UI. By default this resigns first responder, but is customizable by ’s method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/prefersDismissControlVisible
func (t_ ToolPicker) PrefersDismissControlVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("prefersDismissControlVisible"))
	return rv
}
func (t_ ToolPicker) SetPrefersDismissControlVisible(value bool) {
	t_.ID.Send(objc.RegisterName("setPrefersDismissControlVisible:"), value)
}

// The currently selected tool in the tool picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/selectedTool-93ikc
func (t_ ToolPicker) SelectedTool() IPKTool {
	rv := objc.Send[Tool](t_.ID, objc.Sel("selectedTool"))
	return rv
}
func (t_ ToolPicker) SetSelectedTool(value IPKTool) {
	t_.ID.Send(objc.RegisterName("setSelectedTool:"), value)
}

// The currently selected tool item in the tool picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/selectedToolItem
func (t_ ToolPicker) SelectedToolItem() IPKToolPickerItem {
	rv := objc.Send[ToolPickerItem](t_.ID, objc.Sel("selectedToolItem"))
	return rv
}
func (t_ ToolPicker) SetSelectedToolItem(value IPKToolPickerItem) {
	t_.ID.Send(objc.RegisterName("setSelectedToolItem:"), value)
}

// The identifier of the selected tool item in the tool picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/selectedToolItemIdentifier
func (t_ ToolPicker) SelectedToolItemIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("selectedToolItemIdentifier"))
	return rv
}
func (t_ ToolPicker) SetSelectedToolItemIdentifier(value objc.IObject /* cross-framework: NSString */) {
	t_.ID.Send(objc.RegisterName("setSelectedToolItemIdentifier:"), value)
}

// A Boolean value that indicates whether the default drawing policy UI is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/showsDrawingPolicyControls
func (t_ ToolPicker) ShowsDrawingPolicyControls() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("showsDrawingPolicyControls"))
	return rv
}
func (t_ ToolPicker) SetShowsDrawingPolicyControls(value bool) {
	t_.ID.Send(objc.RegisterName("setShowsDrawingPolicyControls:"), value)
}

// The name used to automatically save the tool picker’s state in the defaults system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/stateAutosaveName
func (t_ ToolPicker) StateAutosaveName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("stateAutosaveName"))
	return rv
}
func (t_ ToolPicker) SetStateAutosaveName(value objc.IObject /* cross-framework: NSString */) {
	t_.ID.Send(objc.RegisterName("setStateAutosaveName:"), value)
}

// All tool items in the tool picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/toolItems
func (t_ ToolPicker) ToolItems() []ToolPickerItem {
	rv := objc.Send[[]ToolPickerItem](t_.ID, objc.Sel("toolItems"))
	return rv
}




