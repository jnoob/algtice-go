package bytedance

import "strconv"

func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return nil
	}
	dp := make(map[string][]string)
	return restoreIpSegs(s, 4, dp)
}

func restoreIpSegs(s string, segcount int, dp map[string][]string) []string {
	k := s + "@" + string('0'+segcount)
	_, exists := dp[k]
	if exists {
		return dp[k]
	}
	var result []string
	if segcount == 1 {
		if len(s) > 1 && s[0]=='0' {
			result = nil
		} else {
			seg, _ := strconv.Atoi(s)
			if seg > 255 { // >= 255
				result = nil
			} else {
				result = []string{s}
			}
		}

	} else {
		var minl, maxl int
		if s[0] == '0' {
			minl, maxl = 1, 1
		} else {
			minl, maxl = len(s) - (segcount - 1) * 3, len(s) - segcount + 1
			if minl < 1 {
				minl = 1
			}
			if maxl > 3 {
				maxl = 3
			}
		}

		result = []string{}
		for i:=minl; i <= maxl; i++ {
			if i == 3 {
				seg, _ := strconv.Atoi(s[0:i])
				if seg > 255 {
					break
				}
			}
			subRs := restoreIpSegs(s[i:], segcount-1, dp)
			if subRs != nil {
				for _, subR := range subRs {
					result = append(result, s[0:i] + "." + subR)
				}
			}
		}
	}
	dp[k] = result
	return result
}
