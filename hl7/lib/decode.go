package hl7

import (
    "log"
    //"joeswaminathan.com/Edge_Ble/log"
    "fmt"
    "strconv"
    "strings"
    "reflect"
    "regexp"
)

const (
    noneKey = "none"
)

// An UnmarshalTypeError describes a HL7 value that was
// not appropriate for a value of a specific Go type.
type UnmarshalTypeError struct {
	Value        string       // HL7 value
	Type         reflect.Type // type of Go value it could not be assigned to
	Path         string       // the full path from root node to the field,
                              //     include embedded struct
    Hl7Path      string       // Full path of HL7 message in the format
                              //     S/F[n]/C[n]/c[n]
                              //     Where S, F, C, c are names of
                              //        Segment, Field, Comp, Sub-comp respectively
}

func (e *UnmarshalTypeError) Error() string {
    return ( "hl7: cannot unmarshal (" + e.Value +
             ") to type (" + e.Type.String() +
             "), path (" + e.Path +
             "), hl7 path (" + e.Hl7Path +
             ")" )
}

// An InvalidUnmarshalError describes an invalid argument passed to [Unmarshal].
// (The argument to [Unmarshal] must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "hl7: Unmarshal(nil)"
	}
	if e.Type.Kind() != reflect.Pointer {
		return "hl7: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "hl7: Unmarshal(nil " + e.Type.String() + ")"
}

type decodeState struct {
    level           int
    delimiters      Delimiters
    escapeSequences EscapeSequences
    data            string
    meta            map[string]any
    value           reflect.Value
    path            string
    hl7Path         string
}

func (ds *decodeState) print(prefix string) {
    log.Printf("%sLevel (%d), data (%s), path (%s), hl7Path (%s)\n",
              prefix, ds.level, ds.data, ds.path, ds.hl7Path)
}

func (ds *decodeState) unmarshalError() error {
    return &UnmarshalTypeError{
                ds.data,
                ds.value.Type(),
                ds.path,
                ds.hl7Path,
            }
}

type parseStatus int

const (
    VALID parseStatus = 0
    REPEATABLE parseStatus = 1
    NOT_REPEATABLE parseStatus = -2
    INVALID parseStatus = -1
)


func getMetaData(rValue reflect.Value) (map[string]map[string]any) {
    metaData := make(map[string]map[string]any)
    rType := rValue.Type()
    log.Println("\tstruct ", rType)
    if rType.Kind() == reflect.Struct {
        for i := 0; i < rType.NumField(); i++ {
            field := rType.Field(i)
            log.Println("\t\t", field.Name, ", ", field.Type, ", ", field.Tag)
            tags := strings.Fields(field.Tag.Get("hl7"))
            part := ""
            for _, tag := range tags {
                if strings.HasPrefix(tag, "part") {
                    part = strings.Split(tag, "=")[1]
                    metaData[part] = make(map[string]any)
                    metaData[part]["field"] = rValue.Field(i).Addr()
                    metaData[part]["name"] = field.Name
                    metaData[part]["type"] = field.Type
                    metaData[part]["tag"] = field.Tag
                } else if part != "" && strings.HasPrefix(tag, "table") {
                    tblId := strings.Split(tag, "=")[1]
                    metaData[part]["table"] = reflect.ValueOf(findTable(tblId))
                } else if part != "" && strings.HasPrefix(tag, "length") {
                    length := strings.Split(tag, "=")[1]
                    metaData[part]["length"] = length
                } else if part != "" && strings.HasPrefix(tag, "variedTypeField") {
                    fName := strings.Split(tag, "=")[1]
                    fValue := rValue.FieldByName(fName)
                    //if fValue.IsValid() && !fValue.IsZero() {
                        metaData[part]["intefaceUnderlyingType"] = fValue.Addr()
                    //}
                }
            }
        }
    } else {
        part := noneKey
        metaData[part] = make(map[string]any)
        metaData[part]["field"] = rValue.Addr()
    }
    return metaData
}

func getPartsToField(rValue reflect.Value) (partsToField map[string]reflect.Value) {
    partsToField = make(map[string]reflect.Value)
    rType := rValue.Type()
    log.Println("\tstruct ", rType)
    if rType.Kind() == reflect.Struct {
        for i := 0; i < rType.NumField(); i++ {
            field := rType.Field(i)
            log.Println("\t\t", field.Name, ", ", field.Type, ", ", field.Tag)
            tags := strings.Fields(field.Tag.Get("hl7"))
            for _, tag := range tags {
                if strings.HasPrefix(tag, "part") {
                    part := strings.Split(tag, "=")[1]
                    partsToField[part] = rValue.Field(i).Addr()
                }
            }
        }
    } else {
        partsToField["1"] = rValue.Addr()
    }
    return partsToField
}

func (ds *decodeState) setValue(rValue reflect.Value) (parseStatus, error) {
    if !rValue.IsValid() {
        err := fmt.Errorf("Error setting value: Invalid variable")
        log.Println("\t%s\n", err.Error())
        return INVALID, err
    }
    rType := rValue.Type()
    if rType.Kind() != reflect.Pointer {
        err := fmt.Errorf("Error setting value: Not a pointer")
        log.Println("\t%s\n", err.Error())
        return INVALID, err
    }
    rValue = rValue.Elem()
    rType = rValue.Type()
    if !rValue.CanSet() {
        err := fmt.Errorf("Error setting value: Value unsettable")
        log.Println("\t%s\n", err.Error())
        return INVALID, err
    }

    switch rType.Kind() {
    case reflect.String:
        log.Println("\tSetting string")
        rValue.SetString(ds.data)
        return VALID, nil
    case reflect.Int:
        log.Println("\tSetting int")
        i, err := strconv.ParseInt(ds.data, 10, 64)
        if nil != err {
            return INVALID, err
        }
        rValue.SetInt(i)
        return VALID, nil
    case reflect.Uint:
        log.Println("\tSetting uint")
        u, err := strconv.ParseUint(ds.data, 10, 64)
        if nil != err {
            return INVALID, err
        }
        rValue.SetUint(u)
        return VALID, nil
    case reflect.Uint16:
        log.Println("\tSetting uint16")
        u, err := strconv.ParseUint(ds.data, 10, 16)
        if nil != err {
            return INVALID, err
        }
        rValue.SetUint(u)
        return VALID, nil
    case reflect.Float64:
        log.Println("\tSetting uint16")
        f, err := strconv.ParseFloat(ds.data, 64)
        if nil != err {
            return INVALID, err
        }
        rValue.SetFloat(f)
        return VALID, nil
    case reflect.Struct:
        partsToField := getMetaData(rValue)
        fieldMeta, ok := partsToField["1"]
        if !ok {
            err := fmt.Errorf("Error setting value: No suitable member in struture")
            log.Println("\t%s\n", err.Error())
            return INVALID, err
        }
        return ds.setValue(fieldMeta["field"].(reflect.Value))
    case reflect.Slice:
        rValue.Set(reflect.Append(rValue, reflect.Zero(rType.Elem())))
        return ds.setValue(rValue.Index(rValue.Len() - 1).Addr())
    case reflect.Pointer:
        if rValue.IsNil() {
            rValue.Set(reflect.New(rType.Elem()))
        }
        return ds.setValue(rValue)
    default:
        err := fmt.Errorf("Error setting value: Invalid Kind = %s", rType.Kind())
        log.Printf("\t%s\n", err.Error())
        return INVALID, err
    }
}

func (ds *decodeState) parseValue() (parseStatus, error) {
    if !ds.value.IsValid() {
        log.Println("\tInvalid value\n")
        return INVALID, ds.unmarshalError()
    }
    rValue := ds.value.Elem()
    if rValue.Kind() == reflect.Pointer {
        rValue = rValue.Elem()
    }
    log.Printf("\tparseValue type is %s\n", rValue.Type().String())
    ps, err := ds.setValue(rValue.Addr())
    if err != nil {
        return INVALID, ds.unmarshalError()
    }
    return ps, nil
}

func addIndex(path string, idx int) string {
    pattern := "\\[.+\\]$"
    re := regexp.MustCompile(pattern)
    match := re.FindStringIndex(path)

    index := "[" + strconv.Itoa(idx) + "]"
    if match != nil {
        return path[:match[0]] + index
    } else {
        return path + index
    }
}

func (ds *decodeState) getUnderlyingType() (reflect.Value) {
    fIntf, ok := ds.meta["intefaceUnderlyingType"]
    if !ok {
        return reflect.Value{}
    }
    fValue := fIntf.(reflect.Value)
    return fValue.Elem()
}

func (ds *decodeState) parseHl7() (parseStatus, error) {
    ds.print("\t")
    if len(ds.data) == 0 {
        return VALID, nil
    }

    /*
    status := ds.validate()
    if status != VALID {
        log.Println("\tStatus (%d)\n", status)
        return status, ds.unmarshalError()
    }
    */

    // Verify the value is a pointer, and not nil
    rValue := ds.value
    rType := rValue.Type()
    log.Printf("\tds.value type is %s\n", rType)
    if ( rType.Kind() !=  reflect.Pointer ||
         rValue.IsNil() ) {
        log.Println("\tNot a pointer type. Invalid\n")
        return INVALID, ds.unmarshalError()
	}

    if rType.Elem().Kind() == reflect.Pointer {
        log.Printf("\tds.value.Elem type is %s\n", rValue.Elem().Type())
        rValue = rValue.Elem()
        rType = rValue.Type()
        if rValue.IsNil() {
            rValue.Set(reflect.New(rType.Elem()))
        }
    }

    if rType.Elem().Kind() == reflect.Interface {
        fValue := ds.getUnderlyingType()
        if fValue.IsValid() {
            return ds.instantiateAndParseHl7(fValue.String())
        } else {
            err := fmt.Errorf("Error setting any type field: underlying type not found")
            log.Printf("\t%s\n", err.Error())
            return INVALID, err
        }
    }

    delimiter := ds.delimiters[ds.level]
    repeats := false
    nextLevel := 0
    if rType.Elem().Kind() == reflect.Slice {
        if rValue.IsNil() {
            rValue.Set(reflect.MakeSlice(rType, 0, 0))
        }
        log.Printf("\tJOE: Slice element type (%s)\n", rType.Elem().String())
        if rType.Elem().String() != "hl7.NA" {
            if ds.level == LEVEL_1 {
                delimiter = ds.delimiters[SEGMENT_SEPARATOR]
            } else {
                delimiter = ds.delimiters[REPEAT_SEPARATOR]
            }
        }
        /*
        isStruct := false
        if ( rType.Elem().Elem().Kind() == reflect.Pointer &&

             rType.Elem().Elem().Elem().Kind() == reflect.Struct ) {
            isStruct = true
        } else
        if ( rType.Elem().Elem().Kind() == reflect.Interface ) {
            fValue := ds.getUnderlyingType()
            if fValue.String() == "NA" {
                if ds.level == LEVEL_1 {
                    delimiter = ds.delimiters[SEGMENT_SEPARATOR]
                } else {
                    delimiter = ds.delimiters[REPEAT_SEPARATOR]
                }
            }
        }
        */
        repeats = true
        nextLevel = ds.level
    } else {
        nextLevel = ds.level + 1
    }

    /*
    // the value is of type pointer or interface get the
    // underlying value and type
    if rValue.Type() == reflect.Pointer {
        rValue = rValue.Elem()
        rType = rValue.Type()
    }
    */

    parts := strings.Split(ds.data, delimiter)
    // Check for empty data
    if parts == nil || len(parts) == 0 {
        log.Println("\tEmpty data\n")
        return VALID, nil
    }

    metaData := getMetaData(rValue.Elem())

    if !repeats && len(parts) == 1 && ds.level > 0 {
        return ds.parseValue()
    }

    var dsPart *decodeState = nil
    var fValue reflect.Value
    for fieldIdx, part := range parts {
        if len(part) == 0 {
            continue
        }
        //if dsPart == nil {
            copyPart := *ds
            dsPart = &copyPart
            dsPart.data = part
            dsPart.level = nextLevel
            key := ""
            if rType.Elem().Kind() == reflect.Struct {
                if dsPart.level == LEVEL_1 {
                    key = part[:3]
                } else {
                    key = strconv.Itoa(fieldIdx+1)
                }
            } else {
                key = noneKey
            }
            log.Println("Parsing part (", part, "), Level = ", dsPart.level, ", Key = ", key)
            var ok bool
            fMeta, ok := metaData[key]
            if !ok {
                if rType.Elem().Kind() == reflect.Struct {
                    continue
                } else {
                    break
                }
            }
            if key == noneKey {
                dsPart.meta = ds.meta
            } else {
                dsPart.meta = fMeta
            }
            /*
            if fMeta["name"] == "ObservationValue" {
                fmt.Printf("Handling field of type 'varies'")
            }
            */
            fValue = fMeta["field"].(reflect.Value)
            if !fValue.Elem().CanSet() {
                log.Println("\tUnsettable value")
                return INVALID, dsPart.unmarshalError()
            }
            fType := fValue.Type()
            if repeats {
                fValue = fValue.Elem()
                fType  = fValue.Type()
                fValue.Set(reflect.Append(fValue, reflect.Zero(fType.Elem())))
                dsPart.value = fValue.Index(fValue.Len() - 1).Addr()
            } else {
                dsPart.value = fValue
            }
            log.Println("\tValue = ", fValue, ", Type = ", fType)
            dsPart.path = dsPart.path + "/" + fType.String()
            dsPart.hl7Path = ( dsPart.hl7Path + "/" +
                                strconv.Itoa(dsPart.level) + "-" + key )
        //}
        if repeats {
            dsPart.hl7Path = addIndex(dsPart.hl7Path, fieldIdx+1)
        }
        status, err := dsPart.parseHl7()
        log.Println("Parsed part (", part, "), Status = ", status, ", Error = ", err, "\n")
        if status == INVALID {
            return INVALID, dsPart.unmarshalError()
        } else if status == VALID {
            if repeats || rType.Elem().Kind() == reflect.Struct {
                continue
            } else {
                break
            }
        } else {
            log.Println("Error: status unknown")
            return INVALID, dsPart.unmarshalError()
        }
    }
    return VALID, nil
}

func (ds *decodeState) init(data []byte, v any) error {
    if ds == nil {
        log.Println("Error: nil decode state")
        return  &InvalidUnmarshalError{}
    }
    if v == nil {
        log.Println("Error: nil value")
        return  &InvalidUnmarshalError{}
    }
    ds.value = reflect.ValueOf(v)
    if !ds.value.IsValid() || ds.value.IsNil() || ds.value.Kind() != reflect.Pointer {
        return  &InvalidUnmarshalError{}
    }
    ds.path = reflect.TypeOf(v).String()
    ds.hl7Path = "HL7"
    ds.level = 0
    if len(data) < 8 || string(data[:3]) != "MSH" {
        log.Println("Error: Invalid data")
        return  &InvalidUnmarshalError{}
    }
    ds.delimiters, ds.escapeSequences = newDelimitersAndEscapeSequences(data[3:8])
    for i:=4; i<8; i++ {
        data[i] = '?'
    }
    ds.data = string(data)
    return nil
}


func Unmarshal(data []byte, v any) error {
    rValue := reflect.ValueOf(v)
    if rValue.Kind() != reflect.Pointer || rValue.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}
    ds := &decodeState{}
    if err := ds.init(data, v); err != nil {
        return err
    }
    if _, err := ds.parseHl7(); err != nil {
        return err
    }
    return nil
}
