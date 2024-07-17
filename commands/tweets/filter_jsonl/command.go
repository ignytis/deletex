package filter_jsonl

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/ignytis/deletex/types"
	"github.com/knetic/govaluate"
)

const expr = "tweet.retweetCount() > 20"

func MustRun(inputFile string, outputFile string, expr string) {
	log.Println("Filtering JSON Lines record in file '" + inputFile + "'...")

	in, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("Unable to read input file "+inputFile+". ", err)
	}
	defer in.Close()

	out, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("Unable to read output file "+outputFile+". ", err)
	}
	defer out.Close()

	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		log.Fatal("Failed to initialize an expresstion '"+expr+"'. ", err)
	}

	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)

	for scanner.Scan() {
		line := scanner.Text()

		var record types.Record
		err := json.Unmarshal([]byte(line), &record)
		if err != nil {
			log.Fatal("Unable to read marshal a line: "+line+". ", err)
		}

		toDeleteI, err := expression.Evaluate(map[string]interface{}{
			"created_time":   record.Tweet.CreatedAt().Unix(),
			"favorite_count": record.Tweet.FavoriteCount(),
			"full_text":      record.Tweet.FullText,
			"id":             record.Tweet.Id,
			"retweet_count":  record.Tweet.RetweetCount(),
		})

		if err != nil {
			log.Fatal("Failed to evaluate an expression: "+expr+". ", err)
		}

		toDelete := toDeleteI.(bool)

		if toDelete {
			writer.WriteString(line + "\n")
		}
		writer.Flush()
	}

	log.Println("Created a list of files to delete: " + outputFile)
}
