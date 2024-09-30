package repositories

import (
    "SongsLibrary/app/pkg/controllers"
)

type Repository interface {
    Create()
    Read()
    Update()
    Delete()
}

func NewRepository(repo Repository) *Repository {
    controller := controllers.NewController(repo)
    return &controller.Repository
}
