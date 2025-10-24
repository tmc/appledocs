// Code generated from Apple documentation for ExecutionPolicy. DO NOT EDIT.

package executionpolicy

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class EPDeveloperTool */


/* debug [class_header]: Header for EPDeveloperTool */
// The class instance for the [EPDeveloperTool] class.
var (
	EPDeveloperToolClass     _EPDeveloperToolClass
	EPDeveloperToolClassOnce sync.Once
)

func getEPDeveloperToolClass() _EPDeveloperToolClass {
	EPDeveloperToolClassOnce.Do(func() {
		EPDeveloperToolClass = _EPDeveloperToolClass{objc.GetClass("EPDeveloperTool")}
	})
	return EPDeveloperToolClass
}

type _EPDeveloperToolClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EPDeveloperTool */
// An interface definition for the [EPDeveloperTool] class.
type IEPDeveloperTool interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for EPDeveloperTool */
	// properties:
	AuthorizationStatus() EPDeveloperToolStatus
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EPDeveloperTool */
	// methods:
	RequestDeveloperToolAccessWithCompletionHandler(handler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EPDeveloperTool */
// Alloc allocates a new instance without initialization.
func (ec _EPDeveloperToolClass) Alloc() EPDeveloperTool {
	rv := objc.Send[EPDeveloperTool](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EPDeveloperToolClass) New() EPDeveloperTool {
	rv := objc.Send[EPDeveloperTool](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EPDeveloperTool) Init() EPDeveloperTool {
	rv := objc.Send[EPDeveloperTool](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EPDeveloperTool) Autorelease() EPDeveloperTool {
	rv := objc.Send[EPDeveloperTool](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEPDeveloperTool creates a new EPDeveloperTool instance.
func NewEPDeveloperTool() EPDeveloperTool {
	return getEPDeveloperToolClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EPDeveloperTool */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperTool
type EPDeveloperTool struct {
	objectivec.Object
}

// EPDeveloperToolFrom constructs a [EPDeveloperTool] from an unsafe.Pointer.
func EPDeveloperToolFrom(ptr unsafe.Pointer) EPDeveloperTool {
	return EPDeveloperTool{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EPDeveloperTool */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EPDeveloperTool */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EPDeveloperTool */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EPDeveloperTool */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperTool/requestAccess(completionHandler:)
func (e_ EPDeveloperTool) RequestDeveloperToolAccessWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("requestDeveloperToolAccessWithCompletionHandler:"), handler)
}/* debug [instance_methods/method]: RequestDeveloperToolAccessWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EPDeveloperTool */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperTool/authorizationStatus
func (e_ EPDeveloperTool) AuthorizationStatus() EPDeveloperToolStatus {
	rv := objc.Send[EPDeveloperToolStatus](e_.ID, objc.Sel("authorizationStatus"))
	return rv
}/* debug [instance_properties/getter]: authorizationStatus */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EPDeveloperTool */


