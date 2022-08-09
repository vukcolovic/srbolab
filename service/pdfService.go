package service

import (
	"github.com/go-pdf/fpdf"
	"srbolabApp/model"
	"strconv"
	"time"
)

var (
	PdfService pdfServiceInterface = &pdfService{}
)

type pdfService struct {
}

type pdfServiceInterface interface {
	CreateCertificate(cert *model.Certificate) ([]byte, error)
}

func (pdfService) CreateCertificate(cert *model.Certificate) ([]byte, error) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(15, 20, 15)
	pdf.AddPage()

	var cH float64 = 4

	pdf.SetFont("Arial", "B", 8)
	pdf.Image("img.png", 10, 20, 100, 20, false, "png", 0, "")
	pdf.CellFormat(100, 20, "", "1", 0, "L", false, 0, "")
	pdf.Rect(115, 20, 80, 20, "")
	lines := []string{"SRBOLAB doo Feketic", "Turijski put 17, 21480 Srbobran", "Ogranak, Kontrolno telo za", "kontrolisanje kvaliteta i kvantiteta roba", "telefon: + 381 (21) 310-1533", "mail: vozila@srolab.com"}
	for i, line := range lines {
		y := 24 + i*3
		pdf.Text(118, float64(y), line)
	}
	pdf.Ln(25)

	pdf.SetFont("Arial", "", 7)
	pdf.CellFormat(60, cH, "Referentni broj:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(40, cH, strconv.Itoa(cert.Id), "1", 0, "L", false, 0, "")
	pdf.Ln(-1)
	currentTime := time.Now()
	pdf.CellFormat(60, cH, "Datum izdavanja:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(40, cH, currentTime.Format("02.01.2006"), "1", 0, "L", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(60, cH, "Identifikaciona oznaka vozila (VIN):", "1", 0, "L", false, 0, "")
	pdf.CellFormat(40, cH, "", "1", 0, "L", false, 0, "")

	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 7)
	pdf.Text(60, pdf.GetY(), "IZVOD IZ BAZE O TEHNICKIM KARAKTERISTIKAMA VOZILA")
	pdf.Ln(-1)
	pdf.SetFont("Arial", "", 7)
	pdf.CellFormat(15, 7, "COC", "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 7, "Vozacka dozvola", "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 7, "Tehnicke karakteristike vozila:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(60, 7, "", "1", 0, "L", false, 0, "")

	firstColTxt := []string{"0.1", "0.2", "0.2.1", "-", "16.1", "13", "0.4", "3B", "1", "5", "6", "7", "35", "21", "25", "27", "26", "26.1", "42", "43", "43.1", "44", "49", "16.1"}
	secondColTxt := []string{"(D.1)", "(D.2)", "(D.3)", "(B1)", "(F.1)", "(G)", "(L)", "(J.1)", "(L)", "(5)", "(6)", "(7)", "(35)", "(P)", "(P.1)", "(P.2)", "(P.3)", "(Q)", "(S.1)", "(S.2)", "(43.1)", "(T)", "(V.7)", "(N)"}
	thirdColTxt := []string{"Marka:", "Tip/varijanta/verzija:", "Komercijalna oznaka:", "Procenjena godina proizvodnje:", "Najveca dozvoljena masa vozila (kg):", "Masa vozila spremnog za voznju (kg):", "Kategorija vozila:", "Oznaka oblika za karoseriju:", "Broj osovina i tockova:", "Duzina vozila: (mm)", "Sirina vozila (mm):", "Visina vozila (mm):", "Pneumatik/naplatak kombinacija:", "Oznaka motora:", "Radna zapremina motora (cm3):", "Najveca neto snaga motora(kW):", "Pogonsko gorivo", "Najveca neto snaga/masa vozila (samo za motocikle) (kW/kg):", "Broj mesta za sedenje:", "Broj mesta za stajanje:", "Uredjaj za spajanje vucnog i prikljucnog vozila:", "Najveca brzina (za vozila vrste L)(km/h):", "Nivo izduvne emisije (g/km):", "Najvece dozvoljeno osovinsko opterecenje(kg):"}
	forthColTxt := []string{cert.Brand, cert.TypeVehicle + "/" + cert.Variant + "/" + cert.VersionVehicle, cert.CommercialName, strconv.Itoa(cert.EstimatedProductionYear), strconv.Itoa(cert.MaxMass), strconv.Itoa(cert.RunningMass), cert.Category, cert.BodyworkCode, cert.AxlesTyresNum, strconv.Itoa(cert.Length), strconv.Itoa(cert.Width), strconv.Itoa(cert.Height), cert.TyreWheel, cert.EngineCode, strconv.Itoa(cert.EngineCapacity), strconv.Itoa(cert.EnginePower), cert.Fuel, cert.PowerWeightRatio, strconv.Itoa(cert.SeatNumber), strconv.Itoa(cert.StandingNumber), strconv.Itoa(cert.MaxSpeed), cert.GasLevel, cert.MaxLadenMassAxios, cert.NumberWvta, cert.PollutionCert, cert.NoiseCert, cert.CouplingDeviceApproval}

	var w1 float64 = 15
	var w2 float64 = 25
	var w3 float64 = 80
	var w4 float64 = 60
	for i := 0; i < 24; i++ {
		pdf.Ln(-1)
		pdf.CellFormat(w1, cH, firstColTxt[i], "1", 0, "C", false, 0, "")
		pdf.CellFormat(w2, cH, secondColTxt[i], "1", 0, "C", false, 0, "")
		pdf.CellFormat(w3, cH, thirdColTxt[i], "1", 0, "L", false, 0, "")
		pdf.CellFormat(w4, cH, forthColTxt[i], "1", 0, "L", false, 0, "")
	}

	bottomFirstColTxt := []string{"Broj sertifikata homologacije tipa vozila:", "Broj sertifikata za izduvnu emisiju (datum):", "Broj sertifikata za obuku (datum):"}
	bottomSecondColTxt := []string{cert.NumberWvta, cert.PollutionCert, cert.NoiseCert}
	for i := 0; i < 3; i++ {
		pdf.Ln(-1)
		pdf.CellFormat(80, cH, bottomFirstColTxt[i], "1", 0, "L", false, 0, "")
		pdf.CellFormat(100, cH, bottomSecondColTxt[i], "1", 0, "L", false, 0, "")
	}

	pdf.OutputFileAndClose("xxx.pdf")
	//tempFile, err := ioutil.ReadFile("/home/wolf/GolandProjects/srbolabApp/xxx.pdf")
	//sEnc := base64.StdEncoding.EncodeToString(tempFile)
	return []byte{}, nil
}
