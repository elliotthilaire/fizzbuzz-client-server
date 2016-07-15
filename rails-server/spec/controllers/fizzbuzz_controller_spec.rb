require 'rails_helper'

RSpec.describe FizzbuzzController, type: :controller do

  describe "GET #index" do
    it "returns http success" do
      get :index
      expect(response).to have_http_status(:success)
    end

    context "with ?page=1" do

    end

    context "with ?page=10" do
    end

    context "with ?per_page=20" do
    end

end
