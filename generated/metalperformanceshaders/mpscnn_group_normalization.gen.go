// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNNGroupNormalization] class.
var (
	CNNGroupNormalizationClass     _CNNGroupNormalizationClass
	CNNGroupNormalizationClassOnce sync.Once
)

func getCNNGroupNormalizationClass() _CNNGroupNormalizationClass {
	CNNGroupNormalizationClassOnce.Do(func() {
		CNNGroupNormalizationClass = _CNNGroupNormalizationClass{objc.GetClass("MPSCNNGroupNormalization")}
	})
	return CNNGroupNormalizationClass
}

type _CNNGroupNormalizationClass struct {
	class objc.Class
}

// An interface definition for the [CNNGroupNormalization] class.
type ICNNGroupNormalization interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalization
type CNNGroupNormalization struct {
	objectivec.Object
}

// CNNGroupNormalizationFrom constructs a [CNNGroupNormalization] from an unsafe.Pointer.
func CNNGroupNormalizationFrom(ptr unsafe.Pointer) CNNGroupNormalization {
	return CNNGroupNormalization{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNGroupNormalizationClass) Alloc() CNNGroupNormalization {
	rv := objc.Send[CNNGroupNormalization](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNGroupNormalizationClass) New() CNNGroupNormalization {
	rv := objc.Send[CNNGroupNormalization](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNGroupNormalization) Init() CNNGroupNormalization {
	rv := objc.Send[CNNGroupNormalization](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNGroupNormalization) Autorelease() CNNGroupNormalization {
	rv := objc.Send[CNNGroupNormalization](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNGroupNormalization creates a new CNNGroupNormalization instance.
func NewCNNGroupNormalization() CNNGroupNormalization {
	return getCNNGroupNormalizationClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalization/init(device:dataSource:)
func NewCNNGroupNormalizationWithDeviceDataSource(device objc.ID, dataSource objc.ID) CNNGroupNormalization {
	instance := getCNNGroupNormalizationClass().Alloc()
	rv := objc.Send[CNNGroupNormalization](instance.ID, objc.Sel("initWithDevice:dataSource:"), device, dataSource)
	rv.Autorelease()
	return rv
}
