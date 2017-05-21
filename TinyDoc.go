package main

import (
	"fmt"
	"net/http"
)

// cell number
var apiCellNum = 0
var apiCells []APIs

func init() {
	http.HandleFunc("/api/doc", APIDoc)
}

// APIDoc handle
func APIDoc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "  API Doc \n \n \n")

	apiCellNum = 0
	for _, v := range apiCells {
		showAPICell(w, v)
	}
	fmt.Fprintf(w, " . . . . . . . . . . . . . . . . . . . . . . . . . . . . .\n")
	fmt.Fprintf(w, " \n \n \n \n \n \n \n \n \n \n \n \n \n \n \n \n \n")
}

func showAPICell(w http.ResponseWriter, v APIs) {
	apiCellNum++
	v.defAll()
	fmt.Fprintf(w, " . . . . . . . . . . . . . . . . . . . . . . . . . . . . .\n")
	fmt.Fprintf(w, " | \n")
	fmt.Fprintf(w, " | %2v %-4v %v \n", apiCellNum, v.Info.id, v.Info.title)
	fmt.Fprintf(w, " | \n")
	fmt.Fprintf(w, " | %-5v %-60v \n", v.Info.mothd+":", v.Info.url)
	var tagsNum = len(v.Tags)
	if tagsNum > 0 {
		fmt.Fprintf(w, " | \n")
		fmt.Fprintf(w, " | %-13v %-10v %v \n", "Name", "Type", "Tip")
		for _, v := range v.Tags {
			fmt.Fprintf(w, " | %-13v %-10v %v \n", v.Name, v.Type, v.Tip)
		}
	}
	fmt.Fprintf(w, " | \n")
}

// APIDocAddOnes /api/doc add one cell
func APIDocAddOnes(v APIs) string {
	apiCells = append(apiCells, v)
	return v.Info.url
}

// APIDocAddHandleFun And use http.HandleFunc()
func APIDocAddHandleFun(handler func(http.ResponseWriter, *http.Request), v APIs) {
	apiCells = append(apiCells, v)
	http.HandleFunc(v.Info.url, handler)
}

// APITag -
type APITag struct {
	Name string
	Type string
	Tip  string
}

// Info -
type Info struct {
	id    string
	url   string
	title string
	mothd string
}

// APIs -
type APIs struct {
	Info Info
	Tags []APITag
}

func (v *APIs) defAll() {
	v.Info.defMothd()
	v.Info.defTitle()
}

func (info *Info) defTitle() {
	if info.title == "" {
		info.title = "Interfase title is empty"
	}
}
func (info *Info) defMothd() {
	if info.mothd == "" {
		info.mothd = "GET"
	}
	if info.mothd == "get" {
		info.mothd = "GET"
	}
	if info.mothd == "post" {
		info.mothd = "POST"
	}
	if info.mothd == "push" {
		info.mothd = "PUSH"
	}
}
