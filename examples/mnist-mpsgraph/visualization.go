package main

import (
	"fmt"
	"math"

	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

type VisualizationView struct {
	imageView        appkit.ImageView
	image            appkit.Image
	bitmapRep        appkit.BitmapImageRep
	width            int
	height           int
	smoothedAccuracy float32
	smoothedLoss     float32
	accuracyHistory  []float32
	lossHistory      []float32
	maxHistory       int
	lastSamples      []SampleData // Cache last samples for stable display
	sampleUpdateCount int         // Update samples every N calls
}

type SampleData struct {
	Image      []float32 // 28x28 pixel data
	Label      int       // Actual label
	Predicted  int       // Predicted label
	Confidence float32   // Prediction confidence (0-1)
}

func NewVisualizationView(width, height int) *VisualizationView {
	// Create a bitmap image representation to draw into
	imageSize := foundation.Size{Width: float64(width), Height: float64(height)}
	bitmapRep := appkit.BitmapImageRepClass.Alloc().InitWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBytesPerRowBitsPerPixel(
		nil,
		width, height,
		8, 4,
		true, false,
		appkit.CalibratedRGBColorSpace,
		0, 0,
	)
	objc.Retain(&bitmapRep)

	// Create an image and add the representation
	image := appkit.ImageClass.Alloc().InitWithSize(imageSize)
	objc.Retain(&image)
	image.AddRepresentation(bitmapRep)

	// Create an image view to display the image
	imageView := appkit.NewImageView()
	objc.Retain(&imageView)
	imageView.SetFrame(foundation.Rect{
		Size: imageSize,
	})
	imageView.SetImage(image)

	return &VisualizationView{
		imageView:       imageView,
		image:           image,
		bitmapRep:       bitmapRep,
		width:           width,
		height:          height,
		accuracyHistory: make([]float32, 0, 100),
		maxHistory:      100,
	}
}

func (v *VisualizationView) View() appkit.View {
	return v.imageView.View
}

func (v *VisualizationView) Update(samples []SampleData, testAccuracy, trainLoss float32, iteration int) {
	// Update accuracy history for smoothing
	v.accuracyHistory = append(v.accuracyHistory, testAccuracy)
	if len(v.accuracyHistory) > v.maxHistory {
		v.accuracyHistory = v.accuracyHistory[1:]
	}

	// Update loss history
	v.lossHistory = append(v.lossHistory, trainLoss)
	if len(v.lossHistory) > v.maxHistory {
		v.lossHistory = v.lossHistory[1:]
	}

	// Calculate smoothed metrics (exponential moving average with stronger smoothing)
	alpha := float32(0.05) // Reduced from 0.1 for more stable display
	if v.smoothedAccuracy == 0 {
		v.smoothedAccuracy = testAccuracy
		v.smoothedLoss = trainLoss
	} else {
		v.smoothedAccuracy = alpha*testAccuracy + (1-alpha)*v.smoothedAccuracy
		v.smoothedLoss = alpha*trainLoss + (1-alpha)*v.smoothedLoss
	}

	// Update samples only every 5 calls for more stable display
	v.sampleUpdateCount++
	if v.sampleUpdateCount >= 5 || len(v.lastSamples) == 0 {
		v.lastSamples = samples
		v.sampleUpdateCount = 0
	}

	// Create graphics context from bitmap
	ctx := appkit.GraphicsContext_GraphicsContextWithBitmapImageRep(v.bitmapRep)
	appkit.GraphicsContextClass.SetCurrentContext(ctx)

	// Draw visualization using cached samples
	v.draw(v.lastSamples, testAccuracy, trainLoss, iteration)

	// Trigger redraw
	v.imageView.SetNeedsDisplay(true)
}

func (v *VisualizationView) draw(samples []SampleData, testAccuracy, trainLoss float32, iteration int) {
	// Dark background
	bgColor := appkit.Color_ColorWithCalibratedRedGreenBlueAlpha(0.1, 0.1, 0.1, 1.0)
	bgColor.SetFill()
	appkit.BezierPath_FillRect(foundation.Rect{
		Origin: foundation.Point{X: 0, Y: 0},
		Size:   foundation.Size{Width: float64(v.width), Height: float64(v.height)},
	})

	// Draw summary statistics panel
	panelColor := appkit.Color_ColorWithCalibratedRedGreenBlueAlpha(0.15, 0.15, 0.15, 1.0)
	panelColor.SetFill()
	appkit.BezierPath_FillRect(foundation.Rect{
		Origin: foundation.Point{X: 10, Y: float64(v.height) - 120},
		Size:   foundation.Size{Width: float64(v.width) - 20, Height: 110},
	})

	// Draw statistics
	textColor := appkit.Color_WhiteColor()
	textColor.Set()

	// Training statistics
	stats := fmt.Sprintf("Iteration: %d / 5000", iteration)
	v.drawText(stats, 20, float64(v.height)-30, 18, textColor)

	// Test accuracy (smoothed)
	accuracyText := fmt.Sprintf("Test Accuracy: %.1f%% (smoothed: %.1f%%)", testAccuracy*100, v.smoothedAccuracy*100)
	v.drawText(accuracyText, 20, float64(v.height)-55, 16, textColor)

	// Training loss (smoothed)
	lossText := fmt.Sprintf("Training Loss: %.4f (smoothed: %.4f)", trainLoss, v.smoothedLoss)
	v.drawText(lossText, 20, float64(v.height)-80, 14, textColor)

	// Average confidence
	avgConf := float32(0)
	for _, s := range samples {
		avgConf += s.Confidence
	}
	if len(samples) > 0 {
		avgConf /= float32(len(samples))
	}
	confText := fmt.Sprintf("Avg Confidence: %.1f%%", avgConf*100)
	v.drawText(confText, 20, float64(v.height)-105, 14, textColor)

	// Draw grid of samples (4x4 grid)
	gridSize := 4
	cellSize := 70
	padding := 8
	startX := 20
	startY := float64(v.height) - 140

	numSamples := len(samples)
	if numSamples > gridSize*gridSize {
		numSamples = gridSize * gridSize
	}

	// Draw all samples (up to gridSize*gridSize)
	for i := 0; i < numSamples; i++ {
		row := i / gridSize
		col := i % gridSize
		x := float64(startX + col*(cellSize+padding))
		y := startY - float64(row*(cellSize+padding+30))

		sample := samples[i]

		// Draw digit image
		v.drawDigit(sample.Image, x, y, cellSize)

		// Draw label with confidence (green for correct, red for incorrect)
		var labelColor appkit.Color
		if sample.Predicted == sample.Label {
			labelColor = appkit.Color_ColorWithCalibratedRedGreenBlueAlpha(0.0, 1.0, 0.0, 1.0)
		} else {
			labelColor = appkit.Color_ColorWithCalibratedRedGreenBlueAlpha(1.0, 0.3, 0.0, 1.0)
		}
		label := fmt.Sprintf("L:%d P:%d (%.0f%%)", sample.Label, sample.Predicted, sample.Confidence*100)
		v.drawText(label, x, y-20, 9, labelColor)
	}
}

func (v *VisualizationView) drawDigit(pixels []float32, x, y float64, size int) {
	// Draw 28x28 digit as grayscale squares
	pixelSize := float64(size) / 28.0

	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			pixelValue := pixels[i*28+j]

			// Create grayscale color
			color := appkit.Color_ColorWithCalibratedRedGreenBlueAlpha(
				float64(pixelValue),
				float64(pixelValue),
				float64(pixelValue),
				1.0,
			)
			color.SetFill()

			rect := foundation.Rect{
				Origin: foundation.Point{
					X: x + float64(j)*pixelSize,
					Y: y + float64(27-i)*pixelSize, // Flip Y
				},
				Size: foundation.Size{
					Width:  pixelSize,
					Height: pixelSize,
				},
			}
			appkit.BezierPath_FillRect(rect)
		}
	}
}

func (v *VisualizationView) drawText(text string, x, y, size float64, color appkit.Color) {
	// Create attributed string with font and color
	font := appkit.Font_SystemFontOfSize(size)

	// Create attributes dictionary using objc
	attrs := foundation.NewMutableDictionary()
	objc.Call[objc.Void](attrs, objc.Sel("setObject:forKey:"), font, foundation.String_StringWithString("NSFont"))
	objc.Call[objc.Void](attrs, objc.Sel("setObject:forKey:"), color, foundation.String_StringWithString("NSForegroundColor"))

	// Create NSString and draw it
	str := foundation.String_StringWithString(text)
	objc.Call[objc.Void](str, objc.Sel("drawAtPoint:withAttributes:"),
		foundation.Point{X: x, Y: y},
		attrs,
	)
}

// Helper to generate random MNIST-like sample for visualization
func generateRandomSample() SampleData {
	pixels := make([]float32, 28*28)
	label := int(math.Floor(math.Mod(float64(len(pixels)), 10)))

	// Generate simple pattern
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			// Create a simple digit-like pattern
			centerX, centerY := 14.0, 14.0
			dist := math.Sqrt(math.Pow(float64(i)-centerX, 2) + math.Pow(float64(j)-centerY, 2))
			if dist < 8 {
				pixels[i*28+j] = float32(1.0 - dist/8.0)
			}
		}
	}

	predicted := label
	if math.Mod(float64(label), 3) == 0 {
		predicted = (label + 1) % 10 // Simulate some errors
	}

	return SampleData{
		Image:     pixels,
		Label:     label,
		Predicted: predicted,
	}
}
