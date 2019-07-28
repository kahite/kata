package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"codewarrior/kata/eight"
	"codewarrior/kata/five"
	"codewarrior/kata/four"
	"codewarrior/kata/other"
	"codewarrior/kata/seven"
	"codewarrior/kata/six"
)

var travel = `123 Main Street St. Louisville OH 43071, 432 Main Long Road St. Louisville OH 43071,786 High Street Pollocksville NY 56432,
  54 Holy Grail Street Niagara Town ZP 32908, 3200 Main Rd. Bern AE 56210,1 Gordon St. Atlanta RE 13000,
  10 Pussy Cat Rd. Chicago EX 34342, 10 Gordon St. Atlanta RE 13000, 58 Gordon Road Atlanta RE 13000,
  22 Tokyo Av. Tedmondville SW 43098, 674 Paris bd. Abbeville AA 45521, 10 Surta Alley Goodtown GG 30654,
  45 Holy Grail Al. Niagara Town ZP 32908, 320 Main Al. Bern AE 56210, 14 Gordon Park Atlanta RE 13000,
  100 Pussy Cat Rd. Chicago EX 34342, 2 Gordon St. Atlanta RE 13000, 5 Gordon Road Atlanta RE 13000,
  2200 Tokyo Av. Tedmondville SW 43098, 67 Paris St. Abbeville AA 45521, 11 Surta Avenue Goodtown GG 30654,
  45 Holy Grail Al. Niagara Town ZP 32918, 320 Main Al. Bern AE 56215, 14 Gordon Park Atlanta RE 13200,
  100 Pussy Cat Rd. Chicago EX 34345, 2 Gordon St. Atlanta RE 13222, 5 Gordon Road Atlanta RE 13001,
  2200 Tokyo Av. Tedmondville SW 43198, 67 Paris St. Abbeville AA 45522, 11 Surta Avenue Goodville GG 30655,
  2222 Tokyo Av. Tedmondville SW 43198, 670 Paris St. Abbeville AA 45522, 114 Surta Avenue Goodville GG 30655,
  2 Holy Grail Street Niagara Town ZP 32908, 3 Main Rd. Bern AE 56210, 77 Gordon St. Atlanta RE 13000,
  100 Pussy Cat Rd. Chicago OH 13201`

var _ = Describe("Kata", func() {

	It("Travel", func() {
		Expect(four.Partitions(1)).To(Equal(1))
		Expect(four.Partitions(5)).To(Equal(7))
		Expect(four.Partitions(6)).To(Equal(11))
		Expect(four.Partitions(7)).To(Equal(15))
		Expect(four.Partitions(8)).To(Equal(22))
		Expect(four.Partitions(9)).To(Equal(30))
		Expect(four.Partitions(10)).To(Equal(42))
		Expect(four.Partitions(25)).To(Equal(1958))
	})

	It("Travel", func() {
		Expect(six.Travel(travel, "AA 45522")).To(Equal("AA 45522:Paris St. Abbeville,Paris St. Abbeville/67,670"))
		Expect(six.Travel(travel, "OH 430")).To(Equal("OH 430:/"))
		Expect(six.Travel(travel, "NY 56432")).To(Equal("NY 56432:High Street Pollocksville/786"))
		Expect(six.Travel(travel, "EX 34342")).To(Equal("EX 34342:Pussy Cat Rd. Chicago,Pussy Cat Rd. Chicago/10,100"))
	})

	It("ListSquared", func() {
		Expect(five.ListSquared(1, 250)).To(Equal([][]int{{1, 1}, {42, 2500}, {246, 84100}}))
		Expect(five.ListSquared(250, 500)).To(Equal([][]int{{287, 84100}}))
		Expect(five.ListSquared(300, 600)).To(Equal([][]int{}))
	})

	It("GetDivisors", func() {
		Expect(other.GetDivisors(12)).To(Equal([]int{1, 2, 3, 4, 6, 12}))
		Expect(other.GetDivisors(15)).To(Equal([]int{1, 3, 5, 15}))
		Expect(other.GetDivisors(41)).To(Equal([]int{1, 41}))
	})

	It("JosephusSurvivor", func() {
		Expect(five.JosephusSurvivor(7, 3)).To(Equal(4))
		Expect(five.JosephusSurvivor(11, 19)).To(Equal(10))
		Expect(five.JosephusSurvivor(40, 3)).To(Equal(28))
		Expect(five.JosephusSurvivor(14, 2)).To(Equal(13))
		Expect(five.JosephusSurvivor(100, 1)).To(Equal(100))
	})

	It("SplitStrings", func() {
		Expect(six.SplitStrings("abc")).To(Equal([]string{"ab", "c_"}))
		Expect(six.SplitStrings("abcdef")).To(Equal([]string{"ab", "cd", "ef"}))
	})

	It("ContainAllRots", func() {
		ContainAllRotsTest("bsjq", []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"}, true)
		ContainAllRotsTest("XjYABhR", []string{"TzYxlgfnhf", "yqVAuoLjMLy", "BhRXjYA", "YABhRXj", "hRXjYAB", "jYABhRX", "XjYABhR", "ABhRXjY"}, false)
		ContainAllRotsTest("", []string{}, true)
	})

	It("Revrot", func() {
		RevrotTest("1234", 0, "")
		RevrotTest("", 0, "")
		RevrotTest("1234", 5, "")
		RevrotTest("123456987654", 6, "234561876549")
		RevrotTest("123456987653", 6, "234561356789")
		RevrotTest("66443875", 4, "44668753")
		RevrotTest("66443875", 8, "64438756")
		RevrotTest("664438769", 8, "67834466")
		RevrotTest("123456779", 8, "23456771")
		RevrotTest("733049910872815764", 5, "330479108928157")
		RevrotTest("563000655734469485", 4, "0365065073456944")

	})

	It("PlayPass", func() {
		PlayPassTest("BORN IN 2015!", 1, "!4897 Oj oSpC")
		PlayPassTest("I LOVE YOU!!!", 1, "!!!vPz fWpM J")
		PlayPassTest("I LOVE YOU!!!", 0, "!!!uOy eVoL I")
		PlayPassTest("AAABBCCY", 1, "zDdCcBbB")
	})

	It("InArray", func() {
		var a1 = []string{"live", "arp", "strong"}
		var a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
		var r = []string{"arp", "live", "strong"}
		InArrayTest(a1, a2, r)

		a1 = []string{"cod", "code", "wars", "ewar", "ar"}
		a2 = []string{}
		r = []string{}
		InArrayTest(a1, a2, r)

		a1 = []string{"duplicates", "duplicates"}
		a2 = []string{"duplicates"}
		r = []string{"duplicates"}
		InArrayTest(a1, a2, r)
	})

	It("DigitalRoot", func() {
		Expect(six.DigitalRoot(16)).To(Equal(7))
		Expect(six.DigitalRoot(195)).To(Equal(6))
		Expect(six.DigitalRoot(992)).To(Equal(2))
		Expect(six.DigitalRoot(167346)).To(Equal(9))
		Expect(six.DigitalRoot(0)).To(Equal(0))
	})

	It("PartList", func() {
		PartListTest([]string{"I", "wish", "I", "hadn't", "come"},
			"(I, wish I hadn't come)(I wish, I hadn't come)(I wish I, hadn't come)(I wish I hadn't, come)")
		PartListTest([]string{"cdIw", "tzIy", "xDu", "rThG"},
			"(cdIw, tzIy xDu rThG)(cdIw tzIy, xDu rThG)(cdIw tzIy xDu, rThG)")
	})

	It("TwoSum", func() {
		Expect(six.TwoSum([]int{1, 2, 3}, 4)).To(Equal([2]int{0, 2}))
		Expect(six.TwoSum([]int{1234, 5678, 9012}, 14690)).To(Equal([2]int{1, 2}))
		Expect(six.TwoSum([]int{2, 2, 3}, 4)).To(Equal([2]int{0, 1}))
	})

	It("ReverseString", func() {
		Expect(eight.ReverseString("world")).To(Equal("dlrow"))
		Expect(eight.ReverseString("azerty")).To(Equal("ytreza"))
	})
})

func InArrayTest(array1 []string, array2 []string, exp []string) {
	var ans = six.InArray(array1, array2)
	Expect(ans).To(Equal(exp))
}

func PartListTest(arr []string, exp string) {
	var ans = seven.PartList(arr)
	Expect(ans).To(Equal(exp))
}

func PlayPassTest(a1 string, n int, exp string) {
	var ans = six.PlayPass(a1, n)
	Expect(ans).To(Equal(exp))
}

func RevrotTest(s string, n int, exp string) {
	var ans = six.Revrot(s, n)
	Expect(ans).To(Equal(exp))
}

func ContainAllRotsTest(strng string, arr []string, exp bool) {
	ans := seven.ContainAllRots(strng, arr)
	Expect(ans).To(Equal(exp))
}
