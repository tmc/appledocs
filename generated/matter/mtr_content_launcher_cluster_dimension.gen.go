// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRContentLauncherClusterDimension] class.
var (
	MTRContentLauncherClusterDimensionClass     _MTRContentLauncherClusterDimensionClass
	MTRContentLauncherClusterDimensionClassOnce sync.Once
)

func getMTRContentLauncherClusterDimensionClass() _MTRContentLauncherClusterDimensionClass {
	MTRContentLauncherClusterDimensionClassOnce.Do(func() {
		MTRContentLauncherClusterDimensionClass = _MTRContentLauncherClusterDimensionClass{objc.GetClass("MTRContentLauncherClusterDimension")}
	})
	return MTRContentLauncherClusterDimensionClass
}

type _MTRContentLauncherClusterDimensionClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterDimension] class.
type IMTRContentLauncherClusterDimension interface {
	IMTRContentLauncherClusterDimensionStruct
	// properties:
	Height() objc.IObject /* cross-framework: NSNumber */
	SetHeight(value objc.IObject /* cross-framework: NSNumber */)
	Metric() objc.IObject /* cross-framework: NSNumber */
	SetMetric(value objc.IObject /* cross-framework: NSNumber */)
	Width() objc.IObject /* cross-framework: NSNumber */
	SetWidth(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimension
type MTRContentLauncherClusterDimension struct {
	MTRContentLauncherClusterDimensionStruct
}

// MTRContentLauncherClusterDimensionFrom constructs a [MTRContentLauncherClusterDimension] from an unsafe.Pointer.
func MTRContentLauncherClusterDimensionFrom(ptr unsafe.Pointer) MTRContentLauncherClusterDimension {
	return MTRContentLauncherClusterDimension{
		MTRContentLauncherClusterDimensionStruct: MTRContentLauncherClusterDimensionStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterDimensionClass) Alloc() MTRContentLauncherClusterDimension {
	rv := objc.Send[MTRContentLauncherClusterDimension](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterDimensionClass) New() MTRContentLauncherClusterDimension {
	rv := objc.Send[MTRContentLauncherClusterDimension](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterDimension) Init() MTRContentLauncherClusterDimension {
	rv := objc.Send[MTRContentLauncherClusterDimension](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterDimension) Autorelease() MTRContentLauncherClusterDimension {
	rv := objc.Send[MTRContentLauncherClusterDimension](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterDimension creates a new MTRContentLauncherClusterDimension instance.
func NewMTRContentLauncherClusterDimension() MTRContentLauncherClusterDimension {
	return getMTRContentLauncherClusterDimensionClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterdimension/height
func (m_ MTRContentLauncherClusterDimension) Height() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("height"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterdimension/height
func (m_ MTRContentLauncherClusterDimension) SetHeight(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterdimension/metric
func (m_ MTRContentLauncherClusterDimension) Metric() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("metric"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterdimension/metric
func (m_ MTRContentLauncherClusterDimension) SetMetric(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetric:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterdimension/width
func (m_ MTRContentLauncherClusterDimension) Width() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("width"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterdimension/width
func (m_ MTRContentLauncherClusterDimension) SetWidth(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWidth:"), value)
}



