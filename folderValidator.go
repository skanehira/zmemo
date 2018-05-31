package main

func folderValidation(userName, folderName string) error {

	if !isValidUserName(userName) {
		return InvalidUserName
	}

	if folderName == "" {
		return InvalidFolderName
	}

	return nil
}
