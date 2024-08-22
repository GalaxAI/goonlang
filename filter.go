package main

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
)

func getImageMatrix(img image.Image) [][][]uint8 {
	shape := img.Bounds().Max
	X, Y := shape.X, shape.Y

	/*
		Matrix should look like
		(3, 512, 512) where:
		3 is number of dims
		3 for RGB
		1 for black white
		512,512 is a 2D matrix that has value
	*/

	// Creates RGB part should work same with black white version for now, this could use optimization.
	matrix := make([][][]uint8, Y)
	// iter over [:]
	for i := range matrix {
		// Creates X, Y pixel values
		matrix[i] = make([][]uint8, X)
		for j := range matrix {
			R, G, B, _ := img.At(i, j).RGBA()
			// bitwise operation to normalize image values
			matrix[i][j] = []uint8{uint8(R >> 8), uint8(G >> 8), uint8(B >> 8)}
		}
	}
	return matrix
}

func downSample(matrix [][][]uint8, kernelSize int) [][][]uint8 {
	X, Y, dim := len(matrix), len(matrix[0]), len(matrix[0][0])

	newY := Y / kernelSize
	newX := X / kernelSize
	kernelArea := kernelSize * kernelSize

	downSampledMatrix := make([][][]uint8, newY)
	for i := range downSampledMatrix {
		downSampledMatrix[i] = make([][]uint8, newX)
		for j := range downSampledMatrix[i] {
			sum := make([]uint64, dim)
			// Gets average from the kernel window
			for ky := 0; ky < kernelSize; ky++ {
				y := i*kernelSize + ky
				for kx := 0; kx < kernelSize; kx++ {
					x := j*kernelSize + kx
					for d := 0; d < dim; d++ {
						sum[d] += uint64(matrix[y][x][d])
					}
				}
			}
			avg := make([]uint8, dim)
			for d := 0; d < dim; d++ {
				avg[d] = uint8(sum[d] / uint64(kernelArea))
			}
			downSampledMatrix[i][j] = avg
		}
	}
	return downSampledMatrix
}

func exportImage(matrix [][][]uint8, fileName string) {
	Y, X := len(matrix), len(matrix[1])
	imgCanvas := image.NewNRGBA(image.Rect(0, 0, X, Y))

	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			r := matrix[y][x][0]
			g := matrix[y][x][1]
			b := matrix[y][x][2]
			imgCanvas.Set(y, x, color.RGBA{r, g, b, 255}) // Assuming full opacity
		}
	}

	// Create a file to save the image
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Encode and save the image as PNG
	err = png.Encode(f, imgCanvas)
	if err != nil {
		panic(err)
	}
}

// func main() {

// 	//open the image
// 	file, err := os.Open("static/test.png")
// 	if err != nil {
// 		fmt.Println("Error opening image file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	img, _, err := image.Decode(file)
// 	if err != nil {
// 		fmt.Println("Error decoding image :", err)
// 		return
// 	}
// 	matrix := getImageMatrix(img)
// 	dsmatrix := downSample(matrix, 8)
// 	exportImage(dsmatrix, "static/downsapled.png")
// }
