// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [components] class.
var (
	ComponentsClass     _componentsClass
	ComponentsClassOnce sync.Once
)

func getcomponentsClass() _componentsClass {
	ComponentsClassOnce.Do(func() {
		ComponentsClass = _componentsClass{objc.GetClass("components")}
	})
	return ComponentsClass
}

type _componentsClass struct {
	class objc.Class
}

// An interface definition for the [components] class.
type Icomponents interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortMessage/components-c.ivar
type components struct {
	objectivec.Object
}

// componentsFrom constructs a [components] from an unsafe.Pointer.
func componentsFrom(ptr unsafe.Pointer) components {
	return components{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _componentsClass) Alloc() components {
	rv := objc.Send[components](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _componentsClass) New() components {
	rv := objc.Send[components](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ components) Init() components {
	rv := objc.Send[components](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ components) Autorelease() components {
	rv := objc.Send[components](c_.ID, objc.Sel("autorelease"))
	return rv
}

// Newcomponents creates a new components instance.
func Newcomponents() components {
	return getcomponentsClass().New()
}
