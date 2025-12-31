package errors

import "errors"

var ErrTodoDonotEmpty = errors.New("La todo ne peut pas être vide.")
var ErrTodoHasUpdateEmpty =  errors.New("La todo a modifié doit contenir un texte.")
var ErrTodoNotFound = errors.New("La todo démandée n'existe pas.")
var ErrTodoIdEmpty = errors.New("La recherche n'a pas pu être traîté.")
var ErrTodoAlreadyTrue = errors.New("Une todo complétée ne peut plus être modifié.")