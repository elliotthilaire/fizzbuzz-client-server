require "rails_helper"

RSpec.describe "Favourites", :type => :request do

  let(:headers) do
    {
      "Content-Type" => "application/json",
      "ACCEPT" => "application/json"
    }
  end

  describe 'creating a favourite' do
    context "with valid params" do
      it "returns http status :created" do
        post "/api/v1/favourites", params: {number: 1}, headers: headers, as: :json
        expect(response).to have_http_status(:created)
      end
    end

    context "with invalid params" do
      it "returns http status :unprocessable_entity" do
        post "/api/v1/favourites", params: {number: 'not_a_number'}, headers: headers, as: :json
        expect(response).to have_http_status(:unprocessable_entity)
      end
    end
  end

  describe 'deleting a favourite' do
    context "when it exists" do
      it "returns http status :no_content" do
        Favourite.create(number: 1)
        delete "/api/v1/favourites/1", headers: headers
        expect(response).to have_http_status(:no_content)
      end
    end

    context "when it does not exist" do
      it "returns http status :not_found" do
        delete "/api/v1/favourites/1", headers: headers
        expect(response).to have_http_status(:not_found)
      end
    end
  end
end

