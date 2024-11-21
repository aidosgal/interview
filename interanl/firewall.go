package tasks

import (
	"fmt"
	"strconv"
	"strings"
)

func isAllowed(rules [][2]string, ip string) bool {

    target_ip := strings.Split(ip, ".")

    for _, rule := range rules {
        address := strings.Split(rule[0], "/")

        ip := address[0]
        maskStr := ""
        mask := 32
        
        if (len(address) > 1) { 
            maskStr = address[1]
            mask, _ = strconv.Atoi(maskStr)
        }    

        ip_parts := strings.Split(ip, ".") 
        is_equal := true
        for i := 0; i < mask/8; i++ {
            if(ip_parts[i] != target_ip[i]) {
                is_equal = false
            }
        }

        if(is_equal) {
            if (rule[1] == "ALLOW") {
                return true
            }
        }
    }
    return false
}

func Run(){
    rules1 := [][2]string{
        {"192.168.1.0/24", "DENY"},
        {"10.0.0.0/16", "DENY"},
        {"8.8.8.8", "ALLOW"},
    }

    fmt.Println(isAllowed(rules1, "192.168.1.10"))
}
