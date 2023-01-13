package activity

func (u *UserActivity) Like(initUser, targetUser User) error {
	if u.CheckUserExist(initUser) == false {
		return ErrorUnknownUser{&initUser}
	}

	if u.CheckUserExist(targetUser) == false {
		return ErrorUnknownUser{&targetUser}
	}

	if u.Activity[targetUser].Upload == false {
		return ErrorNoPhoto{}
	}

	for _, liked := range *u.Activity[initUser].Likes {
		if liked == targetUser {
			return nil
		}
	}

	if initUser == targetUser {
		*u.Activity[initUser].Likes = append(*u.Activity[initUser].Likes, targetUser)
		*u.Activity[targetUser].LikeBy = append(*u.Activity[targetUser].LikeBy, initUser)

		for _, notifyFollower := range *u.Activity[initUser].Follower {
			addedNotification := []User{initUser, targetUser}
			*u.Activity[notifyFollower].Notify.Likes = append(*u.Activity[notifyFollower].Notify.Likes, addedNotification)
		}
		return nil
	}

	for _, followed := range *u.Activity[initUser].Notify.Following {
		if followed == targetUser {
			*u.Activity[initUser].Likes = append(*u.Activity[initUser].Likes, targetUser)
			*u.Activity[targetUser].LikeBy = append(*u.Activity[targetUser].LikeBy, initUser)

			for _, notifyFollower := range *u.Activity[initUser].Follower {
				addedNotification := []User{initUser, targetUser}
				*u.Activity[notifyFollower].Notify.Likes = append(*u.Activity[notifyFollower].Notify.Likes, addedNotification)
			}
			return nil
		}
	}
	return ErrorNotFollowed{}
}
