module Api
  module V1
    class FavouritesController < ApplicationController

      def create
        @favourite = Favourite.new(favourite_params)

        if @favourite.save
          render json: @favourite, status: :created
        else
          render json: @favourite.errors, status: :unprocessable_entity
        end
      end

      def destroy
        @favourite = Favourite.find_by(favourite_params)

        if @favourite && @favourite.destroy
          head :no_content
        else
          head :not_found
        end
      end

      def favourite_params
        params.permit(:number)
      end
    end
  end
end
