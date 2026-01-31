package errors

import "errors"

var ErrInfoNotCompleted = errors.New("Veuillez renseignez tous les champs.")
var ErrNameNotTooLong = errors.New("Votre nom doit avoir plus de 1 caractère.")
var ErrPasswordTooShort = errors.New("Le mot de passe doit contenir au moins 8 caractères.")
var ErrEmailAlreadyExists =  errors.New("Cet adresse email est déjà utilisée.")
var ErrEmailNotFound = errors.New("Adresse email n'existe pas.")
var ErrIncorrectCredentials = errors.New("Email ou mot de passe invalides.")