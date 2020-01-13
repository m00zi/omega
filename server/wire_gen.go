// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"github.com/jinzhu/gorm"
	"rest-gin-gorm/pkg/invoice"
	"rest-gin-gorm/pkg/product"
	"rest-gin-gorm/pkg/user"
)

// Injectors from wire.go:

func InitProductAPI(db *gorm.DB) product.ProductAPI {
	productRepository := product.ProvideProductRepostiory(db)
	productService := product.ProvideProductService(productRepository)
	productAPI := product.ProvideProductAPI(productService)
	return productAPI
}

func InitInvoiceAPI(db *gorm.DB) invoice.InvoiceAPI {
	invoiceRepository := invoice.ProvideInvoiceRepostiory(db)
	invoiceService := invoice.ProvideInvoiceService(invoiceRepository)
	invoiceAPI := invoice.ProvideInvoiceAPI(invoiceService)
	return invoiceAPI
}

func InitUserAPI(db *gorm.DB) user.UserAPI {
	userRepository := user.ProvideUserRepostiory(db)
	userService := user.ProvideUserService(userRepository)
	userAPI := user.ProvideUserAPI(userService)
	return userAPI
}