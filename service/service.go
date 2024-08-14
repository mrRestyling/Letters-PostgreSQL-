package service

import (
	"Psql/model"
)

type Service struct {
	Storage StoragerInterface
}

type StoragerInterface interface {
	SetName(fio model.FIO) (model.FIO, error)
	AddLetter(letter model.Letter) (string, error)
	AddMessage(message model.Message) (string, error)
}

func New(s StoragerInterface) Service {
	return Service{
		Storage: s,
	}
}

func (s Service) SetNameServ(fio model.FIO) (int, error) {
	// 1.Здесь все проверки
	// 2.Здесь Бизнес-логика
	// var err error
	fio, err := s.Storage.SetName(fio) // уже другое fio
	if err != nil {
		return -1, err
	}

	return fio.UserID, nil
}

func (s Service) AddLetterServ(letter model.Letter) (string, error) {

	letterStr, err := s.Storage.AddLetter(letter)
	if err != nil {
		return "ошибка в сервисном уровне", err
	}
	return letterStr, nil
}

func (s Service) AddMessageServ(messageFromHand model.Message) (string, error) {

	message, err := s.Storage.AddMessage(messageFromHand)
	if err != nil {
		return "ошибка в сервисном уровне", err
	}

	return message, nil
}