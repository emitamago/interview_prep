package main

var keyChar = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// getKey --> Compute short url from long url
func genKey(count int) string {
	if count == 0 {
		return string(keyChar[0])
	}
	l := len(keyChar)
	s := make([]byte, 20) // FIXME: will overflow. eventually
	i := len(s)
	for count > 0 && i >= 0 {
		i--
		j := count % l
		count = (count - j) / l
		s[i] = keyChar[j]
	}
	return string(s[i:])
}
