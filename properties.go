package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

var properties = map[string]bool{
	"COMPANY":      true,
	"PROJECT":      true,
	"URL":          true,
	"LANGUAGE":     true,
	"DATEFORMAT":   true,
	"REQUESTEDBY":  true,
	"RELATED":      true,
	"TRANSLATIONS": true,
	"SUPERSEDEDBY": true,
	"RESOURCES":    true,
	"PARENT":       true,
	"PERSONS":      true,
	"APPROVED":     true,
}

var unsettableProperties = map[string]bool{
	"REQUESTEDBY":  true,
	"RELATED":      true,
	"TRANSLATIONS": true,
	"PERSONS":      true,
	"SUPERSEDEDBY": true,
	"RESOURCES":    true,
	"PARENT":       true,
}

func (s *speccer) setProperty(property string, val string) error {
	switch property {
	case "COMPANY":
		s.Spec.INFO.Company = val
		return nil
	case "PROJECT":
		s.Spec.INFO.Project = val
		return nil
	// TODO maybe check if it is a valid URI
	case "URL":
		s.Spec.INFO.URL = val
		return nil
	// TODO definitely check, if it is a valid language
	case "LANGUAGE":
		s.Spec.INFO.Language = val
		return nil
	// TODO definitely check, if it is a valid dateformat
	case "DATEFORMAT":
		s.Spec.INFO.DateFormat = val
		return nil
	case "REQUESTEDBY":
		return json.Unmarshal([]byte(val), &s.Spec.INFO.RequestedBy)
	case "RELATED":
		return json.Unmarshal([]byte(val), &s.Spec.INFO.Related)
	case "TRANSLATIONS":
		return json.Unmarshal([]byte(val), &s.Spec.INFO.Translations)
	case "PERSONS":
		return json.Unmarshal([]byte(val), &s.Spec.INFO.Persons)
	case "SUPERSEDEDBY":
		return json.Unmarshal([]byte(val), &s.Spec.INFO.SupersededBy)
	case "RESOURCES":
		return json.Unmarshal([]byte(val), &s.Spec.INFO.Resources)
	// TODO definitely check, if it is a valid url
	case "PARENT":
		s.Spec.INFO.Parent = val
		return nil
	// TODO definitely check, if it is approvable
	case "APPROVED":
		b, err := strconv.ParseBool(val)
		if err != nil {
			return err
		}
		s.Spec.INFO.Approved = b
		return nil
	default:
		panic("unknown property " + property)
	}
}

func (s *speccer) unsetProperty(property string) error {
	switch property {
	case "REQUESTEDBY":
		s.Spec.INFO.RequestedBy = []string{}
		return nil
	case "RELATED":
		s.Spec.INFO.Related = map[string]string{}
		return nil
	case "TRANSLATIONS":
		s.Spec.INFO.Translations = map[string]string{}
		return nil
	case "PERSONS":
		s.Spec.INFO.Persons = map[string]string{}
		return nil
	case "SUPERSEDEDBY":
		s.Spec.INFO.SupersededBy = map[string]string{}
		return nil
	case "RESOURCES":
		s.Spec.INFO.Resources = map[string]string{}
		return nil
	// TODO definitely check, if it is a valid url
	case "PARENT":
		s.Spec.INFO.Parent = ""
		return nil
	default:
		panic("unknown property " + property)
	}
}

func (s *speccer) getProperty(property string) string {
	switch property {
	case "COMPANY":
		return s.Spec.INFO.Company
	case "PROJECT":
		return s.Spec.INFO.Project
	case "URL":
		return s.Spec.INFO.URL
	case "LANGUAGE":
		return s.Spec.INFO.Language
	case "DATEFORMAT":
		return s.Spec.INFO.DateFormat
	case "REQUESTEDBY":
		data, _ := json.Marshal(s.Spec.INFO.RequestedBy)
		return fmt.Sprintf("%s", data)
	case "RELATED":
		data, _ := json.Marshal(s.Spec.INFO.Related)
		return fmt.Sprintf("%s", data)
	case "TRANSLATIONS":
		data, _ := json.Marshal(s.Spec.INFO.Translations)
		return fmt.Sprintf("%s", data)
	case "PERSONS":
		data, _ := json.Marshal(s.Spec.INFO.Persons)
		return fmt.Sprintf("%s", data)
	case "SUPERSEDEDBY":
		data, _ := json.Marshal(s.Spec.INFO.SupersededBy)
		return fmt.Sprintf("%s", data)
	case "RESOURCES":
		data, _ := json.Marshal(s.Spec.INFO.Resources)
		return fmt.Sprintf("%s", data)
	case "PARENT":
		return s.Spec.INFO.Parent
	case "APPROVED":
		return fmt.Sprintf("%v", s.Spec.INFO.Approved)
	default:
		panic("unknown property: " + property)

	}
}

func (s *speccer) allProperties() string {
	var buffer bytes.Buffer
	for property := range properties {
		fmt.Fprintf(&buffer, "%s: %s\n", property, s.getProperty(property))
	}
	return buffer.String()
}
