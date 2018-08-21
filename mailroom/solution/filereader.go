package solution

// Remove _ once package is used.
import (
	_ "encoding/json"
	"os"

	"github.com/SeedJobs/devops-go-problems/mailroom/types"
)

// readContent will take a path to a json file that contains
// an array of all the addresses of letters to be sent.
func readContent(path string) ([]types.Letter, error) {
	if _, err := os.Open(path); os.IsNotExist(err) {
		return nil, err
	}

	// Need to read all the content from the file and marshall it into
	// an array of types that satisfy the return type.
	return nil, nil
}

// validateLetter will examine all the content of the letter
// and determine if the letter is valid or not.
func validateLetter(l types.Letter) error {

	return nil
}
