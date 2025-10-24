// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	Color() objc.IObject /* cross-framework: Color */
	SetColor(value objc.IObject /* cross-framework: Color */)
	Width() float64
	SetWidth(value float64)
	// methods:
}

// A configuration that specifies the appearance and behavior of a custom tool item and its contents.


// A configuration that specifies the appearance and behavior of a custom tool item and its contents.
//
// [Full Topic]
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



// The current color of the custom tool item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pktoolpickercustomitem/color
func (t_ ToolPickerCustomItemConfiguration) Color() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[appkit.Color](t_.ID, objc.Sel("color"))
	return rv
}


// The current color of the custom tool item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pktoolpickercustomitem/color
func (t_ ToolPickerCustomItemConfiguration) SetColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColor:"), value)
}


// The current width of the custom tool item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pktoolpickercustomitem/width
func (t_ ToolPickerCustomItemConfiguration) Width() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("width"))
	return rv
}


// The current width of the custom tool item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pktoolpickercustomitem/width
func (t_ ToolPickerCustomItemConfiguration) SetWidth(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWidth:"), value)
}



