// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EDRMetadata] class.
var (
	eDRMetadataClass     _EDRMetadataClass
	eDRMetadataClassOnce sync.Once
)

func getEDRMetadataClass() _EDRMetadataClass {
	eDRMetadataClassOnce.Do(func() {
		eDRMetadataClass = _EDRMetadataClass{objc.GetClass("CAEDRMetadata")}
	})
	return eDRMetadataClass
}

type _EDRMetadataClass struct {
	class objc.Class
}

// An interface definition for the [EDRMetadata] class.
type IEDRMetadata interface {
	objectivec.IObject
}

// Metadata describing how extended dynamic range (EDR) values should be tone mapped.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata
type EDRMetadata struct {
	objectivec.Object
}

// EDRMetadataFrom constructs a [EDRMetadata] from an unsafe.Pointer.
//
// Metadata describing how extended dynamic range (EDR) values should be tone mapped.
func EDRMetadataFrom(ptr unsafe.Pointer) EDRMetadata {
	return EDRMetadata{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EDRMetadataClass) Alloc() EDRMetadata {
	rv := objc.Send[EDRMetadata](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EDRMetadataClass) New() EDRMetadata {
	rv := objc.Send[EDRMetadata](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EDRMetadata) Init() EDRMetadata {
	rv := objc.Send[EDRMetadata](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EDRMetadata) Autorelease() EDRMetadata {
	rv := objc.Send[EDRMetadata](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEDRMetadata creates a new EDRMetadata instance.
func NewEDRMetadata() EDRMetadata {
	return getEDRMetadataClass().New()
}




