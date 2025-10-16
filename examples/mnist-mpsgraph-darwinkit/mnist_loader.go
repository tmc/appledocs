package main

import (
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	// Updated URLs - trying different mirrors
	trainImagesURL = "https://storage.googleapis.com/cvdf-datasets/mnist/train-images-idx3-ubyte.gz"
	trainLabelsURL = "https://storage.googleapis.com/cvdf-datasets/mnist/train-labels-idx1-ubyte.gz"
)

type MNISTDataset struct {
	Images [][]float32 // Each image is 784 (28x28) floats
	Labels []int
}

// LoadMNIST downloads and loads the MNIST dataset
func LoadMNIST() (*MNISTDataset, error) {
	cacheDir := filepath.Join(os.TempDir(), "mnist_cache")
	os.MkdirAll(cacheDir, 0755)

	imagesPath := filepath.Join(cacheDir, "train-images-idx3-ubyte")
	labelsPath := filepath.Join(cacheDir, "train-labels-idx1-ubyte")

	// Download if not cached
	if _, err := os.Stat(imagesPath); os.IsNotExist(err) {
		fmt.Println("Downloading MNIST images...")
		if err := downloadAndExtract(trainImagesURL, imagesPath); err != nil {
			return nil, fmt.Errorf("failed to download images: %w", err)
		}
	}

	if _, err := os.Stat(labelsPath); os.IsNotExist(err) {
		fmt.Println("Downloading MNIST labels...")
		if err := downloadAndExtract(trainLabelsURL, labelsPath); err != nil {
			return nil, fmt.Errorf("failed to download labels: %w", err)
		}
	}

	// Load images
	images, err := loadImages(imagesPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load images: %w", err)
	}

	// Load labels
	labels, err := loadLabels(labelsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load labels: %w", err)
	}

	fmt.Printf("Loaded %d MNIST images\n", len(images))

	return &MNISTDataset{
		Images: images,
		Labels: labels,
	}, nil
}

func downloadAndExtract(url, outputPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP error: %d %s", resp.StatusCode, resp.Status)
	}

	// Download to temp file first
	tmpFile, err := os.CreateTemp("", "mnist-*.gz")
	if err != nil {
		return err
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)

	_, err = io.Copy(tmpFile, resp.Body)
	tmpFile.Close()
	if err != nil {
		return err
	}

	// Now extract the gzipped file
	gzFile, err := os.Open(tmpPath)
	if err != nil {
		return err
	}
	defer gzFile.Close()

	gz, err := gzip.NewReader(gzFile)
	if err != nil {
		return fmt.Errorf("gzip error (file may not be gzipped): %w", err)
	}
	defer gz.Close()

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, gz)
	return err
}

func loadImages(path string) ([][]float32, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read header
	var magic, numImages, rows, cols int32
	binary.Read(f, binary.BigEndian, &magic)
	binary.Read(f, binary.BigEndian, &numImages)
	binary.Read(f, binary.BigEndian, &rows)
	binary.Read(f, binary.BigEndian, &cols)

	if magic != 2051 {
		return nil, fmt.Errorf("invalid magic number: %d", magic)
	}

	images := make([][]float32, numImages)
	imageSize := int(rows * cols)

	for i := 0; i < int(numImages); i++ {
		image := make([]byte, imageSize)
		if _, err := io.ReadFull(f, image); err != nil {
			return nil, err
		}

		// Convert to float32 and normalize to [0, 1]
		floatImage := make([]float32, imageSize)
		for j := 0; j < imageSize; j++ {
			floatImage[j] = float32(image[j]) / 255.0
		}
		images[i] = floatImage
	}

	return images, nil
}

func loadLabels(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read header
	var magic, numLabels int32
	binary.Read(f, binary.BigEndian, &magic)
	binary.Read(f, binary.BigEndian, &numLabels)

	if magic != 2049 {
		return nil, fmt.Errorf("invalid magic number: %d", magic)
	}

	labels := make([]int, numLabels)
	labelBytes := make([]byte, numLabels)

	if _, err := io.ReadFull(f, labelBytes); err != nil {
		return nil, err
	}

	for i := 0; i < int(numLabels); i++ {
		labels[i] = int(labelBytes[i])
	}

	return labels, nil
}
