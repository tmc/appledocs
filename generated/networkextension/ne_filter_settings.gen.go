// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEFilterSettings */


/* debug [class_header]: Header for NEFilterSettings */
// The class instance for the [NEFilterSettings] class.
var (
	NEFilterSettingsClass     _NEFilterSettingsClass
	NEFilterSettingsClassOnce sync.Once
)

func getNEFilterSettingsClass() _NEFilterSettingsClass {
	NEFilterSettingsClassOnce.Do(func() {
		NEFilterSettingsClass = _NEFilterSettingsClass{objc.GetClass("NEFilterSettings")}
	})
	return NEFilterSettingsClass
}

type _NEFilterSettingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFilterSettings */
// An interface definition for the [NEFilterSettings] class.
type INEFilterSettings interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEFilterSettings */
	// properties:
	DefaultAction() NEFilterAction
	Rules() []NEFilterRule
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFilterSettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFilterSettings */
// Alloc allocates a new instance without initialization.
func (nc _NEFilterSettingsClass) Alloc() NEFilterSettings {
	rv := objc.Send[NEFilterSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterSettingsClass) New() NEFilterSettings {
	rv := objc.Send[NEFilterSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterSettings) Init() NEFilterSettings {
	rv := objc.Send[NEFilterSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterSettings) Autorelease() NEFilterSettings {
	rv := objc.Send[NEFilterSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterSettings creates a new NEFilterSettings instance.
func NewNEFilterSettings() NEFilterSettings {
	return getNEFilterSettingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFilterSettings */
// The rules and other settings that define the operation of a filter.
//
// instances use to communicate the desired settings for the filter to the framework. The framework takes care of applying the contained settings to the system.


// The rules and other settings that define the operation of a filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSettings
type NEFilterSettings struct {
	objectivec.Object
}

// NEFilterSettingsFrom constructs a [NEFilterSettings] from an unsafe.Pointer.
//
// The rules and other settings that define the operation of a filter.
func NEFilterSettingsFrom(ptr unsafe.Pointer) NEFilterSettings {
	return NEFilterSettings{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFilterSettings */

// Creates a new settings instance from an array of rules and a default action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSettings/init(rules:defaultAction:)
func NewNEFilterSettingsWithRulesDefaultAction(rules []NEFilterRule, defaultAction NEFilterAction) NEFilterSettings {
	instance := getNEFilterSettingsClass().Alloc()
	rv := objc.Send[NEFilterSettings](instance.ID, objc.Sel("initWithRules:defaultAction:"), rules, defaultAction)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEFilterSettingsWithRulesDefaultAction */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFilterSettings */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFilterSettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFilterSettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFilterSettings */

// The default action to take for flows of network data that don’t match any of the specified rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSettings/defaultAction
func (n_ NEFilterSettings) DefaultAction() NEFilterAction {
	rv := objc.Send[NEFilterAction](n_.ID, objc.Sel("defaultAction"))
	return rv
}/* debug [instance_properties/getter]: defaultAction */


// An ordered list of rules that define the filter’s operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSettings/rules
func (n_ NEFilterSettings) Rules() []NEFilterRule {
	rv := objc.Send[[]NEFilterRule](n_.ID, objc.Sel("rules"))
	return rv
}/* debug [instance_properties/getter]: rules */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFilterSettings */


