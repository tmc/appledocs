// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INIntentDonationMetadata] class.
var (
	INIntentDonationMetadataClass     _INIntentDonationMetadataClass
	INIntentDonationMetadataClassOnce sync.Once
)

func getINIntentDonationMetadataClass() _INIntentDonationMetadataClass {
	INIntentDonationMetadataClassOnce.Do(func() {
		INIntentDonationMetadataClass = _INIntentDonationMetadataClass{objc.GetClass("INIntentDonationMetadata")}
	})
	return INIntentDonationMetadataClass
}

type _INIntentDonationMetadataClass struct {
	class objc.Class
}

// An interface definition for the [INIntentDonationMetadata] class.
type IINIntentDonationMetadata interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INIntentDonationMetadata
type INIntentDonationMetadata struct {
	objectivec.Object
}

// INIntentDonationMetadataFrom constructs a [INIntentDonationMetadata] from an unsafe.Pointer.
func INIntentDonationMetadataFrom(ptr unsafe.Pointer) INIntentDonationMetadata {
	return INIntentDonationMetadata{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INIntentDonationMetadataClass) Alloc() INIntentDonationMetadata {
	rv := objc.Send[INIntentDonationMetadata](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INIntentDonationMetadataClass) New() INIntentDonationMetadata {
	rv := objc.Send[INIntentDonationMetadata](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INIntentDonationMetadata) Init() INIntentDonationMetadata {
	rv := objc.Send[INIntentDonationMetadata](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INIntentDonationMetadata) Autorelease() INIntentDonationMetadata {
	rv := objc.Send[INIntentDonationMetadata](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINIntentDonationMetadata creates a new INIntentDonationMetadata instance.
func NewINIntentDonationMetadata() INIntentDonationMetadata {
	return getINIntentDonationMetadataClass().New()
}




