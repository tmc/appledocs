// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterAirQuality] class.
var (
	MTRClusterAirQualityClass     _MTRClusterAirQualityClass
	MTRClusterAirQualityClassOnce sync.Once
)

func getMTRClusterAirQualityClass() _MTRClusterAirQualityClass {
	MTRClusterAirQualityClassOnce.Do(func() {
		MTRClusterAirQualityClass = _MTRClusterAirQualityClass{objc.GetClass("MTRClusterAirQuality")}
	})
	return MTRClusterAirQualityClass
}

type _MTRClusterAirQualityClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterAirQuality] class.
type IMTRClusterAirQuality interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterAirQuality
type MTRClusterAirQuality struct {
	MTRGenericCluster
}

// MTRClusterAirQualityFrom constructs a [MTRClusterAirQuality] from an unsafe.Pointer.
func MTRClusterAirQualityFrom(ptr unsafe.Pointer) MTRClusterAirQuality {
	return MTRClusterAirQuality{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterAirQualityClass) Alloc() MTRClusterAirQuality {
	rv := objc.Send[MTRClusterAirQuality](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterAirQualityClass) New() MTRClusterAirQuality {
	rv := objc.Send[MTRClusterAirQuality](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterAirQuality) Init() MTRClusterAirQuality {
	rv := objc.Send[MTRClusterAirQuality](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterAirQuality) Autorelease() MTRClusterAirQuality {
	rv := objc.Send[MTRClusterAirQuality](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterAirQuality creates a new MTRClusterAirQuality instance.
func NewMTRClusterAirQuality() MTRClusterAirQuality {
	return getMTRClusterAirQualityClass().New()
}
