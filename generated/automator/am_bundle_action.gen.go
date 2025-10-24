// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class AMBundleAction */


/* debug [class_header]: Header for AMBundleAction */
// The class instance for the [AMBundleAction] class.
var (
	AMBundleActionClass     _AMBundleActionClass
	AMBundleActionClassOnce sync.Once
)

func getAMBundleActionClass() _AMBundleActionClass {
	AMBundleActionClassOnce.Do(func() {
		AMBundleActionClass = _AMBundleActionClass{objc.GetClass("AMBundleAction")}
	})
	return AMBundleActionClass
}

type _AMBundleActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AMBundleAction */
// An interface definition for the [AMBundleAction] class.
type IAMBundleAction interface {
	IAMAction
	
/* debug [class_interface_properties]: Properties for AMBundleAction */
	// properties:
	Bundle() foundation.Bundle
	HasView() bool
	Parameters() unsafe.Pointer
	SetParameters(value unsafe.Pointer)
	View() appkit.View
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AMBundleAction */
	// methods:
	AwakeFromBundle()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AMBundleAction */
// Alloc allocates a new instance without initialization.
func (ac _AMBundleActionClass) Alloc() AMBundleAction {
	rv := objc.Send[AMBundleAction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AMBundleActionClass) New() AMBundleAction {
	rv := objc.Send[AMBundleAction](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AMBundleAction) Init() AMBundleAction {
	rv := objc.Send[AMBundleAction](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AMBundleAction) Autorelease() AMBundleAction {
	rv := objc.Send[AMBundleAction](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMBundleAction creates a new AMBundleAction instance.
func NewAMBundleAction() AMBundleAction {
	return getAMBundleActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AMBundleAction */
// An object that represents an Automator action that’s a loadable bundle.
//
// Automator loads action bundles from standard locations in the file system: , , and . objects have several important properties: The object associated with the action’s physical bundle The action’s view, which holds its user interface A parameters dictionary that reflects the settings in the user interface When you create a Cocoa Automator Action project in Xcode, the project template includes a custom subclass of . This custom class uses the name of the project. You must provide an implementation of , which is declared by the superclass . If you add any instance variables, you must override the method and the method of to work with them.


// An object that represents an Automator action that’s a loadable bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMBundleAction
type AMBundleAction struct {
	AMAction
}

// AMBundleActionFrom constructs a [AMBundleAction] from an unsafe.Pointer.
//
// An object that represents an Automator action that’s a loadable bundle.
func AMBundleActionFrom(ptr unsafe.Pointer) AMBundleAction {
	return AMBundleAction{
		AMAction: AMActionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AMBundleAction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AMBundleAction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AMBundleAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AMBundleAction */

// Allows the action object to perform setup tasks requiring the presence of all bundle objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMBundleAction/awakeFromBundle()
func (a_ AMBundleAction) AwakeFromBundle() {
	objc.Send[objc.ID](a_.ID, objc.Sel("awakeFromBundle"))
}/* debug [instance_methods/method]: AwakeFromBundle */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AMBundleAction */

// The action’s bundle object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMBundleAction/bundle
func (a_ AMBundleAction) Bundle() foundation.Bundle {
	rv := objc.Send[foundation.Bundle](a_.ID, objc.Sel("bundle"))
	return rv
}/* debug [instance_properties/getter]: bundle */


// A Boolean value that indicates whether the action has a view associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMBundleAction/hasView
func (a_ AMBundleAction) HasView() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasView"))
	return rv
}/* debug [instance_properties/getter]: hasView */


// The action’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMBundleAction/parameters
func (a_ AMBundleAction) Parameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("parameters"))
	return rv
}/* debug [instance_properties/getter]: parameters */


// The action’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMBundleAction/parameters
func (a_ AMBundleAction) SetParameters(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setParameters:"), value)
}/* debug [instance_properties/setter]: parameters */


// The action’s view object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMBundleAction/view
func (a_ AMBundleAction) View() appkit.View {
	rv := objc.Send[appkit.View](a_.ID, objc.Sel("view"))
	return rv
}/* debug [instance_properties/getter]: view */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AMBundleAction */



