/*
Package song defines how we represent individual tracks within the program.
*/
package tunecloud

type Track struct {
	Title string
	Album string
	Artist string
	PathName string
	Length int
	Bitrate int
}
