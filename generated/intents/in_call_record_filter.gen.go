// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [INCallRecordFilter] class.
var (
	INCallRecordFilterClass     _INCallRecordFilterClass
	INCallRecordFilterClassOnce sync.Once
)

func getINCallRecordFilterClass() _INCallRecordFilterClass {
	INCallRecordFilterClassOnce.Do(func() {
		INCallRecordFilterClass = _INCallRecordFilterClass{objc.GetClass("INCallRecordFilter")}
	})
	return INCallRecordFilterClass
}

type _INCallRecordFilterClass struct {
	class objc.Class
}

// An interface definition for the [INCallRecordFilter] class.
type IINCallRecordFilter interface {
	objectivec.IObject
}

// Filters a user specifies to redial a call.
//
// Use this method to create filters contributed by the user to redial a call. The object identifies the person, type of call, and ability to make the call to initiate the user’s request.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallRecordFilter
type INCallRecordFilter struct {
	objectivec.Object
}

// INCallRecordFilterFrom constructs a [INCallRecordFilter] from an unsafe.Pointer.
//
// Filters a user specifies to redial a call.
func INCallRecordFilterFrom(ptr unsafe.Pointer) INCallRecordFilter {
	return INCallRecordFilter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INCallRecordFilterClass) Alloc() INCallRecordFilter {
	rv := objc.Send[INCallRecordFilter](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INCallRecordFilterClass) New() INCallRecordFilter {
	rv := objc.Send[INCallRecordFilter](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INCallRecordFilter) Init() INCallRecordFilter {
	rv := objc.Send[INCallRecordFilter](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INCallRecordFilter) Autorelease() INCallRecordFilter {
	rv := objc.Send[INCallRecordFilter](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINCallRecordFilter creates a new INCallRecordFilter instance.
func NewINCallRecordFilter() INCallRecordFilter {
	return getINCallRecordFilterClass().New()
}




