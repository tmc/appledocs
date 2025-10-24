// Code generated from Apple documentation for ExecutionPolicy. DO NOT EDIT.

package executionpolicy

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class EPExecutionPolicy */


/* debug [class_header]: Header for EPExecutionPolicy */
// The class instance for the [EPExecutionPolicy] class.
var (
	EPExecutionPolicyClass     _EPExecutionPolicyClass
	EPExecutionPolicyClassOnce sync.Once
)

func getEPExecutionPolicyClass() _EPExecutionPolicyClass {
	EPExecutionPolicyClassOnce.Do(func() {
		EPExecutionPolicyClass = _EPExecutionPolicyClass{objc.GetClass("EPExecutionPolicy")}
	})
	return EPExecutionPolicyClass
}

type _EPExecutionPolicyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EPExecutionPolicy */
// An interface definition for the [EPExecutionPolicy] class.
type IEPExecutionPolicy interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for EPExecutionPolicy */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EPExecutionPolicy */
	// methods:
	AddPolicyExceptionForURLError(url objc.IObject /* cross-framework: NSURL */, error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EPExecutionPolicy */
// Alloc allocates a new instance without initialization.
func (ec _EPExecutionPolicyClass) Alloc() EPExecutionPolicy {
	rv := objc.Send[EPExecutionPolicy](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EPExecutionPolicyClass) New() EPExecutionPolicy {
	rv := objc.Send[EPExecutionPolicy](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EPExecutionPolicy) Init() EPExecutionPolicy {
	rv := objc.Send[EPExecutionPolicy](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EPExecutionPolicy) Autorelease() EPExecutionPolicy {
	rv := objc.Send[EPExecutionPolicy](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEPExecutionPolicy creates a new EPExecutionPolicy instance.
func NewEPExecutionPolicy() EPExecutionPolicy {
	return getEPExecutionPolicyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EPExecutionPolicy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPExecutionPolicy
type EPExecutionPolicy struct {
	objectivec.Object
}

// EPExecutionPolicyFrom constructs a [EPExecutionPolicy] from an unsafe.Pointer.
func EPExecutionPolicyFrom(ptr unsafe.Pointer) EPExecutionPolicy {
	return EPExecutionPolicy{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EPExecutionPolicy */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EPExecutionPolicy */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EPExecutionPolicy */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EPExecutionPolicy */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPExecutionPolicy/addException(for:)
func (e_ EPExecutionPolicy) AddPolicyExceptionForURLError(url objc.IObject /* cross-framework: NSURL */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("addPolicyExceptionForURL:error:"), url, error_)
	return rv
}/* debug [instance_methods/method]: AddPolicyExceptionForURLError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EPExecutionPolicy */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EPExecutionPolicy */


