module Api
  module V1
    class FavouritesController < ApplicationController

      def create
        @favourite = Favourite.new(number: params[:number])

        if @favourite.save
          render json: @favourite, status: :created
        else
          render json: @favourite.errors, status: :unprocessable_entity
        end
      end

      def destroy
        @favourite = Favourite.find_by(number: params[:number])

        if @favourite && @favourite.destroy
          @favourite.destroy
          head :no_content
        else
          head :not_found
        end
      end

    end
  end
end
