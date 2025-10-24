// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [SampleBufferRequest] class.
var (
	SampleBufferRequestClass     _SampleBufferRequestClass
	SampleBufferRequestClassOnce sync.Once
)

func getSampleBufferRequestClass() _SampleBufferRequestClass {
	SampleBufferRequestClassOnce.Do(func() {
		SampleBufferRequestClass = _SampleBufferRequestClass{objc.GetClass("AVSampleBufferRequest")}
	})
	return SampleBufferRequestClass
}

type _SampleBufferRequestClass struct {
	class objc.Class
}





// An interface definition for the [SampleBufferRequest] class.
type ISampleBufferRequest interface {
	objectivec.IObject
	

	// properties:
	Direction() SampleBufferRequestDirection
	SetDirection(value SampleBufferRequestDirection)
	LimitCursor() IAVSampleCursor
	SetLimitCursor(value IAVSampleCursor)
	MaxSampleCount() int
	SetMaxSampleCount(value int)
	Mode() SampleBufferRequestMode
	SetMode(value SampleBufferRequestMode)
	OverrideTime() objc.IObject /* cross-framework: Time */
	SetOverrideTime(value objc.IObject /* cross-framework: Time */)
	PreferredMinSampleCount() int
	SetPreferredMinSampleCount(value int)
	StartCursor() IAVSampleCursor


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _SampleBufferRequestClass) Alloc() SampleBufferRequest {
	rv := objc.Send[SampleBufferRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SampleBufferRequestClass) New() SampleBufferRequest {
	rv := objc.Send[SampleBufferRequest](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SampleBufferRequest) Init() SampleBufferRequest {
	rv := objc.Send[SampleBufferRequest](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SampleBufferRequest) Autorelease() SampleBufferRequest {
	rv := objc.Send[SampleBufferRequest](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleBufferRequest creates a new SampleBufferRequest instance.
func NewSampleBufferRequest() SampleBufferRequest {
	return getSampleBufferRequestClass().New()
}





// An object that describes a sample buffer creation request.


// An object that describes a sample buffer creation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest
type SampleBufferRequest struct {
	objectivec.Object
}

// SampleBufferRequestFrom constructs a [SampleBufferRequest] from an unsafe.Pointer.
//
// An object that describes a sample buffer creation request.
func SampleBufferRequestFrom(ptr unsafe.Pointer) SampleBufferRequest {
	return SampleBufferRequest{objectivec.Object{objc.ID(ptr)}}
}






// Creates a newly allocated sample buffer request with the specified sample cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/init(start:)
func NewSampleBufferRequestWithStartCursor(startCursor IAVSampleCursor) SampleBufferRequest {
	instance := getSampleBufferRequestClass().Alloc()
	rv := objc.Send[SampleBufferRequest](instance.ID, objc.Sel("initWithStartCursor:"), startCursor)
	rv.Autorelease()
	return rv
}






















// The buffer sample direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/direction-swift.property
func (s_ SampleBufferRequest) Direction() SampleBufferRequestDirection {
	rv := objc.Send[SampleBufferRequestDirection](s_.ID, objc.Sel("direction"))
	return rv
}


// The buffer sample direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/direction-swift.property
func (s_ SampleBufferRequest) SetDirection(value SampleBufferRequestDirection) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDirection:"), value)
}


// The limiting position for sample loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/limitCursor
func (s_ SampleBufferRequest) LimitCursor() IAVSampleCursor {
	rv := objc.Send[SampleCursor](s_.ID, objc.Sel("limitCursor"))
	return rv
}


// The limiting position for sample loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/limitCursor
func (s_ SampleBufferRequest) SetLimitCursor(value IAVSampleCursor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLimitCursor:"), value)
}


// The maximum number of samples to load.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/maxSampleCount
func (s_ SampleBufferRequest) MaxSampleCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("maxSampleCount"))
	return rv
}


// The maximum number of samples to load.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/maxSampleCount
func (s_ SampleBufferRequest) SetMaxSampleCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxSampleCount:"), value)
}


// The sample buffer request mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/mode-swift.property
func (s_ SampleBufferRequest) Mode() SampleBufferRequestMode {
	rv := objc.Send[SampleBufferRequestMode](s_.ID, objc.Sel("mode"))
	return rv
}


// The sample buffer request mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/mode-swift.property
func (s_ SampleBufferRequest) SetMode(value SampleBufferRequestMode) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMode:"), value)
}


// The deadline for sample data and output PTS for the sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/overrideTime
func (s_ SampleBufferRequest) OverrideTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](s_.ID, objc.Sel("overrideTime"))
	return rv
}


// The deadline for sample data and output PTS for the sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/overrideTime
func (s_ SampleBufferRequest) SetOverrideTime(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setOverrideTime:"), value)
}


// The preferred minimum number of samples to load.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/preferredMinSampleCount
func (s_ SampleBufferRequest) PreferredMinSampleCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("preferredMinSampleCount"))
	return rv
}


// The preferred minimum number of samples to load.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/preferredMinSampleCount
func (s_ SampleBufferRequest) SetPreferredMinSampleCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreferredMinSampleCount:"), value)
}


// The starting cursor position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/startCursor
func (s_ SampleBufferRequest) StartCursor() IAVSampleCursor {
	rv := objc.Send[SampleCursor](s_.ID, objc.Sel("startCursor"))
	return rv
}







