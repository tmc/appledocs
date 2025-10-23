// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MergeConflict] class.
var (
	MergeConflictClass     _MergeConflictClass
	MergeConflictClassOnce sync.Once
)

func getMergeConflictClass() _MergeConflictClass {
	MergeConflictClassOnce.Do(func() {
		MergeConflictClass = _MergeConflictClass{objc.GetClass("NSMergeConflict")}
	})
	return MergeConflictClass
}

type _MergeConflictClass struct {
	class objc.Class
}

// An interface definition for the [MergeConflict] class.
type IMergeConflict interface {
	objectivec.IObject
	OldVersionNumber() uint
	CachedSnapshot() string
	SetCachedSnapshot(value string)
	NewVersionNumber() int
	SetNewVersionNumber(value int)
	ObjectSnapshot() string
	SetObjectSnapshot(value string)
	PersistedSnapshot() string
	SetPersistedSnapshot(value string)
	SourceObject() IManagedObject
	SetSourceObject(value IManagedObject)
}

// An encapsulation of conflicts that occur during an attempt to save changes in a managed object context.
//
// A conflict can occur in two situations: Between the managed object context and its in-memory cached state at the persistent store coordinator layer. Between the cached state at the persistent store coordinator layer and the external store (file, database, and so forth). In this case, the merge conflict has a cached snapshot and a persisted snapshot. The source object is also provided as a convenience, but it is not directly involved in the conflict. Snapshot dictionaries include values for all attributes and to-one relationships, but not to-many relationships. Relationship values are references. To-many relationships must be pulled from the persistent store as needed.


// An encapsulation of conflicts that occur during an attempt to save changes in a managed object context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMergeConflict
type MergeConflict struct {
	objectivec.Object
}

// MergeConflictFrom constructs a [MergeConflict] from an unsafe.Pointer.
//
// An encapsulation of conflicts that occur during an attempt to save changes in a managed object context.
func MergeConflictFrom(ptr unsafe.Pointer) MergeConflict {
	return MergeConflict{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MergeConflictClass) Alloc() MergeConflict {
	rv := objc.Send[MergeConflict](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MergeConflictClass) New() MergeConflict {
	rv := objc.Send[MergeConflict](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MergeConflict) Init() MergeConflict {
	rv := objc.Send[MergeConflict](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MergeConflict) Autorelease() MergeConflict {
	rv := objc.Send[MergeConflict](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMergeConflict creates a new MergeConflict instance.
func NewMergeConflict() MergeConflict {
	return getMergeConflictClass().New()
}



// The old version number for the change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMergeConflict/oldVersionNumber
func (m_ MergeConflict) OldVersionNumber() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("oldVersionNumber"))
	return rv
}


// A dictionary containing the values of the source object held in the persistent store coordinator layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/cachedsnapshot
func (m_ MergeConflict) CachedSnapshot() string {
	rv := objc.Send[string](m_.ID, objc.Sel("cachedSnapshot"))
	return rv
}


// A dictionary containing the values of the source object held in the persistent store coordinator layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/cachedsnapshot
func (m_ MergeConflict) SetCachedSnapshot(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCachedSnapshot:"), objc.String(value))
}


// The new version number for the change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/newversionnumber
func (m_ MergeConflict) NewVersionNumber() int {
	rv := objc.Send[int](m_.ID, objc.Sel("newVersionNumber"))
	return rv
}


// The new version number for the change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/newversionnumber
func (m_ MergeConflict) SetNewVersionNumber(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewVersionNumber:"), value)
}


// A dictionary containing the values of the source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/objectsnapshot
func (m_ MergeConflict) ObjectSnapshot() string {
	rv := objc.Send[string](m_.ID, objc.Sel("objectSnapshot"))
	return rv
}


// A dictionary containing the values of the source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/objectsnapshot
func (m_ MergeConflict) SetObjectSnapshot(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectSnapshot:"), objc.String(value))
}


// A dictionary containing the values of the source object held in the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/persistedsnapshot
func (m_ MergeConflict) PersistedSnapshot() string {
	rv := objc.Send[string](m_.ID, objc.Sel("persistedSnapshot"))
	return rv
}


// A dictionary containing the values of the source object held in the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/persistedsnapshot
func (m_ MergeConflict) SetPersistedSnapshot(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPersistedSnapshot:"), objc.String(value))
}


// The source object for the conflict.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/sourceobject
func (m_ MergeConflict) SourceObject() IManagedObject {
	rv := objc.Send[ManagedObject](m_.ID, objc.Sel("sourceObject"))
	return rv
}


// The source object for the conflict.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmergeconflict/sourceobject
func (m_ MergeConflict) SetSourceObject(value IManagedObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceObject:"), value)
}



