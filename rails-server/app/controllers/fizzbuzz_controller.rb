class FizzbuzzController < ApplicationController
  def index
    @fizzbuzzes = Page.new { |item_number| Fizzbuzz.new(item_number) }
  end
end
