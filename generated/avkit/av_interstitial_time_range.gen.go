// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [InterstitialTimeRange] class.
var (
	InterstitialTimeRangeClass     _InterstitialTimeRangeClass
	InterstitialTimeRangeClassOnce sync.Once
)

func getInterstitialTimeRangeClass() _InterstitialTimeRangeClass {
	InterstitialTimeRangeClassOnce.Do(func() {
		InterstitialTimeRangeClass = _InterstitialTimeRangeClass{objc.GetClass("AVInterstitialTimeRange")}
	})
	return InterstitialTimeRangeClass
}

type _InterstitialTimeRangeClass struct {
	class objc.Class
}

// An interface definition for the [InterstitialTimeRange] class.
type IInterstitialTimeRange interface {
	objectivec.IObject
	TimeRange() unsafe.Pointer
}

// A time range in an audiovisual presentation for content with an interstitial designation, such as advertisements or legal notices.
//
// When you associate interstitial time ranges with an you present with an , you can customize or restrict the presentation of interstitial content. For example, you can allow the user to skip advertisements or prohibit skipping of a legal notice.


// A time range in an audiovisual presentation for content with an interstitial designation, such as advertisements or legal notices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInterstitialTimeRange

type InterstitialTimeRange struct {
	objectivec.Object
}

// InterstitialTimeRangeFrom constructs a [InterstitialTimeRange] from an unsafe.Pointer.
//
// A time range in an audiovisual presentation for content with an interstitial designation, such as advertisements or legal notices.
func InterstitialTimeRangeFrom(ptr unsafe.Pointer) InterstitialTimeRange {
	return InterstitialTimeRange{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _InterstitialTimeRangeClass) Alloc() InterstitialTimeRange {
	rv := objc.Send[InterstitialTimeRange](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InterstitialTimeRangeClass) New() InterstitialTimeRange {
	rv := objc.Send[InterstitialTimeRange](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InterstitialTimeRange) Init() InterstitialTimeRange {
	rv := objc.Send[InterstitialTimeRange](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InterstitialTimeRange) Autorelease() InterstitialTimeRange {
	rv := objc.Send[InterstitialTimeRange](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInterstitialTimeRange creates a new InterstitialTimeRange instance.
func NewInterstitialTimeRange() InterstitialTimeRange {
	return getInterstitialTimeRangeClass().New()
}




// Initializes an interstitial time range object with the specified time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInterstitialTimeRange/init(timeRange:)

func NewInterstitialTimeRangeWithTimeRange(timeRange unsafe.Pointer) InterstitialTimeRange {
	instance := getInterstitialTimeRangeClass().Alloc()
	rv := objc.Send[InterstitialTimeRange](instance.ID, objc.Sel("initWithTimeRange:"), timeRange)
	rv.Autorelease()
	return rv
}



// The time range identified as interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInterstitialTimeRange/timeRange

func (i_ InterstitialTimeRange) TimeRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("timeRange"))
	return rv
}


