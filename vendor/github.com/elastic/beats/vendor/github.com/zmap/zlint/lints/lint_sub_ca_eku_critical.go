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
BRs: 7.1.2.2g extkeyUsage (optional)
For Subordinate CA Certificates to be Technically constrained in line with section 7.1.5, then either the value
id‐kp‐serverAuth [RFC5280] or id‐kp‐clientAuth [RFC5280] or both values MUST be present**.
Other values MAY be present.
If present, this extension SHOULD be marked non‐critical.
************************************************/

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCAEKUCrit struct{}

func (l *subCAEKUCrit) Initialize() error {
	return nil
}

func (l *subCAEKUCrit) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.EkuSynOid)
}

func (l *subCAEKUCrit) Execute(c *x509.Certificate) *LintResult {
	if e := util.GetExtFromCert(c, util.EkuSynOid); e.Critical {
		return &LintResult{Status: Warn}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_ca_eku_critical",
		Description:   "Subordinate CA certificate extkeyUsage extension should be marked non-critical if present",
		Citation:      "BRs: 7.1.2.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABV116Date,
		Lint:          &subCAEKUCrit{},
	})
}
