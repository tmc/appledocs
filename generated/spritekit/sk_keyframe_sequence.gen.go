// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKKeyframeSequence] class.
var sKKeyframeSequenceClass = _SKKeyframeSequenceClass{objc.GetClass("SKKeyframeSequence")}

type _SKKeyframeSequenceClass struct {
	class objc.Class
}

// An interface definition for the [SKKeyframeSequence] class.
type ISKKeyframeSequence interface {
	objectivec.IObject
	AddKeyframeValueTime(value objc.ID, time float64)
	Count() uint
	GetKeyframeTimeForIndex(index uint) float64
	GetKeyframeValueForIndex(index uint) objc.ID
	RemoveKeyframeAtIndex(index uint)
	RemoveLastKeyframe()
	SampleAtTime(time float64) objc.ID
	SetKeyframeTimeForIndex(time float64, index uint)
	SetKeyframeValueForIndex(value objc.ID, index uint)
	SetKeyframeValueTimeForIndex(value objc.ID, time float64, index uint)
}

// An object that performs interpolation between values specified at different times (keyframes). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence

type SKKeyframeSequence struct {
	objectivec.Object
}

// SKKeyframeSequenceFrom constructs a [SKKeyframeSequence] from an unsafe.Pointer.
//
// An object that performs interpolation between values specified at different times (keyframes).
func SKKeyframeSequenceFrom(ptr unsafe.Pointer) SKKeyframeSequence {
	return SKKeyframeSequence{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (sc _SKKeyframeSequenceClass) Alloc() SKKeyframeSequence {
	rv := objc.Send[SKKeyframeSequence](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKKeyframeSequenceClass) New() SKKeyframeSequence {
	rv := objc.Send[SKKeyframeSequence](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKKeyframeSequence) Init() SKKeyframeSequence {
	rv := objc.Send[SKKeyframeSequence](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKKeyframeSequence) Autorelease() SKKeyframeSequence {
	rv := objc.Send[SKKeyframeSequence](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKKeyframeSequence creates a new SKKeyframeSequence instance.
func NewSKKeyframeSequence() SKKeyframeSequence {
	return sKKeyframeSequenceClass.New()
}


// Initializes a new keyframe sequence. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/init(capacity:)
func NewSKKeyframeSequenceWithCapacity(numItems uint) SKKeyframeSequence {
	instance := sKKeyframeSequenceClass.Alloc()
	rv := objc.Send[SKKeyframeSequence](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/init(coder:)
func NewSKKeyframeSequenceWithCoder(aDecoder unsafe.Pointer) SKKeyframeSequence {
	instance := sKKeyframeSequenceClass.Alloc()
	rv := objc.Send[SKKeyframeSequence](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}
// Initializes a keyframe sequence with an initial set of values and times. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/init(keyframeValues:times:)
func NewSKKeyframeSequenceWithKeyframeValuesTimes(values unsafe.Pointer, times unsafe.Pointer) SKKeyframeSequence {
	instance := sKKeyframeSequenceClass.Alloc()
	rv := objc.Send[SKKeyframeSequence](instance.ID, objc.Sel("initWithKeyframeValues:times:"), values, times)
	rv.Autorelease()
	return rv
}


// Adds a keyframe to the sequence. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/addKeyframeValue(_:time:)
func (s_ SKKeyframeSequence) AddKeyframeValueTime(value objc.ID, time float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addKeyframeValue:time:"), value, time)
}
// The number of keyframes in the sequence. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/count()
func (s_ SKKeyframeSequence) Count() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("count"))
	return rv
}
// Gets the time for a keyframe in the sequence. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/getKeyframeTime(for:)
func (s_ SKKeyframeSequence) GetKeyframeTimeForIndex(index uint) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("getKeyframeTimeForIndex:"), index)
	return rv
}
// Gets the value for a keyframe in the sequence. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/getKeyframeValue(for:)
func (s_ SKKeyframeSequence) GetKeyframeValueForIndex(index uint) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("getKeyframeValueForIndex:"), index)
	return rv
}
// Removes a keyframe from the sequence. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/removeKeyframe(at:)
func (s_ SKKeyframeSequence) RemoveKeyframeAtIndex(index uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeKeyframeAtIndex:"), index)
}
// Removes the last value in the sequence. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/removeLastKeyframe()
func (s_ SKKeyframeSequence) RemoveLastKeyframe() {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeLastKeyframe"))
}
// Calculates the sample at a particular time. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/sample(atTime:)
func (s_ SKKeyframeSequence) SampleAtTime(time float64) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("sampleAtTime:"), time)
	return rv
}
// Changes the time for a specific keyframe. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/setKeyframeTime(_:for:)
func (s_ SKKeyframeSequence) SetKeyframeTimeForIndex(time float64, index uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKeyframeTime:forIndex:"), time, index)
}
// Changes the value for a specific keyframe. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/setKeyframeValue(_:for:)
func (s_ SKKeyframeSequence) SetKeyframeValueForIndex(value objc.ID, index uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKeyframeValue:forIndex:"), value, index)
}
// Replaces a keyframe in the sequence with a new keyframe. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/setKeyframeValue(_:time:for:)
func (s_ SKKeyframeSequence) SetKeyframeValueTimeForIndex(value objc.ID, time float64, index uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKeyframeValue:time:forIndex:"), value, time, index)
}

