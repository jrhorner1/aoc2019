package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func get_opcode(opcode_str string) (int, []int) {
    var instr int
        var params []int
        switch len(opcode_str) {
        case 1,2:
            instr, _ = strconv.Atoi(opcode_str)
            for i := 2; i >= 0; i-- {
                param := 0
                params = append(params, param)
            }
        case 3:
            instr, _ = strconv.Atoi(opcode_str[1:])
            for i := 0; i >= -2; i-- {
                if i == 0 {
                    param := int(opcode_str[i] - '0')
                    params = append(params, param)
                } else {
                    param := 0
                    params = append(params, param)
                }
            }
        case 4:
            instr, _ = strconv.Atoi(opcode_str[2:])
            for i := 1; i >= -1; i-- {
                if i == 0 || i == 1 {
                    param := int(opcode_str[i] - '0')
                    params = append(params, param)
                } else {
                    param := 0
                    params = append(params, param)
                }
            }
        case 5:
            instr, _ = strconv.Atoi(opcode_str[3:])
            for i := 2; i >= 0; i-- {
                param := int(opcode_str[i] - '0')
                params = append(params, param)
            }
        }

    return instr, params
}

func get_params(instr int, params []int, intcode_pos int, intcode_str *[]string) []int {
    var param_len int
    switch instr {
    case 1,2,7,8: param_len = len(params)
    case 5,6: param_len = len(params) - 1
    }
    for i := 0; i < param_len; i++ {
        temp, _ := strconv.Atoi((*intcode_str)[intcode_pos + i + 1])
        switch params[i] {
        case 1:
            params[i] = temp
        case 0:
            params[i], _ = strconv.Atoi((*intcode_str)[temp])
        }
    }
    return params
}

func procIntc(intcode_str *[]string, phase int, input int) int {
    // Process each opcode
    var output int
    var break_code bool
    var intcode_adv, inputCount int = 0, 0
    for intcode_pos := 0; break_code == false; intcode_pos += intcode_adv {
        instr, params := get_opcode((*intcode_str)[intcode_pos])
        params = get_params(instr, params, intcode_pos, intcode_str)
        // process opcode
        switch instr {
        case 1: // addition
            intcode_adv = 4
            params[2] = params[0] + params[1]
            temp, _ := strconv.Atoi((*intcode_str)[intcode_pos + 3])
            (*intcode_str)[temp] = strconv.Itoa(params[2])
        case 2: // multiplication
            intcode_adv = 4
            params[2] = params[0] * params[1]
            temp, _ := strconv.Atoi((*intcode_str)[intcode_pos + 3])
            (*intcode_str)[temp] = strconv.Itoa(params[2])
        case 3: // input
            intcode_adv = 2 
            params[0], _ = strconv.Atoi((*intcode_str)[intcode_pos + 1])
            switch inputCount {
            case 0:
                (*intcode_str)[params[0]] = strconv.Itoa(phase)
                inputCount++
            default:
                (*intcode_str)[params[0]] = strconv.Itoa(input)
            }
        case 4: // output
            intcode_adv = 2
            if params[0] == 0 {
                temp, _ := strconv.Atoi((*intcode_str)[intcode_pos+1])
                output, _ = strconv.Atoi((*intcode_str)[temp])
                continue
            }
            fmt.Println((*intcode_str)[intcode_pos+1])
        case 5: // jump-if-true
            if params[0] != 0 {
                intcode_adv = params[1] - intcode_pos
            } else {
                intcode_adv = 3
            }
        case 6: // jump-if-false
            if params[0] == 0 {
                intcode_adv = params[1] - intcode_pos
            } else {
                intcode_adv = 3
            }
        case 7: // less than
            intcode_adv = 4
            temp, _ := strconv.Atoi((*intcode_str)[intcode_pos + 3])
            if params[0] < params[1] {
                (*intcode_str)[temp] = "1"
            } else {
                (*intcode_str)[temp] = "0"
            }
        case 8: // equals
            intcode_adv = 4
            temp, _ := strconv.Atoi((*intcode_str)[intcode_pos + 3])
            if params[0] == params[1] {
                (*intcode_str)[temp] = "1"
            } else {
                (*intcode_str)[temp] = "0"
            }
        case 99: // end program
            break_code = true
        }
    }
    return output
}

func permutations(arr []int)[][]int{
    var helper func([]int, int)
    res := [][]int{}

    helper = func(arr []int, n int){
        if n == 1{
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
}

func main() {
	// change to request the input file from user
    file, _ := os.Open("input")
    defer file.Close()

    var intcode_str []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	intcode_str = strings.Split(scanner.Text(),",")
    }

    // generate list of permutations
    var phasesList [][]int
    phaseVals := []int{0, 1, 2, 3, 4}
    for _, perm := range(permutations(phaseVals)){
       phasesList = append(phasesList, perm)
    }
    fmt.Println(len(phasesList),phasesList)

    // loop through list using each as the values for the incode
    var outputList [][]int
    for i := range phasesList {
        phases := phasesList[i]
        var output []int
        output = append(output, 0)
        for j := range phases {
            phase := phases[j]
            input := output[len(output)-1]
            output = append(output, procIntc(&intcode_str, phase, input))

        }
        outputList = append(outputList, output)
    }
    fmt.Println(outputList)
    
    // loop through the output to determine which final value is largest
    var ans int = 0
    for i := range outputList {
        output := outputList[i]
        if output[5] > ans {
            ans = output[5]
        }
    }
    fmt.Println("Solution:",ans)
}