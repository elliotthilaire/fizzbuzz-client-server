Rails.application.routes.draw do
  root 'fizzbuzz#index'

  get 'fizzbuzz/index'
  post 'favourites/:number', to: 'favourites#create', as: :create_favourite
  delete 'favourites/:number', to: 'favourites#destroy', as: :delete_favourite

  namespace :api, defaults: {format: 'json'} do
    namespace :v1 do
      resources :fizzbuzzes, only: [:index]
      resources :favourites, param: :number, only: [:create, :destroy]
    end
  end
end
