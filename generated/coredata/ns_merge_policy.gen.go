// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MergePolicy] class.
var (
	MergePolicyClass     _MergePolicyClass
	MergePolicyClassOnce sync.Once
)

func getMergePolicyClass() _MergePolicyClass {
	MergePolicyClassOnce.Do(func() {
		MergePolicyClass = _MergePolicyClass{objc.GetClass("NSMergePolicy")}
	})
	return MergePolicyClass
}

type _MergePolicyClass struct {
	class objc.Class
}

// An interface definition for the [MergePolicy] class.
type IMergePolicy interface {
	objectivec.IObject
	// properties:
	MergeType() MergePolicyType /* not a class type */
	// methods:
}

// A policy object that you use to resolve conflicts between the persistent store and in-memory versions of managed objects.
//
// A conflict is a mismatch between state held at two different layers in the Core Data stack. A conflict can arise when you save a managed object context and you have stale data at another layer. There are two places in which a conflict may occur: Between the managed object context layer and its in-memory cached state at the persistent store coordinator layer. Between the cached state at the persistent store coordinator and the external store (file, database, and so forth). Conflicts are represented by instances of .

// A policy object that you use to resolve conflicts between the persistent store and in-memory versions of managed objects.
//
// [Full Topic]
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

// Alloc allocates a new instance without initialization.
func (mc _MergePolicyClass) Alloc() MergePolicy {
	rv := objc.Send[MergePolicy](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MergePolicyClass) New() MergePolicy {
	rv := objc.Send[MergePolicy](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MergePolicy) Init() MergePolicy {
	rv := objc.Send[MergePolicy](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MergePolicy) Autorelease() MergePolicy {
	rv := objc.Send[MergePolicy](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMergePolicy creates a new MergePolicy instance.
func NewMergePolicy() MergePolicy {
	return getMergePolicyClass().New()
}

// The merge type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMergePolicy/mergeType
func (m_ MergePolicy) MergeType() MergePolicyType /* not a class type */ {
	rv := objc.Send[MergePolicyType](m_.ID, objc.Sel("mergeType"))
	return rv
}
