//go:build darwin && ios

// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Bundle


// Returns the current preservation priority for the specified tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/preservationPriority(forTag:)
func (b_ Bundle) PreservationPriorityForTag(tag IString) float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("preservationPriorityForTag:"), tag)
	return rv
}

// A hint to the system of the relative order for purging tagged sets of resources in the bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/setPreservationPriority(_:forTags:)
func (b_ Bundle) SetPreservationPriorityForTags(priority float64, tags unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPreservationPriority:forTags:"), priority, tags)
}

// iOS-only properties




