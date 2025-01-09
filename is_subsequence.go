package main

func IsSubsequence(s string, t string) bool {

	if s == "" {
		return true
	}
	if s == "" && t == "" {
		return true
	}

	main_list := []rune(t)
	sub_list := []rune(s)
	sub_str := ""
	sub_str_ind := 0
	for i := 0; i < len(main_list); i++ {
		if sub_list[sub_str_ind] == main_list[i] {
			sub_str += string(main_list[i])
			sub_str_ind += 1
		}
		if sub_str == s {
			return true
		}
	}

	return false
}
