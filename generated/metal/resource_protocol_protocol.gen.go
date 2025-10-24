// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (

	"github.com/tmc/appledocs/generated/objectivec"
)

// PResource is the MTLResource protocol interface.
//
// An allocation of memory accessible to a GPU.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLResource
type PResource interface {
	// Required methods
	IsAliasable() bool/* debug [protocol_interface/required_method]: IsAliasable */
	MakeAliasable()/* debug [protocol_interface/required_method]: MakeAliasable */
	SetOwnerWithIdentity(task_id_token objectivec.IObject) objectivec.IObject/* debug [protocol_interface/required_method]: SetOwnerWithIdentity */
	SetPurgeableState(state PurgeableState) PurgeableState/* debug [protocol_interface/required_method]: SetPurgeableState */
}
