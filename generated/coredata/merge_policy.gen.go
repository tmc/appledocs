// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MergePolicy] class.
var mergePolicyClass = _MergePolicyClass{objc.GetClass("NSMergePolicy")}

type _MergePolicyClass struct {
	class objc.Class
}

// A policy object that you use to resolve conflicts between the persistent store and in-memory versions of managed objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMergePolicy

type MergePolicy struct {
	objectivec.Object
}

// MergePolicyFrom constructs a [MergePolicy] from an unsafe.Pointer.
//
// A policy object that you use to resolve conflicts between the persistent store and in-memory versions of managed objects.
func MergePolicyFrom(ptr unsafe.Pointer) MergePolicy {
	return MergePolicy{objectivec.Object{objc.ID(ptr)}}
}



