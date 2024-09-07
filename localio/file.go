package localio

import "os"

func CreateOrUpdateFile(filename, content string) error {
	// Convert content string to bytes
	data := []byte(content)

	// Write the data to the file with 0644 permissions (readable and writable by the owner)
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func SaveToFile(filename, text string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		return err
	}

	return nil
}

func ReadFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func SaveBytesToFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644) // 0644 is the file permission mode
	if err != nil {
		return err
	}
	return nil
}

func ReadBytesFromFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}
