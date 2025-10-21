// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [RuleEditor] class.
type IRuleEditor interface {
	IControl
}

// An interface for configuring a rule-based list of options.
//
// A rule editor lets the user visually create and configure a list of options that are expressed as a predicate (as described in ). Each row displayed by the rule editor represents a particular path down a tree of choices. The rule editor’s delegate provides the tree of choices to be displayed. The rule editor presents those choices to the user as a row of popup buttons, static text fields, and custom views. exposes one binding, . You can bind to an ordered collection (such as an instance of ). Each object in the collection should have the following properties:
//
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

// Alloc allocates a new instance without initialization.
func (rc _RuleEditorClass) Alloc() RuleEditor {
	rv := objc.Send[RuleEditor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The key path for the subrows.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/subrowsKeyPath
func (r_ RuleEditor) SubrowsKeyPath() string {
	rv := objc.Send[string](r_.ID, objc.Sel("subrowsKeyPath"))
	return rv
}


// SetSubrowsKeyPath sets the value of the subrowsKeyPath property.
// The key path for the subrows.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/subrowsKeyPath
func (r_ RuleEditor) SetSubrowsKeyPath(value string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSubrowsKeyPath:"), objc.String(value))
}

// A Boolean value that indicates whether all the rows can be removed.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/canremoveallrows
func (r_ RuleEditor) CanRemoveAllRows() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("canRemoveAllRows"))
	return rv
}


// SetCanRemoveAllRows sets the value of the canRemoveAllRows property.
// A Boolean value that indicates whether all the rows can be removed.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/canremoveallrows
func (r_ RuleEditor) SetCanRemoveAllRows(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCanRemoveAllRows:"), value)
}

// The criteria key path.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/criteriakeypath
func (r_ RuleEditor) CriteriaKeyPath() string {
	rv := objc.Send[string](r_.ID, objc.Sel("criteriaKeyPath"))
	return rv
}


// SetCriteriaKeyPath sets the value of the criteriaKeyPath property.
// The criteria key path.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/criteriakeypath
func (r_ RuleEditor) SetCriteriaKeyPath(value string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCriteriaKeyPath:"), objc.String(value))
}

// The rule editor’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/delegate
func (r_ RuleEditor) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The rule editor’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/delegate
func (r_ RuleEditor) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}

// The display values key path.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/displayvalueskeypath
func (r_ RuleEditor) DisplayValuesKeyPath() string {
	rv := objc.Send[string](r_.ID, objc.Sel("displayValuesKeyPath"))
	return rv
}


// SetDisplayValuesKeyPath sets the value of the displayValuesKeyPath property.
// The display values key path.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/displayvalueskeypath
func (r_ RuleEditor) SetDisplayValuesKeyPath(value string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDisplayValuesKeyPath:"), objc.String(value))
}

// The formatting dictionary for the rule editor.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/formattingdictionary
func (r_ RuleEditor) FormattingDictionary() string {
	rv := objc.Send[string](r_.ID, objc.Sel("formattingDictionary"))
	return rv
}


// SetFormattingDictionary sets the value of the formattingDictionary property.
// The formatting dictionary for the rule editor.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/formattingdictionary
func (r_ RuleEditor) SetFormattingDictionary(value string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFormattingDictionary:"), objc.String(value))
}

// The name of the rule editor’s strings file.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/formattingstringsfilename
func (r_ RuleEditor) FormattingStringsFilename() string {
	rv := objc.Send[string](r_.ID, objc.Sel("formattingStringsFilename"))
	return rv
}


// SetFormattingStringsFilename sets the value of the formattingStringsFilename property.
// The name of the rule editor’s strings file.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/formattingstringsfilename
func (r_ RuleEditor) SetFormattingStringsFilename(value string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFormattingStringsFilename:"), objc.String(value))
}

// A Boolean value that determines whether the rule editor is editable.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/iseditable
func (r_ RuleEditor) IsEditable() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isEditable"))
	return rv
}


// SetIsEditable sets the value of the isEditable property.
// A Boolean value that determines whether the rule editor is editable.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/iseditable
func (r_ RuleEditor) SetIsEditable(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsEditable:"), value)
}

// The rule editor’s nesting mode.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/nestingmode-swift.property
func (r_ RuleEditor) NestingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("nestingMode"))
	return rv
}


// SetNestingMode sets the value of the nestingMode property.
// The rule editor’s nesting mode.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/nestingmode-swift.property
func (r_ RuleEditor) SetNestingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNestingMode:"), value)
}

// The number of rows in the rule editor.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/numberofrows
func (r_ RuleEditor) NumberOfRows() int {
	rv := objc.Send[int](r_.ID, objc.Sel("numberOfRows"))
	return rv
}


// SetNumberOfRows sets the value of the numberOfRows property.
// The number of rows in the rule editor.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/numberofrows
func (r_ RuleEditor) SetNumberOfRows(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNumberOfRows:"), value)
}

// The rule editor’s predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/predicate
func (r_ RuleEditor) Predicate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("predicate"))
	return rv
}


// SetPredicate sets the value of the predicate property.
// The rule editor’s predicate.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/predicate
func (r_ RuleEditor) SetPredicate(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPredicate:"), value)
}

// The class used to create a new row in the “rows” binding.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/rowclass
func (r_ RuleEditor) RowClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("rowClass"))
	return rv
}


// SetRowClass sets the value of the rowClass property.
// The class used to create a new row in the “rows” binding.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/rowclass
func (r_ RuleEditor) SetRowClass(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRowClass:"), value)
}

// The rule editor’s row height.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/rowheight
func (r_ RuleEditor) RowHeight() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("rowHeight"))
	return rv
}


// SetRowHeight sets the value of the rowHeight property.
// The rule editor’s row height.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/rowheight
func (r_ RuleEditor) SetRowHeight(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRowHeight:"), value)
}

// The key path for the row type.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/rowtypekeypath
func (r_ RuleEditor) RowTypeKeyPath() string {
	rv := objc.Send[string](r_.ID, objc.Sel("rowTypeKeyPath"))
	return rv
}


// SetRowTypeKeyPath sets the value of the rowTypeKeyPath property.
// The key path for the row type.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/rowtypekeypath
func (r_ RuleEditor) SetRowTypeKeyPath(value string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRowTypeKeyPath:"), objc.String(value))
}

// The indexes of the rule editor’s selected rows.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/selectedrowindexes
func (r_ RuleEditor) SelectedRowIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("selectedRowIndexes"))
	return rv
}


// SetSelectedRowIndexes sets the value of the selectedRowIndexes property.
// The indexes of the rule editor’s selected rows.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditor/selectedrowindexes
func (r_ RuleEditor) SetSelectedRowIndexes(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSelectedRowIndexes:"), value)
}



