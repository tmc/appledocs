// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [identifier] class.
var (
	IdentifierClass     _identifierClass
	IdentifierClassOnce sync.Once
)

func getidentifierClass() _identifierClass {
	IdentifierClassOnce.Do(func() {
		IdentifierClass = _identifierClass{objc.GetClass("identifier")}
	})
	return IdentifierClass
}

type _identifierClass struct {
	class objc.Class
}

// An interface definition for the [identifier] class.
type Iidentifier interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/identifier-c.ivar
type identifier struct {
	objectivec.Object
}

// identifierFrom constructs a [identifier] from an unsafe.Pointer.
func identifierFrom(ptr unsafe.Pointer) identifier {
	return identifier{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _identifierClass) Alloc() identifier {
	rv := objc.Send[identifier](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _identifierClass) New() identifier {
	rv := objc.Send[identifier](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ identifier) Init() identifier {
	rv := objc.Send[identifier](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ identifier) Autorelease() identifier {
	rv := objc.Send[identifier](i_.ID, objc.Sel("autorelease"))
	return rv
}

// Newidentifier creates a new identifier instance.
func Newidentifier() identifier {
	return getidentifierClass().New()
}




