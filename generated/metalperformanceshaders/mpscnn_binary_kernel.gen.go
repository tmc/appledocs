// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNNBinaryKernel] class.
var (
	CNNBinaryKernelClass     _CNNBinaryKernelClass
	CNNBinaryKernelClassOnce sync.Once
)

func getCNNBinaryKernelClass() _CNNBinaryKernelClass {
	CNNBinaryKernelClassOnce.Do(func() {
		CNNBinaryKernelClass = _CNNBinaryKernelClass{objc.GetClass("MPSCNNBinaryKernel")}
	})
	return CNNBinaryKernelClass
}

type _CNNBinaryKernelClass struct {
	class objc.Class
}

// An interface definition for the [CNNBinaryKernel] class.
type ICNNBinaryKernel interface {
	IKernel
}

// A convolution neural network kernel.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel
type CNNBinaryKernel struct {
	Kernel
}

// CNNBinaryKernelFrom constructs a [CNNBinaryKernel] from an unsafe.Pointer.
//
// A convolution neural network kernel.
func CNNBinaryKernelFrom(ptr unsafe.Pointer) CNNBinaryKernel {
	return CNNBinaryKernel{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNBinaryKernelClass) Alloc() CNNBinaryKernel {
	rv := objc.Send[CNNBinaryKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNBinaryKernelClass) New() CNNBinaryKernel {
	rv := objc.Send[CNNBinaryKernel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNBinaryKernel) Init() CNNBinaryKernel {
	rv := objc.Send[CNNBinaryKernel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNBinaryKernel) Autorelease() CNNBinaryKernel {
	rv := objc.Send[CNNBinaryKernel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNBinaryKernel creates a new CNNBinaryKernel instance.
func NewCNNBinaryKernel() CNNBinaryKernel {
	return getCNNBinaryKernelClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarydilationratey
func (c_ CNNBinaryKernel) PrimaryDilationRateY() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primaryDilationRateY"))
	return rv
}


// SetPrimaryDilationRateY sets the value of the primaryDilationRateY property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarydilationratey
func (c_ CNNBinaryKernel) SetPrimaryDilationRateY(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryDilationRateY:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondaryoffset
func (c_ CNNBinaryKernel) SecondaryOffset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("secondaryOffset"))
	return rv
}


// SetSecondaryOffset sets the value of the secondaryOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondaryoffset
func (c_ CNNBinaryKernel) SetSecondaryOffset(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarykernelheight
func (c_ CNNBinaryKernel) SecondaryKernelHeight() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondaryKernelHeight"))
	return rv
}


// SetSecondaryKernelHeight sets the value of the secondaryKernelHeight property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarykernelheight
func (c_ CNNBinaryKernel) SetSecondaryKernelHeight(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryKernelHeight:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarysourcefeaturechanneloffset
func (c_ CNNBinaryKernel) SecondarySourceFeatureChannelOffset() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondarySourceFeatureChannelOffset"))
	return rv
}


// SetSecondarySourceFeatureChannelOffset sets the value of the secondarySourceFeatureChannelOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarysourcefeaturechanneloffset
func (c_ CNNBinaryKernel) SetSecondarySourceFeatureChannelOffset(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondarySourceFeatureChannelOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/padding
func (c_ CNNBinaryKernel) Padding() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("padding"))
	return rv
}


// SetPadding sets the value of the padding property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/padding
func (c_ CNNBinaryKernel) SetPadding(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPadding:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarysourcefeaturechannelmaxcount
func (c_ CNNBinaryKernel) PrimarySourceFeatureChannelMaxCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primarySourceFeatureChannelMaxCount"))
	return rv
}


// SetPrimarySourceFeatureChannelMaxCount sets the value of the primarySourceFeatureChannelMaxCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarysourcefeaturechannelmaxcount
func (c_ CNNBinaryKernel) SetPrimarySourceFeatureChannelMaxCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimarySourceFeatureChannelMaxCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarydilationratex
func (c_ CNNBinaryKernel) SecondaryDilationRateX() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondaryDilationRateX"))
	return rv
}


// SetSecondaryDilationRateX sets the value of the secondaryDilationRateX property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarydilationratex
func (c_ CNNBinaryKernel) SetSecondaryDilationRateX(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryDilationRateX:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarystrideinpixelsx
func (c_ CNNBinaryKernel) PrimaryStrideInPixelsX() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primaryStrideInPixelsX"))
	return rv
}


// SetPrimaryStrideInPixelsX sets the value of the primaryStrideInPixelsX property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarystrideinpixelsx
func (c_ CNNBinaryKernel) SetPrimaryStrideInPixelsX(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryStrideInPixelsX:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarystrideinpixelsy
func (c_ CNNBinaryKernel) PrimaryStrideInPixelsY() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primaryStrideInPixelsY"))
	return rv
}


// SetPrimaryStrideInPixelsY sets the value of the primaryStrideInPixelsY property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarystrideinpixelsy
func (c_ CNNBinaryKernel) SetPrimaryStrideInPixelsY(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryStrideInPixelsY:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarysourcefeaturechannelmaxcount
func (c_ CNNBinaryKernel) SecondarySourceFeatureChannelMaxCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondarySourceFeatureChannelMaxCount"))
	return rv
}


// SetSecondarySourceFeatureChannelMaxCount sets the value of the secondarySourceFeatureChannelMaxCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarysourcefeaturechannelmaxcount
func (c_ CNNBinaryKernel) SetSecondarySourceFeatureChannelMaxCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondarySourceFeatureChannelMaxCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primaryoffset
func (c_ CNNBinaryKernel) PrimaryOffset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("primaryOffset"))
	return rv
}


// SetPrimaryOffset sets the value of the primaryOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primaryoffset
func (c_ CNNBinaryKernel) SetPrimaryOffset(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/cliprect
func (c_ CNNBinaryKernel) ClipRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("clipRect"))
	return rv
}


// SetClipRect sets the value of the clipRect property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/cliprect
func (c_ CNNBinaryKernel) SetClipRect(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClipRect:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarykernelwidth
func (c_ CNNBinaryKernel) SecondaryKernelWidth() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondaryKernelWidth"))
	return rv
}


// SetSecondaryKernelWidth sets the value of the secondaryKernelWidth property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarykernelwidth
func (c_ CNNBinaryKernel) SetSecondaryKernelWidth(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryKernelWidth:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarysourcefeaturechanneloffset
func (c_ CNNBinaryKernel) PrimarySourceFeatureChannelOffset() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primarySourceFeatureChannelOffset"))
	return rv
}


// SetPrimarySourceFeatureChannelOffset sets the value of the primarySourceFeatureChannelOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarysourcefeaturechanneloffset
func (c_ CNNBinaryKernel) SetPrimarySourceFeatureChannelOffset(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimarySourceFeatureChannelOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/destinationfeaturechanneloffset
func (c_ CNNBinaryKernel) DestinationFeatureChannelOffset() int {
	rv := objc.Send[int](c_.ID, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}


// SetDestinationFeatureChannelOffset sets the value of the destinationFeatureChannelOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/destinationfeaturechanneloffset
func (c_ CNNBinaryKernel) SetDestinationFeatureChannelOffset(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarykernelheight
func (c_ CNNBinaryKernel) PrimaryKernelHeight() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primaryKernelHeight"))
	return rv
}


// SetPrimaryKernelHeight sets the value of the primaryKernelHeight property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarykernelheight
func (c_ CNNBinaryKernel) SetPrimaryKernelHeight(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryKernelHeight:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarystrideinpixelsy
func (c_ CNNBinaryKernel) SecondaryStrideInPixelsY() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}


// SetSecondaryStrideInPixelsY sets the value of the secondaryStrideInPixelsY property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarystrideinpixelsy
func (c_ CNNBinaryKernel) SetSecondaryStrideInPixelsY(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarykernelwidth
func (c_ CNNBinaryKernel) PrimaryKernelWidth() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primaryKernelWidth"))
	return rv
}


// SetPrimaryKernelWidth sets the value of the primaryKernelWidth property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarykernelwidth
func (c_ CNNBinaryKernel) SetPrimaryKernelWidth(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryKernelWidth:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondaryedgemode
func (c_ CNNBinaryKernel) SecondaryEdgeMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("secondaryEdgeMode"))
	return rv
}


// SetSecondaryEdgeMode sets the value of the secondaryEdgeMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondaryedgemode
func (c_ CNNBinaryKernel) SetSecondaryEdgeMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryEdgeMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/destinationimageallocator
func (c_ CNNBinaryKernel) DestinationImageAllocator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("destinationImageAllocator"))
	return rv
}


// SetDestinationImageAllocator sets the value of the destinationImageAllocator property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/destinationimageallocator
func (c_ CNNBinaryKernel) SetDestinationImageAllocator(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationImageAllocator:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/isbackwards
func (c_ CNNBinaryKernel) IsBackwards() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isBackwards"))
	return rv
}


// SetIsBackwards sets the value of the isBackwards property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/isbackwards
func (c_ CNNBinaryKernel) SetIsBackwards(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsBackwards:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarystrideinpixelsx
func (c_ CNNBinaryKernel) SecondaryStrideInPixelsX() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}


// SetSecondaryStrideInPixelsX sets the value of the secondaryStrideInPixelsX property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarystrideinpixelsx
func (c_ CNNBinaryKernel) SetSecondaryStrideInPixelsX(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/isstatemodified
func (c_ CNNBinaryKernel) IsStateModified() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStateModified"))
	return rv
}


// SetIsStateModified sets the value of the isStateModified property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/isstatemodified
func (c_ CNNBinaryKernel) SetIsStateModified(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStateModified:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primaryedgemode
func (c_ CNNBinaryKernel) PrimaryEdgeMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("primaryEdgeMode"))
	return rv
}


// SetPrimaryEdgeMode sets the value of the primaryEdgeMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primaryedgemode
func (c_ CNNBinaryKernel) SetPrimaryEdgeMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryEdgeMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarydilationratex
func (c_ CNNBinaryKernel) PrimaryDilationRateX() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primaryDilationRateX"))
	return rv
}


// SetPrimaryDilationRateX sets the value of the primaryDilationRateX property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarydilationratex
func (c_ CNNBinaryKernel) SetPrimaryDilationRateX(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryDilationRateX:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarydilationratey
func (c_ CNNBinaryKernel) SecondaryDilationRateY() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondaryDilationRateY"))
	return rv
}


// SetSecondaryDilationRateY sets the value of the secondaryDilationRateY property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarydilationratey
func (c_ CNNBinaryKernel) SetSecondaryDilationRateY(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryDilationRateY:"), value)
}



