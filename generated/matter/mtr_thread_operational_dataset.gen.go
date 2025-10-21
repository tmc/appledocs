// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThreadOperationalDataset] class.
var (
	MTRThreadOperationalDatasetClass     _MTRThreadOperationalDatasetClass
	MTRThreadOperationalDatasetClassOnce sync.Once
)

func getMTRThreadOperationalDatasetClass() _MTRThreadOperationalDatasetClass {
	MTRThreadOperationalDatasetClassOnce.Do(func() {
		MTRThreadOperationalDatasetClass = _MTRThreadOperationalDatasetClass{objc.GetClass("MTRThreadOperationalDataset")}
	})
	return MTRThreadOperationalDatasetClass
}

type _MTRThreadOperationalDatasetClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadOperationalDataset] class.
type IMTRThreadOperationalDataset interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset
type MTRThreadOperationalDataset struct {
	objectivec.Object
}

// MTRThreadOperationalDatasetFrom constructs a [MTRThreadOperationalDataset] from an unsafe.Pointer.
func MTRThreadOperationalDatasetFrom(ptr unsafe.Pointer) MTRThreadOperationalDataset {
	return MTRThreadOperationalDataset{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadOperationalDatasetClass) Alloc() MTRThreadOperationalDataset {
	rv := objc.Send[MTRThreadOperationalDataset](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadOperationalDatasetClass) New() MTRThreadOperationalDataset {
	rv := objc.Send[MTRThreadOperationalDataset](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadOperationalDataset) Init() MTRThreadOperationalDataset {
	rv := objc.Send[MTRThreadOperationalDataset](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadOperationalDataset) Autorelease() MTRThreadOperationalDataset {
	rv := objc.Send[MTRThreadOperationalDataset](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadOperationalDataset creates a new MTRThreadOperationalDataset instance.
func NewMTRThreadOperationalDataset() MTRThreadOperationalDataset {
	return getMTRThreadOperationalDatasetClass().New()
}




