package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strconv"
	"text/tabwriter"
)

// sortByExpiry orders results so the most urgent (least days, errors first) come first.
func sortByExpiry(results []Result) {
	sort.SliceStable(results, func(i, j int) bool {
		if (results[i].Error != "") != (results[j].Error != "") {
			return results[i].Error != ""
		}
		return results[i].DaysLeft < results[j].DaysLeft
	})
}

func writeJSON(w io.Writer, results []Result) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(results)
}

// exitCode returns 0 (ok), 1 (warn), 2 (crit), or 3 (error).
func exitCode(results []Result, warn, crit int) int {
	code := 0
	for _, r := range results {
		switch {
		case r.Error != "":
			if code < 3 {
				code = 3
			}
		case r.DaysLeft <= crit:
			if code < 2 {
				code = 2
			}
		case r.DaysLeft <= warn:
			if code < 1 {
				code = 1
			}
		}
	}
	return code
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

func writeCSV(w io.Writer, results []Result, warn, crit int) error {
	cw := csv.NewWriter(w)
	cw.Write([]string{"host", "expires", "days_left", "status", "issuer", "error"})
	for _, r := range results {
		s := status(r, warn, crit)
		expires := ""
		if r.Error == "" {
			expires = r.NotAfter.Format("2006-01-02")
		}
		cw.Write([]string{r.Host, expires, strconv.Itoa(r.DaysLeft), s, r.Issuer, r.Error})
	}
	cw.Flush()
	return cw.Error()
}

func writeTable(w io.Writer, results []Result, warn, crit int, color bool) {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "HOST\tEXPIRES\tDAYS\tSTATUS\tISSUER")
	for _, r := range results {
		s := status(r, warn, crit)
		sc := colorize(s, statusColor(s), color)
		if r.Error != "" {
			fmt.Fprintf(tw, "%s\t-\t-\t%s\t%s\n", r.Host, sc, r.Error)
			continue
		}
		fmt.Fprintf(tw, "%s\t%s\t%d\t%s\t%s\n",
			r.Host, r.NotAfter.Format("2006-01-02"), r.DaysLeft, sc, r.Issuer)
	}
	tw.Flush()
}
