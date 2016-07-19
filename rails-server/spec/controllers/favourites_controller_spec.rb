require 'rails_helper'

RSpec.describe FavouritesController, type: :controller do

  describe "POST #create" do
    context "when favourite does not exist" do
      it 'creates a favourite' do
        expect { post :create, favourite: {number: 1} }.to change(Favourite, :count).by(1)
      end
    end

    context "when favourite already exists" do
    end

    context "when given an invalid number" do
    end
  end

  describe "DELETE #destroy" do
    context "when favourite exists" do
      it 'deletes the favourite' do
        Favourite.create(number: 1)
        expect { delete :destroy, number: 1 }.to change(Favourite, :count).by(-1)
      end
    end

    context "when favourite does not exist" do
    end

    context "when given an invalid number" do
    end
  end
end
