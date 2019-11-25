package mockgen_sample

type InvalidPasswordError struct{}

func (e InvalidPasswordError) Error() string {
	return "Password does not meet requirements -- it should have at least 8 characters"
}

type EmptyStorageError struct{}

func (e EmptyStorageError) Error() string {
	return "Storage is empty -- no user was created yet"
}

type UserAlreadyExistsError struct{}

func (e UserAlreadyExistsError) Error() string {
	return "Username taken already"
}

type UserDoesNotExistError struct{}

func (e UserDoesNotExistError) Error() string {
	return "No user registered with this username"
}
