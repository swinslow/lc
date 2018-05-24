// #!/usr/local/bin/python
// # -*- coding: utf-8 -*-
// # SPDX-License-Identifier: MIT

// import os
// import json
// import re
// import sys
// from os import listdir
// from os.path import isfile, join

// '''
// Parses based on the SPDX uploaded licenses in github https://github.com/spdx/license-list-data

// Running this takes a while so be prepared to wait a long time while it churns away
// '''

// def clean_text(text):
//     text = text.lower()
//     text = re.sub('[^a-zA-Z0-9 ]', ' ', text)
//     text = re.sub('\s+', ' ', text)
//     return text

// def find_ngrams(input_list, n):
//     return zip(*[input_list[i:] for i in range(n)])

// def build_database():
//     license_dir = '../examples/licenses/'

//     onlyfiles = [f for f in listdir(license_dir) if isfile(join(license_dir, f))]
//     licenses = []

//     for license in onlyfiles:
//         with open(join(license_dir, license), 'r') as file:
//             temp = file.read()
//             license_json = json.loads(temp)

//             ngrams = []
//             ngramrange = [3, 7, 8]

//             if license_json['licenseId'] in ['Artistic-1.0', 'BSD-3-Clause']:
//                 ngramrange = range(2, 35)

//             cleaned = clean_text(license_json['licenseText']).split()
//             if 'standardLicenseTemplate' in license_json:
//                 cleaned += clean_text(license_json['standardLicenseTemplate']).split()

//             for x in ngramrange:
//                 ngrams = ngrams + find_ngrams(cleaned, x)
//             license_json['ngrams'] = ngrams

//             licenses.append(license_json)

//     fair_source = {
//         'name': 'Fair Source License v0.9',
//         'licenseId': 'Fair-Source-0.9',
//         'licenseText': 'Fair Source License, version 0.9 Copyright (C) [year] [copyright owner] Licensor: [legal name of licensor] Software: [name software and version if applicable] Use Limitation: [number] users License Grant. Licensor hereby grants to each recipient of the Software (\"you\") a non-exclusive, non-transferable, royalty-free and fully-paid-up license, under all of the Licensors copyright and patent rights, to use, copy, distribute, prepare derivative works of, publicly perform and display the Software, subject to the Use Limitation and the conditions set forth below. Use Limitation. The license granted above allows use by up to the number of users per entity set forth above (the \"Use Limitation\"). For determining the number of users, \"you\" includes all affiliates, meaning legal entities controlling, controlled by, or under common control with you. If you exceed the Use Limitation, your use is subject to payment of Licensors then-current list price for licenses. Conditions. Redistribution in source code or other forms must include a copy of this license document to be provided in a reasonable manner. Any redistribution of the Software is only allowed subject to this license. Trademarks. This license does not grant you any right in the trademarks, service marks, brand names or logos of Licensor. DISCLAIMER. THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OR CONDITION, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. LICENSORS HEREBY DISCLAIM ALL LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE. Termination. If you violate the terms of this license, your rights will terminate automatically and will not be reinstated without the prior written consent of Licensor. Any such termination will not affect the right of others who may have received copies of the Software from you.',
//     }
//     for x in [3, 7, 8]:
//         ngrams = ngrams + find_ngrams(fair_source['licenseText'].split(), x)
//         fair_source['ngrams'] = ngrams

//     licenses.append(fair_source)
//     return licenses

// if __name__ == '__main__':
//     print 'Building database...'
//     licenses = build_database()

//     print 'Processing licenses...'
//     for license in licenses:
//         print 'PROCESSING', license['licenseId'], len(license['ngrams'])
//         matches = []

//         for ngram in license['ngrams']:
//             find = ' '.join(ngram)
//             ismatch = True

//             filtered = [x for x in licenses if x['licenseId'] != license['licenseId']]
//             for lic in filtered:
//                 # if find in lic['licenseText']:
//                 #     print find, 'FOUND'
//                 #     ismatch = False
//                 #     break
//                 for ngram2 in lic['ngrams']:
//                     if find == ' '.join(ngram2):
//                         ismatch = False
//                         break

//             if ismatch:
//                 matches.append(find)

//             if len(matches) == 50:
//                 break

//         if len(matches) == 0:
//             print '>>>>', license['licenseId'], len(matches)
//         else:
//             print license['licenseId'], len(matches)

//         license['keywords'] = matches

//     licenses = [{
//         'licenseText': x['licenseText'],
//         'name': x['name'],
//         'licenseId': x['licenseId'],
//         'keywords': x['keywords'][:50]
//     } for x in licenses]

//     with open('database_keywords.json', 'w') as myfile:
//         myfile.write(json.dumps(licenses))

package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"fmt"
)

type License struct {
	LicenseText             string   `json:"licenseText"`
	StandardLicenseTemplate string   `json:"standardLicenseTemplate"`
	Name                    string   `json:"name"`
	LicenseId               string   `json:"licenseId"`
	Keywords                []string `json:"keywords"`
}

var spdxLicenceRegex = regexp.MustCompile(`SPDX-License-Identifier:\s+(.*)[ |\n|\r\n]*?`)
var alphaNumericRegex = regexp.MustCompile("[^a-zA-Z0-9 ]")
var multipleSpacesRegex = regexp.MustCompile("\\s+")

func cleanText(content string) string {
	content = strings.ToLower(content)

	content = alphaNumericRegex.ReplaceAllString(content, " ")
	content = multipleSpacesRegex.ReplaceAllString(content, " ")

	return content
}

func findNgrams(list []string, size int) [][]string {

	for i := 0; i < len(list); i++ {
		var ngram []string

		for j := i; j < i+size; j++ {
			fmt.Println(i, j)
			ngram = append(ngram, list[i+j])
		}

		fmt.Println(ngram)
	}

	return nil
}

func main() {
	files, _ := ioutil.ReadDir("./licenses/")

	var licenses []License

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".json") {
			bytes, _ := ioutil.ReadFile(filepath.Join("./licenses/", f.Name()))

			var license License
			json.Unmarshal(bytes, &license)

			licenses = append(licenses, license)
		}
	}

	findNgrams(strings.Split("Lorem ipsum dolor sit amet consetetur sadipscing elitr", " "), 3)


	//for _, license := range licenses {
	//	split := strings.Split(cleanText(license.LicenseText), " ")
	//	//findNgrams(split, 7)
	//	//findNgrams(split, 8)
	//}

}
