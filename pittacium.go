package pittacium

import (
	"fmt"
	"image/png"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/boombuler/barcode/ean"
	"github.com/go-pdf/fpdf"
	"github.com/nfnt/resize"
)

// FIXME: use a temp directory
const imagePath = "barcodes"

// LabelFormat represents the dimensions of a label format
type LabelFormat struct {
	Name     string
	WidthMM  float64 // Width in millimeters
	HeightMM float64 // Height in millimeters
}

type Product struct {
	Code  string
	Name  string
	Price float64
}

type Document struct {
	Title        string
	MarginTop    float64
	MarginRight  float64
	MarginLeft   float64
	MarginBottom float64
	ShowPage     bool
	PageSize     string
	Orientation  string
	LabelFormat  LabelFormat
}

// Common label formats
var (
	Label2x1     = LabelFormat{Name: "2x1", WidthMM: 50.8, HeightMM: 25.4}
	Label3x1     = LabelFormat{Name: "3x1", WidthMM: 76.2, HeightMM: 25.4}
	Label4x2     = LabelFormat{Name: "4x2", WidthMM: 101.6, HeightMM: 50.8}
	Label225x125 = LabelFormat{Name: "2.25x1.25", WidthMM: 57.15, HeightMM: 31.75}
	Label3x2     = LabelFormat{Name: "3x2", WidthMM: 76.2, HeightMM: 50.8}
)

func GenerateLabels(products []Product, document Document) *fpdf.Fpdf {
	// Create image directory if it doesn't exist
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		os.Mkdir(imagePath, 0755)
	}

	// Generating barcodes from products
	generateBarcode(products)

	// Generating the PDF from the products
	pdf := generatePDF(document, products)

	// Delete the images folder
	os.RemoveAll(imagePath)

	// Output PDF to a file
	// date := fmt.Sprintf("%d-%02d-%02d", time.Now().Year(), time.Now().Month(), time.Now().Day())
	// err := pdf.OutputFileAndClose(date + "-" + "products-labels-" + document.LabelFormat.Name + ".pdf")

	// if err != nil {
		// log.Fatalf("Error saving PDF: %v", err)
	// }

	// fmt.Println("PDF generated successfully!")

	return pdf
}

// Generate barcode image and save it in the imagePath directory
func generateBarcode(products []Product) {
	for _, product := range products {
		barCode, _ := ean.Encode(product.Code)
		newImage := resize.Resize(uint(barCode.Bounds().Dx()), 22, barCode, resize.Lanczos3)
		file, _ := os.Create(filepath.Join(imagePath, product.Code+".png"))
		defer file.Close()
		png.Encode(file, newImage)
	}
}

// Generate PDF from products
func generatePDF(document Document, products []Product) *fpdf.Fpdf {
	// Create new PDF instance
	pdf := fpdf.New(document.Orientation, "mm", document.PageSize, "")
	
	// Retrieve page dimensions
	pageWidth, pageHeight := getPageDimensions(document.PageSize, document.Orientation)

	// Calculate appropriate margins according to the label format and page size and orientation
	// TODO: add support for custom margins or fixed margins

	// Calculate the number of columns and rows per page based on the label format
	columns := int(math.Floor(pageWidth / document.LabelFormat.WidthMM))
	rows := int(math.Floor(pageHeight / document.LabelFormat.HeightMM))
	
	fmt.Println(columns, rows)

	// Calculate the number of pages
	pageCount := int(math.Ceil(float64(len(products)) / float64(columns*rows)))
	fmt.Println(pageCount)

	// Add products and labels to each page
	for page := 1; page <= pageCount; page++ {
		// Add new page
		pdf.AddPage()
		// Add labels for products on the current page
		for i := (page - 1) * columns * rows; i < int(math.Min(float64(len(products)), float64(page*columns*rows))); i++ {
			product := products[i]
			// Calculate position of the label
			row := (i - (page-1)*columns*rows) / columns
			col := (i - (page-1)*columns*rows) % columns
			x := float64(col) * document.LabelFormat.WidthMM
			y := float64(row) * document.LabelFormat.HeightMM

			// Draw rectangle representing the label
			pdf.Rect(x, y, document.LabelFormat.WidthMM, document.LabelFormat.HeightMM, "D")

			// Add product name
			pdf.SetXY(x, y)
			pdf.SetFont("Arial", "B", 10)
			tr := pdf.UnicodeTranslatorFromDescriptor("")
			pdf.CellFormat(document.LabelFormat.WidthMM, 10, tr(product.Name), "", 2, "C", false, 0, "")
			
			// Add product price
			pdf.SetXY(x, y+6)
			pdf.SetFont("Arial", "B", 12)
			pdf.CellFormat(document.LabelFormat.WidthMM, 10, tr(fmt.Sprintf("%.2f€", product.Price)), "", 2, "C", false, 0, "")

			// Add barcode image (assuming the barcode images are named with the product code)
			var opt fpdf.ImageOptions
			barcodeImagePath := fmt.Sprintf("barcodes/%s.png", product.Code)
			pdf.ImageOptions(barcodeImagePath, x+10, y+20, document.LabelFormat.WidthMM-20, 0, false, opt, 0, "")
			
			// Add product code
			pdf.SetXY(x, y+12)
			pdf.SetFont("Arial", "", 9)
			pdf.CellFormat(document.LabelFormat.WidthMM, 10, product.Code, "", 2, "C", false, 0, "")
			
		}
	}
	return pdf
}

// Function to calculate page dimensions based on PageSize and Orientation
func getPageDimensions(pageSize string, orientation string) (float64, float64) {
	var pageWidth, pageHeight float64

	switch pageSize {
	case "A4":
		// A4 dimensions in millimeters
		pageWidth = 210.0
		pageHeight = 297.0
		// Adjust dimensions for landscape orientation
		if orientation == "L" {
			pageWidth, pageHeight = pageHeight, pageWidth
		}
	case "A3":
		// A3 dimensions in millimeters
		pageWidth = 297.0
		pageHeight = 420.0
		// Adjust dimensions for landscape orientation
		if orientation == "L" {
			pageWidth, pageHeight = pageHeight, pageWidth
		}
	default:
		// Default to A4 dimensions if PageSize is unknown
		pageWidth = 210.0
		pageHeight = 297.0
	}

	return pageWidth, pageHeight
}
