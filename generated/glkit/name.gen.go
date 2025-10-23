// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [name] class.
var (
	NameClass     _nameClass
	NameClassOnce sync.Once
)

func getnameClass() _nameClass {
	NameClassOnce.Do(func() {
		NameClass = _nameClass{objc.GetClass("name")}
	})
	return NameClass
}

type _nameClass struct {
	class objc.Class
}

// An interface definition for the [name] class.
type Iname interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/name-c.ivar
type name struct {
	objectivec.Object
}

// nameFrom constructs a [name] from an unsafe.Pointer.
func nameFrom(ptr unsafe.Pointer) name {
	return name{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _nameClass) Alloc() name {
	rv := objc.Send[name](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _nameClass) New() name {
	rv := objc.Send[name](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ name) Init() name {
	rv := objc.Send[name](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ name) Autorelease() name {
	rv := objc.Send[name](n_.ID, objc.Sel("autorelease"))
	return rv
}

// Newname creates a new name instance.
func Newname() name {
	return getnameClass().New()
}




