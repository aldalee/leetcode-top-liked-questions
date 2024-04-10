package back_tracking

func combinationSum(candidates []int, target int) (res [][]int) {
	if target == 0 {
		return [][]int{{}}
	}
	for i, candidate := range candidates {
		if target-candidate >= 0 {
			rest := combinationSum(candidates[i:], target-candidate)
			for _, comb := range rest {
				comb = append(comb, candidate)
				res = append(res, comb)
			}
		}
	}
	return res
}
