// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INMediaSearch] class.
var (
	INMediaSearchClass     _INMediaSearchClass
	INMediaSearchClassOnce sync.Once
)

func getINMediaSearchClass() _INMediaSearchClass {
	INMediaSearchClassOnce.Do(func() {
		INMediaSearchClass = _INMediaSearchClass{objc.GetClass("INMediaSearch")}
	})
	return INMediaSearchClass
}

type _INMediaSearchClass struct {
	class objc.Class
}

// An interface definition for the [INMediaSearch] class.
type IINMediaSearch interface {
	objectivec.IObject
}

// An object that describes a media type to search for, such as a station name, song name, or album name.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INMediaSearch
type INMediaSearch struct {
	objectivec.Object
}

// INMediaSearchFrom constructs a [INMediaSearch] from an unsafe.Pointer.
//
// An object that describes a media type to search for, such as a station name, song name, or album name.
func INMediaSearchFrom(ptr unsafe.Pointer) INMediaSearch {
	return INMediaSearch{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INMediaSearchClass) Alloc() INMediaSearch {
	rv := objc.Send[INMediaSearch](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INMediaSearchClass) New() INMediaSearch {
	rv := objc.Send[INMediaSearch](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INMediaSearch) Init() INMediaSearch {
	rv := objc.Send[INMediaSearch](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INMediaSearch) Autorelease() INMediaSearch {
	rv := objc.Send[INMediaSearch](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINMediaSearch creates a new INMediaSearch instance.
func NewINMediaSearch() INMediaSearch {
	return getINMediaSearchClass().New()
}




