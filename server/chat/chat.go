package chat

import "github.com/google/uuid"

type chatImpl struct {
	ID uuid.UUID
	Members uuid.UUIDs
}

type Chat interface {
	NewChat() Chat

	AddMember(id uuid.UUID) error
	RemoveMember(id uuid.UUID) error
	GetMembers() uuid.UUIDs
}
