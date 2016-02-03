package backend

func (b *Backend) Banned(user_id, channel_id string) (bool, error) {
	return false, nil
}

func (b *Backend) Ban(user_id, channel_id string) error {
	return nil
}
