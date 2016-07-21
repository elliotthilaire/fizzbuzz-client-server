require 'rails_helper'

RSpec.describe FizzbuzzesController, type: :controller do

  describe "GET #index" do
    context "with defaults" do
      it "returns http success" do
        get :index
        expect(response).to have_http_status(:success)
      end
    end

    context "with page specified" do
      it 'returns https success' do
        get :index, params: { page: 2 }
        expect(response).to have_http_status(:success)
      end
    end

    context "with invalid page specified" do
      it "returns http bad request" do
        get :index, params: { page: 'not_a_number' }
        expect(response).to have_http_status(:bad_request)
      end
    end

    context "with per_page specified" do
      it "returns https success" do
        get :index, params: { per_page: 'not_a_number' }
        expect(response).to have_http_status(:success)
      end
    end

    context "with invalid per_page specified" do
      it "returns https success" do
        get :index, params: { per_page: 'not_a_number' }
        expect(response).to have_http_status(:bad_request)
      end
    end
  end

end
