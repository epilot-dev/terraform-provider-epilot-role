// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package utils

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/ericlagergren/decimal"

	"github.com/epilot-dev/terraform-provider-epilot-role/internal/sdk/types"
)

func GenerateURL(_ context.Context, serverURL, path string, pathParams interface{}, globals interface{}) (string, error) {
	uri := strings.TrimSuffix(serverURL, "/") + path

	parsedParameters := map[string]string{}

	globalsAlreadyPopulated, err := populateParsedParameters(pathParams, globals, parsedParameters, []string{})
	if err != nil {
		return "", err
	}

	if globals != nil {
		_, err = populateParsedParameters(globals, nil, parsedParameters, globalsAlreadyPopulated)
		if err != nil {
			return "", err
		}
	}

	// TODO should we handle the case where there are no matching path params?
	return ReplaceParameters(uri, parsedParameters), nil
}

func populateParsedParameters(pathParams interface{}, globals interface{}, parsedParameters map[string]string, skipFields []string) ([]string, error) {
	pathParamsStructType, pathParamsValType := dereferencePointers(reflect.TypeOf(pathParams), reflect.ValueOf(pathParams))

	globalsAlreadyPopulated := []string{}

	for i := 0; i < pathParamsStructType.NumField(); i++ {
		fieldType := pathParamsStructType.Field(i)
		valType := pathParamsValType.Field(i)

		if contains(skipFields, fieldType.Name) {
			continue
		}

		requestTag := getRequestTag(fieldType)
		if requestTag != nil {
			continue
		}

		ppTag := parseParamTag(pathParamTagKey, fieldType, "simple", false)
		if ppTag == nil {
			continue
		}

		if globals != nil {
			var globalFound bool
			fieldType, valType, globalFound = populateFromGlobals(fieldType, valType, pathParamTagKey, globals)
			if globalFound {
				globalsAlreadyPopulated = append(globalsAlreadyPopulated, fieldType.Name)
			}
		}

		if ppTag.Serialization != "" {
			vals, err := populateSerializedParams(ppTag, fieldType.Type, valType)
			if err != nil {
				return nil, err
			}
			for k, v := range vals {
				parsedParameters[k] = url.PathEscape(v)
			}
		} else {
			// TODO: support other styles
			switch ppTag.Style {
			case "simple":
				simpleParams := getSimplePathParams(ppTag.ParamName, fieldType.Type, valType, ppTag.Explode)
				for k, v := range simpleParams {
					parsedParameters[k] = v
				}
			}
		}
	}

	return globalsAlreadyPopulated, nil
}

func getSimplePathParams(parentName string, objType reflect.Type, objValue reflect.Value, explode bool) map[string]string {
	pathParams := make(map[string]string)

	if isNil(objType, objValue) {
		return nil
	}

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
		objValue = objValue.Elem()
	}

	switch objType.Kind() {
	case reflect.Array, reflect.Slice:
		if objValue.Len() == 0 {
			return nil
		}
		var ppVals []string
		for i := 0; i < objValue.Len(); i++ {
			ppVals = append(ppVals, valToString(objValue.Index(i).Interface()))
		}
		pathParams[parentName] = strings.Join(ppVals, ",")
	case reflect.Map:
		if objValue.Len() == 0 {
			return nil
		}
		var ppVals []string
		objMap := objValue.MapRange()
		for objMap.Next() {
			if explode {
				ppVals = append(ppVals, fmt.Sprintf("%s=%s", objMap.Key().String(), valToString(objMap.Value().Interface())))
			} else {
				ppVals = append(ppVals, fmt.Sprintf("%s,%s", objMap.Key().String(), valToString(objMap.Value().Interface())))
			}
		}
		pathParams[parentName] = strings.Join(ppVals, ",")
	case reflect.Struct:
		switch objValue.Interface().(type) {
		case time.Time:
			pathParams[parentName] = valToString(objValue.Interface())
		case types.Date:
			pathParams[parentName] = valToString(objValue.Interface())
		case big.Int:
			pathParams[parentName] = valToString(objValue.Interface())
		case decimal.Big:
			pathParams[parentName] = valToString(objValue.Interface())
		default:
			var ppVals []string
			for i := 0; i < objType.NumField(); i++ {
				fieldType := objType.Field(i)
				valType := objValue.Field(i)

				ppTag := parseParamTag(pathParamTagKey, fieldType, "simple", explode)
				if ppTag == nil {
					continue
				}

				if isNil(fieldType.Type, valType) {
					continue
				}

				if fieldType.Type.Kind() == reflect.Pointer {
					valType = valType.Elem()
				}

				if explode {
					ppVals = append(ppVals, fmt.Sprintf("%s=%s", ppTag.ParamName, valToString(valType.Interface())))
				} else {
					ppVals = append(ppVals, fmt.Sprintf("%s,%s", ppTag.ParamName, valToString(valType.Interface())))
				}
			}
			pathParams[parentName] = strings.Join(ppVals, ",")
		}
	default:
		pathParams[parentName] = valToString(objValue.Interface())
	}

	return pathParams
}
