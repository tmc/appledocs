// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LayoutConstraint] class.
var (
	LayoutConstraintClass     _LayoutConstraintClass
	LayoutConstraintClassOnce sync.Once
)

func getLayoutConstraintClass() _LayoutConstraintClass {
	LayoutConstraintClassOnce.Do(func() {
		LayoutConstraintClass = _LayoutConstraintClass{objc.GetClass("NSLayoutConstraint")}
	})
	return LayoutConstraintClass
}

type _LayoutConstraintClass struct {
	class objc.Class
}

// An interface definition for the [LayoutConstraint] class.
type ILayoutConstraint interface {
	objectivec.IObject
}

// The relationship between two user interface objects that must be satisfied by the constraint-based layout system.
//
// Each constraint is a linear equation with the following format: In this equation, and are the variables that Auto Layout can adjust when solving these constraints. The other values are defined when you create the constraint. For example, If you’re defining the relative position of two buttons, you might say “the leading edge of the second button should be 8 points after the trailing edge of the first button.” The linear equation for this relationship is shown below: Auto Layout then modifies the values of the specified leading and trailing edges until both sides of the equation are equal. Note that Auto Layout does not simply assign the value of the right side of this equation to the left side. Instead, the system can modify either attribute or both attributes as needed to solve for this constraint. The fact that constraints are equations (and not assignment operators) means that you can switch the order of the items in the equation as needed to more clearly express the desired relationship. However, if you switch the order, you must also invert the multiplier and constant. For example, the following two equations produce identical constraints: A valid layout is defined as a set constraints with one and only one possible solution. Valid layouts are also referred to as a nonambiguous, nonconflicting layouts. Constraints with more than one solution are ambiguous. Constraints with no valid solutions are conflicting. For more information on resolving ambiguous and conflicting constraints, see in . Additionally, constraints are not limited to equality relationships. They can also use greater than or equal to (>=) or less than or equal to (<=) to describe the relationship between the two attributes. Constraints also have priorities between 1 and 1,000. Constraints with a priority of 1,000 are required. All priorities less than 1,000 are optional. By default, all constraints are required (priority = 1,000). After solving for the required constraints, Auto Layout tries to solve all the optional constraints in priority order from highest to lowest. If it cannot solve for an optional constraint, it tries to come as close as possible to the desired result, and then moves on to the next constraint. This combination of inequalities, equalities, and priorities gives you a great amount of flexibility and power. By combining multiple constraints, you can define layouts that dynamically adapt as the size and location of the elements in your user interface change. For some example layouts, see in .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint
type LayoutConstraint struct {
	objectivec.Object
}

// LayoutConstraintFrom constructs a [LayoutConstraint] from an unsafe.Pointer.
//
// The relationship between two user interface objects that must be satisfied by the constraint-based layout system.
func LayoutConstraintFrom(ptr unsafe.Pointer) LayoutConstraint {
	return LayoutConstraint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LayoutConstraintClass) Alloc() LayoutConstraint {
	rv := objc.Send[LayoutConstraint](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LayoutConstraintClass) New() LayoutConstraint {
	rv := objc.Send[LayoutConstraint](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LayoutConstraint) Init() LayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LayoutConstraint) Autorelease() LayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayoutConstraint creates a new LayoutConstraint instance.
func NewLayoutConstraint() LayoutConstraint {
	return getLayoutConstraintClass().New()
}


// The second object participating in the constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/secondItem
func (l_ LayoutConstraint) SecondItem() objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("secondItem"))
	return rv
}



