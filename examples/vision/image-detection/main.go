package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/vision"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("Vision Framework Examples")
	fmt.Println("=========================")

	// Example 1: Create Vision requests
	fmt.Println("\n1. Creating Vision Requests:")

	// Face detection request
	faceRequest := vision.NewDetectFaceRectanglesRequest()
	fmt.Printf("   Face rectangles detection request: %v\n", faceRequest)

	// Image classification request
	classifyRequest := vision.NewClassifyImageRequest()
	fmt.Printf("   Image classification request: %v\n", classifyRequest)

	// Saliency request (what draws attention in an image)
	saliencyRequest := vision.NewGenerateAttentionBasedSaliencyImageRequest()
	fmt.Printf("   Attention-based saliency request: %v\n", saliencyRequest)

	// Example 2: Image Request Handler creation
	fmt.Println("\n2. Image Request Handler:")

	// Create handler (requires actual image data in real use)
	handler := vision.NewImageRequestHandler()
	fmt.Printf("   Image request handler created: %v\n", handler)

	// Example 3: Available Vision capabilities
	fmt.Println("\n3. Vision Framework Capabilities:")

	capabilities := []string{
		"Face Detection & Landmarks",
		"Text Recognition (OCR)",
		"Barcode Detection",
		"Object Tracking",
		"Image Classification",
		"Saliency Analysis",
		"Horizon Detection",
		"Rectangle Detection",
		"Human Body Pose Detection",
		"Animal Detection",
		"Document Segmentation",
		"Contour Detection",
	}

	for i, cap := range capabilities {
		fmt.Printf("   %2d. %s\n", i+1, cap)
	}

	// Example 4: Common Vision workflows
	fmt.Println("\n4. Common Vision Workflows:")

	workflows := map[string]string{
		"Face Detection": "VNDetectFaceRectanglesRequest → Process faces",
		"Text Recognition": "VNRecognizeTextRequest → Extract text from images",
		"Object Tracking": "VNTrackObjectRequest → Track objects across frames",
		"Classification": "VNClassifyImageRequest → Categorize image content",
		"QR Code Scanning": "VNDetectBarcodesRequest → Decode barcodes/QR codes",
	}

	for name, flow := range workflows {
		fmt.Printf("   %-20s: %s\n", name, flow)
	}

	// Example 5: Request types by category
	fmt.Println("\n5. Vision Request Categories:")

	categories := map[string][]string{
		"Detection": {
			"Face Rectangles",
			"Face Landmarks",
			"Human Rectangles",
			"Text Rectangles",
			"Barcodes",
			"Rectangles",
			"Horizon",
		},
		"Recognition": {
			"Text (OCR)",
			"Image Classification",
			"Object Classification",
		},
		"Analysis": {
			"Saliency",
			"Image Aesthetics",
			"Face Quality",
			"Document Segmentation",
		},
		"Tracking": {
			"Object Tracking",
			"Rectangle Tracking",
			"Trajectory Detection",
		},
		"Body/Pose": {
			"Human Body Pose",
			"Human Hand Pose",
			"Animal Body Pose",
			"Face Landmarks",
		},
	}

	for category, items := range categories {
		fmt.Printf("\n   %s:\n", category)
		for _, item := range items {
			fmt.Printf("     • %s\n", item)
		}
	}

	fmt.Println("\n✓ Vision framework examples completed!")
	fmt.Println("\nNote: This example demonstrates Vision request creation.")
	fmt.Println("Real applications would:")
	fmt.Println("  - Load images from files or camera")
	fmt.Println("  - Process requests with VNImageRequestHandler")
	fmt.Println("  - Handle completion callbacks")
	fmt.Println("  - Process results and observations")
	fmt.Println("\nFor image processing, you would:")
	fmt.Println("  1. Create a request (VNDetectFaceRectanglesRequest, etc.)")
	fmt.Println("  2. Create a handler with image data")
	fmt.Println("  3. Perform request: handler.perform([request])")
	fmt.Println("  4. Access results through request.results")
}
