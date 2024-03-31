package pittacium

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestGenerateLabels(t *testing.T) {

	// Create a document
	document := Document{
		Title:       "Products",
		PageSize:    "A4",
		Orientation: "P",
		LabelFormat: Label225x125,
	}

	// A list of products
	products := []Product{
		{Code: "5400141850574", Name: "Cereal Flakes", Price: 1.72},
		{Code: "5410091734367", Name: "Déo Divine Moments (Fa)", Price: 1.66},
		{Code: "540091719838", Name: "Déo Glamorous Moments (Fa)", Price: 1.66},
		{Code: "5410091729745", Name: "Déo Pink Passion (Fa)", Price: 2.99},
		{Code: "5400141437102", Name: "Eponges x6 (Boni)", Price: 2.38},
		{Code: "5400141358858", Name: "Frost Cribb (Boni)", Price: 2.98},
		{Code: "5400141367461", Name: "Haricots blancs Bio (Boni)", Price: 1.45},
		{Code: "5400141058383", Name: "Haricots verts entier", Price: 1.49},
		{Code: "5400141710595", Name: "Lait de coco (Boni)", Price: 0.99},
		{Code: "5400141850574", Name: "Cereal Flakes", Price: 1.72},
		{Code: "5410091734367", Name: "Déo Divine Moments (Fa)", Price: 1.66},
		{Code: "540091719838", Name: "Déo Glamorous Moments (Fa)", Price: 1.66},
		{Code: "5410091729745", Name: "Déo Pink Passion (Fa)", Price: 2.99},
		{Code: "5400141437102", Name: "Eponges x6 (Boni)", Price: 2.38},
		{Code: "5400141358858", Name: "Frost Cribb (Boni)", Price: 2.98},
		{Code: "5400141367461", Name: "Haricots blancs Bio (Boni)", Price: 1.45},
		{Code: "5400141058383", Name: "Haricots verts entier", Price: 1.49},
		{Code: "5400141710595", Name: "Lait de coco (Boni)", Price: 0.99},
		{Code: "5400141850574", Name: "Cereal Flakes", Price: 1.72},
		{Code: "5410091734367", Name: "Déo Divine Moments (Fa)", Price: 1.66},
		{Code: "540091719838", Name: "Déo Glamorous Moments (Fa)", Price: 1.66},
		{Code: "5410091729745", Name: "Déo Pink Passion (Fa)", Price: 2.99},
		{Code: "5400141437102", Name: "Eponges x6 (Boni)", Price: 2.38},
		{Code: "5400141358858", Name: "Frost Cribb (Boni)", Price: 2.98},
		{Code: "5400141367461", Name: "Haricots blancs Bio (Boni)", Price: 1.45},
		{Code: "5400141058383", Name: "Haricots verts entier", Price: 1.49},
		{Code: "5400141710595", Name: "Lait de coco (Boni)", Price: 0.99},
		{Code: "5400141850574", Name: "Cereal Flakes", Price: 1.72},
		{Code: "5410091734367", Name: "Déo Divine Moments (Fa)", Price: 1.66},
		{Code: "540091719838", Name: "Déo Glamorous Moments (Fa)", Price: 1.66},
		{Code: "5410091729745", Name: "Déo Pink Passion (Fa)", Price: 2.99},
		{Code: "5400141437102", Name: "Eponges x6 (Boni)", Price: 2.38},
		{Code: "5400141358858", Name: "Frost Cribb (Boni)", Price: 2.98},
		{Code: "5400141367461", Name: "Haricots blancs Bio (Boni)", Price: 1.45},
		{Code: "5400141058383", Name: "Haricots verts entier", Price: 1.49},
		{Code: "5400141710595", Name: "Lait de coco (Boni)", Price: 0.99},
		{Code: "5400141850574", Name: "Cereal Flakes", Price: 1.72},
		{Code: "5410091734367", Name: "Déo Divine Moments (Fa)", Price: 1.66},
		{Code: "540091719838", Name: "Déo Glamorous Moments (Fa)", Price: 1.66},
		{Code: "5410091729745", Name: "Déo Pink Passion (Fa)", Price: 2.99},
		{Code: "5400141437102", Name: "Eponges x6 (Boni)", Price: 2.38},
		{Code: "5400141358858", Name: "Frost Cribb (Boni)", Price: 2.98},
		{Code: "5400141367461", Name: "Haricots blancs Bio (Boni)", Price: 1.45},
		{Code: "5400141058383", Name: "Haricots verts entier", Price: 1.49},
		{Code: "5400141710595", Name: "Lait de coco (Boni)", Price: 0.99},
		{Code: "5400141850574", Name: "Cereal Flakes", Price: 1.72},
		{Code: "5410091734367", Name: "Déo Divine Moments (Fa)", Price: 1.66},
		{Code: "540091719838", Name: "Déo Glamorous Moments (Fa)", Price: 1.66},
		{Code: "5410091729745", Name: "Déo Pink Passion (Fa)", Price: 2.99},
		{Code: "5400141437102", Name: "Eponges x6 (Boni)", Price: 2.38},
		{Code: "5400141358858", Name: "Frost Cribb (Boni)", Price: 2.98},
		{Code: "5400141367461", Name: "Haricots blancs Bio (Boni)", Price: 1.45},
		{Code: "5400141058383", Name: "Haricots verts entier", Price: 1.49},
		{Code: "5400141710595", Name: "Lait de coco (Boni)", Price: 0.99},
		{Code: "5400141850574", Name: "Cereal Flakes", Price: 1.72},
		{Code: "5410091734367", Name: "Déo Divine Moments (Fa)", Price: 1.66},
		{Code: "540091719838", Name: "Déo Glamorous Moments (Fa)", Price: 1.66},
		{Code: "5410091729745", Name: "Déo Pink Passion (Fa)", Price: 2.99},
		{Code: "5400141437102", Name: "Eponges x6 (Boni)", Price: 2.38},
		{Code: "5400141358858", Name: "Frost Cribb (Boni)", Price: 2.98},
		{Code: "5400141367461", Name: "Haricots blancs Bio (Boni)", Price: 1.45},
		{Code: "5400141058383", Name: "Haricots verts entier", Price: 1.49},
		{Code: "5400141710595", Name: "Lait de coco (Boni)", Price: 0.99},
		{Code: "5400141850574", Name: "Cereal Flakes", Price: 1.72},
		{Code: "5410091734367", Name: "Déo Divine Moments (Fa)", Price: 1.66},
		{Code: "540091719838", Name: "Déo Glamorous Moments (Fa)", Price: 1.66},
		{Code: "5410091729745", Name: "Déo Pink Passion (Fa)", Price: 2.99},
		{Code: "5400141437102", Name: "Eponges x6 (Boni)", Price: 2.38},
		{Code: "5400141358858", Name: "Frost Cribb (Boni)", Price: 2.98},
		{Code: "5400141367461", Name: "Haricots blancs Bio (Boni)", Price: 1.45},
		{Code: "5400141058383", Name: "Haricots verts entier", Price: 1.49},
		{Code: "5400141710595", Name: "Lait de coco (Boni)", Price: 0.99},
	}

	// Generate PDF
	pdf := GenerateLabels(products, document)
	if pdf == nil {
		t.Errorf("PDF is nil")
	}

	// Output PDF to a file
	date := fmt.Sprintf("%d-%02d-%02d", time.Now().Year(), time.Now().Month(), time.Now().Day())
	err := pdf.OutputFileAndClose(date + "-" + "products-labels-" + document.LabelFormat.Name + ".pdf")
	if err != nil {
		t.Errorf("Error saving PDF: %v", err)
	}

	// Check if the PDF file exists
	if _, err := os.Stat(date + "-" + "products-labels-" + document.LabelFormat.Name + ".pdf"); os.IsNotExist(err) {
		t.Errorf("PDF file not found")
	}

	// Remove the PDF file
	os.Remove(date + "-" + "products-labels-" + document.LabelFormat.Name + ".pdf")

}
