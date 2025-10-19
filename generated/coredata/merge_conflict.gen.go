// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MergeConflict] class.
var mergeConflictClass = _MergeConflictClass{objc.GetClass("NSMergeConflict")}

type _MergeConflictClass struct {
	class objc.Class
}

// An encapsulation of conflicts that occur during an attempt to save changes in a managed object context. [Full Topic]
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



