require 'rails_helper'

RSpec.describe FizzbuzzController, type: :controller do

  describe "GET #index" do
    context "with defaults" do
      it "returns http success" do
        get :index
        expect(response).to have_http_status(:success)
      end

      it "loads a page of fizzbuzz results" do
      end
    end

    context "with ?page=1" do
    end

    context "with ?page=10" do
    end

    context "with ?per_page=20" do
    end

    context "with invalid ?page" do
      it "returns http bad request" do
        get :index, params: { page: 'not_a_number' }
        expect(response).to have_http_status(:bad_request)
      end
    end

    context "with invalid ?per_page" do
      it "returns http bad request" do
        get :index, params: { per_page: 'not_a_number' }
        expect(response).to have_http_status(:bad_request)
      end
    end
  end

end
