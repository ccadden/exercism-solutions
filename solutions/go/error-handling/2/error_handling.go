package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource

	for resource, err = opener(); err != nil; resource, err = opener() {
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}

	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(FrobError); ok {
				resource.Defrob(e.defrobTag)
			}

			err = r.(error)
		}
	}()

	resource.Frob(input)

	return nil
}
