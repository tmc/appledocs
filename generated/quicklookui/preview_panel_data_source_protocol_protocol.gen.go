// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PPreviewPanelDataSource is the QLPreviewPanelDataSource protocol interface.
//
// A protocol that the Quick Look preview panel uses to access the contents of its data source object.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.quicklookui/documentation/QuickLookUI/QLPreviewPanelDataSource
type PPreviewPanelDataSource interface {
	// Required methods
	NumberOfPreviewItemsInPreviewPanel(panel IQLPreviewPanel) int/* debug [protocol_interface/required_method]: NumberOfPreviewItemsInPreviewPanel */
	PreviewPanelPreviewItemAtIndex(panel IQLPreviewPanel, index int) unsafe.Pointer/* debug [protocol_interface/required_method]: PreviewPanelPreviewItemAtIndex */
}

// PreviewPanelDataSource is a delegate implementation builder for the PPreviewPanelDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PreviewPanelDataSource struct {
	_NumberOfPreviewItemsInPreviewPanel func(panel IQLPreviewPanel) int
	_PreviewPanelPreviewItemAtIndex func(panel IQLPreviewPanel, index int) unsafe.Pointer
}

// SetNumberOfPreviewItemsInPreviewPanel sets the handler for the NumberOfPreviewItemsInPreviewPanel delegate method.
//
// Returns the number of items that the preview panel should preview.
func (d *PreviewPanelDataSource) SetNumberOfPreviewItemsInPreviewPanel(f func(panel IQLPreviewPanel) int) {
	d._NumberOfPreviewItemsInPreviewPanel = f
}

// SetPreviewPanelPreviewItemAtIndex sets the handler for the PreviewPanelPreviewItemAtIndex delegate method.
//
// Returns the item that the preview panel should preview at a given index.
func (d *PreviewPanelDataSource) SetPreviewPanelPreviewItemAtIndex(f func(panel IQLPreviewPanel, index int) unsafe.Pointer) {
	d._PreviewPanelPreviewItemAtIndex = f
}

// NumberOfPreviewItemsInPreviewPanel implements the PPreviewPanelDataSource interface.
func (d *PreviewPanelDataSource) NumberOfPreviewItemsInPreviewPanel(panel IQLPreviewPanel) int {
	if d._NumberOfPreviewItemsInPreviewPanel != nil {
		return d._NumberOfPreviewItemsInPreviewPanel(panel)
	}
	var zero int
	return zero
}

// HasNumberOfPreviewItemsInPreviewPanel returns true if a handler for NumberOfPreviewItemsInPreviewPanel has been set.
func (d *PreviewPanelDataSource) HasNumberOfPreviewItemsInPreviewPanel() bool {
	return d._NumberOfPreviewItemsInPreviewPanel != nil
}

// PreviewPanelPreviewItemAtIndex implements the PPreviewPanelDataSource interface.
func (d *PreviewPanelDataSource) PreviewPanelPreviewItemAtIndex(panel IQLPreviewPanel, index int) unsafe.Pointer {
	if d._PreviewPanelPreviewItemAtIndex != nil {
		return d._PreviewPanelPreviewItemAtIndex(panel, index)
	}
	var zero unsafe.Pointer
	return zero
}

// HasPreviewPanelPreviewItemAtIndex returns true if a handler for PreviewPanelPreviewItemAtIndex has been set.
func (d *PreviewPanelDataSource) HasPreviewPanelPreviewItemAtIndex() bool {
	return d._PreviewPanelPreviewItemAtIndex != nil
}
