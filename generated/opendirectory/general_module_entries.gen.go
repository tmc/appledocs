// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [generalModuleEntries] class.
var (
	GeneralModuleEntriesClass     _generalModuleEntriesClass
	GeneralModuleEntriesClassOnce sync.Once
)

func getgeneralModuleEntriesClass() _generalModuleEntriesClass {
	GeneralModuleEntriesClassOnce.Do(func() {
		GeneralModuleEntriesClass = _generalModuleEntriesClass{objc.GetClass("generalModuleEntries")}
	})
	return GeneralModuleEntriesClass
}

type _generalModuleEntriesClass struct {
	class objc.Class
}

// An interface definition for the [generalModuleEntries] class.
type IgeneralModuleEntries interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/generalModuleEntries-c.ivar
type generalModuleEntries struct {
	objectivec.Object
}

// generalModuleEntriesFrom constructs a [generalModuleEntries] from an unsafe.Pointer.
func generalModuleEntriesFrom(ptr unsafe.Pointer) generalModuleEntries {
	return generalModuleEntries{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _generalModuleEntriesClass) Alloc() generalModuleEntries {
	rv := objc.Send[generalModuleEntries](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _generalModuleEntriesClass) New() generalModuleEntries {
	rv := objc.Send[generalModuleEntries](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ generalModuleEntries) Init() generalModuleEntries {
	rv := objc.Send[generalModuleEntries](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ generalModuleEntries) Autorelease() generalModuleEntries {
	rv := objc.Send[generalModuleEntries](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewgeneralModuleEntries creates a new generalModuleEntries instance.
func NewgeneralModuleEntries() generalModuleEntries {
	return getgeneralModuleEntriesClass().New()
}




