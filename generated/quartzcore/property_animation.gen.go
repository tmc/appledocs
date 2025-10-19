// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PropertyAnimation] class.
var propertyAnimationClass = _PropertyAnimationClass{objc.GetClass("CAPropertyAnimation")}

type _PropertyAnimationClass struct {
	class objc.Class
}

// An interface definition for the [PropertyAnimation] class.
type IPropertyAnimation interface {
	IAnimation
}

// An abstract subclass for creating animations that manipulate the value of layer properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation

type PropertyAnimation struct {
	Animation
}

// PropertyAnimationFrom constructs a [PropertyAnimation] from an unsafe.Pointer.
//
// An abstract subclass for creating animations that manipulate the value of layer properties.
func PropertyAnimationFrom(ptr unsafe.Pointer) PropertyAnimation {
	return PropertyAnimation{
		Animation: AnimationFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (pc _PropertyAnimationClass) Alloc() PropertyAnimation {
	rv := objc.Send[PropertyAnimation](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PropertyAnimationClass) New() PropertyAnimation {
	rv := objc.Send[PropertyAnimation](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PropertyAnimation) Init() PropertyAnimation {
	rv := objc.Send[PropertyAnimation](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PropertyAnimation) Autorelease() PropertyAnimation {
	rv := objc.Send[PropertyAnimation](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPropertyAnimation creates a new PropertyAnimation instance.
func NewPropertyAnimation() PropertyAnimation {
	return propertyAnimationClass.New()
}


// Creates and returns an instance for the specified key path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/init(keyPath:)
func NewAnimationWithKeyPath(path string) PropertyAnimation {
	rv := objc.Send[PropertyAnimation](objc.ID(propertyAnimationClass.class), objc.Sel("animationWithKeyPath:"), path)
	rv.Autorelease()
	return rv
}


// Creates and returns an instance for the specified key path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/init(keyPath:)
func (pc _PropertyAnimationClass) AnimationWithKeyPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("animationWithKeyPath:"), path)
	return rv
}


