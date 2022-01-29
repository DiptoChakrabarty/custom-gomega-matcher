package main

import (
	//"bufio"
	"bytes"
	"fmt"

	"github.com/olekukonko/tablewriter"
)

func main() {
	fmt.Println("This is New Table")
	//stdout := bufio.NewWriter(os.Stdout)
	var output bytes.Buffer

	//tableString := &strings.Builder{}
	data := [][]string{
		[]string{"1/1/2014", "Domain name", "2233", "$10.98"},
	}

	table := tablewriter.NewWriter(&output)
	table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
	table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer
	table.SetBorder(false)                                // Set Border to false
	table.AppendBulk(data)

	table.Render()
	tbl := output.String()
	fmt.Println(tbl)
	fmt.Println(len(tbl))
	//for i, _ := range tbl {
	//	fmt.Println(string(tbl[i]))
	//}
	//fmt.Println(stdout.Size())
	// https://stackoverflow.com/questions/26804642/how-to-test-a-functions-output-stdout-stderr-in-unit-tests
	// https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string

	//fmt.Println(tableString.String())
	//fmt.Println(len(tableString.String()))

	random := []string{
		"DATE   | DESCRIPTION |  CV2  | AMOUNT",
		"-----------+-------------+-------+----------",
		"  1/1/2014 | Domain name |  2233 | $10.98",
		"-----------+-------------+-------+----------",
		"						   TOTAL | $146.93",
		"						 --------+----------",
	}

	for i := range random {
		for j := range random[i] {
			for i := range tbl {
				if tbl[i] != random[i][j] {
					fmt.Println("Not Matching Correctly")
					break
				}
			}
		}
		//fmt.Println("\n")
	}
}
