// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PFontChanging is the NSFontChanging protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSFontChanging
type PFontChanging interface {
	// Optional methods
	ChangeFont(sender IFontManager)
	HasChangeFont() bool
	ValidModesForFontPanel(fontPanel IFontPanel) FontPanelModeMask
	HasValidModesForFontPanel() bool
}
