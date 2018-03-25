package strand

var m = map[byte]byte{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	result := make([]byte, len(dna))

	for i := 0; i < len(dna); i++ {
		result[i] = m[dna[i]]
	}

	return string(result)
}
