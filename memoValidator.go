package main

func memoValidation(memo Memo) error {

	if memo.UserID == "" {
		return NotFoundUserID
	}

	if !isUUID.MatchString(memo.UserID) {
		return InvalidUserID
	}

	if memo.Text == "" {
		return InvalidMemo
	}

	return nil
}

func isValidMemoId(memoId string) bool {
	return isUUID.MatchString(memoId)
}
