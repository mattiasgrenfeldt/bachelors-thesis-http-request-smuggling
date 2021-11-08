package harness

import (
	"encoding/hex"
	"fmt"
	"html"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"
)

const reportDir = "reports"
const mainReport = "report%s.html"

var reportTemplate = template.Must(template.ParseFiles("harness/report.template.html"))

type result struct {
	Name             string
	Responses        []string
	BackendResponses []string
}

func GenerateReport(reportType string, serverResults map[string][]Response, proxyResults map[string][]Response, reqs []Request) error {
	if err := mkdir(reportDir); err != nil {
		return err
	}

	var newReqs []Request
	for _, r := range reqs {
		newReqs = append(newReqs, Request{
			Name: r.Name,
			Data: []byte(fullReport(r.Data)),
		})
	}

	data := struct {
		Columns       int
		Requests      []Request
		ServerResults []result
		ProxyResults  []result
	}{
		Columns:       len(reqs),
		Requests:      newReqs,
		ServerResults: prepareResults(reportType, serverResults),
		ProxyResults:  prepareResults(reportType, proxyResults),
	}

	mini := ""
	if reportType == "mini" {
		mini = ".mini"
	}

	fname := fmt.Sprintf("report-%s-%d_requests%s.html", time.Now().String(), len(reqs), mini)
	path := filepath.Join(reportDir, fname)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := reportTemplate.Execute(f, data); err != nil {
		return err
	}
	return copyFile(path, fmt.Sprintf(mainReport, mini))
}

func prepareResults(reportType string, res map[string][]Response) []result {
	var newRes []result
	for service, resps := range res {
		var data []string
		var backendData []string
		for _, r := range resps {
			if reportType == "full" {
				data = append(data, fullReport(r.Data))
			} else if reportType == "mini" {
				data = append(data, miniReport(r.Data))
			}
			backendData = append(backendData, fullReport(r.BackendData))
		}
		newRes = append(newRes, result{
			Name:             service,
			Responses:        data,
			BackendResponses: backendData,
		})
	}

	sort.Slice(newRes, func(i, j int) bool {
		return newRes[i].Name < newRes[j].Name
	})
	return newRes
}

func fullReport(r []byte) string {
	return highlightNewlines(html.EscapeString(string(r)))
}

func miniReport(r []byte) string {
	s := html.EscapeString(string(r))
	index := strings.Index(s, "\r\n")
	if index == -1 {
		return s
	}
	return s[:index]
}

func copyFile(src, dst string) error {
	return exec.Command("cp", src, dst).Run()
}

func mkdir(dir string) error {
	// drwxr-xr-x
	return os.MkdirAll(dir, os.ModeDir|0755)
}

func highlightNewlines(d string) string {
	var s strings.Builder
	n := len(d)
	for i := 0; i < n; i++ {
		b := d[i]
		if i != n-1 && d[i:i+2] == "\r\n" {
			s.WriteString("<span class=\"crlf\">\\r\\n\r\n</span>")
			i++
		} else if b == '\r' {
			s.WriteString("<span class=\"cr\">\\r\r</span>")
		} else if b == '\n' {
			s.WriteString("<span class=\"lf\">\\n\n</span>")
		} else if b == '\t' {
			s.WriteString("<span class=\"tab\">\\t</span>")
		} else if b < 32 || b >= 127 {
			s.WriteString("<span class=\"hex\">\\x" + hex.EncodeToString([]byte{b}) + "</span>")
		} else {
			s.WriteByte(d[i])
		}
	}
	return s.String()
}
