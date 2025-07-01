package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"io/ioutil"
)

type Bee struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var bees = []Bee{
	{"Honey Bee", "A social bee known for producing honey."},
	{"Bumblebee", "A large bee known for its fuzzy appearance."},
	{"Carpenter Bee", "A solitary bee that nests in wood."},
	{"Leafcutter Bee", "Known for cutting leaves to build nests."},
	{"Sweat Bee", "Attracted to human sweat as a moisture source."},
	{"Mining Bee", "Solitary bees that dig tunnel nests."},
	{"Cuckoo Bee", "Bees that lay their eggs in other bees' nests."},
	{"Mason Bee", "Known for using mud to build their nests."},
	{"Carder Bee", "Bees that collect plant fibers for nests."},
	{"Digger Bee", "Bees that dig soil to make nests."},
	{"Orchid Bee", "Known for their vibrant metallic colors."},
	{"Alkali Bee", "Nests in alkaline soils."},
	{"Blueberry Bee", "Pollinates blueberry plants."},
	{"Squash Bee", "Pollinates squash and gourds exclusively."},
	{"Stingless Bee", "Known for its absence of a stinger."},
	{"Killer Bee", "A hybrid known for aggressive behavior."},
	{"Longhorn Bee", "Named for their long antennae."},
	{"Sunflower Bee", "Specializes in pollinating sunflowers."},
	{"Melipona Bee", "A type of stingless bee."},
	{"Euglossine Bee", "Also known as orchid bees."},
	{"Nomad Bee", "Known for parasitizing other bees' nests."},
	{"Adrena Bee", "Solitary ground-nesting bees."},
	{"Nomia Bee", "Known for striking abdominal bands."},
	{"Tawny Mining Bee", "Known for their bright orange color."},
	{"Common Carder Bee", "Famous for collecting plant fibers."},
	{"Ashy Mining Bee", "Identified by their black and white color."},
	{"Red Mason Bee", "A proficient orchard pollinator."},
	{"Wool Carder Bee", "Collects wool from plants to build nests."},
	{"Plasterer Bee", "Uses its saliva to plaster nest walls."},
	{"Silk Bee", "Our fictive bee known to spin silk from plants."},
}

func getBeeHandler(bee Bee) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(bee)
	}
}

func setupBeesEndpoints() {
	for _, bee := range bees {
		endpoint := fmt.Sprintf("/get%s", bee.Name)
		endpoint = filepath.Clean(endpoint)
		http.HandleFunc(endpoint, getBeeHandler(bee))
	}
}

func initStorage() {
	pvcPath := "/data/bees"
	if err := os.MkdirAll(pvcPath, 0755); err != nil {
		panic(err)
	}
	dataPath := filepath.Join(pvcPath, "bee_data.json")

	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		beeData, _ := json.Marshal(bees)
		ioutil.WriteFile(dataPath, beeData, 0644)
	}
}

func main() {
	initStorage()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to stutestalpha API!")
	})
	setupBeesEndpoints()
	http.ListenAndServe(":8080", nil)
}
