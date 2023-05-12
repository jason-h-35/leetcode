import "strings"

func addBinary(a string, b string) string {
    var br strings.Builder
    carry := false
    if len(a) < len(b) {
        pad := strings.Repeat("0", len(b)-len(a))
        a = pad + a
    } else {
        pad := strings.Repeat("0", len(a)-len(b))
        b = pad + b
    }
    for i:=len(a)-1;0 <= i; i-- {
        if a[i] == '1' && b[i] == '1' {
            if carry {
                br.WriteRune('1')
                carry = true
            } else {
                br.WriteRune('0')
                carry = true
            }
        }
        if (a[i] == '1' && b[i] == '0') || (a[i] == '0' && b[i] == '1') {
            if carry {
                br.WriteRune('0')
                carry = true
            } else {
                br.WriteRune('1')
            }
        }
        if a[i] == '0' && b[i] == '0' {
            if carry {
                br.WriteRune('1')
                carry = false
            } else {
                br.WriteRune('0')
            }
        }
    }
    if carry {
        br.WriteRune('1')
    }
    rev := br.String()
    runes := []rune(rev)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}