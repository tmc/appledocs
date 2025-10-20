# Vision Framework Image Detection Example

This example demonstrates the Vision framework's image analysis capabilities in Go.

## What it demonstrates

- Creating Vision request objects
- Available Vision detection capabilities
- Common Vision workflows
- Request categories and types
- Framework overview and capabilities

## Running the example

```bash
go run main.go
# or with e2e flag
go run main.go -e2e
```

## Key Concepts

### Vision Framework

Vision is Apple's computer vision framework providing:
- Face detection and analysis
- Text recognition (OCR)
- Object detection and tracking
- Image classification
- Barcode/QR code scanning
- Body and pose detection
- Saliency analysis
- And much more

### Request-Based Architecture

Vision uses a request-based pattern:

```go
// 1. Create a request
request := vision.NewDetectFaceRectanglesRequest()

// 2. Create handler with image
handler := vision.NewImageRequestHandlerWithURL(imageURL, nil)

// 3. Perform request
handler.PerformRequests([]Request{request})

// 4. Get results
results := request.Results()
```

### Common Request Types

**Detection Requests:**
- `VNDetectFaceRectanglesRequest` - Detect faces
- `VNDetectTextRectanglesRequest` - Find text regions
- `VNDetectBarcodesRequest` - Scan barcodes/QR codes
- `VNDetectHumanRectanglesRequest` - Detect people
- `VNDetectRectanglesRequest` - Find rectangles

**Recognition Requests:**
- `VNRecognizeTextRequest` - OCR (text recognition)
- `VNClassifyImageRequest` - Categorize images
- `VNGenerateImageFeaturePrintRequest` - Image fingerprinting

**Analysis Requests:**
- `VNGenerateAttentionBasedSaliencyImageRequest` - What draws attention
- `VNCalculateImageAestheticsScoresRequest` - Image quality
- `VNDetectFaceCaptureQualityRequest` - Face photo quality

**Tracking Requests:**
- `VNTrackObjectRequest` - Track objects in video
- `VNTrackRectangleRequest` - Track rectangles
- `VNDetectTrajectoriesRequest` - Detect motion paths

**Body/Pose Requests:**
- `VNDetectHumanBodyPoseRequest` - Full body pose
- `VNDetectHumanHandPoseRequest` - Hand pose/gestures
- `VNDetectAnimalBodyPoseRequest` - Animal pose
- `VNDetectFaceLandmarksRequest` - Facial features

## Real-World Usage

### Face Detection Example

```go
// Load image
url := foundation.NewURLWithString("file:///path/to/image.jpg")

// Create face detection request
faceRequest := vision.NewDetectFaceRectanglesRequest()

// Create handler with image
handler := vision.NewImageRequestHandlerWithURLOptions(
    unsafe.Pointer(url.ID),
    nil, // options
)

// Perform request
requestsArray := []vision.Request{faceRequest}
handler.PerformRequests(requestsArray)

// Get results
results := faceRequest.Results()
// Process face rectangles...
```

### Text Recognition (OCR)

```go
// Create text recognition request
textRequest := vision.NewRecognizeTextRequest()

// Set recognition level
// VNRequestTextRecognitionLevelAccurate = highest accuracy

// Perform request with handler
handler.PerformRequests([]vision.Request{textRequest})

// Get recognized text
for _, observation := range textRequest.Results() {
    text := observation.Text()
    confidence := observation.Confidence()
    // Use recognized text...
}
```

### QR Code Scanning

```go
// Create barcode detection request
barcodeRequest := vision.NewDetectBarcodesRequest()

// Perform request
handler.PerformRequests([]vision.Request{barcodeRequest})

// Get barcodes
for _, observation := range barcodeRequest.Results() {
    payload := observation.PayloadStringValue()
    symbology := observation.Symbology()
    // Process QR code data...
}
```

## Use Cases

### Photo Organization
- Face detection for photo albums
- Scene classification for categorization
- Duplicate detection via fingerprints

### Document Processing
- OCR for document digitization
- Document segmentation
- Table/form detection

### AR/Computer Vision
- Object tracking in video
- Body pose for fitness apps
- Hand gestures for UI control

### Quality Control
- Image aesthetics scoring
- Face photo quality assessment
- Horizon detection for straightening

### Security
- Face detection for authentication
- QR code scanning
- Document verification

## Performance Considerations

1. **Request Reuse**: Create requests once, reuse for multiple images
2. **Batch Processing**: Process multiple requests together when possible
3. **Background Processing**: Run Vision on background threads (not main)
4. **Resolution**: Lower resolution = faster processing
5. **Request Options**: Configure accuracy vs speed tradeoffs

## Image Input Sources

Vision accepts images from:
- File URLs (`NewImageRequestHandlerWithURLOptions`)
- In-memory data (`NewImageRequestHandlerWithDataOptions`)
- Core Video pixel buffers (`NewImageRequestHandlerWithCVPixelBufferOptions`)
- Core Graphics images
- CIImage objects

## Results and Observations

Different requests return different observation types:
- `VNFaceObservation` - Face detection results
- `VNRecognizedTextObservation` - OCR results
- `VNBarcodeObservation` - Barcode/QR results
- `VNClassificationObservation` - Classification results
- `VNHumanBodyPoseObservation` - Body pose results

Each observation includes:
- Bounding box (`boundingBox`)
- Confidence score (`confidence`)
- Type-specific data

## Error Handling

```go
// Perform requests returns error
err := handler.PerformRequests(requests)
if err != nil {
    // Handle error
}

// Check observation confidence
for _, obs := range results {
    if obs.Confidence() > 0.8 {
        // High confidence result
    }
}
```

## Best Practices

1. **Check availability**: Some features require specific iOS/macOS versions
2. **Handle failures gracefully**: Vision can fail on corrupted images
3. **Validate confidence scores**: Use appropriate thresholds
4. **Optimize image size**: Larger images = slower processing
5. **Use appropriate request level**: Balance speed vs accuracy
6. **Clean up resources**: Release handlers and requests when done

## Integration with Other Frameworks

Vision works with:
- **Core Image**: Image preprocessing
- **Core ML**: Custom vision models
- **AVFoundation**: Real-time camera analysis
- **Photos**: Photo library analysis
- **UIKit/AppKit**: Display results

## References

- [Vision Documentation](https://developer.apple.com/documentation/vision)
- [Vision Programming Guide](https://developer.apple.com/documentation/vision/recognizing_text_in_images)
- [VNRequest](https://developer.apple.com/documentation/vision/vnrequest)
- [VNImageRequestHandler](https://developer.apple.com/documentation/vision/vnimagerequesthandler)
