// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PRuleEditorDelegate is the NSRuleEditorDelegate protocol interface.
//
// The   protocol defines the optional methods implemented by delegates of   objects.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSRuleEditorDelegate
type PRuleEditorDelegate interface {
	// Required methods
	RuleEditorChildForCriterionWithRowType(editor IRuleEditor, index int, criterion objectivec.IObject, rowType RuleEditorRowType) objc.ID
	RuleEditorDisplayValueForCriterionInRow(editor IRuleEditor, criterion objectivec.IObject, row int) objc.ID
	RuleEditorNumberOfChildrenForCriterionWithRowType(editor IRuleEditor, criterion objectivec.IObject, rowType RuleEditorRowType) int
	// Optional methods
	RuleEditorPredicatePartsForCriterionWithDisplayValueInRow(editor IRuleEditor, criterion objectivec.IObject, value objectivec.IObject, row int) foundation.IDictionary
	HasRuleEditorPredicatePartsForCriterionWithDisplayValueInRow() bool
	RuleEditorRowsDidChange(notification foundation.foundation.INSNotification)
	HasRuleEditorRowsDidChange() bool
}

// RuleEditorDelegate is a delegate implementation builder for the PRuleEditorDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type RuleEditorDelegate struct {
	_RuleEditorPredicatePartsForCriterionWithDisplayValueInRow func(editor IRuleEditor, criterion objectivec.IObject, value objectivec.IObject, row int) foundation.IDictionary
	_RuleEditorRowsDidChange func(notification foundation.foundation.INSNotification)
	_RuleEditorChildForCriterionWithRowType func(editor IRuleEditor, index int, criterion objectivec.IObject, rowType RuleEditorRowType) objc.ID
	_RuleEditorDisplayValueForCriterionInRow func(editor IRuleEditor, criterion objectivec.IObject, row int) objc.ID
	_RuleEditorNumberOfChildrenForCriterionWithRowType func(editor IRuleEditor, criterion objectivec.IObject, rowType RuleEditorRowType) int
}

// SetRuleEditorPredicatePartsForCriterionWithDisplayValueInRow sets the handler for the RuleEditorPredicatePartsForCriterionWithDisplayValueInRow delegate method.
//
// Returns a dictionary representing the parts of the predicate determined by the given criterion and value.
func (d *RuleEditorDelegate) SetRuleEditorPredicatePartsForCriterionWithDisplayValueInRow(f func(editor IRuleEditor, criterion objectivec.IObject, value objectivec.IObject, row int) foundation.IDictionary) {
	d._RuleEditorPredicatePartsForCriterionWithDisplayValueInRow = f
}

// SetRuleEditorRowsDidChange sets the handler for the RuleEditorRowsDidChange delegate method.
//
// Notifies the receiver that a rule editor’s rows changed.
func (d *RuleEditorDelegate) SetRuleEditorRowsDidChange(f func(notification foundation.foundation.INSNotification)) {
	d._RuleEditorRowsDidChange = f
}

// SetRuleEditorChildForCriterionWithRowType sets the handler for the RuleEditorChildForCriterionWithRowType delegate method.
//
// Returns the child of a given item at a given index.
func (d *RuleEditorDelegate) SetRuleEditorChildForCriterionWithRowType(f func(editor IRuleEditor, index int, criterion objectivec.IObject, rowType RuleEditorRowType) objc.ID) {
	d._RuleEditorChildForCriterionWithRowType = f
}

// SetRuleEditorDisplayValueForCriterionInRow sets the handler for the RuleEditorDisplayValueForCriterionInRow delegate method.
//
// Returns the value for a given criterion.
func (d *RuleEditorDelegate) SetRuleEditorDisplayValueForCriterionInRow(f func(editor IRuleEditor, criterion objectivec.IObject, row int) objc.ID) {
	d._RuleEditorDisplayValueForCriterionInRow = f
}

// SetRuleEditorNumberOfChildrenForCriterionWithRowType sets the handler for the RuleEditorNumberOfChildrenForCriterionWithRowType delegate method.
//
// Returns the number of child items of a given criterion or row type.
func (d *RuleEditorDelegate) SetRuleEditorNumberOfChildrenForCriterionWithRowType(f func(editor IRuleEditor, criterion objectivec.IObject, rowType RuleEditorRowType) int) {
	d._RuleEditorNumberOfChildrenForCriterionWithRowType = f
}

// RuleEditorPredicatePartsForCriterionWithDisplayValueInRow implements the PRuleEditorDelegate interface.
func (d *RuleEditorDelegate) RuleEditorPredicatePartsForCriterionWithDisplayValueInRow(editor IRuleEditor, criterion objectivec.IObject, value objectivec.IObject, row int) foundation.IDictionary {
	if d._RuleEditorPredicatePartsForCriterionWithDisplayValueInRow != nil {
		return d._RuleEditorPredicatePartsForCriterionWithDisplayValueInRow(editor, criterion, value, row)
	}
	var zero foundation.IDictionary
	return zero
}

// HasRuleEditorPredicatePartsForCriterionWithDisplayValueInRow returns true if a handler for RuleEditorPredicatePartsForCriterionWithDisplayValueInRow has been set.
func (d *RuleEditorDelegate) HasRuleEditorPredicatePartsForCriterionWithDisplayValueInRow() bool {
	return d._RuleEditorPredicatePartsForCriterionWithDisplayValueInRow != nil
}

// RuleEditorRowsDidChange implements the PRuleEditorDelegate interface.
func (d *RuleEditorDelegate) RuleEditorRowsDidChange(notification foundation.foundation.INSNotification) {
	if d._RuleEditorRowsDidChange != nil {
		d._RuleEditorRowsDidChange(notification)
	}
}

// HasRuleEditorRowsDidChange returns true if a handler for RuleEditorRowsDidChange has been set.
func (d *RuleEditorDelegate) HasRuleEditorRowsDidChange() bool {
	return d._RuleEditorRowsDidChange != nil
}

// RuleEditorChildForCriterionWithRowType implements the PRuleEditorDelegate interface.
func (d *RuleEditorDelegate) RuleEditorChildForCriterionWithRowType(editor IRuleEditor, index int, criterion objectivec.IObject, rowType RuleEditorRowType) objc.ID {
	if d._RuleEditorChildForCriterionWithRowType != nil {
		return d._RuleEditorChildForCriterionWithRowType(editor, index, criterion, rowType)
	}
	var zero objc.ID
	return zero
}

// HasRuleEditorChildForCriterionWithRowType returns true if a handler for RuleEditorChildForCriterionWithRowType has been set.
func (d *RuleEditorDelegate) HasRuleEditorChildForCriterionWithRowType() bool {
	return d._RuleEditorChildForCriterionWithRowType != nil
}

// RuleEditorDisplayValueForCriterionInRow implements the PRuleEditorDelegate interface.
func (d *RuleEditorDelegate) RuleEditorDisplayValueForCriterionInRow(editor IRuleEditor, criterion objectivec.IObject, row int) objc.ID {
	if d._RuleEditorDisplayValueForCriterionInRow != nil {
		return d._RuleEditorDisplayValueForCriterionInRow(editor, criterion, row)
	}
	var zero objc.ID
	return zero
}

// HasRuleEditorDisplayValueForCriterionInRow returns true if a handler for RuleEditorDisplayValueForCriterionInRow has been set.
func (d *RuleEditorDelegate) HasRuleEditorDisplayValueForCriterionInRow() bool {
	return d._RuleEditorDisplayValueForCriterionInRow != nil
}

// RuleEditorNumberOfChildrenForCriterionWithRowType implements the PRuleEditorDelegate interface.
func (d *RuleEditorDelegate) RuleEditorNumberOfChildrenForCriterionWithRowType(editor IRuleEditor, criterion objectivec.IObject, rowType RuleEditorRowType) int {
	if d._RuleEditorNumberOfChildrenForCriterionWithRowType != nil {
		return d._RuleEditorNumberOfChildrenForCriterionWithRowType(editor, criterion, rowType)
	}
	var zero int
	return zero
}

// HasRuleEditorNumberOfChildrenForCriterionWithRowType returns true if a handler for RuleEditorNumberOfChildrenForCriterionWithRowType has been set.
func (d *RuleEditorDelegate) HasRuleEditorNumberOfChildrenForCriterionWithRowType() bool {
	return d._RuleEditorNumberOfChildrenForCriterionWithRowType != nil
}

// RuleEditorDelegateObject wraps an existing Objective-C object that conforms to the PRuleEditorDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type RuleEditorDelegateObject struct {
	objectivec.Object
}

// NewRuleEditorDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSRuleEditorDelegate protocol.
func NewRuleEditorDelegateObject(obj objectivec.Object) *RuleEditorDelegateObject {
	return &RuleEditorDelegateObject{obj}
}

// Make sure RuleEditorDelegateObject implements PRuleEditorDelegate.
var _ PRuleEditorDelegate = (*RuleEditorDelegateObject)(nil)

// RuleEditorChildForCriterionWithRowType implements the PRuleEditorDelegate interface.
// This required method is always available on objects conforming to RuleEditorChildForCriterionWithRowType.
func (o *RuleEditorDelegateObject) RuleEditorChildForCriterionWithRowType(editor IRuleEditor, index int, criterion objectivec.IObject, rowType RuleEditorRowType) objc.ID {
	return objc.Send[objc.ID](o.ID, objc.Sel("ruleEditor:child:forCriterion:withRowType:"), editor, index, criterion, rowType)
}

// RuleEditorDisplayValueForCriterionInRow implements the PRuleEditorDelegate interface.
// This required method is always available on objects conforming to RuleEditorDisplayValueForCriterionInRow.
func (o *RuleEditorDelegateObject) RuleEditorDisplayValueForCriterionInRow(editor IRuleEditor, criterion objectivec.IObject, row int) objc.ID {
	return objc.Send[objc.ID](o.ID, objc.Sel("ruleEditor:displayValueForCriterion:inRow:"), editor, criterion, row)
}

// RuleEditorNumberOfChildrenForCriterionWithRowType implements the PRuleEditorDelegate interface.
// This required method is always available on objects conforming to RuleEditorNumberOfChildrenForCriterionWithRowType.
func (o *RuleEditorDelegateObject) RuleEditorNumberOfChildrenForCriterionWithRowType(editor IRuleEditor, criterion objectivec.IObject, rowType RuleEditorRowType) int {
	return objc.Send[int](o.ID, objc.Sel("ruleEditor:numberOfChildrenForCriterion:withRowType:"), editor, criterion, rowType)
}

// RuleEditorPredicatePartsForCriterionWithDisplayValueInRow implements the PRuleEditorDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *RuleEditorDelegateObject) RuleEditorPredicatePartsForCriterionWithDisplayValueInRow(editor IRuleEditor, criterion objectivec.IObject, value objectivec.IObject, row int) foundation.IDictionary {
	return objc.Send[foundation.IDictionary](o.ID, objc.Sel("ruleEditor:predicatePartsForCriterion:withDisplayValue:inRow:"), editor, criterion, value, row)
}

// HasRuleEditorPredicatePartsForCriterionWithDisplayValueInRow returns true; this is a placeholder for optional method checks.
func (o *RuleEditorDelegateObject) HasRuleEditorPredicatePartsForCriterionWithDisplayValueInRow() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// RuleEditorRowsDidChange implements the PRuleEditorDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *RuleEditorDelegateObject) RuleEditorRowsDidChange(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("ruleEditorRowsDidChange:"), notification)
}

// HasRuleEditorRowsDidChange returns true; this is a placeholder for optional method checks.
func (o *RuleEditorDelegateObject) HasRuleEditorRowsDidChange() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
