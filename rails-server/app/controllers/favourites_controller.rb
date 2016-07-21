class FavouritesController < ApplicationController

  def create
    @favourite = Favourite.new(favourite_params)

    if @favourite.save
      redirect_back notice: 'Favourite added.', fallback_location: root_path
    else
      redirect_back notice: @favourite.errors, fallback_location: root_path
    end
  end

  def destroy
    @favourite = Favourite.find_by(favourite_params)

    if @favourite && @favourite.destroy
      redirect_back notice: 'Favourite removed.', fallback_location: root_path
    else
      redirect_back notice: 'Favourite not found', fallback_location: root_path
    end
  end

  private

  def favourite_params
    params.permit(:number)
  end
end
