package test

import (
	"testing"

	burnaboy "github.com/ertush/burna"
)

type testBurna interface {
	runTest()
}

type directorArtistTest struct {
	testResults string
}

type managerArtistTest struct {
	testResults string
}

func (ma *managerArtistTest) runTest(t *testing.T, m *burnaboy.ArtistManager, a *burnaboy.Artist, expectName string) {
	if a.GetArtistManager(a, m) != expectName {
		t.Fail()
		ma.testResults = "FAIL"
	}
	ma.testResults = "PASS"
}

func (da *directorArtistTest) runTest(artistName, expectName string, d *burnaboy.ArtistDirector, t *testing.T) {
	if d.GetArtistDirector(artistName) != expectName {
		t.Fail()
		da.testResults = "FAIL"
	}
	da.testResults = "PASS"

}

func entryPoint() {

	var ts *testing.T
	var m managerArtistTest

	artists := []*burnaboy.Artist{}

	artist := burnaboy.CreateArtist("Burna Boy", "31", "0723928832", "Afro Pop", "They Call Me Burna !", map[string]string{"location": "Lagos Nigeria", "street": "Obina 22", "ZipCode": "934"}, []string{"BritAwards2019", "MtvAwards2018", "KalashaAwrads2016"})
	artists = append(artists, artist)
	manager := &burnaboy.ArtistManager{UserKYCinfo: burnaboy.UserKYCinfo{Name: "John Doe", Age: "34", Tel: "019-322-443", Address: map[string]string{"location": "Abuja", "street": "joka 03", "ZipCode": "883"}}, Manages: artists}

	//director := &burnaboy.ArtistDirector{UserKYCinfo: burnaboy.UserKYCinfo{"Alex Martin", "24", "123-334-635", map[string]string{"location": "Montreal", "street": "OpenValley 13 st", "ZipCode": "443"}}, Boss: manager}
	m.runTest(ts, manager, artist, "John Doe")

	// if artist.GetArtistManager(artist, manager) != "John Doe" {
	// 	t.Fail()
	// }

	// if director.GetArtistDirector("Burna Boy") != "Alex Martin" {
	// 	t.Fail()
	// }
}
