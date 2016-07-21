Rails.application.routes.draw do
  root 'fizzbuzzes#index'

  resources :fizzbuzzes, only: [:index]
  resources :favourites, param: :number, only: [:create, :destroy]

  namespace :api, defaults: {format: 'json'} do
    namespace :v1 do
      resources :fizzbuzzes, only: [:index]
      resources :favourites, param: :number, only: [:create, :destroy]
    end
  end
end
