// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterAirQuality] class.
var (
	MTRBaseClusterAirQualityClass     _MTRBaseClusterAirQualityClass
	MTRBaseClusterAirQualityClassOnce sync.Once
)

func getMTRBaseClusterAirQualityClass() _MTRBaseClusterAirQualityClass {
	MTRBaseClusterAirQualityClassOnce.Do(func() {
		MTRBaseClusterAirQualityClass = _MTRBaseClusterAirQualityClass{objc.GetClass("MTRBaseClusterAirQuality")}
	})
	return MTRBaseClusterAirQualityClass
}

type _MTRBaseClusterAirQualityClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterAirQuality] class.
type IMTRBaseClusterAirQuality interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterAirQuality
type MTRBaseClusterAirQuality struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterAirQualityFrom constructs a [MTRBaseClusterAirQuality] from an unsafe.Pointer.
func MTRBaseClusterAirQualityFrom(ptr unsafe.Pointer) MTRBaseClusterAirQuality {
	return MTRBaseClusterAirQuality{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterAirQualityClass) Alloc() MTRBaseClusterAirQuality {
	rv := objc.Send[MTRBaseClusterAirQuality](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterAirQualityClass) New() MTRBaseClusterAirQuality {
	rv := objc.Send[MTRBaseClusterAirQuality](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterAirQuality) Init() MTRBaseClusterAirQuality {
	rv := objc.Send[MTRBaseClusterAirQuality](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterAirQuality) Autorelease() MTRBaseClusterAirQuality {
	rv := objc.Send[MTRBaseClusterAirQuality](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterAirQuality creates a new MTRBaseClusterAirQuality instance.
func NewMTRBaseClusterAirQuality() MTRBaseClusterAirQuality {
	return getMTRBaseClusterAirQualityClass().New()
}
