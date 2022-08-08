package service

import (
	"github.com/go-pdf/fpdf"
	"srbolabApp/model"
)

var (
	PdfService pdfServiceInterface = &pdfService{}
)

type pdfService struct {
}

type pdfServiceInterface interface {
	CreateCertificate(cert model.Certificate) ([]byte, error)
}

func (pdfService) CreateCertificate(cert model.Certificate) ([]byte, error) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
	return []byte{}, err
}
