// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ToolPicker] class.
var (
	ToolPickerClass     _ToolPickerClass
	ToolPickerClassOnce sync.Once
)

func getToolPickerClass() _ToolPickerClass {
	ToolPickerClassOnce.Do(func() {
		ToolPickerClass = _ToolPickerClass{objc.GetClass("PKToolPicker")}
	})
	return ToolPickerClass
}

type _ToolPickerClass struct {
	class objc.Class
}

// An interface definition for the [ToolPicker] class.
type IToolPicker interface {
	objectivec.IObject
	AddObserver(observer objectivec.IObject)
	FrameObscuredInView(view appkit.IView) coregraphics.CGRect
	RemoveObserver(observer objectivec.IObject)
	SetVisibleForFirstResponder(visible bool, responder appkit.IResponder)
	AccessoryItem() unsafe.Pointer
	SetAccessoryItem(value unsafe.Pointer)
	ColorMaximumLinearExposure() float64
	SetColorMaximumLinearExposure(value float64)
	ColorUserInterfaceStyle() unsafe.Pointer
	SetColorUserInterfaceStyle(value unsafe.Pointer)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	RulerActive() bool
	SetRulerActive(value bool)
	IsVisible() bool
	MaximumSupportedContentVersion() ContentVersion
	SetMaximumSupportedContentVersion(value IContentVersion)
	OverrideUserInterfaceStyle() unsafe.Pointer
	SetOverrideUserInterfaceStyle(value unsafe.Pointer)
	PrefersDismissControlVisible() bool
	SetPrefersDismissControlVisible(value bool)
	SelectedTool() PKTool
	SetSelectedTool(value IPKTool)
	SelectedToolItem() PKToolPickerItem
	SetSelectedToolItem(value IPKToolPickerItem)
	SelectedToolItemIdentifier() string
	SetSelectedToolItemIdentifier(value string)
	ShowsDrawingPolicyControls() bool
	SetShowsDrawingPolicyControls(value bool)
	StateAutosaveName() string
	SetStateAutosaveName(value string)
	ToolItems() []ToolPickerItem
	IsRulerActive() bool
	SetIsRulerActive(value bool)
}

// A tool palette that displays a selection of drawing tools and colors for tools that a person can choose from.
//
// A manages a draggable palette that displays drawing tools, colors, and additional options. You add a tool picker to your interface and configure it to display its palette at appropriate times. While the palette is onscreen, a person may reposition it anywhere within the current window. When a person interacts with the palette, the tool picker notifies registered observers of the changes so that they can respond. When configuring your interface, call the method to associate the tool picker with one or more views in your interface. Each window manages its own tool picker, and the window’s first responder determines the visibility of that tool picker. When one of the registered objects becomes first responder, the tool picker automatically adds its palette view to the current window. When there isn’t a registered object as first responder, the tool picker hides its palette view. implements the observer protocol for detecting tool picker changes. Adding your canvas view as an observer to a tool picker automatically updates the current drawing tools. For more information about implementing custom observer objects, see .
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker
type ToolPicker struct {
	objectivec.Object
}

// ToolPickerFrom constructs a [ToolPicker] from an unsafe.Pointer.
//
// A tool palette that displays a selection of drawing tools and colors for tools that a person can choose from.
func ToolPickerFrom(ptr unsafe.Pointer) ToolPicker {
	return ToolPicker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _ToolPickerClass) Alloc() ToolPicker {
	rv := objc.Send[ToolPicker](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ToolPickerClass) New() ToolPicker {
	rv := objc.Send[ToolPicker](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ToolPicker) Init() ToolPicker {
	rv := objc.Send[ToolPicker](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ToolPicker) Autorelease() ToolPicker {
	rv := objc.Send[ToolPicker](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewToolPicker creates a new ToolPicker instance.
func NewToolPicker() ToolPicker {
	return getToolPickerClass().New()
}




// Creates a new tool picker with the tools you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/init(toolItems:)
func NewToolPickerWithToolItems(items []ToolPickerItem) ToolPicker {
	instance := getToolPickerClass().Alloc()
	rv := objc.Send[ToolPicker](instance.ID, objc.Sel("initWithToolItems:"), items)
	rv.Autorelease()
	return rv
}


// Returns the tool picker object to use for the specified window.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/shared(for:)
func (tc _ToolPickerClass) SharedToolPickerForWindow(window appkit.IWindow) ToolPicker {
	rv := objc.Send[ToolPicker](objc.ID(tc.class), objc.Sel("sharedToolPickerForWindow:"), window)
	return rv
}

// The default tool items for new tool pickers.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/defaultToolItems
func (tc _ToolPickerClass) DefaultToolItems() []ToolPickerItem {
	rv := objc.Send[[]ToolPickerItem](objc.ID(tc.class), objc.Sel("defaultToolItems"))
	return rv
}
// Adds the specified object to the list of objects to notify when the picker configuration changes.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/addObserver(_:)
func (t_ ToolPicker) AddObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addObserver:"), observer)
}

// Returns the portion of the specified view that the tool picker obscures.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/frameObscured(in:)
func (t_ ToolPicker) FrameObscuredInView(view appkit.IView) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("frameObscuredInView:"), view)
	return rv
}

// Removes the specified object from the list of objects to notify when the picker configuration changes.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/removeObserver(_:)
func (t_ ToolPicker) RemoveObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeObserver:"), observer)
}

// Sets the visibility for the tool picker, based on when the specified responder object becomes active.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/setVisible(_:forFirstResponder:)
func (t_ ToolPicker) SetVisibleForFirstResponder(visible bool, responder appkit.IResponder) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisible:forFirstResponder:"), visible, responder)
}

// An optional button that appears at the trailing edge of the tool picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/accessoryItem
func (t_ ToolPicker) AccessoryItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("accessoryItem"))
	return rv
}


// SetAccessoryItem sets the value of the accessoryItem property.
// An optional button that appears at the trailing edge of the tool picker.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/accessoryItem
func (t_ ToolPicker) SetAccessoryItem(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAccessoryItem:"), value)
}

// Maximum linear exposure for the color picker used by the tool picker. Can be used to enable picking HDR colors.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/colorMaximumLinearExposure
func (t_ ToolPicker) ColorMaximumLinearExposure() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("colorMaximumLinearExposure"))
	return rv
}


// SetColorMaximumLinearExposure sets the value of the colorMaximumLinearExposure property.
// Maximum linear exposure for the color picker used by the tool picker. Can be used to enable picking HDR colors.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/colorMaximumLinearExposure
func (t_ ToolPicker) SetColorMaximumLinearExposure(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColorMaximumLinearExposure:"), value)
}

// The user interface style for the tool picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/colorUserInterfaceStyle
func (t_ ToolPicker) ColorUserInterfaceStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("colorUserInterfaceStyle"))
	return rv
}


// SetColorUserInterfaceStyle sets the value of the colorUserInterfaceStyle property.
// The user interface style for the tool picker.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/colorUserInterfaceStyle
func (t_ ToolPicker) SetColorUserInterfaceStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColorUserInterfaceStyle:"), value)
}

// The default tool items for new tool pickers.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/defaultToolItems
func (t_ ToolPicker) DefaultToolItems() []ToolPickerItem {
	rv := objc.Send[[]ToolPickerItem](t_.ID, objc.Sel("defaultToolItems"))
	return rv
}

// The delegate for the tool picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/delegate-swift.property
func (t_ ToolPicker) Delegate() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate for the tool picker.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/delegate-swift.property
func (t_ ToolPicker) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the ruler is visible on the canvas.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/isRulerActive
func (t_ ToolPicker) RulerActive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rulerActive"))
	return rv
}


// SetRulerActive sets the value of the rulerActive property.
// A Boolean value that indicates whether the ruler is visible on the canvas.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/isRulerActive
func (t_ ToolPicker) SetRulerActive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRulerActive:"), value)
}

// A Boolean value that indicates whether the tool picker is currently visible.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/isVisible
func (t_ ToolPicker) IsVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isVisible"))
	return rv
}

// The maximum version of PencilKit to support.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/maximumSupportedContentVersion
func (t_ ToolPicker) MaximumSupportedContentVersion() ContentVersion {
	rv := objc.Send[ContentVersion](t_.ID, objc.Sel("maximumSupportedContentVersion"))
	return rv
}


// SetMaximumSupportedContentVersion sets the value of the maximumSupportedContentVersion property.
// The maximum version of PencilKit to support.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/maximumSupportedContentVersion
func (t_ ToolPicker) SetMaximumSupportedContentVersion(value IContentVersion) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaximumSupportedContentVersion:"), value)
}

// The specific user interface style to apply to the tool picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/overrideUserInterfaceStyle
func (t_ ToolPicker) OverrideUserInterfaceStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("overrideUserInterfaceStyle"))
	return rv
}


// SetOverrideUserInterfaceStyle sets the value of the overrideUserInterfaceStyle property.
// The specific user interface style to apply to the tool picker.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/overrideUserInterfaceStyle
func (t_ ToolPicker) SetOverrideUserInterfaceStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setOverrideUserInterfaceStyle:"), value)
}

// If this is true the tool picker may show UI that allows dismissing it. If this is false the tool picker will not show this UI. By default this resigns first responder, but is customizable by ’s method.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/prefersDismissControlVisible
func (t_ ToolPicker) PrefersDismissControlVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("prefersDismissControlVisible"))
	return rv
}


// SetPrefersDismissControlVisible sets the value of the prefersDismissControlVisible property.
// If this is true the tool picker may show UI that allows dismissing it. If this is false the tool picker will not show this UI. By default this resigns first responder, but is customizable by ’s method.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/prefersDismissControlVisible
func (t_ ToolPicker) SetPrefersDismissControlVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPrefersDismissControlVisible:"), value)
}

// The currently selected tool in the tool picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/selectedTool-93ikc
func (t_ ToolPicker) SelectedTool() PKTool {
	rv := objc.Send[PKTool](t_.ID, objc.Sel("selectedTool"))
	return rv
}


// SetSelectedTool sets the value of the selectedTool property.
// The currently selected tool in the tool picker.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/selectedTool-93ikc
func (t_ ToolPicker) SetSelectedTool(value IPKTool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedTool:"), value)
}

// The currently selected tool item in the tool picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/selectedToolItem
func (t_ ToolPicker) SelectedToolItem() PKToolPickerItem {
	rv := objc.Send[PKToolPickerItem](t_.ID, objc.Sel("selectedToolItem"))
	return rv
}


// SetSelectedToolItem sets the value of the selectedToolItem property.
// The currently selected tool item in the tool picker.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/selectedToolItem
func (t_ ToolPicker) SetSelectedToolItem(value IPKToolPickerItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedToolItem:"), value)
}

// The identifier of the selected tool item in the tool picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/selectedToolItemIdentifier
func (t_ ToolPicker) SelectedToolItemIdentifier() string {
	rv := objc.Send[string](t_.ID, objc.Sel("selectedToolItemIdentifier"))
	return rv
}


// SetSelectedToolItemIdentifier sets the value of the selectedToolItemIdentifier property.
// The identifier of the selected tool item in the tool picker.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/selectedToolItemIdentifier
func (t_ ToolPicker) SetSelectedToolItemIdentifier(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedToolItemIdentifier:"), objc.String(value))
}

// A Boolean value that indicates whether the default drawing policy UI is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/showsDrawingPolicyControls
func (t_ ToolPicker) ShowsDrawingPolicyControls() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("showsDrawingPolicyControls"))
	return rv
}


// SetShowsDrawingPolicyControls sets the value of the showsDrawingPolicyControls property.
// A Boolean value that indicates whether the default drawing policy UI is visible.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/showsDrawingPolicyControls
func (t_ ToolPicker) SetShowsDrawingPolicyControls(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setShowsDrawingPolicyControls:"), value)
}

// The name used to automatically save the tool picker’s state in the defaults system.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/stateAutosaveName
func (t_ ToolPicker) StateAutosaveName() string {
	rv := objc.Send[string](t_.ID, objc.Sel("stateAutosaveName"))
	return rv
}


// SetStateAutosaveName sets the value of the stateAutosaveName property.
// The name used to automatically save the tool picker’s state in the defaults system.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/stateAutosaveName
func (t_ ToolPicker) SetStateAutosaveName(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStateAutosaveName:"), objc.String(value))
}

// All tool items in the tool picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/toolItems
func (t_ ToolPicker) ToolItems() []ToolPickerItem {
	rv := objc.Send[[]ToolPickerItem](t_.ID, objc.Sel("toolItems"))
	return rv
}

// A Boolean value that indicates whether the ruler is visible on the canvas.
//
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pktoolpicker/isruleractive
func (t_ ToolPicker) IsRulerActive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRulerActive"))
	return rv
}


// SetIsRulerActive sets the value of the isRulerActive property.
// A Boolean value that indicates whether the ruler is visible on the canvas.

//
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pktoolpicker/isruleractive
func (t_ ToolPicker) SetIsRulerActive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRulerActive:"), value)
}


