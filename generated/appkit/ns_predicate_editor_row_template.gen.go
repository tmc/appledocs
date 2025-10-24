// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coredata"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PredicateEditorRowTemplate] class.
var (
	PredicateEditorRowTemplateClass     _PredicateEditorRowTemplateClass
	PredicateEditorRowTemplateClassOnce sync.Once
)

func getPredicateEditorRowTemplateClass() _PredicateEditorRowTemplateClass {
	PredicateEditorRowTemplateClassOnce.Do(func() {
		PredicateEditorRowTemplateClass = _PredicateEditorRowTemplateClass{objc.GetClass("NSPredicateEditorRowTemplate")}
	})
	return PredicateEditorRowTemplateClass
}

type _PredicateEditorRowTemplateClass struct {
	class objc.Class
}

// An interface definition for the [PredicateEditorRowTemplate] class.
type IPredicateEditorRowTemplate interface {
	objectivec.IObject
	// properties:
	CompoundTypes() []foundation.Number
	LeftExpressions() []coredata.Expression
	Modifier() ComparisonPredicateModifier /* not a class type */
	Operators() []foundation.Number
	Options() uint
	RightExpressionAttributeType() AttributeType /* not a class type */
	RightExpressions() []coredata.Expression
	TemplateViews() []View
	RowTemplates() IPredicateEditorRowTemplate
	SetRowTemplates(value IPredicateEditorRowTemplate)
	// methods:
	DisplayableSubpredicatesOfPredicate(predicate foundation.Predicate) []foundation.Predicate
	MatchForPredicate(predicate foundation.Predicate) float64
	PredicateWithSubpredicates(subpredicates []foundation.Predicate) foundation.Predicate
	SetPredicate(predicate foundation.Predicate)
}

// A template that describes available predicates and how to display them.
//
// You can create instances of programmatically or in Interface Builder. By default, a noncompound row template has three views: a popup (or static text field) on the left, a popup or static text field for operators, and either a popup or other view on the right.  You can subclass to create a row template with different numbers or types of views. is a concrete class, but it has five primitive methods that are called by : , , , , and . implements all of these methods, but you can override them for custom templates. The primitive methods are used by an instance of as follows. First, an instance of is created, and some row templates are set on it—either through a nib file or programmatically. The first thing predicate editor does is ask each of the templates for their views, using . After setting up the predicate editor, you typically send it a message to restore a saved predicate. needs to determine which of its templates should display each predicate in the predicate tree. It does this by sending each of its row templates a message and choosing the one that returns the highest value. After finding the best match for a predicate, copies that template to get fresh views, inserts them into the proper row, and then sets the predicate on the template using . Within that method, the object must set its views’ values to represent that predicate. next asks the template for the “displayable sub-predicates” of the predicate by sending a message. If a template represents a predicate in its entirety, or if the predicate has no subpredicates, it can return for this.  Otherwise, it should return a list of predicates to be made into sub-rows of that template’s row. The whole process repeats for each sub-predicate. At this point, the user sees the predicate that was saved.  If the user then makes some changes to the views of the templates, this causes to recompute its predicate by asking each of the templates to return the predicate represented by the new view values, passing in the subpredicates represented by the sub-rows (an empty array if there are none, or if they aren’t supported by that predicate type):


// A template that describes available predicates and how to display them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate
type PredicateEditorRowTemplate struct {
	objectivec.Object
}

// PredicateEditorRowTemplateFrom constructs a [PredicateEditorRowTemplate] from an unsafe.Pointer.
//
// A template that describes available predicates and how to display them.
func PredicateEditorRowTemplateFrom(ptr unsafe.Pointer) PredicateEditorRowTemplate {
	return PredicateEditorRowTemplate{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PredicateEditorRowTemplateClass) Alloc() PredicateEditorRowTemplate {
	rv := objc.Send[PredicateEditorRowTemplate](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PredicateEditorRowTemplateClass) New() PredicateEditorRowTemplate {
	rv := objc.Send[PredicateEditorRowTemplate](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PredicateEditorRowTemplate) Init() PredicateEditorRowTemplate {
	rv := objc.Send[PredicateEditorRowTemplate](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PredicateEditorRowTemplate) Autorelease() PredicateEditorRowTemplate {
	rv := objc.Send[PredicateEditorRowTemplate](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPredicateEditorRowTemplate creates a new PredicateEditorRowTemplate instance.
func NewPredicateEditorRowTemplate() PredicateEditorRowTemplate {
	return getPredicateEditorRowTemplateClass().New()
}



// Initializes and returns a row template suitable for displaying compound predicates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/init(compoundTypes:)
func NewPredicateEditorRowTemplateWithCompoundTypes(compoundTypes []foundation.Number) PredicateEditorRowTemplate {
	instance := getPredicateEditorRowTemplateClass().Alloc()
	rv := objc.Send[PredicateEditorRowTemplate](instance.ID, objc.Sel("initWithCompoundTypes:"), compoundTypes)
	rv.Autorelease()
	return rv
}


// Initializes and returns a “pop-up-pop-up-view”–style row template.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/init(leftExpressions:rightExpressionAttributeType:modifier:operators:options:)
func NewPredicateEditorRowTemplateWithLeftExpressionsRightExpressionAttributeTypeModifierOperatorsOptions(leftExpressions []coredata.Expression, attributeType AttributeType /* not a class type */, modifier ComparisonPredicateModifier /* not a class type */, operators []foundation.Number, options uint) PredicateEditorRowTemplate {
	instance := getPredicateEditorRowTemplateClass().Alloc()
	rv := objc.Send[PredicateEditorRowTemplate](instance.ID, objc.Sel("initWithLeftExpressions:rightExpressionAttributeType:modifier:operators:options:"), leftExpressions, attributeType, modifier, operators, options)
	rv.Autorelease()
	return rv
}


// Initializes and returns a “pop-up-pop-up-pop-up”–style row template.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/init(leftExpressions:rightExpressions:modifier:operators:options:)
func NewPredicateEditorRowTemplateWithLeftExpressionsRightExpressionsModifierOperatorsOptions(leftExpressions []coredata.Expression, rightExpressions []coredata.Expression, modifier ComparisonPredicateModifier /* not a class type */, operators []foundation.Number, options uint) PredicateEditorRowTemplate {
	instance := getPredicateEditorRowTemplateClass().Alloc()
	rv := objc.Send[PredicateEditorRowTemplate](instance.ID, objc.Sel("initWithLeftExpressions:rightExpressions:modifier:operators:options:"), leftExpressions, rightExpressions, modifier, operators, options)
	rv.Autorelease()
	return rv
}



// Returns an array of predicate templates for the given attribute key paths for a given entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/templates(withAttributeKeyPaths:in:)
func (pc _PredicateEditorRowTemplateClass) TemplatesWithAttributeKeyPathsInEntityDescription(keyPaths []string, entityDescription coredata.EntityDescription) []PredicateEditorRowTemplate {
	rv := objc.Send[[]PredicateEditorRowTemplate](objc.ID(pc.class), objc.Sel("templatesWithAttributeKeyPaths:inEntityDescription:"), keyPaths, entityDescription)
	return rv
}


// Returns the subpredicates that should be made sub-rows of a given predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/displayableSubpredicates(of:)
func (p_ PredicateEditorRowTemplate) DisplayableSubpredicatesOfPredicate(predicate foundation.Predicate) []foundation.Predicate {
	rv := objc.Send[[]foundation.Predicate](p_.ID, objc.Sel("displayableSubpredicatesOfPredicate:"), predicate)
	return rv
}


// Returns a positive number if the receiver can represent a given predicate, and if it cannot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/match(for:)
func (p_ PredicateEditorRowTemplate) MatchForPredicate(predicate foundation.Predicate) float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("matchForPredicate:"), predicate)
	return rv
}


// Returns the predicate represented by the receiver’s views’ values and the given sub-predicates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/predicate(withSubpredicates:)
func (p_ PredicateEditorRowTemplate) PredicateWithSubpredicates(subpredicates []foundation.Predicate) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](p_.ID, objc.Sel("predicateWithSubpredicates:"), subpredicates)
	return rv
}


// Sets the value of the views according to the given predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/setPredicate(_:)
func (p_ PredicateEditorRowTemplate) SetPredicate(predicate foundation.Predicate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPredicate:"), predicate)
}


// Returns the compound predicate types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/compoundTypes
func (p_ PredicateEditorRowTemplate) CompoundTypes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](p_.ID, objc.Sel("compoundTypes"))
	return rv
}


// Returns the left hand expressions for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/leftExpressions
func (p_ PredicateEditorRowTemplate) LeftExpressions() []coredata.Expression {
	rv := objc.Send[[]coredata.Expression](p_.ID, objc.Sel("leftExpressions"))
	return rv
}


// Returns the comparison predicate modifier for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/modifier
func (p_ PredicateEditorRowTemplate) Modifier() ComparisonPredicateModifier /* not a class type */ {
	rv := objc.Send[ComparisonPredicateModifier](p_.ID, objc.Sel("modifier"))
	return rv
}


// Returns the array of comparison predicate operators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/operators
func (p_ PredicateEditorRowTemplate) Operators() []foundation.Number {
	rv := objc.Send[[]foundation.Number](p_.ID, objc.Sel("operators"))
	return rv
}


// Returns the comparison predicate options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/options
func (p_ PredicateEditorRowTemplate) Options() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("options"))
	return rv
}


// Returns the attribute type of the receiver’s right expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/rightExpressionAttributeType
func (p_ PredicateEditorRowTemplate) RightExpressionAttributeType() AttributeType /* not a class type */ {
	rv := objc.Send[AttributeType](p_.ID, objc.Sel("rightExpressionAttributeType"))
	return rv
}


// Returns the right hand expressions for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/rightExpressions
func (p_ PredicateEditorRowTemplate) RightExpressions() []coredata.Expression {
	rv := objc.Send[[]coredata.Expression](p_.ID, objc.Sel("rightExpressions"))
	return rv
}


// Returns the views that display this template’s predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/templateViews
func (p_ PredicateEditorRowTemplate) TemplateViews() []View {
	rv := objc.Send[[]View](p_.ID, objc.Sel("templateViews"))
	return rv
}


// The row templates for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditor/rowtemplates
func (p_ PredicateEditorRowTemplate) RowTemplates() IPredicateEditorRowTemplate {
	rv := objc.Send[PredicateEditorRowTemplate](p_.ID, objc.Sel("rowTemplates"))
	return rv
}


// The row templates for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditor/rowtemplates
func (p_ PredicateEditorRowTemplate) SetRowTemplates(value IPredicateEditorRowTemplate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRowTemplates:"), value)
}


