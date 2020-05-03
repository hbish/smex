package stdout

import (
	"testing"

	"github.com/hbish/smex/pkg/xml"
	"github.com/stretchr/testify/assert"
)

var urls = []xml.URL{
	{Loc: "http://www.example.com/", LastMod: "2005-01-01", ChangeFreq: "monthly", Priority: 0.8},
	{Loc: "http://www.example.com/catalog?item=12&desc=vacation_hawaii", LastMod: "", ChangeFreq: "weekly", Priority: 0},
	{Loc: "http://www.example.com/catalog?item=73&desc=vacation_new_zealand", LastMod: "2004-12-23", ChangeFreq: "weekly", Priority: 0},
	{Loc: "http://www.example.com/catalog?item=74&desc=vacation_newfoundland", LastMod: "2004-12-23T18:00:15+00:00", ChangeFreq: "", Priority: 0.3},
	{Loc: "http://www.example.com/catalog?item=83&desc=vacation_usa", LastMod: "2004-11-23", ChangeFreq: "", Priority: 0},
}

func TestNewWriter(t *testing.T) {
	//given
	writer := NewWriter()
	//when
	//then
	assert.NotNil(t, writer)
}

func TestWriter_Write_All(t *testing.T) {
	// given
	writer := NewWriter()

	// when
	actual := writer.Write(urls, false)

	// then
	expected := "http://www.example.com/                                          \t2005-01-01          monthly\t0.80\n" +
		"http://www.example.com/catalog?item=12&desc=vacation_hawaii      \tweekly \t\n" +
		"http://www.example.com/catalog?item=73&desc=vacation_new_zealand \t2004-12-23          weekly \t\n" +
		"http://www.example.com/catalog?item=74&desc=vacation_newfoundland\t2004-12-23T18:00:15+00:000.30\n" +
		"http://www.example.com/catalog?item=83&desc=vacation_usa         \t2004-11-23          \n"
	assert.Equal(t, expected, actual)
}

func TestWriter_Write_OnlyLoc(t *testing.T) {
	// given
	writer := NewWriter()

	// when
	actual := writer.Write(urls, true)

	// then
	expected := "http://www.example.com/\t\n" +
		"http://www.example.com/catalog?item=12&desc=vacation_hawaii\t\n" +
		"http://www.example.com/catalog?item=73&desc=vacation_new_zealand\t\n" +
		"http://www.example.com/catalog?item=74&desc=vacation_newfoundland\t\n" +
		"http://www.example.com/catalog?item=83&desc=vacation_usa\t\n"
	assert.Equal(t, expected, actual)
}
