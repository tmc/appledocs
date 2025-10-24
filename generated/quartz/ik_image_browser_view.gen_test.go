// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz_test

import (
	"github.com/tmc/appledocs/generated/quartz"
)

// Suppress unused import errors
var _ = quartz.NewIKImageBrowserView

// ExampleNewIKImageBrowserViewWithFrame demonstrates how to create a IKImageBrowserView instance using NewIKImageBrowserViewWithFrame.
// Initializes a newly allocated image browser view with the provided frame rectangle.
func ExampleNewIKImageBrowserViewWithFrame() {
	_ = quartz.NewIKImageBrowserViewWithFrame(
		quartz.Rect /* not a class type */{}, // frame Rect /* not a class type */
	)
	// Output:
}
// ExampleIKImageBrowserView_AllowsDroppingOnItems demonstrates using AllowsDroppingOnItems on a IKImageBrowserView instance.
// Returns whether the user can drop on items.
func ExampleIKImageBrowserView_AllowsDroppingOnItems() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.AllowsDroppingOnItems()
	// Output:
	}

// ExampleIKImageBrowserView_AllowsEmptySelection demonstrates using AllowsEmptySelection on a IKImageBrowserView instance.
// Returns whether an empty selection is allowed.
func ExampleIKImageBrowserView_AllowsEmptySelection() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.AllowsEmptySelection()
	// Output:
	}

// ExampleIKImageBrowserView_AllowsMultipleSelection demonstrates using AllowsMultipleSelection on a IKImageBrowserView instance.
// Returns whether multiple selections are allowed.
func ExampleIKImageBrowserView_AllowsMultipleSelection() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.AllowsMultipleSelection()
	// Output:
	}

// ExampleIKImageBrowserView_AllowsReordering demonstrates using AllowsReordering on a IKImageBrowserView instance.
// Returns whether the user can reorder items.
func ExampleIKImageBrowserView_AllowsReordering() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.AllowsReordering()
	// Output:
	}

// ExampleIKImageBrowserView_Animates demonstrates using Animates on a IKImageBrowserView instance.
// Returns whether the receiver animates reordering and changes of the data source.
func ExampleIKImageBrowserView_Animates() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.Animates()
	// Output:
	}

// ExampleIKImageBrowserView_BackgroundLayer demonstrates using BackgroundLayer on a IKImageBrowserView instance.
// Returns the foreground Core Animation layer
func ExampleIKImageBrowserView_BackgroundLayer() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.BackgroundLayer()
	// Output:
	}

// ExampleIKImageBrowserView_CanControlQuickLookPanel demonstrates using CanControlQuickLookPanel on a IKImageBrowserView instance.
// Returns whether the view can automatically take control of the QuickLook panel.
func ExampleIKImageBrowserView_CanControlQuickLookPanel() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.CanControlQuickLookPanel()
	// Output:
	}

// ExampleIKImageBrowserView_CellSize demonstrates using CellSize on a IKImageBrowserView instance.
// Returns the cell size.
func ExampleIKImageBrowserView_CellSize() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.CellSize()
	// Output:
	}

// ExampleIKImageBrowserView_CellsStyleMask demonstrates using CellsStyleMask on a IKImageBrowserView instance.
// Returns the appearance style mask for the cell.
func ExampleIKImageBrowserView_CellsStyleMask() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.CellsStyleMask()
	// Output:
	}

// ExampleIKImageBrowserView_ConstrainsToOriginalSize demonstrates using ConstrainsToOriginalSize on a IKImageBrowserView instance.
// Returns whether the receiver constrains the cell’s image to its original size.
func ExampleIKImageBrowserView_ConstrainsToOriginalSize() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.ConstrainsToOriginalSize()
	// Output:
	}

// ExampleIKImageBrowserView_ContentResizingMask demonstrates using ContentResizingMask on a IKImageBrowserView instance.
// Returns the receiver’s content resizing mask, which determines how its content is resized while zooming.
func ExampleIKImageBrowserView_ContentResizingMask() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.ContentResizingMask()
	// Output:
	}

// ExampleIKImageBrowserView_DraggingDestinationDelegate demonstrates using DraggingDestinationDelegate on a IKImageBrowserView instance.
// Returns the dragging destination delegate of the receiver.
func ExampleIKImageBrowserView_DraggingDestinationDelegate() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.DraggingDestinationDelegate()
	// Output:
	}

// ExampleIKImageBrowserView_DropOperation demonstrates using DropOperation on a IKImageBrowserView instance.
// Returns the current drop operation.
func ExampleIKImageBrowserView_DropOperation() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.DropOperation()
	// Output:
	}

// ExampleIKImageBrowserView_ForegroundLayer demonstrates using ForegroundLayer on a IKImageBrowserView instance.
// Returns the foreground Core Animation layer
func ExampleIKImageBrowserView_ForegroundLayer() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.ForegroundLayer()
	// Output:
	}

// ExampleIKImageBrowserView_IndexAtLocationOfDroppedItem demonstrates using IndexAtLocationOfDroppedItem on a IKImageBrowserView instance.
// Returns the index of the cell where the drop operation occurred.
func ExampleIKImageBrowserView_IndexAtLocationOfDroppedItem() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.IndexAtLocationOfDroppedItem()
	// Output:
	}

// ExampleIKImageBrowserView_IntercellSpacing demonstrates using IntercellSpacing on a IKImageBrowserView instance.
// Returns the spacing between cells in the view.
func ExampleIKImageBrowserView_IntercellSpacing() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.IntercellSpacing()
	// Output:
	}

// ExampleIKImageBrowserView_NumberOfColumns demonstrates using NumberOfColumns on a IKImageBrowserView instance.
// Returns the current number of columns.
func ExampleIKImageBrowserView_NumberOfColumns() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.NumberOfColumns()
	// Output:
	}

// ExampleIKImageBrowserView_NumberOfRows demonstrates using NumberOfRows on a IKImageBrowserView instance.
// Returns the current number of rows.
func ExampleIKImageBrowserView_NumberOfRows() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.NumberOfRows()
	// Output:
	}

// ExampleIKImageBrowserView_ReloadData demonstrates using ReloadData on a IKImageBrowserView instance.
// Marks the receiver as needing its data reloaded.
func ExampleIKImageBrowserView_ReloadData() {
	obj := quartz.NewIKImageBrowserView()
	obj.ReloadData()
	// Output:
	}

// ExampleIKImageBrowserView_SelectionIndexes demonstrates using SelectionIndexes on a IKImageBrowserView instance.
// Returns the indexes of the selected cells.
func ExampleIKImageBrowserView_SelectionIndexes() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.SelectionIndexes()
	// Output:
	}

// ExampleIKImageBrowserView_VisibleItemIndexes demonstrates using VisibleItemIndexes on a IKImageBrowserView instance.
// Returns the indexes of the view’s currently visible items.
func ExampleIKImageBrowserView_VisibleItemIndexes() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.VisibleItemIndexes()
	// Output:
	}

// ExampleIKImageBrowserView_ZoomValue demonstrates using ZoomValue on a IKImageBrowserView instance.
// Returns the current zoom value.
func ExampleIKImageBrowserView_ZoomValue() {
	obj := quartz.NewIKImageBrowserView()
	_ = obj.ZoomValue()
	// Output:
	}

