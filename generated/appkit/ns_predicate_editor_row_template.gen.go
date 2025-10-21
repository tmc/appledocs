// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A template that describes available predicates and how to display them.
//
// You can create instances of programmatically or in Interface Builder. By default, a noncompound row template has three views: a popup (or static text field) on the left, a popup or static text field for operators, and either a popup or other view on the right.  You can subclass to create a row template with different numbers or types of views. is a concrete class, but it has five primitive methods that are called by : , , , , and . implements all of these methods, but you can override them for custom templates. The primitive methods are used by an instance of as follows. First, an instance of is created, and some row templates are set on it—either through a nib file or programmatically. The first thing predicate editor does is ask each of the templates for their views, using . After setting up the predicate editor, you typically send it a message to restore a saved predicate. needs to determine which of its templates should display each predicate in the predicate tree. It does this by sending each of its row templates a message and choosing the one that returns the highest value. After finding the best match for a predicate, copies that template to get fresh views, inserts them into the proper row, and then sets the predicate on the template using . Within that method, the object must set its views’ values to represent that predicate. next asks the template for the “displayable sub-predicates” of the predicate by sending a message. If a template represents a predicate in its entirety, or if the predicate has no subpredicates, it can return for this.  Otherwise, it should return a list of predicates to be made into sub-rows of that template’s row. The whole process repeats for each sub-predicate. At this point, the user sees the predicate that was saved.  If the user then makes some changes to the views of the templates, this causes to recompute its predicate by asking each of the templates to return the predicate represented by the new view values, passing in the subpredicates represented by the sub-rows (an empty array if there are none, or if they aren’t supported by that predicate type):
//
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


// The value of the receiver’s cell as an Objective-C object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/objectvalue
func (p_ PredicateEditorRowTemplate) ObjectValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("objectValue"))
	return rv
}


// SetObjectValue sets the value of the objectValue property.
// The value of the receiver’s cell as an Objective-C object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/objectvalue
func (p_ PredicateEditorRowTemplate) SetObjectValue(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObjectValue:"), value)
}

// The row templates for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditor/rowtemplates
func (p_ PredicateEditorRowTemplate) RowTemplates() NSPredicateEditorRowTemplate {
	rv := objc.Send[NSPredicateEditorRowTemplate](p_.ID, objc.Sel("rowTemplates"))
	return rv
}


// SetRowTemplates sets the value of the rowTemplates property.
// The row templates for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditor/rowtemplates
func (p_ PredicateEditorRowTemplate) SetRowTemplates(value IPredicateEditorRowTemplate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRowTemplates:"), value)
}

// Returns the compound predicate types.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/compoundtypes
func (p_ PredicateEditorRowTemplate) CompoundTypes() foundation.Number {
	rv := objc.Send[foundation.Number](p_.ID, objc.Sel("compoundTypes"))
	return rv
}


// SetCompoundTypes sets the value of the compoundTypes property.
// Returns the compound predicate types.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/compoundtypes
func (p_ PredicateEditorRowTemplate) SetCompoundTypes(value foundation.INumber) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCompoundTypes:"), value)
}

// Returns the left hand expressions for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/leftexpressions
func (p_ PredicateEditorRowTemplate) LeftExpressions() Expression {
	rv := objc.Send[Expression](p_.ID, objc.Sel("leftExpressions"))
	return rv
}


// SetLeftExpressions sets the value of the leftExpressions property.
// Returns the left hand expressions for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/leftexpressions
func (p_ PredicateEditorRowTemplate) SetLeftExpressions(value IExpression) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLeftExpressions:"), value)
}

// Returns the comparison predicate modifier for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/modifier
func (p_ PredicateEditorRowTemplate) Modifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("modifier"))
	return rv
}


// SetModifier sets the value of the modifier property.
// Returns the comparison predicate modifier for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/modifier
func (p_ PredicateEditorRowTemplate) SetModifier(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModifier:"), value)
}

// Returns the array of comparison predicate operators.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/operators
func (p_ PredicateEditorRowTemplate) Operators() foundation.Number {
	rv := objc.Send[foundation.Number](p_.ID, objc.Sel("operators"))
	return rv
}


// SetOperators sets the value of the operators property.
// Returns the array of comparison predicate operators.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/operators
func (p_ PredicateEditorRowTemplate) SetOperators(value foundation.INumber) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOperators:"), value)
}

// Returns the comparison predicate options.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/options
func (p_ PredicateEditorRowTemplate) Options() int {
	rv := objc.Send[int](p_.ID, objc.Sel("options"))
	return rv
}


// SetOptions sets the value of the options property.
// Returns the comparison predicate options.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/options
func (p_ PredicateEditorRowTemplate) SetOptions(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOptions:"), value)
}

// Returns the attribute type of the receiver’s right expression.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/rightexpressionattributetype
func (p_ PredicateEditorRowTemplate) RightExpressionAttributeType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("rightExpressionAttributeType"))
	return rv
}


// SetRightExpressionAttributeType sets the value of the rightExpressionAttributeType property.
// Returns the attribute type of the receiver’s right expression.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/rightexpressionattributetype
func (p_ PredicateEditorRowTemplate) SetRightExpressionAttributeType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRightExpressionAttributeType:"), value)
}

// Returns the right hand expressions for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/rightexpressions
func (p_ PredicateEditorRowTemplate) RightExpressions() Expression {
	rv := objc.Send[Expression](p_.ID, objc.Sel("rightExpressions"))
	return rv
}


// SetRightExpressions sets the value of the rightExpressions property.
// Returns the right hand expressions for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/rightexpressions
func (p_ PredicateEditorRowTemplate) SetRightExpressions(value IExpression) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRightExpressions:"), value)
}

// Returns the views that display this template’s predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/templateviews
func (p_ PredicateEditorRowTemplate) TemplateViews() NSView {
	rv := objc.Send[NSView](p_.ID, objc.Sel("templateViews"))
	return rv
}


// SetTemplateViews sets the value of the templateViews property.
// Returns the views that display this template’s predicate.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/templateviews
func (p_ PredicateEditorRowTemplate) SetTemplateViews(value IView) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTemplateViews:"), value)
}



