require 'rails_helper'

RSpec.describe FavouritesController, type: :controller do

  before { request.env["HTTP_REFERER"] = root_url }

  describe "POST #create" do
    context "when favourite does not exist" do
      it 'creates a favourite' do
        expect { post :create, params: {number: 1} }.to change(Favourite, :count).by(1)
      end

      it 'redirects to previous url' do
        pending
        expect(response).to redirect_to :back
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
        expect { delete :destroy, params: {number: 1} }.to change(Favourite, :count).by(-1)
      end

      it 'redirects to previous url' do
        pending
        expect(response).to redirect_to :back
      end
    end

    context "when favourite does not exist" do
    end

    context "when given an invalid number" do
    end
  end
end
