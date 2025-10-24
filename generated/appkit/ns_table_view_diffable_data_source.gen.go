// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTableViewDiffableDataSource */


/* debug [class_header]: Header for NSTableViewDiffableDataSource */
// The class instance for the [TableViewDiffableDataSource] class.
var (
	TableViewDiffableDataSourceClass     _TableViewDiffableDataSourceClass
	TableViewDiffableDataSourceClassOnce sync.Once
)

func getTableViewDiffableDataSourceClass() _TableViewDiffableDataSourceClass {
	TableViewDiffableDataSourceClassOnce.Do(func() {
		TableViewDiffableDataSourceClass = _TableViewDiffableDataSourceClass{objc.GetClass("NSTableViewDiffableDataSource")}
	})
	return TableViewDiffableDataSourceClass
}

type _TableViewDiffableDataSourceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TableViewDiffableDataSource */
// An interface definition for the [TableViewDiffableDataSource] class.
type ITableViewDiffableDataSource interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TableViewDiffableDataSource */
	// properties:
	DefaultRowAnimation() TableViewAnimationOptions
	SetDefaultRowAnimation(value TableViewAnimationOptions)
	RowViewProvider() TableViewDiffableDataSourceRowProvider /* not a class type */
	SetRowViewProvider(value TableViewDiffableDataSourceRowProvider /* not a class type */)
	SectionHeaderViewProvider() TableViewDiffableDataSourceSectionHeaderViewProvider /* not a class type */
	SetSectionHeaderViewProvider(value TableViewDiffableDataSourceSectionHeaderViewProvider /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TableViewDiffableDataSource */
	// methods:
	ApplySnapshotAnimatingDifferences(snapshot unsafe.Pointer, animatingDifferences bool)
	ApplySnapshotAnimatingDifferencesCompletion(snapshot unsafe.Pointer, animatingDifferences bool, completion unsafe.Pointer)
	ItemIdentifierForRow(row int) objectivec.IObject
	RowForItemIdentifier(identifier objectivec.IObject) int
	RowForSectionIdentifier(identifier objectivec.IObject) int
	SectionIdentifierForRow(row int) objectivec.IObject
	Snapshot() unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TableViewDiffableDataSource */
// Alloc allocates a new instance without initialization.
func (tc _TableViewDiffableDataSourceClass) Alloc() TableViewDiffableDataSource {
	rv := objc.Send[TableViewDiffableDataSource](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TableViewDiffableDataSourceClass) New() TableViewDiffableDataSource {
	rv := objc.Send[TableViewDiffableDataSource](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TableViewDiffableDataSource) Init() TableViewDiffableDataSource {
	rv := objc.Send[TableViewDiffableDataSource](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TableViewDiffableDataSource) Autorelease() TableViewDiffableDataSource {
	rv := objc.Send[TableViewDiffableDataSource](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTableViewDiffableDataSource creates a new TableViewDiffableDataSource instance.
func NewTableViewDiffableDataSource() TableViewDiffableDataSource {
	return getTableViewDiffableDataSourceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TableViewDiffableDataSource */
// The object you use to manage data and provide items for a table view.
//
// A object is a specialized type of data source that works together with your table view object. It provides the behavior you need to manage updates to your table view’s data and UI in a simple, efficient way. It also conforms to the protocol. To fill a table view with data: Connect a diffable data source to your table view. Implement a cell provider to configure your table view’s cells. Generate the current state of the data. Display the data in the UI. To connect a diffable data source to a table view, you create the diffable data source using its initializer, passing in the table view you want to associate with that data source. You also pass in a cell provider, where you configure each of your cells to determine how to display your data in the UI. Then, you generate the current state of the data and display the data in the UI by constructing and applying a snapshot. For more information, see .


// The object you use to manage data and provide items for a table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference
type TableViewDiffableDataSource struct {
	objectivec.Object
}

// TableViewDiffableDataSourceFrom constructs a [TableViewDiffableDataSource] from an unsafe.Pointer.
//
// The object you use to manage data and provide items for a table view.
func TableViewDiffableDataSourceFrom(ptr unsafe.Pointer) TableViewDiffableDataSource {
	return TableViewDiffableDataSource{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TableViewDiffableDataSource */

// Creates a diffable data source with the specified cell provider, and connects it to the specified table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/init(tableView:cellProvider:)
func NewTableViewDiffableDataSourceWithTableViewCellProvider(tableView ITableView, cellProvider TableViewDiffableDataSourceCellProvider /* not a class type */) TableViewDiffableDataSource {
	instance := getTableViewDiffableDataSourceClass().Alloc()
	rv := objc.Send[TableViewDiffableDataSource](instance.ID, objc.Sel("initWithTableView:cellProvider:"), tableView, cellProvider)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTableViewDiffableDataSourceWithTableViewCellProvider */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TableViewDiffableDataSource */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TableViewDiffableDataSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TableViewDiffableDataSource */

// Updates the UI to reflect the state of the data in the specified snapshot, optionally animating the UI changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/applySnapshot(_:animatingDifferences:)
func (t_ TableViewDiffableDataSource) ApplySnapshotAnimatingDifferences(snapshot unsafe.Pointer, animatingDifferences bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("applySnapshot:animatingDifferences:"), snapshot, animatingDifferences)
}/* debug [instance_methods/method]: ApplySnapshotAnimatingDifferences */


// Updates the UI to reflect the state of the data in the specified snapshot, optionally animating the UI changes and executing a completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/applySnapshot(_:animatingDifferences:completion:)
func (t_ TableViewDiffableDataSource) ApplySnapshotAnimatingDifferencesCompletion(snapshot unsafe.Pointer, animatingDifferences bool, completion unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("applySnapshot:animatingDifferences:completion:"), snapshot, animatingDifferences, completion)
}/* debug [instance_methods/method]: ApplySnapshotAnimatingDifferencesCompletion */


// Returns an identifier for the item at the specified row in the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/itemIdentifier(forRow:)
func (t_ TableViewDiffableDataSource) ItemIdentifierForRow(row int) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("itemIdentifierForRow:"), row)
	return rv
}/* debug [instance_methods/method]: ItemIdentifierForRow */


// Returns a row for the item with the specified identifier in the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/row(forItemIdentifier:)
func (t_ TableViewDiffableDataSource) RowForItemIdentifier(identifier objectivec.IObject) int {
	rv := objc.Send[int](t_.ID, objc.Sel("rowForItemIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: RowForItemIdentifier */


// Returns a row for the section with the specified identifier in the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/row(forSectionIdentifier:)
func (t_ TableViewDiffableDataSource) RowForSectionIdentifier(identifier objectivec.IObject) int {
	rv := objc.Send[int](t_.ID, objc.Sel("rowForSectionIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: RowForSectionIdentifier */


// Returns the identifier of the section containing the specified row in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/sectionIdentifier(forRow:)
func (t_ TableViewDiffableDataSource) SectionIdentifierForRow(row int) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("sectionIdentifierForRow:"), row)
	return rv
}/* debug [instance_methods/method]: SectionIdentifierForRow */


// Returns a representation of the current state of the data in the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/snapshot()
func (t_ TableViewDiffableDataSource) Snapshot() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("snapshot"))
	return rv
}/* debug [instance_methods/method]: Snapshot */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TableViewDiffableDataSource */

// The default animation the UI uses to show differences between rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/defaultRowAnimation
func (t_ TableViewDiffableDataSource) DefaultRowAnimation() TableViewAnimationOptions {
	rv := objc.Send[TableViewAnimationOptions](t_.ID, objc.Sel("defaultRowAnimation"))
	return rv
}/* debug [instance_properties/getter]: defaultRowAnimation */


// The default animation the UI uses to show differences between rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/defaultRowAnimation
func (t_ TableViewDiffableDataSource) SetDefaultRowAnimation(value TableViewAnimationOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDefaultRowAnimation:"), value)
}/* debug [instance_properties/setter]: defaultRowAnimation */


// The closure that configures and returns the table view’s row views from the diffable data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/rowViewProvider
func (t_ TableViewDiffableDataSource) RowViewProvider() TableViewDiffableDataSourceRowProvider /* not a class type */ {
	rv := objc.Send[TableViewDiffableDataSourceRowProvider](t_.ID, objc.Sel("rowViewProvider"))
	return rv
}/* debug [instance_properties/getter]: rowViewProvider */


// The closure that configures and returns the table view’s row views from the diffable data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/rowViewProvider
func (t_ TableViewDiffableDataSource) SetRowViewProvider(value TableViewDiffableDataSourceRowProvider /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRowViewProvider:"), value)
}/* debug [instance_properties/setter]: rowViewProvider */


// The closure that configures and returns the table view’s section header views from the diffable data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/sectionHeaderViewProvider
func (t_ TableViewDiffableDataSource) SectionHeaderViewProvider() TableViewDiffableDataSourceSectionHeaderViewProvider /* not a class type */ {
	rv := objc.Send[TableViewDiffableDataSourceSectionHeaderViewProvider](t_.ID, objc.Sel("sectionHeaderViewProvider"))
	return rv
}/* debug [instance_properties/getter]: sectionHeaderViewProvider */


// The closure that configures and returns the table view’s section header views from the diffable data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/sectionHeaderViewProvider
func (t_ TableViewDiffableDataSource) SetSectionHeaderViewProvider(value TableViewDiffableDataSourceSectionHeaderViewProvider /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSectionHeaderViewProvider:"), value)
}/* debug [instance_properties/setter]: sectionHeaderViewProvider */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTableViewDiffableDataSource */


