package test

import (
	"bytes"
	"fmt"
	"testing"
	"reflect"

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
			table.Render()
			//fmt.Println(output.String())
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
	tableString := actual.(string)
	tableString = strings.Split(tableString, "\n")

	for i := range t.tableData {
		if !reflect.DeepEqual(t.tableData[i], tableString[i]) {
			fmt.Println("THIS IS WRONG")
			return false,fmt.Errorf("InvaliD Tables given")
		}
	return true, nil
}

func (t *Table) FailureMessage(actual interface{}) string {
	return "Expected and received table not equal"
}

func (t *Table) NegatedFailureMessage(actual interface{}) string {
	return "Expected and received table not equal"
}
