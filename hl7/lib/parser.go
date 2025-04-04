package hl7

import (
    "fmt"
    "unicode"
    "strings"
)

var typeIsStruct = make(map[string]bool)

func (def *Definitions) defineType(dtn string) error {
    var typeDef strings.Builder
    //fmt.Printf("--- Definining %s ---\n", dtn)
    if _, ok := typeIsStruct[dtn]; ok {
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
                if typeIsStruct[part.Type] {
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
                typeDef.WriteString(fmt.Sprintf(" `json:\"%s\"", part.Name))
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
                typeDef.WriteString(fmt.Sprintf("\"`"))

            }
            typeDef.WriteString(fmt.Sprintf("\n"))
        }
        if dt.Separator != "" {
            typeDef.WriteString(fmt.Sprintf("}"))
        }
        typeDef.WriteString(fmt.Sprintf("\n"))
    }
    typeIsStruct[dtn] = (dt.Separator != "")
    //fmt.Printf("--- Defined %s ---\n", dtn)
    fmt.Println(typeDef.String())
    return nil
}

func (def *Definitions) MakeTypes() error {
    if def.DataTypes == nil {
        return nil
    }
    for k := range def.DataTypes {
        //fmt.Printf("// %s/n", k)
        def.defineType(k)
    }
    return nil
}
