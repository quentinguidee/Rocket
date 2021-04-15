package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

type Launch struct {
	Name     string
	Details  string
	Rocket   string
	Date_utc string
}

type Rocket struct {
	Name string
}

func main() {
	flag.String("command", "next", "Command to execute")
	flag.Parse()

	launch := new(Launch)
	rocket := new(Rocket)

	fmt.Println()
	font := color.New(color.FgHiMagenta).Add(color.Bold)

	if flag.Arg(0) == "latest" {
		font.Println("=> LATEST LAUNCH\n")
		GetLatestLaunch(launch)
	} else {
		font.Println("=> NEXT LAUNCH\n")
		GetNextLaunch(launch)
	}
	GetRocket(rocket, launch.Rocket)
	PrintLaunch(launch, rocket)
}

func PrintLaunch(launch *Launch, rocket *Rocket) {
	PrintLine("Name", launch.Name)
	PrintLine("Date", launch.Date_utc)
	PrintLine("Rocket", rocket.Name)
	PrintLine("Description", launch.Details)
}

func PrintLine(tag string, content string) {
	titleFont := color.New(color.FgYellow).Add(color.Bold)

	titleFont.Printf(tag + " ")
	fmt.Println(content)
}

func GetLatestLaunch(launch interface{}) error {
	return Request("launches/latest", launch)
}

func GetNextLaunch(launch interface{}) error {
	return Request("launches/next", launch)
}

func GetRocket(rocket interface{}, id string) error {
	return Request("rockets/"+id, rocket)
}

func Request(path string, object interface{}) error {
	resp, err := http.Get("https://api.spacexdata.com/v4/" + path)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(object)
}
