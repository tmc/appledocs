// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MethodSignature] class.
var (
	MethodSignatureClass     _MethodSignatureClass
	MethodSignatureClassOnce sync.Once
)

func getMethodSignatureClass() _MethodSignatureClass {
	MethodSignatureClassOnce.Do(func() {
		MethodSignatureClass = _MethodSignatureClass{objc.GetClass("NSMethodSignature")}
	})
	return MethodSignatureClass
}

type _MethodSignatureClass struct {
	class objc.Class
}

// An interface definition for the [MethodSignature] class.
type IMethodSignature interface {
	objectivec.IObject
	// properties:
	FrameLength() uint
	MethodReturnLength() uint
	MethodReturnType() unsafe.Pointer
	NumberOfArguments() uint
	// methods:
	GetArgumentTypeAtIndex(idx uint) unsafe.Pointer
	IsOneway() bool
}

// A record of the type information for the return value and parameters of a method.
//
// Use an object to forward messages that the receiving object does not respond to—most notably in the case of distributed objects. You typically create an object using the instance method (in macOS 10.5 and later you can also use ). It is then used to create an object, which is passed as the argument to a message to send the invocation on to whatever other object can handle the message. In the default case, invokes , which raises an exception. For distributed objects, the object is encoded using the information in the object and sent to the real object represented by the receiver of the message.


// A record of the type information for the return value and parameters of a method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMethodSignature
type MethodSignature struct {
	objectivec.Object
}

// MethodSignatureFrom constructs a [MethodSignature] from an unsafe.Pointer.
//
// A record of the type information for the return value and parameters of a method.
func MethodSignatureFrom(ptr unsafe.Pointer) MethodSignature {
	return MethodSignature{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MethodSignatureClass) Alloc() MethodSignature {
	rv := objc.Send[MethodSignature](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MethodSignatureClass) New() MethodSignature {
	rv := objc.Send[MethodSignature](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MethodSignature) Init() MethodSignature {
	rv := objc.Send[MethodSignature](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MethodSignature) Autorelease() MethodSignature {
	rv := objc.Send[MethodSignature](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMethodSignature creates a new MethodSignature instance.
func NewMethodSignature() MethodSignature {
	return getMethodSignatureClass().New()
}



// Returns an object for the given Objective-C method type string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMethodSignature/signatureWithObjCTypes:
func (mc _MethodSignatureClass) SignatureWithObjCTypes(types unsafe.Pointer) IMethodSignature {
	rv := objc.Send[MethodSignature](objc.ID(mc.class), objc.Sel("signatureWithObjCTypes:"), types)
	return rv
}


// Returns the type encoding for the argument at a given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMethodSignature/getArgumentTypeAtIndex:
func (m_ MethodSignature) GetArgumentTypeAtIndex(idx uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("getArgumentTypeAtIndex:"), idx)
	return rv
}


// Whether the receiver is asynchronous when invoked through distributed objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMethodSignature/isOneway
func (m_ MethodSignature) IsOneway() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isOneway"))
	return rv
}


// The number of bytes that the arguments, taken together, occupy on the stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMethodSignature/frameLength
func (m_ MethodSignature) FrameLength() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("frameLength"))
	return rv
}


// The number of bytes required for the return value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMethodSignature/methodReturnLength
func (m_ MethodSignature) MethodReturnLength() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("methodReturnLength"))
	return rv
}


// A C string encoding the return type of the method in Objective-C type encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMethodSignature/methodReturnType
func (m_ MethodSignature) MethodReturnType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("methodReturnType"))
	return rv
}


// The number of arguments recorded in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMethodSignature/numberOfArguments
func (m_ MethodSignature) NumberOfArguments() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("numberOfArguments"))
	return rv
}



