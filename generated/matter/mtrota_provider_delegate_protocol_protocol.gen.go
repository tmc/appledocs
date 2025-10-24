// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTROTAProviderDelegate is the MTROTAProviderDelegate protocol interface.
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
// See: doc://com.apple.matter/documentation/Matter/MTROTAProviderDelegate
type PMTROTAProviderDelegate interface {
	// Optional methods
	HandleApplyUpdateRequestForNodeIDControllerParamsCompletion(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams, completion unsafe.Pointer)
	HasHandleApplyUpdateRequestForNodeIDControllerParamsCompletion() bool
	HandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams, completionHandler unsafe.Pointer)
	HasHandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler() bool
	HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, blockSize objc.IObject /* cross-framework: NSNumber */, blockIndex objc.IObject /* cross-framework: NSNumber */, bytesToSkip objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	HasHandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion() bool
	HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, blockSize objc.IObject /* cross-framework: NSNumber */, blockIndex objc.IObject /* cross-framework: NSNumber */, bytesToSkip objc.IObject /* cross-framework: NSNumber */, completionHandler unsafe.Pointer)
	HasHandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler() bool
	HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, fileDesignator objc.IObject /* cross-framework: NSString */, offset objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	HasHandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion() bool
	HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, fileDesignator objc.IObject /* cross-framework: NSString */, offset objc.IObject /* cross-framework: NSNumber */, completionHandler unsafe.Pointer)
	HasHandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler() bool
	HandleBDXTransferSessionEndForNodeIDControllerError(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */)
	HasHandleBDXTransferSessionEndForNodeIDControllerError() bool
	HandleBDXTransferSessionEndForNodeIDControllerMetricsError(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, metrics objc.IObject /* cross-framework: MTRMetrics */, error_ objc.IObject /* cross-framework: Error */)
	HasHandleBDXTransferSessionEndForNodeIDControllerMetricsError() bool
	HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams, completion unsafe.Pointer)
	HasHandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion() bool
	HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams, completionHandler unsafe.Pointer)
	HasHandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler() bool
	HandleQueryImageForNodeIDControllerParamsCompletion(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROTASoftwareUpdateProviderClusterQueryImageParams, completion unsafe.Pointer)
	HasHandleQueryImageForNodeIDControllerParamsCompletion() bool
	HandleQueryImageForNodeIDControllerParamsCompletionHandler(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROtaSoftwareUpdateProviderClusterQueryImageParams, completionHandler unsafe.Pointer)
	HasHandleQueryImageForNodeIDControllerParamsCompletionHandler() bool
}

// MTROTAProviderDelegate is a delegate implementation builder for the PMTROTAProviderDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MTROTAProviderDelegate struct {
	_HandleApplyUpdateRequestForNodeIDControllerParamsCompletion func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams, completion unsafe.Pointer)
	_HandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams, completionHandler unsafe.Pointer)
	_HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, blockSize objc.IObject /* cross-framework: NSNumber */, blockIndex objc.IObject /* cross-framework: NSNumber */, bytesToSkip objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	_HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, blockSize objc.IObject /* cross-framework: NSNumber */, blockIndex objc.IObject /* cross-framework: NSNumber */, bytesToSkip objc.IObject /* cross-framework: NSNumber */, completionHandler unsafe.Pointer)
	_HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, fileDesignator objc.IObject /* cross-framework: NSString */, offset objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	_HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, fileDesignator objc.IObject /* cross-framework: NSString */, offset objc.IObject /* cross-framework: NSNumber */, completionHandler unsafe.Pointer)
	_HandleBDXTransferSessionEndForNodeIDControllerError func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */)
	_HandleBDXTransferSessionEndForNodeIDControllerMetricsError func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, metrics objc.IObject /* cross-framework: MTRMetrics */, error_ objc.IObject /* cross-framework: Error */)
	_HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams, completion unsafe.Pointer)
	_HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams, completionHandler unsafe.Pointer)
	_HandleQueryImageForNodeIDControllerParamsCompletion func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROTASoftwareUpdateProviderClusterQueryImageParams, completion unsafe.Pointer)
	_HandleQueryImageForNodeIDControllerParamsCompletionHandler func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROtaSoftwareUpdateProviderClusterQueryImageParams, completionHandler unsafe.Pointer)
}

// SetHandleApplyUpdateRequestForNodeIDControllerParamsCompletion sets the handler for the HandleApplyUpdateRequestForNodeIDControllerParamsCompletion delegate method.
func (d *MTROTAProviderDelegate) SetHandleApplyUpdateRequestForNodeIDControllerParamsCompletion(f func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams, completion unsafe.Pointer)) {
	d._HandleApplyUpdateRequestForNodeIDControllerParamsCompletion = f
}

// SetHandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler sets the handler for the HandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler delegate method.
func (d *MTROTAProviderDelegate) SetHandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler(f func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams, completionHandler unsafe.Pointer)) {
	d._HandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler = f
}

// SetHandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion sets the handler for the HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion delegate method.
func (d *MTROTAProviderDelegate) SetHandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion(f func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, blockSize objc.IObject /* cross-framework: NSNumber */, blockIndex objc.IObject /* cross-framework: NSNumber */, bytesToSkip objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)) {
	d._HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion = f
}

// SetHandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler sets the handler for the HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler delegate method.
func (d *MTROTAProviderDelegate) SetHandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler(f func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, blockSize objc.IObject /* cross-framework: NSNumber */, blockIndex objc.IObject /* cross-framework: NSNumber */, bytesToSkip objc.IObject /* cross-framework: NSNumber */, completionHandler unsafe.Pointer)) {
	d._HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler = f
}

// SetHandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion sets the handler for the HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion delegate method.
func (d *MTROTAProviderDelegate) SetHandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion(f func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, fileDesignator objc.IObject /* cross-framework: NSString */, offset objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)) {
	d._HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion = f
}

// SetHandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler sets the handler for the HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler delegate method.
func (d *MTROTAProviderDelegate) SetHandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler(f func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, fileDesignator objc.IObject /* cross-framework: NSString */, offset objc.IObject /* cross-framework: NSNumber */, completionHandler unsafe.Pointer)) {
	d._HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler = f
}

// SetHandleBDXTransferSessionEndForNodeIDControllerError sets the handler for the HandleBDXTransferSessionEndForNodeIDControllerError delegate method.
func (d *MTROTAProviderDelegate) SetHandleBDXTransferSessionEndForNodeIDControllerError(f func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */)) {
	d._HandleBDXTransferSessionEndForNodeIDControllerError = f
}

// SetHandleBDXTransferSessionEndForNodeIDControllerMetricsError sets the handler for the HandleBDXTransferSessionEndForNodeIDControllerMetricsError delegate method.
//
// Notify the delegate when a BDX Session ends for some node.  The controller   identifies the fabric the node is on, and the nodeID identifies the node   within that fabric.
func (d *MTROTAProviderDelegate) SetHandleBDXTransferSessionEndForNodeIDControllerMetricsError(f func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, metrics objc.IObject /* cross-framework: MTRMetrics */, error_ objc.IObject /* cross-framework: Error */)) {
	d._HandleBDXTransferSessionEndForNodeIDControllerMetricsError = f
}

// SetHandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion sets the handler for the HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion delegate method.
func (d *MTROTAProviderDelegate) SetHandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion(f func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams, completion unsafe.Pointer)) {
	d._HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion = f
}

// SetHandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler sets the handler for the HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler delegate method.
func (d *MTROTAProviderDelegate) SetHandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler(f func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams, completionHandler unsafe.Pointer)) {
	d._HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler = f
}

// SetHandleQueryImageForNodeIDControllerParamsCompletion sets the handler for the HandleQueryImageForNodeIDControllerParamsCompletion delegate method.
func (d *MTROTAProviderDelegate) SetHandleQueryImageForNodeIDControllerParamsCompletion(f func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROTASoftwareUpdateProviderClusterQueryImageParams, completion unsafe.Pointer)) {
	d._HandleQueryImageForNodeIDControllerParamsCompletion = f
}

// SetHandleQueryImageForNodeIDControllerParamsCompletionHandler sets the handler for the HandleQueryImageForNodeIDControllerParamsCompletionHandler delegate method.
func (d *MTROTAProviderDelegate) SetHandleQueryImageForNodeIDControllerParamsCompletionHandler(f func(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROtaSoftwareUpdateProviderClusterQueryImageParams, completionHandler unsafe.Pointer)) {
	d._HandleQueryImageForNodeIDControllerParamsCompletionHandler = f
}

// HandleApplyUpdateRequestForNodeIDControllerParamsCompletion implements the PMTROTAProviderDelegate interface.
func (d *MTROTAProviderDelegate) HandleApplyUpdateRequestForNodeIDControllerParamsCompletion(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams, completion unsafe.Pointer) {
	if d._HandleApplyUpdateRequestForNodeIDControllerParamsCompletion != nil {
		d._HandleApplyUpdateRequestForNodeIDControllerParamsCompletion(nodeID, controller, params, completion)
	}
}

// HasHandleApplyUpdateRequestForNodeIDControllerParamsCompletion returns true if a handler for HandleApplyUpdateRequestForNodeIDControllerParamsCompletion has been set.
func (d *MTROTAProviderDelegate) HasHandleApplyUpdateRequestForNodeIDControllerParamsCompletion() bool {
	return d._HandleApplyUpdateRequestForNodeIDControllerParamsCompletion != nil
}

// HandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler implements the PMTROTAProviderDelegate interface.
func (d *MTROTAProviderDelegate) HandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams, completionHandler unsafe.Pointer) {
	if d._HandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler != nil {
		d._HandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler(nodeID, controller, params, completionHandler)
	}
}

// HasHandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler returns true if a handler for HandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler has been set.
func (d *MTROTAProviderDelegate) HasHandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler() bool {
	return d._HandleApplyUpdateRequestForNodeIDControllerParamsCompletionHandler != nil
}

// HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion implements the PMTROTAProviderDelegate interface.
func (d *MTROTAProviderDelegate) HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, blockSize objc.IObject /* cross-framework: NSNumber */, blockIndex objc.IObject /* cross-framework: NSNumber */, bytesToSkip objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	if d._HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion != nil {
		d._HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion(nodeID, controller, blockSize, blockIndex, bytesToSkip, completion)
	}
}

// HasHandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion returns true if a handler for HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion has been set.
func (d *MTROTAProviderDelegate) HasHandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion() bool {
	return d._HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletion != nil
}

// HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler implements the PMTROTAProviderDelegate interface.
func (d *MTROTAProviderDelegate) HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, blockSize objc.IObject /* cross-framework: NSNumber */, blockIndex objc.IObject /* cross-framework: NSNumber */, bytesToSkip objc.IObject /* cross-framework: NSNumber */, completionHandler unsafe.Pointer) {
	if d._HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler != nil {
		d._HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler(nodeID, controller, blockSize, blockIndex, bytesToSkip, completionHandler)
	}
}

// HasHandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler returns true if a handler for HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler has been set.
func (d *MTROTAProviderDelegate) HasHandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler() bool {
	return d._HandleBDXQueryForNodeIDControllerBlockSizeBlockIndexBytesToSkipCompletionHandler != nil
}

// HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion implements the PMTROTAProviderDelegate interface.
func (d *MTROTAProviderDelegate) HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, fileDesignator objc.IObject /* cross-framework: NSString */, offset objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	if d._HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion != nil {
		d._HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion(nodeID, controller, fileDesignator, offset, completion)
	}
}

// HasHandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion returns true if a handler for HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion has been set.
func (d *MTROTAProviderDelegate) HasHandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion() bool {
	return d._HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletion != nil
}

// HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler implements the PMTROTAProviderDelegate interface.
func (d *MTROTAProviderDelegate) HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, fileDesignator objc.IObject /* cross-framework: NSString */, offset objc.IObject /* cross-framework: NSNumber */, completionHandler unsafe.Pointer) {
	if d._HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler != nil {
		d._HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler(nodeID, controller, fileDesignator, offset, completionHandler)
	}
}

// HasHandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler returns true if a handler for HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler has been set.
func (d *MTROTAProviderDelegate) HasHandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler() bool {
	return d._HandleBDXTransferSessionBeginForNodeIDControllerFileDesignatorOffsetCompletionHandler != nil
}

// HandleBDXTransferSessionEndForNodeIDControllerError implements the PMTROTAProviderDelegate interface.
func (d *MTROTAProviderDelegate) HandleBDXTransferSessionEndForNodeIDControllerError(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */) {
	if d._HandleBDXTransferSessionEndForNodeIDControllerError != nil {
		d._HandleBDXTransferSessionEndForNodeIDControllerError(nodeID, controller, error_)
	}
}

// HasHandleBDXTransferSessionEndForNodeIDControllerError returns true if a handler for HandleBDXTransferSessionEndForNodeIDControllerError has been set.
func (d *MTROTAProviderDelegate) HasHandleBDXTransferSessionEndForNodeIDControllerError() bool {
	return d._HandleBDXTransferSessionEndForNodeIDControllerError != nil
}

// HandleBDXTransferSessionEndForNodeIDControllerMetricsError implements the PMTROTAProviderDelegate interface.
func (d *MTROTAProviderDelegate) HandleBDXTransferSessionEndForNodeIDControllerMetricsError(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, metrics objc.IObject /* cross-framework: MTRMetrics */, error_ objc.IObject /* cross-framework: Error */) {
	if d._HandleBDXTransferSessionEndForNodeIDControllerMetricsError != nil {
		d._HandleBDXTransferSessionEndForNodeIDControllerMetricsError(nodeID, controller, metrics, error_)
	}
}

// HasHandleBDXTransferSessionEndForNodeIDControllerMetricsError returns true if a handler for HandleBDXTransferSessionEndForNodeIDControllerMetricsError has been set.
func (d *MTROTAProviderDelegate) HasHandleBDXTransferSessionEndForNodeIDControllerMetricsError() bool {
	return d._HandleBDXTransferSessionEndForNodeIDControllerMetricsError != nil
}

// HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion implements the PMTROTAProviderDelegate interface.
func (d *MTROTAProviderDelegate) HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams, completion unsafe.Pointer) {
	if d._HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion != nil {
		d._HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion(nodeID, controller, params, completion)
	}
}

// HasHandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion returns true if a handler for HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion has been set.
func (d *MTROTAProviderDelegate) HasHandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion() bool {
	return d._HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletion != nil
}

// HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler implements the PMTROTAProviderDelegate interface.
func (d *MTROTAProviderDelegate) HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams, completionHandler unsafe.Pointer) {
	if d._HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler != nil {
		d._HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler(nodeID, controller, params, completionHandler)
	}
}

// HasHandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler returns true if a handler for HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler has been set.
func (d *MTROTAProviderDelegate) HasHandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler() bool {
	return d._HandleNotifyUpdateAppliedForNodeIDControllerParamsCompletionHandler != nil
}

// HandleQueryImageForNodeIDControllerParamsCompletion implements the PMTROTAProviderDelegate interface.
func (d *MTROTAProviderDelegate) HandleQueryImageForNodeIDControllerParamsCompletion(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROTASoftwareUpdateProviderClusterQueryImageParams, completion unsafe.Pointer) {
	if d._HandleQueryImageForNodeIDControllerParamsCompletion != nil {
		d._HandleQueryImageForNodeIDControllerParamsCompletion(nodeID, controller, params, completion)
	}
}

// HasHandleQueryImageForNodeIDControllerParamsCompletion returns true if a handler for HandleQueryImageForNodeIDControllerParamsCompletion has been set.
func (d *MTROTAProviderDelegate) HasHandleQueryImageForNodeIDControllerParamsCompletion() bool {
	return d._HandleQueryImageForNodeIDControllerParamsCompletion != nil
}

// HandleQueryImageForNodeIDControllerParamsCompletionHandler implements the PMTROTAProviderDelegate interface.
func (d *MTROTAProviderDelegate) HandleQueryImageForNodeIDControllerParamsCompletionHandler(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController, params IMTROtaSoftwareUpdateProviderClusterQueryImageParams, completionHandler unsafe.Pointer) {
	if d._HandleQueryImageForNodeIDControllerParamsCompletionHandler != nil {
		d._HandleQueryImageForNodeIDControllerParamsCompletionHandler(nodeID, controller, params, completionHandler)
	}
}

// HasHandleQueryImageForNodeIDControllerParamsCompletionHandler returns true if a handler for HandleQueryImageForNodeIDControllerParamsCompletionHandler has been set.
func (d *MTROTAProviderDelegate) HasHandleQueryImageForNodeIDControllerParamsCompletionHandler() bool {
	return d._HandleQueryImageForNodeIDControllerParamsCompletionHandler != nil
}
