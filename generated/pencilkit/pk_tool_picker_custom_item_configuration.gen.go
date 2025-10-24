// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PKToolPickerCustomItemConfiguration */


/* debug [class_header]: Header for PKToolPickerCustomItemConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ToolPickerCustomItemConfiguration */
// An interface definition for the [ToolPickerCustomItemConfiguration] class.
type IToolPickerCustomItemConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ToolPickerCustomItemConfiguration */
	// properties:
	Color() appkit.Color
	SetColor(value appkit.Color)
	Width() float64
	SetWidth(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ToolPickerCustomItemConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ToolPickerCustomItemConfiguration */
// Alloc allocates a new instance without initialization.
func (tc _ToolPickerCustomItemConfigurationClass) Alloc() ToolPickerCustomItemConfiguration {
	rv := objc.Send[ToolPickerCustomItemConfiguration](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ToolPickerCustomItemConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ToolPickerCustomItemConfiguration */

// Create a new configuration with an identifier and a name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemConfiguration/initWithIdentifier:name:
func NewToolPickerCustomItemConfigurationWithIdentifierName(identifier objc.IObject /* cross-framework: NSString */, name objc.IObject /* cross-framework: NSString */) ToolPickerCustomItemConfiguration {
	instance := getToolPickerCustomItemConfigurationClass().Alloc()
	rv := objc.Send[ToolPickerCustomItemConfiguration](instance.ID, objc.Sel("initWithIdentifier:name:"), identifier, name)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewToolPickerCustomItemConfigurationWithIdentifierName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ToolPickerCustomItemConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ToolPickerCustomItemConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ToolPickerCustomItemConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ToolPickerCustomItemConfiguration */

// The current color of the custom tool item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pktoolpickercustomitem/color
func (t_ ToolPickerCustomItemConfiguration) Color() appkit.Color {
	rv := objc.Send[appkit.Color](t_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// The current color of the custom tool item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pktoolpickercustomitem/color
func (t_ ToolPickerCustomItemConfiguration) SetColor(value appkit.Color) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// The current width of the custom tool item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pktoolpickercustomitem/width
func (t_ ToolPickerCustomItemConfiguration) Width() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// The current width of the custom tool item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pktoolpickercustomitem/width
func (t_ ToolPickerCustomItemConfiguration) SetWidth(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKToolPickerCustomItemConfiguration */


