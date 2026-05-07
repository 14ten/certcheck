package main

import (
	"fmt"
	"io"
	"text/tabwriter"
)

func writeTable(w io.Writer, results []Result) {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "HOST\tEXPIRES\tDAYS\tISSUER")
	for _, r := range results {
		if r.Error != "" {
			fmt.Fprintf(tw, "%s\t-\t-\t%s\n", r.Host, r.Error)
			continue
		}
		fmt.Fprintf(tw, "%s\t%s\t%d\t%s\n",
			r.Host, r.NotAfter.Format("2006-01-02"), r.DaysLeft, r.Issuer)
	}
	tw.Flush()
}
