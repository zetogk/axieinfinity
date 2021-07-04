package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Axie struct {
	name         string
	namekey      string
	parents      []*Axie
	currentBreed int
	slpPricing   int
}

func calcBreedingPrice(currentBreed int) int {

	if currentBreed > 7 {

		return 0

	}

	slpPricing := make(map[int]int)

	slpPricing[1] = 150
	slpPricing[2] = 300
	slpPricing[3] = 450
	slpPricing[4] = 750
	slpPricing[5] = 1200
	slpPricing[6] = 1950
	slpPricing[7] = 3150

	return slpPricing[currentBreed]

}

func main() {

	a := &Axie{
		name:         "santorini",
		namekey:      "A",
		currentBreed: 4,
	}

	b := &Axie{
		name:         "ankara",
		namekey:      "B",
		currentBreed: 4,
	}

	c := &Axie{
		name:         "firenze",
		namekey:      "C",
		currentBreed: 4,
	}

	executeTree(a, b, c)

}

func createChildren(parentA, parentB *Axie) (Axie, error) {

	if parentA.currentBreed == 7 || parentB.currentBreed == 7 {

		return Axie{}, errors.New("currentBreed exceeds the limit")

	}

	parentA.currentBreed = parentA.currentBreed + 1
	parentB.currentBreed = parentB.currentBreed + 1

	fmt.Println("s")

	return Axie{
		name:         generateName(),
		namekey:      fmt.Sprintf("%s%s", parentA.namekey, parentB.namekey),
		parents:      []*Axie{parentA, parentB},
		currentBreed: 0,
		slpPricing:   calcBreedingPrice(parentA.currentBreed) + calcBreedingPrice(parentB.currentBreed),
	}, nil

}

func executeTree(a, b, c *Axie) {

	axs := []*Axie{}

	axs = append(axs, a)
	axs = append(axs, b)
	axs = append(axs, c)

	ab, _ := createChildren(a, b)
	axs = append(axs, &ab)

	abc, _ := createChildren(&ab, c)
	axs = append(axs, &abc)

	aabc, _ := createChildren(a, &abc)
	axs = append(axs, &aabc)

	aabcc, _ := createChildren(&aabc, c)
	axs = append(axs, &aabcc)

	abcaabcc, _ := createChildren(&abc, &aabcc)
	axs = append(axs, &abcaabcc)

	abcaabccc, _ := createChildren(&abcaabcc, c)
	axs = append(axs, &abcaabccc)

	_a, _ := createChildren(b, &abcaabcc)
	axs = append(axs, &_a)

	_b, _ := createChildren(&ab, &abcaabccc)
	axs = append(axs, &_b)

	_c, _ := createChildren(&aabcc, &abcaabccc)
	axs = append(axs, &_c)

	fmt.Println("axs", axs)

	storeInFile(axs, "tree.csv")

}

func storeInFile(axs []*Axie, fileName string) {

	record := []string{
		"axies_name", "axies_key_name", "final_breeding_count", "parent_1", "parent_2", "slp_cost",
	}

	records := []string{}

	records = append(records, strings.Join(record, ","))

	for _, axie := range axs {

		parent1 := "-"
		parent2 := "-"

		if len(axie.parents) == 2 {

			parent1 = axie.parents[0].name
			parent2 = axie.parents[1].name

		}

		record := []string{
			axie.name, axie.namekey, strconv.Itoa(axie.currentBreed), parent1, parent2, strconv.Itoa(axie.slpPricing),
		}

		records = append(records, strings.Join(record, ","))

	}

	pathFile := fmt.Sprintf("./%s", fileName)
	writeLines(records, pathFile)

}

func generateName() string {

	syl := []string{"wa", "we", "wi", "wo", "wu", "ra", "re", "ri", "ro", "ru", "ta", "te", "ti", "to", "tu", "ya", "ye", "yi", "yo", "yu", "pa", "pe", "pi", "po", "pu", "sa", "se", "si", "so", "su", "da", "de", "di", "do", "du", "fa", "fe", "fi", "fo", "fu", "ga", "ge", "gi", "go", "gu", "ha", "he", "hi", "ho", "hu", "ja", "je", "ji", "jo", "ju", "ka", "ke", "ki", "ko", "ku", "la", "le", "li", "lo", "lu", "za", "ze", "zi", "zo", "zu", "xa", "xe", "xi", "xo", "xu", "ca", "ce", "ci", "co", "cu", "va", "ve", "vi", "vo", "vu", "ba", "be", "bi", "bo", "bu", "na", "ne", "ni", "no", "nu", "ma", "me", "mi", "mo", "mu"}

	numSyl := getRandomInt(3, 6)

	name := ""

	for i := 0; i < numSyl; i++ {

		sylIndex := getRandomInt(0, len(syl))
		name = name + syl[sylIndex]

	}

	return name

}

func getRandomInt(min, max int) int {

	return rand.Intn(max-min) + min

}

func writeLines(lines []string, path string) error {

	// overwrite file if it exists
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	// new writer w/ default 4096 buffer size
	w := bufio.NewWriter(file)

	for _, line := range lines {
		_, err := w.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	// flush outstanding data
	return w.Flush()
}
