// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentStoreDescription] class.
var (
	persistentStoreDescriptionClass     _PersistentStoreDescriptionClass
	persistentStoreDescriptionClassOnce sync.Once
)

func getPersistentStoreDescriptionClass() _PersistentStoreDescriptionClass {
	persistentStoreDescriptionClassOnce.Do(func() {
		persistentStoreDescriptionClass = _PersistentStoreDescriptionClass{objc.GetClass("NSPersistentStoreDescription")}
	})
	return persistentStoreDescriptionClass
}

type _PersistentStoreDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [PersistentStoreDescription] class.
type IPersistentStoreDescription interface {
	objectivec.IObject
	SetOptionForKey(option unsafe.Pointer, key string)
	SetValueForPragmaNamed(value unsafe.Pointer, name string)
}

// A description object used to create and load a persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription
type PersistentStoreDescription struct {
	objectivec.Object
}

// PersistentStoreDescriptionFrom constructs a [PersistentStoreDescription] from an unsafe.Pointer.
//
// A description object used to create and load a persistent store.
func PersistentStoreDescriptionFrom(ptr unsafe.Pointer) PersistentStoreDescription {
	return PersistentStoreDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentStoreDescriptionClass) Alloc() PersistentStoreDescription {
	rv := objc.Send[PersistentStoreDescription](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentStoreDescriptionClass) New() PersistentStoreDescription {
	rv := objc.Send[PersistentStoreDescription](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentStoreDescription) Init() PersistentStoreDescription {
	rv := objc.Send[PersistentStoreDescription](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentStoreDescription) Autorelease() PersistentStoreDescription {
	rv := objc.Send[PersistentStoreDescription](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentStoreDescription creates a new PersistentStoreDescription instance.
func NewPersistentStoreDescription() PersistentStoreDescription {
	return getPersistentStoreDescriptionClass().New()
}


// Initializes the receiver with a URL for the store. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/init(url:)
func NewPersistentStoreDescriptionWithURL(url unsafe.Pointer) PersistentStoreDescription {
	instance := getPersistentStoreDescriptionClass().Alloc()
	rv := objc.Send[PersistentStoreDescription](instance.ID, objc.Sel("initWithURL:"), url)
	rv.Autorelease()
	return rv
}


// Initializes and returns a persistent store description with the given URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/persistentStoreDescriptionWithURL:
func (pc _PersistentStoreDescriptionClass) PersistentStoreDescriptionWithURL(URL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("persistentStoreDescriptionWithURL:"), URL)
	return rv
}
// Sets an option on the store. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/setOption(_:forKey:)
func (p_ PersistentStoreDescription) SetOptionForKey(option unsafe.Pointer, key string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOption:forKey:"), option, objc.String(key))
}
// Allows you to set pragmas for the SQLite store. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/setValue(_:forPragmaNamed:)
func (p_ PersistentStoreDescription) SetValueForPragmaNamed(value unsafe.Pointer, name string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValue:forPragmaNamed:"), value, objc.String(name))
}

