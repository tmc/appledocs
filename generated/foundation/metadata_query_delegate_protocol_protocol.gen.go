// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PMetadataQueryDelegate is the NSMetadataQueryDelegate protocol interface.
//
// An interface that enables the delegate of a metadata query to provide substitute results or attributes.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSMetadataQueryDelegate
type PMetadataQueryDelegate interface {
	// Optional methods
	MetadataQueryReplacementValueForAttributeValue(query IMetadataQuery, attrName IString, attrValue objc.IObject) objc.ID
	HasMetadataQueryReplacementValueForAttributeValue() bool
}

// MetadataQueryDelegate is a delegate implementation builder for the PMetadataQueryDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MetadataQueryDelegate struct {
	_MetadataQueryReplacementValueForAttributeValue func(query IMetadataQuery, attrName IString, attrValue objc.IObject) objc.ID
}

// SetMetadataQueryReplacementValueForAttributeValue sets the handler for the MetadataQueryReplacementValueForAttributeValue delegate method.
//
// Returns a different value for a given attribute and value.
func (d *MetadataQueryDelegate) SetMetadataQueryReplacementValueForAttributeValue(f func(query IMetadataQuery, attrName IString, attrValue objc.IObject) objc.ID) {
	d._MetadataQueryReplacementValueForAttributeValue = f
}

// MetadataQueryReplacementValueForAttributeValue implements the PMetadataQueryDelegate interface.
func (d *MetadataQueryDelegate) MetadataQueryReplacementValueForAttributeValue(query IMetadataQuery, attrName IString, attrValue objc.IObject) objc.ID {
	if d._MetadataQueryReplacementValueForAttributeValue != nil {
		return d._MetadataQueryReplacementValueForAttributeValue(query, attrName, attrValue)
	}
	var zero objc.ID
	return zero
}

// HasMetadataQueryReplacementValueForAttributeValue returns true if a handler for MetadataQueryReplacementValueForAttributeValue has been set.
func (d *MetadataQueryDelegate) HasMetadataQueryReplacementValueForAttributeValue() bool {
	return d._MetadataQueryReplacementValueForAttributeValue != nil
}
