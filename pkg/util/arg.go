package util

import "fmt"

func ArgToStringArray(name string, value interface{}) [2]string {
    result := [2]string{}
    result[0] = fmt.Sprintf("--%s", name)
    result[1] = fmt.Sprintf("%s", value)
    return result
}

func ArgsToStringArray(args map[string]interface{}) []string {
    result := make([]string, 10)
    for name, value := range args {
        arg := ArgToStringArray(name, value)
        result = append(append(result, arg[0], ), arg[1])
    }
    return result
}
