// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKAttachment] class.
var (
	HKAttachmentClass     _HKAttachmentClass
	HKAttachmentClassOnce sync.Once
)

func getHKAttachmentClass() _HKAttachmentClass {
	HKAttachmentClassOnce.Do(func() {
		HKAttachmentClass = _HKAttachmentClass{objc.GetClass("HKAttachment")}
	})
	return HKAttachmentClass
}

type _HKAttachmentClass struct {
	class objc.Class
}

// An interface definition for the [HKAttachment] class.
type IHKAttachment interface {
	objectivec.IObject
}

// A file that is attached to a sample in the HealthKit store.
//
// To access the attachment’s data, get a data reader from the attachment store for each attachment. You can then asynchronously access the whole data object. Alternatively, you can access the file’s contents as an asynchronous sequence of bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachment
type HKAttachment struct {
	objectivec.Object
}

// HKAttachmentFrom constructs a [HKAttachment] from an unsafe.Pointer.
//
// A file that is attached to a sample in the HealthKit store.
func HKAttachmentFrom(ptr unsafe.Pointer) HKAttachment {
	return HKAttachment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKAttachmentClass) Alloc() HKAttachment {
	rv := objc.Send[HKAttachment](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKAttachmentClass) New() HKAttachment {
	rv := objc.Send[HKAttachment](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAttachment) Init() HKAttachment {
	rv := objc.Send[HKAttachment](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAttachment) Autorelease() HKAttachment {
	rv := objc.Send[HKAttachment](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAttachment creates a new HKAttachment instance.
func NewHKAttachment() HKAttachment {
	return getHKAttachmentClass().New()
}




