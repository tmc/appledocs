// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CoreDataCoreSpotlightDelegate] class.
var (
	coreDataCoreSpotlightDelegateClass     _CoreDataCoreSpotlightDelegateClass
	coreDataCoreSpotlightDelegateClassOnce sync.Once
)

func getCoreDataCoreSpotlightDelegateClass() _CoreDataCoreSpotlightDelegateClass {
	coreDataCoreSpotlightDelegateClassOnce.Do(func() {
		coreDataCoreSpotlightDelegateClass = _CoreDataCoreSpotlightDelegateClass{objc.GetClass("NSCoreDataCoreSpotlightDelegate")}
	})
	return coreDataCoreSpotlightDelegateClass
}

type _CoreDataCoreSpotlightDelegateClass struct {
	class objc.Class
}

// An interface definition for the [CoreDataCoreSpotlightDelegate] class.
type ICoreDataCoreSpotlightDelegate interface {
	objectivec.IObject
}

// A set of methods that enable integration with Core Spotlight.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate
type CoreDataCoreSpotlightDelegate struct {
	objectivec.Object
}

// CoreDataCoreSpotlightDelegateFrom constructs a [CoreDataCoreSpotlightDelegate] from an unsafe.Pointer.
//
// A set of methods that enable integration with Core Spotlight.
func CoreDataCoreSpotlightDelegateFrom(ptr unsafe.Pointer) CoreDataCoreSpotlightDelegate {
	return CoreDataCoreSpotlightDelegate{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CoreDataCoreSpotlightDelegateClass) Alloc() CoreDataCoreSpotlightDelegate {
	rv := objc.Send[CoreDataCoreSpotlightDelegate](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CoreDataCoreSpotlightDelegateClass) New() CoreDataCoreSpotlightDelegate {
	rv := objc.Send[CoreDataCoreSpotlightDelegate](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CoreDataCoreSpotlightDelegate) Init() CoreDataCoreSpotlightDelegate {
	rv := objc.Send[CoreDataCoreSpotlightDelegate](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CoreDataCoreSpotlightDelegate) Autorelease() CoreDataCoreSpotlightDelegate {
	rv := objc.Send[CoreDataCoreSpotlightDelegate](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreDataCoreSpotlightDelegate creates a new CoreDataCoreSpotlightDelegate instance.
func NewCoreDataCoreSpotlightDelegate() CoreDataCoreSpotlightDelegate {
	return getCoreDataCoreSpotlightDelegateClass().New()
}




