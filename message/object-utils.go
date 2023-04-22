package message

func MakeObjectWithKeyAndValueSame(inputMessage string) string {
	return "{" + inputMessage + ":" + inputMessage + "}"
}

func MakeObjectWithKeyAndValue(inputMessageKey string, inputMessageValue string) string {
	return "{" + inputMessageKey + ":" + inputMessageValue + "}"
}
