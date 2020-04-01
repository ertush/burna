package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var usrname string
var pass string
var useBurna bool

func main() {

	flag.StringVar(&usrname, "username", "", "Username to be registered")
	flag.StringVar(&pass, "pass", "", "Password to be registered")
	flag.BoolVar(&useBurna, "useBurna", false, "Create an artist using burnaboy package")
	flag.Parse()

	viper.AddConfigPath(".")
	viper.SetConfigName("creds")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Unable to read Configuration file")

	}

	if usrname == "" {
		//viper.Set("credentials.username", usrname)
		fmt.Println("Username will be read from the configration file")
	}

	if pass == "" {
		//viper.Set("credentials.password", pass)
		fmt.Println("Password will be read from the configuration file")
	}

	if usrname != "" {
		viper.Set("credentials.username", usrname)
	}

	if pass != "" {
		viper.Set("credentials.password", pass)
	}

	if useBurna == true {
		artist := burnaboy.CreateArtist(
			"Burna Boy",
			"31",
			"(+112)723928832",
			"Afro Pop",
			"They Call Me Burna !",
			map[string]string{
				"location": "Lagos Nigeria",
				"street":   "Obina 22",
				"ZipCode":  "934"},
			[]string{"BritAwards2019", "MtvAwards2018", "KalashaAwrads2016"},
		)
		fmt.Printf("Artist Bio: \n")

		for k, i := range artist.GetArtistBio() {
			fmt.Println(k, i)
		}

	}

	fmt.Printf("Username -> %s\n", viper.GetString("credentials.username"))
	fmt.Printf("Password -> %s\n", viper.GetString("credentials.password"))
}
