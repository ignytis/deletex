package converter

import (
	"bufio"
	"strings"
)

func JavascriptToJSONL(s *bufio.Scanner, w *bufio.Writer) error {
	s.Scan() // skip the first line, as there is variable assignment + beginning of array
	if s.Err() != nil {
		return s.Err()
	}

	nestingLevel := 0
	for s.Scan() {
		line := strings.TrimSpace(s.Text())

		// Looks up curly braces in string, BUT skips values between quotation marks
		// The idea: to split by quotation marks and to skip items with even index
		// Examples:
		//   { "test": "ab{c" } -> SKIPS "test" and "ab{c", checks '  { ', ': ' and ' }
		for i, part := range strings.Split(line, "\"") {
			if i%2 == 1 {
				continue
			}

			nestingLevel = nestingLevel + strings.Count(part, "{") - strings.Count(part, "}")
		}

		if 0 == nestingLevel {
			if strings.HasSuffix(line, ",") { // end if element
				line = line[:len(line)-1] + "\n"
			} else if strings.HasSuffix(line, "]") {
				// close square bracket at the end of file
				// open bracket was eliminated by first invocation of scanner.Scan() method outside of loop
				break
			}
		}
		_, err := w.Write([]byte(line))
		if err != nil {
			return err
		}
	}

	if s.Err() != nil {
		return s.Err()
	}

	err := w.Flush()
	if err != nil {
		return err
	}

	return nil
}
