package erratum

func Use(o ResourceOpener, input string) (err error) {
	var resource Resource
	for {
		resource, err = o()

		if err == nil {
			break
		}

		switch err.(type) {
		case TransientError:
			continue
		default:
			return
		}
	}
	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {
			if asFrob, ok := r.(FrobError); ok {
				resource.Defrob(asFrob.defrobTag)
			}

			if asError, ok := r.(error); ok {
				err = asError
			}
		}
	}()

	resource.Frob(input)
	return
}
