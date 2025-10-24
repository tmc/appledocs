// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UIWindowSceneActivationConfiguration */


/* debug [class_header]: Header for UIWindowSceneActivationConfiguration */
// The class instance for the [WindowSceneActivationConfiguration] class.
var (
	WindowSceneActivationConfigurationClass     _WindowSceneActivationConfigurationClass
	WindowSceneActivationConfigurationClassOnce sync.Once
)

func getWindowSceneActivationConfigurationClass() _WindowSceneActivationConfigurationClass {
	WindowSceneActivationConfigurationClassOnce.Do(func() {
		WindowSceneActivationConfigurationClass = _WindowSceneActivationConfigurationClass{objc.GetClass("UIWindowSceneActivationConfiguration")}
	})
	return WindowSceneActivationConfigurationClass
}

type _WindowSceneActivationConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WindowSceneActivationConfiguration */
// An interface definition for the [WindowSceneActivationConfiguration] class.
type IWindowSceneActivationConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WindowSceneActivationConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WindowSceneActivationConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WindowSceneActivationConfiguration */
// Alloc allocates a new instance without initialization.
func (wc _WindowSceneActivationConfigurationClass) Alloc() WindowSceneActivationConfiguration {
	rv := objc.Send[WindowSceneActivationConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WindowSceneActivationConfigurationClass) New() WindowSceneActivationConfiguration {
	rv := objc.Send[WindowSceneActivationConfiguration](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WindowSceneActivationConfiguration) Init() WindowSceneActivationConfiguration {
	rv := objc.Send[WindowSceneActivationConfiguration](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WindowSceneActivationConfiguration) Autorelease() WindowSceneActivationConfiguration {
	rv := objc.Send[WindowSceneActivationConfiguration](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWindowSceneActivationConfiguration creates a new WindowSceneActivationConfiguration instance.
func NewWindowSceneActivationConfiguration() WindowSceneActivationConfiguration {
	return getWindowSceneActivationConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WindowSceneActivationConfiguration */
// A parent class referenced by other QuickLook classes.


// A parent class referenced by other QuickLook classes. [Full Topic]
type WindowSceneActivationConfiguration struct {
	objectivec.Object
}

// WindowSceneActivationConfigurationFrom constructs a [WindowSceneActivationConfiguration] from an unsafe.Pointer.
//
// A parent class referenced by other QuickLook classes.
func WindowSceneActivationConfigurationFrom(ptr unsafe.Pointer) WindowSceneActivationConfiguration {
	return WindowSceneActivationConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WindowSceneActivationConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WindowSceneActivationConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WindowSceneActivationConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WindowSceneActivationConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WindowSceneActivationConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class UIWindowSceneActivationConfiguration */



