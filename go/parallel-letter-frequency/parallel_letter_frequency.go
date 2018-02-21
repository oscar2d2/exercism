package letter

func ConcurrentFrequency(strs []string) FreqMap {
	result := FreqMap{}
	c := make(chan FreqMap)
	for _, str := range strs {
		go func(s string) {
			c <- Frequency(s)
		}(str)
	}

	for range strs {
		for k, v := range <-c {
			result[k] += v
		}
	}

	return result

}
