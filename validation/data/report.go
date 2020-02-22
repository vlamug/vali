package data

import (
	"bufio"
	"io"
)

type Report interface {
	Add(msg string)
	Publish() error
}

type WriterReport struct {
	msgs   []string
	writer *bufio.Writer
}

func NewWriterReport(writer io.Writer) *WriterReport {
	return &WriterReport{writer: bufio.NewWriter(writer), msgs: make([]string, 0)}
}

func (r *WriterReport) Add(msg string) {
	r.msgs = append(r.msgs, msg)
}

func (r *WriterReport) Publish() error {
	for _, v := range r.msgs {
		if _, err := r.writer.WriteString(v + "\n"); err != nil {
			return err
		}
	}

	return r.writer.Flush()
}
