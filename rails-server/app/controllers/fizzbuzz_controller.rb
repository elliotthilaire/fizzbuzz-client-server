class FizzbuzzController < ApplicationController
  def index
    page = params[:page] || 1
    per_page = params[:per_page] || 100

    @fizzbuzzes = Page.new(page: page, per_page: per_page) do |item_number|
      Fizzbuzz.new(item_number)
    end
  end
end
