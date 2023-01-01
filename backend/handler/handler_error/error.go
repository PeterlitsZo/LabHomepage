package handlerError

import "errors"

var (
	AuthorizationIsEmpty error = errors.New("Authorization is empty")
	TokenIsInvalid       error = errors.New("Token is invalid")
	JwtIsInvalid         error = errors.New("JWT is invalid")
	PermissionDenied     error = errors.New("Permission denied")
	NewsNotExist         error = errors.New("News not exist")
	NewsAlreadyExist     error = errors.New("News already exist")
	NewsTitleEmpty       error = errors.New("News title is empty")
	PaperNotExist        error = errors.New("Paper not exist")
	PaperAlreadyExist    error = errors.New("Paper Already exist")
	PaperTitleEmpty      error = errors.New("Paper title is empty")
	PersonNotExist       error = errors.New("Person not exist")
	PersonAlreadyExist   error = errors.New("Person Already exist")
	PersonNameEmpty      error = errors.New("Person name is empty")
	ResourceNotExist     error = errors.New("Resource not exist")
	ResourceAlreadtExist error = errors.New("Resource Already exist")
	ResourceTitleEmpty   error = errors.New("Resource title is empty")
	UserNotExist         error = errors.New("User not exist")
	PasswordNotCorrect   error = errors.New("Password not correct")
	UserAlreadyExist     error = errors.New("User already exist")
	UserNameEmpty        error = errors.New("User name is empty")
	UserIdIsInvalid      error = errors.New("User id is invalid")
	UsernameIsInvalid    error = errors.New("Username is invalid")
)
