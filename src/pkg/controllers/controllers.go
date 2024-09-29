package controllers

import "SongsLibrary/src/pkg/repositories"

type Controller struct {
    Repository repositories.Repository
}

func NewController(repository repositories.Repository) *Controller {
    return &Controller{
        Repository: repository,
    }
}
