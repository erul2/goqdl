package goqdl

import "regexp"

// getURLInfo returns the type of the url and the id.
func getURLInfo(url string) (string, string) {
	// Returns the type of the url and the id.
	//
	// Compatible with urls of the form:
	//     https://www.qobuz.com/us-en/{type}/{name}/{id}
	//     https://open.qobuz.com/{type}/{id}
	//     https://play.qobuz.com/{type}/{id}
	//     /us-en/{type}/-/{id}

	r := regexp.MustCompile(`(?:https:\/\/(?:w{3}|open|play)\.qobuz\.com)?(?:\/[a-z]{2}-[a-z]{2})` +
		`?\/(album|artist|track|playlist|label)(?:\/[-\w\d]+)?\/([\w\d]+)`)
	matches := r.FindStringSubmatch(url)
	if len(matches) == 3 {
		return matches[1], matches[2]
	}
	return "", ""
}
