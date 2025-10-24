// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PABImageClient is the ABImageClient protocol interface.
//
// Methods for responding to a request to load images associated with a contact.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.addressbook/documentation/AddressBook/ABImageClient
type PABImageClient interface {
	// Required methods
	ConsumeImageDataForTag(data objc.IObject /* cross-framework: NSData */, tag int)/* debug [protocol_interface/required_method]: ConsumeImageDataForTag */
}
