package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

import (
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SubjectDNSerialNumberNotPrintableString struct{}

func (l *SubjectDNSerialNumberNotPrintableString) Initialize() error {
	return nil
}

func (l *SubjectDNSerialNumberNotPrintableString) CheckApplies(c *x509.Certificate) bool {
	return len(c.Subject.SerialNumber) > 0
}

func (l *SubjectDNSerialNumberNotPrintableString) Execute(c *x509.Certificate) *LintResult {
	rdnSequence := util.RawRDNSequence{}
	rest, err := asn1.Unmarshal(c.RawSubject, &rdnSequence)
	if err != nil {
		return &LintResult{Status: Fatal}
	}
	if len(rest) > 0 {
		return &LintResult{Status: Fatal}
	}

	for _, attrTypeAndValueSet := range rdnSequence {
		for _, attrTypeAndValue := range attrTypeAndValueSet {
			if attrTypeAndValue.Type.Equal(util.SerialOID) && attrTypeAndValue.Value.Tag != asn1.TagPrintableString {
				return &LintResult{Status: Error}
			}
		}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_dn_serial_number_not_printable_string",
		Description:   "X520 Distinguished Name SerialNumber MUST be encoded as PrintableString",
		Citation:      "RFC 5280: Appendix A",
		Source:        RFC5280,
		EffectiveDate: util.ZeroDate,
		Lint:          &SubjectDNSerialNumberNotPrintableString{},
	})
}
