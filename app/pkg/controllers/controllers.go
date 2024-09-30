package controllers

import (
    "SongsLibrary/app/pkg/repositories"
)

type Controller struct {
    Repository repositories.Repository
}

func NewController(repository repositories.Repository) *Controller {
    return &Controller{
        Repository: repository,
    }
}
