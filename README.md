# Pittacium

Pittacium is a Go library for generating PDF documents with labels, primarily
designed for product labeling purposes.

## Overview

The Pittacium library provides functionality to create PDF documents containing
labels for products. These labels include product information such as name,
price, and barcode.

## Features

* **Label Generation**: Generate labels for products with customizable formats.
* **Barcode Support**: Automatically generate and include barcodes for each product.
* **PDF Output**: Produce PDF documents with labeled products ready for printing.

## Dependencies

The Pittacium library relies on the following third-party dependencies:

* **[fpdf](https://github.com/go-pdf/fpdf)**: A library for generating PDF documents.
* **[barcode](https://github.com/boombuler/barcode)**: A library for generating barcodes.
* **[resize](https://github.com/nfnt/resize)**: A library for resizing images.

## Installation

To use Pittacium in your Go project, simply import it:

```go
import "github.com/alexandrelamberty/pittacium"
```

Then, run `go get to install the package:

```bash
go get github.com/alexandrelamberty/pittacium
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/your-username/pittacium"
)

func main() {
    // Define products
    products := []pittacium.Product{
        {Code: "1234567890123", Name: "Product 1", Price: 10.99},
        {Code: "9876543210987", Name: "Product 2", Price: 15.99},
        // Add more products as needed
    }

    // Define document settings
    document := pittacium.Document{
        Title:       "Product Labels",
        MarginTop:   10.0,
        MarginRight: 10.0,
        MarginLeft:  10.0,
        MarginBottom: 10.0,
        ShowPage:    true,
        PageSize:    "A4",
        Orientation: "P",
        LabelFormat: pittacium.Label2x1,
    }

    // Generate labels and create PDF document
    pdf := pittacium.GenerateLabels(products, document)

    // Save the PDF to a file
    err := pdf.OutputFile("product_labels.pdf")
    if err != nil {
        fmt.Println("Error saving PDF:", err)
        return
    }

    fmt.Println("PDF generated successfully!")
}
```

## Label Formats

Pittacium supports several common label formats, including:

* **2x1:**
* **3x1:**
* **4x2:**
* **2.25x1.25:**
* **3x2:**
