package main

func main() {
	var contryCapitalMap map[string] string
	contryCapitalMap = make(map[string] string)

	contryCapitalMap [ "France" ] = "Paris"
	contryCapitalMap [ "Italy" ] = "Rome"
	contryCapitalMap [ "Japan" ] = "Tokyo"
	contryCapitalMap [ "India" ] = "New delhi"

	for contry := range contryCapitalMap {
		println(contry, contryCapitalMap [contry])
	}

	capital, ok := contryCapitalMap [ "American" ]
	println(capital, ok)	

	delete( contryCapitalMap, "France" )
	for contry := range contryCapitalMap {
		println(contry, contryCapitalMap [ contry ])
	}
}
