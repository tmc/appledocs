// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ABPeoplePickerView */


/* debug [class_header]: Header for ABPeoplePickerView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ABPeoplePickerView */
// An interface definition for the [ABPeoplePickerView] class.
type IABPeoplePickerView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for ABPeoplePickerView */
	// properties:
	AccessoryView() appkit.View
	SetAccessoryView(value appkit.View)
	AllowsGroupSelection() bool
	SetAllowsGroupSelection(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	AutosaveName() objc.IObject /* cross-framework: NSString */
	SetAutosaveName(value objc.IObject /* cross-framework: NSString */)
	DisplayedProperty() objc.IObject /* cross-framework: NSString */
	SetDisplayedProperty(value objc.IObject /* cross-framework: NSString */)
	GroupDoubleAction() objc.SEL
	SetGroupDoubleAction(value objc.SEL)
	NameDoubleAction() objc.SEL
	SetNameDoubleAction(value objc.SEL)
	SelectedGroups() objc.IObject /* cross-framework: NSArray */
	SelectedRecords() objc.IObject /* cross-framework: NSArray */
	Target() objc.ID
	SetTarget(value objc.ID)
	ValueSelectionBehavior() unsafe.Pointer
	SetValueSelectionBehavior(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ABPeoplePickerView */
	// methods:
	AddProperty(property objc.IObject /* cross-framework: NSString */)
	ClearSearchField(sender objc.IObject)
	ColumnTitleForProperty(property objc.IObject /* cross-framework: NSString */) foundation.String
	DeselectRecord(record IABRecord)
	DeselectGroup(group IABGroup)
	DeselectAll(sender objc.IObject)
	DeselectIdentifierForPerson(identifier objc.IObject /* cross-framework: NSString */, person IABPerson)
	EditInAddressBook(sender objc.IObject)
	Properties() foundation.Array
	RemoveProperty(property objc.IObject /* cross-framework: NSString */)
	SelectGroupByExtendingSelection(group IABGroup, extend bool)
	SelectRecordByExtendingSelection(record IABRecord, extend bool)
	SelectedIdentifiersForPerson(person IABPerson) foundation.Array
	SelectedValues() foundation.Array
	SelectIdentifierForPersonByExtendingSelection(identifier objc.IObject /* cross-framework: NSString */, person IABPerson, extend bool)
	SelectInAddressBook(sender objc.IObject)
	SetColumnTitleForProperty(title objc.IObject /* cross-framework: NSString */, property objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ABPeoplePickerView */
// Alloc allocates a new instance without initialization.
func (ac _ABPeoplePickerViewClass) Alloc() ABPeoplePickerView {
	rv := objc.Send[ABPeoplePickerView](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ABPeoplePickerView */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ABPeoplePickerView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ABPeoplePickerView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ABPeoplePickerView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ABPeoplePickerView */

// Adds a property to the group of properties whose values are shown in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/addProperty(_:)
func (a_ ABPeoplePickerView) AddProperty(property objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addProperty:"), property)
}/* debug [instance_methods/method]: AddProperty */


// Clears the search field and resets the list of displayed records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/clearSearchField(_:)
func (a_ ABPeoplePickerView) ClearSearchField(sender objc.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("clearSearchField:"), sender)
}/* debug [instance_methods/method]: ClearSearchField */


// Returns the title of a custom property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/columnTitle(forProperty:)
func (a_ ABPeoplePickerView) ColumnTitleForProperty(property objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("columnTitleForProperty:"), property)
	return rv
}/* debug [instance_methods/method]: ColumnTitleForProperty */


// Deselects a record selected in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/deselect(_:)-1yy11
func (a_ ABPeoplePickerView) DeselectRecord(record IABRecord) {
	objc.Send[objc.ID](a_.ID, objc.Sel("deselectRecord:"), record)
}/* debug [instance_methods/method]: DeselectRecord */


// Deselects a group selected in the group list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/deselect(_:)-3x7tl
func (a_ ABPeoplePickerView) DeselectGroup(group IABGroup) {
	objc.Send[objc.ID](a_.ID, objc.Sel("deselectGroup:"), group)
}/* debug [instance_methods/method]: DeselectGroup */


// Deselects all selected groups, records, and values in multivalue properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/deselectAll(_:)
func (a_ ABPeoplePickerView) DeselectAll(sender objc.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("deselectAll:"), sender)
}/* debug [instance_methods/method]: DeselectAll */


// Deselects a value selected in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/deselectIdentifier(_:for:)
func (a_ ABPeoplePickerView) DeselectIdentifierForPerson(identifier objc.IObject /* cross-framework: NSString */, person IABPerson) {
	objc.Send[objc.ID](a_.ID, objc.Sel("deselectIdentifier:forPerson:"), identifier, person)
}/* debug [instance_methods/method]: DeselectIdentifierForPerson */


// Launches Address Book to edit the item selected in the people picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/editInAddressBook(_:)
func (a_ ABPeoplePickerView) EditInAddressBook(sender objc.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("editInAddressBook:"), sender)
}/* debug [instance_methods/method]: EditInAddressBook */


// Returns an array of the properties whose values are shown in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/properties()
func (a_ ABPeoplePickerView) Properties() foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("properties"))
	return rv
}/* debug [instance_methods/method]: Properties */


// Removes a property from the group of properties whose values are shown in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/removeProperty(_:)
func (a_ ABPeoplePickerView) RemoveProperty(property objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeProperty:"), property)
}/* debug [instance_methods/method]: RemoveProperty */


// Selects a group or a set of groups in the group list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/select(_:byExtendingSelection:)-6mrii
func (a_ ABPeoplePickerView) SelectGroupByExtendingSelection(group IABGroup, extend bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("selectGroup:byExtendingSelection:"), group, extend)
}/* debug [instance_methods/method]: SelectGroupByExtendingSelection */


// Selects a record or a set of records in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/select(_:byExtendingSelection:)-9eldk
func (a_ ABPeoplePickerView) SelectRecordByExtendingSelection(record IABRecord, extend bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("selectRecord:byExtendingSelection:"), record, extend)
}/* debug [instance_methods/method]: SelectRecordByExtendingSelection */


// Returns the identifiers of the selected values in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/selectedIdentifiers(for:)
func (a_ ABPeoplePickerView) SelectedIdentifiersForPerson(person IABPerson) foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("selectedIdentifiersForPerson:"), person)
	return rv
}/* debug [instance_methods/method]: SelectedIdentifiersForPerson */


// Returns an array of all the values selected in the displayed multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/selectedValues()
func (a_ ABPeoplePickerView) SelectedValues() foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("selectedValues"))
	return rv
}/* debug [instance_methods/method]: SelectedValues */


// Selects a value or a set of values in a multivalue property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/selectIdentifier(_:for:byExtendingSelection:)
func (a_ ABPeoplePickerView) SelectIdentifierForPersonByExtendingSelection(identifier objc.IObject /* cross-framework: NSString */, person IABPerson, extend bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("selectIdentifier:forPerson:byExtendingSelection:"), identifier, person, extend)
}/* debug [instance_methods/method]: SelectIdentifierForPersonByExtendingSelection */


// Launches Address Book and selects the item selected in the people picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/selectInAddressBook(_:)
func (a_ ABPeoplePickerView) SelectInAddressBook(sender objc.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("selectInAddressBook:"), sender)
}/* debug [instance_methods/method]: SelectInAddressBook */


// Sets the title displayed in the people picker for a property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/setColumnTitle(_:forProperty:)
func (a_ ABPeoplePickerView) SetColumnTitleForProperty(title objc.IObject /* cross-framework: NSString */, property objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setColumnTitle:forProperty:"), title, property)
}/* debug [instance_methods/method]: SetColumnTitleForProperty */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ABPeoplePickerView */

// The view that is placed to the left of the search field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/accessoryView
func (a_ ABPeoplePickerView) AccessoryView() appkit.View {
	rv := objc.Send[appkit.View](a_.ID, objc.Sel("accessoryView"))
	return rv
}/* debug [instance_properties/getter]: accessoryView */


// The view that is placed to the left of the search field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/accessoryView
func (a_ ABPeoplePickerView) SetAccessoryView(value appkit.View) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAccessoryView:"), value)
}/* debug [instance_properties/setter]: accessoryView */


// A Boolean value that specifies whether the user can select entire groups in the group column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/allowsGroupSelection
func (a_ ABPeoplePickerView) AllowsGroupSelection() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsGroupSelection"))
	return rv
}/* debug [instance_properties/getter]: allowsGroupSelection */


// A Boolean value that specifies whether the user can select entire groups in the group column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/allowsGroupSelection
func (a_ ABPeoplePickerView) SetAllowsGroupSelection(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsGroupSelection:"), value)
}/* debug [instance_properties/setter]: allowsGroupSelection */


// A Boolean value that specifies whether multiple groups, records, or values of multivalue properties can be selected at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/allowsMultipleSelection
func (a_ ABPeoplePickerView) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}/* debug [instance_properties/getter]: allowsMultipleSelection */


// A Boolean value that specifies whether multiple groups, records, or values of multivalue properties can be selected at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/allowsMultipleSelection
func (a_ ABPeoplePickerView) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}/* debug [instance_properties/setter]: allowsMultipleSelection */


// The name under which the column positions and the filter selection are saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/autosaveName
func (a_ ABPeoplePickerView) AutosaveName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("autosaveName"))
	return rv
}/* debug [instance_properties/getter]: autosaveName */


// The name under which the column positions and the filter selection are saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/autosaveName
func (a_ ABPeoplePickerView) SetAutosaveName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAutosaveName:"), value)
}/* debug [instance_properties/setter]: autosaveName */


// The property currently displayed in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/displayedProperty
func (a_ ABPeoplePickerView) DisplayedProperty() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("displayedProperty"))
	return rv
}/* debug [instance_properties/getter]: displayedProperty */


// The property currently displayed in the record list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/displayedProperty
func (a_ ABPeoplePickerView) SetDisplayedProperty(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDisplayedProperty:"), value)
}/* debug [instance_properties/setter]: displayedProperty */


// The action to be invoked when a group is double-clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/groupDoubleAction
func (a_ ABPeoplePickerView) GroupDoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](a_.ID, objc.Sel("groupDoubleAction"))
	return rv
}/* debug [instance_properties/getter]: groupDoubleAction */


// The action to be invoked when a group is double-clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/groupDoubleAction
func (a_ ABPeoplePickerView) SetGroupDoubleAction(value objc.SEL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setGroupDoubleAction:"), value)
}/* debug [instance_properties/setter]: groupDoubleAction */


// The action to be invoked when a name is double-clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/nameDoubleAction
func (a_ ABPeoplePickerView) NameDoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](a_.ID, objc.Sel("nameDoubleAction"))
	return rv
}/* debug [instance_properties/getter]: nameDoubleAction */


// The action to be invoked when a name is double-clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/nameDoubleAction
func (a_ ABPeoplePickerView) SetNameDoubleAction(value objc.SEL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNameDoubleAction:"), value)
}/* debug [instance_properties/setter]: nameDoubleAction */


// The groups selected in the group list. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/selectedGroups
func (a_ ABPeoplePickerView) SelectedGroups() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](a_.ID, objc.Sel("selectedGroups"))
	return rv
}/* debug [instance_properties/getter]: selectedGroups */


// The selection in the records list. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/selectedRecords
func (a_ ABPeoplePickerView) SelectedRecords() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](a_.ID, objc.Sel("selectedRecords"))
	return rv
}/* debug [instance_properties/getter]: selectedRecords */


// The target for double-click actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/target
func (a_ ABPeoplePickerView) Target() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */


// The target for double-click actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/target
func (a_ ABPeoplePickerView) SetTarget(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTarget:"), value)
}/* debug [instance_properties/setter]: target */


// The current selection behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/valueSelectionBehavior
func (a_ ABPeoplePickerView) ValueSelectionBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("valueSelectionBehavior"))
	return rv
}/* debug [instance_properties/getter]: valueSelectionBehavior */


// The current selection behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABPeoplePickerView/valueSelectionBehavior
func (a_ ABPeoplePickerView) SetValueSelectionBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setValueSelectionBehavior:"), value)
}/* debug [instance_properties/setter]: valueSelectionBehavior */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ABPeoplePickerView */



