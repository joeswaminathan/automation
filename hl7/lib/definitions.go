package hl7

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
    VariedTypeField string  `json:"variedTypeField,omitempty"`
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
    initialized bool
    typeIsStruct map[string]bool
    definedTables []string
}

