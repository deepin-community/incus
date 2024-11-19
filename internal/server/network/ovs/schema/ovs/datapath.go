// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package ovsmodel

const DatapathTable = "Datapath"

// Datapath defines an object in Datapath table
type Datapath struct {
	UUID            string            `ovsdb:"_uuid"`
	Capabilities    map[string]string `ovsdb:"capabilities"`
	CTZones         map[int]string    `ovsdb:"ct_zones"`
	DatapathVersion string            `ovsdb:"datapath_version"`
	ExternalIDs     map[string]string `ovsdb:"external_ids"`
}
