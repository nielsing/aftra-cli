/*
Web backend for Vikingr

The main api for Vikingr

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// Sourceinformation struct for Sourceinformation
type Sourceinformation struct {
	AWSIntegration *AWSIntegration
	AmassEnum *AmassEnum
	AmassIntel *AmassIntel
	BitBucketIntegration *BitBucketIntegration
	DNSScan *DNSScan
	DNSScanReverse *DNSScanReverse
	GitLabIntegration *GitLabIntegration
	GitRepo *GitRepo
	GithubIntegration *GithubIntegration
	GoogleIntegration *GoogleIntegration
	MicrosoftIntegration *MicrosoftIntegration
	PassDB *PassDB
	SubdomainEnumerationScan *SubdomainEnumerationScan
	UserEntry *UserEntry
	WebCrawl *WebCrawl
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Sourceinformation) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AWSIntegration
	err = json.Unmarshal(data, &dst.AWSIntegration);
	if err == nil {
		jsonAWSIntegration, _ := json.Marshal(dst.AWSIntegration)
		if string(jsonAWSIntegration) == "{}" { // empty struct
			dst.AWSIntegration = nil
		} else {
			return nil // data stored in dst.AWSIntegration, return on the first match
		}
	} else {
		dst.AWSIntegration = nil
	}

	// try to unmarshal JSON data into AmassEnum
	err = json.Unmarshal(data, &dst.AmassEnum);
	if err == nil {
		jsonAmassEnum, _ := json.Marshal(dst.AmassEnum)
		if string(jsonAmassEnum) == "{}" { // empty struct
			dst.AmassEnum = nil
		} else {
			return nil // data stored in dst.AmassEnum, return on the first match
		}
	} else {
		dst.AmassEnum = nil
	}

	// try to unmarshal JSON data into AmassIntel
	err = json.Unmarshal(data, &dst.AmassIntel);
	if err == nil {
		jsonAmassIntel, _ := json.Marshal(dst.AmassIntel)
		if string(jsonAmassIntel) == "{}" { // empty struct
			dst.AmassIntel = nil
		} else {
			return nil // data stored in dst.AmassIntel, return on the first match
		}
	} else {
		dst.AmassIntel = nil
	}

	// try to unmarshal JSON data into BitBucketIntegration
	err = json.Unmarshal(data, &dst.BitBucketIntegration);
	if err == nil {
		jsonBitBucketIntegration, _ := json.Marshal(dst.BitBucketIntegration)
		if string(jsonBitBucketIntegration) == "{}" { // empty struct
			dst.BitBucketIntegration = nil
		} else {
			return nil // data stored in dst.BitBucketIntegration, return on the first match
		}
	} else {
		dst.BitBucketIntegration = nil
	}

	// try to unmarshal JSON data into DNSScan
	err = json.Unmarshal(data, &dst.DNSScan);
	if err == nil {
		jsonDNSScan, _ := json.Marshal(dst.DNSScan)
		if string(jsonDNSScan) == "{}" { // empty struct
			dst.DNSScan = nil
		} else {
			return nil // data stored in dst.DNSScan, return on the first match
		}
	} else {
		dst.DNSScan = nil
	}

	// try to unmarshal JSON data into DNSScanReverse
	err = json.Unmarshal(data, &dst.DNSScanReverse);
	if err == nil {
		jsonDNSScanReverse, _ := json.Marshal(dst.DNSScanReverse)
		if string(jsonDNSScanReverse) == "{}" { // empty struct
			dst.DNSScanReverse = nil
		} else {
			return nil // data stored in dst.DNSScanReverse, return on the first match
		}
	} else {
		dst.DNSScanReverse = nil
	}

	// try to unmarshal JSON data into GitLabIntegration
	err = json.Unmarshal(data, &dst.GitLabIntegration);
	if err == nil {
		jsonGitLabIntegration, _ := json.Marshal(dst.GitLabIntegration)
		if string(jsonGitLabIntegration) == "{}" { // empty struct
			dst.GitLabIntegration = nil
		} else {
			return nil // data stored in dst.GitLabIntegration, return on the first match
		}
	} else {
		dst.GitLabIntegration = nil
	}

	// try to unmarshal JSON data into GitRepo
	err = json.Unmarshal(data, &dst.GitRepo);
	if err == nil {
		jsonGitRepo, _ := json.Marshal(dst.GitRepo)
		if string(jsonGitRepo) == "{}" { // empty struct
			dst.GitRepo = nil
		} else {
			return nil // data stored in dst.GitRepo, return on the first match
		}
	} else {
		dst.GitRepo = nil
	}

	// try to unmarshal JSON data into GithubIntegration
	err = json.Unmarshal(data, &dst.GithubIntegration);
	if err == nil {
		jsonGithubIntegration, _ := json.Marshal(dst.GithubIntegration)
		if string(jsonGithubIntegration) == "{}" { // empty struct
			dst.GithubIntegration = nil
		} else {
			return nil // data stored in dst.GithubIntegration, return on the first match
		}
	} else {
		dst.GithubIntegration = nil
	}

	// try to unmarshal JSON data into GoogleIntegration
	err = json.Unmarshal(data, &dst.GoogleIntegration);
	if err == nil {
		jsonGoogleIntegration, _ := json.Marshal(dst.GoogleIntegration)
		if string(jsonGoogleIntegration) == "{}" { // empty struct
			dst.GoogleIntegration = nil
		} else {
			return nil // data stored in dst.GoogleIntegration, return on the first match
		}
	} else {
		dst.GoogleIntegration = nil
	}

	// try to unmarshal JSON data into MicrosoftIntegration
	err = json.Unmarshal(data, &dst.MicrosoftIntegration);
	if err == nil {
		jsonMicrosoftIntegration, _ := json.Marshal(dst.MicrosoftIntegration)
		if string(jsonMicrosoftIntegration) == "{}" { // empty struct
			dst.MicrosoftIntegration = nil
		} else {
			return nil // data stored in dst.MicrosoftIntegration, return on the first match
		}
	} else {
		dst.MicrosoftIntegration = nil
	}

	// try to unmarshal JSON data into PassDB
	err = json.Unmarshal(data, &dst.PassDB);
	if err == nil {
		jsonPassDB, _ := json.Marshal(dst.PassDB)
		if string(jsonPassDB) == "{}" { // empty struct
			dst.PassDB = nil
		} else {
			return nil // data stored in dst.PassDB, return on the first match
		}
	} else {
		dst.PassDB = nil
	}

	// try to unmarshal JSON data into SubdomainEnumerationScan
	err = json.Unmarshal(data, &dst.SubdomainEnumerationScan);
	if err == nil {
		jsonSubdomainEnumerationScan, _ := json.Marshal(dst.SubdomainEnumerationScan)
		if string(jsonSubdomainEnumerationScan) == "{}" { // empty struct
			dst.SubdomainEnumerationScan = nil
		} else {
			return nil // data stored in dst.SubdomainEnumerationScan, return on the first match
		}
	} else {
		dst.SubdomainEnumerationScan = nil
	}

	// try to unmarshal JSON data into UserEntry
	err = json.Unmarshal(data, &dst.UserEntry);
	if err == nil {
		jsonUserEntry, _ := json.Marshal(dst.UserEntry)
		if string(jsonUserEntry) == "{}" { // empty struct
			dst.UserEntry = nil
		} else {
			return nil // data stored in dst.UserEntry, return on the first match
		}
	} else {
		dst.UserEntry = nil
	}

	// try to unmarshal JSON data into WebCrawl
	err = json.Unmarshal(data, &dst.WebCrawl);
	if err == nil {
		jsonWebCrawl, _ := json.Marshal(dst.WebCrawl)
		if string(jsonWebCrawl) == "{}" { // empty struct
			dst.WebCrawl = nil
		} else {
			return nil // data stored in dst.WebCrawl, return on the first match
		}
	} else {
		dst.WebCrawl = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(Sourceinformation)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Sourceinformation) MarshalJSON() ([]byte, error) {
	if src.AWSIntegration != nil {
		return json.Marshal(&src.AWSIntegration)
	}

	if src.AmassEnum != nil {
		return json.Marshal(&src.AmassEnum)
	}

	if src.AmassIntel != nil {
		return json.Marshal(&src.AmassIntel)
	}

	if src.BitBucketIntegration != nil {
		return json.Marshal(&src.BitBucketIntegration)
	}

	if src.DNSScan != nil {
		return json.Marshal(&src.DNSScan)
	}

	if src.DNSScanReverse != nil {
		return json.Marshal(&src.DNSScanReverse)
	}

	if src.GitLabIntegration != nil {
		return json.Marshal(&src.GitLabIntegration)
	}

	if src.GitRepo != nil {
		return json.Marshal(&src.GitRepo)
	}

	if src.GithubIntegration != nil {
		return json.Marshal(&src.GithubIntegration)
	}

	if src.GoogleIntegration != nil {
		return json.Marshal(&src.GoogleIntegration)
	}

	if src.MicrosoftIntegration != nil {
		return json.Marshal(&src.MicrosoftIntegration)
	}

	if src.PassDB != nil {
		return json.Marshal(&src.PassDB)
	}

	if src.SubdomainEnumerationScan != nil {
		return json.Marshal(&src.SubdomainEnumerationScan)
	}

	if src.UserEntry != nil {
		return json.Marshal(&src.UserEntry)
	}

	if src.WebCrawl != nil {
		return json.Marshal(&src.WebCrawl)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSourceinformation struct {
	value *Sourceinformation
	isSet bool
}

func (v NullableSourceinformation) Get() *Sourceinformation {
	return v.value
}

func (v *NullableSourceinformation) Set(val *Sourceinformation) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceinformation) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceinformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceinformation(val *Sourceinformation) *NullableSourceinformation {
	return &NullableSourceinformation{value: val, isSet: true}
}

func (v NullableSourceinformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceinformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


