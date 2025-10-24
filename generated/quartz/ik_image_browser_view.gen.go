// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class IKImageBrowserView */


/* debug [class_header]: Header for IKImageBrowserView */
// The class instance for the [IKImageBrowserView] class.
var (
	IKImageBrowserViewClass     _IKImageBrowserViewClass
	IKImageBrowserViewClassOnce sync.Once
)

func getIKImageBrowserViewClass() _IKImageBrowserViewClass {
	IKImageBrowserViewClassOnce.Do(func() {
		IKImageBrowserViewClass = _IKImageBrowserViewClass{objc.GetClass("IKImageBrowserView")}
	})
	return IKImageBrowserViewClass
}

type _IKImageBrowserViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IKImageBrowserView */
// An interface definition for the [IKImageBrowserView] class.
type IIKImageBrowserView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for IKImageBrowserView */
	// properties:
	DataSource() objc.ID
	SetDataSource(value objc.ID)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IKImageBrowserView */
	// methods:
	AllowsDroppingOnItems() bool
	AllowsEmptySelection() bool
	AllowsMultipleSelection() bool
	AllowsReordering() bool
	Animates() bool
	BackgroundLayer() avfoundation.Layer
	CanControlQuickLookPanel() bool
	CellForItemAtIndex(index uint) IKImageBrowserCell
	CellSize() Size /* not a class type */
	CellsStyleMask() uint
	CollapseGroupAtIndex(index uint)
	ColumnIndexesInRect(rect Rect /* not a class type */) foundation.IndexSet
	ConstrainsToOriginalSize() bool
	ContentResizingMask() uint
	DraggingDestinationDelegate() objc.ID
	DropOperation() objectivec.IObject
	ExpandGroupAtIndex(index uint)
	ForegroundLayer() avfoundation.Layer
	IndexAtLocationOfDroppedItem() uint
	IndexOfItemAtPoint(point vision.Point) int
	IntercellSpacing() Size /* not a class type */
	IsGroupExpandedAtIndex(index uint) bool
	ItemFrameAtIndex(index int) Rect /* not a class type */
	NewCellForRepresentedItem(anItem objc.IObject) IKImageBrowserCell
	NumberOfColumns() uint
	NumberOfRows() uint
	RectOfColumn(columnIndex uint) Rect /* not a class type */
	RectOfRow(rowIndex uint) Rect /* not a class type */
	ReloadData()
	RowIndexesInRect(rect Rect /* not a class type */) foundation.IndexSet
	ScrollIndexToVisible(index int)
	SelectionIndexes() foundation.IndexSet
	SetAllowsDroppingOnItems(flag bool)
	SetAllowsEmptySelection(flag bool)
	SetAllowsMultipleSelection(flag bool)
	SetAllowsReordering(flag bool)
	SetAnimates(flag bool)
	SetBackgroundLayer(aLayer avfoundation.Layer)
	SetCanControlQuickLookPanel(flag bool)
	SetCellSize(size Size /* not a class type */)
	SetCellsStyleMask(mask uint)
	SetConstrainsToOriginalSize(flag bool)
	SetContentResizingMask(mask uint)
	SetDraggingDestinationDelegate(delegate objc.IObject)
	SetDropIndexDropOperation(index int, operation objectivec.IObject)
	SetForegroundLayer(aLayer avfoundation.Layer)
	SetIntercellSpacing(aSize Size /* not a class type */)
	SetSelectionIndexesByExtendingSelection(indexes foundation.IndexSet, extendSelection bool)
	SetZoomValue(aValue float32)
	VisibleItemIndexes() foundation.IndexSet
	ZoomValue() float32
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IKImageBrowserView */
// Alloc allocates a new instance without initialization.
func (ic _IKImageBrowserViewClass) Alloc() IKImageBrowserView {
	rv := objc.Send[IKImageBrowserView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _IKImageBrowserViewClass) New() IKImageBrowserView {
	rv := objc.Send[IKImageBrowserView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKImageBrowserView) Init() IKImageBrowserView {
	rv := objc.Send[IKImageBrowserView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKImageBrowserView) Autorelease() IKImageBrowserView {
	rv := objc.Send[IKImageBrowserView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKImageBrowserView creates a new IKImageBrowserView instance.
func NewIKImageBrowserView() IKImageBrowserView {
	return getIKImageBrowserViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IKImageBrowserView */
// A view for displaying and browsing a large collection of images and movies.
//
// The class is a view for displaying and browsing a large amount of images and movies efficiently. This class will be deprecated in a future release. Please switch to instead. You must set a datasource for the view and implement, at a minimum, the and described in . The items must conform to the IKImageBrowserItem Protocol protocol. The class’s delegate object must conform to IKImageBrowserDelegate Protocol protocol. It receives notification of changes in selection, as well as mouse events in the cells.


// A view for displaying and browsing a large collection of images and movies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView
type IKImageBrowserView struct {
	appkit.View
}

// IKImageBrowserViewFrom constructs a [IKImageBrowserView] from an unsafe.Pointer.
//
// A view for displaying and browsing a large collection of images and movies.
func IKImageBrowserViewFrom(ptr unsafe.Pointer) IKImageBrowserView {
	return IKImageBrowserView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IKImageBrowserView */

// Initializes a newly allocated image browser view with the provided frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/initWithFrame:
func NewIKImageBrowserViewWithFrame(frame Rect /* not a class type */) IKImageBrowserView {
	instance := getIKImageBrowserViewClass().Alloc()
	rv := objc.Send[IKImageBrowserView](instance.ID, objc.Sel("initWithFrame:"), frame)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewIKImageBrowserViewWithFrame */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IKImageBrowserView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IKImageBrowserView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IKImageBrowserView */

// Returns whether the user can drop on items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/allowsDroppingOnItems()
func (i_ IKImageBrowserView) AllowsDroppingOnItems() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("allowsDroppingOnItems"))
	return rv
}/* debug [instance_methods/method]: AllowsDroppingOnItems */


// Returns whether an empty selection is allowed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/allowsEmptySelection()
func (i_ IKImageBrowserView) AllowsEmptySelection() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("allowsEmptySelection"))
	return rv
}/* debug [instance_methods/method]: AllowsEmptySelection */


// Returns whether multiple selections are allowed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/allowsMultipleSelection()
func (i_ IKImageBrowserView) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}/* debug [instance_methods/method]: AllowsMultipleSelection */


// Returns whether the user can reorder items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/allowsReordering()
func (i_ IKImageBrowserView) AllowsReordering() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("allowsReordering"))
	return rv
}/* debug [instance_methods/method]: AllowsReordering */


// Returns whether the receiver animates reordering and changes of the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/animates()
func (i_ IKImageBrowserView) Animates() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("animates"))
	return rv
}/* debug [instance_methods/method]: Animates */


// Returns the foreground Core Animation layer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/backgroundLayer()
func (i_ IKImageBrowserView) BackgroundLayer() avfoundation.Layer {
	rv := objc.Send[avfoundation.Layer](i_.ID, objc.Sel("backgroundLayer"))
	return rv
}/* debug [instance_methods/method]: BackgroundLayer */


// Returns whether the view can automatically take control of the QuickLook panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/canControlQuickLookPanel()
func (i_ IKImageBrowserView) CanControlQuickLookPanel() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("canControlQuickLookPanel"))
	return rv
}/* debug [instance_methods/method]: CanControlQuickLookPanel */


// Returns the browser cell for the item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/cellForItem(at:)
func (i_ IKImageBrowserView) CellForItemAtIndex(index uint) IKImageBrowserCell {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("cellForItemAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: CellForItemAtIndex */


// Returns the cell size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/cellSize()
func (i_ IKImageBrowserView) CellSize() Size /* not a class type */ {
	rv := objc.Send[Size](i_.ID, objc.Sel("cellSize"))
	return rv
}/* debug [instance_methods/method]: CellSize */


// Returns the appearance style mask for the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/cellsStyleMask()
func (i_ IKImageBrowserView) CellsStyleMask() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("cellsStyleMask"))
	return rv
}/* debug [instance_methods/method]: CellsStyleMask */


// Collapses a group at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/collapseGroup(at:)
func (i_ IKImageBrowserView) CollapseGroupAtIndex(index uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("collapseGroupAtIndex:"), index)
}/* debug [instance_methods/method]: CollapseGroupAtIndex */


// Returns the column indexes in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/columnIndexes(in:)
func (i_ IKImageBrowserView) ColumnIndexesInRect(rect Rect /* not a class type */) foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](i_.ID, objc.Sel("columnIndexesInRect:"), rect)
	return rv
}/* debug [instance_methods/method]: ColumnIndexesInRect */


// Returns whether the receiver constrains the cell’s image to its original size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/constrainsToOriginalSize()
func (i_ IKImageBrowserView) ConstrainsToOriginalSize() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("constrainsToOriginalSize"))
	return rv
}/* debug [instance_methods/method]: ConstrainsToOriginalSize */


// Returns the receiver’s content resizing mask, which determines how its content is resized while zooming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/contentResizingMask()
func (i_ IKImageBrowserView) ContentResizingMask() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("contentResizingMask"))
	return rv
}/* debug [instance_methods/method]: ContentResizingMask */


// Returns the dragging destination delegate of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/draggingDestinationDelegate()
func (i_ IKImageBrowserView) DraggingDestinationDelegate() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("draggingDestinationDelegate"))
	return rv
}/* debug [instance_methods/method]: DraggingDestinationDelegate */


// Returns the current drop operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/dropOperation()
func (i_ IKImageBrowserView) DropOperation() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("dropOperation"))
	return rv
}/* debug [instance_methods/method]: DropOperation */


// Expands a group at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/expandGroup(at:)
func (i_ IKImageBrowserView) ExpandGroupAtIndex(index uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("expandGroupAtIndex:"), index)
}/* debug [instance_methods/method]: ExpandGroupAtIndex */


// Returns the foreground Core Animation layer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/foregroundLayer()
func (i_ IKImageBrowserView) ForegroundLayer() avfoundation.Layer {
	rv := objc.Send[avfoundation.Layer](i_.ID, objc.Sel("foregroundLayer"))
	return rv
}/* debug [instance_methods/method]: ForegroundLayer */


// Returns the index of the cell where the drop operation occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/indexAtLocationOfDroppedItem()
func (i_ IKImageBrowserView) IndexAtLocationOfDroppedItem() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("indexAtLocationOfDroppedItem"))
	return rv
}/* debug [instance_methods/method]: IndexAtLocationOfDroppedItem */


// Returns the index of the item at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/indexOfItem(at:)
func (i_ IKImageBrowserView) IndexOfItemAtPoint(point vision.Point) int {
	rv := objc.Send[int](i_.ID, objc.Sel("indexOfItemAtPoint:"), point)
	return rv
}/* debug [instance_methods/method]: IndexOfItemAtPoint */


// Returns the spacing between cells in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/intercellSpacing()
func (i_ IKImageBrowserView) IntercellSpacing() Size /* not a class type */ {
	rv := objc.Send[Size](i_.ID, objc.Sel("intercellSpacing"))
	return rv
}/* debug [instance_methods/method]: IntercellSpacing */


// Returns whether the group at the provided index is expanded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/isGroupExpanded(at:)
func (i_ IKImageBrowserView) IsGroupExpandedAtIndex(index uint) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isGroupExpandedAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: IsGroupExpandedAtIndex */


// Returns the frame rectangle for the item located at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/itemFrame(at:)
func (i_ IKImageBrowserView) ItemFrameAtIndex(index int) Rect /* not a class type */ {
	rv := objc.Send[Rect](i_.ID, objc.Sel("itemFrameAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ItemFrameAtIndex */


// Returns the cell to use for the specified item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/newCell(forRepresentedItem:)
func (i_ IKImageBrowserView) NewCellForRepresentedItem(anItem objc.IObject) IKImageBrowserCell {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("newCellForRepresentedItem:"), anItem)
	return rv
}/* debug [instance_methods/method]: NewCellForRepresentedItem */


// Returns the current number of columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/numberOfColumns()
func (i_ IKImageBrowserView) NumberOfColumns() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("numberOfColumns"))
	return rv
}/* debug [instance_methods/method]: NumberOfColumns */


// Returns the current number of rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/numberOfRows()
func (i_ IKImageBrowserView) NumberOfRows() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("numberOfRows"))
	return rv
}/* debug [instance_methods/method]: NumberOfRows */


// Returns the rectangle containing the specified column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/rect(ofColumn:)
func (i_ IKImageBrowserView) RectOfColumn(columnIndex uint) Rect /* not a class type */ {
	rv := objc.Send[Rect](i_.ID, objc.Sel("rectOfColumn:"), columnIndex)
	return rv
}/* debug [instance_methods/method]: RectOfColumn */


// Returns the rectangle containing the specified row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/rect(ofRow:)
func (i_ IKImageBrowserView) RectOfRow(rowIndex uint) Rect /* not a class type */ {
	rv := objc.Send[Rect](i_.ID, objc.Sel("rectOfRow:"), rowIndex)
	return rv
}/* debug [instance_methods/method]: RectOfRow */


// Marks the receiver as needing its data reloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/reloadData()
func (i_ IKImageBrowserView) ReloadData() {
	objc.Send[objc.ID](i_.ID, objc.Sel("reloadData"))
}/* debug [instance_methods/method]: ReloadData */


// Returns the row indexes in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/rowIndexes(in:)
func (i_ IKImageBrowserView) RowIndexesInRect(rect Rect /* not a class type */) foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](i_.ID, objc.Sel("rowIndexesInRect:"), rect)
	return rv
}/* debug [instance_methods/method]: RowIndexesInRect */


// Scrolls the receiver to the item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/scrollIndexToVisible(_:)
func (i_ IKImageBrowserView) ScrollIndexToVisible(index int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("scrollIndexToVisible:"), index)
}/* debug [instance_methods/method]: ScrollIndexToVisible */


// Returns the indexes of the selected cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/selectionIndexes()
func (i_ IKImageBrowserView) SelectionIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](i_.ID, objc.Sel("selectionIndexes"))
	return rv
}/* debug [instance_methods/method]: SelectionIndexes */


// Specifies whether the user can drop on items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setAllowsDroppingOnItems(_:)
func (i_ IKImageBrowserView) SetAllowsDroppingOnItems(flag bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAllowsDroppingOnItems:"), flag)
}/* debug [instance_methods/method]: SetAllowsDroppingOnItems */


// Controls whether an empty selection is allowed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setAllowsEmptySelection(_:)
func (i_ IKImageBrowserView) SetAllowsEmptySelection(flag bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAllowsEmptySelection:"), flag)
}/* debug [instance_methods/method]: SetAllowsEmptySelection */


// Controls whether the user can select more than one cell at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setAllowsMultipleSelection(_:)
func (i_ IKImageBrowserView) SetAllowsMultipleSelection(flag bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAllowsMultipleSelection:"), flag)
}/* debug [instance_methods/method]: SetAllowsMultipleSelection */


// Controls whether the user can reorder items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setAllowsReordering(_:)
func (i_ IKImageBrowserView) SetAllowsReordering(flag bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAllowsReordering:"), flag)
}/* debug [instance_methods/method]: SetAllowsReordering */


// Controls whether the receiver animates reordering and changes of the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setAnimates(_:)
func (i_ IKImageBrowserView) SetAnimates(flag bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAnimates:"), flag)
}/* debug [instance_methods/method]: SetAnimates */


// The Core Animation layer used as the view’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setBackgroundLayer(_:)
func (i_ IKImageBrowserView) SetBackgroundLayer(aLayer avfoundation.Layer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBackgroundLayer:"), aLayer)
}/* debug [instance_methods/method]: SetBackgroundLayer */


// Specifies whether the view can automatically take control of the QuickLook panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setCanControlQuickLookPanel(_:)
func (i_ IKImageBrowserView) SetCanControlQuickLookPanel(flag bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCanControlQuickLookPanel:"), flag)
}/* debug [instance_methods/method]: SetCanControlQuickLookPanel */


// Sets the cell size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setCellSize(_:)
func (i_ IKImageBrowserView) SetCellSize(size Size /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCellSize:"), size)
}/* debug [instance_methods/method]: SetCellSize */


// Defines the appearance style of the cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setCellsStyleMask(_:)
func (i_ IKImageBrowserView) SetCellsStyleMask(mask uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCellsStyleMask:"), mask)
}/* debug [instance_methods/method]: SetCellsStyleMask */


// Sets whether the receiver constrains the cell’s image to its original size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setConstrainsToOriginalSize(_:)
func (i_ IKImageBrowserView) SetConstrainsToOriginalSize(flag bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setConstrainsToOriginalSize:"), flag)
}/* debug [instance_methods/method]: SetConstrainsToOriginalSize */


// Determines how the receiver resizes its content when zooming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setContentResizingMask(_:)
func (i_ IKImageBrowserView) SetContentResizingMask(mask uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContentResizingMask:"), mask)
}/* debug [instance_methods/method]: SetContentResizingMask */


// Sets the dragging destination delegate of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setDraggingDestinationDelegate(_:)
func (i_ IKImageBrowserView) SetDraggingDestinationDelegate(delegate objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDraggingDestinationDelegate:"), delegate)
}/* debug [instance_methods/method]: SetDraggingDestinationDelegate */


// Allows the class to retarget the drop action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setDrop(_:dropOperation:)
func (i_ IKImageBrowserView) SetDropIndexDropOperation(index int, operation objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDropIndex:dropOperation:"), index, operation)
}/* debug [instance_methods/method]: SetDropIndexDropOperation */


// The Core Animation layer used as the foreground overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setForegroundLayer(_:)
func (i_ IKImageBrowserView) SetForegroundLayer(aLayer avfoundation.Layer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setForegroundLayer:"), aLayer)
}/* debug [instance_methods/method]: SetForegroundLayer */


// Sets the spacing between cells in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setIntercellSpacing(_:)
func (i_ IKImageBrowserView) SetIntercellSpacing(aSize Size /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIntercellSpacing:"), aSize)
}/* debug [instance_methods/method]: SetIntercellSpacing */


// Selects cells at the specified indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setSelectionIndexes(_:byExtendingSelection:)
func (i_ IKImageBrowserView) SetSelectionIndexesByExtendingSelection(indexes foundation.IndexSet, extendSelection bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSelectionIndexes:byExtendingSelection:"), indexes, extendSelection)
}/* debug [instance_methods/method]: SetSelectionIndexesByExtendingSelection */


// Sets the zoom value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setZoomValue(_:)
func (i_ IKImageBrowserView) SetZoomValue(aValue float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setZoomValue:"), aValue)
}/* debug [instance_methods/method]: SetZoomValue */


// Returns the indexes of the view’s currently visible items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/visibleItemIndexes()
func (i_ IKImageBrowserView) VisibleItemIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](i_.ID, objc.Sel("visibleItemIndexes"))
	return rv
}/* debug [instance_methods/method]: VisibleItemIndexes */


// Returns the current zoom value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/zoomValue()
func (i_ IKImageBrowserView) ZoomValue() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("zoomValue"))
	return rv
}/* debug [instance_methods/method]: ZoomValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IKImageBrowserView */

// Returns the data source of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/dataSource
func (i_ IKImageBrowserView) DataSource() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// Returns the data source of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/dataSource
func (i_ IKImageBrowserView) SetDataSource(value objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDataSource:"), value)
}/* debug [instance_properties/setter]: dataSource */


// Returns the delegate of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/delegate
func (i_ IKImageBrowserView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// Returns the delegate of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/delegate
func (i_ IKImageBrowserView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IKImageBrowserView */


