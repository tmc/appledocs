// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [rootObject] class.
var (
	RootObjectClass     _rootObjectClass
	RootObjectClassOnce sync.Once
)

func getrootObjectClass() _rootObjectClass {
	RootObjectClassOnce.Do(func() {
		RootObjectClass = _rootObjectClass{objc.GetClass("rootObject")}
	})
	return RootObjectClass
}

type _rootObjectClass struct {
	class objc.Class
}

// An interface definition for the [rootObject] class.
type IrootObject interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/rootObject-c.ivar

type rootObject struct {
	objectivec.Object
}

// rootObjectFrom constructs a [rootObject] from an unsafe.Pointer.
func rootObjectFrom(ptr unsafe.Pointer) rootObject {
	return rootObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _rootObjectClass) Alloc() rootObject {
	rv := objc.Send[rootObject](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _rootObjectClass) New() rootObject {
	rv := objc.Send[rootObject](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ rootObject) Init() rootObject {
	rv := objc.Send[rootObject](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ rootObject) Autorelease() rootObject {
	rv := objc.Send[rootObject](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrootObject creates a new rootObject instance.
func NewrootObject() rootObject {
	return getrootObjectClass().New()
}




