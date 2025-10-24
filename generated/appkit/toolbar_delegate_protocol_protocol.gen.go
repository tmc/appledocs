// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PToolbarDelegate is the NSToolbarDelegate protocol interface.
//
// A set of optional methods you use to configure the toolbar and respond to changes.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSToolbarDelegate
type PToolbarDelegate interface {
	// Optional methods
	ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar IToolbar, itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, flag bool) ToolbarItem
	HasToolbarItemForItemIdentifierWillBeInsertedIntoToolbar() bool
	ToolbarItemIdentifierCanBeInsertedAtIndex(toolbar IToolbar, itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, index int) bool
	HasToolbarItemIdentifierCanBeInsertedAtIndex() bool
	ToolbarAllowedItemIdentifiers(toolbar IToolbar) []string
	HasToolbarAllowedItemIdentifiers() bool
	ToolbarDefaultItemIdentifiers(toolbar IToolbar) []string
	HasToolbarDefaultItemIdentifiers() bool
	ToolbarDidRemoveItem(notification foundation.Notification)
	HasToolbarDidRemoveItem() bool
	ToolbarImmovableItemIdentifiers(toolbar IToolbar) unsafe.Pointer
	HasToolbarImmovableItemIdentifiers() bool
	ToolbarSelectableItemIdentifiers(toolbar IToolbar) []string
	HasToolbarSelectableItemIdentifiers() bool
	ToolbarWillAddItem(notification foundation.Notification)
	HasToolbarWillAddItem() bool
}

// ToolbarDelegate is a delegate implementation builder for the PToolbarDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ToolbarDelegate struct {
	_ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar func(toolbar IToolbar, itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, flag bool) ToolbarItem
	_ToolbarItemIdentifierCanBeInsertedAtIndex func(toolbar IToolbar, itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, index int) bool
	_ToolbarAllowedItemIdentifiers func(toolbar IToolbar) []string
	_ToolbarDefaultItemIdentifiers func(toolbar IToolbar) []string
	_ToolbarDidRemoveItem func(notification foundation.Notification)
	_ToolbarImmovableItemIdentifiers func(toolbar IToolbar) unsafe.Pointer
	_ToolbarSelectableItemIdentifiers func(toolbar IToolbar) []string
	_ToolbarWillAddItem func(notification foundation.Notification)
}

// SetToolbarItemForItemIdentifierWillBeInsertedIntoToolbar sets the handler for the ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar delegate method.
//
// Asks the delegate for the toolbar item associated with the specified identifier.
func (d *ToolbarDelegate) SetToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(f func(toolbar IToolbar, itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, flag bool) ToolbarItem) {
	d._ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar = f
}

// SetToolbarItemIdentifierCanBeInsertedAtIndex sets the handler for the ToolbarItemIdentifierCanBeInsertedAtIndex delegate method.
//
// Asks the delegate for a Boolean value that indicates whether the toolbar can place the item at the specified position.
func (d *ToolbarDelegate) SetToolbarItemIdentifierCanBeInsertedAtIndex(f func(toolbar IToolbar, itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, index int) bool) {
	d._ToolbarItemIdentifierCanBeInsertedAtIndex = f
}

// SetToolbarAllowedItemIdentifiers sets the handler for the ToolbarAllowedItemIdentifiers delegate method.
//
// Asks the delegate to provide the items allowed on the toolbar.
func (d *ToolbarDelegate) SetToolbarAllowedItemIdentifiers(f func(toolbar IToolbar) []string) {
	d._ToolbarAllowedItemIdentifiers = f
}

// SetToolbarDefaultItemIdentifiers sets the handler for the ToolbarDefaultItemIdentifiers delegate method.
//
// Asks the delegate to provide the default items to display on the toolbar.
func (d *ToolbarDelegate) SetToolbarDefaultItemIdentifiers(f func(toolbar IToolbar) []string) {
	d._ToolbarDefaultItemIdentifiers = f
}

// SetToolbarDidRemoveItem sets the handler for the ToolbarDidRemoveItem delegate method.
//
// Tells the delegate that the toolbar removed the specified item.
func (d *ToolbarDelegate) SetToolbarDidRemoveItem(f func(notification foundation.Notification)) {
	d._ToolbarDidRemoveItem = f
}

// SetToolbarImmovableItemIdentifiers sets the handler for the ToolbarImmovableItemIdentifiers delegate method.
//
// Asks the delegate to provide the items that people can’t remove from the toolbar or rearrange during the customization process.
func (d *ToolbarDelegate) SetToolbarImmovableItemIdentifiers(f func(toolbar IToolbar) unsafe.Pointer) {
	d._ToolbarImmovableItemIdentifiers = f
}

// SetToolbarSelectableItemIdentifiers sets the handler for the ToolbarSelectableItemIdentifiers delegate method.
//
// Asks the delegate to provide the set of selectable items in the toolbar.
func (d *ToolbarDelegate) SetToolbarSelectableItemIdentifiers(f func(toolbar IToolbar) []string) {
	d._ToolbarSelectableItemIdentifiers = f
}

// SetToolbarWillAddItem sets the handler for the ToolbarWillAddItem delegate method.
//
// Tells the delegate that the toolbar is about to add the specified item.
func (d *ToolbarDelegate) SetToolbarWillAddItem(f func(notification foundation.Notification)) {
	d._ToolbarWillAddItem = f
}

// ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar implements the PToolbarDelegate interface.
func (d *ToolbarDelegate) ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar IToolbar, itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, flag bool) ToolbarItem {
	if d._ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar != nil {
		return d._ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar, itemIdentifier, flag)
	}
	var zero ToolbarItem
	return zero
}

// HasToolbarItemForItemIdentifierWillBeInsertedIntoToolbar returns true if a handler for ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar has been set.
func (d *ToolbarDelegate) HasToolbarItemForItemIdentifierWillBeInsertedIntoToolbar() bool {
	return d._ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar != nil
}

// ToolbarItemIdentifierCanBeInsertedAtIndex implements the PToolbarDelegate interface.
func (d *ToolbarDelegate) ToolbarItemIdentifierCanBeInsertedAtIndex(toolbar IToolbar, itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, index int) bool {
	if d._ToolbarItemIdentifierCanBeInsertedAtIndex != nil {
		return d._ToolbarItemIdentifierCanBeInsertedAtIndex(toolbar, itemIdentifier, index)
	}
	var zero bool
	return zero
}

// HasToolbarItemIdentifierCanBeInsertedAtIndex returns true if a handler for ToolbarItemIdentifierCanBeInsertedAtIndex has been set.
func (d *ToolbarDelegate) HasToolbarItemIdentifierCanBeInsertedAtIndex() bool {
	return d._ToolbarItemIdentifierCanBeInsertedAtIndex != nil
}

// ToolbarAllowedItemIdentifiers implements the PToolbarDelegate interface.
func (d *ToolbarDelegate) ToolbarAllowedItemIdentifiers(toolbar IToolbar) []string {
	if d._ToolbarAllowedItemIdentifiers != nil {
		return d._ToolbarAllowedItemIdentifiers(toolbar)
	}
	var zero []string
	return zero
}

// HasToolbarAllowedItemIdentifiers returns true if a handler for ToolbarAllowedItemIdentifiers has been set.
func (d *ToolbarDelegate) HasToolbarAllowedItemIdentifiers() bool {
	return d._ToolbarAllowedItemIdentifiers != nil
}

// ToolbarDefaultItemIdentifiers implements the PToolbarDelegate interface.
func (d *ToolbarDelegate) ToolbarDefaultItemIdentifiers(toolbar IToolbar) []string {
	if d._ToolbarDefaultItemIdentifiers != nil {
		return d._ToolbarDefaultItemIdentifiers(toolbar)
	}
	var zero []string
	return zero
}

// HasToolbarDefaultItemIdentifiers returns true if a handler for ToolbarDefaultItemIdentifiers has been set.
func (d *ToolbarDelegate) HasToolbarDefaultItemIdentifiers() bool {
	return d._ToolbarDefaultItemIdentifiers != nil
}

// ToolbarDidRemoveItem implements the PToolbarDelegate interface.
func (d *ToolbarDelegate) ToolbarDidRemoveItem(notification foundation.Notification) {
	if d._ToolbarDidRemoveItem != nil {
		d._ToolbarDidRemoveItem(notification)
	}
}

// HasToolbarDidRemoveItem returns true if a handler for ToolbarDidRemoveItem has been set.
func (d *ToolbarDelegate) HasToolbarDidRemoveItem() bool {
	return d._ToolbarDidRemoveItem != nil
}

// ToolbarImmovableItemIdentifiers implements the PToolbarDelegate interface.
func (d *ToolbarDelegate) ToolbarImmovableItemIdentifiers(toolbar IToolbar) unsafe.Pointer {
	if d._ToolbarImmovableItemIdentifiers != nil {
		return d._ToolbarImmovableItemIdentifiers(toolbar)
	}
	var zero unsafe.Pointer
	return zero
}

// HasToolbarImmovableItemIdentifiers returns true if a handler for ToolbarImmovableItemIdentifiers has been set.
func (d *ToolbarDelegate) HasToolbarImmovableItemIdentifiers() bool {
	return d._ToolbarImmovableItemIdentifiers != nil
}

// ToolbarSelectableItemIdentifiers implements the PToolbarDelegate interface.
func (d *ToolbarDelegate) ToolbarSelectableItemIdentifiers(toolbar IToolbar) []string {
	if d._ToolbarSelectableItemIdentifiers != nil {
		return d._ToolbarSelectableItemIdentifiers(toolbar)
	}
	var zero []string
	return zero
}

// HasToolbarSelectableItemIdentifiers returns true if a handler for ToolbarSelectableItemIdentifiers has been set.
func (d *ToolbarDelegate) HasToolbarSelectableItemIdentifiers() bool {
	return d._ToolbarSelectableItemIdentifiers != nil
}

// ToolbarWillAddItem implements the PToolbarDelegate interface.
func (d *ToolbarDelegate) ToolbarWillAddItem(notification foundation.Notification) {
	if d._ToolbarWillAddItem != nil {
		d._ToolbarWillAddItem(notification)
	}
}

// HasToolbarWillAddItem returns true if a handler for ToolbarWillAddItem has been set.
func (d *ToolbarDelegate) HasToolbarWillAddItem() bool {
	return d._ToolbarWillAddItem != nil
}
