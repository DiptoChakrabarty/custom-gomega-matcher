package test

import (
	"bytes"
	"fmt"
	"testing"

	//	. "github.com/onsi/ginkgo/v2"

	"github.com/olekukonko/tablewriter"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

func TestTable(t *testing.T) {
	tests := []struct {
		name     string
		expected []string
		want     bool
	}{
		{
			name: "Testing Random Table",
			want: true,
			expected: []string{
				"+----------+-------------+-------+---------+\n",
				"|   DATE   | DESCRIPTION |  CV2  | AMOUNT  |\n",
				"+----------+-------------+-------+---------+\n",
				"| 1/1/2014 | Domain name |  2233 | $10.98  |\n",
				"+----------+-------------+-------+---------+\n",
				"|                          TOTAL | $146.93 |\n",
				"+----------+-------------+-------+---------+\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)

			data := [][]string{
				[]string{"1/1/2014", "Domain name", "2233", "$10.98"},
			}
			var output bytes.Buffer
			table := tablewriter.NewWriter(&output)
			table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
			table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer
			table.SetBorder(true)                                 // Set Border to false
			table.AppendBulk(data)

			g.Expect(output.String()).Should(MatchTable(tt.expected))
		})
	}
}

type Table struct {
	tableData []string
}

func MatchTable(expected []string) types.GomegaMatcher {
	return &Table{tableData: expected}
}

func (t *Table) Match(actual interface{}) (bool, error) {
	fmt.Println(t.tableData)
	fmt.Println("Data Table")
	//table := actual.(bytes.Buffer)
	tableString := actual.(string)
	fmt.Println(tableString)
	fmt.Println("Real Table")
	tableLength := 0

	for i := range t.tableData {
		for j := range t.tableData[i] {
			if string(t.tableData[i][j]) != string(tableString[tableLength]) {
				return false, fmt.Errorf("Wrong Table Given")
			}
			tableLength += 1
		}
	}
	return true, nil
}

func (t *Table) FailureMessage(actual interface{}) string {
	return "Expected and received table not equal"
}

func (t *Table) NegatedFailureMessage(actual interface{}) string {
	return "Expected and received table not equal"
}
