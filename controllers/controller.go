package controllers

// Service interface del crud
type Service interface {
	Create()
	Update()
	Delete()
	FindAll()
	FindBy()
}
