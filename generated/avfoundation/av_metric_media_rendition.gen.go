// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MetricMediaRendition] class.
var (
	MetricMediaRenditionClass     _MetricMediaRenditionClass
	MetricMediaRenditionClassOnce sync.Once
)

func getMetricMediaRenditionClass() _MetricMediaRenditionClass {
	MetricMediaRenditionClassOnce.Do(func() {
		MetricMediaRenditionClass = _MetricMediaRenditionClass{objc.GetClass("AVMetricMediaRendition")}
	})
	return MetricMediaRenditionClass
}

type _MetricMediaRenditionClass struct {
	class objc.Class
}





// An interface definition for the [MetricMediaRendition] class.
type IMetricMediaRendition interface {
	objectivec.IObject
	

	// properties:
	StableID() foundation.foundation.INSString
	URL() foundation.foundation.INSURL


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetricMediaRenditionClass) Alloc() MetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricMediaRenditionClass) New() MetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricMediaRendition) Init() MetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricMediaRendition) Autorelease() MetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricMediaRendition creates a new MetricMediaRendition instance.
func NewMetricMediaRendition() MetricMediaRendition {
	return getMetricMediaRenditionClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaRendition
type MetricMediaRendition struct {
	objectivec.Object
}

// MetricMediaRenditionFrom constructs a [MetricMediaRendition] from an unsafe.Pointer.
func MetricMediaRenditionFrom(ptr unsafe.Pointer) MetricMediaRendition {
	return MetricMediaRendition{objectivec.Object{objc.ID(ptr)}}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaRendition/stableID
func (m_ MetricMediaRendition) StableID() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("stableID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaRendition/url
func (m_ MetricMediaRendition) URL() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("URL"))
	return rv
}








