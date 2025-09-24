package main

import "fmt"

func getNameCounts(names []string) map[string]map[string]int {
	result := map[string]map[string]int{}
	for _, name := range names {
		// fmt.Println(name)
		firstChar := string(name[0]) // Need to convert first char to string or else you can't use it as key in map
		_, charExists := result[firstChar]
		// fmt.Println(result, firstChar, result[firstChar], charExists)
		if !charExists {
			result[firstChar] = map[string]int{}
		}

		// NOTE: Not needed since the line "result[firstChar] = map[string]int{}" initialized key-value with a default value of 0 since its type is int
		// _, nameExists := result[firstChar][name]
		// if !nameExists {
		// 	result[firstChar][name] = 0
		// }

		result[firstChar][name]++
	}

	return result
}

// NOTE: If you wanna use runes instead of string in 1st level of map
// func getNameCounts2(names []string) map[rune]map[string]int {
// 	result := map[rune]map[string]int{}
// 	for _, name := range names {
// 		// fmt.Println(name)
// 		firstChar := rune(name[0]) // Need to convert first char to string or else you can't use it as key in map
// 		_, charExists := result[firstChar]
// 		// fmt.Println(result, firstChar, result[firstChar], charExists)
// 		if !charExists {
// 			result[firstChar] = map[string]int{}
// 		}

// 		// #NOTE: Not needed since the line "result[firstChar] = map[string]int{}" initialized key-value with a default value of 0 since its type is int
// 		// _, nameExists := result[firstChar][name]
// 		// if !nameExists {
// 		// 	result[firstChar][name] = 0
// 		// }

// 		result[firstChar][name]++
// 	}

// 	return result
// }

func getTestNames(length int) []string {
	return []string{
		"Grant", "Eduardo", "Peter", "Matthew", "Matthew", "Matthew", "Peter", "Peter", "Henry", "Parker",
		"Parker", "Parker", "Collin", "Hayden", "George", "Bradley", "Mitchell", "Devon", "Ricardo", "Shawn",
		"Taylor", "Nicolas", "Gregory", "Francisco", "Liam", "Kaleb", "Preston", "Erik", "Alexis", "Owen",
		"Omar", "Diego", "Dustin", "Corey", "Fernando", "Clayton", "Carter", "Ivan", "Jaden", "Javier",
		"Alec", "Johnathan", "Scott", "Manuel", "Cristian", "Alan", "Raymond", "Brett", "Max", "Drew",
		"Andres", "Gage", "Mario", "Dawson", "Dillon", "Cesar", "Wesley", "Levi", "Jakob", "Chandler",
		"Martin", "Malik", "Edgar", "Sergio", "Trenton", "Josiah", "Nolan", "Marco", "Drew", "Peyton",
		"Harrison", "Drew", "Hector", "Micah", "Roberto", "Drew", "Brady", "Erick", "Conner", "Jonah",
		"Casey", "Jayden", "Edwin", "Emmanuel", "Andre", "Phillip", "Brayden", "Landon", "Giovanni", "Bailey",
		"Ronald", "Braden", "Damian", "Donovan", "Ruben", "Frank", "Gerardo", "Pedro", "Andy", "Chance",
		"Abraham", "Calvin", "Trey", "Cade", "Donald", "Derrick", "Payton", "Darius", "Enrique", "Keith",
		"Raul", "Jaylen", "Troy", "Jonathon", "Cory", "Marc", "Eli", "Skyler", "Rafael", "Trent",
		"Griffin", "Colby", "Johnny", "Chad", "Armando", "Kobe", "Caden", "Marcos", "Cooper", "Elias",
		"Brenden", "Israel", "Avery", "Zane", "Zane", "Zane", "Zane", "Dante", "Josue", "Zackary",
		"Allen", "Philip", "Mathew", "Dennis", "Leonardo", "Ashton", "Philip", "Philip", "Philip", "Julio",
		"Miles", "Damien", "Ty", "Gustavo", "Drake", "Jaime", "Simon", "Jerry", "Curtis", "Kameron",
		"Lance", "Brock", "Bryson", "Alberto", "Dominick", "Jimmy", "Kaden", "Douglas", "Gary", "Brennan",
		"Zachery", "Randy", "Louis", "Larry", "Nickolas", "Albert", "Tony", "Fabian", "Keegan", "Saul",
		"Danny", "Tucker", "Myles", "Damon", "Arturo", "Corbin", "Deandre", "Ricky", "Kristopher", "Lane",
		"Pablo", "Darren", "Jarrett", "Zion", "Alfredo", "Micheal", "Angelo", "Carl", "Oliver", "Kyler",
		"Tommy", "Walter", "Dallas", "Jace", "Quinn", "Theodore", "Grayson", "Lorenzo", "Joe", "Arthur",
		"Bryant", "Roman", "Brent", "Russell", "Ramon", "Lawrence", "Moises", "Aiden", "Quentin", "Jay",
		"Tyrese", "Tristen", "Emanuel", "Salvador", "Terry", "Morgan", "Jeffery", "Esteban", "Tyson", "Braxton",
		"Branden", "Marvin", "Brody", "Craig", "Ismael", "Rodney", "Isiah", "Marshall", "Maurice", "Ernesto",
		"Emilio", "Brendon", "Kody", "Eddie", "Malachi", "Abel", "Keaton", "Jon", "Shaun", "Skylar",
		"Ezekiel", "Nikolas", "Santiago", "Kendall", "Axel", "Camden", "Trevon", "Bobby", "Conor", "Jamal",
		"Lukas", "Malcolm", "Zackery", "Jayson", "Javon", "Roger", "Reginald", "Zachariah", "Desmond", "Felix",
		"Johnathon", "Dean", "Quinton", "Ali", "Davis", "Gerald", "Rodrigo", "Demetrius", "Billy", "Rene",
		"Reece", "Kelvin", "Leo", "Justice", "Chris", "Guillermo", "Matthew", "Matthew", "Matthew", "Kevon",
		"Steve", "Frederick", "Clay", "Weston", "Dorian", "Hugo", "Roy", "Orlando", "Terrance", "😊",
		"Kai", "Khalil", "Khalil", "Khalil", "Graham", "Noel", "Willie", "Nathanael", "Terrell",
	}[:length]
}

func main() {
	testCase1 := getTestNames(10)
	res1 := getNameCounts(testCase1)
	fmt.Println("res1", res1)

	testCase2 := getTestNames(50)
	res2 := getNameCounts(testCase2)
	fmt.Println("res2", res2)

	testCase3 := getTestNames(100)
	res3 := getNameCounts(testCase3)
	fmt.Println("res3", res3)
}
