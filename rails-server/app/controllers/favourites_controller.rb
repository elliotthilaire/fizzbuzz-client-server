class FavouritesController < ApplicationController

  def create
    @favourite = Favourite.new(number: params[:number])

    respond_to do |format|
      if @favourite.save
        format.html { redirect_to :back, notice: 'Favourite added.' }
        format.json { render json: @favourite, status: :created }
      else
        format.html { redirect_to :back, notice: @favourite.errors }
        format.json { render json: @favourite.errors, status: :unprocessable_entity }
      end
    end
  end

  def destroy
    @favourite = Favourite.find_by(number: params[:number])

    respond_to do |format|
      if @favourite.destroy
        format.html { redirect_to :back, notice: 'Favourite removed.' }
        format.json { head :no_content }
      else
        format.html { render 'what'}
        format.json { render json: @favourite.errors, status: :unprocessable_entity }
      end
    end
  end
end
