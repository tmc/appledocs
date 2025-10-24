// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [discoveryModuleEntries] class.
var (
	DiscoveryModuleEntriesClass     _discoveryModuleEntriesClass
	DiscoveryModuleEntriesClassOnce sync.Once
)

func getdiscoveryModuleEntriesClass() _discoveryModuleEntriesClass {
	DiscoveryModuleEntriesClassOnce.Do(func() {
		DiscoveryModuleEntriesClass = _discoveryModuleEntriesClass{objc.GetClass("discoveryModuleEntries")}
	})
	return DiscoveryModuleEntriesClass
}

type _discoveryModuleEntriesClass struct {
	class objc.Class
}

// An interface definition for the [discoveryModuleEntries] class.
type IdiscoveryModuleEntries interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/discoveryModuleEntries-c.ivar
type discoveryModuleEntries struct {
	objectivec.Object
}

// discoveryModuleEntriesFrom constructs a [discoveryModuleEntries] from an unsafe.Pointer.
func discoveryModuleEntriesFrom(ptr unsafe.Pointer) discoveryModuleEntries {
	return discoveryModuleEntries{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _discoveryModuleEntriesClass) Alloc() discoveryModuleEntries {
	rv := objc.Send[discoveryModuleEntries](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _discoveryModuleEntriesClass) New() discoveryModuleEntries {
	rv := objc.Send[discoveryModuleEntries](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ discoveryModuleEntries) Init() discoveryModuleEntries {
	rv := objc.Send[discoveryModuleEntries](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ discoveryModuleEntries) Autorelease() discoveryModuleEntries {
	rv := objc.Send[discoveryModuleEntries](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewdiscoveryModuleEntries creates a new discoveryModuleEntries instance.
func NewdiscoveryModuleEntries() discoveryModuleEntries {
	return getdiscoveryModuleEntriesClass().New()
}




