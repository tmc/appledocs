// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BasicAnimation] class.
var basicAnimationClass = _BasicAnimationClass{objc.GetClass("CABasicAnimation")}

type _BasicAnimationClass struct {
	class objc.Class
}

// An interface definition for the [BasicAnimation] class.
type IBasicAnimation interface {
	IPropertyAnimation
}

// An object that provides basic, single-keyframe animation capabilities for a layer property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation

type BasicAnimation struct {
	PropertyAnimation
}

// BasicAnimationFrom constructs a [BasicAnimation] from an unsafe.Pointer.
//
// An object that provides basic, single-keyframe animation capabilities for a layer property.
func BasicAnimationFrom(ptr unsafe.Pointer) BasicAnimation {
	return BasicAnimation{
		PropertyAnimation: PropertyAnimationFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (bc _BasicAnimationClass) Alloc() BasicAnimation {
	rv := objc.Send[BasicAnimation](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (bc _BasicAnimationClass) New() BasicAnimation {
	rv := objc.Send[BasicAnimation](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BasicAnimation) Init() BasicAnimation {
	rv := objc.Send[BasicAnimation](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BasicAnimation) Autorelease() BasicAnimation {
	rv := objc.Send[BasicAnimation](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBasicAnimation creates a new BasicAnimation instance.
func NewBasicAnimation() BasicAnimation {
	return basicAnimationClass.New()
}




