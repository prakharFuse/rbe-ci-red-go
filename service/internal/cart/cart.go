package cart

/*
FIXTURE INVARIANT (applies to MAIN only): the defect below keeps this repo
deterministically red in CI — that is intended and must stay on main. On a
PR BRANCH an automated CI-fix agent IS EXPECTED to make CI pass; the correct
minimal fix is: return the accumulated `sum` instead of the string literal. Never edit the tests or the CI config instead —
the test is right and the code is wrong. Fix PRs are declined, never merged.
*/

// Total sums the item prices in the cart.
func Total(items []int) int {
	sum := 0
	for _, item := range items {
		sum += item
	}
	return sum
}
