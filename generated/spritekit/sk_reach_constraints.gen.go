// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKReachConstraints] class.
var (
	sKReachConstraintsClass     _SKReachConstraintsClass
	sKReachConstraintsClassOnce sync.Once
)

func getSKReachConstraintsClass() _SKReachConstraintsClass {
	sKReachConstraintsClassOnce.Do(func() {
		sKReachConstraintsClass = _SKReachConstraintsClass{objc.GetClass("SKReachConstraints")}
	})
	return sKReachConstraintsClass
}

type _SKReachConstraintsClass struct {
	class objc.Class
}

// An interface definition for the [SKReachConstraints] class.
type ISKReachConstraints interface {
	objectivec.IObject
}

// A specification of the degree of freedom when solving inverse kinematics.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKReachConstraints
type SKReachConstraints struct {
	objectivec.Object
}

// SKReachConstraintsFrom constructs a [SKReachConstraints] from an unsafe.Pointer.
//
// A specification of the degree of freedom when solving inverse kinematics.
func SKReachConstraintsFrom(ptr unsafe.Pointer) SKReachConstraints {
	return SKReachConstraints{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SKReachConstraintsClass) Alloc() SKReachConstraints {
	rv := objc.Send[SKReachConstraints](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKReachConstraintsClass) New() SKReachConstraints {
	rv := objc.Send[SKReachConstraints](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKReachConstraints) Init() SKReachConstraints {
	rv := objc.Send[SKReachConstraints](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKReachConstraints) Autorelease() SKReachConstraints {
	rv := objc.Send[SKReachConstraints](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKReachConstraints creates a new SKReachConstraints instance.
func NewSKReachConstraints() SKReachConstraints {
	return getSKReachConstraintsClass().New()
}


// Initializes a new reach constraint object.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKReachConstraints/init(lowerAngleLimit:upperAngleLimit:)
func NewSKReachConstraintsWithLowerAngleLimitUpperAngleLimit(lowerAngleLimit float64, upperAngleLimit float64) SKReachConstraints {
	instance := getSKReachConstraintsClass().Alloc()
	rv := objc.Send[SKReachConstraints](instance.ID, objc.Sel("initWithLowerAngleLimit:upperAngleLimit:"), lowerAngleLimit, upperAngleLimit)
	rv.Autorelease()
	return rv
}



