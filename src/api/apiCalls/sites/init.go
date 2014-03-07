/*
 * Copyright (c) 2013 GRNET S.A., SRCE, IN2P3 CNRS Computing Centre
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the
 * License. You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an "AS
 * IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language
 * governing permissions and limitations under the License.
 *
 * The views and conclusions contained in the software and
 * documentation are those of the authors and should not be
 * interpreted as representing official policies, either expressed
 * or implied, of either GRNET S.A., SRCE or IN2P3 CNRS Computing
 * Centre
 *
 * The work represented by this source file is partially funded by
 * the EGI-InSPIRE project through the European Commission's 7th
 * Framework Programme (contract # INFSO-RI-261323)
 */

package sites

var customForm []string

type ApiSiteAvailabilityInProfileInput struct {
	// mandatory values
	start_time         string   // UTC time in W3C format
	end_time           string   // UTC time in W3C format
	profile_name       []string // may appear more than once. (eg: CMS_CRITICAL)
	group_type         []string // may appear more than once. (eg: CMS_Site)
	availabilityperiod string   // availability period; possible values: `HOURLY`, `DAILY`, `WEEKLY`, `MONTHLY`
	// optional values
	output     string   // default XML; possible values are: XML, JSON
	namespace  []string // profile namespace; may appear more than once. (eg: ch.cern.sam)
	group_name []string // site name; may appear more than once
}

type ApiSiteAvailabilityInProfileOutput struct {
	SiteScope     string  "ss"
	Scope         string  "sc"
	Date          string  "dt"
	Namespace     string  "ns"
	Profile       string  "p"
	Production    string  "pr"
	Monitored     string  "m"
	Ngi           string  "n"
	Site          string  "s"
	Infastructure string  "i"
	CertStatus    string  "cs"
	Availability  float64 "a"
	Reliability   float64 "r"
}

func init() {
	customForm = []string{"20060102", "2006-01-02T15:04:05Z"} //{"Format that is returned by the database" , "Format that will be used in the generated report"}
}
