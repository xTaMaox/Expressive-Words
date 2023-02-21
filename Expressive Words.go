func expressiveWords(s string, words []string) int {
    res := 0
    
    for i := range words {
        if len(words[i]) > len(s) {
            continue
        }
        p1, p2 := 0, 0 

        for p1 < len(s) && p2 < len(words[i]) {
            if s[p1] != words[i][p2] {
                break
            }
            c1, c2 := 1, 1
            for p1 < len(s)-1 && s[p1] == s[p1+1] {
                p1++
                c1++
            }

            for p2 < len(words[i])-1 && words[i][p2] == words[i][p2+1] {
                p2++
                c2++
            }
            if (c1 != c2 && c1 < 3 ) || c1 < c2 {
                break
            }
            p1++
            p2++
            if p1 == len(s) && p2 == len(words[i]) {
                res++
            }
        }
        
    }
    return res
}