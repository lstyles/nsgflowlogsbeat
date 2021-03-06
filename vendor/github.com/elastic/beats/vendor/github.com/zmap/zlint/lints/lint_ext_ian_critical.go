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

/************************************************
Issuer Alternative Name
   As with Section 4.2.1.6, this extension is used to associate Internet style identities with the certificate issuer. Issuer alternative name MUST be encoded as in 4.2.1.6.  Issuer alternative names are not processed as part of the certification path validation algorithm in Section 6. (That is, issuer alternative names are not used in name chaining and name constraints are not enforced.)
   Where present, conforming CAs SHOULD mark this extension as non-critical.
************************************************/

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type ExtIANCritical struct{}

func (l *ExtIANCritical) Initialize() error {
	return nil
}

func (l *ExtIANCritical) CheckApplies(cert *x509.Certificate) bool {
	return util.IsExtInCert(cert, util.IssuerAlternateNameOID)
}

func (l *ExtIANCritical) Execute(cert *x509.Certificate) *LintResult {
	if util.GetExtFromCert(cert, util.IssuerAlternateNameOID).Critical {
		return &LintResult{Status: Warn}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ext_ian_critical",
		Description:   "Issuer alternate name should be marked as non-critical",
		Citation:      "RFC 5280: 4.2.1.7",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &ExtIANCritical{},
	})
}
