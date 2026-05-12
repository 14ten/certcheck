package main

import (
	"encoding/json"
	"fmt"
	"io"
	"text/tabwriter"
)

func writeJSON(w io.Writer, results []Result) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(results)
}

func status(r Result, warn, crit int) string {
	if r.Error != "" {
		return "ERR"
	}
	switch {
	case r.DaysLeft <= crit:
		return "CRIT"
	case r.DaysLeft <= warn:
		return "WARN"
	default:
		return "OK"
	}
}

func writeTable(w io.Writer, results []Result, warn, crit int) {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "HOST\tEXPIRES\tDAYS\tSTATUS\tISSUER")
	for _, r := range results {
		s := status(r, warn, crit)
		if r.Error != "" {
			fmt.Fprintf(tw, "%s\t-\t-\t%s\t%s\n", r.Host, s, r.Error)
			continue
		}
		fmt.Fprintf(tw, "%s\t%s\t%d\t%s\t%s\n",
			r.Host, r.NotAfter.Format("2006-01-02"), r.DaysLeft, s, r.Issuer)
	}
	tw.Flush()
}
