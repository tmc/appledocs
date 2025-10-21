// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AXDataSeriesDescriptor] class.
var (
	AXDataSeriesDescriptorClass     _AXDataSeriesDescriptorClass
	AXDataSeriesDescriptorClassOnce sync.Once
)

func getAXDataSeriesDescriptorClass() _AXDataSeriesDescriptorClass {
	AXDataSeriesDescriptorClassOnce.Do(func() {
		AXDataSeriesDescriptorClass = _AXDataSeriesDescriptorClass{objc.GetClass("AXDataSeriesDescriptor")}
	})
	return AXDataSeriesDescriptorClass
}

type _AXDataSeriesDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AXDataSeriesDescriptor] class.
type IAXDataSeriesDescriptor interface {
	objectivec.IObject
}

// An object that represents a series of data points.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor
type AXDataSeriesDescriptor struct {
	objectivec.Object
}

// AXDataSeriesDescriptorFrom constructs a [AXDataSeriesDescriptor] from an unsafe.Pointer.
//
// An object that represents a series of data points.
func AXDataSeriesDescriptorFrom(ptr unsafe.Pointer) AXDataSeriesDescriptor {
	return AXDataSeriesDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AXDataSeriesDescriptorClass) Alloc() AXDataSeriesDescriptor {
	rv := objc.Send[AXDataSeriesDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXDataSeriesDescriptorClass) New() AXDataSeriesDescriptor {
	rv := objc.Send[AXDataSeriesDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXDataSeriesDescriptor) Init() AXDataSeriesDescriptor {
	rv := objc.Send[AXDataSeriesDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXDataSeriesDescriptor) Autorelease() AXDataSeriesDescriptor {
	rv := objc.Send[AXDataSeriesDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXDataSeriesDescriptor creates a new AXDataSeriesDescriptor instance.
func NewAXDataSeriesDescriptor() AXDataSeriesDescriptor {
	return getAXDataSeriesDescriptorClass().New()
}


// Creates a data series with the specified attributed name, a Boolean value that indicates whether the series is continuous, and data points.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/init(attributedName:isContinuous:dataPoints:)
func NewAXDataSeriesDescriptorWithAttributedNameIsContinuousDataPoints(attributedName unsafe.Pointer, isContinuous bool, dataPoints unsafe.Pointer) AXDataSeriesDescriptor {
	instance := getAXDataSeriesDescriptorClass().Alloc()
	rv := objc.Send[AXDataSeriesDescriptor](instance.ID, objc.Sel("initWithAttributedName:isContinuous:dataPoints:"), attributedName, isContinuous, dataPoints)
	rv.Autorelease()
	return rv
}

// Creates a data series with the specified name, a Boolean value that indicates whether the series is continuous, and data points.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/init(name:isContinuous:dataPoints:)
func NewAXDataSeriesDescriptorWithNameIsContinuousDataPoints(name string, isContinuous bool, dataPoints unsafe.Pointer) AXDataSeriesDescriptor {
	instance := getAXDataSeriesDescriptorClass().Alloc()
	rv := objc.Send[AXDataSeriesDescriptor](instance.ID, objc.Sel("initWithName:isContinuous:dataPoints:"), objc.String(name), isContinuous, dataPoints)
	rv.Autorelease()
	return rv
}


// An attributed version of the data series name.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/attributedName
func (a_ AXDataSeriesDescriptor) AttributedName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("attributedName"))
	return rv
}


// SetAttributedName sets the value of the attributedName property.
// An attributed version of the data series name.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/attributedName
func (a_ AXDataSeriesDescriptor) SetAttributedName(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributedName:"), value)
}
// The data points that the series contains.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/dataPoints
func (a_ AXDataSeriesDescriptor) DataPoints() []AXDataPoint {
	rv := objc.Send[[]AXDataPoint](a_.ID, objc.Sel("dataPoints"))
	return rv
}


// SetDataPoints sets the value of the dataPoints property.
// The data points that the series contains.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/dataPoints
func (a_ AXDataSeriesDescriptor) SetDataPoints(value []AXDataPoint) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setDataPoints:"), nsArray)
}
// A Boolean value that determines whether the data series is continuous.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/isContinuous
func (a_ AXDataSeriesDescriptor) IsContinuous() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isContinuous"))
	return rv
}


// SetIsContinuous sets the value of the isContinuous property.
// A Boolean value that determines whether the data series is continuous.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/isContinuous
func (a_ AXDataSeriesDescriptor) SetIsContinuous(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsContinuous:"), value)
}
// The name of the data series.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/name
func (a_ AXDataSeriesDescriptor) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the data series.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/name
func (a_ AXDataSeriesDescriptor) SetName(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setName:"), value)
}

