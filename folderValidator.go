package main

func folderValidation(userId, folderName string) error {

	if !isValidUserId(userId) {
		return InvalidUserID
	}

	if folderName == "" {
		return InvalidFolderName
	}

	return nil
}
