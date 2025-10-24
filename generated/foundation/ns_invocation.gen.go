// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Invocation] class.
type IInvocation interface {
	objectivec.IObject
	// properties:
	ArgumentsRetained() bool
	MethodSignature() objc.IObject /* cross-framework: MethodSignature */
	Selector() objc.SEL
	SetSelector(value objc.SEL)
	Target() objc.ID
	SetTarget(value objc.ID)
	// methods:
	GetArgumentAtIndex(argumentLocation unsafe.Pointer, idx int)
	GetReturnValue(retLoc unsafe.Pointer)
	Invoke()
	InvokeUsingIMP(imp unsafe.Pointer)
	InvokeWithTarget(target objectivec.IObject)
	RetainArguments()
	SetArgumentAtIndex(argumentLocation unsafe.Pointer, idx int)
	SetReturnValue(retLoc unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (ic _InvocationClass) Alloc() Invocation {
	rv := objc.Send[Invocation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns an object able to construct messages using a given method signature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/invocationWithMethodSignature:
func (ic _InvocationClass) InvocationWithMethodSignature(sig objc.IObject /* cross-framework: MethodSignature */) IInvocation {
	rv := objc.Send[Invocation](objc.ID(ic.class), objc.Sel("invocationWithMethodSignature:"), sig)
	return rv
}


// Returns by indirection the receiver’s argument at a specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/getArgument:atIndex:
func (i_ Invocation) GetArgumentAtIndex(argumentLocation unsafe.Pointer, idx int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("getArgument:atIndex:"), argumentLocation, idx)
}


// Gets the invocation’s return value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/getReturnValue:
func (i_ Invocation) GetReturnValue(retLoc unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("getReturnValue:"), retLoc)
}


// Sends the receiver’s message (with arguments) to its target and sets the return value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/invoke
func (i_ Invocation) Invoke() {
	objc.Send[objc.ID](i_.ID, objc.Sel("invoke"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/invokeUsingIMP:
func (i_ Invocation) InvokeUsingIMP(imp unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("invokeUsingIMP:"), imp)
}


// Sets the receiver’s target, sends the receiver’s message (with arguments) to that target, and sets the return value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/invokeWithTarget:
func (i_ Invocation) InvokeWithTarget(target objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("invokeWithTarget:"), target)
}


// If the receiver hasn’t already done so, retains the target and all object arguments of the receiver and copies all of its C-string arguments and blocks. If a returnvalue has been set, this is also retained or copied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/retainArguments
func (i_ Invocation) RetainArguments() {
	objc.Send[objc.ID](i_.ID, objc.Sel("retainArguments"))
}


// Sets an argument of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/setArgument:atIndex:
func (i_ Invocation) SetArgumentAtIndex(argumentLocation unsafe.Pointer, idx int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setArgument:atIndex:"), argumentLocation, idx)
}


// Sets the receiver’s return value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/setReturnValue:
func (i_ Invocation) SetReturnValue(retLoc unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReturnValue:"), retLoc)
}


// A Boolean value that indicates if the receiver has retained its arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/argumentsRetained
func (i_ Invocation) ArgumentsRetained() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("argumentsRetained"))
	return rv
}


// The receiver’s method signature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/methodSignature
func (i_ Invocation) MethodSignature() objc.IObject /* cross-framework: MethodSignature */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("methodSignature"))
	return rv
}


// The receiver’s selector, or 0 if it hasn’t been set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/selector
func (i_ Invocation) Selector() objc.SEL {
	rv := objc.Send[objc.SEL](i_.ID, objc.Sel("selector"))
	return rv
}


// The receiver’s selector, or 0 if it hasn’t been set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/selector
func (i_ Invocation) SetSelector(value objc.SEL) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSelector:"), value)
}


// The receiver’s target, or if the receiver has no target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/target
func (i_ Invocation) Target() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("target"))
	return rv
}


// The receiver’s target, or if the receiver has no target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/target
func (i_ Invocation) SetTarget(value objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTarget:"), value)
}



