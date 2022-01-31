package test

import (
	"bytes"
	"testing"

	//	. "github.com/onsi/ginkgo/v2"

	"github.com/olekukonko/tablewriter"
	. "github.com/onsi/gomega"
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

			g.Expect(table).Should(MatchTable(tt.expected))
		})
	}
}
