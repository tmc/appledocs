// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INRelevantShortcutStore] class.
var (
	INRelevantShortcutStoreClass     _INRelevantShortcutStoreClass
	INRelevantShortcutStoreClassOnce sync.Once
)

func getINRelevantShortcutStoreClass() _INRelevantShortcutStoreClass {
	INRelevantShortcutStoreClassOnce.Do(func() {
		INRelevantShortcutStoreClass = _INRelevantShortcutStoreClass{objc.GetClass("INRelevantShortcutStore")}
	})
	return INRelevantShortcutStoreClass
}

type _INRelevantShortcutStoreClass struct {
	class objc.Class
}

// An interface definition for the [INRelevantShortcutStore] class.
type IINRelevantShortcutStore interface {
	objectivec.IObject
}

// An object that saves relevant shortcuts.

// An object that saves relevant shortcuts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRelevantShortcutStore
type INRelevantShortcutStore struct {
	objectivec.Object
}

// INRelevantShortcutStoreFrom constructs a [INRelevantShortcutStore] from an unsafe.Pointer.
//
// An object that saves relevant shortcuts.
func INRelevantShortcutStoreFrom(ptr unsafe.Pointer) INRelevantShortcutStore {
	return INRelevantShortcutStore{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INRelevantShortcutStoreClass) Alloc() INRelevantShortcutStore {
	rv := objc.Send[INRelevantShortcutStore](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INRelevantShortcutStoreClass) New() INRelevantShortcutStore {
	rv := objc.Send[INRelevantShortcutStore](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INRelevantShortcutStore) Init() INRelevantShortcutStore {
	rv := objc.Send[INRelevantShortcutStore](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INRelevantShortcutStore) Autorelease() INRelevantShortcutStore {
	rv := objc.Send[INRelevantShortcutStore](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINRelevantShortcutStore creates a new INRelevantShortcutStore instance.
func NewINRelevantShortcutStore() INRelevantShortcutStore {
	return getINRelevantShortcutStoreClass().New()
}

// The default relevant shortcut store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRelevantShortcutStore/default
func (ic _INRelevantShortcutStoreClass) DefaultStore() INRelevantShortcutStore {
	rv := objc.Send[INRelevantShortcutStore](objc.ID(ic.class), objc.Sel("defaultStore"))
	return rv
}

// The default relevant shortcut store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRelevantShortcutStore/default
func (i_ INRelevantShortcutStore) DefaultStore() INRelevantShortcutStore {
	rv := objc.Send[INRelevantShortcutStore](i_.ID, objc.Sel("defaultStore"))
	return rv
}
