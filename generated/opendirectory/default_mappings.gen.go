// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [defaultMappings] class.
var (
	DefaultMappingsClass     _defaultMappingsClass
	DefaultMappingsClassOnce sync.Once
)

func getdefaultMappingsClass() _defaultMappingsClass {
	DefaultMappingsClassOnce.Do(func() {
		DefaultMappingsClass = _defaultMappingsClass{objc.GetClass("defaultMappings")}
	})
	return DefaultMappingsClass
}

type _defaultMappingsClass struct {
	class objc.Class
}

// An interface definition for the [defaultMappings] class.
type IdefaultMappings interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultMappings-c.ivar
type defaultMappings struct {
	objectivec.Object
}

// defaultMappingsFrom constructs a [defaultMappings] from an unsafe.Pointer.
func defaultMappingsFrom(ptr unsafe.Pointer) defaultMappings {
	return defaultMappings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _defaultMappingsClass) Alloc() defaultMappings {
	rv := objc.Send[defaultMappings](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _defaultMappingsClass) New() defaultMappings {
	rv := objc.Send[defaultMappings](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ defaultMappings) Init() defaultMappings {
	rv := objc.Send[defaultMappings](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ defaultMappings) Autorelease() defaultMappings {
	rv := objc.Send[defaultMappings](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewdefaultMappings creates a new defaultMappings instance.
func NewdefaultMappings() defaultMappings {
	return getdefaultMappingsClass().New()
}




