// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ABPeoplePickerView] class.
var (
	ABPeoplePickerViewClass     _ABPeoplePickerViewClass
	ABPeoplePickerViewClassOnce sync.Once
)

func getABPeoplePickerViewClass() _ABPeoplePickerViewClass {
	ABPeoplePickerViewClassOnce.Do(func() {
		ABPeoplePickerViewClass = _ABPeoplePickerViewClass{objc.GetClass("ABPeoplePickerView")}
	})
	return ABPeoplePickerViewClass
}

type _ABPeoplePickerViewClass struct {
	class objc.Class
}

// An interface definition for the [ABPeoplePickerView] class.
type IABPeoplePickerView interface {
	appkit.IView
	// properties:
	AccessoryView() objc.IObject /* cross-framework: View */
	SetAccessoryView(value objc.IObject /* cross-framework: View */)
	AllowsGroupSelection() bool /* primitive/slice/pointer. */
	SetAllowsGroupSelection(value bool /* primitive/slice/pointer. */)
	AllowsMultipleSelection() bool /* primitive/slice/pointer. */
	SetAllowsMultipleSelection(value bool /* primitive/slice/pointer. */)
	AutosaveName() string /* primitive/slice/pointer. */
	SetAutosaveName(value string /* primitive/slice/pointer. */)
	DisplayedProperty() string /* primitive/slice/pointer. */
	SetDisplayedProperty(value string /* primitive/slice/pointer. */)
	GroupDoubleAction() objc.SEL
	SetGroupDoubleAction(value objc.SEL)
	NameDoubleAction() objc.SEL
	SetNameDoubleAction(value objc.SEL)
	SelectedGroups() objc.ID
	SelectedRecords() objc.ID
	Target() objc.ID
	SetTarget(value objc.ID)
	ValueSelectionBehavior() unsafe.Pointer
	SetValueSelectionBehavior(value unsafe.Pointer)
	// methods:
	AddProperty(property string /* primitive/slice/pointer. */)
	ClearSearchField(sender objectivec.IObject)
	ColumnTitleForProperty(property string /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */
	DeselectRecord(record IABRecord)
	DeselectGroup(group IABGroup)
	DeselectAll(sender objectivec.IObject)
	DeselectIdentifierForPerson(identifier string /* primitive/slice/pointer. */, person IABPerson)
	EditInAddressBook(sender objectivec.IObject)
	Properties() objc.IObject /* cross-framework: Array */
	RemoveProperty(property string /* primitive/slice/pointer. */)
	SelectGroupByExtendingSelection(group IABGroup, extend bool /* primitive/slice/pointer. */)
	SelectRecordByExtendingSelection(record IABRecord, extend bool /* primitive/slice/pointer. */)
	SelectIdentifierForPersonByExtendingSelection(identifier string /* primitive/slice/pointer. */, person IABPerson, extend bool /* primitive/slice/pointer. */)
	SelectInAddressBook(sender objectivec.IObject)
	SelectedIdentifiersForPerson(person IABPerson) objc.IObject /* cross-framework: Array */
	SelectedValues() objc.IObject /* cross-framework: Array */
	SetColumnTitleForProperty(title string /* primitive/slice/pointer. */, property string /* primitive/slice/pointer. */)
}

// An object you use to customize the behavior of people-picker views in an app’s user interface.


// An object you use to customize the behavior of people-picker views in an app’s user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView
type ABPeoplePickerView struct {
	appkit.View
}

// ABPeoplePickerViewFrom constructs a [ABPeoplePickerView] from an unsafe.Pointer.
//
// An object you use to customize the behavior of people-picker views in an app’s user interface.
func ABPeoplePickerViewFrom(ptr unsafe.Pointer) ABPeoplePickerView {
	return ABPeoplePickerView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _ABPeoplePickerViewClass) Alloc() ABPeoplePickerView {
	rv := objc.Send[ABPeoplePickerView](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ABPeoplePickerViewClass) New() ABPeoplePickerView {
	rv := objc.Send[ABPeoplePickerView](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ABPeoplePickerView) Init() ABPeoplePickerView {
	rv := objc.Send[ABPeoplePickerView](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ABPeoplePickerView) Autorelease() ABPeoplePickerView {
	rv := objc.Send[ABPeoplePickerView](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewABPeoplePickerView creates a new ABPeoplePickerView instance.
func NewABPeoplePickerView() ABPeoplePickerView {
	return getABPeoplePickerViewClass().New()
}



// Adds a property to the group of properties whose values are shown in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/addProperty(_:)
func (a_ ABPeoplePickerView) AddProperty(property string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addProperty:"), objc.String(property))
}


// Clears the search field and resets the list of displayed records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/clearSearchField(_:)
func (a_ ABPeoplePickerView) ClearSearchField(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("clearSearchField:"), sender)
}


// Returns the title of a custom property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/columnTitle(forProperty:)
func (a_ ABPeoplePickerView) ColumnTitleForProperty(property string /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](a_.ID, objc.Sel("columnTitleForProperty:"), objc.String(property))
	return rv
}


// Deselects a record selected in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/deselect(_:)-1yy11
func (a_ ABPeoplePickerView) DeselectRecord(record IABRecord) {
	objc.Send[objc.ID](a_.ID, objc.Sel("deselectRecord:"), record)
}


// Deselects a group selected in the group list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/deselect(_:)-3x7tl
func (a_ ABPeoplePickerView) DeselectGroup(group IABGroup) {
	objc.Send[objc.ID](a_.ID, objc.Sel("deselectGroup:"), group)
}


// Deselects all selected groups, records, and values in multivalue properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/deselectAll(_:)
func (a_ ABPeoplePickerView) DeselectAll(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("deselectAll:"), sender)
}


// Deselects a value selected in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/deselectIdentifier(_:for:)
func (a_ ABPeoplePickerView) DeselectIdentifierForPerson(identifier string /* primitive/slice/pointer. */, person IABPerson) {
	objc.Send[objc.ID](a_.ID, objc.Sel("deselectIdentifier:forPerson:"), objc.String(identifier), person)
}


// Launches Address Book to edit the item selected in the people picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/editInAddressBook(_:)
func (a_ ABPeoplePickerView) EditInAddressBook(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("editInAddressBook:"), sender)
}


// Returns an array of the properties whose values are shown in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/properties()
func (a_ ABPeoplePickerView) Properties() objc.IObject /* cross-framework: Array */ {
	rv := objc.Send[Array](a_.ID, objc.Sel("properties"))
	return rv
}


// Removes a property from the group of properties whose values are shown in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/removeProperty(_:)
func (a_ ABPeoplePickerView) RemoveProperty(property string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeProperty:"), objc.String(property))
}


// Selects a group or a set of groups in the group list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/select(_:byExtendingSelection:)-6mrii
func (a_ ABPeoplePickerView) SelectGroupByExtendingSelection(group IABGroup, extend bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("selectGroup:byExtendingSelection:"), group, extend)
}


// Selects a record or a set of records in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/select(_:byExtendingSelection:)-9eldk
func (a_ ABPeoplePickerView) SelectRecordByExtendingSelection(record IABRecord, extend bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("selectRecord:byExtendingSelection:"), record, extend)
}


// Selects a value or a set of values in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/selectIdentifier(_:for:byExtendingSelection:)
func (a_ ABPeoplePickerView) SelectIdentifierForPersonByExtendingSelection(identifier string /* primitive/slice/pointer. */, person IABPerson, extend bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("selectIdentifier:forPerson:byExtendingSelection:"), objc.String(identifier), person, extend)
}


// Launches Address Book and selects the item selected in the people picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/selectInAddressBook(_:)
func (a_ ABPeoplePickerView) SelectInAddressBook(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("selectInAddressBook:"), sender)
}


// Returns the identifiers of the selected values in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/selectedIdentifiers(for:)
func (a_ ABPeoplePickerView) SelectedIdentifiersForPerson(person IABPerson) objc.IObject /* cross-framework: Array */ {
	rv := objc.Send[Array](a_.ID, objc.Sel("selectedIdentifiersForPerson:"), person)
	return rv
}


// Returns an array of all the values selected in the displayed multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/selectedValues()
func (a_ ABPeoplePickerView) SelectedValues() objc.IObject /* cross-framework: Array */ {
	rv := objc.Send[Array](a_.ID, objc.Sel("selectedValues"))
	return rv
}


// Sets the title displayed in the people picker for a property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/setColumnTitle(_:forProperty:)
func (a_ ABPeoplePickerView) SetColumnTitleForProperty(title string /* primitive/slice/pointer. */, property string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setColumnTitle:forProperty:"), objc.String(title), objc.String(property))
}


// The view that is placed to the left of the search field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/accessoryView
func (a_ ABPeoplePickerView) AccessoryView() objc.IObject /* cross-framework: View */ {
	rv := objc.Send[View](a_.ID, objc.Sel("accessoryView"))
	return rv
}


// The view that is placed to the left of the search field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/accessoryView
func (a_ ABPeoplePickerView) SetAccessoryView(value objc.IObject /* cross-framework: View */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAccessoryView:"), value)
}


// A Boolean value that specifies whether the user can select entire groups in the group column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/allowsGroupSelection
func (a_ ABPeoplePickerView) AllowsGroupSelection() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsGroupSelection"))
	return rv
}


// A Boolean value that specifies whether the user can select entire groups in the group column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/allowsGroupSelection
func (a_ ABPeoplePickerView) SetAllowsGroupSelection(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsGroupSelection:"), value)
}


// A Boolean value that specifies whether multiple groups, records, or values of multivalue properties can be selected at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/allowsMultipleSelection
func (a_ ABPeoplePickerView) AllowsMultipleSelection() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}


// A Boolean value that specifies whether multiple groups, records, or values of multivalue properties can be selected at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/allowsMultipleSelection
func (a_ ABPeoplePickerView) SetAllowsMultipleSelection(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}


// The name under which the column positions and the filter selection are saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/autosaveName
func (a_ ABPeoplePickerView) AutosaveName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("autosaveName"))
	return rv
}


// The name under which the column positions and the filter selection are saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/autosaveName
func (a_ ABPeoplePickerView) SetAutosaveName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAutosaveName:"), objc.String(value))
}


// The property currently displayed in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/displayedProperty
func (a_ ABPeoplePickerView) DisplayedProperty() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("displayedProperty"))
	return rv
}


// The property currently displayed in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/displayedProperty
func (a_ ABPeoplePickerView) SetDisplayedProperty(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDisplayedProperty:"), objc.String(value))
}


// The action to be invoked when a group is double-clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/groupDoubleAction
func (a_ ABPeoplePickerView) GroupDoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](a_.ID, objc.Sel("groupDoubleAction"))
	return rv
}


// The action to be invoked when a group is double-clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/groupDoubleAction
func (a_ ABPeoplePickerView) SetGroupDoubleAction(value objc.SEL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setGroupDoubleAction:"), value)
}


// The action to be invoked when a name is double-clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/nameDoubleAction
func (a_ ABPeoplePickerView) NameDoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](a_.ID, objc.Sel("nameDoubleAction"))
	return rv
}


// The action to be invoked when a name is double-clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/nameDoubleAction
func (a_ ABPeoplePickerView) SetNameDoubleAction(value objc.SEL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNameDoubleAction:"), value)
}


// The groups selected in the group list. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/selectedGroups
func (a_ ABPeoplePickerView) SelectedGroups() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("selectedGroups"))
	return rv
}


// The selection in the records list. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/selectedRecords
func (a_ ABPeoplePickerView) SelectedRecords() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("selectedRecords"))
	return rv
}


// The target for double-click actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/target
func (a_ ABPeoplePickerView) Target() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("target"))
	return rv
}


// The target for double-click actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/target
func (a_ ABPeoplePickerView) SetTarget(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTarget:"), value)
}


// The current selection behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/valueSelectionBehavior
func (a_ ABPeoplePickerView) ValueSelectionBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("valueSelectionBehavior"))
	return rv
}


// The current selection behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/valueSelectionBehavior
func (a_ ABPeoplePickerView) SetValueSelectionBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setValueSelectionBehavior:"), value)
}



