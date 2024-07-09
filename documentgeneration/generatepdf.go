package documentgeneration

import (
	"html/template"
	"os"
	"strconv"
	// Set the font location
	"strings"

	"github.com/jung-kurt/gofpdf"
	// "unicode/utf8"
)

// Company represents information about the company.
type Company struct {
	Name         string
	Office       string
	Objectives   []string
	Liability    string
	ShareCapital string
}

// Subscriber represents information about a subscriber.
type Subscriber struct {
	Name       string
	Occupation string
	Shares     string
	Signature  string
}

type Date struct {
	Day   string
	Month string
	Year  string
}

// PageData represents the data passed to the HTML and PDF templates.
type PageData struct {
	// Title       string
	// News        string
	// Body        string
	CompanyName string
	Company     Company
	Subscribers []Subscriber
	Date        Date
	CheckOption CheckOption
}

type CheckOption struct {
	AdoptTableII                 bool `json:"adopt_table_a_part_II"`
	AdoptTableIIWithModification bool `json:"adopt_table_a_part_II_with_modification"`
}

func GeneratePDF(p PageData, filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	// pdf.UnicodeTranslatorFromDescriptor("UTF-8")
    pdf.SetFontLocation("./font")

	// Register the DejaVuSans font
	pdf.AddUTF8Font("DejaVuSans", "", "../ttf/DejaVuSans.ttf")
	pdf.AddUTF8Font("DejaVuSans-Bold", "", "../ttf/DejaVuSans-Bold.ttf")


	pdf.Ln(31)

	// Centered title
	// pdf.SetFont("Arial", "B", 25)
	// pdf.CellFormat(0, 10, "URSB", "", 1, "C", false, 0, "")

	lf, tp, _, _ := pdf.GetMargins()
	newLeftMargin := lf + 80
	newTopMargin := tp + 24
	//adding an image
	imagePath := "./images/ursblogo.jpg"
	pdf.ImageOptions(imagePath, newLeftMargin, newTopMargin, 34, 0, false, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "")

	pdf.Ln(9)

	lineHeight := 5.5

	//setting the font
	pdf.SetFont("Arial", "B", 10)

	upperCasedName := strings.ToUpper(p.Company.Name)

	// Text content
	lines := []string{
		"THE REPUBLIC OF UGANDA",
		"THE COMPANIES ACT NO. 1 OF 2012",
		"MEMORANDUM AND ARTICLES OF ASSOCIATION OF A COMPANY LIMITED BY",
		"SHARES",
		"MEMORANDUM AND ARTICLES OF " + upperCasedName,
	}

	// Centered underlined content
	for _, line := range lines {
		// Get current X and Y
		// x := pdf.GetX()
		y := pdf.GetY()

		// Print the text
		pdf.CellFormat(0, lineHeight, line, "", 1, "C", false, 0, "")

		// Get text width
		width := pdf.GetStringWidth(line)

		// Calculate starting X position for the underline
		startX := (210.0 - width) / 2 // Assuming A4 page width (210mm)

		// Draw underline
		pdf.Line(startX, y+lineHeight, startX+width, y+lineHeight)
	}

	pdf.Ln(6)

	// Ordered list
	pdf.SetFont("Arial", "", 10)
	// \""+upperCasedName,"\"",
	pdf.CellFormat(0, 7, "          1. The name of the company is \""+upperCasedName+"\"", "", 1, "", false, 0, "")
	pdf.CellFormat(0, 7, "          2. The registered office of the company will be situated in "+p.Company.Office, "", 1, "", false, 0, "")
	pdf.CellFormat(0, 7, "          3. The objects for which the company is established are:\"", "", 1, "", false, 0, "")

	//increasing the left margin
	leftMargin := 30.0
	for i, obj := range p.Company.Objectives {
		pdf.SetX(leftMargin)
		pdf.CellFormat(0, 7, "  "+string('1'+i)+". "+obj, "", 1, "", false, 0, "")
		// pdf.CellFormat(0, 7, "    -"+obj, "", 1, "", false, 0, "")
	}
	// fmt.Printf(obj)

	// pdf.CellFormat(0, 7, "          4. Liability Details "+p.Company.Liability, "", 1, "", false, 0, "")
	pdf.CellFormat(0, 7, "          4. The liability of the members is "+p.Company.Liability+".", "", 1, "", false, 0, "")
	pdf.CellFormat(0, 7, "          5. The share capital of the company is "+p.Company.ShareCapital+" divided into 100 Ordinary shares of 10,000 UGX each.", "", 1, "", false, 0, "")
	pdf.Ln(1)

	// Paragraph
	pdf.MultiCell(0, 4, "WE the several persons whose names are subscribed, desire to be formed into a company under this memorandum of association, and we respectively agree to take the number of shares in the capital of the company set opposite our respective names.", "", "L", false)
	pdf.Ln(5)

	// Adjust cell widths to fit within page width
	cellWidths := []float64{80, 55, 55}
	totalWidth := 0.0
	for _, width := range cellWidths {
		totalWidth += width
	}

	// Print table headers
	pdf.SetFont("Arial", "B", 10)
	cellWidth1 := 80.0
	cellWidth2 := 55.0
	// cellWidth3 := 135.0

	// Initial position
	x := pdf.GetX()
	y := pdf.GetY()

	// First cell
	pdf.SetFont("Arial", "", 9)
	pdf.MultiCell(cellWidth1, 4, "Name of postal addresses and occupations of subscribers", "1", "L", false)

	// Second cell (move to the right of the first cell)
	pdf.SetXY(x+cellWidth1, y) // Move to the right of the previous cell
	pdf.MultiCell(cellWidth2, 4, "Number of shares taken by each subscriber", "1", "L", false)

	// Third cell (move to the right of the second cell)
	pdf.SetXY(x+cellWidth1+cellWidth2, y) // Move to the right of the previous cell
	pdf.MultiCell(cellWidth2, 8, "Signature of subscribers", "1", "L", false)

	// Move to the next line
	pdf.Ln(0) // Adjust the spacing as needed

	// Table headers
	// pdf.SetFont("Arial", "B", 12)
	// pdf.MultiCell(80, 10, "Name of postal addresses and occupations of subscribers", "1", "C", false)
	// pdf.MultiCell(80, 10, "Number of shares taken by each subscriber", "1", "C", false)
	// pdf.MultiCell(80, 10, "Signature of subscribers", "1", "C", false)
	// pdf.CellFormat(cellWidths[0], 10, "Name of postal addresses and occupations of subscribers", "1", 0, "C", false, 0, "")
	// pdf.CellFormat(cellWidths[1], 10, Number of shares taken" by each subscriber", "1", 0, "C", false, 0, "")
	// pdf.CellFormat(cellWidths[2], 10, "Signature of subscribers", "1", 1, "C", false, 0, "")

	// Table rows
	pdf.SetFont("Arial", "", 9)
	totalShares := 0
	for _, subscriber := range p.Subscribers {
		pdf.MultiCell(cellWidths[0], 5, subscriber.Name+"\n"+subscriber.Occupation, "1", "L", false)
		// Get the current position after MultiCell
		x := pdf.GetX()
		y := pdf.GetY()

		// Second cell
		pdf.SetXY(x+cellWidths[0], y-10) // Adjust Y position after MultiCell
		pdf.MultiCell(cellWidths[1], 5, subscriber.Shares+"\n"+"TOTAL   "+subscriber.Shares, "1", "L", false)

		// Third cell
		pdf.SetXY(x+cellWidths[0]+cellWidths[1], y-10) // Adjust Y position after MultiCell
		pdf.CellFormat(cellWidths[2], 10, subscriber.Signature, "1", 1, "L", false, 0, "")

		// pdf.CellFormat(cellWidths[1], 10, subscriber.Shares, "1", 0, "L", false, 0, "")
		// pdf.CellFormat(cellWidths[2], 10, "", "1", 1, "L", false, 0, "")

		shares, err := strconv.Atoi(subscriber.Shares)
		if err != nil {
			err.Error()
			// panic(err.Error())
		}
		totalShares += shares
	}

	// total of shares
	// pdf.SetFont("Arial", "", 10)
	// pdf.MultiCell(cellWidths[0], 4, "Total", "1", "L", false)
	// pdf.MultiCell(cellWidths[1], 4, strconv.Itoa(totalShares), "1", "L", false)
	// pdf.MultiCell(cellWidths[2], 4, "", "1", "L", false)




	
	// Display the total at the bottom
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(cellWidths[0], 5, "Total Shares taken", "1", 0, "L", false, 0, "")
	pdf.CellFormat(cellWidths[1], 5, strconv.Itoa(totalShares), "1", 0, "L", false, 0, "")
	pdf.CellFormat(cellWidths[2], 5, "", "1", 1, "L", false, 0, "")

	pdf.Ln(3)

	// Footer
	pdf.SetFont("Arial", "", 10)
	// pdf.MultiCell(0, 6, "Dated the \"21st\" day of \"June\", \"2024\"\nWitness to the above signatures ................................\n", "", "L", false)
	// pdf.MultiCell(0, 6, "Dated the \"21st\" day of \"June\", \"2024\"\nWitness to the above signatures ................................\n", "", "L", false)

	pdf.MultiCell(0, 6, "Dated the \""+p.Date.Day+"\" day of \""+p.Date.Month+"\", \""+p.Date.Year+"\"\nWitness to the above signatures .................................................\n", "", "L", false)
	// }

	// fmt.Println(p.Date)
	pdf.Ln(3)

	// Additional content
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(0, 4, "ARTICLES OF ASSOCIATION OF "+upperCasedName+".", "", 1, "L", false, 0, "")
	pdf.Ln(3)

	

	// Initialize CheckOption with default values
	// checkOpts := CheckOption{
	// 	AdoptTableII:                 true, // Set to true for adoption
	// 	// AdoptTableIIWithModification: false,
	// }


	// Use Unicode checkbox characters
	   checkboxChecked := "\u2611" // ☑
	   checkboxEmpty := "\u2610" // ☐

	pdf.SetFont("DejaVuSans", "", 9)
	if p.CheckOption.AdoptTableII {
		pdf.CellFormat(0, 5, checkboxChecked+"Adopt table A Part 11 of Companies Act 2012", "", 1, "", false, 0, "")
	 } else {
		pdf.CellFormat(0, 5, checkboxEmpty+"Adopt table A Part 11 of Companies Act 2012", "", 1, "", false, 0, "")
	}
	
	if p.CheckOption.AdoptTableIIWithModification {
		pdf.CellFormat(0, 5, checkboxChecked+" Adopt table A Part II of Companies Act 2012 with modification (attach the modification)", "", 1, "", false, 0, "")
	} else {
		pdf.CellFormat(0, 5, checkboxEmpty+" Adopt table A Part II of Companies Act 2012 with modification (attach the modification)", "", 1, "", false, 0, "")
	}
	// fmt.Println(checkboxChecked)
	// fmt.Println(checkboxEmpty)
	pdf.Ln(5)

	// pdf.SetFont("Arial", "", 10)
	// pdf.CellFormat(0, 6, "[ ] Adopt table A Part 11 of Companies Act 2012", "", 1, "", false, 0, "")
	// pdf.CellFormat(0, 6, "[ ] Adopt table A Part II of Companies Act 2012 with modification (attach the modification)", "", 1, "", false, 0, "")
	// pdf.Ln(5)

	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(0, 4, "WE, the several persons whose names are subscribed, desire to be formed into a company under this association, and we respectively agree to take the number of shares in the capital of the company set opposite our respective names.", "", "L", false)
	pdf.Ln(7)

	// Second table
	pdf.SetFont("Arial", "", 9)
	pdf.CellFormat(120, 4, "Names, postal addresses and occupations of subscribers", "1", 0, "L", false, 0, "")
	pdf.CellFormat(70, 4, "Signature of subscribers", "1", 1, "L", false, 0, "")

	// Table rows
	pdf.SetFont("Arial", "", 10)
	for _, subscriber := range p.Subscribers {
		pdf.MultiCell(120, 5, subscriber.Name+"\n"+subscriber.Occupation, "1", "L", false)
		// pdf.CellFormat(120, 10, subscriber.Name+"\n"+subscriber.Occupation, "1", 0, "L", false, 0, "")
		x := pdf.GetX()
		y := pdf.GetY()
		pdf.SetXY(x+120, y-10)
		pdf.MultiCell(70, 10, "", "1", "L", false)
	}

	pdf.Ln(5)

	pdf.MultiCell(0, 6, "Dated the \""+p.Date.Day+"\" day of \""+p.Date.Month+"\", \""+p.Date.Year+"\"\nWitness to the above signatures .................................................\n\n\nRecommended, not more than 2 pages\n", "", "L", false)

	// pdf.MultiCell(0, 7, "Dated the \"21st\" day of \"June\", \"2024\"\nWitness to the above signatures ................................\n\n\nRecommended, not more than 2 pages\n", "", "L", false)
	// pdf.Ln(5)

	// pdf.SetFont("Arial", "B", 12)
	// pdf.CellFormat(0, 7, "REGISTRAR OF COMPANIES", "", 1, "C", false, 0, "")

	// Output PDF
	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		return err
	}

	return nil
}


func GenerateHTML(p PageData, filename string) error {
	// Parse template
	htmlTemplate, err := template.ParseFiles("basictemplating.html")
	if err != nil {
		return err
	}

	// Create HTML file
	htmlFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer htmlFile.Close()

	// Execute template
	err = htmlTemplate.Execute(htmlFile, p)
	if err != nil {
		return err
	}

	return nil
}
