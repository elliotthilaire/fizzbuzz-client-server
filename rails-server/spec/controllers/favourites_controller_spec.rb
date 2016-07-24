require 'rails_helper'

RSpec.describe FavouritesController, type: :controller do

  before { request.env["HTTP_REFERER"] = root_url }

  describe "POST #create" do
    context "when favourite does not exist" do
      it 'creates a favourite' do
        expect { post :create, params: {number: 1} }.to change(Favourite, :count).by(1)
      end

      it 'redirects to previous url' do
        post :create, params: {number: 1}
        expect(response).to redirect_to :back
      end
    end

    context "when favourite already exists" do
      it 'redirects to previous url' do
        Favourite.create(number: 1)
        post :create, params: {number: 1}
        expect(response).to redirect_to :back
      end
    end

    context "when given an invalid number" do
        it 'redirects to previous url' do
        post :create, params: {number: 'invalid_number'}
        expect(response).to redirect_to :back
      end
    end
  end

  describe "DELETE #destroy" do
    context "when favourite exists" do
      it 'deletes the favourite' do
        Favourite.create(number: 1)
        expect { delete :destroy, params: {number: 1} }.to change(Favourite, :count).by(-1)
      end

      it 'redirects to previous url' do
        delete :destroy, params: {number: 1}
        expect(response).to redirect_to :back
      end
    end

    context "when favourite does not exist" do
      it 'redirects to previous url' do
        Favourite.create(number: 1)
        delete :destroy, params: {number: 1}
        expect(response).to redirect_to :back
      end
    end

    context "when given an invalid number" do
      it 'redirects to previous url' do
        delete :destroy, params: {number: 'invalid_number'}
        expect(response).to redirect_to :back
      end
    end

  end
end
