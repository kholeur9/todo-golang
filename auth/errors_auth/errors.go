package errors

import "errors"

var ErrInfoNotCompleted = errors.New("Veuillez renseignez tous les champs.")
var ErrNameNotTooLong = errors.New("Votre nom doit avoir plus de 1 caractère.")