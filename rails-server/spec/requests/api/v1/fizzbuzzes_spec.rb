require "rails_helper"

RSpec.describe "Fizzbuzzes", :type => :request do

  let(:headers) do
    { "ACCEPT" => "application/json" }
  end

  describe 'listing fizzbuzzes' do
    it "lists fizzbuzzes" do
      get "/api/v1/fizzbuzzes", {}, headers
      expect(response.content_type).to eq("application/json")
      expect(response).to have_http_status(:ok)
    end
  end
end
