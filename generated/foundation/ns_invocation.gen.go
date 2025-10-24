// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSInvocation */


/* debug [class_header]: Header for NSInvocation */
// The class instance for the [Invocation] class.
var (
	InvocationClass     _InvocationClass
	InvocationClassOnce sync.Once
)

func getInvocationClass() _InvocationClass {
	InvocationClassOnce.Do(func() {
		InvocationClass = _InvocationClass{objc.GetClass("NSInvocation")}
	})
	return InvocationClass
}

type _InvocationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Invocation */
// An interface definition for the [Invocation] class.
type IInvocation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Invocation */
	// properties:
	ArgumentsRetained() bool
	MethodSignature() MethodSignature /* not a class type */
	Selector() objc.SEL
	SetSelector(value objc.SEL)
	Target() objc.ID
	SetTarget(value objc.ID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Invocation */
	// methods:
	GetArgumentAtIndex(argumentLocation objectivec.IObject, idx int)
	GetReturnValue(retLoc objectivec.IObject)
	Invoke()
	InvokeUsingIMP(imp objectivec.IObject)
	InvokeWithTarget(target objc.IObject)
	RetainArguments()
	SetArgumentAtIndex(argumentLocation objectivec.IObject, idx int)
	SetReturnValue(retLoc objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Invocation */
// Alloc allocates a new instance without initialization.
func (ic _InvocationClass) Alloc() Invocation {
	rv := objc.Send[Invocation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _InvocationClass) New() Invocation {
	rv := objc.Send[Invocation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ Invocation) Init() Invocation {
	rv := objc.Send[Invocation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ Invocation) Autorelease() Invocation {
	rv := objc.Send[Invocation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInvocation creates a new Invocation instance.
func NewInvocation() Invocation {
	return getInvocationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Invocation */
// An Objective-C message rendered as an object.
//
// objects are used to store and forward messages between objects and between applications, primarily by objects and the distributed objects system. An object contains all the elements of an Objective-C message: a target, a selector, arguments, and the return value. Each of these elements can be set directly, and the return value is set automatically when the object is dispatched. An object can be repeatedly dispatched to different targets; its arguments can be modified between dispatch for varying results; even its selector can be changed to another with the same method signature (argument and return types). This flexibility makes useful for repeating messages with many arguments and variations; rather than retyping a slightly different expression for each message, you modify the object as needed each time before dispatching it to a new target. does not support invocations of methods with either variable numbers of arguments or arguments. You should use the class method to create objects; you should not create these objects using and . This class does not retain the arguments for the contained invocation by default. If those objects might disappear between the time you create your instance of and the time you use it, you should explicitly retain the objects yourself or invoke the method to have the invocation object retain them itself.


// An Objective-C message rendered as an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation
type Invocation struct {
	objectivec.Object
}

// InvocationFrom constructs a [Invocation] from an unsafe.Pointer.
//
// An Objective-C message rendered as an object.
func InvocationFrom(ptr unsafe.Pointer) Invocation {
	return Invocation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Invocation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Invocation */

// Returns an object able to construct messages using a given method signature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/invocationWithMethodSignature:
func (ic _InvocationClass) InvocationWithMethodSignature(sig MethodSignature /* not a class type */) IInvocation {
	rv := objc.Send[Invocation](objc.ID(ic.class), objc.Sel("invocationWithMethodSignature:"), sig)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InvocationWithMethodSignature) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Invocation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Invocation */

// Returns by indirection the receiver’s argument at a specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/getArgument:atIndex:
func (i_ Invocation) GetArgumentAtIndex(argumentLocation objectivec.IObject, idx int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("getArgument:atIndex:"), argumentLocation, idx)
}/* debug [instance_methods/method]: GetArgumentAtIndex */


// Gets the invocation’s return value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/getReturnValue:
func (i_ Invocation) GetReturnValue(retLoc objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("getReturnValue:"), retLoc)
}/* debug [instance_methods/method]: GetReturnValue */


// Sends the receiver’s message (with arguments) to its target and sets the return value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/invoke
func (i_ Invocation) Invoke() {
	objc.Send[objc.ID](i_.ID, objc.Sel("invoke"))
}/* debug [instance_methods/method]: Invoke */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/invokeUsingIMP:
func (i_ Invocation) InvokeUsingIMP(imp objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("invokeUsingIMP:"), imp)
}/* debug [instance_methods/method]: InvokeUsingIMP */


// Sets the receiver’s target, sends the receiver’s message (with arguments) to that target, and sets the return value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/invokeWithTarget:
func (i_ Invocation) InvokeWithTarget(target objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("invokeWithTarget:"), target)
}/* debug [instance_methods/method]: InvokeWithTarget */


// If the receiver hasn’t already done so, retains the target and all object arguments of the receiver and copies all of its C-string arguments and blocks. If a returnvalue has been set, this is also retained or copied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/retainArguments
func (i_ Invocation) RetainArguments() {
	objc.Send[objc.ID](i_.ID, objc.Sel("retainArguments"))
}/* debug [instance_methods/method]: RetainArguments */


// Sets an argument of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/setArgument:atIndex:
func (i_ Invocation) SetArgumentAtIndex(argumentLocation objectivec.IObject, idx int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setArgument:atIndex:"), argumentLocation, idx)
}/* debug [instance_methods/method]: SetArgumentAtIndex */


// Sets the receiver’s return value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/setReturnValue:
func (i_ Invocation) SetReturnValue(retLoc objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReturnValue:"), retLoc)
}/* debug [instance_methods/method]: SetReturnValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Invocation */

// A Boolean value that indicates if the receiver has retained its arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/argumentsRetained
func (i_ Invocation) ArgumentsRetained() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("argumentsRetained"))
	return rv
}/* debug [instance_properties/getter]: argumentsRetained */


// The receiver’s method signature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/methodSignature
func (i_ Invocation) MethodSignature() MethodSignature /* not a class type */ {
	rv := objc.Send[MethodSignature](i_.ID, objc.Sel("methodSignature"))
	return rv
}/* debug [instance_properties/getter]: methodSignature */


// The receiver’s selector, or 0 if it hasn’t been set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/selector
func (i_ Invocation) Selector() objc.SEL {
	rv := objc.Send[objc.SEL](i_.ID, objc.Sel("selector"))
	return rv
}/* debug [instance_properties/getter]: selector */


// The receiver’s selector, or 0 if it hasn’t been set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/selector
func (i_ Invocation) SetSelector(value objc.SEL) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSelector:"), value)
}/* debug [instance_properties/setter]: selector */


// The receiver’s target, or if the receiver has no target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/target
func (i_ Invocation) Target() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */


// The receiver’s target, or if the receiver has no target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/target
func (i_ Invocation) SetTarget(value objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTarget:"), value)
}/* debug [instance_properties/setter]: target */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSInvocation */



