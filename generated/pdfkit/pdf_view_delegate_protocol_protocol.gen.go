// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/appkit"

	"github.com/tmc/appledocs/generated/foundation"
)

// PPDFViewDelegate is the PDFViewDelegate protocol interface.
//
// The delegate for the   object.
//
// Availability:
//   - Mac Catalyst 11.0+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.4+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.pdfkit/documentation/PDFKit/PDFViewDelegate
type PPDFViewDelegate interface {
	// Optional methods
	PDFViewOpenPDFForRemoteGoToAction(sender IPDFView, action IPDFActionRemoteGoTo)
	HasPDFViewOpenPDFForRemoteGoToAction() bool
	PDFViewParentViewController() appkit.ViewController
	HasPDFViewParentViewController() bool
	PDFViewPerformFind(sender IPDFView)
	HasPDFViewPerformFind() bool
	PDFViewPerformGoToPage(sender IPDFView)
	HasPDFViewPerformGoToPage() bool
	PDFViewPerformPrint(sender IPDFView)
	HasPDFViewPerformPrint() bool
	PDFViewPrintJobTitle(sender IPDFView) foundation.String
	HasPDFViewPrintJobTitle() bool
	PDFViewWillChangeScaleFactorToScale(sender IPDFView, scaler float64) float64
	HasPDFViewWillChangeScaleFactorToScale() bool
	PDFViewWillClickOnLinkWithURL(sender IPDFView, url objc.IObject /* cross-framework: NSURL */)
	HasPDFViewWillClickOnLinkWithURL() bool
}

// PDFViewDelegate is a delegate implementation builder for the PPDFViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PDFViewDelegate struct {
	_PDFViewOpenPDFForRemoteGoToAction func(sender IPDFView, action IPDFActionRemoteGoTo)
	_PDFViewParentViewController func() appkit.ViewController
	_PDFViewPerformFind func(sender IPDFView)
	_PDFViewPerformGoToPage func(sender IPDFView)
	_PDFViewPerformPrint func(sender IPDFView)
	_PDFViewPrintJobTitle func(sender IPDFView) foundation.String
	_PDFViewWillChangeScaleFactorToScale func(sender IPDFView, scaler float64) float64
	_PDFViewWillClickOnLinkWithURL func(sender IPDFView, url objc.IObject /* cross-framework: NSURL */)
}

// SetPDFViewOpenPDFForRemoteGoToAction sets the handler for the PDFViewOpenPDFForRemoteGoToAction delegate method.
//
// Opens a specified page.
func (d *PDFViewDelegate) SetPDFViewOpenPDFForRemoteGoToAction(f func(sender IPDFView, action IPDFActionRemoteGoTo)) {
	d._PDFViewOpenPDFForRemoteGoToAction = f
}

// SetPDFViewParentViewController sets the handler for the PDFViewParentViewController delegate method.
func (d *PDFViewDelegate) SetPDFViewParentViewController(f func() appkit.ViewController) {
	d._PDFViewParentViewController = f
}

// SetPDFViewPerformFind sets the handler for the PDFViewPerformFind delegate method.
//
// Performs a find operation.
func (d *PDFViewDelegate) SetPDFViewPerformFind(f func(sender IPDFView)) {
	d._PDFViewPerformFind = f
}

// SetPDFViewPerformGoToPage sets the handler for the PDFViewPerformGoToPage delegate method.
//
// Performs a go-to operation.
func (d *PDFViewDelegate) SetPDFViewPerformGoToPage(f func(sender IPDFView)) {
	d._PDFViewPerformGoToPage = f
}

// SetPDFViewPerformPrint sets the handler for the PDFViewPerformPrint delegate method.
//
// Prints the current document.
func (d *PDFViewDelegate) SetPDFViewPerformPrint(f func(sender IPDFView)) {
	d._PDFViewPerformPrint = f
}

// SetPDFViewPrintJobTitle sets the handler for the PDFViewPrintJobTitle delegate method.
//
// Overrides the job title used when the   is printed.
func (d *PDFViewDelegate) SetPDFViewPrintJobTitle(f func(sender IPDFView) foundation.String) {
	d._PDFViewPrintJobTitle = f
}

// SetPDFViewWillChangeScaleFactorToScale sets the handler for the PDFViewWillChangeScaleFactorToScale delegate method.
//
// Overrides changes to the scale factor.
func (d *PDFViewDelegate) SetPDFViewWillChangeScaleFactorToScale(f func(sender IPDFView, scaler float64) float64) {
	d._PDFViewWillChangeScaleFactorToScale = f
}

// SetPDFViewWillClickOnLinkWithURL sets the handler for the PDFViewWillClickOnLinkWithURL delegate method.
//
// Handle clicks on URL links in a view.
func (d *PDFViewDelegate) SetPDFViewWillClickOnLinkWithURL(f func(sender IPDFView, url objc.IObject /* cross-framework: NSURL */)) {
	d._PDFViewWillClickOnLinkWithURL = f
}

// PDFViewOpenPDFForRemoteGoToAction implements the PPDFViewDelegate interface.
func (d *PDFViewDelegate) PDFViewOpenPDFForRemoteGoToAction(sender IPDFView, action IPDFActionRemoteGoTo) {
	if d._PDFViewOpenPDFForRemoteGoToAction != nil {
		d._PDFViewOpenPDFForRemoteGoToAction(sender, action)
	}
}

// HasPDFViewOpenPDFForRemoteGoToAction returns true if a handler for PDFViewOpenPDFForRemoteGoToAction has been set.
func (d *PDFViewDelegate) HasPDFViewOpenPDFForRemoteGoToAction() bool {
	return d._PDFViewOpenPDFForRemoteGoToAction != nil
}

// PDFViewParentViewController implements the PPDFViewDelegate interface.
func (d *PDFViewDelegate) PDFViewParentViewController() appkit.ViewController {
	if d._PDFViewParentViewController != nil {
		return d._PDFViewParentViewController()
	}
	var zero appkit.ViewController
	return zero
}

// HasPDFViewParentViewController returns true if a handler for PDFViewParentViewController has been set.
func (d *PDFViewDelegate) HasPDFViewParentViewController() bool {
	return d._PDFViewParentViewController != nil
}

// PDFViewPerformFind implements the PPDFViewDelegate interface.
func (d *PDFViewDelegate) PDFViewPerformFind(sender IPDFView) {
	if d._PDFViewPerformFind != nil {
		d._PDFViewPerformFind(sender)
	}
}

// HasPDFViewPerformFind returns true if a handler for PDFViewPerformFind has been set.
func (d *PDFViewDelegate) HasPDFViewPerformFind() bool {
	return d._PDFViewPerformFind != nil
}

// PDFViewPerformGoToPage implements the PPDFViewDelegate interface.
func (d *PDFViewDelegate) PDFViewPerformGoToPage(sender IPDFView) {
	if d._PDFViewPerformGoToPage != nil {
		d._PDFViewPerformGoToPage(sender)
	}
}

// HasPDFViewPerformGoToPage returns true if a handler for PDFViewPerformGoToPage has been set.
func (d *PDFViewDelegate) HasPDFViewPerformGoToPage() bool {
	return d._PDFViewPerformGoToPage != nil
}

// PDFViewPerformPrint implements the PPDFViewDelegate interface.
func (d *PDFViewDelegate) PDFViewPerformPrint(sender IPDFView) {
	if d._PDFViewPerformPrint != nil {
		d._PDFViewPerformPrint(sender)
	}
}

// HasPDFViewPerformPrint returns true if a handler for PDFViewPerformPrint has been set.
func (d *PDFViewDelegate) HasPDFViewPerformPrint() bool {
	return d._PDFViewPerformPrint != nil
}

// PDFViewPrintJobTitle implements the PPDFViewDelegate interface.
func (d *PDFViewDelegate) PDFViewPrintJobTitle(sender IPDFView) foundation.String {
	if d._PDFViewPrintJobTitle != nil {
		return d._PDFViewPrintJobTitle(sender)
	}
	var zero foundation.String
	return zero
}

// HasPDFViewPrintJobTitle returns true if a handler for PDFViewPrintJobTitle has been set.
func (d *PDFViewDelegate) HasPDFViewPrintJobTitle() bool {
	return d._PDFViewPrintJobTitle != nil
}

// PDFViewWillChangeScaleFactorToScale implements the PPDFViewDelegate interface.
func (d *PDFViewDelegate) PDFViewWillChangeScaleFactorToScale(sender IPDFView, scaler float64) float64 {
	if d._PDFViewWillChangeScaleFactorToScale != nil {
		return d._PDFViewWillChangeScaleFactorToScale(sender, scaler)
	}
	var zero float64
	return zero
}

// HasPDFViewWillChangeScaleFactorToScale returns true if a handler for PDFViewWillChangeScaleFactorToScale has been set.
func (d *PDFViewDelegate) HasPDFViewWillChangeScaleFactorToScale() bool {
	return d._PDFViewWillChangeScaleFactorToScale != nil
}

// PDFViewWillClickOnLinkWithURL implements the PPDFViewDelegate interface.
func (d *PDFViewDelegate) PDFViewWillClickOnLinkWithURL(sender IPDFView, url objc.IObject /* cross-framework: NSURL */) {
	if d._PDFViewWillClickOnLinkWithURL != nil {
		d._PDFViewWillClickOnLinkWithURL(sender, url)
	}
}

// HasPDFViewWillClickOnLinkWithURL returns true if a handler for PDFViewWillClickOnLinkWithURL has been set.
func (d *PDFViewDelegate) HasPDFViewWillClickOnLinkWithURL() bool {
	return d._PDFViewWillClickOnLinkWithURL != nil
}
