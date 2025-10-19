// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MethodSignature] class.
var methodSignatureClass = _MethodSignatureClass{objc.GetClass("NSMethodSignature")}

type _MethodSignatureClass struct {
	class objc.Class
}

// An interface definition for the [MethodSignature] class.
type IMethodSignature interface {
	objectivec.IObject
}

// A record of the type information for the return value and parameters of a method. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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
	return methodSignatureClass.New()
}




