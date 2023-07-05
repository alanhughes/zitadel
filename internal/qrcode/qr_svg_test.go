//go:build unit

package qrcode

import (
	"bytes"
	"os"
	"testing"

	"github.com/ajstarks/svgo"
	"github.com/boombuler/barcode/qr"
)

func Test_goqrsvg(t *testing.T) {
	buf := bytes.NewBufferString("")
	s := svg.New(buf)

	// Create the barcode
	qrCode, _ := qr.Encode("Hello World", qr.M, qr.Auto)

	// Write QR code to SVG
	qs := NewQrSVG(qrCode, 5)
	qs.StartQrSVG(s)
	qs.WriteQrSVG(s)

	s.End()

	// Check if output the same as correctOutput
	if buf.String() != correctOutput {
		t.Error("Something is not right... The SVG created is not the same as correctOutput.")
	}
}

func ExampleNewQrSVG() {
	s := svg.New(os.Stdout)

	// Create the barcode
	qrCode, _ := qr.Encode("Hello World", qr.M, qr.Auto)

	// Write QR code to SVG
	qs := NewQrSVG(qrCode, 5)
	qs.StartQrSVG(s)
	qs.WriteQrSVG(s)

	s.End()
}

const correctOutput = `<?xml version="1.0"?>
<!-- Generated by SVGo -->
<svg width="145" height="145"
     xmlns="http://www.w3.org/2000/svg"
     xmlns:xlink="http://www.w3.org/1999/xlink">
<rect x="20" y="20" width="5" height="5" class="color" />
<rect x="25" y="20" width="5" height="5" class="color" />
<rect x="30" y="20" width="5" height="5" class="color" />
<rect x="35" y="20" width="5" height="5" class="color" />
<rect x="40" y="20" width="5" height="5" class="color" />
<rect x="45" y="20" width="5" height="5" class="color" />
<rect x="50" y="20" width="5" height="5" class="color" />
<rect x="55" y="20" width="5" height="5" class="bg-color" />
<rect x="60" y="20" width="5" height="5" class="color" />
<rect x="65" y="20" width="5" height="5" class="color" />
<rect x="70" y="20" width="5" height="5" class="bg-color" />
<rect x="75" y="20" width="5" height="5" class="color" />
<rect x="80" y="20" width="5" height="5" class="bg-color" />
<rect x="85" y="20" width="5" height="5" class="bg-color" />
<rect x="90" y="20" width="5" height="5" class="color" />
<rect x="95" y="20" width="5" height="5" class="color" />
<rect x="100" y="20" width="5" height="5" class="color" />
<rect x="105" y="20" width="5" height="5" class="color" />
<rect x="110" y="20" width="5" height="5" class="color" />
<rect x="115" y="20" width="5" height="5" class="color" />
<rect x="120" y="20" width="5" height="5" class="color" />
<rect x="20" y="25" width="5" height="5" class="color" />
<rect x="25" y="25" width="5" height="5" class="bg-color" />
<rect x="30" y="25" width="5" height="5" class="bg-color" />
<rect x="35" y="25" width="5" height="5" class="bg-color" />
<rect x="40" y="25" width="5" height="5" class="bg-color" />
<rect x="45" y="25" width="5" height="5" class="bg-color" />
<rect x="50" y="25" width="5" height="5" class="color" />
<rect x="55" y="25" width="5" height="5" class="bg-color" />
<rect x="60" y="25" width="5" height="5" class="bg-color" />
<rect x="65" y="25" width="5" height="5" class="bg-color" />
<rect x="70" y="25" width="5" height="5" class="color" />
<rect x="75" y="25" width="5" height="5" class="bg-color" />
<rect x="80" y="25" width="5" height="5" class="color" />
<rect x="85" y="25" width="5" height="5" class="bg-color" />
<rect x="90" y="25" width="5" height="5" class="color" />
<rect x="95" y="25" width="5" height="5" class="bg-color" />
<rect x="100" y="25" width="5" height="5" class="bg-color" />
<rect x="105" y="25" width="5" height="5" class="bg-color" />
<rect x="110" y="25" width="5" height="5" class="bg-color" />
<rect x="115" y="25" width="5" height="5" class="bg-color" />
<rect x="120" y="25" width="5" height="5" class="color" />
<rect x="20" y="30" width="5" height="5" class="color" />
<rect x="25" y="30" width="5" height="5" class="bg-color" />
<rect x="30" y="30" width="5" height="5" class="color" />
<rect x="35" y="30" width="5" height="5" class="color" />
<rect x="40" y="30" width="5" height="5" class="color" />
<rect x="45" y="30" width="5" height="5" class="bg-color" />
<rect x="50" y="30" width="5" height="5" class="color" />
<rect x="55" y="30" width="5" height="5" class="bg-color" />
<rect x="60" y="30" width="5" height="5" class="color" />
<rect x="65" y="30" width="5" height="5" class="bg-color" />
<rect x="70" y="30" width="5" height="5" class="color" />
<rect x="75" y="30" width="5" height="5" class="color" />
<rect x="80" y="30" width="5" height="5" class="color" />
<rect x="85" y="30" width="5" height="5" class="bg-color" />
<rect x="90" y="30" width="5" height="5" class="color" />
<rect x="95" y="30" width="5" height="5" class="bg-color" />
<rect x="100" y="30" width="5" height="5" class="color" />
<rect x="105" y="30" width="5" height="5" class="color" />
<rect x="110" y="30" width="5" height="5" class="color" />
<rect x="115" y="30" width="5" height="5" class="bg-color" />
<rect x="120" y="30" width="5" height="5" class="color" />
<rect x="20" y="35" width="5" height="5" class="color" />
<rect x="25" y="35" width="5" height="5" class="bg-color" />
<rect x="30" y="35" width="5" height="5" class="color" />
<rect x="35" y="35" width="5" height="5" class="color" />
<rect x="40" y="35" width="5" height="5" class="color" />
<rect x="45" y="35" width="5" height="5" class="bg-color" />
<rect x="50" y="35" width="5" height="5" class="color" />
<rect x="55" y="35" width="5" height="5" class="bg-color" />
<rect x="60" y="35" width="5" height="5" class="color" />
<rect x="65" y="35" width="5" height="5" class="color" />
<rect x="70" y="35" width="5" height="5" class="color" />
<rect x="75" y="35" width="5" height="5" class="bg-color" />
<rect x="80" y="35" width="5" height="5" class="bg-color" />
<rect x="85" y="35" width="5" height="5" class="bg-color" />
<rect x="90" y="35" width="5" height="5" class="color" />
<rect x="95" y="35" width="5" height="5" class="bg-color" />
<rect x="100" y="35" width="5" height="5" class="color" />
<rect x="105" y="35" width="5" height="5" class="color" />
<rect x="110" y="35" width="5" height="5" class="color" />
<rect x="115" y="35" width="5" height="5" class="bg-color" />
<rect x="120" y="35" width="5" height="5" class="color" />
<rect x="20" y="40" width="5" height="5" class="color" />
<rect x="25" y="40" width="5" height="5" class="bg-color" />
<rect x="30" y="40" width="5" height="5" class="color" />
<rect x="35" y="40" width="5" height="5" class="color" />
<rect x="40" y="40" width="5" height="5" class="color" />
<rect x="45" y="40" width="5" height="5" class="bg-color" />
<rect x="50" y="40" width="5" height="5" class="color" />
<rect x="55" y="40" width="5" height="5" class="bg-color" />
<rect x="60" y="40" width="5" height="5" class="color" />
<rect x="65" y="40" width="5" height="5" class="color" />
<rect x="70" y="40" width="5" height="5" class="color" />
<rect x="75" y="40" width="5" height="5" class="color" />
<rect x="80" y="40" width="5" height="5" class="bg-color" />
<rect x="85" y="40" width="5" height="5" class="bg-color" />
<rect x="90" y="40" width="5" height="5" class="color" />
<rect x="95" y="40" width="5" height="5" class="bg-color" />
<rect x="100" y="40" width="5" height="5" class="color" />
<rect x="105" y="40" width="5" height="5" class="color" />
<rect x="110" y="40" width="5" height="5" class="color" />
<rect x="115" y="40" width="5" height="5" class="bg-color" />
<rect x="120" y="40" width="5" height="5" class="color" />
<rect x="20" y="45" width="5" height="5" class="color" />
<rect x="25" y="45" width="5" height="5" class="bg-color" />
<rect x="30" y="45" width="5" height="5" class="bg-color" />
<rect x="35" y="45" width="5" height="5" class="bg-color" />
<rect x="40" y="45" width="5" height="5" class="bg-color" />
<rect x="45" y="45" width="5" height="5" class="bg-color" />
<rect x="50" y="45" width="5" height="5" class="color" />
<rect x="55" y="45" width="5" height="5" class="bg-color" />
<rect x="60" y="45" width="5" height="5" class="color" />
<rect x="65" y="45" width="5" height="5" class="color" />
<rect x="70" y="45" width="5" height="5" class="bg-color" />
<rect x="75" y="45" width="5" height="5" class="bg-color" />
<rect x="80" y="45" width="5" height="5" class="bg-color" />
<rect x="85" y="45" width="5" height="5" class="bg-color" />
<rect x="90" y="45" width="5" height="5" class="color" />
<rect x="95" y="45" width="5" height="5" class="bg-color" />
<rect x="100" y="45" width="5" height="5" class="bg-color" />
<rect x="105" y="45" width="5" height="5" class="bg-color" />
<rect x="110" y="45" width="5" height="5" class="bg-color" />
<rect x="115" y="45" width="5" height="5" class="bg-color" />
<rect x="120" y="45" width="5" height="5" class="color" />
<rect x="20" y="50" width="5" height="5" class="color" />
<rect x="25" y="50" width="5" height="5" class="color" />
<rect x="30" y="50" width="5" height="5" class="color" />
<rect x="35" y="50" width="5" height="5" class="color" />
<rect x="40" y="50" width="5" height="5" class="color" />
<rect x="45" y="50" width="5" height="5" class="color" />
<rect x="50" y="50" width="5" height="5" class="color" />
<rect x="55" y="50" width="5" height="5" class="bg-color" />
<rect x="60" y="50" width="5" height="5" class="color" />
<rect x="65" y="50" width="5" height="5" class="bg-color" />
<rect x="70" y="50" width="5" height="5" class="color" />
<rect x="75" y="50" width="5" height="5" class="bg-color" />
<rect x="80" y="50" width="5" height="5" class="color" />
<rect x="85" y="50" width="5" height="5" class="bg-color" />
<rect x="90" y="50" width="5" height="5" class="color" />
<rect x="95" y="50" width="5" height="5" class="color" />
<rect x="100" y="50" width="5" height="5" class="color" />
<rect x="105" y="50" width="5" height="5" class="color" />
<rect x="110" y="50" width="5" height="5" class="color" />
<rect x="115" y="50" width="5" height="5" class="color" />
<rect x="120" y="50" width="5" height="5" class="color" />
<rect x="20" y="55" width="5" height="5" class="bg-color" />
<rect x="25" y="55" width="5" height="5" class="bg-color" />
<rect x="30" y="55" width="5" height="5" class="bg-color" />
<rect x="35" y="55" width="5" height="5" class="bg-color" />
<rect x="40" y="55" width="5" height="5" class="bg-color" />
<rect x="45" y="55" width="5" height="5" class="bg-color" />
<rect x="50" y="55" width="5" height="5" class="bg-color" />
<rect x="55" y="55" width="5" height="5" class="bg-color" />
<rect x="60" y="55" width="5" height="5" class="bg-color" />
<rect x="65" y="55" width="5" height="5" class="bg-color" />
<rect x="70" y="55" width="5" height="5" class="color" />
<rect x="75" y="55" width="5" height="5" class="bg-color" />
<rect x="80" y="55" width="5" height="5" class="bg-color" />
<rect x="85" y="55" width="5" height="5" class="bg-color" />
<rect x="90" y="55" width="5" height="5" class="bg-color" />
<rect x="95" y="55" width="5" height="5" class="bg-color" />
<rect x="100" y="55" width="5" height="5" class="bg-color" />
<rect x="105" y="55" width="5" height="5" class="bg-color" />
<rect x="110" y="55" width="5" height="5" class="bg-color" />
<rect x="115" y="55" width="5" height="5" class="bg-color" />
<rect x="120" y="55" width="5" height="5" class="bg-color" />
<rect x="20" y="60" width="5" height="5" class="bg-color" />
<rect x="25" y="60" width="5" height="5" class="bg-color" />
<rect x="30" y="60" width="5" height="5" class="color" />
<rect x="35" y="60" width="5" height="5" class="color" />
<rect x="40" y="60" width="5" height="5" class="color" />
<rect x="45" y="60" width="5" height="5" class="color" />
<rect x="50" y="60" width="5" height="5" class="color" />
<rect x="55" y="60" width="5" height="5" class="color" />
<rect x="60" y="60" width="5" height="5" class="bg-color" />
<rect x="65" y="60" width="5" height="5" class="color" />
<rect x="70" y="60" width="5" height="5" class="bg-color" />
<rect x="75" y="60" width="5" height="5" class="color" />
<rect x="80" y="60" width="5" height="5" class="bg-color" />
<rect x="85" y="60" width="5" height="5" class="color" />
<rect x="90" y="60" width="5" height="5" class="bg-color" />
<rect x="95" y="60" width="5" height="5" class="color" />
<rect x="100" y="60" width="5" height="5" class="color" />
<rect x="105" y="60" width="5" height="5" class="color" />
<rect x="110" y="60" width="5" height="5" class="color" />
<rect x="115" y="60" width="5" height="5" class="bg-color" />
<rect x="120" y="60" width="5" height="5" class="color" />
<rect x="20" y="65" width="5" height="5" class="bg-color" />
<rect x="25" y="65" width="5" height="5" class="bg-color" />
<rect x="30" y="65" width="5" height="5" class="bg-color" />
<rect x="35" y="65" width="5" height="5" class="color" />
<rect x="40" y="65" width="5" height="5" class="bg-color" />
<rect x="45" y="65" width="5" height="5" class="bg-color" />
<rect x="50" y="65" width="5" height="5" class="bg-color" />
<rect x="55" y="65" width="5" height="5" class="color" />
<rect x="60" y="65" width="5" height="5" class="bg-color" />
<rect x="65" y="65" width="5" height="5" class="bg-color" />
<rect x="70" y="65" width="5" height="5" class="color" />
<rect x="75" y="65" width="5" height="5" class="color" />
<rect x="80" y="65" width="5" height="5" class="bg-color" />
<rect x="85" y="65" width="5" height="5" class="bg-color" />
<rect x="90" y="65" width="5" height="5" class="bg-color" />
<rect x="95" y="65" width="5" height="5" class="bg-color" />
<rect x="100" y="65" width="5" height="5" class="bg-color" />
<rect x="105" y="65" width="5" height="5" class="color" />
<rect x="110" y="65" width="5" height="5" class="color" />
<rect x="115" y="65" width="5" height="5" class="bg-color" />
<rect x="120" y="65" width="5" height="5" class="bg-color" />
<rect x="20" y="70" width="5" height="5" class="color" />
<rect x="25" y="70" width="5" height="5" class="color" />
<rect x="30" y="70" width="5" height="5" class="color" />
<rect x="35" y="70" width="5" height="5" class="bg-color" />
<rect x="40" y="70" width="5" height="5" class="bg-color" />
<rect x="45" y="70" width="5" height="5" class="color" />
<rect x="50" y="70" width="5" height="5" class="color" />
<rect x="55" y="70" width="5" height="5" class="bg-color" />
<rect x="60" y="70" width="5" height="5" class="bg-color" />
<rect x="65" y="70" width="5" height="5" class="bg-color" />
<rect x="70" y="70" width="5" height="5" class="bg-color" />
<rect x="75" y="70" width="5" height="5" class="bg-color" />
<rect x="80" y="70" width="5" height="5" class="bg-color" />
<rect x="85" y="70" width="5" height="5" class="color" />
<rect x="90" y="70" width="5" height="5" class="bg-color" />
<rect x="95" y="70" width="5" height="5" class="bg-color" />
<rect x="100" y="70" width="5" height="5" class="bg-color" />
<rect x="105" y="70" width="5" height="5" class="color" />
<rect x="110" y="70" width="5" height="5" class="bg-color" />
<rect x="115" y="70" width="5" height="5" class="color" />
<rect x="120" y="70" width="5" height="5" class="color" />
<rect x="20" y="75" width="5" height="5" class="bg-color" />
<rect x="25" y="75" width="5" height="5" class="bg-color" />
<rect x="30" y="75" width="5" height="5" class="bg-color" />
<rect x="35" y="75" width="5" height="5" class="color" />
<rect x="40" y="75" width="5" height="5" class="bg-color" />
<rect x="45" y="75" width="5" height="5" class="color" />
<rect x="50" y="75" width="5" height="5" class="bg-color" />
<rect x="55" y="75" width="5" height="5" class="bg-color" />
<rect x="60" y="75" width="5" height="5" class="color" />
<rect x="65" y="75" width="5" height="5" class="color" />
<rect x="70" y="75" width="5" height="5" class="bg-color" />
<rect x="75" y="75" width="5" height="5" class="color" />
<rect x="80" y="75" width="5" height="5" class="bg-color" />
<rect x="85" y="75" width="5" height="5" class="bg-color" />
<rect x="90" y="75" width="5" height="5" class="color" />
<rect x="95" y="75" width="5" height="5" class="bg-color" />
<rect x="100" y="75" width="5" height="5" class="color" />
<rect x="105" y="75" width="5" height="5" class="bg-color" />
<rect x="110" y="75" width="5" height="5" class="bg-color" />
<rect x="115" y="75" width="5" height="5" class="bg-color" />
<rect x="120" y="75" width="5" height="5" class="color" />
<rect x="20" y="80" width="5" height="5" class="color" />
<rect x="25" y="80" width="5" height="5" class="color" />
<rect x="30" y="80" width="5" height="5" class="color" />
<rect x="35" y="80" width="5" height="5" class="color" />
<rect x="40" y="80" width="5" height="5" class="color" />
<rect x="45" y="80" width="5" height="5" class="bg-color" />
<rect x="50" y="80" width="5" height="5" class="color" />
<rect x="55" y="80" width="5" height="5" class="bg-color" />
<rect x="60" y="80" width="5" height="5" class="bg-color" />
<rect x="65" y="80" width="5" height="5" class="color" />
<rect x="70" y="80" width="5" height="5" class="color" />
<rect x="75" y="80" width="5" height="5" class="color" />
<rect x="80" y="80" width="5" height="5" class="color" />
<rect x="85" y="80" width="5" height="5" class="color" />
<rect x="90" y="80" width="5" height="5" class="bg-color" />
<rect x="95" y="80" width="5" height="5" class="bg-color" />
<rect x="100" y="80" width="5" height="5" class="bg-color" />
<rect x="105" y="80" width="5" height="5" class="color" />
<rect x="110" y="80" width="5" height="5" class="color" />
<rect x="115" y="80" width="5" height="5" class="color" />
<rect x="120" y="80" width="5" height="5" class="color" />
<rect x="20" y="85" width="5" height="5" class="bg-color" />
<rect x="25" y="85" width="5" height="5" class="bg-color" />
<rect x="30" y="85" width="5" height="5" class="bg-color" />
<rect x="35" y="85" width="5" height="5" class="bg-color" />
<rect x="40" y="85" width="5" height="5" class="bg-color" />
<rect x="45" y="85" width="5" height="5" class="bg-color" />
<rect x="50" y="85" width="5" height="5" class="bg-color" />
<rect x="55" y="85" width="5" height="5" class="bg-color" />
<rect x="60" y="85" width="5" height="5" class="bg-color" />
<rect x="65" y="85" width="5" height="5" class="color" />
<rect x="70" y="85" width="5" height="5" class="color" />
<rect x="75" y="85" width="5" height="5" class="color" />
<rect x="80" y="85" width="5" height="5" class="color" />
<rect x="85" y="85" width="5" height="5" class="bg-color" />
<rect x="90" y="85" width="5" height="5" class="bg-color" />
<rect x="95" y="85" width="5" height="5" class="color" />
<rect x="100" y="85" width="5" height="5" class="bg-color" />
<rect x="105" y="85" width="5" height="5" class="color" />
<rect x="110" y="85" width="5" height="5" class="bg-color" />
<rect x="115" y="85" width="5" height="5" class="color" />
<rect x="120" y="85" width="5" height="5" class="bg-color" />
<rect x="20" y="90" width="5" height="5" class="color" />
<rect x="25" y="90" width="5" height="5" class="color" />
<rect x="30" y="90" width="5" height="5" class="color" />
<rect x="35" y="90" width="5" height="5" class="color" />
<rect x="40" y="90" width="5" height="5" class="color" />
<rect x="45" y="90" width="5" height="5" class="color" />
<rect x="50" y="90" width="5" height="5" class="color" />
<rect x="55" y="90" width="5" height="5" class="bg-color" />
<rect x="60" y="90" width="5" height="5" class="color" />
<rect x="65" y="90" width="5" height="5" class="color" />
<rect x="70" y="90" width="5" height="5" class="color" />
<rect x="75" y="90" width="5" height="5" class="bg-color" />
<rect x="80" y="90" width="5" height="5" class="color" />
<rect x="85" y="90" width="5" height="5" class="bg-color" />
<rect x="90" y="90" width="5" height="5" class="color" />
<rect x="95" y="90" width="5" height="5" class="bg-color" />
<rect x="100" y="90" width="5" height="5" class="color" />
<rect x="105" y="90" width="5" height="5" class="color" />
<rect x="110" y="90" width="5" height="5" class="bg-color" />
<rect x="115" y="90" width="5" height="5" class="bg-color" />
<rect x="120" y="90" width="5" height="5" class="color" />
<rect x="20" y="95" width="5" height="5" class="color" />
<rect x="25" y="95" width="5" height="5" class="bg-color" />
<rect x="30" y="95" width="5" height="5" class="bg-color" />
<rect x="35" y="95" width="5" height="5" class="bg-color" />
<rect x="40" y="95" width="5" height="5" class="bg-color" />
<rect x="45" y="95" width="5" height="5" class="bg-color" />
<rect x="50" y="95" width="5" height="5" class="color" />
<rect x="55" y="95" width="5" height="5" class="bg-color" />
<rect x="60" y="95" width="5" height="5" class="color" />
<rect x="65" y="95" width="5" height="5" class="color" />
<rect x="70" y="95" width="5" height="5" class="bg-color" />
<rect x="75" y="95" width="5" height="5" class="bg-color" />
<rect x="80" y="95" width="5" height="5" class="bg-color" />
<rect x="85" y="95" width="5" height="5" class="bg-color" />
<rect x="90" y="95" width="5" height="5" class="bg-color" />
<rect x="95" y="95" width="5" height="5" class="color" />
<rect x="100" y="95" width="5" height="5" class="color" />
<rect x="105" y="95" width="5" height="5" class="color" />
<rect x="110" y="95" width="5" height="5" class="color" />
<rect x="115" y="95" width="5" height="5" class="bg-color" />
<rect x="120" y="95" width="5" height="5" class="bg-color" />
<rect x="20" y="100" width="5" height="5" class="color" />
<rect x="25" y="100" width="5" height="5" class="bg-color" />
<rect x="30" y="100" width="5" height="5" class="color" />
<rect x="35" y="100" width="5" height="5" class="color" />
<rect x="40" y="100" width="5" height="5" class="color" />
<rect x="45" y="100" width="5" height="5" class="bg-color" />
<rect x="50" y="100" width="5" height="5" class="color" />
<rect x="55" y="100" width="5" height="5" class="bg-color" />
<rect x="60" y="100" width="5" height="5" class="color" />
<rect x="65" y="100" width="5" height="5" class="color" />
<rect x="70" y="100" width="5" height="5" class="bg-color" />
<rect x="75" y="100" width="5" height="5" class="color" />
<rect x="80" y="100" width="5" height="5" class="bg-color" />
<rect x="85" y="100" width="5" height="5" class="bg-color" />
<rect x="90" y="100" width="5" height="5" class="bg-color" />
<rect x="95" y="100" width="5" height="5" class="bg-color" />
<rect x="100" y="100" width="5" height="5" class="bg-color" />
<rect x="105" y="100" width="5" height="5" class="color" />
<rect x="110" y="100" width="5" height="5" class="bg-color" />
<rect x="115" y="100" width="5" height="5" class="color" />
<rect x="120" y="100" width="5" height="5" class="color" />
<rect x="20" y="105" width="5" height="5" class="color" />
<rect x="25" y="105" width="5" height="5" class="bg-color" />
<rect x="30" y="105" width="5" height="5" class="color" />
<rect x="35" y="105" width="5" height="5" class="color" />
<rect x="40" y="105" width="5" height="5" class="color" />
<rect x="45" y="105" width="5" height="5" class="bg-color" />
<rect x="50" y="105" width="5" height="5" class="color" />
<rect x="55" y="105" width="5" height="5" class="bg-color" />
<rect x="60" y="105" width="5" height="5" class="color" />
<rect x="65" y="105" width="5" height="5" class="color" />
<rect x="70" y="105" width="5" height="5" class="color" />
<rect x="75" y="105" width="5" height="5" class="color" />
<rect x="80" y="105" width="5" height="5" class="bg-color" />
<rect x="85" y="105" width="5" height="5" class="color" />
<rect x="90" y="105" width="5" height="5" class="bg-color" />
<rect x="95" y="105" width="5" height="5" class="color" />
<rect x="100" y="105" width="5" height="5" class="bg-color" />
<rect x="105" y="105" width="5" height="5" class="color" />
<rect x="110" y="105" width="5" height="5" class="bg-color" />
<rect x="115" y="105" width="5" height="5" class="color" />
<rect x="120" y="105" width="5" height="5" class="bg-color" />
<rect x="20" y="110" width="5" height="5" class="color" />
<rect x="25" y="110" width="5" height="5" class="bg-color" />
<rect x="30" y="110" width="5" height="5" class="color" />
<rect x="35" y="110" width="5" height="5" class="color" />
<rect x="40" y="110" width="5" height="5" class="color" />
<rect x="45" y="110" width="5" height="5" class="bg-color" />
<rect x="50" y="110" width="5" height="5" class="color" />
<rect x="55" y="110" width="5" height="5" class="bg-color" />
<rect x="60" y="110" width="5" height="5" class="color" />
<rect x="65" y="110" width="5" height="5" class="color" />
<rect x="70" y="110" width="5" height="5" class="color" />
<rect x="75" y="110" width="5" height="5" class="color" />
<rect x="80" y="110" width="5" height="5" class="bg-color" />
<rect x="85" y="110" width="5" height="5" class="bg-color" />
<rect x="90" y="110" width="5" height="5" class="color" />
<rect x="95" y="110" width="5" height="5" class="color" />
<rect x="100" y="110" width="5" height="5" class="bg-color" />
<rect x="105" y="110" width="5" height="5" class="bg-color" />
<rect x="110" y="110" width="5" height="5" class="color" />
<rect x="115" y="110" width="5" height="5" class="color" />
<rect x="120" y="110" width="5" height="5" class="bg-color" />
<rect x="20" y="115" width="5" height="5" class="color" />
<rect x="25" y="115" width="5" height="5" class="bg-color" />
<rect x="30" y="115" width="5" height="5" class="bg-color" />
<rect x="35" y="115" width="5" height="5" class="bg-color" />
<rect x="40" y="115" width="5" height="5" class="bg-color" />
<rect x="45" y="115" width="5" height="5" class="bg-color" />
<rect x="50" y="115" width="5" height="5" class="color" />
<rect x="55" y="115" width="5" height="5" class="bg-color" />
<rect x="60" y="115" width="5" height="5" class="bg-color" />
<rect x="65" y="115" width="5" height="5" class="bg-color" />
<rect x="70" y="115" width="5" height="5" class="color" />
<rect x="75" y="115" width="5" height="5" class="bg-color" />
<rect x="80" y="115" width="5" height="5" class="bg-color" />
<rect x="85" y="115" width="5" height="5" class="bg-color" />
<rect x="90" y="115" width="5" height="5" class="color" />
<rect x="95" y="115" width="5" height="5" class="color" />
<rect x="100" y="115" width="5" height="5" class="bg-color" />
<rect x="105" y="115" width="5" height="5" class="bg-color" />
<rect x="110" y="115" width="5" height="5" class="bg-color" />
<rect x="115" y="115" width="5" height="5" class="bg-color" />
<rect x="120" y="115" width="5" height="5" class="color" />
<rect x="20" y="120" width="5" height="5" class="color" />
<rect x="25" y="120" width="5" height="5" class="color" />
<rect x="30" y="120" width="5" height="5" class="color" />
<rect x="35" y="120" width="5" height="5" class="color" />
<rect x="40" y="120" width="5" height="5" class="color" />
<rect x="45" y="120" width="5" height="5" class="color" />
<rect x="50" y="120" width="5" height="5" class="color" />
<rect x="55" y="120" width="5" height="5" class="bg-color" />
<rect x="60" y="120" width="5" height="5" class="bg-color" />
<rect x="65" y="120" width="5" height="5" class="color" />
<rect x="70" y="120" width="5" height="5" class="bg-color" />
<rect x="75" y="120" width="5" height="5" class="bg-color" />
<rect x="80" y="120" width="5" height="5" class="color" />
<rect x="85" y="120" width="5" height="5" class="bg-color" />
<rect x="90" y="120" width="5" height="5" class="bg-color" />
<rect x="95" y="120" width="5" height="5" class="color" />
<rect x="100" y="120" width="5" height="5" class="color" />
<rect x="105" y="120" width="5" height="5" class="bg-color" />
<rect x="110" y="120" width="5" height="5" class="bg-color" />
<rect x="115" y="120" width="5" height="5" class="bg-color" />
<rect x="120" y="120" width="5" height="5" class="bg-color" />
</svg>
`
