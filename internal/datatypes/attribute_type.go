package datatypes

type AttributeType string

const (
	STRING   AttributeType = "STRING"
	DATE     AttributeType = "DATE"
	DATETIME AttributeType = "DATETIME"
	NUMBER   AttributeType = "NUMBER"
)

func (ct *AttributeType) Scan(value interface{}) error {
	*ct = AttributeType(value.(string))
	return nil
}

func (ct AttributeType) Value() (string, error) {
	return string(ct), nil
}

func (ct AttributeType) IsValid() bool {
	switch ct {
	case STRING, DATE, DATETIME, NUMBER:
		return true
	}

	return false
}

func (ct AttributeType) IsNotValid() bool {
	return !ct.IsValid()
}
