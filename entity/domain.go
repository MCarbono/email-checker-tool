package entity

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type Domain struct {
	Domain      string
	MX          bool
	SPF         bool
	DMARC       bool
	SpfRecord   string
	DmarcRecord string
}

func (d Domain) String() string {
	return fmt.Sprintf(
		"\nDomain: %v\nHas MX: %v\nHas SPF: %v\nSPF Record: %v\nHas DMARC %v\nDMARC Record: %v\n",
		d.Domain, d.MX, d.SPF, d.SpfRecord, d.DMARC, d.DmarcRecord)
}

func NewDomain(domain string) (*Domain, error) {
	var totalErr []error

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		totalErr = append(totalErr, err)
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		totalErr = append(totalErr, err)
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		totalErr = append(totalErr, err)
	}

	if len(totalErr) == 3 {
		return nil, errors.New("Invalid domain")
	}

	d := &Domain{
		Domain: domain,
	}

	if len(mxRecords) > 0 {
		d.MX = true
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			d.SPF = true
			d.SpfRecord = record
			break
		}
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			d.DMARC = true
			d.DmarcRecord = record
			break
		}
	}

	return d, nil
}
