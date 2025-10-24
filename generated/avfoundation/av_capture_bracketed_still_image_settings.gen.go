// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureBracketedStillImageSettings */


/* debug [class_header]: Header for AVCaptureBracketedStillImageSettings */
// The class instance for the [CaptureBracketedStillImageSettings] class.
var (
	CaptureBracketedStillImageSettingsClass     _CaptureBracketedStillImageSettingsClass
	CaptureBracketedStillImageSettingsClassOnce sync.Once
)

func getCaptureBracketedStillImageSettingsClass() _CaptureBracketedStillImageSettingsClass {
	CaptureBracketedStillImageSettingsClassOnce.Do(func() {
		CaptureBracketedStillImageSettingsClass = _CaptureBracketedStillImageSettingsClass{objc.GetClass("AVCaptureBracketedStillImageSettings")}
	})
	return CaptureBracketedStillImageSettingsClass
}

type _CaptureBracketedStillImageSettingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureBracketedStillImageSettings */
// An interface definition for the [CaptureBracketedStillImageSettings] class.
type ICaptureBracketedStillImageSettings interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureBracketedStillImageSettings */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureBracketedStillImageSettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureBracketedStillImageSettings */
// Alloc allocates a new instance without initialization.
func (cc _CaptureBracketedStillImageSettingsClass) Alloc() CaptureBracketedStillImageSettings {
	rv := objc.Send[CaptureBracketedStillImageSettings](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureBracketedStillImageSettingsClass) New() CaptureBracketedStillImageSettings {
	rv := objc.Send[CaptureBracketedStillImageSettings](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureBracketedStillImageSettings) Init() CaptureBracketedStillImageSettings {
	rv := objc.Send[CaptureBracketedStillImageSettings](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureBracketedStillImageSettings) Autorelease() CaptureBracketedStillImageSettings {
	rv := objc.Send[CaptureBracketedStillImageSettings](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureBracketedStillImageSettings creates a new CaptureBracketedStillImageSettings instance.
func NewCaptureBracketedStillImageSettings() CaptureBracketedStillImageSettings {
	return getCaptureBracketedStillImageSettingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureBracketedStillImageSettings */
// The abstract superclass for bracketed photo capture settings.


// The abstract superclass for bracketed photo capture settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureBracketedStillImageSettings
type CaptureBracketedStillImageSettings struct {
	objectivec.Object
}

// CaptureBracketedStillImageSettingsFrom constructs a [CaptureBracketedStillImageSettings] from an unsafe.Pointer.
//
// The abstract superclass for bracketed photo capture settings.
func CaptureBracketedStillImageSettingsFrom(ptr unsafe.Pointer) CaptureBracketedStillImageSettings {
	return CaptureBracketedStillImageSettings{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureBracketedStillImageSettings *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureBracketedStillImageSettings */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureBracketedStillImageSettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureBracketedStillImageSettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureBracketedStillImageSettings */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureBracketedStillImageSettings */



