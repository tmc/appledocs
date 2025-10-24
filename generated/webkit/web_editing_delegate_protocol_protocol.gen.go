// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// PWebEditingDelegate is the WebEditingDelegate protocol interface.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/WebEditingDelegate
type PWebEditingDelegate interface {
	// Optional methods
	UndoManagerForWebView(webView IWebView) foundation.UndoManager
	HasUndoManagerForWebView() bool
	WebViewDoCommandBySelector(webView IWebView, selector objc.SEL) bool
	HasWebViewDoCommandBySelector() bool
	WebViewShouldApplyStyleToElementsInDOMRange(webView IWebView, style IDOMCSSStyleDeclaration, range_ IDOMRange) bool
	HasWebViewShouldApplyStyleToElementsInDOMRange() bool
	WebViewShouldBeginEditingInDOMRange(webView IWebView, range_ IDOMRange) bool
	HasWebViewShouldBeginEditingInDOMRange() bool
	WebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting(webView IWebView, currentRange IDOMRange, proposedRange IDOMRange, selectionAffinity SelectionAffinity /* not a class type */, flag bool) bool
	HasWebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting() bool
	WebViewShouldChangeTypingStyleToStyle(webView IWebView, currentStyle IDOMCSSStyleDeclaration, proposedStyle IDOMCSSStyleDeclaration) bool
	HasWebViewShouldChangeTypingStyleToStyle() bool
	WebViewShouldDeleteDOMRange(webView IWebView, range_ IDOMRange) bool
	HasWebViewShouldDeleteDOMRange() bool
	WebViewShouldEndEditingInDOMRange(webView IWebView, range_ IDOMRange) bool
	HasWebViewShouldEndEditingInDOMRange() bool
	WebViewShouldInsertNodeReplacingDOMRangeGivenAction(webView IWebView, node IDOMNode, range_ IDOMRange, action WebViewInsertAction) bool
	HasWebViewShouldInsertNodeReplacingDOMRangeGivenAction() bool
	WebViewShouldInsertTextReplacingDOMRangeGivenAction(webView IWebView, text objc.IObject /* cross-framework: NSString */, range_ IDOMRange, action WebViewInsertAction) bool
	HasWebViewShouldInsertTextReplacingDOMRangeGivenAction() bool
	WebViewDidBeginEditing(notification foundation.Notification)
	HasWebViewDidBeginEditing() bool
	WebViewDidChange(notification foundation.Notification)
	HasWebViewDidChange() bool
	WebViewDidChangeSelection(notification foundation.Notification)
	HasWebViewDidChangeSelection() bool
	WebViewDidChangeTypingStyle(notification foundation.Notification)
	HasWebViewDidChangeTypingStyle() bool
	WebViewDidEndEditing(notification foundation.Notification)
	HasWebViewDidEndEditing() bool
}

// WebEditingDelegate is a delegate implementation builder for the PWebEditingDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type WebEditingDelegate struct {
	_UndoManagerForWebView                                               func(webView IWebView) foundation.UndoManager
	_WebViewDoCommandBySelector                                          func(webView IWebView, selector objc.SEL) bool
	_WebViewShouldApplyStyleToElementsInDOMRange                         func(webView IWebView, style IDOMCSSStyleDeclaration, range_ IDOMRange) bool
	_WebViewShouldBeginEditingInDOMRange                                 func(webView IWebView, range_ IDOMRange) bool
	_WebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting func(webView IWebView, currentRange IDOMRange, proposedRange IDOMRange, selectionAffinity SelectionAffinity /* not a class type */, flag bool) bool
	_WebViewShouldChangeTypingStyleToStyle                               func(webView IWebView, currentStyle IDOMCSSStyleDeclaration, proposedStyle IDOMCSSStyleDeclaration) bool
	_WebViewShouldDeleteDOMRange                                         func(webView IWebView, range_ IDOMRange) bool
	_WebViewShouldEndEditingInDOMRange                                   func(webView IWebView, range_ IDOMRange) bool
	_WebViewShouldInsertNodeReplacingDOMRangeGivenAction                 func(webView IWebView, node IDOMNode, range_ IDOMRange, action WebViewInsertAction) bool
	_WebViewShouldInsertTextReplacingDOMRangeGivenAction                 func(webView IWebView, text objc.IObject /* cross-framework: NSString */, range_ IDOMRange, action WebViewInsertAction) bool
	_WebViewDidBeginEditing                                              func(notification foundation.Notification)
	_WebViewDidChange                                                    func(notification foundation.Notification)
	_WebViewDidChangeSelection                                           func(notification foundation.Notification)
	_WebViewDidChangeTypingStyle                                         func(notification foundation.Notification)
	_WebViewDidEndEditing                                                func(notification foundation.Notification)
}

// SetUndoManagerForWebView sets the handler for the UndoManagerForWebView delegate method.
//
// Returns the undo manager to be used by a web view.
func (d *WebEditingDelegate) SetUndoManagerForWebView(f func(webView IWebView) foundation.UndoManager) {
	d._UndoManagerForWebView = f
}

// SetWebViewDoCommandBySelector sets the handler for the WebViewDoCommandBySelector delegate method.
//
// Returns whether the receiver performs a command instead of the web view.
func (d *WebEditingDelegate) SetWebViewDoCommandBySelector(f func(webView IWebView, selector objc.SEL) bool) {
	d._WebViewDoCommandBySelector = f
}

// SetWebViewShouldApplyStyleToElementsInDOMRange sets the handler for the WebViewShouldApplyStyleToElementsInDOMRange delegate method.
//
// Returns whether the user should be allowed to apply a style to a range of content.
func (d *WebEditingDelegate) SetWebViewShouldApplyStyleToElementsInDOMRange(f func(webView IWebView, style IDOMCSSStyleDeclaration, range_ IDOMRange) bool) {
	d._WebViewShouldApplyStyleToElementsInDOMRange = f
}

// SetWebViewShouldBeginEditingInDOMRange sets the handler for the WebViewShouldBeginEditingInDOMRange delegate method.
//
// Returns whether the user is allowed to edit a range of content in a web view.
func (d *WebEditingDelegate) SetWebViewShouldBeginEditingInDOMRange(f func(webView IWebView, range_ IDOMRange) bool) {
	d._WebViewShouldBeginEditingInDOMRange = f
}

// SetWebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting sets the handler for the WebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting delegate method.
//
// Returns whether the user should be allowed to change the selected range.
func (d *WebEditingDelegate) SetWebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting(f func(webView IWebView, currentRange IDOMRange, proposedRange IDOMRange, selectionAffinity SelectionAffinity /* not a class type */, flag bool) bool) {
	d._WebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting = f
}

// SetWebViewShouldChangeTypingStyleToStyle sets the handler for the WebViewShouldChangeTypingStyleToStyle delegate method.
//
// Returns whether the user should be allowed to change the typing style in a web view.
func (d *WebEditingDelegate) SetWebViewShouldChangeTypingStyleToStyle(f func(webView IWebView, currentStyle IDOMCSSStyleDeclaration, proposedStyle IDOMCSSStyleDeclaration) bool) {
	d._WebViewShouldChangeTypingStyleToStyle = f
}

// SetWebViewShouldDeleteDOMRange sets the handler for the WebViewShouldDeleteDOMRange delegate method.
//
// Returns whether the user should be allowed to delete a range of content.
func (d *WebEditingDelegate) SetWebViewShouldDeleteDOMRange(f func(webView IWebView, range_ IDOMRange) bool) {
	d._WebViewShouldDeleteDOMRange = f
}

// SetWebViewShouldEndEditingInDOMRange sets the handler for the WebViewShouldEndEditingInDOMRange delegate method.
//
// Returns whether the user should be allowed to end editing.
func (d *WebEditingDelegate) SetWebViewShouldEndEditingInDOMRange(f func(webView IWebView, range_ IDOMRange) bool) {
	d._WebViewShouldEndEditingInDOMRange = f
}

// SetWebViewShouldInsertNodeReplacingDOMRangeGivenAction sets the handler for the WebViewShouldInsertNodeReplacingDOMRangeGivenAction delegate method.
//
// Returns whether the user should be allowed to insert a node in place of a range of content.
func (d *WebEditingDelegate) SetWebViewShouldInsertNodeReplacingDOMRangeGivenAction(f func(webView IWebView, node IDOMNode, range_ IDOMRange, action WebViewInsertAction) bool) {
	d._WebViewShouldInsertNodeReplacingDOMRangeGivenAction = f
}

// SetWebViewShouldInsertTextReplacingDOMRangeGivenAction sets the handler for the WebViewShouldInsertTextReplacingDOMRangeGivenAction delegate method.
//
// Returns whether a user should be allowed to insert text in place of a range of content.
func (d *WebEditingDelegate) SetWebViewShouldInsertTextReplacingDOMRangeGivenAction(f func(webView IWebView, text objc.IObject /* cross-framework: NSString */, range_ IDOMRange, action WebViewInsertAction) bool) {
	d._WebViewShouldInsertTextReplacingDOMRangeGivenAction = f
}

// SetWebViewDidBeginEditing sets the handler for the WebViewDidBeginEditing delegate method.
//
// Sent by the default notification center when the user begins editing the web view.
func (d *WebEditingDelegate) SetWebViewDidBeginEditing(f func(notification foundation.Notification)) {
	d._WebViewDidBeginEditing = f
}

// SetWebViewDidChange sets the handler for the WebViewDidChange delegate method.
//
// Sent by the default notification center when the user changes content in the web view.
func (d *WebEditingDelegate) SetWebViewDidChange(f func(notification foundation.Notification)) {
	d._WebViewDidChange = f
}

// SetWebViewDidChangeSelection sets the handler for the WebViewDidChangeSelection delegate method.
//
// Sent by the default notification center when the user changes the selection in the web view.
func (d *WebEditingDelegate) SetWebViewDidChangeSelection(f func(notification foundation.Notification)) {
	d._WebViewDidChangeSelection = f
}

// SetWebViewDidChangeTypingStyle sets the handler for the WebViewDidChangeTypingStyle delegate method.
//
// Sent by the default notification center when the user changes the typing style in the web view.
func (d *WebEditingDelegate) SetWebViewDidChangeTypingStyle(f func(notification foundation.Notification)) {
	d._WebViewDidChangeTypingStyle = f
}

// SetWebViewDidEndEditing sets the handler for the WebViewDidEndEditing delegate method.
//
// Sent by the default notification center when the user stops editing the web view.
func (d *WebEditingDelegate) SetWebViewDidEndEditing(f func(notification foundation.Notification)) {
	d._WebViewDidEndEditing = f
}

// UndoManagerForWebView implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) UndoManagerForWebView(webView IWebView) foundation.UndoManager {
	if d._UndoManagerForWebView != nil {
		return d._UndoManagerForWebView(webView)
	}
	var zero foundation.UndoManager
	return zero
}

// HasUndoManagerForWebView returns true if a handler for UndoManagerForWebView has been set.
func (d *WebEditingDelegate) HasUndoManagerForWebView() bool {
	return d._UndoManagerForWebView != nil
}

// WebViewDoCommandBySelector implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) WebViewDoCommandBySelector(webView IWebView, selector objc.SEL) bool {
	if d._WebViewDoCommandBySelector != nil {
		return d._WebViewDoCommandBySelector(webView, selector)
	}
	var zero bool
	return zero
}

// HasWebViewDoCommandBySelector returns true if a handler for WebViewDoCommandBySelector has been set.
func (d *WebEditingDelegate) HasWebViewDoCommandBySelector() bool {
	return d._WebViewDoCommandBySelector != nil
}

// WebViewShouldApplyStyleToElementsInDOMRange implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) WebViewShouldApplyStyleToElementsInDOMRange(webView IWebView, style IDOMCSSStyleDeclaration, range_ IDOMRange) bool {
	if d._WebViewShouldApplyStyleToElementsInDOMRange != nil {
		return d._WebViewShouldApplyStyleToElementsInDOMRange(webView, style, range_)
	}
	var zero bool
	return zero
}

// HasWebViewShouldApplyStyleToElementsInDOMRange returns true if a handler for WebViewShouldApplyStyleToElementsInDOMRange has been set.
func (d *WebEditingDelegate) HasWebViewShouldApplyStyleToElementsInDOMRange() bool {
	return d._WebViewShouldApplyStyleToElementsInDOMRange != nil
}

// WebViewShouldBeginEditingInDOMRange implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) WebViewShouldBeginEditingInDOMRange(webView IWebView, range_ IDOMRange) bool {
	if d._WebViewShouldBeginEditingInDOMRange != nil {
		return d._WebViewShouldBeginEditingInDOMRange(webView, range_)
	}
	var zero bool
	return zero
}

// HasWebViewShouldBeginEditingInDOMRange returns true if a handler for WebViewShouldBeginEditingInDOMRange has been set.
func (d *WebEditingDelegate) HasWebViewShouldBeginEditingInDOMRange() bool {
	return d._WebViewShouldBeginEditingInDOMRange != nil
}

// WebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) WebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting(webView IWebView, currentRange IDOMRange, proposedRange IDOMRange, selectionAffinity SelectionAffinity /* not a class type */, flag bool) bool {
	if d._WebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting != nil {
		return d._WebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting(webView, currentRange, proposedRange, selectionAffinity, flag)
	}
	var zero bool
	return zero
}

// HasWebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting returns true if a handler for WebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting has been set.
func (d *WebEditingDelegate) HasWebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting() bool {
	return d._WebViewShouldChangeSelectedDOMRangeToDOMRangeAffinityStillSelecting != nil
}

// WebViewShouldChangeTypingStyleToStyle implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) WebViewShouldChangeTypingStyleToStyle(webView IWebView, currentStyle IDOMCSSStyleDeclaration, proposedStyle IDOMCSSStyleDeclaration) bool {
	if d._WebViewShouldChangeTypingStyleToStyle != nil {
		return d._WebViewShouldChangeTypingStyleToStyle(webView, currentStyle, proposedStyle)
	}
	var zero bool
	return zero
}

// HasWebViewShouldChangeTypingStyleToStyle returns true if a handler for WebViewShouldChangeTypingStyleToStyle has been set.
func (d *WebEditingDelegate) HasWebViewShouldChangeTypingStyleToStyle() bool {
	return d._WebViewShouldChangeTypingStyleToStyle != nil
}

// WebViewShouldDeleteDOMRange implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) WebViewShouldDeleteDOMRange(webView IWebView, range_ IDOMRange) bool {
	if d._WebViewShouldDeleteDOMRange != nil {
		return d._WebViewShouldDeleteDOMRange(webView, range_)
	}
	var zero bool
	return zero
}

// HasWebViewShouldDeleteDOMRange returns true if a handler for WebViewShouldDeleteDOMRange has been set.
func (d *WebEditingDelegate) HasWebViewShouldDeleteDOMRange() bool {
	return d._WebViewShouldDeleteDOMRange != nil
}

// WebViewShouldEndEditingInDOMRange implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) WebViewShouldEndEditingInDOMRange(webView IWebView, range_ IDOMRange) bool {
	if d._WebViewShouldEndEditingInDOMRange != nil {
		return d._WebViewShouldEndEditingInDOMRange(webView, range_)
	}
	var zero bool
	return zero
}

// HasWebViewShouldEndEditingInDOMRange returns true if a handler for WebViewShouldEndEditingInDOMRange has been set.
func (d *WebEditingDelegate) HasWebViewShouldEndEditingInDOMRange() bool {
	return d._WebViewShouldEndEditingInDOMRange != nil
}

// WebViewShouldInsertNodeReplacingDOMRangeGivenAction implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) WebViewShouldInsertNodeReplacingDOMRangeGivenAction(webView IWebView, node IDOMNode, range_ IDOMRange, action WebViewInsertAction) bool {
	if d._WebViewShouldInsertNodeReplacingDOMRangeGivenAction != nil {
		return d._WebViewShouldInsertNodeReplacingDOMRangeGivenAction(webView, node, range_, action)
	}
	var zero bool
	return zero
}

// HasWebViewShouldInsertNodeReplacingDOMRangeGivenAction returns true if a handler for WebViewShouldInsertNodeReplacingDOMRangeGivenAction has been set.
func (d *WebEditingDelegate) HasWebViewShouldInsertNodeReplacingDOMRangeGivenAction() bool {
	return d._WebViewShouldInsertNodeReplacingDOMRangeGivenAction != nil
}

// WebViewShouldInsertTextReplacingDOMRangeGivenAction implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) WebViewShouldInsertTextReplacingDOMRangeGivenAction(webView IWebView, text objc.IObject /* cross-framework: NSString */, range_ IDOMRange, action WebViewInsertAction) bool {
	if d._WebViewShouldInsertTextReplacingDOMRangeGivenAction != nil {
		return d._WebViewShouldInsertTextReplacingDOMRangeGivenAction(webView, text, range_, action)
	}
	var zero bool
	return zero
}

// HasWebViewShouldInsertTextReplacingDOMRangeGivenAction returns true if a handler for WebViewShouldInsertTextReplacingDOMRangeGivenAction has been set.
func (d *WebEditingDelegate) HasWebViewShouldInsertTextReplacingDOMRangeGivenAction() bool {
	return d._WebViewShouldInsertTextReplacingDOMRangeGivenAction != nil
}

// WebViewDidBeginEditing implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) WebViewDidBeginEditing(notification foundation.Notification) {
	if d._WebViewDidBeginEditing != nil {
		d._WebViewDidBeginEditing(notification)
	}
}

// HasWebViewDidBeginEditing returns true if a handler for WebViewDidBeginEditing has been set.
func (d *WebEditingDelegate) HasWebViewDidBeginEditing() bool {
	return d._WebViewDidBeginEditing != nil
}

// WebViewDidChange implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) WebViewDidChange(notification foundation.Notification) {
	if d._WebViewDidChange != nil {
		d._WebViewDidChange(notification)
	}
}

// HasWebViewDidChange returns true if a handler for WebViewDidChange has been set.
func (d *WebEditingDelegate) HasWebViewDidChange() bool {
	return d._WebViewDidChange != nil
}

// WebViewDidChangeSelection implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) WebViewDidChangeSelection(notification foundation.Notification) {
	if d._WebViewDidChangeSelection != nil {
		d._WebViewDidChangeSelection(notification)
	}
}

// HasWebViewDidChangeSelection returns true if a handler for WebViewDidChangeSelection has been set.
func (d *WebEditingDelegate) HasWebViewDidChangeSelection() bool {
	return d._WebViewDidChangeSelection != nil
}

// WebViewDidChangeTypingStyle implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) WebViewDidChangeTypingStyle(notification foundation.Notification) {
	if d._WebViewDidChangeTypingStyle != nil {
		d._WebViewDidChangeTypingStyle(notification)
	}
}

// HasWebViewDidChangeTypingStyle returns true if a handler for WebViewDidChangeTypingStyle has been set.
func (d *WebEditingDelegate) HasWebViewDidChangeTypingStyle() bool {
	return d._WebViewDidChangeTypingStyle != nil
}

// WebViewDidEndEditing implements the PWebEditingDelegate interface.
func (d *WebEditingDelegate) WebViewDidEndEditing(notification foundation.Notification) {
	if d._WebViewDidEndEditing != nil {
		d._WebViewDidEndEditing(notification)
	}
}

// HasWebViewDidEndEditing returns true if a handler for WebViewDidEndEditing has been set.
func (d *WebEditingDelegate) HasWebViewDidEndEditing() bool {
	return d._WebViewDidEndEditing != nil
}
