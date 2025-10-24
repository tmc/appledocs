// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PFSManageableResourceMaintenanceOperations is the FSManageableResourceMaintenanceOperations protocol interface.
//
// Maintenance operations for a file system’s resources.
//
// Availability:
//   - macOS 15.4+
//
// See: doc://FSKit/documentation/FSKit/FSManageableResourceMaintenanceOperations
type PFSManageableResourceMaintenanceOperations interface {
	// Required methods
	StartCheckWithTaskOptionsError(task IFSTask, options IFSTaskOptions, error_ unsafe.Pointer) foundation.Progress/* debug [protocol_interface/required_method]: StartCheckWithTaskOptionsError */
	StartFormatWithTaskOptionsError(task IFSTask, options IFSTaskOptions, error_ unsafe.Pointer) foundation.Progress/* debug [protocol_interface/required_method]: StartFormatWithTaskOptionsError */
}
