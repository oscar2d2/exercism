package protein

var m = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromCodon(codon string) string {
	return m[codon]
}

func FromRNA(rna string) []string {
	result := make([]string, 0)
	for i := 0; i < len(rna); i += 3 {
		if temp := m[rna[i:i+3]]; temp == "STOP" {
			break
		} else {
			result = append(result, temp)
		}
	}
	return result
}
