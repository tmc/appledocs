// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PurgeableData] class.
var (
	purgeableDataClass     _PurgeableDataClass
	purgeableDataClassOnce sync.Once
)

func getPurgeableDataClass() _PurgeableDataClass {
	purgeableDataClassOnce.Do(func() {
		purgeableDataClass = _PurgeableDataClass{objc.GetClass("NSPurgeableData")}
	})
	return purgeableDataClass
}

type _PurgeableDataClass struct {
	class objc.Class
}

// An interface definition for the [PurgeableData] class.
type IPurgeableData interface {
	IMutableData
}

// A mutable data object containing bytes that can be discarded when they’re no longer needed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPurgeableData
type PurgeableData struct {
	MutableData
}

// PurgeableDataFrom constructs a [PurgeableData] from an unsafe.Pointer.
//
// A mutable data object containing bytes that can be discarded when they’re no longer needed.
func PurgeableDataFrom(ptr unsafe.Pointer) PurgeableData {
	return PurgeableData{
		MutableData: MutableDataFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PurgeableDataClass) Alloc() PurgeableData {
	rv := objc.Send[PurgeableData](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PurgeableDataClass) New() PurgeableData {
	rv := objc.Send[PurgeableData](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PurgeableData) Init() PurgeableData {
	rv := objc.Send[PurgeableData](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PurgeableData) Autorelease() PurgeableData {
	rv := objc.Send[PurgeableData](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPurgeableData creates a new PurgeableData instance.
func NewPurgeableData() PurgeableData {
	return getPurgeableDataClass().New()
}




