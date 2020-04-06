package main;

func groupAnagrams(strs []string) [][]string {
    l :=  make(map[int][]string);
    
    for _, w := range strs{
        s:= hash(w);

        c, ok := l[s];
        if ok {
            l[s] = append(c, w);
        } else{
            l[s] = []string { w };
        }
    }

    r:= [][]string { };
    for _, val := range l {
        r = append(r, val)
    }

    return r;
}

var prime = []int {2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103};
func hash(s string) int{
    hash:= 1;
    for _, c := range s{
        hash *= prime[c - 'a'];
    }
    return hash;
}

func main(){
	groupAnagrams([]string{ "test", "lel"});
}