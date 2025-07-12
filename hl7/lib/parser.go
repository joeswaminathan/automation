package hl7

import (
    "fmt"
    "unicode"
    "strings"
    "io/ioutil"
    "encoding/json"
)

func (def *Definitions) defineType(dtn string) error {
    var typeDef strings.Builder
    //fmt.Printf("--- Definining %s ---\n", dtn)
    if _, ok := def.typeIsStruct[dtn]; ok {
        //fmt.Printf("// %s is already defined\n", dtn)
        return nil
    }
    dt, ok := def.DataTypes[dtn]
    if !ok {
        return fmt.Errorf("Error data type (%s) not found\n", dtn)
    }
    //fmt.Printf("Key (%s), Value (%+v)\n", dtn, dt)
    for /*version*/_, vdt := range dt.Versions {
        typeDef.WriteString(fmt.Sprintf("type %s ", dtn))
        /*
        typeDef.WriteString(fmt.Sprintf("type %s_%s ",
                   dtn, strings.ReplaceAll(version, ".", "_")))
        */
        indent := ""
        if dt.Separator != "" {
            typeDef.WriteString(fmt.Sprintf("struct {\n"))
            indent = "    "
        }
        for idx, part := range vdt.Parts {
            repeat := ""
            if part.Repeats {
                repeat = "[]"
            }
            pointer := ""
            if unicode.IsUpper(rune(part.Type[0])) {
                if err := def.defineType(part.Type); err != nil {
                    fmt.Println(err)
                    return err
                }
                if def.typeIsStruct[part.Type] {
                    pointer = "*"
                }
            }
            name := ""
            if part.Name != "" {
                name = ( string(unicode.ToUpper(rune(part.Name[0]))) +
                          part.Name[1:] )
            }
            typeDef.WriteString(fmt.Sprintf("%s%s %s%s%s",
                       indent, name, repeat, pointer, part.Type))
            if dt.Separator != "" {
                typeDef.WriteString(fmt.Sprintf(" `json:\"%s,omitempty\"", part.Name))
                if dt.PartId == "type" {
                    typeDef.WriteString(fmt.Sprintf(" hl7:\"part=%s", part.Type))
                } else {
                    typeDef.WriteString(fmt.Sprintf(" hl7:\"part=%d", idx+1))
                }
                if part.Length != 0 {
                    typeDef.WriteString(fmt.Sprintf(" length=%d", part.Length))
                }
                if part.TableId != "" {
                    typeDef.WriteString(fmt.Sprintf(" table=%s", part.TableId))
                }
                if part.VariedTypeField != "" {
                    vtf := ( string(unicode.ToUpper(rune(part.VariedTypeField[0]))) +
                             part.VariedTypeField[1:] )
                    typeDef.WriteString(fmt.Sprintf(" variedTypeField=%s", vtf))
                }
                typeDef.WriteString(fmt.Sprintf("\"`"))

            }
            typeDef.WriteString(fmt.Sprintf("\n"))
        }
        if dt.Separator != "" {
            typeDef.WriteString(fmt.Sprintf("}"))
        }
        typeDef.WriteString(fmt.Sprintf("\n"))
    }
    def.typeIsStruct[dtn] = (dt.Separator != "")
    //fmt.Printf("--- Defined %s ---\n", dtn)
    fmt.Println(typeDef.String())
    return nil
}

func (def *Definitions) defineTypeFunctions() error {
    var instFunc strings.Builder
    var structFunc strings.Builder
    //fmt.Printf("--- Definining Table %s ---\n", tbn)
    if len(def.typeIsStruct) == 0 {
        return nil
    }
    instFunc.WriteString("func (ds *decodeState) instantiateAndParseHl7(dtn string) (parseStatus, error) {\n")
    instFunc.WriteString("    switch dtn {\n")
    //instFunc.WriteString("    dtn, _ := ds.meta[\"intefaceUnderlyingType\"]\n")
    structFunc.WriteString("func (ds *decodeState) isTypeStruct(dtn string) bool {\n")
    structFunc.WriteString("    switch dtn {\n")
    for dtn, isStruct := range def.typeIsStruct {
        instFunc.WriteString(fmt.Sprintf("    case \"%s\":\n", dtn))
        instFunc.WriteString(fmt.Sprintf("        rValue := reflect.New(reflect.TypeFor[%s]())\n", dtn))
        instFunc.WriteString(fmt.Sprintf("        ds.value.Elem().Set(rValue)\n"))
        instFunc.WriteString(fmt.Sprintf("        ds.value = rValue\n"))
        structFunc.WriteString(fmt.Sprintf("    case \"%s\":\n", dtn))
        if isStruct {
            structFunc.WriteString("        return true\n")
        } else {
            structFunc.WriteString("        return false\n")
        }
    }
    instFunc.WriteString("    default:\n")
    instFunc.WriteString(fmt.Sprintf("        err := fmt.Errorf(\"Error setting any type field: type (%%s) not found\", dtn)\n"))
    instFunc.WriteString(fmt.Sprintf("        log.Printf(\"\\t%%s\\n\", err.Error())\n"))
    instFunc.WriteString("        return INVALID, err\n")
    instFunc.WriteString("    }\n")
    instFunc.WriteString("    return ds.parseHl7()\n")
    instFunc.WriteString("}\n")
    structFunc.WriteString("    default:\n")
    structFunc.WriteString("        return false\n")
    structFunc.WriteString("    }\n")
    structFunc.WriteString("}\n")
    //fmt.Printf("--- Defined Table %s ---\n", tbn)
    fmt.Println(instFunc.String())
    fmt.Println(structFunc.String())
    return nil
}

func (def *Definitions) MakeTypes() error {
    if !def.initialized || def.DataTypes == nil {
        return nil
    }

    def.typeIsStruct = make(map[string]bool)
    for k := range def.DataTypes {
        //fmt.Printf("// %s/n", k)
        def.defineType(k)
    }

    def.defineTypeFunctions()
    return nil
}

func (def *Definitions) defineTable(tbn string) error {
    var tableDef strings.Builder
    //fmt.Printf("--- Definining Table %s ---\n", tbn)
    dt, ok := def.Tables[tbn]
    if !ok {
        return fmt.Errorf("Error data type (%s) not found\n", tbn)
    }
    //fmt.Printf("Key (%s), Value (%+v)\n", tbn, dt)
    //fmt.Printf("Key (%s), Length (%d)\n", tbn, len(*dt))
    tableDef.WriteString(fmt.Sprintf("var table_%s []string = []string{\n", tbn))
    for _, value := range *dt {
        tableDef.WriteString(fmt.Sprintf("    \"%s\",\n", value.Value))
    }
    tableDef.WriteString("}\n")
    //fmt.Printf("--- Defined Table %s ---\n", tbn)
    fmt.Println(tableDef.String())
    return nil
}

func (def *Definitions) defineTableFinder() error {
    var tableDef strings.Builder
    //fmt.Printf("--- Definining defineTableFinder %s ---\n", tbn)
    if len(def.definedTables) == 0 {
        return nil
    }
    tableDef.WriteString("func findTable(tn string) []string {\n")
    tableDef.WriteString("    switch tn {\n")
    for _, tn := range def.definedTables {
        tableDef.WriteString(fmt.Sprintf("    case \"%s\":\n", tn))
        tableDef.WriteString(fmt.Sprintf("        return table_%s\n", tn))
    }
    tableDef.WriteString("    default:\n        return nil\n")
    tableDef.WriteString("    }\n")
    tableDef.WriteString("}\n")
    //fmt.Printf("--- Defined defineTableFinder %s ---\n", tbn)
    fmt.Println(tableDef.String())
    return nil
}

func (def *Definitions) MakeTables() error {
    if !def.initialized || def.Tables == nil {
        return nil
    }

    def.definedTables = make([]string, len(def.Tables), len(def.Tables))
    idx := 0
    for k := range def.Tables {
        //fmt.Printf("// %s/n", k)
        def.defineTable(k)
        def.definedTables[idx] = k
        idx++
    }

    def.defineTableFinder()
    return nil
}

func (def *Definitions) Initialize() error {
    var tableDef strings.Builder
    tableDef.WriteString("package hl7\n")
    tableDef.WriteString("import (\n")
    tableDef.WriteString("    \"reflect\"\n")
    tableDef.WriteString("    \"log\"\n")
    tableDef.WriteString("    \"fmt\"\n")
    tableDef.WriteString(")")
    def.initialized = true
    fmt.Println(tableDef.String())
    return nil
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

