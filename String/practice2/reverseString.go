package main

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}

	reverseString(s)
}

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1

	for i < j {
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
		i++
		j--
	}
}
