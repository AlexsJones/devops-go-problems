package solution

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/AlexsJones/devops-go-problems/mailroom/types"
)

// readContent will take a path to a json file that contains
// an array of all the addresses of letters to be sent.
func readContent(path string) ([]types.Letter, error) {
	if _, err := os.Open(path); os.IsNotExist(err) {
		return nil, err
	}
	// Need to read all the content from the file and marshall it into
	// an array of types that satisfy the return type.
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var (
		letters []post
		ret     []types.Letter
	)
	err = json.Unmarshal(buff, &letters)
	for _, l := range letters {
		ret = append(ret, l)
	}
	return ret, err
}
