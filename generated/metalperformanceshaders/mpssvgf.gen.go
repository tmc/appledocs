// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SVGF] class.
var (
	SVGFClass     _SVGFClass
	SVGFClassOnce sync.Once
)

func getSVGFClass() _SVGFClass {
	SVGFClassOnce.Do(func() {
		SVGFClass = _SVGFClass{objc.GetClass("MPSSVGF")}
	})
	return SVGFClass
}

type _SVGFClass struct {
	class objc.Class
}

// An interface definition for the [SVGF] class.
type ISVGF interface {
	IKernel
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF
type SVGF struct {
	Kernel
}

// SVGFFrom constructs a [SVGF] from an unsafe.Pointer.
func SVGFFrom(ptr unsafe.Pointer) SVGF {
	return SVGF{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SVGFClass) Alloc() SVGF {
	rv := objc.Send[SVGF](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SVGFClass) New() SVGF {
	rv := objc.Send[SVGF](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SVGF) Init() SVGF {
	rv := objc.Send[SVGF](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SVGF) Autorelease() SVGF {
	rv := objc.Send[SVGF](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSVGF creates a new SVGF instance.
func NewSVGF() SVGF {
	return getSVGFClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/variancePrefilterRadius
func (s_ SVGF) VariancePrefilterRadius() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("variancePrefilterRadius"))
	return rv
}


// SetVariancePrefilterRadius sets the value of the variancePrefilterRadius property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/variancePrefilterRadius
func (s_ SVGF) SetVariancePrefilterRadius(value uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVariancePrefilterRadius:"), value)
}


