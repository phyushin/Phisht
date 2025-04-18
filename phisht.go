package main

import (
	"phisht/helpers"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices: []string{"Add"},
	}
}

func main() {

	//profile := helpers.Sending_Profile{Profile_name: "foo", From_address: "testing@example.com", Host: "localhost:587", Username: "phyu", Password: "phyu"}

	//helpers.Create_gophish_sending_profile(profile)
	helpers.Get_gophish_sending_profiles()

}
