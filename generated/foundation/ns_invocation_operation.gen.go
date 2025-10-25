// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSInvocationOperation */


/* debug [class_header]: Header for NSInvocationOperation */
// The class instance for the [InvocationOperation] class.
var (
	InvocationOperationClass     _InvocationOperationClass
	InvocationOperationClassOnce sync.Once
)

func getInvocationOperationClass() _InvocationOperationClass {
	InvocationOperationClassOnce.Do(func() {
		InvocationOperationClass = _InvocationOperationClass{objc.GetClass("NSInvocationOperation")}
	})
	return InvocationOperationClass
}

type _InvocationOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InvocationOperation */
// An interface definition for the [InvocationOperation] class.
type IInvocationOperation interface {
	IOperation
	
/* debug [class_interface_properties]: Properties for InvocationOperation */
	// properties:
	Invocation() IInvocation
	Result() objc.ID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InvocationOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InvocationOperation */
// Alloc allocates a new instance without initialization.
func (ic _InvocationOperationClass) Alloc() InvocationOperation {
	rv := objc.Send[InvocationOperation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _InvocationOperationClass) New() InvocationOperation {
	rv := objc.Send[InvocationOperation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InvocationOperation) Init() InvocationOperation {
	rv := objc.Send[InvocationOperation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InvocationOperation) Autorelease() InvocationOperation {
	rv := objc.Send[InvocationOperation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInvocationOperation creates a new InvocationOperation instance.
func NewInvocationOperation() InvocationOperation {
	return getInvocationOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InvocationOperation */
// An operation that manages the execution of a single encapsulated task specified as an invocation.
//
// The class is a concrete subclass of that you use to initiate an operation that consists of invoking a selector on a specified object. This class implements a non-concurrent operation. For more information on concurrent versus non-concurrent operations, see .


// An operation that manages the execution of a single encapsulated task specified as an invocation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocationOperation
type InvocationOperation struct {
	Operation
}

// InvocationOperationFrom constructs a [InvocationOperation] from an unsafe.Pointer.
//
// An operation that manages the execution of a single encapsulated task specified as an invocation.
func InvocationOperationFrom(ptr unsafe.Pointer) InvocationOperation {
	return InvocationOperation{
		Operation: OperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InvocationOperation */

// Returns an object initialized with the specified invocation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocationOperation/initWithInvocation:
func NewInvocationOperationWithInvocation(inv IInvocation) InvocationOperation {
	instance := getInvocationOperationClass().Alloc()
	rv := objc.Send[InvocationOperation](instance.ID, objc.Sel("initWithInvocation:"), inv)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewInvocationOperationWithInvocation */


// Returns an object initialized with the specified target and selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocationOperation/initWithTarget:selector:object:
func NewInvocationOperationWithTargetSelectorObject(target objc.IObject, sel objc.SEL, arg objc.IObject) InvocationOperation {
	instance := getInvocationOperationClass().Alloc()
	rv := objc.Send[InvocationOperation](instance.ID, objc.Sel("initWithTarget:selector:object:"), target, sel, arg)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewInvocationOperationWithTargetSelectorObject */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InvocationOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InvocationOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InvocationOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InvocationOperation */

// The receiver’s invocation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocationOperation/invocation
func (i_ InvocationOperation) Invocation() IInvocation {
	rv := objc.Send[Invocation](i_.ID, objc.Sel("invocation"))
	return rv
}/* debug [instance_properties/getter]: invocation */


// The result of the invocation or method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocationOperation/result
func (i_ InvocationOperation) Result() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("result"))
	return rv
}/* debug [instance_properties/getter]: result */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSInvocationOperation */


