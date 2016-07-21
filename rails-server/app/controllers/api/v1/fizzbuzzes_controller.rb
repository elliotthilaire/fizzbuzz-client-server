module Api
  module V1
    class FizzbuzzesController < ApplicationController

      def index
        page = params[:page] || 1
        per_page = params[:per_page] || 100

        @page_of_fizzbuzzes = Page.new(page_number: page, per_page: per_page) do |item_number|
          Fizzbuzz.new(item_number, favourites: favourites)
        end

        render json: @page_of_fizzbuzzes, status: :ok
      end

      private

      def favourites
        Favourite.all.map(&:number)
      end
    end
  end
end
