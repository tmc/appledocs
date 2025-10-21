// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRAccessControlClusterAccessControlEntry] class.
var (
	MTRAccessControlClusterAccessControlEntryClass     _MTRAccessControlClusterAccessControlEntryClass
	MTRAccessControlClusterAccessControlEntryClassOnce sync.Once
)

func getMTRAccessControlClusterAccessControlEntryClass() _MTRAccessControlClusterAccessControlEntryClass {
	MTRAccessControlClusterAccessControlEntryClassOnce.Do(func() {
		MTRAccessControlClusterAccessControlEntryClass = _MTRAccessControlClusterAccessControlEntryClass{objc.GetClass("MTRAccessControlClusterAccessControlEntry")}
	})
	return MTRAccessControlClusterAccessControlEntryClass
}

type _MTRAccessControlClusterAccessControlEntryClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterAccessControlEntry] class.
type IMTRAccessControlClusterAccessControlEntry interface {
	IMTRAccessControlClusterAccessControlEntryStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessControlEntry
type MTRAccessControlClusterAccessControlEntry struct {
	MTRAccessControlClusterAccessControlEntryStruct
}

// MTRAccessControlClusterAccessControlEntryFrom constructs a [MTRAccessControlClusterAccessControlEntry] from an unsafe.Pointer.
func MTRAccessControlClusterAccessControlEntryFrom(ptr unsafe.Pointer) MTRAccessControlClusterAccessControlEntry {
	return MTRAccessControlClusterAccessControlEntry{
		MTRAccessControlClusterAccessControlEntryStruct: MTRAccessControlClusterAccessControlEntryStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterAccessControlEntryClass) Alloc() MTRAccessControlClusterAccessControlEntry {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntry](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterAccessControlEntryClass) New() MTRAccessControlClusterAccessControlEntry {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntry](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterAccessControlEntry) Init() MTRAccessControlClusterAccessControlEntry {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntry](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterAccessControlEntry) Autorelease() MTRAccessControlClusterAccessControlEntry {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntry](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterAccessControlEntry creates a new MTRAccessControlClusterAccessControlEntry instance.
func NewMTRAccessControlClusterAccessControlEntry() MTRAccessControlClusterAccessControlEntry {
	return getMTRAccessControlClusterAccessControlEntryClass().New()
}




