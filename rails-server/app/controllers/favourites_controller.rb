class FavouritesController < ApplicationController

  def create
    @favourite = Favourite.new(number: params[:number])

    respond_to do |format|
      if @favourite.save
        format.html { redirect_to :back, notice: 'Favourite was successfully created.' }
        format.json { render :show, status: :created, location: @favourite }
      else
        format.html { }
        format.json { render json: @favourite.errors, status: :unprocessable_entity }
      end
    end
  end

  def destroy
    @favourite = Favourite.find_by(number: params[:number])

    @favourite.destroy
    respond_to do |format|
      format.html { redirect_to :back, notice: 'Favourite was successfully destroyed.' }
      format.json { head :no_content }
    end
  end

end
