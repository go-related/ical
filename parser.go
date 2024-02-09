package ical

import (
	"bufio"
	"context"
	"io"
	"time"
)

type Parser struct {
	Start   *time.Time
	End     *time.Time
	scanner *bufio.Scanner
	events  []Event
	current *Event
}

type Event struct {
	Uid              string
	Summary          string
	Description      string
	Start            *time.Time
	End              *time.Time
	Created          *time.Time
	LastModified     *time.Time
	IsRecurring      bool
	RecurrenceID     *time.Time
	RecurrenceRule   map[string]string
	ExcludeDates     []time.Time
	CustomAttributes map[string]string
}

func NewParser(reader io.Reader) *Parser {
	return &Parser{
		scanner: bufio.NewScanner(reader),
		events:  make([]Event, 0),
	}
}

func (p *Parser) Parse(ctx context.Context) error {
	if p.Start == nil {
		p.Start = datePtr(time.Now().UTC())
	}
	if p.End == nil {
		p.End = datePtr(time.Now().UTC().AddDate(0, 0, DefaultTimeRangeInDays))
	}

	return nil
}
