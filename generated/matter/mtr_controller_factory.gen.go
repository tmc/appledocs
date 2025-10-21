// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRControllerFactory] class.
var (
	MTRControllerFactoryClass     _MTRControllerFactoryClass
	MTRControllerFactoryClassOnce sync.Once
)

func getMTRControllerFactoryClass() _MTRControllerFactoryClass {
	MTRControllerFactoryClassOnce.Do(func() {
		MTRControllerFactoryClass = _MTRControllerFactoryClass{objc.GetClass("MTRControllerFactory")}
	})
	return MTRControllerFactoryClass
}

type _MTRControllerFactoryClass struct {
	class objc.Class
}

// An interface definition for the [MTRControllerFactory] class.
type IMTRControllerFactory interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRControllerFactory
type MTRControllerFactory struct {
	objectivec.Object
}

// MTRControllerFactoryFrom constructs a [MTRControllerFactory] from an unsafe.Pointer.
func MTRControllerFactoryFrom(ptr unsafe.Pointer) MTRControllerFactory {
	return MTRControllerFactory{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRControllerFactoryClass) Alloc() MTRControllerFactory {
	rv := objc.Send[MTRControllerFactory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRControllerFactoryClass) New() MTRControllerFactory {
	rv := objc.Send[MTRControllerFactory](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRControllerFactory) Init() MTRControllerFactory {
	rv := objc.Send[MTRControllerFactory](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRControllerFactory) Autorelease() MTRControllerFactory {
	rv := objc.Send[MTRControllerFactory](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRControllerFactory creates a new MTRControllerFactory instance.
func NewMTRControllerFactory() MTRControllerFactory {
	return getMTRControllerFactoryClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontrollerfactory/isrunning
func (m_ MTRControllerFactory) IsRunning() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isRunning"))
	return rv
}


// SetIsRunning sets the value of the isRunning property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontrollerfactory/isrunning
func (m_ MTRControllerFactory) SetIsRunning(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsRunning:"), value)
}



