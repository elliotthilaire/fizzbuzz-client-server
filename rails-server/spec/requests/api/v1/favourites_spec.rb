require "rails_helper"

RSpec.describe "Favourites", :type => :request do

  let(:headers) do
    { "ACCEPT" => "application/json" }
  end

  describe 'creating a favourite' do
    it "creates a favourite" do
      post "/api/v1/favourites", { number: 1 }, headers
      expect(response).to have_http_status(:created)
    end
  end

  describe 'deleting a favourite' do
    it "deletes a favourite" do
      Favourite.create(number: 1)
      delete "/api/v1/favourites/1", {}, headers
      expect(response).to have_http_status(:no_content)
    end
  end
end

