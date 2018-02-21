package main

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	_, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("値が存在しない場合は authAvatar.GetAvatarURL は " +
			"ErrNoAvaterURL を返すべきです")
	}

	testUrl := "http://url-to-avatar/"
	client.userData = map[string]interface{}{
		"avatar_url": testUrl,
	}
	url, err := authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("avatar_url が設定されているのに authAvatar.GetAvatarURL は"+
			"エラーを返しては行けません。 err: ", err)
	} else {
		if url != testUrl {
			t.Error("authAvatar.GetAvatarURL は正しい URL を返すべきです。url", url)
		}
	}
}
