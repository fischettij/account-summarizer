package summarizer

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

const tableTemplate = `
		<!DOCTYPE html>
			<html>
			   <head>
				  <style>
					 table {
					 border-collapse: collapse;
					 width: 100%;
					 }
					 th, td {
					 border: none;
					 text-align: left;
					 padding: 8px;
					 }
				  </style>
			   </head>
			   <body>
				  <table>
					 {{range .}}
					 <tr>
						<td>{{index . 0}}</td>
						<td>{{index . 1}}</td>
					 </tr>
					 {{end}}
				  </table>
			   </body>
			</html>
	`

func generateHTMLBody(summary Summary) string {
	totalBalance := fmt.Sprintf("Total balance is %.2f", summary.Balance())
	debitAverage := fmt.Sprintf("Average debit amount: %.2f", summary.AverageDebitAmount())
	creditAverage := fmt.Sprintf("Average credit amount: %.2f", summary.AverageCreditAmount())

	data := [][]string{
		{totalBalance, debitAverage},
		{"", creditAverage},
	}

	for month := time.January; month <= time.December; month++ {
		value, found := summary.TransactionsByMonth()[month]
		if found && value > 0 {
			// Set first row of transactions by month
			if data[1][0] == "" {
				data[1] = []string{fmt.Sprintf("Number of transactions in %s: %d", month.String(), value), creditAverage}
			} else {
				data = append(data, []string{fmt.Sprintf("Number of transactions in %s: %d", month.String(), value), ""})
			}
		}
	}

	tmpl, err := template.New("table").Parse(tableTemplate)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, data)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
