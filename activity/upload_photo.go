package activity

func (u *UserActivity) Upload(user User) error {
	if u.CheckUserExist(user) == false {
		return ErrorUnknownUser{&user}
	}

	u.Activity[user].Upload = true

	for _, notifyFollower := range *u.Activity[user].Follower {
		for _, followedNotifyUpload := range *u.Activity[notifyFollower].Notify.Upload {
			if followedNotifyUpload == user {
				return nil
			}
		}
		*u.Activity[notifyFollower].Notify.Upload = append(*u.Activity[notifyFollower].Notify.Upload, user)
	}
	return nil
}
