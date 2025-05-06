package utils

import (
	"encoding/json"
	"fmt"
	structure "groupie-tracker/models"
	"net/http"
)

func FetchArtists() ([]structure.Artists, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []structure.Artists

	// Désérialiser le JSON dans la structure
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func FetchRelation(Id string) (structure.Relation, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", Id)
	resp, err := http.Get(url)
	if err != nil {
		return structure.Relation{}, err
	}
	defer resp.Body.Close()

	var relation structure.Relation

	// Désérialiser le JSON dans la structure
	err = json.NewDecoder(resp.Body).Decode(&relation)
	if err != nil {
		return structure.Relation{}, err
	}

	return relation, nil
}
