package main

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    var letters [26]int
    for i := range s { //Use either because they are the same len
        letters[s[i] - 'a']++
        letters[t[i] - 'a']--
    }
    for _, v := range letters {
        if v != 0{
            return false
        }
    }
    return true
}
