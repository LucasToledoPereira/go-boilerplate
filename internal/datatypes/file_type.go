package datatypes

type FileType string

const (
	COVER  FileType = "COVER"
	ANY    FileType = "ANY"
	AVATAR FileType = "AVATAR"
)

func (ct *FileType) Scan(value interface{}) error {
	*ct = FileType(value.(string))
	return nil
}

func (ct FileType) Value() (string, error) {
	return string(ct), nil
}

func (ct FileType) IsValid() bool {
	switch ct {
	case COVER, ANY, AVATAR:
		return true
	}

	return false
}

func (ct FileType) IsNotValid() bool {
	return !ct.IsValid()
}

func (ct FileType) IsType(compareWith FileType) bool {
	return ct == compareWith
}
