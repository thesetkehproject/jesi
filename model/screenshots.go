package model

type Screenshot struct {
	Name      string 	`json:"Name"`         //Image Name / Filename
	Size      int64 	`json:"Size"`         //Image File Size
	User      string 	`json:"User"`         //Name of the user uploading the file
	UnixTime  int64     `json:"UnixTime"`     //Unix Time the Image was Captured
	Thumb64   string    `json:"Thumb64"`      //Base64 Encoded Thumbnail Image
	Image64   string    `json:"Image64"`      //Base64 Encoded Full Image
}
