// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [INCallRecord] class.
var (
	INCallRecordClass     _INCallRecordClass
	INCallRecordClassOnce sync.Once
)

func getINCallRecordClass() _INCallRecordClass {
	INCallRecordClassOnce.Do(func() {
		INCallRecordClass = _INCallRecordClass{objc.GetClass("INCallRecord")}
	})
	return INCallRecordClass
}

type _INCallRecordClass struct {
	class objc.Class
}

// An interface definition for the [INCallRecord] class.
type IINCallRecord interface {
	objectivec.IObject
}

// The details about a call handled by your app.
//
// An object stores details about calls made by the user through your app. You use call record objects to communicate basic information about calls to SiriKit. A call record identifies the type of call, the duration of the call, the date and time of the call, and the person on the other end of the call. You create call record objects when reporting search results back to SiriKit and when identifying voicemails to play.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallRecord
type INCallRecord struct {
	objectivec.Object
}

// INCallRecordFrom constructs a [INCallRecord] from an unsafe.Pointer.
//
// The details about a call handled by your app.
func INCallRecordFrom(ptr unsafe.Pointer) INCallRecord {
	return INCallRecord{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INCallRecordClass) Alloc() INCallRecord {
	rv := objc.Send[INCallRecord](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INCallRecordClass) New() INCallRecord {
	rv := objc.Send[INCallRecord](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INCallRecord) Init() INCallRecord {
	rv := objc.Send[INCallRecord](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INCallRecord) Autorelease() INCallRecord {
	rv := objc.Send[INCallRecord](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINCallRecord creates a new INCallRecord instance.
func NewINCallRecord() INCallRecord {
	return getINCallRecordClass().New()
}




