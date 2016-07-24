require "rails_helper"

RSpec.describe "Fizzbuzz", :type => :request do

  let(:headers) do
    {
      "Content-Type" => "application/json",
      "ACCEPT" => "application/json"
    }
  end

  context "with defaults" do
    it "returns http success" do
      get "/api/v1/fizzbuzzes", headers: headers
      expect(response.content_type).to eq("application/json")
      expect(response).to have_http_status(:ok)
    end
  end

  context "with page specified" do
    it 'returns http success' do
      get "/api/v1/fizzbuzzes?page=2", headers: headers
      expect(response).to have_http_status(:ok)
    end
  end

  context "with invalid page specified" do
    it "returns http bad request" do
      get "/api/v1/fizzbuzzes?page=not_a_number", headers: headers
      expect(response).to have_http_status(:bad_request)
    end
  end

  context "with per_page specified" do
    it "returns https success" do
      get "/api/v1/fizzbuzzes?per_page=20", headers: headers
      expect(response).to have_http_status(:success)
    end
  end

  context "with invalid per_page specified" do
    it "returns https bad_request" do
      get "/api/v1/fizzbuzzes?per_page=not_a_number", headers: headers
      expect(response).to have_http_status(:bad_request)
    end
  end
end
