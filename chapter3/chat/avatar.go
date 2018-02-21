package main

import (
	"errors"
)

var ErrNoAvatarURL = errors.New("chat: アバターのURLを返せません")

type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct {
}

var UseGravatarAvatar GravatarAvatar

func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userId, ok := c.userData["userid"]; ok {
		if userIdStr, ok := userId.(string); ok {
			return "//www.gravatar.com/avatar/" + userIdStr, nil
		}
	}
	return "", ErrNoAvatarURL
}
