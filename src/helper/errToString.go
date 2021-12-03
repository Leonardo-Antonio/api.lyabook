package helper

func ErrToString(errs []error) (errsMessage []string) {
	for i := 0; i < len(errs); i++ {
		errsMessage = append(errsMessage, errs[i].Error())
	}
	return errsMessage
}
