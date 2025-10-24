// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PSFSafariViewControllerDelegate is the SFSafariViewControllerDelegate protocol interface.
//
// A protocol used to implement custom event handling for a Safari view controller.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//
// See: doc://com.apple.safariservices/documentation/SafariServices/SFSafariViewControllerDelegate
type PSFSafariViewControllerDelegate interface {
	// Optional methods
	SafariViewControllerActivityItemsForURLTitle(controller ISFSafariViewController, URL objc.IObject /* cross-framework: NSURL */, title objc.IObject /* cross-framework: NSString */) []Activity
	HasSafariViewControllerActivityItemsForURLTitle() bool
	SafariViewControllerDidCompleteInitialLoad(controller ISFSafariViewController, didLoadSuccessfully bool)
	HasSafariViewControllerDidCompleteInitialLoad() bool
	SafariViewControllerExcludedActivityTypesForURLTitle(controller ISFSafariViewController, URL objc.IObject /* cross-framework: NSURL */, title objc.IObject /* cross-framework: NSString */) []string
	HasSafariViewControllerExcludedActivityTypesForURLTitle() bool
	SafariViewControllerInitialLoadDidRedirectToURL(controller ISFSafariViewController, URL objc.IObject /* cross-framework: NSURL */)
	HasSafariViewControllerInitialLoadDidRedirectToURL() bool
	SafariViewControllerDidFinish(controller ISFSafariViewController)
	HasSafariViewControllerDidFinish() bool
	SafariViewControllerWillOpenInBrowser(controller ISFSafariViewController)
	HasSafariViewControllerWillOpenInBrowser() bool
}

// SFSafariViewControllerDelegate is a delegate implementation builder for the PSFSafariViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SFSafariViewControllerDelegate struct {
	_SafariViewControllerActivityItemsForURLTitle func(controller ISFSafariViewController, URL objc.IObject /* cross-framework: NSURL */, title objc.IObject /* cross-framework: NSString */) []Activity
	_SafariViewControllerDidCompleteInitialLoad func(controller ISFSafariViewController, didLoadSuccessfully bool)
	_SafariViewControllerExcludedActivityTypesForURLTitle func(controller ISFSafariViewController, URL objc.IObject /* cross-framework: NSURL */, title objc.IObject /* cross-framework: NSString */) []string
	_SafariViewControllerInitialLoadDidRedirectToURL func(controller ISFSafariViewController, URL objc.IObject /* cross-framework: NSURL */)
	_SafariViewControllerDidFinish func(controller ISFSafariViewController)
	_SafariViewControllerWillOpenInBrowser func(controller ISFSafariViewController)
}

// SetSafariViewControllerActivityItemsForURLTitle sets the handler for the SafariViewControllerActivityItemsForURLTitle delegate method.
//
// Tells the delegate that the user tapped an Action button.
func (d *SFSafariViewControllerDelegate) SetSafariViewControllerActivityItemsForURLTitle(f func(controller ISFSafariViewController, URL objc.IObject /* cross-framework: NSURL */, title objc.IObject /* cross-framework: NSString */) []Activity) {
	d._SafariViewControllerActivityItemsForURLTitle = f
}

// SetSafariViewControllerDidCompleteInitialLoad sets the handler for the SafariViewControllerDidCompleteInitialLoad delegate method.
//
// Tells the delegate that the initial URL load completed.
func (d *SFSafariViewControllerDelegate) SetSafariViewControllerDidCompleteInitialLoad(f func(controller ISFSafariViewController, didLoadSuccessfully bool)) {
	d._SafariViewControllerDidCompleteInitialLoad = f
}

// SetSafariViewControllerExcludedActivityTypesForURLTitle sets the handler for the SafariViewControllerExcludedActivityTypesForURLTitle delegate method.
func (d *SFSafariViewControllerDelegate) SetSafariViewControllerExcludedActivityTypesForURLTitle(f func(controller ISFSafariViewController, URL objc.IObject /* cross-framework: NSURL */, title objc.IObject /* cross-framework: NSString */) []string) {
	d._SafariViewControllerExcludedActivityTypesForURLTitle = f
}

// SetSafariViewControllerInitialLoadDidRedirectToURL sets the handler for the SafariViewControllerInitialLoadDidRedirectToURL delegate method.
func (d *SFSafariViewControllerDelegate) SetSafariViewControllerInitialLoadDidRedirectToURL(f func(controller ISFSafariViewController, URL objc.IObject /* cross-framework: NSURL */)) {
	d._SafariViewControllerInitialLoadDidRedirectToURL = f
}

// SetSafariViewControllerDidFinish sets the handler for the SafariViewControllerDidFinish delegate method.
//
// Tells the delegate that the user dismissed the view.
func (d *SFSafariViewControllerDelegate) SetSafariViewControllerDidFinish(f func(controller ISFSafariViewController)) {
	d._SafariViewControllerDidFinish = f
}

// SetSafariViewControllerWillOpenInBrowser sets the handler for the SafariViewControllerWillOpenInBrowser delegate method.
func (d *SFSafariViewControllerDelegate) SetSafariViewControllerWillOpenInBrowser(f func(controller ISFSafariViewController)) {
	d._SafariViewControllerWillOpenInBrowser = f
}

// SafariViewControllerActivityItemsForURLTitle implements the PSFSafariViewControllerDelegate interface.
func (d *SFSafariViewControllerDelegate) SafariViewControllerActivityItemsForURLTitle(controller ISFSafariViewController, URL objc.IObject /* cross-framework: NSURL */, title objc.IObject /* cross-framework: NSString */) []Activity {
	if d._SafariViewControllerActivityItemsForURLTitle != nil {
		return d._SafariViewControllerActivityItemsForURLTitle(controller, URL, title)
	}
	var zero []Activity
	return zero
}

// HasSafariViewControllerActivityItemsForURLTitle returns true if a handler for SafariViewControllerActivityItemsForURLTitle has been set.
func (d *SFSafariViewControllerDelegate) HasSafariViewControllerActivityItemsForURLTitle() bool {
	return d._SafariViewControllerActivityItemsForURLTitle != nil
}

// SafariViewControllerDidCompleteInitialLoad implements the PSFSafariViewControllerDelegate interface.
func (d *SFSafariViewControllerDelegate) SafariViewControllerDidCompleteInitialLoad(controller ISFSafariViewController, didLoadSuccessfully bool) {
	if d._SafariViewControllerDidCompleteInitialLoad != nil {
		d._SafariViewControllerDidCompleteInitialLoad(controller, didLoadSuccessfully)
	}
}

// HasSafariViewControllerDidCompleteInitialLoad returns true if a handler for SafariViewControllerDidCompleteInitialLoad has been set.
func (d *SFSafariViewControllerDelegate) HasSafariViewControllerDidCompleteInitialLoad() bool {
	return d._SafariViewControllerDidCompleteInitialLoad != nil
}

// SafariViewControllerExcludedActivityTypesForURLTitle implements the PSFSafariViewControllerDelegate interface.
func (d *SFSafariViewControllerDelegate) SafariViewControllerExcludedActivityTypesForURLTitle(controller ISFSafariViewController, URL objc.IObject /* cross-framework: NSURL */, title objc.IObject /* cross-framework: NSString */) []string {
	if d._SafariViewControllerExcludedActivityTypesForURLTitle != nil {
		return d._SafariViewControllerExcludedActivityTypesForURLTitle(controller, URL, title)
	}
	var zero []string
	return zero
}

// HasSafariViewControllerExcludedActivityTypesForURLTitle returns true if a handler for SafariViewControllerExcludedActivityTypesForURLTitle has been set.
func (d *SFSafariViewControllerDelegate) HasSafariViewControllerExcludedActivityTypesForURLTitle() bool {
	return d._SafariViewControllerExcludedActivityTypesForURLTitle != nil
}

// SafariViewControllerInitialLoadDidRedirectToURL implements the PSFSafariViewControllerDelegate interface.
func (d *SFSafariViewControllerDelegate) SafariViewControllerInitialLoadDidRedirectToURL(controller ISFSafariViewController, URL objc.IObject /* cross-framework: NSURL */) {
	if d._SafariViewControllerInitialLoadDidRedirectToURL != nil {
		d._SafariViewControllerInitialLoadDidRedirectToURL(controller, URL)
	}
}

// HasSafariViewControllerInitialLoadDidRedirectToURL returns true if a handler for SafariViewControllerInitialLoadDidRedirectToURL has been set.
func (d *SFSafariViewControllerDelegate) HasSafariViewControllerInitialLoadDidRedirectToURL() bool {
	return d._SafariViewControllerInitialLoadDidRedirectToURL != nil
}

// SafariViewControllerDidFinish implements the PSFSafariViewControllerDelegate interface.
func (d *SFSafariViewControllerDelegate) SafariViewControllerDidFinish(controller ISFSafariViewController) {
	if d._SafariViewControllerDidFinish != nil {
		d._SafariViewControllerDidFinish(controller)
	}
}

// HasSafariViewControllerDidFinish returns true if a handler for SafariViewControllerDidFinish has been set.
func (d *SFSafariViewControllerDelegate) HasSafariViewControllerDidFinish() bool {
	return d._SafariViewControllerDidFinish != nil
}

// SafariViewControllerWillOpenInBrowser implements the PSFSafariViewControllerDelegate interface.
func (d *SFSafariViewControllerDelegate) SafariViewControllerWillOpenInBrowser(controller ISFSafariViewController) {
	if d._SafariViewControllerWillOpenInBrowser != nil {
		d._SafariViewControllerWillOpenInBrowser(controller)
	}
}

// HasSafariViewControllerWillOpenInBrowser returns true if a handler for SafariViewControllerWillOpenInBrowser has been set.
func (d *SFSafariViewControllerDelegate) HasSafariViewControllerWillOpenInBrowser() bool {
	return d._SafariViewControllerWillOpenInBrowser != nil
}
