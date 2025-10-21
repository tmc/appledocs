// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PropertyAnimation] class.
var (
	PropertyAnimationClass     _PropertyAnimationClass
	PropertyAnimationClassOnce sync.Once
)

func getPropertyAnimationClass() _PropertyAnimationClass {
	PropertyAnimationClassOnce.Do(func() {
		PropertyAnimationClass = _PropertyAnimationClass{objc.GetClass("CAPropertyAnimation")}
	})
	return PropertyAnimationClass
}

type _PropertyAnimationClass struct {
	class objc.Class
}

// An interface definition for the [PropertyAnimation] class.
type IPropertyAnimation interface {
	IAnimation
}

// An abstract subclass for creating animations that manipulate the value of layer properties.
//
// The property to animate is specified using a key path that is relative to the layer using the animation. You do not create instances of : to animate the properties of a Core Animation layer, create instance of the concrete subclasses or .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getPropertyAnimationClass().New()
}




// Creates and returns an instance for the specified key path.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/init(keyPath:)
func NewPropertyAnimationWithKeyPath(path string) PropertyAnimation {
	rv := objc.Send[PropertyAnimation](objc.ID(getPropertyAnimationClass().class), objc.Sel("animationWithKeyPath:"), objc.String(path))
	return rv
}


// Creates and returns an instance for the specified key path.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/init(keyPath:)
func (pc _PropertyAnimationClass) AnimationWithKeyPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("animationWithKeyPath:"), objc.String(path))
	return rv
}

// Determines if the value specified by the animation is added to the current render tree value to produce the new render tree value.
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/isadditive
func (p_ PropertyAnimation) IsAdditive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isAdditive"))
	return rv
}


// SetIsAdditive sets the value of the isAdditive property.
// Determines if the value specified by the animation is added to the current render tree value to produce the new render tree value.

//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/isadditive
func (p_ PropertyAnimation) SetIsAdditive(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsAdditive:"), value)
}

// Determines if the value of the property is the value at the end of the previous repeat cycle, plus the value of the current repeat cycle.
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/iscumulative
func (p_ PropertyAnimation) IsCumulative() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isCumulative"))
	return rv
}


// SetIsCumulative sets the value of the isCumulative property.
// Determines if the value of the property is the value at the end of the previous repeat cycle, plus the value of the current repeat cycle.

//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/iscumulative
func (p_ PropertyAnimation) SetIsCumulative(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsCumulative:"), value)
}

// Determines if the value specified by the animation is added to the current render tree value to produce the new render tree value.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/isAdditive
func (p_ PropertyAnimation) Additive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("additive"))
	return rv
}


// SetAdditive sets the value of the additive property.
// Determines if the value specified by the animation is added to the current render tree value to produce the new render tree value.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/isAdditive
func (p_ PropertyAnimation) SetAdditive(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAdditive:"), value)
}

// Determines if the value of the property is the value at the end of the previous repeat cycle, plus the value of the current repeat cycle.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/isCumulative
func (p_ PropertyAnimation) Cumulative() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("cumulative"))
	return rv
}


// SetCumulative sets the value of the cumulative property.
// Determines if the value of the property is the value at the end of the previous repeat cycle, plus the value of the current repeat cycle.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/isCumulative
func (p_ PropertyAnimation) SetCumulative(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCumulative:"), value)
}

// Specifies the key path the receiver animates.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/keyPath
func (p_ PropertyAnimation) KeyPath() string {
	rv := objc.Send[string](p_.ID, objc.Sel("keyPath"))
	return rv
}


// SetKeyPath sets the value of the keyPath property.
// Specifies the key path the receiver animates.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/keyPath
func (p_ PropertyAnimation) SetKeyPath(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setKeyPath:"), objc.String(value))
}

// An optional value function that is applied to interpolated values.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/valueFunction
func (p_ PropertyAnimation) ValueFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("valueFunction"))
	return rv
}


// SetValueFunction sets the value of the valueFunction property.
// An optional value function that is applied to interpolated values.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/valueFunction
func (p_ PropertyAnimation) SetValueFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValueFunction:"), value)
}



