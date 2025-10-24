// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PTelephonyNetworkInfoDelegate is the CTTelephonyNetworkInfoDelegate protocol interface.
//
// The methods that the system invokes when data service changes occur.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//
// See: doc://com.apple.coretelephony/documentation/CoreTelephony/CTTelephonyNetworkInfoDelegate
type PTelephonyNetworkInfoDelegate interface {
	// Optional methods
	DataServiceIdentifierDidChange(identifier objc.IObject /* cross-framework: NSString */)
	HasDataServiceIdentifierDidChange() bool
}

// TelephonyNetworkInfoDelegate is a delegate implementation builder for the PTelephonyNetworkInfoDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TelephonyNetworkInfoDelegate struct {
	_DataServiceIdentifierDidChange func(identifier objc.IObject /* cross-framework: NSString */)
}

// SetDataServiceIdentifierDidChange sets the handler for the DataServiceIdentifierDidChange delegate method.
//
// Informs the delegate when the identifier changes for the service that’s currently providing data.
func (d *TelephonyNetworkInfoDelegate) SetDataServiceIdentifierDidChange(f func(identifier objc.IObject /* cross-framework: NSString */)) {
	d._DataServiceIdentifierDidChange = f
}

// DataServiceIdentifierDidChange implements the PTelephonyNetworkInfoDelegate interface.
func (d *TelephonyNetworkInfoDelegate) DataServiceIdentifierDidChange(identifier objc.IObject /* cross-framework: NSString */) {
	if d._DataServiceIdentifierDidChange != nil {
		d._DataServiceIdentifierDidChange(identifier)
	}
}

// HasDataServiceIdentifierDidChange returns true if a handler for DataServiceIdentifierDidChange has been set.
func (d *TelephonyNetworkInfoDelegate) HasDataServiceIdentifierDidChange() bool {
	return d._DataServiceIdentifierDidChange != nil
}
