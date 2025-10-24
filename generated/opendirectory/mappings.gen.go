// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mappings] class.
var (
	MappingsClass     _mappingsClass
	MappingsClassOnce sync.Once
)

func getmappingsClass() _mappingsClass {
	MappingsClassOnce.Do(func() {
		MappingsClass = _mappingsClass{objc.GetClass("mappings")}
	})
	return MappingsClass
}

type _mappingsClass struct {
	class objc.Class
}

// An interface definition for the [mappings] class.
type Imappings interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/mappings-c.ivar
type mappings struct {
	objectivec.Object
}

// mappingsFrom constructs a [mappings] from an unsafe.Pointer.
func mappingsFrom(ptr unsafe.Pointer) mappings {
	return mappings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mappingsClass) Alloc() mappings {
	rv := objc.Send[mappings](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mappingsClass) New() mappings {
	rv := objc.Send[mappings](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mappings) Init() mappings {
	rv := objc.Send[mappings](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mappings) Autorelease() mappings {
	rv := objc.Send[mappings](m_.ID, objc.Sel("autorelease"))
	return rv
}

// Newmappings creates a new mappings instance.
func Newmappings() mappings {
	return getmappingsClass().New()
}




