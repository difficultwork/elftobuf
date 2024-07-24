package lang

import (
	"encoding/hex"
	"fmt"
	"os"
)

func ElfToGo(elfFile, targetFile, packageName, variableName string) error {
	data, err := readElf(elfFile)
	if err != nil {
		fmt.Printf("readElf(%s) failed: %v.\n", elfFile, err)
		return err
	}

	os.Remove(targetFile)
	f, err := os.OpenFile(targetFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Printf("open(%s) failed: %v.\n", targetFile, err)
		return err
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("package %s\n\n", packageName))
	if err != nil {
		fmt.Printf("write(%s) failed: %v.\n", targetFile, err)
		return err
	}
	f.Sync()

	_, err = f.WriteString(fmt.Sprintf("var %s = []byte(\"", variableName))
	if err != nil {
		fmt.Printf("write(%s) failed: %v.\n", targetFile, err)
		return err
	}
	f.Sync()

	dataBuf := make([]byte, 2048)
	for len(data) >= 1024 {
		rdata := data[:1024]
		hex.Encode(dataBuf, rdata)
		if _, err = f.Write(dataBuf); err != nil {
			fmt.Printf("write(%s) failed: %v.\n", targetFile, err)
			return err
		}
		data = data[1024:]
	}

	if len(data) > 0 {
		wlen := hex.Encode(dataBuf, data)
		if _, err = f.Write(dataBuf[:wlen]); err != nil {
			fmt.Printf("write(%s) failed: %v.\n", targetFile, err)
			return err
		}
	}

	_, err = f.Write([]byte("\")\n"))
	if err != nil {
		fmt.Printf("write(%s) failed.\n", targetFile)
		return err
	}
	return nil
}
