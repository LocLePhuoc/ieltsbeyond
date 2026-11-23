package llm

import "context"

type UserPart struct {
	Text      string
	ImageData []byte
	ImageMIME string
}

func TextPart(s string) UserPart {
	return UserPart{Text: s}
}

func ImagePart(data []byte, mime string) UserPart {
	return UserPart{ImageData: data, ImageMIME: mime}
}

func (p UserPart) IsImage() bool {
	return len(p.ImageData) > 0
}

type Provider interface {
	Complete(ctx context.Context, systemPrompt string, parts []UserPart) (string, error)
}
