package entities

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

const (
	AgentSessionMinLength = 3
	AgentSessionMaxLength = 255
)

type Session struct {
	ID           IdSession
	UserID       IdUser
	IP           IpSession
	Agent        AgentSession
	RefreshToken RefreshTokenSession
	ExpiredAt    *time.Time
	DeclinedAt   *time.Time
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}

func NewSession(id string, userId string, ip string, agent string, expiredAt *time.Time, declinedAt *time.Time, createdAt *time.Time, updatedAt *time.Time, deletedAt *time.Time) (*Session, error) {
	idValueObject, err := NewIdSession(id)

	if err != nil {
		return nil, err
	}

	userIdValueObject, err := NewIdUser(userId)

	if err != nil {
		return nil, err
	}

	ipValueObject, err := NewIpSession(ip)

	if err != nil {
		return nil, err
	}

	agentValueObject, err := NewAgentSession(agent)

	if err != nil {
		return nil, err
	}

	refreshTokenValueObject, err := NewRefreshTokenSession()

	if err != nil {
		return nil, err
	}

	session := Session{
		ID:           *idValueObject,
		UserID:       *userIdValueObject,
		IP:           *ipValueObject,
		Agent:        *agentValueObject,
		RefreshToken: *refreshTokenValueObject,
	}

	if expiredAt != nil {
		expiredAtCopy := *expiredAt
		session.ExpiredAt = &expiredAtCopy
	}

	if declinedAt != nil {
		declinedAtCopy := *declinedAt
		session.DeclinedAt = &declinedAtCopy
	}

	if createdAt != nil {
		createdAtCopy := *createdAt
		session.CreatedAt = &createdAtCopy
	}

	if updatedAt != nil {
		updatedAtCopy := *updatedAt
		session.UpdatedAt = &updatedAtCopy
	}

	if deletedAt != nil {
		deletedAtCopy := *deletedAt
		session.DeletedAt = &deletedAtCopy
	}

	return &session, nil
}

func (session *Session) Decline(declinedAt *time.Time) error {
	if declinedAt == nil {
		tempDeclinedAt := time.Now()
		declinedAt = &tempDeclinedAt
	}

	session.DeclinedAt = declinedAt

	return nil
}

type IdSession struct {
	id string
}

func NewIdSession(id string) (*IdSession, error) {
	if id == "" {
		newUUID, err := uuid.NewRandom()

		if err != nil {
			return nil, err
		}

		id = newUUID.String()
	}

	err := validate.Var(id, "required,uuid")

	if err != nil {
		return nil, err
	}

	return &IdSession{id: id}, nil
}

func (valueObject IdSession) GetId() string {
	return valueObject.id
}

type IpSession struct {
	ip string
}

func NewIpSession(ip string) (*IpSession, error) {
	err := validate.Var(ip, "required,ip")

	if err != nil {
		return nil, err
	}

	return &IpSession{ip: ip}, nil
}

func (valueObject IpSession) GetIp() string {
	return valueObject.ip
}

type AgentSession struct {
	agent string
}

func NewAgentSession(agent string) (*AgentSession, error) {
	err := validate.Var(agent, fmt.Sprintf("required,gte=%d,lte=%d", AgentSessionMinLength, AgentSessionMaxLength))

	if err != nil {
		return nil, err
	}

	return &AgentSession{agent: agent}, nil
}

func (valueObject AgentSession) GetAgent() string {
	return valueObject.agent
}

type RefreshTokenSession struct {
	token string
}

func NewRefreshTokenSession() (*RefreshTokenSession, error) {
	newUUID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	token := newUUID.String()

	return &RefreshTokenSession{token: token}, nil
}

func (valueObject RefreshTokenSession) GetToken() string {
	return valueObject.token
}
