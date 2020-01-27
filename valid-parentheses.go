

package main

func isValid(s string) bool {
	var stk []string
	var top = ""
	for _, _c := range s {
		if len(stk) > 0 {
			top = stk[len(stk)-1]
		}
		var c = string(_c)

		switch c {
		case "(":
			{
				stk = append(stk, "(")
				break
			}
		case "[":
			{
				stk = append(stk, "[")
				break
			}
		case "{":
			{
				stk = append(stk, "{")
				break
			}
		case ")":
			{
				if top != "(" {
					return false
				}
				stk = stk[0:len(stk)-1]

				break
			}
		case "]":
			{
				if top != "[" {
					return false
				}
				stk = stk[0:len(stk)-1]
				break
			}
		case "}":
			{
				if top != "{" {
					return false
				}
				stk = stk[0:len(stk)-1]
				break
			}
		default:
			{
				return false
			}
		}
	}

    return len(stk)==0
}
func main() {
	var s = ""
	println (isValid(s))
}