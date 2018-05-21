package main

func memoValidation(memo Memo) error {

	if memo.UserID == "" {
		return NotFuondMemoID
	}

	if !isUUID.MatchString(memo.UserID) {
		return InvalidMemoID
	}

	if memo.Text == "" {
		return InvalidMemo
	}

	return nil
}

func isValidMemoId(memoId string) bool {
	return isUUID.MatchString(memoId)
}
