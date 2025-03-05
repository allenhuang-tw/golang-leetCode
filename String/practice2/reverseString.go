package main

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}

	reverseString(s)
}

func reverseString(s []byte) {

	i := 0
	j := len(s) - 1

	for i < j {
		j := s[i]
		s[i] = tmp
		tmp = j
	}
}
