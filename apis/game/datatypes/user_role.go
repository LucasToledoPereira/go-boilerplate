package datatypes

type UserRole string

const (
	OWNER          UserRole = "OWNER"
	ADMINISTRATOR  UserRole = "ADMINISTRATOR"
	DEVELOPER      UserRole = "DEVELOPER"
	LEVEL_DESIGNER UserRole = "LEVEL_DESIGNER"
	CREATOR        UserRole = "CREATOR"
	COMMON         UserRole = "COMMON"
)

func (ct *UserRole) Scan(value interface{}) error {
	*ct = UserRole(value.(string))
	return nil
}

func (ct UserRole) Value() (string, error) {
	return string(ct), nil
}

func (ct UserRole) IsValid() bool {
	switch ct {
	case OWNER, ADMINISTRATOR, DEVELOPER, LEVEL_DESIGNER, CREATOR, COMMON:
		return true
	}

	return false
}

func (ct UserRole) IsNotValid() bool {
	return !ct.IsValid()
}
