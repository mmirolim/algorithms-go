package checks

func NeqErrs(e1, e2 error) bool {
	if e1 == e2 {
		return false
	}

	if e1 != nil && e2 != nil && e1.Error() == e2.Error() {
		return false
	}

	return true
}
