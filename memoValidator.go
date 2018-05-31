package main

func memoValidation(memo Memo) error {

	if !isValidUserName(memo.UserName) {
		return InvalidUserName
	}

	if memo.Title == "" {
		return NotFoundTitle
	}

	if memo.Text == "" {
		return InvalidMemo
	}

	return nil
}

func isValidMemoId(memoId string) bool {
	return isValidUUID(memoId)
}
