package bytedance

const slash = '/'
const oneDot = "."
const twoDot = ".."

func simplifyPath(path string) string {
	if len(path) == 0 || path[0] != slash {
		return "/"
	}

	parts := make([]string, 0)
	tail, pIndex := -1, -1
	for i, c := range path {
		if c == slash {
			if tail == i-1 { // first / or //
				tail = i
			} else {
				part := path[tail+1 : i]
				tail = i
				parts, pIndex = handlePart(parts, part, pIndex)
			}
		}
	}
	if tail < len(path)-1 {
		part := path[tail+1:]
		parts, pIndex = handlePart(parts, part, pIndex)
	}

	r := "/"
	if pIndex >= 0 {
		for i := 0; i < pIndex; i++ {
			r += parts[i] + "/"
		}
		r += parts[pIndex]
	}
	return r
}

func handlePart(parts []string, part string, pIndex int) ([]string, int) {
	if part == twoDot {
		if pIndex > -1 {
			pIndex--
		}
	} else if part != oneDot {
		if pIndex == len(parts)-1 { // has empty position
			parts = append(parts, part)
		} else {
			parts[pIndex+1] = part
		}
		pIndex++
	}
	return parts, pIndex
}
