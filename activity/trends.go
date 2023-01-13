package activity

import "fmt"

func (u UserActivity) GetMostTrending(pastUsers []User) User {
	tempMax := 0
	var tempMaxUser User

loopMax:
	for key, val := range u.Activity {
		for _, val2 := range pastUsers {
			if key == val2 {
				continue loopMax
			}
		}
		if len(*val.LikeBy) > tempMax {
			tempMaxUser = key
			tempMax = len(*val.LikeBy)
		}
	}
	return tempMaxUser
}

func (u UserActivity) Trending() string {
	sortedTrend := "Trending photos:\n"
	var trendUser []User

	for i := 0; i < 3; i++ {
		trendUser = append(trendUser, u.GetMostTrending(trendUser))
	}

	for i, val := range trendUser {
		if val.Name != "" {
			sortedTrend += fmt.Sprintf("%v. %v photo got %v likes\n", i+1, val.Name, len(*u.Activity[val].LikeBy))
		}
	}
	return sortedTrend
}
