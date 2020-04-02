package main

import (
	"fmt"
	"github.com/ertush/burna"
)


func main(){
 
artists := []*burnaboy.Artist{}

artist := burnaboy.CreateArtist(

	"Burna Boy", 
	"31", 
	"0723928832", 
	"Afro Pop", 
	"They Call Me Burna !", 
	map[string]string{
	 "location" : "Lagos Nigeria",
	 "street" : "Obina 22",
	 "ZipCode" : "934",
	},
	[]string{"BritAwards2019", "MtvAwards2018", "KalashaAwrads2016"},
	
	)
 
 artists = append(artists, artist)

 manager := &burnaboy.ArtistManager{
	UserKYCinfo:burnaboy.UserKYCinfo{
		"John Doe", 
		"34", 
		"019-322-443", 
		map[string]string{"location": "Abuja", "street": "joka 03", "ZipCode":"883"}}, 
		Manages:artists,}
 
 director := &burnaboy.ArtistDirector{
		UserKYCinfo:burnaboy.UserKYCinfo{
		"Monte Carol", 
		"44", 
		"019-399-743", 
		map[string]string{"location": "Ugenya", "street": "mypopy 03", "ZipCode":"435"}},
		Boss:manager,}


 fmt.Println(artist.GetArtistManager(artist, manager))

 fmt.Println(director.GetArtistDirector(artist.UserKYCinfo.Name))

 //fmt.Printf("%s has won the following awards %v\n", artist.GetArtistBio()["Name"], artist.GetArtistAwards())
}
