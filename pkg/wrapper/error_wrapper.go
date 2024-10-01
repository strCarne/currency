package wrapper

import "fmt"

func Wrap(location string, message string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("[%s] %s: %w", location, message, err)
}
