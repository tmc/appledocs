// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PICCameraDeviceDelegate is the ICCameraDeviceDelegate protocol interface.
//
// Methods for detecting cameras, getting metadata and thumbnails, handling access and capability changes, and performing other actions on connected cameras.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - visionOS +
//
// See: doc://com.apple.imagecapturecore/documentation/ImageCaptureCore/ICCameraDeviceDelegate
type PICCameraDeviceDelegate interface {
	// Required methods
	CameraDeviceDidChangeCapability()/* debug [protocol_interface/required_method]: CameraDeviceDidChangeCapability */
	DeviceDidBecomeReady()/* debug [protocol_interface/required_method]: DeviceDidBecomeReady */
	DeviceDidBecomeReadyWithCompleteContentCatalog(device ICCameraDevice)/* debug [protocol_interface/required_method]: DeviceDidBecomeReadyWithCompleteContentCatalog */
	CameraDeviceDidEnableAccessRestriction()/* debug [protocol_interface/required_method]: CameraDeviceDidEnableAccessRestriction */
	CameraDeviceDidRemoveAccessRestriction()/* debug [protocol_interface/required_method]: CameraDeviceDidRemoveAccessRestriction */
	CameraDeviceDidAddItems(camera ICCameraDevice, items []CCameraItem)/* debug [protocol_interface/required_method]: CameraDeviceDidAddItems */
	CameraDeviceDidReceiveMetadataForItemError(camera ICCameraDevice, metadata objc.IObject /* cross-framework: NSDictionary */, item ICCameraItem, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: CameraDeviceDidReceiveMetadataForItemError */
	CameraDeviceDidReceivePTPEvent(camera ICCameraDevice, eventData objc.IObject /* cross-framework: NSData */)/* debug [protocol_interface/required_method]: CameraDeviceDidReceivePTPEvent */
	CameraDeviceDidReceiveThumbnailForItemError(camera ICCameraDevice, thumbnail ImageRef /* not a class type */, item ICCameraItem, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: CameraDeviceDidReceiveThumbnailForItemError */
	CameraDeviceDidRemoveItems(camera ICCameraDevice, items []CCameraItem)/* debug [protocol_interface/required_method]: CameraDeviceDidRemoveItems */
	CameraDeviceDidRenameItems(camera ICCameraDevice, items []CCameraItem)/* debug [protocol_interface/required_method]: CameraDeviceDidRenameItems */
	// Optional methods
	CameraDevice()
	HasCameraDevice() bool
	CameraDeviceDidAddItem(camera ICCameraDevice, item ICCameraItem)
	HasCameraDeviceDidAddItem() bool
	CameraDeviceDidCompleteDeleteFilesWithError(camera ICCameraDevice, error_ objc.IObject /* cross-framework: Error */)
	HasCameraDeviceDidCompleteDeleteFilesWithError() bool
	CameraDeviceDidReceiveMetadataForItem(camera ICCameraDevice, item ICCameraItem)
	HasCameraDeviceDidReceiveMetadataForItem() bool
	CameraDeviceDidReceiveThumbnailForItem(camera ICCameraDevice, item ICCameraItem)
	HasCameraDeviceDidReceiveThumbnailForItem() bool
	CameraDeviceDidRemoveItem(camera ICCameraDevice, item ICCameraItem)
	HasCameraDeviceDidRemoveItem() bool
	CameraDeviceShouldGetMetadataOfItem(cameraDevice ICCameraDevice, item ICCameraItem) bool
	HasCameraDeviceShouldGetMetadataOfItem() bool
	CameraDeviceShouldGetThumbnailOfItem(cameraDevice ICCameraDevice, item ICCameraItem) bool
	HasCameraDeviceShouldGetThumbnailOfItem() bool
}

// ICCameraDeviceDelegate is a delegate implementation builder for the PICCameraDeviceDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ICCameraDeviceDelegate struct {
	_CameraDevice func()
	_CameraDeviceDidAddItem func(camera ICCameraDevice, item ICCameraItem)
	_CameraDeviceDidCompleteDeleteFilesWithError func(camera ICCameraDevice, error_ objc.IObject /* cross-framework: Error */)
	_CameraDeviceDidReceiveMetadataForItem func(camera ICCameraDevice, item ICCameraItem)
	_CameraDeviceDidReceiveThumbnailForItem func(camera ICCameraDevice, item ICCameraItem)
	_CameraDeviceDidRemoveItem func(camera ICCameraDevice, item ICCameraItem)
	_CameraDeviceShouldGetMetadataOfItem func(cameraDevice ICCameraDevice, item ICCameraItem) bool
	_CameraDeviceShouldGetThumbnailOfItem func(cameraDevice ICCameraDevice, item ICCameraItem) bool
	_CameraDeviceDidChangeCapability func()
	_DeviceDidBecomeReady func()
	_DeviceDidBecomeReadyWithCompleteContentCatalog func(device ICCameraDevice)
	_CameraDeviceDidEnableAccessRestriction func()
	_CameraDeviceDidRemoveAccessRestriction func()
	_CameraDeviceDidAddItems func(camera ICCameraDevice, items []CCameraItem)
	_CameraDeviceDidReceiveMetadataForItemError func(camera ICCameraDevice, metadata objc.IObject /* cross-framework: NSDictionary */, item ICCameraItem, error_ objc.IObject /* cross-framework: Error */)
	_CameraDeviceDidReceivePTPEvent func(camera ICCameraDevice, eventData objc.IObject /* cross-framework: NSData */)
	_CameraDeviceDidReceiveThumbnailForItemError func(camera ICCameraDevice, thumbnail ImageRef /* not a class type */, item ICCameraItem, error_ objc.IObject /* cross-framework: Error */)
	_CameraDeviceDidRemoveItems func(camera ICCameraDevice, items []CCameraItem)
	_CameraDeviceDidRenameItems func(camera ICCameraDevice, items []CCameraItem)
}

// SetCameraDevice sets the handler for the CameraDevice delegate method.
//
// Tells the client when the metadata requested for an item on a camera is available.
func (d *ICCameraDeviceDelegate) SetCameraDevice(f func()) {
	d._CameraDevice = f
}

// SetCameraDeviceDidAddItem sets the handler for the CameraDeviceDidAddItem delegate method.
//
// Tells the client when an object is added to the device.
func (d *ICCameraDeviceDelegate) SetCameraDeviceDidAddItem(f func(camera ICCameraDevice, item ICCameraItem)) {
	d._CameraDeviceDidAddItem = f
}

// SetCameraDeviceDidCompleteDeleteFilesWithError sets the handler for the CameraDeviceDidCompleteDeleteFilesWithError delegate method.
//
// Tells the client when the camera completes a delete operation.
func (d *ICCameraDeviceDelegate) SetCameraDeviceDidCompleteDeleteFilesWithError(f func(camera ICCameraDevice, error_ objc.IObject /* cross-framework: Error */)) {
	d._CameraDeviceDidCompleteDeleteFilesWithError = f
}

// SetCameraDeviceDidReceiveMetadataForItem sets the handler for the CameraDeviceDidReceiveMetadataForItem delegate method.
//
// Tells the client when the metadata requested for an item on a camera is available.
func (d *ICCameraDeviceDelegate) SetCameraDeviceDidReceiveMetadataForItem(f func(camera ICCameraDevice, item ICCameraItem)) {
	d._CameraDeviceDidReceiveMetadataForItem = f
}

// SetCameraDeviceDidReceiveThumbnailForItem sets the handler for the CameraDeviceDidReceiveThumbnailForItem delegate method.
//
// Tells the client when the requested thumbnail is available.
func (d *ICCameraDeviceDelegate) SetCameraDeviceDidReceiveThumbnailForItem(f func(camera ICCameraDevice, item ICCameraItem)) {
	d._CameraDeviceDidReceiveThumbnailForItem = f
}

// SetCameraDeviceDidRemoveItem sets the handler for the CameraDeviceDidRemoveItem delegate method.
//
// Tells the client when an object is removed from the device.
func (d *ICCameraDeviceDelegate) SetCameraDeviceDidRemoveItem(f func(camera ICCameraDevice, item ICCameraItem)) {
	d._CameraDeviceDidRemoveItem = f
}

// SetCameraDeviceShouldGetMetadataOfItem sets the handler for the CameraDeviceShouldGetMetadataOfItem delegate method.
//
// Tells the client when the camera is about to execute queued requests for the metadata of a specific item.
func (d *ICCameraDeviceDelegate) SetCameraDeviceShouldGetMetadataOfItem(f func(cameraDevice ICCameraDevice, item ICCameraItem) bool) {
	d._CameraDeviceShouldGetMetadataOfItem = f
}

// SetCameraDeviceShouldGetThumbnailOfItem sets the handler for the CameraDeviceShouldGetThumbnailOfItem delegate method.
//
// Tells the client when the camera is about to execute queued requests for the thumbnail of a specific item.
func (d *ICCameraDeviceDelegate) SetCameraDeviceShouldGetThumbnailOfItem(f func(cameraDevice ICCameraDevice, item ICCameraItem) bool) {
	d._CameraDeviceShouldGetThumbnailOfItem = f
}

// SetCameraDeviceDidChangeCapability sets the handler for the CameraDeviceDidChangeCapability delegate method.
//
// Tells the client when a capability of a camera changes.
func (d *ICCameraDeviceDelegate) SetCameraDeviceDidChangeCapability(f func()) {
	d._CameraDeviceDidChangeCapability = f
}

// SetDeviceDidBecomeReady sets the handler for the DeviceDidBecomeReady delegate method.
//
// Tells the client that the camera device is done enumerating its content and is ready to receive requests.
func (d *ICCameraDeviceDelegate) SetDeviceDidBecomeReady(f func()) {
	d._DeviceDidBecomeReady = f
}

// SetDeviceDidBecomeReadyWithCompleteContentCatalog sets the handler for the DeviceDidBecomeReadyWithCompleteContentCatalog delegate method.
//
// Tells the client that the camera device is done enumerating its content and is ready to receive requests.
func (d *ICCameraDeviceDelegate) SetDeviceDidBecomeReadyWithCompleteContentCatalog(f func(device ICCameraDevice)) {
	d._DeviceDidBecomeReadyWithCompleteContentCatalog = f
}

// SetCameraDeviceDidEnableAccessRestriction sets the handler for the CameraDeviceDidEnableAccessRestriction delegate method.
//
// Tells the client when an Apple device has been locked, and media is unavailable until the restriction has been removed.
func (d *ICCameraDeviceDelegate) SetCameraDeviceDidEnableAccessRestriction(f func()) {
	d._CameraDeviceDidEnableAccessRestriction = f
}

// SetCameraDeviceDidRemoveAccessRestriction sets the handler for the CameraDeviceDidRemoveAccessRestriction delegate method.
//
// Tells the client when an Apple device has been unlocked, paired to the host, and media is available.
func (d *ICCameraDeviceDelegate) SetCameraDeviceDidRemoveAccessRestriction(f func()) {
	d._CameraDeviceDidRemoveAccessRestriction = f
}

// SetCameraDeviceDidAddItems sets the handler for the CameraDeviceDidAddItems delegate method.
//
// Tells the client when objects are added to the device.
func (d *ICCameraDeviceDelegate) SetCameraDeviceDidAddItems(f func(camera ICCameraDevice, items []CCameraItem)) {
	d._CameraDeviceDidAddItems = f
}

// SetCameraDeviceDidReceiveMetadataForItemError sets the handler for the CameraDeviceDidReceiveMetadataForItemError delegate method.
//
// Tells the client when the metadata requested for an item on a camera is available.
func (d *ICCameraDeviceDelegate) SetCameraDeviceDidReceiveMetadataForItemError(f func(camera ICCameraDevice, metadata objc.IObject /* cross-framework: NSDictionary */, item ICCameraItem, error_ objc.IObject /* cross-framework: Error */)) {
	d._CameraDeviceDidReceiveMetadataForItemError = f
}

// SetCameraDeviceDidReceivePTPEvent sets the handler for the CameraDeviceDidReceivePTPEvent delegate method.
//
// Tells the client about a PTP event.
func (d *ICCameraDeviceDelegate) SetCameraDeviceDidReceivePTPEvent(f func(camera ICCameraDevice, eventData objc.IObject /* cross-framework: NSData */)) {
	d._CameraDeviceDidReceivePTPEvent = f
}

// SetCameraDeviceDidReceiveThumbnailForItemError sets the handler for the CameraDeviceDidReceiveThumbnailForItemError delegate method.
//
// Tells the client when the requested thumbnail is available.
func (d *ICCameraDeviceDelegate) SetCameraDeviceDidReceiveThumbnailForItemError(f func(camera ICCameraDevice, thumbnail ImageRef /* not a class type */, item ICCameraItem, error_ objc.IObject /* cross-framework: Error */)) {
	d._CameraDeviceDidReceiveThumbnailForItemError = f
}

// SetCameraDeviceDidRemoveItems sets the handler for the CameraDeviceDidRemoveItems delegate method.
//
// Tells the client when objects are removed from the device.
func (d *ICCameraDeviceDelegate) SetCameraDeviceDidRemoveItems(f func(camera ICCameraDevice, items []CCameraItem)) {
	d._CameraDeviceDidRemoveItems = f
}

// SetCameraDeviceDidRenameItems sets the handler for the CameraDeviceDidRenameItems delegate method.
//
// Tells the client when one or more objects are renamed on the device.
func (d *ICCameraDeviceDelegate) SetCameraDeviceDidRenameItems(f func(camera ICCameraDevice, items []CCameraItem)) {
	d._CameraDeviceDidRenameItems = f
}

// CameraDevice implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDevice() {
	if d._CameraDevice != nil {
		d._CameraDevice()
	}
}

// HasCameraDevice returns true if a handler for CameraDevice has been set.
func (d *ICCameraDeviceDelegate) HasCameraDevice() bool {
	return d._CameraDevice != nil
}

// CameraDeviceDidAddItem implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceDidAddItem(camera ICCameraDevice, item ICCameraItem) {
	if d._CameraDeviceDidAddItem != nil {
		d._CameraDeviceDidAddItem(camera, item)
	}
}

// HasCameraDeviceDidAddItem returns true if a handler for CameraDeviceDidAddItem has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceDidAddItem() bool {
	return d._CameraDeviceDidAddItem != nil
}

// CameraDeviceDidCompleteDeleteFilesWithError implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceDidCompleteDeleteFilesWithError(camera ICCameraDevice, error_ objc.IObject /* cross-framework: Error */) {
	if d._CameraDeviceDidCompleteDeleteFilesWithError != nil {
		d._CameraDeviceDidCompleteDeleteFilesWithError(camera, error_)
	}
}

// HasCameraDeviceDidCompleteDeleteFilesWithError returns true if a handler for CameraDeviceDidCompleteDeleteFilesWithError has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceDidCompleteDeleteFilesWithError() bool {
	return d._CameraDeviceDidCompleteDeleteFilesWithError != nil
}

// CameraDeviceDidReceiveMetadataForItem implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceDidReceiveMetadataForItem(camera ICCameraDevice, item ICCameraItem) {
	if d._CameraDeviceDidReceiveMetadataForItem != nil {
		d._CameraDeviceDidReceiveMetadataForItem(camera, item)
	}
}

// HasCameraDeviceDidReceiveMetadataForItem returns true if a handler for CameraDeviceDidReceiveMetadataForItem has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceDidReceiveMetadataForItem() bool {
	return d._CameraDeviceDidReceiveMetadataForItem != nil
}

// CameraDeviceDidReceiveThumbnailForItem implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceDidReceiveThumbnailForItem(camera ICCameraDevice, item ICCameraItem) {
	if d._CameraDeviceDidReceiveThumbnailForItem != nil {
		d._CameraDeviceDidReceiveThumbnailForItem(camera, item)
	}
}

// HasCameraDeviceDidReceiveThumbnailForItem returns true if a handler for CameraDeviceDidReceiveThumbnailForItem has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceDidReceiveThumbnailForItem() bool {
	return d._CameraDeviceDidReceiveThumbnailForItem != nil
}

// CameraDeviceDidRemoveItem implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceDidRemoveItem(camera ICCameraDevice, item ICCameraItem) {
	if d._CameraDeviceDidRemoveItem != nil {
		d._CameraDeviceDidRemoveItem(camera, item)
	}
}

// HasCameraDeviceDidRemoveItem returns true if a handler for CameraDeviceDidRemoveItem has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceDidRemoveItem() bool {
	return d._CameraDeviceDidRemoveItem != nil
}

// CameraDeviceShouldGetMetadataOfItem implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceShouldGetMetadataOfItem(cameraDevice ICCameraDevice, item ICCameraItem) bool {
	if d._CameraDeviceShouldGetMetadataOfItem != nil {
		return d._CameraDeviceShouldGetMetadataOfItem(cameraDevice, item)
	}
	var zero bool
	return zero
}

// HasCameraDeviceShouldGetMetadataOfItem returns true if a handler for CameraDeviceShouldGetMetadataOfItem has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceShouldGetMetadataOfItem() bool {
	return d._CameraDeviceShouldGetMetadataOfItem != nil
}

// CameraDeviceShouldGetThumbnailOfItem implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceShouldGetThumbnailOfItem(cameraDevice ICCameraDevice, item ICCameraItem) bool {
	if d._CameraDeviceShouldGetThumbnailOfItem != nil {
		return d._CameraDeviceShouldGetThumbnailOfItem(cameraDevice, item)
	}
	var zero bool
	return zero
}

// HasCameraDeviceShouldGetThumbnailOfItem returns true if a handler for CameraDeviceShouldGetThumbnailOfItem has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceShouldGetThumbnailOfItem() bool {
	return d._CameraDeviceShouldGetThumbnailOfItem != nil
}

// CameraDeviceDidChangeCapability implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceDidChangeCapability() {
	if d._CameraDeviceDidChangeCapability != nil {
		d._CameraDeviceDidChangeCapability()
	}
}

// HasCameraDeviceDidChangeCapability returns true if a handler for CameraDeviceDidChangeCapability has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceDidChangeCapability() bool {
	return d._CameraDeviceDidChangeCapability != nil
}

// DeviceDidBecomeReady implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) DeviceDidBecomeReady() {
	if d._DeviceDidBecomeReady != nil {
		d._DeviceDidBecomeReady()
	}
}

// HasDeviceDidBecomeReady returns true if a handler for DeviceDidBecomeReady has been set.
func (d *ICCameraDeviceDelegate) HasDeviceDidBecomeReady() bool {
	return d._DeviceDidBecomeReady != nil
}

// DeviceDidBecomeReadyWithCompleteContentCatalog implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) DeviceDidBecomeReadyWithCompleteContentCatalog(device ICCameraDevice) {
	if d._DeviceDidBecomeReadyWithCompleteContentCatalog != nil {
		d._DeviceDidBecomeReadyWithCompleteContentCatalog(device)
	}
}

// HasDeviceDidBecomeReadyWithCompleteContentCatalog returns true if a handler for DeviceDidBecomeReadyWithCompleteContentCatalog has been set.
func (d *ICCameraDeviceDelegate) HasDeviceDidBecomeReadyWithCompleteContentCatalog() bool {
	return d._DeviceDidBecomeReadyWithCompleteContentCatalog != nil
}

// CameraDeviceDidEnableAccessRestriction implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceDidEnableAccessRestriction() {
	if d._CameraDeviceDidEnableAccessRestriction != nil {
		d._CameraDeviceDidEnableAccessRestriction()
	}
}

// HasCameraDeviceDidEnableAccessRestriction returns true if a handler for CameraDeviceDidEnableAccessRestriction has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceDidEnableAccessRestriction() bool {
	return d._CameraDeviceDidEnableAccessRestriction != nil
}

// CameraDeviceDidRemoveAccessRestriction implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceDidRemoveAccessRestriction() {
	if d._CameraDeviceDidRemoveAccessRestriction != nil {
		d._CameraDeviceDidRemoveAccessRestriction()
	}
}

// HasCameraDeviceDidRemoveAccessRestriction returns true if a handler for CameraDeviceDidRemoveAccessRestriction has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceDidRemoveAccessRestriction() bool {
	return d._CameraDeviceDidRemoveAccessRestriction != nil
}

// CameraDeviceDidAddItems implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceDidAddItems(camera ICCameraDevice, items []CCameraItem) {
	if d._CameraDeviceDidAddItems != nil {
		d._CameraDeviceDidAddItems(camera, items)
	}
}

// HasCameraDeviceDidAddItems returns true if a handler for CameraDeviceDidAddItems has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceDidAddItems() bool {
	return d._CameraDeviceDidAddItems != nil
}

// CameraDeviceDidReceiveMetadataForItemError implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceDidReceiveMetadataForItemError(camera ICCameraDevice, metadata objc.IObject /* cross-framework: NSDictionary */, item ICCameraItem, error_ objc.IObject /* cross-framework: Error */) {
	if d._CameraDeviceDidReceiveMetadataForItemError != nil {
		d._CameraDeviceDidReceiveMetadataForItemError(camera, metadata, item, error_)
	}
}

// HasCameraDeviceDidReceiveMetadataForItemError returns true if a handler for CameraDeviceDidReceiveMetadataForItemError has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceDidReceiveMetadataForItemError() bool {
	return d._CameraDeviceDidReceiveMetadataForItemError != nil
}

// CameraDeviceDidReceivePTPEvent implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceDidReceivePTPEvent(camera ICCameraDevice, eventData objc.IObject /* cross-framework: NSData */) {
	if d._CameraDeviceDidReceivePTPEvent != nil {
		d._CameraDeviceDidReceivePTPEvent(camera, eventData)
	}
}

// HasCameraDeviceDidReceivePTPEvent returns true if a handler for CameraDeviceDidReceivePTPEvent has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceDidReceivePTPEvent() bool {
	return d._CameraDeviceDidReceivePTPEvent != nil
}

// CameraDeviceDidReceiveThumbnailForItemError implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceDidReceiveThumbnailForItemError(camera ICCameraDevice, thumbnail ImageRef /* not a class type */, item ICCameraItem, error_ objc.IObject /* cross-framework: Error */) {
	if d._CameraDeviceDidReceiveThumbnailForItemError != nil {
		d._CameraDeviceDidReceiveThumbnailForItemError(camera, thumbnail, item, error_)
	}
}

// HasCameraDeviceDidReceiveThumbnailForItemError returns true if a handler for CameraDeviceDidReceiveThumbnailForItemError has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceDidReceiveThumbnailForItemError() bool {
	return d._CameraDeviceDidReceiveThumbnailForItemError != nil
}

// CameraDeviceDidRemoveItems implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceDidRemoveItems(camera ICCameraDevice, items []CCameraItem) {
	if d._CameraDeviceDidRemoveItems != nil {
		d._CameraDeviceDidRemoveItems(camera, items)
	}
}

// HasCameraDeviceDidRemoveItems returns true if a handler for CameraDeviceDidRemoveItems has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceDidRemoveItems() bool {
	return d._CameraDeviceDidRemoveItems != nil
}

// CameraDeviceDidRenameItems implements the PICCameraDeviceDelegate interface.
func (d *ICCameraDeviceDelegate) CameraDeviceDidRenameItems(camera ICCameraDevice, items []CCameraItem) {
	if d._CameraDeviceDidRenameItems != nil {
		d._CameraDeviceDidRenameItems(camera, items)
	}
}

// HasCameraDeviceDidRenameItems returns true if a handler for CameraDeviceDidRenameItems has been set.
func (d *ICCameraDeviceDelegate) HasCameraDeviceDidRenameItems() bool {
	return d._CameraDeviceDidRenameItems != nil
}
