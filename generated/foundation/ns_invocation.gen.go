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
	RetainArguments()
	Target() objc.ID
	SetTarget(value objc.ID)
}

// An Objective-C message rendered as an object.
//
// objects are used to store and forward messages between objects and between applications, primarily by objects and the distributed objects system. An object contains all the elements of an Objective-C message: a target, a selector, arguments, and the return value. Each of these elements can be set directly, and the return value is set automatically when the object is dispatched. An object can be repeatedly dispatched to different targets; its arguments can be modified between dispatch for varying results; even its selector can be changed to another with the same method signature (argument and return types). This flexibility makes useful for repeating messages with many arguments and variations; rather than retyping a slightly different expression for each message, you modify the object as needed each time before dispatching it to a new target. does not support invocations of methods with either variable numbers of arguments or arguments. You should use the class method to create objects; you should not create these objects using and . This class does not retain the arguments for the contained invocation by default. If those objects might disappear between the time you create your instance of and the time you use it, you should explicitly retain the objects yourself or invoke the method to have the invocation object retain them itself.
//
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/invocationWithMethodSignature:
func (ic _InvocationClass) InvocationWithMethodSignature(sig IMethodSignature) Invocation {
	rv := objc.Send[Invocation](objc.ID(ic.class), objc.Sel("invocationWithMethodSignature:"), sig)
	return rv
}

// If the receiver hasn’t already done so, retains the target and all object arguments of the receiver and copies all of its C-string arguments and blocks. If a returnvalue has been set, this is also retained or copied.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/retainArguments
func (i_ Invocation) RetainArguments() {
	objc.Send[objc.ID](i_.ID, objc.Sel("retainArguments"))
}

// The receiver’s target, or if the receiver has no target.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/target
func (i_ Invocation) Target() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
// The receiver’s target, or if the receiver has no target.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation/target
func (i_ Invocation) SetTarget(value objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTarget:"), value)
}



