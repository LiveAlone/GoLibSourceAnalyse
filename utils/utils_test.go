package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type WorkResourceItem struct {
	WorksUrl        string `json:"worksUrl"`
	WorksPreviewUrl string `json:"worksPreviewUrl"`
	WorksOid        string `json:"worksOid"`
	Extend          string `json:"extend"`
}

type WorkResourceInfo struct {
	Source       int                `json:"source"`
	WorksFiles   []WorkResourceItem `json:"worksFiles"`
	WorksDesigns []interface{}      `json:"worksDesigns"`
}

func TestNone(t *testing.T) {
	item := &WorkResourceInfo{
		Source: 1,
		WorksFiles: []WorkResourceItem{
			{
				WorksUrl:        "123",
				WorksPreviewUrl: "44",
				WorksOid:        "111",
				Extend:          "ext",
			},
		},
		WorksDesigns: make([]interface{}, 0),
	}
	b, err := json.Marshal(item)
	fmt.Println(string(b), err)
}
