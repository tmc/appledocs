// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ToolPickerCustomItemConfiguration] class.
var (
	ToolPickerCustomItemConfigurationClass     _ToolPickerCustomItemConfigurationClass
	ToolPickerCustomItemConfigurationClassOnce sync.Once
)

func getToolPickerCustomItemConfigurationClass() _ToolPickerCustomItemConfigurationClass {
	ToolPickerCustomItemConfigurationClassOnce.Do(func() {
		ToolPickerCustomItemConfigurationClass = _ToolPickerCustomItemConfigurationClass{objc.GetClass("PKToolPickerCustomItemConfiguration")}
	})
	return ToolPickerCustomItemConfigurationClass
}

type _ToolPickerCustomItemConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [ToolPickerCustomItemConfiguration] class.
type IToolPickerCustomItemConfiguration interface {
	objectivec.IObject
}

// A configuration that specifies the appearance and behavior of a custom tool item and its contents.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration
type ToolPickerCustomItemConfiguration struct {
	objectivec.Object
}

// ToolPickerCustomItemConfigurationFrom constructs a [ToolPickerCustomItemConfiguration] from an unsafe.Pointer.
//
// A configuration that specifies the appearance and behavior of a custom tool item and its contents.
func ToolPickerCustomItemConfigurationFrom(ptr unsafe.Pointer) ToolPickerCustomItemConfiguration {
	return ToolPickerCustomItemConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _ToolPickerCustomItemConfigurationClass) Alloc() ToolPickerCustomItemConfiguration {
	rv := objc.Send[ToolPickerCustomItemConfiguration](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ToolPickerCustomItemConfigurationClass) New() ToolPickerCustomItemConfiguration {
	rv := objc.Send[ToolPickerCustomItemConfiguration](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ToolPickerCustomItemConfiguration) Init() ToolPickerCustomItemConfiguration {
	rv := objc.Send[ToolPickerCustomItemConfiguration](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ToolPickerCustomItemConfiguration) Autorelease() ToolPickerCustomItemConfiguration {
	rv := objc.Send[ToolPickerCustomItemConfiguration](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewToolPickerCustomItemConfiguration creates a new ToolPickerCustomItemConfiguration instance.
func NewToolPickerCustomItemConfiguration() ToolPickerCustomItemConfiguration {
	return getToolPickerCustomItemConfigurationClass().New()
}


// Create a new configuration with an identifier and a name.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/initWithIdentifier:name:
func NewToolPickerCustomItemConfigurationWithIdentifierName(identifier string, name string) ToolPickerCustomItemConfiguration {
	instance := getToolPickerCustomItemConfigurationClass().Alloc()
	rv := objc.Send[ToolPickerCustomItemConfiguration](instance.ID, objc.Sel("initWithIdentifier:name:"), objc.String(identifier), objc.String(name))
	rv.Autorelease()
	return rv
}


// A Boolean value that determines whether to show the color selection UI for the tool.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/allowsColorSelection
func (t_ ToolPickerCustomItemConfiguration) AllowsColorSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsColorSelection"))
	return rv
}


// SetAllowsColorSelection sets the value of the allowsColorSelection property.
// A Boolean value that determines whether to show the color selection UI for the tool.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/allowsColorSelection
func (t_ ToolPickerCustomItemConfiguration) SetAllowsColorSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsColorSelection:"), value)
}
// The default width for the tool.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/defaultWidth
func (t_ ToolPickerCustomItemConfiguration) DefaultWidth() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("defaultWidth"))
	return rv
}


// SetDefaultWidth sets the value of the defaultWidth property.
// The default width for the tool.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/defaultWidth
func (t_ ToolPickerCustomItemConfiguration) SetDefaultWidth(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDefaultWidth:"), value)
}
// A string that uniquely identifies the tool in the picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/identifier
func (t_ ToolPickerCustomItemConfiguration) Identifier() string {
	rv := objc.Send[string](t_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// A string that uniquely identifies the tool in the picker.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/identifier
func (t_ ToolPickerCustomItemConfiguration) SetIdentifier(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}
// A closure that provides an image for the tool.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/imageProvider
func (t_ ToolPickerCustomItemConfiguration) ImageProvider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("imageProvider"))
	return rv
}


// SetImageProvider sets the value of the imageProvider property.
// A closure that provides an image for the tool.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/imageProvider
func (t_ ToolPickerCustomItemConfiguration) SetImageProvider(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImageProvider:"), value)
}
// A short string to show as the name of the tool in the UI.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/name
func (t_ ToolPickerCustomItemConfiguration) Name() string {
	rv := objc.Send[string](t_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// A short string to show as the name of the tool in the UI.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/name
func (t_ ToolPickerCustomItemConfiguration) SetName(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setName:"), objc.String(value))
}
// Defines which attribute controls are available to be presented in UI such as the tool attributes popover, or inline in the picker presented from a pencil squeeze. Controls for properties which the tool item does not support will not be presented. Excluding a control here does not hide all UI for adjusting that value. For example, excluding the opacity control here will not remove it from the color picker, if the color picker is otherwise available. Defaults to all controls.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/toolAttributeControls
func (t_ ToolPickerCustomItemConfiguration) ToolAttributeControls() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("toolAttributeControls"))
	return rv
}


// SetToolAttributeControls sets the value of the toolAttributeControls property.
// Defines which attribute controls are available to be presented in UI such as the tool attributes popover, or inline in the picker presented from a pencil squeeze. Controls for properties which the tool item does not support will not be presented. Excluding a control here does not hide all UI for adjusting that value. For example, excluding the opacity control here will not remove it from the color picker, if the color picker is otherwise available. Defaults to all controls.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/toolAttributeControls
func (t_ ToolPickerCustomItemConfiguration) SetToolAttributeControls(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setToolAttributeControls:"), value)
}
// A closure to provide a view controller above the system controls in the tool attributes popover.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/viewControllerProvider
func (t_ ToolPickerCustomItemConfiguration) ViewControllerProvider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("viewControllerProvider"))
	return rv
}


// SetViewControllerProvider sets the value of the viewControllerProvider property.
// A closure to provide a view controller above the system controls in the tool attributes popover.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/viewControllerProvider
func (t_ ToolPickerCustomItemConfiguration) SetViewControllerProvider(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setViewControllerProvider:"), value)
}
// A dictionary with UI options for selecting width, with each element containing a width value and its corresponding image.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/widthVariants
func (t_ ToolPickerCustomItemConfiguration) WidthVariants() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("widthVariants"))
	return rv
}


// SetWidthVariants sets the value of the widthVariants property.
// A dictionary with UI options for selecting width, with each element containing a width value and its corresponding image.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/widthVariants
func (t_ ToolPickerCustomItemConfiguration) SetWidthVariants(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWidthVariants:"), value)
}

