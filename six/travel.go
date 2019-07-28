package six

import (
	"strconv"
	"strings"
)

// Travel returns the zipcode followed by the streets and addresses
func Travel(r, zipcode string) string {
	addresses := strings.Split(r, ",")

	streets := []string{}
	numbers := []string{}

	for _, a := range addresses {
		elems := strings.Split(a, " ")
		length := len(elems)
		if elems[length-2]+" "+elems[length-1] == zipcode {
			streetIndex := 0
			for i := 0; i < length; i++ {
				if _, err := strconv.Atoi(elems[i]); err == nil {
					streetIndex = i
					break
				}
			}
			numbers = append(numbers, elems[streetIndex])
			street := []string{}
			for index, e := range elems {
				if !(index <= streetIndex || index == length-1 || index == length-2) {
					street = append(street, string(e))
				}
			}
			streets = append(streets, strings.Join(street, " "))
		}
	}

	return zipcode + ":" + strings.Join(streets, ",") + "/" + strings.Join(numbers, ",")
}
