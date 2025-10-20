// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [defaultModuleEntries] class.
var (
	DefaultModuleEntriesClass     _defaultModuleEntriesClass
	DefaultModuleEntriesClassOnce sync.Once
)

func getdefaultModuleEntriesClass() _defaultModuleEntriesClass {
	DefaultModuleEntriesClassOnce.Do(func() {
		DefaultModuleEntriesClass = _defaultModuleEntriesClass{objc.GetClass("defaultModuleEntries")}
	})
	return DefaultModuleEntriesClass
}

type _defaultModuleEntriesClass struct {
	class objc.Class
}

// An interface definition for the [defaultModuleEntries] class.
type IdefaultModuleEntries interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultModuleEntries-c.ivar
type defaultModuleEntries struct {
	objectivec.Object
}

// defaultModuleEntriesFrom constructs a [defaultModuleEntries] from an unsafe.Pointer.
func defaultModuleEntriesFrom(ptr unsafe.Pointer) defaultModuleEntries {
	return defaultModuleEntries{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _defaultModuleEntriesClass) Alloc() defaultModuleEntries {
	rv := objc.Send[defaultModuleEntries](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _defaultModuleEntriesClass) New() defaultModuleEntries {
	rv := objc.Send[defaultModuleEntries](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ defaultModuleEntries) Init() defaultModuleEntries {
	rv := objc.Send[defaultModuleEntries](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ defaultModuleEntries) Autorelease() defaultModuleEntries {
	rv := objc.Send[defaultModuleEntries](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewdefaultModuleEntries creates a new defaultModuleEntries instance.
func NewdefaultModuleEntries() defaultModuleEntries {
	return getdefaultModuleEntriesClass().New()
}




