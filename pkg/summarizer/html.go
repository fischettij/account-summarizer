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
					 font-family: 'Arial, sans-serif';
					 font-size: 14px;
					 table {
					 border-collapse: collapse;
					 width: 100%;
					 }
					 th, td {
					 border: none;
					 text-align: left;
					 padding: 8px;
					 }
					 .logo-header {
					 background-color: rgb(0 209 128);
					 padding: 20px;
					 }
				  </style>
			   </head>
			   <body>
				  <div class="logo-header">
					   <a href="https://storicard.com/">
						  <svg xmlns="http://www.w3.org/2000/svg" width="129" height="48" viewBox="0 0 129 48" fill="#003A40" class="w-[88px] h-8 lg:w-[129px] lg:h-12 fill-white" style="fill: white;">
							 <path d="M31.7527 16.1824V0.54541H15.8764V6.46813C15.8764 7.54372 14.9917 8.415 13.8997 8.415H0V24H13.8997C14.9917 24 15.8764 24.8712 15.8764 25.9468V29.8706C15.8764 30.9462 14.9917 31.8175 13.8997 31.8175H0V47.4545H15.8764V41.5839C15.8764 40.5083 16.761 39.637 17.853 39.637H31.7527V24H17.853C16.761 24 15.8764 23.1287 15.8764 22.0531V18.1293C15.8764 17.0537 16.761 16.1824 17.853 16.1824H31.7527Z"></path>
							 <path fill-rule="evenodd" clip-rule="evenodd" d="M90.7134 40.1557H88.6351C82.3533 40.1557 77.2408 35.1203 77.2408 28.9332V26.8862C77.2408 20.6991 82.3533 15.6637 88.6351 15.6637H90.7134C96.9952 15.6637 102.108 20.6971 102.108 26.8862V28.9332C102.108 35.1203 96.9952 40.1557 90.7134 40.1557ZM88.6351 21.4542C85.5928 21.4542 83.1199 23.8918 83.1199 26.8862V28.9332C83.1199 31.9276 85.5948 34.3652 88.6351 34.3652H90.7134C93.7537 34.3652 96.2286 31.9276 96.2286 28.9332V26.8862C96.2286 23.8918 93.7537 21.4542 90.7134 21.4542H88.6351Z"></path>
							 <path d="M39.0715 32.857C39.3663 38.0106 43.4519 39.3666 44.552 39.681H44.5439C45.6563 39.9995 47.1815 40.1617 49.754 40.1617C53.75 40.1617 56.5503 39.5248 58.348 38.2669C60.1274 37.0211 60.9571 35.1423 60.9571 32.5645C60.9571 30.4134 60.2677 28.7329 58.9276 27.5411C57.7115 26.4616 55.9179 25.7866 53.7236 25.6123L48.5989 25.1156C47.7895 25.0294 47.1429 24.8031 46.6975 24.4266C46.244 24.044 46.0142 23.5212 46.0142 22.8883C46.0142 22.1993 46.2786 21.6324 46.7707 21.2419C47.2588 20.8573 47.9502 20.661 48.7819 20.661H50.5776C51.5476 20.661 52.3123 20.8413 52.8695 21.2499C53.3921 21.6344 53.7073 22.2033 53.8538 22.9444H60.644C60.4833 20.5769 59.4889 18.7682 57.6952 17.5484C55.8935 16.3206 53.8578 15.6897 49.9736 15.6897C48.001 15.6897 46.665 15.8379 45.5953 16.1263C44.5297 16.4168 39.8097 17.7467 39.8097 23.1367C39.8097 25.2217 40.4747 26.8962 41.705 28.098C42.9374 29.3017 44.7554 30.0488 47.0941 30.2371L52.1699 30.7339C52.721 30.7879 53.3474 30.9281 53.8416 31.2546C54.3479 31.5911 54.7058 32.1199 54.7058 32.9111C54.7058 33.6321 54.4415 34.215 53.9249 34.6096C53.4165 34.9981 52.6885 35.1884 51.7876 35.1884H49.2741C48.2939 35.1884 47.5272 35.0041 46.9537 34.5875C46.4128 34.195 46.0732 33.6161 45.8759 32.857H39.0715Z"></path>
							 <path d="M122.661 8.32286H129V14.542H123.497C123.035 14.542 122.661 14.1735 122.661 13.7188V8.32286Z"></path>
							 <path d="M120.947 22.0411H122.222C122.684 22.0411 123.058 22.4096 123.058 22.8643V39.637H129V16.2065H120.947V22.0411Z"></path>
							 <path d="M110.584 22.8643V39.637H104.644V24.3244C104.644 19.8398 108.334 16.2045 112.888 16.2045H119.009V22.0411H111.419C110.958 22.0411 110.584 22.4096 110.584 22.8643Z"></path>
							 <path d="M71.8857 33.8485H75.4933L75.4953 33.8465V39.637H71.8247C67.0721 39.637 63.2205 35.8434 63.2205 31.1625V9.96127H69.1627V15.3813C69.1627 15.8339 69.5369 16.2045 69.9985 16.2045H75.4933V22.0411H69.9985C69.5369 22.0411 69.1607 22.4096 69.1607 22.8643V31.1645C69.1607 32.6467 70.3808 33.8485 71.8857 33.8485Z"></path>
						  </svg>
					   </a>
					</div>
					   <h2>Hi! Here is the balance of your year</h2>
				  <table>
					 {{range .}}
					 <tr>
						<td>{{index . 0}}</td>
						<td>{{index . 1}}</td>
					 </tr>
					 {{end}}
				  </table>
                  <p>Regards, Stori team</p>
			   </body>
			</html>
	`

func GenerateHTMLBody(summary Summary) (string, error) {
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
		return "", fmt.Errorf("error generating html temaplte: %w", err)
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", fmt.Errorf("error generating filling email body: %w", err)
	}

	return buf.String(), nil
}
