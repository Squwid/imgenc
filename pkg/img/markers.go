package img

type Marker byte

type MarkerData struct {
	Name        string `json:"Name"`
	Description string `json:"Description,omitempty"`
	Marker      string `json:"Marker"`
}

// List of Jpeg Markers https://www.disktuna.com/list-of-jpeg-markers/
const (
	MarkerSOF0  Marker = 0xC0
	MarkerSOF1  Marker = 0xC1
	MarkerSOF2  Marker = 0xC2
	MarkerSOF3  Marker = 0xC3
	MarkerDHT   Marker = 0xC4
	MarkerSOF5  Marker = 0xC5
	MarkerSOF6  Marker = 0xC6
	MarkerSOF7  Marker = 0xC7
	MarkerJPG   Marker = 0xC8
	MarkerSOF9  Marker = 0xC9
	MarkerSOF10 Marker = 0xCA
	MarkerSOF11 Marker = 0xCB
	MarkerDAC   Marker = 0xCC
	MarkerSOF13 Marker = 0xCD
	MarkerSOF14 Marker = 0xCE
	MarkerSOF15 Marker = 0xCF
	MarkerRST0  Marker = 0xD0
	MarkerRST1  Marker = 0xD1
	MarkerRST2  Marker = 0xD2
	MarkerRST3  Marker = 0xD3
	MarkerRST4  Marker = 0xD4
	MarkerRST5  Marker = 0xD5
	MarkerRST6  Marker = 0xD6
	MarkerRST7  Marker = 0xD7
	MarkerSOI   Marker = 0xD8
	MarkerEOI   Marker = 0xD9
	MarkerSOS   Marker = 0xDA
	MarkerDQT   Marker = 0xDB
	MarkerDNL   Marker = 0xDC
	MarkerDRI   Marker = 0xDD
	MarkerDHP   Marker = 0xDE
	MarkerEXP   Marker = 0xDF
	MarkerAPP0  Marker = 0xE0
	MarkerAPP1  Marker = 0xE1
	MarkerAPP2  Marker = 0xE2
	MarkerAPP3  Marker = 0xE3
	MarkerAPP4  Marker = 0xE4
	MarkerAPP5  Marker = 0xE5
	MarkerAPP6  Marker = 0xE6
	MarkerAPP7  Marker = 0xE7
	MarkerAPP8  Marker = 0xE8
	MarkerAPP9  Marker = 0xE9
	MarkerAPP10 Marker = 0xEA
	MarkerAPP11 Marker = 0xEB
	MarkerAPP12 Marker = 0xEC
	MarkerAPP13 Marker = 0xED
	MarkerAPP14 Marker = 0xEE
	MarkerAPP15 Marker = 0xEF
	MarkerJPG0  Marker = 0xF0
	MarkerJPG1  Marker = 0xF1
	MarkerJPG2  Marker = 0xF2
	MarkerJPG3  Marker = 0xF3
	MarkerJPG4  Marker = 0xF4
	MarkerJPG5  Marker = 0xF5
	MarkerJPG6  Marker = 0xF6
	MarkerJPG7  Marker = 0xF7
	MarkerJPG8  Marker = 0xF8
	MarkerJPG9  Marker = 0xF9
	MarkerJPG10 Marker = 0xFA
	MarkerJPG11 Marker = 0xFB
	MarkerJPG12 Marker = 0xFC
	MarkerJPG13 Marker = 0xFD
	MarkerCOM   Marker = 0xFE
)

var Markers = map[Marker]MarkerData{
	MarkerSOF0: {
		Marker:      "S0F0",
		Name:        "Start of Frame 0",
		Description: "Baseline DCT",
	},
	MarkerSOF1: {
		Marker:      "SOF1",
		Name:        "Start of Frame 1",
		Description: "Extended Sequential DCT",
	},
	MarkerSOF2: {
		Marker:      "SOF2",
		Name:        "Start of Frame 2",
		Description: "Progressive DCT",
	},
	MarkerSOF3: {
		Marker:      "SOF3",
		Name:        "Start of Frame 3",
		Description: "Lossless (sequential)",
	},
	MarkerDHT: {
		Marker: "DHT",
		Name:   "Define Huffman Table",
	},
	MarkerSOF5: {
		Marker:      "SOF5",
		Name:        "Start of Frame 5",
		Description: "Differential sequential DCT",
	},
	MarkerSOF6: {
		Marker:      "SOF6",
		Name:        "Start of Frame 6",
		Description: "Differential progressive DCT",
	},
	MarkerSOF7: {
		Marker:      "SOF7",
		Name:        "Start of Frame 7",
		Description: "Differential lossless (sequential)",
	},
	MarkerJPG: {
		Marker: "JPG",
		Name:   "JPEG Extensions",
	},
	MarkerSOF9: {
		Marker:      "SOF9",
		Name:        "Start of Frame 9",
		Description: "Extended sequential DCT, Arithmetic coding",
	},
	MarkerSOF10: {
		Marker:      "SOF10",
		Name:        "Start of Frame 10",
		Description: "Progressive DCT, Arithmetic coding",
	},
	MarkerSOF11: {
		Marker:      "SOF11",
		Name:        "Start of Frame 11",
		Description: "Lossless (sequential), Arithmetic coding",
	},
	MarkerDAC: {
		Marker: "DAC",
		Name:   "Define Arithmetic Coding",
	},
	MarkerSOF13: {
		Marker:      "SOF13",
		Name:        "Start of Frame 13",
		Description: "Differential sequential DCT, Arithmetic coding",
	},
	MarkerSOF14: {
		Marker:      "SOF14",
		Name:        "Start of Frame 14",
		Description: "Differential progressive DCT, Arithmetic coding",
	},
	MarkerSOF15: {
		Marker:      "SOF15",
		Name:        "Start of Frame 15",
		Description: "Differential lossless (sequential), Arithmetic coding",
	},
	MarkerRST0: {
		Marker: "RST0",
		Name:   "Restart Marker 0",
	},
	MarkerRST1: {
		Marker: "RST1",
		Name:   "Restart Marker 1",
	},
	MarkerRST2: {
		Marker: "RST2",
		Name:   "Restart Marker 2",
	},
	MarkerRST3: {
		Marker: "RST3",
		Name:   "Restart Marker 3",
	},
	MarkerRST4: {
		Marker: "RST4",
		Name:   "Restart Marker 4",
	},
	MarkerRST5: {
		Marker: "RST5",
		Name:   "Restart Marker 5",
	},
	MarkerRST6: {
		Marker: "RST6",
		Name:   "Restart Marker 6",
	},
	MarkerRST7: {
		Marker: "RST7",
		Name:   "Restart Marker 7",
	},
	MarkerSOI: {
		Marker: "SOI",
		Name:   "Start of Image",
	},
	MarkerEOI: {
		Marker: "EOI",
		Name:   "End of Image",
	},
	MarkerSOS: {
		Marker: "SOS",
		Name:   "Start of Scan",
	},
	MarkerDQT: {
		Marker: "DQT",
		Name:   "Define Quantization Table",
	},
	MarkerDNL: {
		Marker:      "DNL",
		Name:        "Define Number of Lines",
		Description: "(Not common)",
	},
	MarkerDRI: {
		Marker: "DRI",
		Name:   "Define Restart Interval",
	},
	MarkerDHP: {
		Marker:      "DHP",
		Name:        "Define Hierarchical Progression",
		Description: "(Not common)",
	},
	MarkerEXP: {
		Marker:      "EXP",
		Name:        "Expand Reference Component",
		Description: "(Not common)",
	},
	MarkerAPP0: {
		Marker:      "APP0",
		Name:        "Application Segment 0",
		Description: "JFIF – JFIF JPEG image\nAVI1 – Motion JPEG (MJPG)",
	},
	MarkerAPP1: {
		Marker:      "APP1",
		Name:        "Application Segment 1",
		Description: "EXIF Metadata, TIFF IFD format,\nJPEG Thumbnail (160×120)\nAdobe XMP",
	},
	MarkerAPP2: {
		Marker:      "APP2",
		Name:        "Application Segment 2",
		Description: "ICC color profile,\nFlashPix",
	},
	MarkerAPP3: {
		Marker:      "APP3",
		Name:        "Application Segment 3",
		Description: "(Not common)\nJPS Tag for Stereoscopic JPEG images",
	},
	MarkerAPP4: {
		Marker:      "APP4",
		Name:        "Application Segment 4",
		Description: "(Not common)",
	},
	MarkerAPP5: {
		Marker:      "APP5",
		Name:        "Application Segment 5",
		Description: "(Not common)",
	},
	MarkerAPP6: {
		Marker:      "APP6",
		Name:        "Application Segment 6",
		Description: "(Not common)\nNITF Lossles profile",
	},
	MarkerAPP7: {
		Marker:      "APP7",
		Name:        "Application Segment 7",
		Description: "(Not common)",
	},
	MarkerAPP8: {
		Marker:      "APP8",
		Name:        "Application Segment 8",
		Description: "(Not common)",
	},
	MarkerAPP9: {
		Marker:      "APP9",
		Name:        "Application Segment 9",
		Description: "(Not common)",
	},
	MarkerAPP10: {
		Marker:      "APP10",
		Name:        "Application Segment 10\nPhoTags",
		Description: "(Not common)\nActiveObject (multimedia messages / captions)",
	},
	MarkerAPP11: {
		Marker:      "APP11",
		Name:        "Application Segment 11",
		Description: "(Not common)\nHELIOS JPEG Resources (OPI Postscript)",
	},
	MarkerAPP12: {
		Marker:      "APP12",
		Name:        "Application Segment 12",
		Description: "Picture Info (older digicams),\nPhotoshop Save for Web: Ducky",
	},
	MarkerAPP13: {
		Marker:      "APP13",
		Name:        "Application Segment 13",
		Description: "Photoshop Save As: IRB, 8BIM, IPTC",
	},
	MarkerAPP14: {
		Marker:      "APP14",
		Name:        "Application Segment 14",
		Description: "(Not common)",
	},
	MarkerAPP15: {
		Marker:      "APP15",
		Name:        "Application Segment 15",
		Description: "(Not common)",
	},
	MarkerJPG0: {
		Marker:      "JPG0",
		Name:        "JPEG Extension 0",
		Description: "(Not common)",
	},
	MarkerJPG1: {
		Marker:      "JPG1",
		Name:        "JPEG Extension 1",
		Description: "(Not common)",
	},
	MarkerJPG2: {
		Marker:      "JPG2",
		Name:        "JPEG Extension 2",
		Description: "(Not common)",
	},
	MarkerJPG3: {
		Marker:      "JPG3",
		Name:        "JPEG Extension 3",
		Description: "(Not common)",
	},
	MarkerJPG4: {
		Marker:      "JPG4",
		Name:        "JPEG Extension 4",
		Description: "(Not common)",
	},
	MarkerJPG5: {
		Marker:      "JPG5",
		Name:        "JPEG Extension 5",
		Description: "(Not common)",
	},
	MarkerJPG6: {
		Marker:      "JPG6",
		Name:        "JPEG Extension 6",
		Description: "(Not common)",
	},
	MarkerJPG7: {
		Marker:      "JPG7\nSOF48",
		Name:        "JPEG Extension 7\nJPEG-LS",
		Description: "Lossless JPEG",
	},
	MarkerJPG8: {
		Marker:      "JPG8\nLSE",
		Name:        "JPEG Extension 8\nJPEG-LS Extension",
		Description: "Lossless JPEG Extension Parameters",
	},
	MarkerJPG9: {
		Marker:      "JPG9",
		Name:        "JPEG Extension 9",
		Description: "(Not common)",
	},
	MarkerJPG10: {
		Marker:      "JPG10",
		Name:        "JPEG Extension 10",
		Description: "(Not common)",
	},
	MarkerJPG11: {
		Marker:      "JPG11",
		Name:        "JPEG Extension 11",
		Description: "(Not common)",
	},
	MarkerJPG12: {
		Marker:      "JPG12",
		Name:        "JPEG Extension 12",
		Description: "(Not common)",
	},
	MarkerJPG13: {
		Marker:      "JPG13",
		Name:        "JPEG Extension 13",
		Description: "(Not common)",
	},
	MarkerCOM: {
		Marker: "COM",
		Name:   "Comment",
	},
}
