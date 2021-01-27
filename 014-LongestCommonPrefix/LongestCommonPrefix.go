package longestcommonprefix

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    var result []byte
    for i := 0; i < len(strs[0]); i++ {
        result = append(result, strs[0][i])
        for j := 1; j < len(strs); j++ {
            if len(strs[j])-1 < i || strs[j][i] != result[len(result)-1] {
                return string(result[:len(result)-1])
            }
        }
    }
    return strs[0]
}