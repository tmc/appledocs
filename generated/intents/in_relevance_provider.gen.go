// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INRelevanceProvider] class.
var (
	INRelevanceProviderClass     _INRelevanceProviderClass
	INRelevanceProviderClassOnce sync.Once
)

func getINRelevanceProviderClass() _INRelevanceProviderClass {
	INRelevanceProviderClassOnce.Do(func() {
		INRelevanceProviderClass = _INRelevanceProviderClass{objc.GetClass("INRelevanceProvider")}
	})
	return INRelevanceProviderClass
}

type _INRelevanceProviderClass struct {
	class objc.Class
}

// An interface definition for the [INRelevanceProvider] class.
type IINRelevanceProvider interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other Intents classes.

// A parent class referenced by other Intents classes. [Full Topic]
type INRelevanceProvider struct {
	objectivec.Object
}

// INRelevanceProviderFrom constructs a [INRelevanceProvider] from an unsafe.Pointer.
//
// A parent class referenced by other Intents classes.
func INRelevanceProviderFrom(ptr unsafe.Pointer) INRelevanceProvider {
	return INRelevanceProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INRelevanceProviderClass) Alloc() INRelevanceProvider {
	rv := objc.Send[INRelevanceProvider](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INRelevanceProviderClass) New() INRelevanceProvider {
	rv := objc.Send[INRelevanceProvider](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INRelevanceProvider) Init() INRelevanceProvider {
	rv := objc.Send[INRelevanceProvider](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INRelevanceProvider) Autorelease() INRelevanceProvider {
	rv := objc.Send[INRelevanceProvider](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINRelevanceProvider creates a new INRelevanceProvider instance.
func NewINRelevanceProvider() INRelevanceProvider {
	return getINRelevanceProviderClass().New()
}
