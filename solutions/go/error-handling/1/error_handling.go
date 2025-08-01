package erratum

import "errors"

func Use(opener ResourceOpener, input string) error {
	resource, err := retry(opener)

	if err != nil {
		return err
	}

	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				resource.Defrob(e.defrobTag)
				err = e
			default:
				err = errors.New("Unrecoverable error")
			}
		}
	}()

	resource.Frob(input)

	return err
}

func retry(opener ResourceOpener) (Resource, error) {
	for {
		resource, err := opener()

		if err != nil {
			switch err.(type) {
			case TransientError:
				continue
			default:
				return nil, err
			}

		}

		return resource, nil
	}
}
