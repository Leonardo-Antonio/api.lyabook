package enum

type format struct {
	Digital,
	Fisico string
}

var Format = format{
	Digital: "Digital",
	Fisico:  "Fisico",
}

var Formats []string = []string{"d", "f", "df"}
