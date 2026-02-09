package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Bag struct {
	name string
	contains map[string]uint
	goldLink bool
}

func parseBags(path string) map[string]*Bag {
	strs := readLines(path)
	bags := make(map[string]*Bag)
	for _,str := range strs {
		newBag := &Bag{contains: make(map[string]uint)}

		x := strings.Split(str, "contain")
		bagInfo := strings.Fields(x[0])
		newBag.name = bagInfo[0] + " " + bagInfo[1]

		rules := strings.Split(x[1], ",")
		//fmt.Printf("Bag %s has:\n", newBag.name)
		for _,rule := range rules {
			parts := strings.Fields(rule)
			if parts[0] == "no" {
				break
			}
			amount, err := strconv.ParseUint(parts[0], 10, 64)
			if err != nil {
				panic(err)
			}
			ruleName := parts[1] + " " + parts[2]
			newBag.contains[ruleName] = uint(amount)
			if ruleName == "shiny gold" {
				newBag.goldLink = true
			}
			//fmt.Printf("\t%d x %s\n", amount, ruleName)
		}
		bags[newBag.name] = newBag
		//fmt.Printf("\n")
	}
	return bags
}

func adventDay7A(path string) {
	bags := parseBags(path)

	numChanged := 1
	for numChanged > 0 {
		numChanged = 0
		for _,bag := range bags {
			for contains := range bag.contains {
				if bags[contains].goldLink && bag.goldLink == false {
					bag.goldLink = true
					numChanged++
				}
			}
		}
	}
	numValid := 0
	for _,bag := range bags {
		if bag.goldLink {
			numValid++
			//fmt.Printf("%s could contain shiny gold\n", name)
		}
	}
	fmt.Printf("%d could contain shiny gold\n", numValid)

}


func numberOfBagsInsideBag(start *Bag, bags map[string]*Bag) uint {
	sum := uint(0)
	for name,value := range start.contains {
		sum += value * numberOfBagsInsideBag(bags[name], bags)
		sum += value
	}
	return sum
}

func adventDay7B(path string) {
	bags := parseBags(path)

	fmt.Printf("shiny gold contains %d bags\n", numberOfBagsInsideBag(bags["shiny gold"], bags))

}

