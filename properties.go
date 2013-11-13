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
	"TITLE":        true,
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
		s.Spec.Company = val
		return nil
	case "PROJECT":
		s.Spec.Project = val
		return nil
	case "TITLE":
		s.Spec.Title = val
		return nil
	// TODO maybe check if it is a valid URI
	case "URL":
		s.Spec.URL = val
		return nil
	// TODO definitely check, if it is a valid language
	case "LANGUAGE":
		s.Spec.Language = val
		return nil
	// TODO definitely check, if it is a valid dateformat
	case "DATEFORMAT":
		s.Spec.DateFormat = val
		return nil
	case "REQUESTEDBY":
		return json.Unmarshal([]byte(val), &s.Spec.RequestedBy)
	case "RELATED":
		return json.Unmarshal([]byte(val), &s.Spec.Related)
	case "TRANSLATIONS":
		return json.Unmarshal([]byte(val), &s.Spec.Translations)
	case "PERSONS":
		return json.Unmarshal([]byte(val), &s.Spec.Persons)
	case "SUPERSEDEDBY":
		return json.Unmarshal([]byte(val), &s.Spec.SupersededBy)
	case "RESOURCES":
		return json.Unmarshal([]byte(val), &s.Spec.Resources)
	// TODO definitely check, if it is a valid url
	case "PARENT":
		s.Spec.Parent = val
		return nil
	// TODO definitely check, if it is approvable
	case "APPROVED":
		b, err := strconv.ParseBool(val)
		if err != nil {
			return err
		}
		s.Spec.Approved = b
		return nil
	default:
		panic("unknown property " + property)
	}
}

func (s *speccer) unsetProperty(property string) error {
	switch property {
	case "REQUESTEDBY":
		s.Spec.RequestedBy = []string{}
		return nil
	case "RELATED":
		s.Spec.Related = map[string]string{}
		return nil
	case "TRANSLATIONS":
		s.Spec.Translations = map[string]string{}
		return nil
	case "PERSONS":
		s.Spec.Persons = map[string]string{}
		return nil
	case "SUPERSEDEDBY":
		s.Spec.SupersededBy = map[string]string{}
		return nil
	case "RESOURCES":
		s.Spec.Resources = map[string]string{}
		return nil
	// TODO definitely check, if it is a valid url
	case "PARENT":
		s.Spec.Parent = ""
		return nil
	default:
		panic("unknown property " + property)
	}
}

func (s *speccer) getProperty(property string) string {
	switch property {
	case "COMPANY":
		return s.Spec.Company
	case "PROJECT":
		return s.Spec.Project
	case "TITLE":
		return s.Spec.Title
	case "URL":
		return s.Spec.URL
	case "LANGUAGE":
		return s.Spec.Language
	case "DATEFORMAT":
		return s.Spec.DateFormat
	case "REQUESTEDBY":
		data, _ := json.Marshal(s.Spec.RequestedBy)
		return fmt.Sprintf("%s", data)
	case "RELATED":
		data, _ := json.Marshal(s.Spec.Related)
		return fmt.Sprintf("%s", data)
	case "TRANSLATIONS":
		data, _ := json.Marshal(s.Spec.Translations)
		return fmt.Sprintf("%s", data)
	case "PERSONS":
		data, _ := json.Marshal(s.Spec.Persons)
		return fmt.Sprintf("%s", data)
	case "SUPERSEDEDBY":
		data, _ := json.Marshal(s.Spec.SupersededBy)
		return fmt.Sprintf("%s", data)
	case "RESOURCES":
		data, _ := json.Marshal(s.Spec.Resources)
		return fmt.Sprintf("%s", data)
	case "PARENT":
		return s.Spec.Parent
	case "APPROVED":
		return fmt.Sprintf("%v", s.Spec.Approved)
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
