// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSRuleEditor */


/* debug [class_header]: Header for NSRuleEditor */
// The class instance for the [RuleEditor] class.
var (
	RuleEditorClass     _RuleEditorClass
	RuleEditorClassOnce sync.Once
)

func getRuleEditorClass() _RuleEditorClass {
	RuleEditorClassOnce.Do(func() {
		RuleEditorClass = _RuleEditorClass{objc.GetClass("NSRuleEditor")}
	})
	return RuleEditorClass
}

type _RuleEditorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RuleEditor */
// An interface definition for the [RuleEditor] class.
type IRuleEditor interface {
	IControl
	
/* debug [class_interface_properties]: Properties for RuleEditor */
	// properties:
	CanRemoveAllRows() bool
	SetCanRemoveAllRows(value bool)
	CriteriaKeyPath() objc.IObject /* cross-framework: NSString */
	SetCriteriaKeyPath(value objc.IObject /* cross-framework: NSString */)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DisplayValuesKeyPath() objc.IObject /* cross-framework: NSString */
	SetDisplayValuesKeyPath(value objc.IObject /* cross-framework: NSString */)
	FormattingDictionary() foundation.IDictionary
	SetFormattingDictionary(value foundation.IDictionary)
	FormattingStringsFilename() objc.IObject /* cross-framework: NSString */
	SetFormattingStringsFilename(value objc.IObject /* cross-framework: NSString */)
	Editable() bool
	SetEditable(value bool)
	NestingMode() RuleEditorNestingMode
	SetNestingMode(value RuleEditorNestingMode)
	NumberOfRows() int
	Predicate() foundation.Predicate
	RowClass() objc.Class
	SetRowClass(value objc.Class)
	RowHeight() float64
	SetRowHeight(value float64)
	RowTypeKeyPath() objc.IObject /* cross-framework: NSString */
	SetRowTypeKeyPath(value objc.IObject /* cross-framework: NSString */)
	SelectedRowIndexes() foundation.IndexSet
	SubrowsKeyPath() objc.IObject /* cross-framework: NSString */
	SetSubrowsKeyPath(value objc.IObject /* cross-framework: NSString */)
	IsEditable() bool
	SetIsEditable(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RuleEditor */
	// methods:
	AddRow(sender objc.IObject)
	CriteriaForRow(row int) foundation.Array
	DisplayValuesForRow(row int) foundation.Array
	InsertRowAtIndexWithTypeAsSubrowOfRowAnimate(rowIndex int, rowType RuleEditorRowType, parentRow int, shouldAnimate bool)
	ParentRowForRow(rowIndex int) int
	PredicateForRow(row int) foundation.Predicate
	ReloadCriteria()
	ReloadPredicate()
	RemoveRowAtIndex(rowIndex int)
	RemoveRowsAtIndexesIncludeSubrows(rowIndexes foundation.IndexSet, includeSubrows bool)
	RowForDisplayValue(displayValue objc.IObject) int
	RowTypeForRow(rowIndex int) RuleEditorRowType
	SelectRowIndexesByExtendingSelection(indexes foundation.IndexSet, extend bool)
	SetCriteriaAndDisplayValuesForRowAtIndex(criteria objc.IObject /* cross-framework: NSArray */, values objc.IObject /* cross-framework: NSArray */, rowIndex int)
	SubrowIndexesForRow(rowIndex int) foundation.IndexSet
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RuleEditor */
// Alloc allocates a new instance without initialization.
func (rc _RuleEditorClass) Alloc() RuleEditor {
	rv := objc.Send[RuleEditor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RuleEditorClass) New() RuleEditor {
	rv := objc.Send[RuleEditor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RuleEditor) Init() RuleEditor {
	rv := objc.Send[RuleEditor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RuleEditor) Autorelease() RuleEditor {
	rv := objc.Send[RuleEditor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRuleEditor creates a new RuleEditor instance.
func NewRuleEditor() RuleEditor {
	return getRuleEditorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RuleEditor */
// An interface for configuring a rule-based list of options.
//
// A rule editor lets the user visually create and configure a list of options that are expressed as a predicate (as described in ). Each row displayed by the rule editor represents a particular path down a tree of choices. The rule editor’s delegate provides the tree of choices to be displayed. The rule editor presents those choices to the user as a row of popup buttons, static text fields, and custom views. exposes one binding, . You can bind to an ordered collection (such as an instance of ). Each object in the collection should have the following properties:


// An interface for configuring a rule-based list of options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor
type RuleEditor struct {
	Control
}

// RuleEditorFrom constructs a [RuleEditor] from an unsafe.Pointer.
//
// An interface for configuring a rule-based list of options.
func RuleEditorFrom(ptr unsafe.Pointer) RuleEditor {
	return RuleEditor{
		Control: ControlFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RuleEditor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RuleEditor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RuleEditor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RuleEditor */

// Adds a row to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/addRow(_:)
func (r_ RuleEditor) AddRow(sender objc.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addRow:"), sender)
}/* debug [instance_methods/method]: AddRow */


// Returns the currently chosen items for a given row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/criteria(forRow:)
func (r_ RuleEditor) CriteriaForRow(row int) foundation.Array {
	rv := objc.Send[foundation.Array](r_.ID, objc.Sel("criteriaForRow:"), row)
	return rv
}/* debug [instance_methods/method]: CriteriaForRow */


// Returns the chosen values for a given row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/displayValues(forRow:)
func (r_ RuleEditor) DisplayValuesForRow(row int) foundation.Array {
	rv := objc.Send[foundation.Array](r_.ID, objc.Sel("displayValuesForRow:"), row)
	return rv
}/* debug [instance_methods/method]: DisplayValuesForRow */


// Adds a new row of a given type at a given location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/insertRow(at:with:asSubrowOfRow:animate:)
func (r_ RuleEditor) InsertRowAtIndexWithTypeAsSubrowOfRowAnimate(rowIndex int, rowType RuleEditorRowType, parentRow int, shouldAnimate bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("insertRowAtIndex:withType:asSubrowOfRow:animate:"), rowIndex, rowType, parentRow, shouldAnimate)
}/* debug [instance_methods/method]: InsertRowAtIndexWithTypeAsSubrowOfRowAnimate */


// Returns the index of the parent of a given row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/parentRow(forRow:)
func (r_ RuleEditor) ParentRowForRow(rowIndex int) int {
	rv := objc.Send[int](r_.ID, objc.Sel("parentRowForRow:"), rowIndex)
	return rv
}/* debug [instance_methods/method]: ParentRowForRow */


// Returns the predicate for a given row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/predicate(forRow:)
func (r_ RuleEditor) PredicateForRow(row int) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](r_.ID, objc.Sel("predicateForRow:"), row)
	return rv
}/* debug [instance_methods/method]: PredicateForRow */


// Instructs the receiver to refetch criteria from its delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/reloadCriteria()
func (r_ RuleEditor) ReloadCriteria() {
	objc.Send[objc.ID](r_.ID, objc.Sel("reloadCriteria"))
}/* debug [instance_methods/method]: ReloadCriteria */


// Instructs the receiver to regenerate its predicate by invoking the corresponding delegate method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/reloadPredicate()
func (r_ RuleEditor) ReloadPredicate() {
	objc.Send[objc.ID](r_.ID, objc.Sel("reloadPredicate"))
}/* debug [instance_methods/method]: ReloadPredicate */


// Removes the row at a given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/removeRow(at:)
func (r_ RuleEditor) RemoveRowAtIndex(rowIndex int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("removeRowAtIndex:"), rowIndex)
}/* debug [instance_methods/method]: RemoveRowAtIndex */


// Removes the rows at given indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/removeRows(at:includeSubrows:)
func (r_ RuleEditor) RemoveRowsAtIndexesIncludeSubrows(rowIndexes foundation.IndexSet, includeSubrows bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("removeRowsAtIndexes:includeSubrows:"), rowIndexes, includeSubrows)
}/* debug [instance_methods/method]: RemoveRowsAtIndexesIncludeSubrows */


// Returns the index of the row containing a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/row(forDisplayValue:)
func (r_ RuleEditor) RowForDisplayValue(displayValue objc.IObject) int {
	rv := objc.Send[int](r_.ID, objc.Sel("rowForDisplayValue:"), displayValue)
	return rv
}/* debug [instance_methods/method]: RowForDisplayValue */


// Returns the type of a given row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/rowType(forRow:)
func (r_ RuleEditor) RowTypeForRow(rowIndex int) RuleEditorRowType {
	rv := objc.Send[RuleEditorRowType](r_.ID, objc.Sel("rowTypeForRow:"), rowIndex)
	return rv
}/* debug [instance_methods/method]: RowTypeForRow */


// Sets in the receiver the indexes of rows that are selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/selectRowIndexes(_:byExtendingSelection:)
func (r_ RuleEditor) SelectRowIndexesByExtendingSelection(indexes foundation.IndexSet, extend bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("selectRowIndexes:byExtendingSelection:"), indexes, extend)
}/* debug [instance_methods/method]: SelectRowIndexesByExtendingSelection */


// Modifies the row at a given index to contain the given items and values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/setCriteria(_:andDisplayValues:forRowAt:)
func (r_ RuleEditor) SetCriteriaAndDisplayValuesForRowAtIndex(criteria objc.IObject /* cross-framework: NSArray */, values objc.IObject /* cross-framework: NSArray */, rowIndex int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCriteria:andDisplayValues:forRowAtIndex:"), criteria, values, rowIndex)
}/* debug [instance_methods/method]: SetCriteriaAndDisplayValuesForRowAtIndex */


// Returns the immediate subrows of a given row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/subrowIndexes(forRow:)
func (r_ RuleEditor) SubrowIndexesForRow(rowIndex int) foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](r_.ID, objc.Sel("subrowIndexesForRow:"), rowIndex)
	return rv
}/* debug [instance_methods/method]: SubrowIndexesForRow */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RuleEditor */

// A Boolean value that indicates whether all the rows can be removed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/canRemoveAllRows
func (r_ RuleEditor) CanRemoveAllRows() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("canRemoveAllRows"))
	return rv
}/* debug [instance_properties/getter]: canRemoveAllRows */


// A Boolean value that indicates whether all the rows can be removed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/canRemoveAllRows
func (r_ RuleEditor) SetCanRemoveAllRows(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCanRemoveAllRows:"), value)
}/* debug [instance_properties/setter]: canRemoveAllRows */


// The criteria key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/criteriaKeyPath
func (r_ RuleEditor) CriteriaKeyPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("criteriaKeyPath"))
	return rv
}/* debug [instance_properties/getter]: criteriaKeyPath */


// The criteria key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/criteriaKeyPath
func (r_ RuleEditor) SetCriteriaKeyPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCriteriaKeyPath:"), value)
}/* debug [instance_properties/setter]: criteriaKeyPath */


// The rule editor’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/delegate
func (r_ RuleEditor) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The rule editor’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/delegate
func (r_ RuleEditor) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The display values key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/displayValuesKeyPath
func (r_ RuleEditor) DisplayValuesKeyPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("displayValuesKeyPath"))
	return rv
}/* debug [instance_properties/getter]: displayValuesKeyPath */


// The display values key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/displayValuesKeyPath
func (r_ RuleEditor) SetDisplayValuesKeyPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDisplayValuesKeyPath:"), value)
}/* debug [instance_properties/setter]: displayValuesKeyPath */


// The formatting dictionary for the rule editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/formattingDictionary
func (r_ RuleEditor) FormattingDictionary() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](r_.ID, objc.Sel("formattingDictionary"))
	return rv
}/* debug [instance_properties/getter]: formattingDictionary */


// The formatting dictionary for the rule editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/formattingDictionary
func (r_ RuleEditor) SetFormattingDictionary(value foundation.IDictionary) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFormattingDictionary:"), value)
}/* debug [instance_properties/setter]: formattingDictionary */


// The name of the rule editor’s strings file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/formattingStringsFilename
func (r_ RuleEditor) FormattingStringsFilename() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("formattingStringsFilename"))
	return rv
}/* debug [instance_properties/getter]: formattingStringsFilename */


// The name of the rule editor’s strings file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/formattingStringsFilename
func (r_ RuleEditor) SetFormattingStringsFilename(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFormattingStringsFilename:"), value)
}/* debug [instance_properties/setter]: formattingStringsFilename */


// A Boolean value that determines whether the rule editor is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/isEditable
func (r_ RuleEditor) Editable() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("editable"))
	return rv
}/* debug [instance_properties/getter]: editable */


// A Boolean value that determines whether the rule editor is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/isEditable
func (r_ RuleEditor) SetEditable(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setEditable:"), value)
}/* debug [instance_properties/setter]: editable */


// The rule editor’s nesting mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/nestingMode-swift.property
func (r_ RuleEditor) NestingMode() RuleEditorNestingMode {
	rv := objc.Send[RuleEditorNestingMode](r_.ID, objc.Sel("nestingMode"))
	return rv
}/* debug [instance_properties/getter]: nestingMode */


// The rule editor’s nesting mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/nestingMode-swift.property
func (r_ RuleEditor) SetNestingMode(value RuleEditorNestingMode) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNestingMode:"), value)
}/* debug [instance_properties/setter]: nestingMode */


// The number of rows in the rule editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/numberOfRows
func (r_ RuleEditor) NumberOfRows() int {
	rv := objc.Send[int](r_.ID, objc.Sel("numberOfRows"))
	return rv
}/* debug [instance_properties/getter]: numberOfRows */


// The rule editor’s predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/predicate
func (r_ RuleEditor) Predicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](r_.ID, objc.Sel("predicate"))
	return rv
}/* debug [instance_properties/getter]: predicate */


// The class used to create a new row in the “rows” binding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/rowClass
func (r_ RuleEditor) RowClass() objc.Class {
	rv := objc.Send[objc.Class](r_.ID, objc.Sel("rowClass"))
	return rv
}/* debug [instance_properties/getter]: rowClass */


// The class used to create a new row in the “rows” binding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/rowClass
func (r_ RuleEditor) SetRowClass(value objc.Class) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRowClass:"), value)
}/* debug [instance_properties/setter]: rowClass */


// The rule editor’s row height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/rowHeight
func (r_ RuleEditor) RowHeight() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("rowHeight"))
	return rv
}/* debug [instance_properties/getter]: rowHeight */


// The rule editor’s row height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/rowHeight
func (r_ RuleEditor) SetRowHeight(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRowHeight:"), value)
}/* debug [instance_properties/setter]: rowHeight */


// The key path for the row type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/rowTypeKeyPath
func (r_ RuleEditor) RowTypeKeyPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("rowTypeKeyPath"))
	return rv
}/* debug [instance_properties/getter]: rowTypeKeyPath */


// The key path for the row type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/rowTypeKeyPath
func (r_ RuleEditor) SetRowTypeKeyPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRowTypeKeyPath:"), value)
}/* debug [instance_properties/setter]: rowTypeKeyPath */


// The indexes of the rule editor’s selected rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/selectedRowIndexes
func (r_ RuleEditor) SelectedRowIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](r_.ID, objc.Sel("selectedRowIndexes"))
	return rv
}/* debug [instance_properties/getter]: selectedRowIndexes */


// The key path for the subrows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/subrowsKeyPath
func (r_ RuleEditor) SubrowsKeyPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("subrowsKeyPath"))
	return rv
}/* debug [instance_properties/getter]: subrowsKeyPath */


// The key path for the subrows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/subrowsKeyPath
func (r_ RuleEditor) SetSubrowsKeyPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSubrowsKeyPath:"), value)
}/* debug [instance_properties/setter]: subrowsKeyPath */


// A Boolean value that determines whether the rule editor is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/iseditable
func (r_ RuleEditor) IsEditable() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isEditable"))
	return rv
}/* debug [instance_properties/getter]: isEditable */


// A Boolean value that determines whether the rule editor is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/iseditable
func (r_ RuleEditor) SetIsEditable(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsEditable:"), value)
}/* debug [instance_properties/setter]: isEditable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSRuleEditor */



