// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentHistoryToken] class.
var (
	persistentHistoryTokenClass     _PersistentHistoryTokenClass
	persistentHistoryTokenClassOnce sync.Once
)

func getPersistentHistoryTokenClass() _PersistentHistoryTokenClass {
	persistentHistoryTokenClassOnce.Do(func() {
		persistentHistoryTokenClass = _PersistentHistoryTokenClass{objc.GetClass("NSPersistentHistoryToken")}
	})
	return persistentHistoryTokenClass
}

type _PersistentHistoryTokenClass struct {
	class objc.Class
}

// An interface definition for the [PersistentHistoryToken] class.
type IPersistentHistoryToken interface {
	objectivec.IObject
}

// A bookmark for keeping track the most recent history that you’ve processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryToken
type PersistentHistoryToken struct {
	objectivec.Object
}

// PersistentHistoryTokenFrom constructs a [PersistentHistoryToken] from an unsafe.Pointer.
//
// A bookmark for keeping track the most recent history that you’ve processed.
func PersistentHistoryTokenFrom(ptr unsafe.Pointer) PersistentHistoryToken {
	return PersistentHistoryToken{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentHistoryTokenClass) Alloc() PersistentHistoryToken {
	rv := objc.Send[PersistentHistoryToken](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentHistoryTokenClass) New() PersistentHistoryToken {
	rv := objc.Send[PersistentHistoryToken](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentHistoryToken) Init() PersistentHistoryToken {
	rv := objc.Send[PersistentHistoryToken](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentHistoryToken) Autorelease() PersistentHistoryToken {
	rv := objc.Send[PersistentHistoryToken](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentHistoryToken creates a new PersistentHistoryToken instance.
func NewPersistentHistoryToken() PersistentHistoryToken {
	return getPersistentHistoryTokenClass().New()
}




