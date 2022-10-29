package main

import (
	"log"
	"m-theory/clients"
)

func main() {
	myClient := clients.NewPhysicistsInfoClient()

	physicists, err := myClient.GetAllPhysicists()
	if err != nil {
		log.Println("Could not fetch the physcists")
	}

	var albertUuid = ""

	for _, physicist := range physicists.Physicists {
		if physicist.FirstName == "Albert" {
			albertUuid = physicist.PhysicistId.UuidInString
		}
		log.Println(physicist.FirstName)
	}

	albert, _ := myClient.GetPhysicistByUuid(albertUuid)

	log.Println(albert.FirstName)

	myClient.GetPhysicistByCountryOfBirth("Germany")

}
