// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ExtensionContext] class.
var (
	ExtensionContextClass     _ExtensionContextClass
	ExtensionContextClassOnce sync.Once
)

func getExtensionContextClass() _ExtensionContextClass {
	ExtensionContextClassOnce.Do(func() {
		ExtensionContextClass = _ExtensionContextClass{objc.GetClass("NSExtensionContext")}
	})
	return ExtensionContextClass
}

type _ExtensionContextClass struct {
	class objc.Class
}

// An interface definition for the [ExtensionContext] class.
type IExtensionContext interface {
	objectivec.IObject
}

// A parent class referenced by other CallKit classes.


// A parent class referenced by other CallKit classes. [Full Topic]

type ExtensionContext struct {
	objectivec.Object
}

// ExtensionContextFrom constructs a [ExtensionContext] from an unsafe.Pointer.
//
// A parent class referenced by other CallKit classes.
func ExtensionContextFrom(ptr unsafe.Pointer) ExtensionContext {
	return ExtensionContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _ExtensionContextClass) Alloc() ExtensionContext {
	rv := objc.Send[ExtensionContext](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _ExtensionContextClass) New() ExtensionContext {
	rv := objc.Send[ExtensionContext](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExtensionContext) Init() ExtensionContext {
	rv := objc.Send[ExtensionContext](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExtensionContext) Autorelease() ExtensionContext {
	rv := objc.Send[ExtensionContext](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExtensionContext creates a new ExtensionContext instance.
func NewExtensionContext() ExtensionContext {
	return getExtensionContextClass().New()
}




