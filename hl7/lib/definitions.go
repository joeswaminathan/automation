package hl7

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type TableEntry struct {
    Value string        `json:"value"`
    Description string  `json:"description,omitempty"`
    Comment string      `json:"comment,omitempty"`
}

type Table []*TableEntry

type DataTypePart struct {
    Mandatory bool          `json:"mandatory,omitempty"`
    Name string             `json:"name"`
    Repeats bool            `json:"repeats,omitempty"`
    Type string             `json:"type"`
    Length int              `json:"length"`
    TableId string          `json:"table,omitempty"`
    Composites *DataType    `json:"composites,omitempty"`
}

type DataTypeVersion struct {
    AppliesTo string        `json:"appliesTo"`
    Length  int             `json:"length"`
    Parts []*DataTypePart   `json:"parts"`
}

type DataType struct {
    Separator string                        `json:"separator"`
    PartId string                           `json:"partId"`
    Versions map[string]*DataTypeVersion    `json:"versions"`
}

type Definitions struct {
    Tables map[string]*Table        `json:"tables,omitempty"`
    DataTypes map[string]*DataType  `json:"dataTypes,omitempty"`
}

func (def *Definitions) Load( file string ) error {
    data, err := ioutil.ReadFile( file )
    if err != nil {
        fmt.Printf("Error loading file (%s), error (%s)\n", file, err.Error() )
        return err
    }

    if err := json.Unmarshal(data, def); err != nil {
        fmt.Printf("Error in Unmarshal, error (%s)\n", err.Error() )
        return err
    }

    /*
    dataNew, err := json.MarshalIndent(def, "", "  ")
    fmt.Println(string(dataNew))
    */
    return nil
}

