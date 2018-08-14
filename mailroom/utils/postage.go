package utils

import (
  "os"
  "regexp"
  "bufio"
)

var (
  whitespace = regexp.MustCompile(`^\s*$`)
)

// PostageReader takes a path and returns each address as
// a byte array.
func PostageReader(path string) ([][]byte, error) {
  file, err := os.Open(path)
  if os.IsNotExist(err) {
    return nil, err
  }
  defer file.Close()
  contents := [][]byte{}
  scanner := bufio.NewScanner(file)
  content := []byte{}
  for scanner.Scan() {
    switch {
    case whitespace.Match(scanner.Bytes()):
      if len(content) > 0 {
        // Trims off the additional new line character
        contents = append(contents, content[:len(content)-1])
      }
      content = []byte{}
    default:
      // Seperates content by new line to preserve file format
      content = append(content, append(scanner.Bytes(), '\n')...)
    }
  }
  if len(content) >0 {
    contents = append(contents, content[:len(content)-1])
  }
  if err = scanner.Err(); err != nil {
    return nil, err
  }
  return contents, nil
}
