package vlgo

// Enum a kind of alias for int values.
type Enum int

type imageTypeList struct {
	Identity Enum
	Profile  Enum
}

// ImageType types of image requests.
var ImageType = &imageTypeList{
	Identity: 0,
	Profile:  1,
}

// ValueOfImageType returns value for ImageType.
func (mode Enum) ValueOfImageType() string {
	modes := [...]string{
		"uploadIdentity",
		"uploadProfile",
	}
	return modes[mode]
}
