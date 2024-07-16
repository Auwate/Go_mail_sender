package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func EnvSetup(buffer io.Reader) error {

	if buffer == nil {
		buffer = os.Stdin
	}

	reader := bufio.NewReader(buffer)
	fmt.Println("Please enter your gmail address.")
	email, err := reader.ReadString('\n')

	if err != nil {
		return err
	}

	os.Setenv("EMAIL", email)

	fmt.Println("Please enter your gmail's app password (For instructions, look up app password on Google. It will require 2FA.)")
	pass, err := reader.ReadString('\n')

	if err != nil {
		return err
	}

	os.Setenv("PASSWORD", pass)
	return nil

}
