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
	mergeConflictClass     _MergeConflictClass
	mergeConflictClassOnce sync.Once
)

func getMergeConflictClass() _MergeConflictClass {
	mergeConflictClassOnce.Do(func() {
		mergeConflictClass = _MergeConflictClass{objc.GetClass("NSMergeConflict")}
	})
	return mergeConflictClass
}

type _MergeConflictClass struct {
	class objc.Class
}

// An interface definition for the [MergeConflict] class.
type IMergeConflict interface {
	objectivec.IObject
}

// An encapsulation of conflicts that occur during an attempt to save changes in a managed object context.
//
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




