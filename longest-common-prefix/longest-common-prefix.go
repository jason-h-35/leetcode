import "strings"
func longestCommonPrefix(strs []string) string {
    minlen := 200
    min := ""
    // O(n), 1st pass to find min length.
    for _, s := range strs {
        if len(s) < minlen {
            min = s
            minlen = len(s)
        }
    }
    // O(n^2)
    for j:=0;j<minlen;j++ {
        check := strs[0][j] // grab letter to check others
        for i:=1;i<len(strs);i++ {
            if strs[i][j] != check {
                return strs[0][:j]
            }
        }
    }
    return min
}