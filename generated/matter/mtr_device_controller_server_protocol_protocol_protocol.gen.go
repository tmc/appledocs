// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTRDeviceControllerServerProtocol is the MTRDeviceControllerServerProtocol protocol interface.
//
// Availability:
//   - Mac Catalyst 16.1+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+
//
// See: doc://com.apple.matter/documentation/Matter/MTRDeviceControllerServerProtocol
type PMTRDeviceControllerServerProtocol interface {
	// Required methods
	GetAnyDeviceControllerWithCompletion(completion unsafe.Pointer)/* debug [protocol_interface/required_method]: GetAnyDeviceControllerWithCompletion */
	InvokeCommandWithControllerNodeIdEndpointIdClusterIdCommandIdFieldsTimedInvokeTimeoutCompletion(controller objc.IObject, nodeId uint64, endpointId objc.IObject /* cross-framework: NSNumber */, clusterId objc.IObject /* cross-framework: NSNumber */, commandId objc.IObject /* cross-framework: NSNumber */, fields objc.IObject, timeoutMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: InvokeCommandWithControllerNodeIdEndpointIdClusterIdCommandIdFieldsTimedInvokeTimeoutCompletion */
	ReadAttributeWithControllerNodeIdEndpointIdClusterIdAttributeIdParamsCompletion(controller objc.IObject, nodeId uint64, endpointId objc.IObject /* cross-framework: NSNumber */, clusterId objc.IObject /* cross-framework: NSNumber */, attributeId objc.IObject /* cross-framework: NSNumber */, params foundation.IDictionary, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: ReadAttributeWithControllerNodeIdEndpointIdClusterIdAttributeIdParamsCompletion */
	ReadAttributeCacheWithControllerNodeIdEndpointIdClusterIdAttributeIdCompletion(controller objc.IObject, nodeId uint64, endpointId objc.IObject /* cross-framework: NSNumber */, clusterId objc.IObject /* cross-framework: NSNumber */, attributeId objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: ReadAttributeCacheWithControllerNodeIdEndpointIdClusterIdAttributeIdCompletion */
	StopReportsWithControllerNodeIdCompletion(controller objc.IObject, nodeId uint64, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: StopReportsWithControllerNodeIdCompletion */
	SubscribeWithControllerNodeIdMinIntervalMaxIntervalParamsShouldCacheCompletion(controller objc.IObject, nodeId uint64, minInterval objc.IObject /* cross-framework: NSNumber */, maxInterval objc.IObject /* cross-framework: NSNumber */, params foundation.IDictionary, shouldCache bool, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: SubscribeWithControllerNodeIdMinIntervalMaxIntervalParamsShouldCacheCompletion */
	SubscribeAttributeWithControllerNodeIdEndpointIdClusterIdAttributeIdMinIntervalMaxIntervalParamsEstablishedHandler(controller objc.IObject, nodeId uint64, endpointId objc.IObject /* cross-framework: NSNumber */, clusterId objc.IObject /* cross-framework: NSNumber */, attributeId objc.IObject /* cross-framework: NSNumber */, minInterval objc.IObject /* cross-framework: NSNumber */, maxInterval objc.IObject /* cross-framework: NSNumber */, params foundation.IDictionary, establishedHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: SubscribeAttributeWithControllerNodeIdEndpointIdClusterIdAttributeIdMinIntervalMaxIntervalParamsEstablishedHandler */
	WriteAttributeWithControllerNodeIdEndpointIdClusterIdAttributeIdValueTimedWriteTimeoutCompletion(controller objc.IObject, nodeId uint64, endpointId objc.IObject /* cross-framework: NSNumber */, clusterId objc.IObject /* cross-framework: NSNumber */, attributeId objc.IObject /* cross-framework: NSNumber */, value objc.IObject, timeoutMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: WriteAttributeWithControllerNodeIdEndpointIdClusterIdAttributeIdValueTimedWriteTimeoutCompletion */
	// Optional methods
	DownloadLogWithControllerNodeIdTypeTimeoutCompletion(controller objc.IObject, nodeId objc.IObject /* cross-framework: NSNumber */, type_ unsafe.Pointer, timeout float64, completion unsafe.Pointer)
	HasDownloadLogWithControllerNodeIdTypeTimeoutCompletion() bool
	GetDeviceControllerWithFabricIdCompletion(fabricId uint64, completion unsafe.Pointer)
	HasGetDeviceControllerWithFabricIdCompletion() bool
}
