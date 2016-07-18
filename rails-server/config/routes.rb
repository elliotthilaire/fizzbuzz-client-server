Rails.application.routes.draw do
  root 'fizzbuzz#index'

  get 'fizzbuzz/index'
  resources :favourites, only: [:create, :destroy]
end
