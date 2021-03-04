package util

//IDGenerator id generator
var IDGenerator = NopIDGenerator

//NopIDGenerator nop id  generator which alawys return empty string
func NopIDGenerator() string {
	return ""
}
