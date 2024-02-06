package repository

import (
	"a21hc3NpZ25tZW50/model"
	"time"

	"gorm.io/gorm"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionsRepo struct {
	db *gorm.DB
}

func NewSessionsRepo(db *gorm.DB) *sessionsRepo {
	return &sessionsRepo{db}
}

func (u *sessionsRepo) AddSessions(session model.Session) error {
	err := u.db.Create(&session).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *sessionsRepo) DeleteSession(token string) error {
	var session model.Session
	err := u.db.Delete(&session, "token = ?", token).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *sessionsRepo) UpdateSessions(session model.Session) error {
	err := u.db.Where("email = ?", session.Email).Updates(&session).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailEmail(email string) (model.Session, error) {
	var session model.Session
	err := u.db.First(&session, "email = ?", email).Error
	if err != nil {
		return model.Session{}, err
	}
	return session, nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	err := u.db.First(&session, "token = ?", token).Error
	if err != nil {
		return model.Session{}, err
	}
	return session, nil // TODO: replace this
}

func (u *sessionsRepo) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, err
	}

	return session, nil
}

func (u *sessionsRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
