// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PSDataStoreDelegate is the CLSDataStoreDelegate protocol interface.
//
// An interface the data store uses to request new contexts.
//
// Availability:
//   - Mac Catalyst 11.3+
//   - iOS 11.3+
//   - iPadOS 11.3+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.classkit/documentation/ClassKit/CLSDataStoreDelegate
type PSDataStoreDelegate interface {
	// Required methods
	CreateContextForIdentifierParentContextParentIdentifierPath(identifier objc.IObject /* cross-framework: NSString */, parentContext ICLSContext, parentIdentifierPath []string) SContext/* debug [protocol_interface/required_method]: CreateContextForIdentifierParentContextParentIdentifierPath */
}

// SDataStoreDelegate is a delegate implementation builder for the PSDataStoreDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SDataStoreDelegate struct {
	_CreateContextForIdentifierParentContextParentIdentifierPath func(identifier objc.IObject /* cross-framework: NSString */, parentContext ICLSContext, parentIdentifierPath []string) SContext
}

// SetCreateContextForIdentifierParentContextParentIdentifierPath sets the handler for the CreateContextForIdentifierParentContextParentIdentifierPath delegate method.
//
// Asks the delegate for a new context with the given identifier for the given parent context.
func (d *SDataStoreDelegate) SetCreateContextForIdentifierParentContextParentIdentifierPath(f func(identifier objc.IObject /* cross-framework: NSString */, parentContext ICLSContext, parentIdentifierPath []string) SContext) {
	d._CreateContextForIdentifierParentContextParentIdentifierPath = f
}

// CreateContextForIdentifierParentContextParentIdentifierPath implements the PSDataStoreDelegate interface.
func (d *SDataStoreDelegate) CreateContextForIdentifierParentContextParentIdentifierPath(identifier objc.IObject /* cross-framework: NSString */, parentContext ICLSContext, parentIdentifierPath []string) SContext {
	if d._CreateContextForIdentifierParentContextParentIdentifierPath != nil {
		return d._CreateContextForIdentifierParentContextParentIdentifierPath(identifier, parentContext, parentIdentifierPath)
	}
	var zero SContext
	return zero
}

// HasCreateContextForIdentifierParentContextParentIdentifierPath returns true if a handler for CreateContextForIdentifierParentContextParentIdentifierPath has been set.
func (d *SDataStoreDelegate) HasCreateContextForIdentifierParentContextParentIdentifierPath() bool {
	return d._CreateContextForIdentifierParentContextParentIdentifierPath != nil
}
