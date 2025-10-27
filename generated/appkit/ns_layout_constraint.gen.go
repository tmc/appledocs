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
	

	// properties:
	Constant() float64
	SetConstant(value float64)
	FirstAnchor() ILayoutAnchor
	FirstAttribute() LayoutAttribute
	FirstItem() objc.ID
	Identifier() foundation.foundation.INSString
	SetIdentifier(value foundation.foundation.INSString)
	Active() bool
	SetActive(value bool)
	Multiplier() float64
	Priority() LayoutPriority
	SetPriority(value LayoutPriority)
	Relation() LayoutRelation
	SecondAnchor() ILayoutAnchor
	SecondAttribute() LayoutAttribute
	SecondItem() objc.ID
	ShouldBeArchived() bool
	SetShouldBeArchived(value bool)
	IsActive() bool
	SetIsActive(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (lc _LayoutConstraintClass) Alloc() LayoutConstraint {
	rv := objc.Send[LayoutConstraint](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// The relationship between two user interface objects that must be satisfied by the constraint-based layout system.
//
// Each constraint is a linear equation with the following format: In this equation, and are the variables that Auto Layout can adjust when solving these constraints. The other values are defined when you create the constraint. For example, If you’re defining the relative position of two buttons, you might say “the leading edge of the second button should be 8 points after the trailing edge of the first button.” The linear equation for this relationship is shown below: Auto Layout then modifies the values of the specified leading and trailing edges until both sides of the equation are equal. Note that Auto Layout does not simply assign the value of the right side of this equation to the left side. Instead, the system can modify either attribute or both attributes as needed to solve for this constraint. The fact that constraints are equations (and not assignment operators) means that you can switch the order of the items in the equation as needed to more clearly express the desired relationship. However, if you switch the order, you must also invert the multiplier and constant. For example, the following two equations produce identical constraints: A valid layout is defined as a set constraints with one and only one possible solution. Valid layouts are also referred to as a nonambiguous, nonconflicting layouts. Constraints with more than one solution are ambiguous. Constraints with no valid solutions are conflicting. For more information on resolving ambiguous and conflicting constraints, see in . Additionally, constraints are not limited to equality relationships. They can also use greater than or equal to (>=) or less than or equal to (<=) to describe the relationship between the two attributes. Constraints also have priorities between 1 and 1,000. Constraints with a priority of 1,000 are required. All priorities less than 1,000 are optional. By default, all constraints are required (priority = 1,000). After solving for the required constraints, Auto Layout tries to solve all the optional constraints in priority order from highest to lowest. If it cannot solve for an optional constraint, it tries to come as close as possible to the desired result, and then moves on to the next constraint. This combination of inequalities, equalities, and priorities gives you a great amount of flexibility and power. By combining multiple constraints, you can define layouts that dynamically adapt as the size and location of the elements in your user interface change. For some example layouts, see in .


// The relationship between two user interface objects that must be satisfied by the constraint-based layout system.
//
// [Full Topic]
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






// Creates a constraint that defines the relationship between the specified attributes of the given views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/init(item:attribute:relatedBy:toItem:attribute:multiplier:constant:)
func NewLayoutConstraintWithItemAttributeRelatedByToItemAttributeMultiplierConstant(view1 objectivec.IObject, attr1 LayoutAttribute, relation LayoutRelation, view2 objectivec.IObject, attr2 LayoutAttribute, multiplier float64, c float64) LayoutConstraint {
	rv := objc.Send[LayoutConstraint](objc.ID(getLayoutConstraintClass().class), objc.Sel("constraintWithItem:attribute:relatedBy:toItem:attribute:multiplier:constant:"), view1, attr1, relation, view2, attr2, multiplier, c)
	return rv
}







// Activates each constraint in the specified array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/activate(_:)
func (lc _LayoutConstraintClass) ActivateConstraints(constraints []LayoutConstraint) {
	objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("activateConstraints:"), constraints)
}


// Creates constraints described by an ASCII art-like visual format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/constraints(withVisualFormat:options:metrics:views:)
func (lc _LayoutConstraintClass) ConstraintsWithVisualFormatOptionsMetricsViews(format foundation.foundation.INSString, opts LayoutFormatOptions, metrics foundation.IDictionary, views foundation.IDictionary) []LayoutConstraint {
	rv := objc.Send[[]LayoutConstraint](objc.ID(lc.class), objc.Sel("constraintsWithVisualFormat:options:metrics:views:"), format, opts, metrics, views)
	return rv
}


// Deactivates each constraint in the specified array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/deactivate(_:)
func (lc _LayoutConstraintClass) DeactivateConstraints(constraints []LayoutConstraint) {
	objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("deactivateConstraints:"), constraints)
}


// Creates a constraint that defines the relationship between the specified attributes of the given views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/init(item:attribute:relatedBy:toItem:attribute:multiplier:constant:)
func (lc _LayoutConstraintClass) ConstraintWithItemAttributeRelatedByToItemAttributeMultiplierConstant(view1 objectivec.IObject, attr1 LayoutAttribute, relation LayoutRelation, view2 objectivec.IObject, attr2 LayoutAttribute, multiplier float64, c float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(lc.class), objc.Sel("constraintWithItem:attribute:relatedBy:toItem:attribute:multiplier:constant:"), view1, attr1, relation, view2, attr2, multiplier, c)
	return rv
}

















// The constant added to the multiplied second attribute participating in the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/constant
func (l_ LayoutConstraint) Constant() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("constant"))
	return rv
}


// The constant added to the multiplied second attribute participating in the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/constant
func (l_ LayoutConstraint) SetConstant(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setConstant:"), value)
}


// The first anchor that defines the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/firstAnchor
func (l_ LayoutConstraint) FirstAnchor() ILayoutAnchor {
	rv := objc.Send[LayoutAnchor](l_.ID, objc.Sel("firstAnchor"))
	return rv
}


// The attribute of the first object participating in the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/firstAttribute
func (l_ LayoutConstraint) FirstAttribute() LayoutAttribute {
	rv := objc.Send[LayoutAttribute](l_.ID, objc.Sel("firstAttribute"))
	return rv
}


// The first object participating in the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/firstItem
func (l_ LayoutConstraint) FirstItem() objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("firstItem"))
	return rv
}


// The name that identifies the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/identifier
func (l_ LayoutConstraint) Identifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("identifier"))
	return rv
}


// The name that identifies the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/identifier
func (l_ LayoutConstraint) SetIdentifier(value foundation.foundation.INSString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIdentifier:"), value)
}


// The active state of the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/isActive
func (l_ LayoutConstraint) Active() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("active"))
	return rv
}


// The active state of the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/isActive
func (l_ LayoutConstraint) SetActive(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setActive:"), value)
}


// The multiplier applied to the second attribute participating in the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/multiplier
func (l_ LayoutConstraint) Multiplier() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("multiplier"))
	return rv
}


// The priority of the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/priority-swift.property
func (l_ LayoutConstraint) Priority() LayoutPriority {
	rv := objc.Send[LayoutPriority](l_.ID, objc.Sel("priority"))
	return rv
}


// The priority of the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/priority-swift.property
func (l_ LayoutConstraint) SetPriority(value LayoutPriority) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPriority:"), value)
}


// The relation between the two attributes in the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/relation-swift.property
func (l_ LayoutConstraint) Relation() LayoutRelation {
	rv := objc.Send[LayoutRelation](l_.ID, objc.Sel("relation"))
	return rv
}


// The second anchor that defines the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/secondAnchor
func (l_ LayoutConstraint) SecondAnchor() ILayoutAnchor {
	rv := objc.Send[LayoutAnchor](l_.ID, objc.Sel("secondAnchor"))
	return rv
}


// The attribute of the second object participating in the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/secondAttribute
func (l_ LayoutConstraint) SecondAttribute() LayoutAttribute {
	rv := objc.Send[LayoutAttribute](l_.ID, objc.Sel("secondAttribute"))
	return rv
}


// The second object participating in the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/secondItem
func (l_ LayoutConstraint) SecondItem() objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("secondItem"))
	return rv
}


// A Boolean value that determines whether the constraint should be archived by its owning view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/shouldBeArchived
func (l_ LayoutConstraint) ShouldBeArchived() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("shouldBeArchived"))
	return rv
}


// A Boolean value that determines whether the constraint should be archived by its owning view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/shouldBeArchived
func (l_ LayoutConstraint) SetShouldBeArchived(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShouldBeArchived:"), value)
}


// The active state of the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutconstraint/isactive
func (l_ LayoutConstraint) IsActive() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isActive"))
	return rv
}


// The active state of the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutconstraint/isactive
func (l_ LayoutConstraint) SetIsActive(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsActive:"), value)
}







