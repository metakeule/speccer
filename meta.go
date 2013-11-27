package main

import (
	"bytes"
	"fmt"
	"strconv"
)

var metas = map[string]bool{
	"RESPONSIBLE":    true,
	"STATE":          true,
	"DEADLINE":       true,
	"LASTUPDATE":     true,
	"ESTIMATEDHOURS": true,
	"TITLE":          true,
}

func (s *speccer) getMeta(meta string) string {
	switch meta {
	case "UUID":
		return s.Paragraph.UUID
	case "TITLE":
		return s.Paragraph.Title
	case "RESPONSIBLE":
		return s.Paragraph.Responsible
	case "STATE":
		return s.Paragraph.State
	case "DEADLINE":
		return s.Paragraph.Deadline
	case "LASTUPDATE":
		return s.Paragraph.LastUpdate
	case "ESTIMATEDHOURS":
		return fmt.Sprintf("%v", s.Paragraph.EstimatedHours)
	default:
		panic(fmt.Sprintf("unknown meta: %#v", meta))
	}
}

func (s *speccer) setMeta(meta string, val string) error {
	switch meta {
	case "TITLE":
		s.Paragraph.Title = val
		return nil
	case "LASTUPDATE":
		return fmt.Errorf("LastUpdate is only readable")
	case "RESPONSIBLE":
		s.Paragraph.Responsible = val
		return nil
	// TODO check if state is a valid state
	case "STATE":
		s.Paragraph.State = val
		return nil
	// TODO: check if deadline is in correct format
	case "DEADLINE":
		s.Paragraph.Deadline = val
		return nil
	case "ESTIMATEDHOURS":
		i, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		s.Paragraph.EstimatedHours = i
		return nil
	default:
		panic(fmt.Sprintf("unknown meta: %#v", meta))
	}
}

func (s *speccer) allMetas() string {
	var buffer bytes.Buffer
	for meta := range metas {
		fmt.Fprintf(&buffer, "%s: %s\n", meta, s.getMeta(meta))
	}
	return buffer.String()
}
