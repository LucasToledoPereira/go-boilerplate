package codes

const (
	UserWithNicknameAlreadyExists Codes = "already.exists.user.with.nickname"
	UserWithEmailAlreadyExists    Codes = "already.exists.user.with.email"
	UserNotFound                  Codes = "user.not.found"
	UserAlreadyExists             Codes = "user.already.exists"
	UserInvalidFields             Codes = "user.invalid.fields"
	UserCreateFailed              Codes = "user.create.failed"
	UserCreateSuccess             Codes = "user.create.success"
	UserListsFailed               Codes = "user.list.failed"
	UserListsSuccess              Codes = "user.list.success"
	UserInvalidIdentity           Codes = "user.identity.invalid"
	UserReadSuccess               Codes = "user.read.success"
	UserDeleteSuccess             Codes = "user.delete.success"
)
