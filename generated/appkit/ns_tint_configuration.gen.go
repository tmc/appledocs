// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTintConfiguration */


/* debug [class_header]: Header for NSTintConfiguration */
// The class instance for the [TintConfiguration] class.
var (
	TintConfigurationClass     _TintConfigurationClass
	TintConfigurationClassOnce sync.Once
)

func getTintConfigurationClass() _TintConfigurationClass {
	TintConfigurationClassOnce.Do(func() {
		TintConfigurationClass = _TintConfigurationClass{objc.GetClass("NSTintConfiguration")}
	})
	return TintConfigurationClass
}

type _TintConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TintConfiguration */
// An interface definition for the [TintConfiguration] class.
type ITintConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TintConfiguration */
	// properties:
	AdaptsToUserAccentColor() bool
	BaseTintColor() IColor
	EquivalentContentTintColor() IColor
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TintConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TintConfiguration */
// Alloc allocates a new instance without initialization.
func (tc _TintConfigurationClass) Alloc() TintConfiguration {
	rv := objc.Send[TintConfiguration](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TintConfigurationClass) New() TintConfiguration {
	rv := objc.Send[TintConfiguration](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TintConfiguration) Init() TintConfiguration {
	rv := objc.Send[TintConfiguration](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TintConfiguration) Autorelease() TintConfiguration {
	rv := objc.Send[TintConfiguration](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTintConfiguration creates a new TintConfiguration instance.
func NewTintConfiguration() TintConfiguration {
	return getTintConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TintConfiguration */
// An object that gives you the ability to choose from system-provided tinting behaviors.


// An object that gives you the ability to choose from system-provided tinting behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintConfiguration
type TintConfiguration struct {
	objectivec.Object
}

// TintConfigurationFrom constructs a [TintConfiguration] from an unsafe.Pointer.
//
// An object that gives you the ability to choose from system-provided tinting behaviors.
func TintConfigurationFrom(ptr unsafe.Pointer) TintConfiguration {
	return TintConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TintConfiguration */

// Creates a new tint configuration using a specific color value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/init(fixedColor:)
func NewTintConfigurationWithFixedColor(color IColor) TintConfiguration {
	rv := objc.Send[TintConfiguration](objc.ID(getTintConfigurationClass().class), objc.Sel("tintConfigurationWithFixedColor:"), color)
	return rv
}/* debug [class_init_methods/constructor]: NewTintConfigurationWithFixedColor */


// Creates a new tint configuration for the system to use when the app’s preferred accent color is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/init(preferredColor:)
func NewTintConfigurationWithPreferredColor(color IColor) TintConfiguration {
	rv := objc.Send[TintConfiguration](objc.ID(getTintConfigurationClass().class), objc.Sel("tintConfigurationWithPreferredColor:"), color)
	return rv
}/* debug [class_init_methods/constructor]: NewTintConfigurationWithPreferredColor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TintConfiguration */

// Creates a new tint configuration using a specific color value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/init(fixedColor:)
func (tc _TintConfigurationClass) TintConfigurationWithFixedColor(color IColor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("tintConfigurationWithFixedColor:"), color)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TintConfigurationWithFixedColor) */


// Creates a new tint configuration for the system to use when the app’s preferred accent color is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/init(preferredColor:)
func (tc _TintConfigurationClass) TintConfigurationWithPreferredColor(color IColor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("tintConfigurationWithPreferredColor:"), color)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TintConfigurationWithPreferredColor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TintConfiguration */

// The system tints the content using the system default value for its context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/default
func (tc _TintConfigurationClass) DefaultTintConfiguration() TintConfiguration {
	rv := objc.Send[TintConfiguration](objc.ID(tc.class), objc.Sel("defaultTintConfiguration"))
	return rv
}/* debug [class_properties_class/property]: defaultTintConfiguration */

// The content always displays in monochrome.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/monochrome
func (tc _TintConfigurationClass) MonochromeTintConfiguration() TintConfiguration {
	rv := objc.Send[TintConfiguration](objc.ID(tc.class), objc.Sel("monochromeTintConfiguration"))
	return rv
}/* debug [class_properties_class/property]: monochromeTintConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TintConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TintConfiguration */

// A Boolean value that indicates whether the tint configuration alters its effect based on the user’s preferred accent color choice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/adaptsToUserAccentColor
func (t_ TintConfiguration) AdaptsToUserAccentColor() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("adaptsToUserAccentColor"))
	return rv
}/* debug [instance_properties/getter]: adaptsToUserAccentColor */


// The color the system supplies when you create a tint configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/baseTintColor
func (t_ TintConfiguration) BaseTintColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("baseTintColor"))
	return rv
}/* debug [instance_properties/getter]: baseTintColor */


// The system tints the content using the system default value for its context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/default
func (t_ TintConfiguration) DefaultTintConfiguration() ITintConfiguration {
	rv := objc.Send[TintConfiguration](t_.ID, objc.Sel("defaultTintConfiguration"))
	return rv
}/* debug [instance_properties/getter]: defaultTintConfiguration */


// A color object that matches the effective content tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/equivalentContentTintColor
func (t_ TintConfiguration) EquivalentContentTintColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("equivalentContentTintColor"))
	return rv
}/* debug [instance_properties/getter]: equivalentContentTintColor */


// The content always displays in monochrome.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/monochrome
func (t_ TintConfiguration) MonochromeTintConfiguration() ITintConfiguration {
	rv := objc.Send[TintConfiguration](t_.ID, objc.Sel("monochromeTintConfiguration"))
	return rv
}/* debug [instance_properties/getter]: monochromeTintConfiguration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTintConfiguration */


