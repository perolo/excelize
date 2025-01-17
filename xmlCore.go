// Copyright 2016 - 2022 The myexcelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package myexcelize providing a set of functions that allow you to write to and
// read from XLAM / XLSM / XLSX / XLTM / XLTX files. Supports reading and
// writing spreadsheet documents generated by Microsoft Excel™ 2007 and later.
// Supports complex components by high compatibility, and provided streaming
// API for generating or reading data from a worksheet with huge amounts of
// data. This library needs Go version 1.15 or later.

package myexcelize

import "encoding/xml"

// DocProperties directly maps the document core properties.
type DocProperties struct {
	Category       string
	ContentStatus  string
	Created        string
	Creator        string
	Description    string
	Identifier     string
	Keywords       string
	LastModifiedBy string
	Modified       string
	Revision       string
	Subject        string
	Title          string
	Language       string
	Version        string
}

// decodeCoreProperties directly maps the root element for a part of this
// content type shall coreProperties. In order to solve the problem that the
// label structure is changed after serialization and deserialization, two
// different structures are defined. decodeCoreProperties just for
// deserialization.
type decodeCoreProperties struct {
	XMLName        xml.Name `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties coreProperties"`
	Title          string   `xml:"http://purl.org/dc/elements/1.1/ title,omitempty"`
	Subject        string   `xml:"http://purl.org/dc/elements/1.1/ subject,omitempty"`
	Creator        string   `xml:"http://purl.org/dc/elements/1.1/ creator"`
	Keywords       string   `xml:"keywords,omitempty"`
	Description    string   `xml:"http://purl.org/dc/elements/1.1/ description,omitempty"`
	LastModifiedBy string   `xml:"lastModifiedBy"`
	Language       string   `xml:"http://purl.org/dc/elements/1.1/ language,omitempty"`
	Identifier     string   `xml:"http://purl.org/dc/elements/1.1/ identifier,omitempty"`
	Revision       string   `xml:"revision,omitempty"`
	Created        struct {
		Text string `xml:",chardata"`
		Type string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	} `xml:"http://purl.org/dc/terms/ created"`
	Modified struct {
		Text string `xml:",chardata"`
		Type string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	} `xml:"http://purl.org/dc/terms/ modified"`
	ContentStatus string `xml:"contentStatus,omitempty"`
	Category      string `xml:"category,omitempty"`
	Version       string `xml:"version,omitempty"`
}

// xlsxCoreProperties directly maps the root element for a part of this
// content type shall coreProperties.
type xlsxCoreProperties struct {
	XMLName        xml.Name `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties coreProperties"`
	Dc             string   `xml:"xmlns:dc,attr"`
	Dcterms        string   `xml:"xmlns:dcterms,attr"`
	Dcmitype       string   `xml:"xmlns:dcmitype,attr"`
	XSI            string   `xml:"xmlns:xsi,attr"`
	Title          string   `xml:"dc:title,omitempty"`
	Subject        string   `xml:"dc:subject,omitempty"`
	Creator        string   `xml:"dc:creator"`
	Keywords       string   `xml:"keywords,omitempty"`
	Description    string   `xml:"dc:description,omitempty"`
	LastModifiedBy string   `xml:"lastModifiedBy"`
	Language       string   `xml:"dc:language,omitempty"`
	Identifier     string   `xml:"dc:identifier,omitempty"`
	Revision       string   `xml:"revision,omitempty"`
	Created        struct {
		Text string `xml:",chardata"`
		Type string `xml:"xsi:type,attr"`
	} `xml:"dcterms:created"`
	Modified struct {
		Text string `xml:",chardata"`
		Type string `xml:"xsi:type,attr"`
	} `xml:"dcterms:modified"`
	ContentStatus string `xml:"contentStatus,omitempty"`
	Category      string `xml:"category,omitempty"`
	Version       string `xml:"version,omitempty"`
}
