


package burnaboy

//UserKYCinfo is to be exported
type UserKYCinfo struct{
	Name string
	Age string
	Tel string
	Address map[string]string
}

//Artist is to be exported
type Artist struct{ 
		UserKYCinfo 
		Genre  string
		Vibe string  
		Awards []string
		
}

//ArtistDirector is to be exported
type ArtistDirector struct{
  UserKYCinfo
  Boss *ArtistManager
}

//ArtistManager is to be exported
type ArtistManager struct{
   UserKYCinfo
   Manages []*Artist
}


func (m *ArtistManager) getManagerByName(aName string) string{

	var artist_name string
 	
	for _, i:=range m.Manages {
		if i.Name == aName {
		   artist_name = m.Name 
		}
	}
	return artist_name
}

//GetArtistManager is to be exported
func (a *Artist) GetArtistManager(ar *Artist, m *ArtistManager) string{
	return  m.getManagerByName(ar.Name)
}


//GetArtistDirector is to be exported
func (d *ArtistDirector) GetArtistDirector(aName string) string{
	
	var artistDirector string

	for _,i:=range d.Boss.Manages {
		if i.Name == aName{
			artistDirector = d.Name
		}
	}
	return artistDirector
}

//CreateArtistManager is to be exported
func CreateArtistManager(uname, age, tel string, addr map[string]string, mngs []*Artist) *ArtistManager {
 	return &ArtistManager{UserKYCinfo:UserKYCinfo{Name:uname, Age:age, Tel:tel, Address:addr}, Manages:mngs,}
}

//CreateArtistDirector is to be exported
func CreateArtistDirector(uname, age, tel string, addr map[string]string, boss *ArtistManager) *ArtistDirector {
	return &ArtistDirector{UserKYCinfo:UserKYCinfo{Name:uname, Age:age, Tel:tel, Address:addr}, Boss:boss,}
}

//GetArtistDetails is to be exported
func (a *Artist) GetArtistDetails() map[string]string {
	return map[string]string {
		"Details": a.UserKYCinfo.Name, 
		"Age":a.UserKYCinfo.Age, 
		"Tel":a.UserKYCinfo.Tel, 
		"Genre":a.Genre,
		"Vibe":a.Vibe,
		}
}

//GetArtistBio is to be exported
func (a *Artist) GetArtistBio() map[string]string {

     artistBio := map[string]string {"Name":a.Name, "Age":a.Age}
     
     for k, prop := range artistBio {
	artistBio[k] = prop	
     }
	return artistBio
}

//CreateArtist is to be exported
func CreateArtist ( aName, aAge, atel, aGenre, aVibe string, aAddress map[string]string, aAwards []string) *Artist{
 return &Artist{UserKYCinfo:UserKYCinfo{Name:aName, Age:aAge, Tel:atel, Address:aAddress}, Genre:aGenre, Vibe:aVibe, Awards:aAwards,}
}
