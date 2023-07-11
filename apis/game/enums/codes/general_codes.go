package codes

const (
	NotAuthorized            Codes = "auth.unauthorized"
	InvalidParamID           Codes = "id.params.invalid"
	FileRetreiveFailed       Codes = "file.retrieve.error"
	FileUploadFailed         Codes = "file.upload.fail"
	FileUploadSuccess        Codes = "file.upload.success"
	FileNotImage             Codes = "file.not.image"
	FileNotFound             Codes = "file.not.found"
	FileListFailed           Codes = "file.lists.failed"
	FileListSuccess          Codes = "file.lists.success"
	FileDeleteSuccess        Codes = "file.delete.success"
	FileDeleteFailed         Codes = "file.delete.failed"
	CoverImageAlreadyExists  Codes = "cover.image.already.exists"
	InvalidLoginFields       Codes = "login.fields.invalid"
	InvalidRegisterFields    Codes = "register.fields.invalid"
	DatastoreAdapterNotFound Codes = "datastore.adapter.not.found"
	FilestoreAdapterNotFound Codes = "filestore.adapter.not.found"
)
