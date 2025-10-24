// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AMWorkspace */


/* debug [class_header]: Header for AMWorkspace */
// The class instance for the [AMWorkspace] class.
var (
	AMWorkspaceClass     _AMWorkspaceClass
	AMWorkspaceClassOnce sync.Once
)

func getAMWorkspaceClass() _AMWorkspaceClass {
	AMWorkspaceClassOnce.Do(func() {
		AMWorkspaceClass = _AMWorkspaceClass{objc.GetClass("AMWorkspace")}
	})
	return AMWorkspaceClass
}

type _AMWorkspaceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AMWorkspace */
// An interface definition for the [AMWorkspace] class.
type IAMWorkspace interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AMWorkspace */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AMWorkspace */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AMWorkspace */
// Alloc allocates a new instance without initialization.
func (ac _AMWorkspaceClass) Alloc() AMWorkspace {
	rv := objc.Send[AMWorkspace](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AMWorkspaceClass) New() AMWorkspace {
	rv := objc.Send[AMWorkspace](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AMWorkspace) Init() AMWorkspace {
	rv := objc.Send[AMWorkspace](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AMWorkspace) Autorelease() AMWorkspace {
	rv := objc.Send[AMWorkspace](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMWorkspace creates a new AMWorkspace instance.
func NewAMWorkspace() AMWorkspace {
	return getAMWorkspaceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AMWorkspace */
// A workspace for running an Automator workflow.
//
// The class provides access to the shared workspace in the Automator framework, where you can run workflows without a workflow controller. Use to access the shared workspace and to run your workflow in it.


// A workspace for running an Automator workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkspace
type AMWorkspace struct {
	objectivec.Object
}

// AMWorkspaceFrom constructs a [AMWorkspace] from an unsafe.Pointer.
//
// A workspace for running an Automator workflow.
func AMWorkspaceFrom(ptr unsafe.Pointer) AMWorkspace {
	return AMWorkspace{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AMWorkspace *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AMWorkspace */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AMWorkspace */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AMWorkspace */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AMWorkspace */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AMWorkspace */





