package gomod

import (
	"errors"
	"fmt"
)

//HelloWorld return hello
func HelloWorld(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hello World %s!", name), nil
	case "pt":
		return fmt.Sprintf("Oi,%s!", name), nil
	default:
		return "", errors.New("unknown language")

	}

}
