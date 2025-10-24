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
	// properties:
	ClipRect() objc.IObject /* cross-framework: MTLRegion */
	SetClipRect(value objc.IObject /* cross-framework: MTLRegion */)
	DestinationFeatureChannelOffset() int
	SetDestinationFeatureChannelOffset(value int)
	DestinationImageAllocator() ImageAllocator /* not a class type */
	SetDestinationImageAllocator(value ImageAllocator /* not a class type */)
	IsBackwards() bool
	SetIsBackwards(value bool)
	IsStateModified() bool
	SetIsStateModified(value bool)
	Padding() Padding /* not a class type */
	SetPadding(value Padding /* not a class type */)
	PrimaryDilationRateX() int
	SetPrimaryDilationRateX(value int)
	PrimaryDilationRateY() int
	SetPrimaryDilationRateY(value int)
	PrimaryEdgeMode() ImageEdgeMode
	SetPrimaryEdgeMode(value ImageEdgeMode)
	PrimaryKernelHeight() int
	SetPrimaryKernelHeight(value int)
	PrimaryKernelWidth() int
	SetPrimaryKernelWidth(value int)
	PrimaryOffset() MPSOffset /* not a class type */
	SetPrimaryOffset(value MPSOffset /* not a class type */)
	PrimarySourceFeatureChannelMaxCount() int
	SetPrimarySourceFeatureChannelMaxCount(value int)
	PrimarySourceFeatureChannelOffset() int
	SetPrimarySourceFeatureChannelOffset(value int)
	PrimaryStrideInPixelsX() int
	SetPrimaryStrideInPixelsX(value int)
	PrimaryStrideInPixelsY() int
	SetPrimaryStrideInPixelsY(value int)
	SecondaryDilationRateX() int
	SetSecondaryDilationRateX(value int)
	SecondaryDilationRateY() int
	SetSecondaryDilationRateY(value int)
	SecondaryEdgeMode() ImageEdgeMode
	SetSecondaryEdgeMode(value ImageEdgeMode)
	SecondaryKernelHeight() int
	SetSecondaryKernelHeight(value int)
	SecondaryKernelWidth() int
	SetSecondaryKernelWidth(value int)
	SecondaryOffset() MPSOffset /* not a class type */
	SetSecondaryOffset(value MPSOffset /* not a class type */)
	SecondarySourceFeatureChannelMaxCount() int
	SetSecondarySourceFeatureChannelMaxCount(value int)
	SecondarySourceFeatureChannelOffset() int
	SetSecondarySourceFeatureChannelOffset(value int)
	SecondaryStrideInPixelsX() int
	SetSecondaryStrideInPixelsX(value int)
	SecondaryStrideInPixelsY() int
	SetSecondaryStrideInPixelsY(value int)
	// methods:
}

// A convolution neural network kernel.


// A convolution neural network kernel.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/cliprect
func (c_ CNNBinaryKernel) ClipRect() objc.IObject /* cross-framework: MTLRegion */ {
	rv := objc.Send[Region](c_.ID, objc.Sel("clipRect"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/cliprect
func (c_ CNNBinaryKernel) SetClipRect(value objc.IObject /* cross-framework: MTLRegion */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClipRect:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/destinationfeaturechanneloffset
func (c_ CNNBinaryKernel) DestinationFeatureChannelOffset() int {
	rv := objc.Send[int](c_.ID, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/destinationfeaturechanneloffset
func (c_ CNNBinaryKernel) SetDestinationFeatureChannelOffset(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/destinationimageallocator
func (c_ CNNBinaryKernel) DestinationImageAllocator() ImageAllocator /* not a class type */ {
	rv := objc.Send[ImageAllocator](c_.ID, objc.Sel("destinationImageAllocator"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/destinationimageallocator
func (c_ CNNBinaryKernel) SetDestinationImageAllocator(value ImageAllocator /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDestinationImageAllocator:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/isbackwards
func (c_ CNNBinaryKernel) IsBackwards() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isBackwards"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/isbackwards
func (c_ CNNBinaryKernel) SetIsBackwards(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsBackwards:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/isstatemodified
func (c_ CNNBinaryKernel) IsStateModified() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStateModified"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/isstatemodified
func (c_ CNNBinaryKernel) SetIsStateModified(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStateModified:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/padding
func (c_ CNNBinaryKernel) Padding() Padding /* not a class type */ {
	rv := objc.Send[Padding](c_.ID, objc.Sel("padding"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/padding
func (c_ CNNBinaryKernel) SetPadding(value Padding /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPadding:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarydilationratex
func (c_ CNNBinaryKernel) PrimaryDilationRateX() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primaryDilationRateX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarydilationratex
func (c_ CNNBinaryKernel) SetPrimaryDilationRateX(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryDilationRateX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarydilationratey
func (c_ CNNBinaryKernel) PrimaryDilationRateY() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primaryDilationRateY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarydilationratey
func (c_ CNNBinaryKernel) SetPrimaryDilationRateY(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryDilationRateY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primaryedgemode
func (c_ CNNBinaryKernel) PrimaryEdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](c_.ID, objc.Sel("primaryEdgeMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primaryedgemode
func (c_ CNNBinaryKernel) SetPrimaryEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryEdgeMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarykernelheight
func (c_ CNNBinaryKernel) PrimaryKernelHeight() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primaryKernelHeight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarykernelheight
func (c_ CNNBinaryKernel) SetPrimaryKernelHeight(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryKernelHeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarykernelwidth
func (c_ CNNBinaryKernel) PrimaryKernelWidth() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primaryKernelWidth"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarykernelwidth
func (c_ CNNBinaryKernel) SetPrimaryKernelWidth(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryKernelWidth:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primaryoffset
func (c_ CNNBinaryKernel) PrimaryOffset() MPSOffset /* not a class type */ {
	rv := objc.Send[Offset](c_.ID, objc.Sel("primaryOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primaryoffset
func (c_ CNNBinaryKernel) SetPrimaryOffset(value MPSOffset /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarysourcefeaturechannelmaxcount
func (c_ CNNBinaryKernel) PrimarySourceFeatureChannelMaxCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primarySourceFeatureChannelMaxCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarysourcefeaturechannelmaxcount
func (c_ CNNBinaryKernel) SetPrimarySourceFeatureChannelMaxCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimarySourceFeatureChannelMaxCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarysourcefeaturechanneloffset
func (c_ CNNBinaryKernel) PrimarySourceFeatureChannelOffset() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primarySourceFeatureChannelOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarysourcefeaturechanneloffset
func (c_ CNNBinaryKernel) SetPrimarySourceFeatureChannelOffset(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimarySourceFeatureChannelOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarystrideinpixelsx
func (c_ CNNBinaryKernel) PrimaryStrideInPixelsX() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primaryStrideInPixelsX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarystrideinpixelsx
func (c_ CNNBinaryKernel) SetPrimaryStrideInPixelsX(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryStrideInPixelsX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarystrideinpixelsy
func (c_ CNNBinaryKernel) PrimaryStrideInPixelsY() int {
	rv := objc.Send[int](c_.ID, objc.Sel("primaryStrideInPixelsY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/primarystrideinpixelsy
func (c_ CNNBinaryKernel) SetPrimaryStrideInPixelsY(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryStrideInPixelsY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarydilationratex
func (c_ CNNBinaryKernel) SecondaryDilationRateX() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondaryDilationRateX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarydilationratex
func (c_ CNNBinaryKernel) SetSecondaryDilationRateX(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryDilationRateX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarydilationratey
func (c_ CNNBinaryKernel) SecondaryDilationRateY() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondaryDilationRateY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarydilationratey
func (c_ CNNBinaryKernel) SetSecondaryDilationRateY(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryDilationRateY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondaryedgemode
func (c_ CNNBinaryKernel) SecondaryEdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](c_.ID, objc.Sel("secondaryEdgeMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondaryedgemode
func (c_ CNNBinaryKernel) SetSecondaryEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryEdgeMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarykernelheight
func (c_ CNNBinaryKernel) SecondaryKernelHeight() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondaryKernelHeight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarykernelheight
func (c_ CNNBinaryKernel) SetSecondaryKernelHeight(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryKernelHeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarykernelwidth
func (c_ CNNBinaryKernel) SecondaryKernelWidth() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondaryKernelWidth"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarykernelwidth
func (c_ CNNBinaryKernel) SetSecondaryKernelWidth(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryKernelWidth:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondaryoffset
func (c_ CNNBinaryKernel) SecondaryOffset() MPSOffset /* not a class type */ {
	rv := objc.Send[Offset](c_.ID, objc.Sel("secondaryOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondaryoffset
func (c_ CNNBinaryKernel) SetSecondaryOffset(value MPSOffset /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarysourcefeaturechannelmaxcount
func (c_ CNNBinaryKernel) SecondarySourceFeatureChannelMaxCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondarySourceFeatureChannelMaxCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarysourcefeaturechannelmaxcount
func (c_ CNNBinaryKernel) SetSecondarySourceFeatureChannelMaxCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondarySourceFeatureChannelMaxCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarysourcefeaturechanneloffset
func (c_ CNNBinaryKernel) SecondarySourceFeatureChannelOffset() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondarySourceFeatureChannelOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarysourcefeaturechanneloffset
func (c_ CNNBinaryKernel) SetSecondarySourceFeatureChannelOffset(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondarySourceFeatureChannelOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarystrideinpixelsx
func (c_ CNNBinaryKernel) SecondaryStrideInPixelsX() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarystrideinpixelsx
func (c_ CNNBinaryKernel) SetSecondaryStrideInPixelsX(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarystrideinpixelsy
func (c_ CNNBinaryKernel) SecondaryStrideInPixelsY() int {
	rv := objc.Send[int](c_.ID, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/secondarystrideinpixelsy
func (c_ CNNBinaryKernel) SetSecondaryStrideInPixelsY(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}



