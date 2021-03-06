/*
 * Copyright (c) 2014 GRNET S.A., SRCE, IN2P3 CNRS Computing Centre
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

package statusSites

import "encoding/xml"

type StatusSitesInput struct {
	start_time string // UTC time in W3C format
	end_time   string
	vo         string
	profile    string
	group_type string
	group      string
}

type StatusSitesOutput struct {
	Timestamp string `bson:"ts"`
	Roc       string `bson:"roc"`
	Site      string `bson:"site"`
	Status    string `bson:"s"`
	Time_int  int    `bson:"ti"`
	P_status  string `bson:"ps"`
	Profile   string `bson:"p"`
}

type ReadRoot struct {
	XMLName xml.Name `xml:"root"`
	Profile *Profile
}

type Profile struct {
	XMLName xml.Name `xml:"profile"`
	Name    string   `xml:"name,attr"`
	Groups  []*Group
}

type Group struct {
	XMLName xml.Name `xml:"group"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
	Groups  []*Group
	Sites   []*Site
}

type Site struct {
	XMLName  xml.Name `xml:"endpoint"`
	Name     string   `xml:"name,attr"`
	Timeline []*Status
}

type Status struct {
	XMLName   xml.Name `xml:"status"`
	Timestamp string   `xml:"timestamp,attr"`
	Status    string   `xml:"status,attr"`
}
